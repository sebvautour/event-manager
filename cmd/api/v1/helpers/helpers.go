package helpers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func NotImplemented(c *gin.Context) {
	Error(c, "not implemented", http.StatusNotImplemented)
}

func Error(c *gin.Context, msg string, status int, details ...string) {

	if len(details) == 0 {
		c.JSON(status, gin.H{
			"message": msg,
			"path":    c.FullPath(),
			"method":  status,
		})
		return
	}

	c.JSON(status, gin.H{
		"message": msg,
		"path":    c.FullPath(),
		"method":  status,
		"details": strings.Join(details, ", "),
	})
}
