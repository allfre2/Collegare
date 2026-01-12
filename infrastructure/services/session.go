package services

import (
	"collegare/application/services"
	"collegare/domain"
	"collegare/infrastructure/errors"
	"collegare/infrastructure/util"
	"fmt"
	"time"

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

	ws.spinSession(connection, session)

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

	go ws.inboundMessages(connection, session)
	go ws.outboundMessages(connection, session)

	if err := ws.sendConfigMessage(session); err != nil {
		fmt.Println("Couldn't send configuration message")
		connection.Close()
		session.Close()
		return
	}
}

func (ws *WebSocketService) inboundMessages(connection *websocket.Conn, session *domain.Session) {
	defer connection.Close()

	for {
		_, message, err := connection.ReadMessage()
		if err != nil {
			fmt.Println("error receiving command: " + err.Error() + ", Id: " + session.Id.String())
			break //(!)
		}

		fmt.Printf("\nReceived: \n%s\n", message)

		if err := ws.cmdHandler.Handle(string(message), session); err != nil {
			fmt.Println("error handling command")
		}

		if !session.IsActive {
			fmt.Println("Closing connection: " + session.Id.String())
			session.Close()
			closeWebsocket(connection, 0, domain.Disconnect)
			return
		}
	}
}

func (ws *WebSocketService) outboundMessages(connection *websocket.Conn, session *domain.Session) {
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

func closeWebsocket(connection *websocket.Conn, code int, reason string) {
	message := websocket.FormatCloseMessage(code, reason)

	connection.SetWriteDeadline(time.Now().Add(5 * time.Second))

	// Send the close message (Close control frame)
	if err := connection.WriteMessage(websocket.CloseMessage, message); err != nil {
		fmt.Println("websocket close message error:", err)
	}
}

func (ws *WebSocketService) sendConfigMessage(session *domain.Session) error {
	config, err := ws.GetSessionConfiguration()
	if err != nil {
		return err
	}

	session.Outbound <- config

	return nil
}
