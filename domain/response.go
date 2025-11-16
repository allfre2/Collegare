package domain

type LoginResponse struct {
	SessionMessage
}

type ErrorResponse struct {
	SessionMessage
	Message string
}

type CallFlowsResponse struct {
	Organization string
	Flows        []string
}
