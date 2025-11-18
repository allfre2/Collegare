package services

import (
	"collegare/application/services"
	"collegare/domain"
	"collegare/infrastructure/errors"
	"collegare/infrastructure/util"
	"fmt"
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

	var result any

	switch cmd.Type {
	case domain.Connect:
		result, err = util.Deserialize[domain.ConnectCommand](message)
	case domain.Login:
		result, err = util.Deserialize[domain.LoginCommand](message)
	case domain.Call:
		result, err = util.Deserialize[domain.CallCommand](message)
	case domain.HangUp:
		result, err = util.Deserialize[domain.HangUpCommand](message)
	case domain.Digit:
		result, err = util.Deserialize[domain.DigitCommand](message)
	case domain.Disconnect:
		result, err = util.Deserialize[domain.DisconnectCommand](message)
	default:
		return nil, errors.ErrInvalidCommand
	}

	return result, err
}

func (h *DefaultCommandHandler) Handle(message string, session *domain.Session) (string, error) {

	command, err := h.Parse(message)
	if err != nil {
		fmt.Println("Error Parsing command.", err)
		return "", err
	}

	var response string

	switch cmd := command.(type) {
	case *domain.ConnectCommand:
		response, _ = util.Serialize(cmd)
	case *domain.LoginCommand:
		response, _ = util.Serialize(cmd)
	case *domain.CallCommand:
		if cmd.Mode == domain.AdHoc {
			fmt.Println("matched callCommand type.")
			continuation, err := h.callHandler.Route(cmd, session)
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
		response, _ = util.Serialize(cmd)
	case *domain.HangUpCommand:
		response, _ = util.Serialize(cmd)
	case *domain.DigitCommand:
		response, _ = util.Serialize(cmd)
	case *domain.DisconnectCommand:
		return cmd.Type, nil
	default:
		response = errors.ErrInvalidCommand.Error()
	}

	return "=> Received: \n" + response + "\n", nil
}

func (h *DefaultCommandHandler) GetCallConfiguration() (*domain.CallConfiguration, error) {
	return h.callHandler.GetConfiguration(), nil
}
