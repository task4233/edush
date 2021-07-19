package shell

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/docker/docker/client"
	"github.com/gorilla/websocket"
	"github.com/taise-hub/edush/container"
	"github.com/taise-hub/edush/model"
)

func StdInListner(conn *websocket.Conn, que *model.CmdQueue) {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		cmdResult, err := cmdExecOnContainer("hogehoge_container", p)
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

func cmdExecOnContainer(name string, p []byte) ([]byte, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	reader, err := container.Exec(name, fmt.Sprintf("%s", p), cli)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//TODO: stdout, stderr両方返すようにする。
	output, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return output, nil
}
