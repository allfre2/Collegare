package infrastructure

import (
	"github.com/gin-gonic/gin"
)

func SetupV1(router *gin.Engine) {
	api := router.Group("api")
	{
		v1 := api.Group("v1")
		{
			sessionRouter := v1.Group("session")
			{
				sessionRouter.GET("", sessionController.Session)
			}
		}
	}

	router.GET("health", healthController.Status)
}
