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
