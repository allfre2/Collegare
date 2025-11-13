package main

import (
	"collegare/infrastructure"
)

func main() {
	server := infrastructure.NewRouter()
	infrastructure.SetupV1(server)
	infrastructure.Start(server)
}
