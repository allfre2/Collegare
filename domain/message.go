package domain

// Message Types
const (
	SessionConfig = "configuration"
)

type SessionMessage struct {
	Type string
}

type ConfigurationMessage struct {
	SessionMessage
	CallConfiguration
}
