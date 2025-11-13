package constants

import "errors"

var ErrInvalidContext error = errors.New("invalid context for request")
var ErrWebsocketConnection error = errors.New("couldn't open websockets connection")
var ErrInvalidCommand error = errors.New("invalid command")
