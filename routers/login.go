package routers

import (
	"github.com/gin-gonic/gin"
	"misso/handlers/login"
)

func Login(rg *gin.RouterGroup) {
	rg.GET("/login", login.Login)
	rg.GET("/misskey", login.MisskeyAuthCallback)
}
