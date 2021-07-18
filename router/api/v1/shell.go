package v1

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/taise-hub/edush/model"
	"github.com/taise-hub/edush/shell"
)

func GetHome(c *gin.Context) {
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
