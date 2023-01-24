package routers

import (
	"github.com/gin-gonic/gin"
	"misso/handlers/user"
)

func User(rg *gin.RouterGroup) {
	rg.GET("/userinfo", user.UserInfo)
}
