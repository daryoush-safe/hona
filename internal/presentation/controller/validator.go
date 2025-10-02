package controller

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/en"
	"github.com/go-playground/validator/v10/translations/fa"
)

var validate *validator.Validate = validator.New(validator.WithRequiredStructEnabled())

func setupTranslation(c *gin.Context) {
	trans, _ := c.Get("translator")
	translator := trans.(ut.Translator)

	fa.RegisterDefaultTranslations(validate, translator)
	en.RegisterDefaultTranslations(validate, translator)
}

func Validator[T any](c *gin.Context) T {
	var params T

	setupTranslation(c)

	trans, _ := c.Get("translator")
	translator := trans.(ut.Translator)

	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(422, gin.H{"msg": "wrong input"})
	}

	if err := c.ShouldBind(&params); err != nil {
		c.JSON(422, gin.H{"msg": "wrong input"})
	}

	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(422, gin.H{"msg": "wrong input"})
	}

	if err := validate.Struct(&params); err != nil {
		validationErrors, _ := err.(validator.ValidationErrors)

		messages := make(map[string]string)
		for _, err := range validationErrors {
			messages[err.Tag()] = err.Translate(translator)
		}
		c.JSON(422, messages)
	}

	return params
}
