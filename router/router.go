package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/taise-hub/edush/router/api/v1"
	"github.com/taise-hub/edush/util"
)

func Init() *gin.Engine {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*")
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/join", v1.GetJoin)
	r.POST("/join", v1.PostClientInfo)

	r.Use(util.SessionCheck())
	{
		r.GET("/", v1.GetHome)
		r.GET("/ws", v1.WsCmd)
		r.POST("/judge",v1.Judge)
	}
	return r
}
