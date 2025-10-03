package controller

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New(validator.WithRequiredStructEnabled())

func Validator[T any](c *gin.Context) T {
	var params T

	translator := GetTranslator(c, "translator")

	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(400, gin.H{"msg": "wrong input"})
	}

	if err := c.ShouldBind(&params); err != nil {
		c.JSON(400, gin.H{"msg": "wrong input"})
	}

	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(400, gin.H{"msg": "wrong input"})
	}

	if err := validate.Struct(&params); err != nil {
		validationErrors, _ := err.(validator.ValidationErrors)

		messages := make(map[string]string)
		for _, err := range validationErrors {
			tagValue := formatValidationError[T](err)
			translatedTagValue, _ := translator.T(tagValue)
			translatedTag, _ := translator.T("errors."+err.Tag(), translatedTagValue)
			messages[err.Tag()] = translatedTag
		}
		c.JSON(422, messages)
	}

	return params
}

func formatValidationError[T any](err validator.FieldError) string {
	var params T

	tagTypes := []string{"json", "uri", "form"}

	reflectType := reflect.TypeOf(params)

	field, _ := reflectType.FieldByName(err.StructField())
	tagValue := getAnyTag(field, tagTypes...)

	return tagValue
}

func getAnyTag(field reflect.StructField, tagNames ...string) string {
	for _, tagName := range tagNames {
		if tag := field.Tag.Get(tagName); tag != "" {
			return tag
		}
	}
	return field.Name
}
