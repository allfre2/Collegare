package errors

import "errors"

var ErrWebsocketConnection error = errors.New("couldn't open websockets connection")
var ErrInvalidContext error = errors.New("invalid context for request")
