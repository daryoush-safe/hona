package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type singleResponseMessage struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type multipleResponseMessage struct {
	StatusCode int               `json:"statusCode"`
	Messages   map[string]string `json:"messages"`
	Data       interface{}       `json:"data"`
}

func Response[T string | map[string]string](c *gin.Context, statusCode int, message T, data interface{}) {
	switch msg := any(message).(type) {
	case map[string]string:
		c.JSON(statusCode, multipleResponseMessage{
			StatusCode: statusCode,
			Messages:   msg,
			Data:       data,
		})
	case string:
		if msg == "" {
			msg = http.StatusText(statusCode)
		}
		c.JSON(statusCode, singleResponseMessage{
			StatusCode: statusCode,
			Message:    msg,
			Data:       data,
		})
	}
}
