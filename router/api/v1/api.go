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
	go shell.StdInListner(conn, q)
	go shell.StdOut(conn, q)
}
