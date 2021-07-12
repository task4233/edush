package v1

import(
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"os/exec"
	"log"
	"fmt"
)

func PostCmd(c *gin.Context) {
	cmd := c.PostForm("cmd")
	// judge.Judge(cmd)
	// judgeで判定。正解だったら勝ち、その他のユーザに負けを通知する。
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		c.JSON(500, gin.H{
			"result": err,
		})
	}
	c.JSON(200, gin.H{
		"result": out,
	})
}

func GetHome(c *gin.Context) {
	c.HTML(200, "edush.html",nil)
}

//QUEUE=====================================================
type CmdQueue struct {
	StdIn chan []byte
	StdOut chan []byte
}

func NewCmdQueue() *CmdQueue {
	return &CmdQueue{
		StdIn: make(chan []byte),
		StdOut: make(chan []byte),
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
func StdInListner(conn *websocket.Conn,que *CmdQueue) {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		que.StdIn <- p
		defer conn.Close()
	}
}

//pipeで受け取ったコマンドを実行するだけ
func CmdExec(conn *websocket.Conn, que *CmdQueue) []byte {
	for {
		select {
		case p := <- que.StdIn:
			cmd := fmt.Sprintf("%s", p)
			out, err := exec.Command("bash", "-c", cmd).Output()
			if err != nil {
				log.Println(err)
				return nil
			}
			return out
		}
	}
}