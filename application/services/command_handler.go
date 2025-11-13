package services

type CommandHandler interface {
	Parse(command string) (any, error)
	Handle(command string) (string, error)
}
