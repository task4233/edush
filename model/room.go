package model

import (
	"github.com/taise-hub/edush/shell"
)

type Room struct {
	ID string
	Clients []*Client 
	Forward chan shell.ExecResult
}

func NewRoom(id string) *Room {
	return &Room{
		ID: id,
		Clients: make([]*Client,0 , 2),
		Forward: make(chan shell.ExecResult),
	}
}

func (r *Room) run() {
	for {
		execResult:= <- r.Forward
		for _, client := range r.Clients {
			client.Message <- execResult
		}
	}
}