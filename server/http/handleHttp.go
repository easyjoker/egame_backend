package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", helloHandler)
}

func helloHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, HTTP!")
}
