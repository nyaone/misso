package user

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"misso/consts"
	"misso/global"
	"misso/types"
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

	// Retrieve context
	global.Logger.Debugf("Retrieving context...")
	var acceptCtx types.SessionContext
	sessKey := fmt.Sprintf(consts.REDIS_KEY_SHARE_CONTEXT, *tokenInfo.Sub)
	acceptCtxBytes, err := global.Redis.Get(context.Background(), sessKey).Bytes()
	if err != nil {
		global.Logger.Errorf("Failed to retrieve context with error: %v", err)
		ctx.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
			"error": "Failed to retrieve context",
		})
		return
	}

	global.Logger.Debugf("Decoding context...")
	err = json.Unmarshal(acceptCtxBytes, &acceptCtx)
	if err != nil {
		global.Logger.Errorf("Failed to parse context with error: %v", err)
		ctx.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
			"error": "Failed to parse context",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"email": tokenInfo.Sub,
		"user":  acceptCtx.User,
	})

}
