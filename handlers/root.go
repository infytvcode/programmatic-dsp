package handlers

import (
	"github.com/gin-gonic/gin"
)

func RootHandler(c *gin.Context) {
	c.String(200, "Hello, World! 1234")
}
