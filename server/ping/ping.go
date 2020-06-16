package ping

import (
	"github.com/bbiao/nephele/server"
	"github.com/gin-gonic/gin"
)

func init() {
	server.Register(configurater)
}

func configurater(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
}
