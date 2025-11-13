package constants

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
	Disconnect = "diconnect"
)
