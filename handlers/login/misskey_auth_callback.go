package login

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	client "github.com/ory/hydra-client-go/v2"
	"misso/config"
	"misso/consts"
	"misso/global"
	"misso/misskey"
	"net/http"
)

func MisskeyAuthCallback(ctx *gin.Context) {
	// Misskey Callback login
	misskeySession := ctx.Query("token")
	if misskeySession == "" {
		ctx.HTML(http.StatusBadRequest, "error.tmpl", gin.H{
			"error": "Necessary token not provided",
		})
		return
	}

	// Get saved login challenge from redis
	global.Logger.Debugf("Getting saved login challenge...")
	sessKey := fmt.Sprintf(consts.REDIS_KEY_LOGIN_SESSION, misskeySession)
	oauth2challenge, err := global.Redis.Get(context.Background(), sessKey).Result()
	if err != nil {
		global.Logger.Errorf("Failed to get session from redis with error: %v", err)
		ctx.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
			"error": "Failed to get session",
		})
		return
	}

	// Delete used challenge
	global.Redis.Del(context.Background(), sessKey)

	// Check if session token is valid
	global.Logger.Debugf("Checking if session is valid...")
	usermeta, err := misskey.GetUserkey(misskeySession)
	if err != nil {
		global.Logger.Errorf("Failed to verify session from misskey with error: %v", err)
		ctx.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
			"error": "Failed to verify session (pending or invalid)",
		})
		return
	}

	// Accept user login
	// Get request with oauth2challenge
	global.Logger.Debugf("Getting OAuth2 Consent Request...")
	_, _, err = global.Hydra.Admin.OAuth2Api.GetOAuth2LoginRequest(context.Background()).LoginChallenge(oauth2challenge).Execute()
	if err != nil {
		global.Logger.Errorf("Failed to get required login request with error: %v", err)
		ctx.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
			"error": "Failed to get required login request",
		})
		return
	}

	// Save user key
	userIdentifier := fmt.Sprintf("%s@%s", usermeta.User.Username, config.Config.Misskey.Instance)

	sessAccessTokenKey := fmt.Sprintf(consts.REDIS_KEY_USER_ACCESS_TOKEN, userIdentifier)
	err = global.Redis.Set(context.Background(), sessAccessTokenKey, usermeta.AccessToken, 0).Err()
	if err != nil {
		global.Logger.Errorf("Failed to save session access token into redis with error: %v", err)
		ctx.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
			"error": "Failed to save context",
		})
		return
	}

	global.Logger.Debugf("User accepted the request, reporting back to hydra...")
	acceptReq := client.AcceptOAuth2LoginRequest{
		Subject: userIdentifier,
	}
	if config.Config.Time.LoginRemember > 0 {
		remember := true
		acceptReq.Remember = &remember
		acceptReq.RememberFor = &config.Config.Time.LoginRemember
	}
	acceptRes, _, err := global.Hydra.Admin.OAuth2Api.AcceptOAuth2LoginRequest(context.Background()).LoginChallenge(oauth2challenge).AcceptOAuth2LoginRequest(acceptReq).Execute()
	if err != nil {
		global.Logger.Errorf("Failed to accept login request with error: %v", err)
		ctx.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
			"error": "Failed to accept login request",
		})
		return
	}

	// Redirect to target uri
	ctx.Redirect(http.StatusTemporaryRedirect, acceptRes.RedirectTo)

	global.Logger.Debugf("User should now be redirecting to target URI.")

}
