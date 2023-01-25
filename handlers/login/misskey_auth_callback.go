package login

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	client "github.com/ory/hydra-client-go/v2"
	"misso/config"
	"misso/consts"
	"misso/global"
	"misso/misskey"
	"misso/types"
	"net/http"
	"time"
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

	userid := fmt.Sprintf("%s@%s", usermeta.User.Username, config.Config.Misskey.Instance)

	// Save context into redis
	userinfoCtxBytes, err := json.Marshal(&types.SessionContext{
		MisskeyToken: usermeta.AccessToken,
		User:         usermeta.User,
	})
	if err != nil {
		global.Logger.Errorf("Failed to parse accept context with error: %v", err)
		ctx.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
			"error": "Failed to parse accept context",
		})
		return
	}
	sessKey = fmt.Sprintf(consts.REDIS_KEY_SHARE_CONTEXT, userid)
	err = global.Redis.Set(context.Background(), sessKey, userinfoCtxBytes, consts.TIME_LOGIN_SESSION_VALID).Err()
	if err != nil {
		global.Logger.Errorf("Failed to save session into redis with error: %v", err)
		ctx.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
			"error": "Failed to save context",
		})
		return
	}

	global.Logger.Debugf("User accepted the request, reporting back to hydra...")
	remember := true
	rememberFor := int64(consts.TIME_LOGIN_SESSION_VALID / time.Second)
	acceptReq, _, err := global.Hydra.Admin.OAuth2Api.AcceptOAuth2LoginRequest(context.Background()).LoginChallenge(oauth2challenge).AcceptOAuth2LoginRequest(client.AcceptOAuth2LoginRequest{
		Subject:     userid,
		Remember:    &remember,
		RememberFor: &rememberFor,
	}).Execute()
	if err != nil {
		global.Logger.Errorf("Failed to accept login request with error: %v", err)
		ctx.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
			"error": "Failed to accept login request",
		})
		return
	}

	// Redirect to target uri
	ctx.Redirect(http.StatusTemporaryRedirect, acceptReq.RedirectTo)

	global.Logger.Debugf("User should now be redirecting to target URI.")

}
