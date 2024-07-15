package server

import (
	"egame_backend/server/http"
	"egame_backend/server/ws"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	http.RegisterRoutes(r)

	r.GET("/ws", func(c *gin.Context) {
		ws.HandleWebSocket(c.Writer, c.Request)
	})

	return r
}
