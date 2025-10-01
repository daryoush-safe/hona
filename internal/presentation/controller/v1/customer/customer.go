package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Salam(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "aleyk",
	})
}

func Adder(c *gin.Context) {
}