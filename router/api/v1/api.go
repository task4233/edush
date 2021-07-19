package v1

import (
	"log"

	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/taise-hub/edush/container"
	"github.com/taise-hub/edush/model"
	"github.com/taise-hub/edush/shell"
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
	q :=  make(chan model.ExecResult)
	go func() {
		for {
			execResult := StdInListner(conn)
			q <- execResult
		}
	}()

	for {
		select {
		case execResult := <-q:
			if err := conn.WriteMessage(websocket.TextMessage, execResult.StdOut); err != nil {
				log.Println(err)
				return
			}
		}
	}
}

func StdInListner(conn *websocket.Conn) model.ExecResult {
	_, p, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return model.ExecResult{}
	}
	execResult, err := shell.CmdExecOnContainer("hogehoge_container", p)
	if err != nil {
		log.Println(err)
		return model.ExecResult{}
	}
	return execResult
}
