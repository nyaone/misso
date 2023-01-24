package routers

import (
	"github.com/gin-gonic/gin"
	"misso/handlers/consent"
)

func Consent(rg *gin.RouterGroup) {
	rg.GET("/consent", consent.ConsentCheck)
	rg.POST("/consent", consent.ConsentConfirm)
}
