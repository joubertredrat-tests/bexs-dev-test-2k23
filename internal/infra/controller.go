package infra

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ApiBaseController struct {
}

func NewApiBaseController() ApiBaseController {
	return ApiBaseController{}
}

func (c ApiBaseController) HandleStatus(ctx *gin.Context) {
	t := time.Now()
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"time":   DatetimeCanonical(&t),
	})
}

func (c ApiBaseController) HandleNotFound(ctx *gin.Context) {
	t := time.Now()
	ctx.JSON(http.StatusNotFound, gin.H{
		"error": "page not found",
		"time":  DatetimeCanonical(&t),
	})
}
