package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"misso/global"
	"misso/utils"
	"net/http"
	"strings"
)

func UserInfo(ctx *gin.Context) {
	// Get token from header
	accessToken := strings.Replace(ctx.GetHeader("Authorization"), "Bearer ", "", 1)
	if accessToken == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "No authorization token found",
		})
		return
	}

	// Retrieve token info
	global.Logger.Debugf("Retrieving access token info...")
	tokenInfo, _, err := global.Hydra.Admin.OAuth2Api.IntrospectOAuth2Token(context.Background()).Token(accessToken).Execute()
	if err != nil {
		global.Logger.Errorf("Failed to retrieve access token info with error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve access token info",
		})
		return
	}

	if !tokenInfo.Active {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": "This token is inactive",
		})
		return
	}

	// Return user info
	global.Logger.Debugf("Retrieving context...")
	userinfo, err := utils.GetUserinfo(*tokenInfo.Sub)
	if err != nil {
		global.Logger.Errorf("Failed to retrieve userinfo with error: %v", err)
		ctx.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
			"error": "Failed to get userinfo",
		})
		return
	}

	userinfoRes := gin.H{} // map[string]interface{}

	// Get scopes
	if tokenInfo.Scope != nil && *tokenInfo.Scope != "" {
		// Has scopes
		scopes := strings.Split(*tokenInfo.Scope, " ")
		for _, s := range scopes {
			if value, ok := (*userinfo)[s]; ok {
				userinfoRes[s] = value
			}
		}
	}

	ctx.JSON(http.StatusOK, userinfoRes)

}
