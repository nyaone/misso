package public

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthCheck(ctx *gin.Context) {
	ctx.String(http.StatusOK, "OK")
}
