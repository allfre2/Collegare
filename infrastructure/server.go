package infrastructure

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func NewRouter() *gin.Engine {
	if AppConfig.Mode == gin.DebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(cors.AllowAll())

	router.RedirectTrailingSlash = true
	router.UseH2C = true

	return router
}

func Start(server *gin.Engine) {
	server.Run(AppConfig.Server.Port)
}
