package infrastructure

import (
	"net/http"

	interfaces "collegare/application/services"
	"collegare/infrastructure/config"
	"collegare/infrastructure/controllers"
	"collegare/infrastructure/services"

	"github.com/gorilla/websocket"
)

// Configuration
var AppConfig *config.AppConfiguration = config.ParseConfiguration()

// Persistence

// Services
var sessionCache interfaces.SessionCache = services.NewSessionCache()
var callHandler interfaces.CallHandler = services.NewWebRTCService(
	AppConfig.WebRTC.ICEServers,
	&sessionCache,
)
var commandHandler interfaces.CommandHandler = services.NewDefaultCommandHandler(callHandler)
var sessionService interfaces.SessionService = services.NewWebSocketService(websocket.Upgrader{
	ReadBufferSize:  AppConfig.Session.WebSockets.ReadBufferSize,
	WriteBufferSize: AppConfig.Session.WebSockets.WriteBufferSize,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}, commandHandler)

// Controllers
var sessionController *controllers.SessionController = controllers.NewSessionController(sessionService)
var healthController *controllers.HealthController = controllers.NewHealthController()
