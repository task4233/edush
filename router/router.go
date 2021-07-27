package router

import (
	"github.com/gin-gonic/gin"
	"github.com/taise-hub/edush/router/api/v1"
)

func Init() *gin.Engine {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", v1.GetHome)
	r.GET("/ws", v1.WsCmd)
	r.POST("/judge",v1.Judge)
	return r
}
