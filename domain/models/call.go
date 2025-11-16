package models

import "github.com/gofrs/uuid/v5"

type CallFlow struct {
	Id   uuid.UUID
	Name string
	Script
}

type Script struct {
	Start int
	Nodes map[int]Node
}

type Node struct {
	Id      int
	Track   Track
	Loop    bool
	Name    string
	Options map[string]int
}

type Track struct {
	Id       uuid.UUID
	Filename string
}
