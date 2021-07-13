package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/taise-hub/edush/model"
	"log"
	"os/exec"
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
	go StdInListner(conn, q)
	go StdOut(conn, q)
}

// websocketからのコマンドをチャネルで排他制御する。
func StdInListner(conn *websocket.Conn, que *model.CmdQueue) {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		que.Pipe <- p
		defer conn.Close()
	}
}

//pipeで受け取ったコマンドを実行するだけ
func StdOut(conn *websocket.Conn, que *model.CmdQueue) {
	for {
		select {
		case p := <-que.Pipe:
			out, err := CmdExec(p)
			if err != nil {
				log.Println(err)
				return
			}
			if err := conn.WriteMessage(websocket.TextMessage, out); err != nil {
				log.Println(err)
				return
			}
		}
	}
}

func CmdExec(p []byte) ([]byte, error) {
	cmd := fmt.Sprintf("%s", p)
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return nil, err
	}
	return out, nil
}
