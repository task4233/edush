package model

import (
	"github.com/gorilla/websocket"
)

type Client struct {

	Name string //コンテナを識別するためのID
	Conn *websocket.Conn
	Room *Room
	Message chan []byte
}



func NewClient(name string, conn *websocket.Conn, room *Room) *Client {
	return &Client{
		Name: name,
		Conn: conn,
		Room: room,
	}
}

func (c *Client) read() {
	for {
		if _, msg, err := c.Conn.ReadMessage(); err == nil {
			c.Room.Forward <- msg
		} else {
			break
		}
	}
	c.Conn.Close()
}

func (c *Client) write() {
	for {
		select {
		case msg := <- c.Message:
			if err := c.Conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				break
			}
		}
		break
	}
}