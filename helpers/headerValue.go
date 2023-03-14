package helpers

import (
	"github.com/gin-gonic/gin"
)

const (
	AppJSON string = "application/json"
)

func GetContentType(c *gin.Context) string {
	return c.Request.Header.Get("Content-Type")
}
