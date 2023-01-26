package logout

import (
	"context"
	"github.com/gin-gonic/gin"
	"misso/global"
	"net/http"
)

func Logout(ctx *gin.Context) {
	oauth2challenge := ctx.Query("logout_challenge") // OAuth2 login
	if oauth2challenge == "" {
		ctx.HTML(http.StatusBadRequest, "error.tmpl", gin.H{
			"error": "Necessary challenge not provided",
		})
		return
	}

	acceptReq, _, err := global.Hydra.Admin.OAuth2Api.AcceptOAuth2LogoutRequest(context.Background()).LogoutChallenge(oauth2challenge).Execute()
	if err != nil {
		global.Logger.Errorf("Failed to accept logout request with error: %v", err)
		ctx.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
			"error": "Failed to accept logout request",
		})
		return
	}

	ctx.Redirect(http.StatusTemporaryRedirect, acceptReq.RedirectTo)

}
