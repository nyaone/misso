package routers

import (
	"github.com/gin-gonic/gin"
	"misso/handlers/public"
)

func Public(rg *gin.RouterGroup) {
	rg.GET("/", public.Index)
	rg.GET("/healthcheck", public.HealthCheck)
}
