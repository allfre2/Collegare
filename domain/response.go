package domain

type LoginResponse struct {
	SessionMessage
}

type ErrorResponse struct {
	SessionMessage
	Message string
}
