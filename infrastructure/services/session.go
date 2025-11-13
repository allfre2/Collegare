package services

import (
	"collegare/application/services"
	"collegare/domain"
	"collegare/infrastructure/constants"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebSocketService struct {
	upgrader   websocket.Upgrader
	cmdHandler services.CommandHandler
}

func NewWebSocketService(upgrader websocket.Upgrader, cmdHandler services.CommandHandler) *WebSocketService {
	return &WebSocketService{
		upgrader:   upgrader,
		cmdHandler: cmdHandler,
	}
}

func (ws *WebSocketService) Connect(request *domain.SessionRequest) (*domain.Session, error) {
	ctx, ok := request.Context.(*gin.Context)
	if !ok {
		return nil, constants.ErrInvalidContext
	}

	connection, err := ws.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return nil, err
	}

	session := domain.NewSession()

	go ws.SpinSession(connection, session)

	return session, nil
}

func (ws *WebSocketService) Disconnect(session *domain.Session) error {
	return nil
}

func (ws *WebSocketService) SpinSession(connection *websocket.Conn, session *domain.Session) {
	connection.SetCloseHandler(closeHandler(session))
	defer connection.Close()

	for {
		messageType, message, err := connection.ReadMessage()
		if err != nil {
			fmt.Println("error receiving command: " + err.Error() + ", Id: " + session.Id.String())
			break //(!)
		}

		fmt.Printf("\nReceived: \n%s\n", message)

		response, err := ws.cmdHandler.Handle(string(message))
		if err != nil {
			fmt.Println("error handling command")
		}

		err = connection.WriteMessage(messageType, []byte(response))
		if err != nil {
			fmt.Println("error sending response")
		}
	}
}

func closeHandler(session *domain.Session) func(code int, text string) error {
	return func(code int, text string) error {
		fmt.Println("Closing connection: " + session.Id.String())
		return nil
	}
}
