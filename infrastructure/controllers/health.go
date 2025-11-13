package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (h HealthController) Status(ctx *gin.Context) {
	ctx.String(http.StatusOK, "OK")
}
