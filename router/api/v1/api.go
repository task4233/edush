package v1

import (
	"log"
	"github.com/google/uuid"

	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gorilla/websocket"
	"github.com/taise-hub/edush/container"
	"github.com/taise-hub/edush/model"
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

var spv = model.NewSupervisor()

func WsCmd(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
	}
	session := sessions.Default(c)
	id := session.Get("id").(string)
	room := session.Get("room").(string)
	spv.Append(id, room, conn)	
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
		id := c.PostForm("name") + "-" + uuid.NewString()
		room := c.PostForm("room")

		session.Set("id", id)
		session.Set("room", room)
		session.Save()
	}
	c.Redirect(302, "/")
}