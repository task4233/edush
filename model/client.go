package model

import (
	"log"
	"github.com/gorilla/websocket"
	"github.com/taise-hub/edush/shell"
)

type Client struct {
	Name string
	Conn *websocket.Conn
	Room *Room
	Message chan shell.ExecResult
	Owner bool
}

func NewClient(name string, conn *websocket.Conn, room *Room) *Client {
	return &Client{
		Name: name,
		Conn: conn,
		Room: room,
		Message: make(chan shell.ExecResult),
	}
}

func (c *Client) joinRoom() bool {
	if len(c.Room.Clients) < 2 {
		c.Room.Clients = append(c.Room.Clients, c)
		return true
	}
	return false
}

func (c *Client) read() {
	for {
		_, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		c.Owner = true
		execResult, err := shell.CmdExecOnContainer(c.Name, p)
		c.Room.Forward <- execResult
	}
}

func (c *Client) write() {
	for {
		execResult := <- c.Message
		if c.Owner == true {
			execResult.Owner = true
		}else {
			execResult.Owner = false
		}
		if err := c.Conn.WriteJSON(execResult); err != nil {
			break
		}
		c.Owner = false
	}
}
