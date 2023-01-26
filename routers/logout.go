package routers

import (
	"github.com/gin-gonic/gin"
	"misso/handlers/logout"
)

func Logout(rg *gin.RouterGroup) {
	rg.GET("/logout", logout.Logout)
}
