package domain

type Command struct {
	Type string
}

type ConnectCommand struct {
	Command
}

type LoginCommand struct {
	Command
}

type CallCommand struct {
	Command
}

type HangUpCommand struct {
	Command
}

type DigitCommand struct {
	Command
	Key string
}

type DisconnectCommand struct {
	Command
}
