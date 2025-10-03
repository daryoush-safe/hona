package controller

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
)

func GetTranslator(c *gin.Context, key string) ut.Translator {
	trans, _ := c.Get(key)
	return trans.(ut.Translator)
}
