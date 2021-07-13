package v1

import(
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"os/exec"
	"log"
	"fmt"
)

func GetHome(c *gin.Context) {
	c.HTML(200, "edush.html",nil)
}

//QUEUE=====================================================
type CmdQueue struct {
	Pipe chan []byte
}

func NewCmdQueue() *CmdQueue {
	return &CmdQueue{
		Pipe: make(chan []byte),
	}
}
//=========================================================

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func WsCmd(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
	}
	q := NewCmdQueue()
	go StdInListner(conn, q)
	go CmdExec(conn, q)
}

// websocketからのコマンドをチャネルで排他制御する。
func StdInListner(conn *websocket.Conn, que *CmdQueue) {
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
func CmdExec(conn *websocket.Conn, que *CmdQueue) {
	for {
		select {
		case p := <- que.Pipe:
			cmd := fmt.Sprintf("%s", p)
			out, err := exec.Command("bash", "-c", cmd).Output()
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