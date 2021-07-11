package v1

import(
	"github.com/gin-gonic/gin"
	"os/exec"
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