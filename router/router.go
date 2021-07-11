package router

import (
	"github.com/gin-gonic/gin"
	"github.com/taise-hub/edush/router/api/v1"
)

func Init() *gin.Engine{
	r := gin.Default()

	r.POST("/exec", v1.PostCmd)
	return r
}