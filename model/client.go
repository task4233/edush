package model

import (
	"log"
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

func (c *Client) joinRoom() bool {
	log.Println(len(c.Room.Clients))
	if len(c.Room.Clients) < 2 {
		c.Room.Clients = append(c.Room.Clients, c)
		return true
	}
	return false
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