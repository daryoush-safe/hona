package middleware

import (
	"example/hona/internal/infrastructure/translation"

	"github.com/gin-gonic/gin"
)

func Localization(c *gin.Context) {
	locale := c.Request.Header.Get("Accept-Language")
	if locale == "" {
		locale = "fa_IR"
	}
	translator := translation.GetTranslator(locale)

	c.Set("translator", translator)

	c.Next()
}
