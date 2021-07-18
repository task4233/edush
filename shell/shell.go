package shell

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/gorilla/websocket"
	"github.com/taise-hub/edush/model"
)

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
			out, err := cmdExec(p)
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

func cmdExec(p []byte) ([]byte, error) {
	cmd := fmt.Sprintf("%s", p)
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return nil, err
	}
	return out, nil
}
