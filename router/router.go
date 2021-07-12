package router

import (
	"github.com/gin-gonic/gin"
	"github.com/taise-hub/edush/router/api/v1"
)

func Init() *gin.Engine{
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", v1.GetHome)
	r.POST("/exec", v1.PostCmd)
	return r
}