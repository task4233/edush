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
		Message: make(chan []byte),
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
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		} 
		c.Room.Forward <- msg
	}
}

func (c *Client) write() {
	for {
		msg := <- c.Message
		if err := c.Conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
}