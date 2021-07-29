package util

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gin-contrib/sessions"
)

func SessionCheck() gin.HandlerFunc {
    return func(c *gin.Context) {
        s := sessions.Default(c)
        id := s.Get("id")

        if id == nil {
            c.Redirect(http.StatusMovedPermanently, "/join")
            c.Abort()
        } else {
            c.Set("id", id)
            c.Next()
        }
    }
}
