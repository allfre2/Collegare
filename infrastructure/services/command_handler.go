package services

import (
	"collegare/application/services"
	"collegare/domain"
	"collegare/infrastructure/errors"
	"collegare/infrastructure/util"
	"fmt"
	"reflect"
)

type DefaultCommandHandler struct {
	callHandler services.CallHandler
}

func NewDefaultCommandHandler(callHandler services.CallHandler) *DefaultCommandHandler {
	return &DefaultCommandHandler{
		callHandler: callHandler,
	}
}

func (h *DefaultCommandHandler) Parse(message string) (any, error) {
	cmd, err := util.Deserialize[domain.SessionMessage](message)
	if err != nil {
		return nil, err
	}

	if parse, ok := GetParser(cmd.Type); ok {
		return parse(message)
	}

	return nil, errors.ErrInvalidCommand
}

func (h *DefaultCommandHandler) Handle(message string, session *domain.Session) (string, error) {

	command, err := h.Parse(message)
	if err != nil {
		fmt.Println("Error Parsing command.", err)
		return "", err
	}

	if handler, ok := GetHandler(command); ok {
		return handler.Handle(command, session, h)
	}

	return errors.ErrInvalidCommand.Error(), nil // Fix this
}

func (h *DefaultCommandHandler) GetCallConfiguration() (*domain.CallConfiguration, error) {
	return h.callHandler.GetConfiguration(), nil
}

// Parsers
var Parsers map[string]func(string) (any, error) = map[string]func(string) (any, error){
	domain.Connect:    func(message string) (any, error) { return util.Deserialize[domain.ConnectCommand](message) },
	domain.Login:      func(message string) (any, error) { return util.Deserialize[domain.LoginCommand](message) },
	domain.Call:       func(message string) (any, error) { return util.Deserialize[domain.CallCommand](message) },
	domain.HangUp:     func(message string) (any, error) { return util.Deserialize[domain.HangUpCommand](message) },
	domain.Digit:      func(message string) (any, error) { return util.Deserialize[domain.DigitCommand](message) },
	domain.Disconnect: func(message string) (any, error) { return util.Deserialize[domain.DisconnectCommand](message) },
}

func GetParser(cmdType string) (func(string) (any, error), bool) {
	parser, ok := Parsers[cmdType]
	return parser, ok
}

// Handlers
type Handler interface {
	Handle(cmd any, session *domain.Session, h *DefaultCommandHandler) (string, error)
}

type ConnectHandler struct{}

func (ConnectHandler) Handle(cmd any, session *domain.Session, h *DefaultCommandHandler) (string, error) {
	response, _ := util.Serialize(cmd.(*domain.ConnectCommand))
	return "=> Received: \n" + response + "\n", nil
}

type LoginHandler struct{}

func (LoginHandler) Handle(cmd any, session *domain.Session, h *DefaultCommandHandler) (string, error) {
	response, _ := util.Serialize(cmd.(*domain.LoginCommand))
	return "=> Received: \n" + response + "\n", nil
}

type CallHandler struct{}

func (CallHandler) Handle(cmd any, session *domain.Session, h *DefaultCommandHandler) (string, error) {
	call := cmd.(*domain.CallCommand)
	if call.Mode == domain.AdHoc {
		fmt.Println("matched callCommand type.")
		continuation, err := h.callHandler.Route(call, session)
		if err != nil {
			fmt.Println("call handler Route error")
			return "", err
		}

		if continuation != nil {
			response, _ := util.Serialize(continuation)
			return response, nil
		}

		return "", nil
	}

	// Other types pending to handle
	response, _ := util.Serialize(call)
	return "=> Received: \n" + response + "\n", nil
}

type HangupHandler struct{}

func (HangupHandler) Handle(cmd any, session *domain.Session, h *DefaultCommandHandler) (string, error) {
	response, _ := util.Serialize(cmd.(*domain.HangUpCommand))
	return "=> Received: \n" + response + "\n", nil
}

type DigitHandler struct{}

func (DigitHandler) Handle(cmd any, session *domain.Session, h *DefaultCommandHandler) (string, error) {
	response, _ := util.Serialize(cmd.(*domain.DigitCommand))
	return "=> Received: \n" + response + "\n", nil
}

type DisconnectHandler struct{}

func (DisconnectHandler) Handle(cmd any, session *domain.Session, h *DefaultCommandHandler) (string, error) {
	disconnect := cmd.(*domain.DisconnectCommand)
	return disconnect.Type, nil
}

var Handlers map[reflect.Type]Handler = map[reflect.Type]Handler{
	reflect.TypeOf(&domain.ConnectCommand{}):    ConnectHandler{},
	reflect.TypeOf(&domain.LoginCommand{}):      LoginHandler{},
	reflect.TypeOf(&domain.CallCommand{}):       CallHandler{},
	reflect.TypeOf(&domain.HangUpCommand{}):     HangupHandler{},
	reflect.TypeOf(&domain.DigitCommand{}):      DigitHandler{},
	reflect.TypeOf(&domain.DisconnectCommand{}): DisconnectHandler{},
}

func GetHandler(cmd any) (Handler, bool) {
	handler, ok := Handlers[reflect.TypeOf(cmd)]
	return handler, ok
}
