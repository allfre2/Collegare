package services

import (
	"collegare/application/services"
	"collegare/domain"
	"collegare/infrastructure/errors"
	"collegare/infrastructure/util"
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
		return nil, errors.ErrInvalidContext
	}

	connection, err := ws.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return nil, err
	}

	session := domain.NewSession()

	go ws.routingMessages(connection, session)
	go ws.spinSession(connection, session)

	return session, nil
}

func (ws *WebSocketService) Disconnect(session *domain.Session) error {
	return nil
}

func (ws *WebSocketService) GetSessionConfiguration() (*domain.ConfigurationMessage, error) {
	callConfig, err := ws.cmdHandler.GetCallConfiguration()
	if err != nil {
		return nil, err
	}

	message := domain.ConfigurationMessage{}
	message.Type = domain.SessionConfig
	message.CallConfiguration = *callConfig

	return &message, nil
}

func (ws *WebSocketService) spinSession(connection *websocket.Conn, session *domain.Session) {
	connection.SetCloseHandler(websocketCloseHandler(session))
	defer connection.Close()

	if err := ws.sendConfigMessage(connection); err != nil {
		fmt.Println("Couldn't send configuration message")
		return // (!)
	}

	for {
		messageType, message, err := connection.ReadMessage()
		if err != nil {
			fmt.Println("error receiving command: " + err.Error() + ", Id: " + session.Id.String())
			break //(!)
		}

		fmt.Printf("\nReceived: \n%s\n", message)

		response, err := ws.cmdHandler.Handle(string(message), session)
		if err != nil {
			fmt.Println("error handling command")
		}

		if response == "" {
			continue
		}

		err = connection.WriteMessage(messageType, []byte(response))
		if err != nil {
			fmt.Println("error sending response")
		}
	}
}

func (ws *WebSocketService) routingMessages(connection *websocket.Conn, session *domain.Session) {
	for msg := range session.Outbound {
		fmt.Println("outbound channel: ", msg)
		message, _ := util.Serialize(msg)
		connection.WriteMessage(websocket.TextMessage, []byte(message))
	}
}

func websocketCloseHandler(session *domain.Session) func(code int, text string) error {
	return func(code int, text string) error {
		fmt.Println("Closing connection: " + session.Id.String())
		return nil
	}
}

func (ws *WebSocketService) sendConfigMessage(connection *websocket.Conn) error {
	config, err := ws.GetSessionConfiguration()
	if err != nil {
		return err
	}

	message, err := util.Serialize(config)
	if err != nil {
		return err
	}

	if err := connection.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		fmt.Println("couldn't send session configuration")
		return err
	}

	return nil
}
