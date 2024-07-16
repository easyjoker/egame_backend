package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", helloHandler)
	r.POST("/player/:id/deposit", AddMoneyHandler)
	r.POST("/player", CreatePlayerHandler)
	r.POST("/login", LoginHandler)

}

func helloHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, HTTP!")
}
