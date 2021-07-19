package v1

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/docker/docker/client"
	"github.com/taise-hub/edush/model"
	"github.com/taise-hub/edush/shell"
	"github.com/taise-hub/edush/container"
)

func GetHome(c *gin.Context) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return
	}
	if err := container.Run("hogehoge_container", cli); err != nil {
		return
	}
	c.HTML(200, "edush.html", nil)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WsCmd(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
	}
	q := model.NewCmdQueue()
	go StdInListner(conn, q)
	go StdOut(conn, q)
}

func StdInListner(conn *websocket.Conn, que *model.CmdQueue) {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		cmdResult, err := shell.CmdExecOnContainer("hogehoge_container", p)
		if err != nil {
			log.Println(err)
			return
		}
		que.Pipe <- cmdResult
	}
}

func StdOut(conn *websocket.Conn, que *model.CmdQueue) {
	for {
		select {
		case output := <-que.Pipe:
			if err := conn.WriteMessage(websocket.TextMessage, output); err != nil {
				log.Println(err)
				return
			}
		}
	}
}