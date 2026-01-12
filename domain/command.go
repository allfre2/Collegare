package domain

import "github.com/gofrs/uuid/v5"

// Command Types
const (
	Connect    = "connect"
	Login      = "login"
	Call       = "call"
	HangUp     = "hangup"
	Digit      = "digit"
	Disconnect = "disconnect"
)

type ConnectCommand struct {
	SessionMessage
}

type LoginCommand struct {
	SessionMessage
}

type CallCommand struct {
	SessionMessage
	Mode       string
	AdHocId    uuid.UUID
	Parameters map[string]any
}

type HangUpCommand struct {
	SessionMessage
	AdHocId uuid.UUID
}

type DigitCommand struct {
	SessionMessage
	Key string
}

type DisconnectCommand struct {
	SessionMessage
}
