package login

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	client "github.com/ory/hydra-client-go/v2"
	"misso/consts"
	"misso/global"
	"misso/misskey"
	"net/http"
)

func Login(ctx *gin.Context) {
	oauth2challenge := ctx.Query("login_challenge") // OAuth2 login
	if oauth2challenge == "" {
		ctx.HTML(http.StatusBadRequest, "error.tmpl", gin.H{
			"error": "Necessary challenge not provided",
		})
		return
	}

	// OAuth2 Login
	// Get request with oauth2challenge
	global.Logger.Debugf("Getting OAuth2 Login Request...")
	loginReq, _, err := global.Hydra.Admin.OAuth2Api.GetOAuth2LoginRequest(context.Background()).LoginChallenge(oauth2challenge).Execute()
	if err != nil {
		global.Logger.Errorf("Failed to get required login request with error: %v", err)
		ctx.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
			"error": "Failed to get required login request",
		})
		return
	}

	// Check if is able to skip
	if loginReq.Skip {
		global.Logger.Debugf("Login is able to skip, skipping...")
		// Skip auth redirect, simply go to final
		acceptReq, _, err := global.Hydra.Admin.OAuth2Api.AcceptOAuth2LoginRequest(context.Background()).LoginChallenge(oauth2challenge).AcceptOAuth2LoginRequest(client.AcceptOAuth2LoginRequest{
			Context: context.Background(),
			Subject: loginReq.Subject,
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
	} else {
		global.Logger.Debugf("Login is NOT able to skip, redirecting to Misskey for authentication...")
		// Generate auth session
		global.Logger.Debugf("Generating auth session...")
		authSess, err := misskey.GenerateAuthSession()
		if err != nil {
			global.Logger.Errorf("Failed to generate misskey auth session with error: %v", err)
			ctx.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
				"error": "Failed to generate misskey auth session",
			})
			return
		}

		global.Logger.Debugf("Grabbed auth session: %s", authSess.Token)

		// Save login challenge state into redis (misskey cannot keep state info)
		sessKey := fmt.Sprintf(consts.REDIS_KEY_LOGIN_SESSION, authSess.Token)
		err = global.Redis.Set(context.Background(), sessKey, oauth2challenge, consts.TIME_REQUEST_VALID).Err()
		if err != nil {
			global.Logger.Errorf("Failed to save session into redis with error: %v", err)
			ctx.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
				"error": "Failed to save session",
			})
			return
		}

		// Redirect to misskey auth
		ctx.Redirect(http.StatusTemporaryRedirect, authSess.Url)

		global.Logger.Debugf("User should now be redirecting to target URI.")
	}
}
