package infrastructure

import (
	"net/http"

	repositories "collegare/application/persistence"
	interfaces "collegare/application/services"
	"collegare/infrastructure/config"
	"collegare/infrastructure/controllers"
	"collegare/infrastructure/persistence"
	"collegare/infrastructure/services"

	"github.com/gorilla/websocket"
)

// Configuration
var AppConfig *config.AppConfiguration = config.ParseConfiguration()

// Persistence
var callFlowRepository repositories.CallFlowRepository = persistence.NewMockCallFlowRepository()

// Services
var fileCache interfaces.FileCache = services.NewMockFileCache()
var sessionCache interfaces.SessionCache = services.NewMockSessionCache()
var callHandler interfaces.CallHandler = services.NewWebRTCService(
	AppConfig.WebRTC.ICEServers,
	callFlowRepository,
	fileCache,
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
