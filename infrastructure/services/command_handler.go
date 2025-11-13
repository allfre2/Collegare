package services

import (
	"collegare/domain"
	"collegare/domain/constants"
	infra "collegare/infrastructure/constants"
	"collegare/infrastructure/util"
)

type DefaultCommandHandler struct {
}

func NewDefaultCommandHandler() *DefaultCommandHandler {
	return &DefaultCommandHandler{}
}

func (h *DefaultCommandHandler) Parse(command string) (any, error) {
	cmd, err := util.Deserialize[domain.Command](command)
	if err != nil {
		return nil, err
	}

	var result any

	switch cmd.Type {
	case constants.Connect:
		result, err = util.Deserialize[domain.ConnectCommand](command)
	case constants.Login:
		result, err = util.Deserialize[domain.LoginCommand](command)
	case constants.Call:
		result, err = util.Deserialize[domain.CallCommand](command)
	case constants.HangUp:
		result, err = util.Deserialize[domain.HangUpCommand](command)
	case constants.Digit:
		result, err = util.Deserialize[domain.DigitCommand](command)
	case constants.Disconnect:
		result, err = util.Deserialize[domain.DisconnectCommand](command)
	default:
		return nil, infra.ErrInvalidCommand
	}

	return result, err
}

func (h *DefaultCommandHandler) Handle(commandStr string) (string, error) {

	command, err := h.Parse(commandStr)
	if err != nil {
		return "", err
	}

	var response string

	switch cmd := command.(type) {
	case *domain.ConnectCommand:
		response, _ = util.Serialize(cmd)
	case *domain.LoginCommand:
		response, _ = util.Serialize(cmd)
	case *domain.CallCommand:
		response, _ = util.Serialize(cmd)
	case *domain.HangUpCommand:
		response, _ = util.Serialize(cmd)
	case *domain.DigitCommand:
		response, _ = util.Serialize(cmd)
	case *domain.DisconnectCommand:
		response, _ = util.Serialize(cmd)
	default:
		response = infra.ErrInvalidCommand.Error()
	}

	return "=> Received: \n" + response + "\n", nil
}
