package domain

import "github.com/gofrs/uuid/v5"

const (
	Zero     = "0"
	One      = "1"
	Two      = "2"
	Three    = "3"
	Four     = "4"
	Five     = "5"
	Six      = "6"
	Seven    = "7"
	Eight    = "8"
	Nine     = "9"
	Pound    = "#"
	Asterisk = "*"
)

var KeyPad = []string{
	Zero, One, Two, Three, Four, Five, Six, Seven, Eight, Nine, Pound, Asterisk,
}

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
	Mode         string
	Organization string
	Flow         uuid.UUID
	AdHocId      uuid.UUID
}

type HangUpCommand struct {
	SessionMessage
}

type DigitCommand struct {
	SessionMessage
	Key string
}

type DisconnectCommand struct {
	SessionMessage
}

type GetFlowCommand struct {
	SessionMessage
	Organization string
}
