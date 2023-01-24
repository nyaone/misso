package public

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.tmpl", nil)
}
