package v1

import (
	"log"
	"github.com/google/uuid"

	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gorilla/websocket"
	"github.com/taise-hub/edush/container"
	"github.com/taise-hub/edush/judge"
	"github.com/taise-hub/edush/model"
	"github.com/taise-hub/edush/shell"
)

func GetHome(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("id").(string)

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Print(err)
		return
	}
	if exist := container.IsContainerExists(id, cli); !exist {
		if err := container.Run(id, cli); err != nil {
			log.Print(err)
			return
		}
	}

	c.HTML(200, "index.html", nil)
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
	q := make(chan model.ExecResult)
	session := sessions.Default(c)
	id := session.Get("id").(string)
	go func() {
		for {
			execResult := StdInListner(conn, id)
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

func StdInListner(conn *websocket.Conn, id string) model.ExecResult {
	_, p, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return model.ExecResult{}
	}
	execResult, err := shell.CmdExecOnContainer(id, p)
	if err != nil {
		log.Println(err)
		return model.ExecResult{}
	}
	return execResult
}

func Judge(c *gin.Context) {
	ans := c.PostForm("answer")
	result := judge.Problem1(ans)
	c.JSON(200, gin.H{
		"message": ans,
		"result":  result,
	})
}

func GetJoin(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("id") == nil {
		c.HTML(200, "join.html", nil)
	}
	c.Redirect(302, "/")
}

func PostClientInfo(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("id") == nil {
		name := c.PostForm("name")
		id := name + "-" + uuid.NewString()
		log.Print(id)
		session.Set("id", id)
		session.Save()	
	}
	c.Redirect(302, "/")
}