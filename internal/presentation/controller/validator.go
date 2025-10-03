package controller

import (
	"example/hona/bootstrap"
	"example/hona/internal/domain/exceptions"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New(validator.WithRequiredStructEnabled())

func Validator[T any](c *gin.Context, constants *bootstrap.Constants) T {
	var params T

	if err := c.ShouldBindUri(&params); err != nil {
		bindingError := exceptions.NewBindingError(err)
		panic(bindingError)
	}

	if err := c.ShouldBind(&params); err != nil {
		bindingError := exceptions.NewBindingError(err)
		panic(bindingError)
	}

	if err := c.ShouldBindQuery(&params); err != nil {
		bindingError := exceptions.NewBindingError(err)
		panic(bindingError)
	}

	if err := validate.Struct(&params); err != nil {
		validationErrors, _ := err.(validator.ValidationErrors)
		customValidationError := exceptions.NewValidationErrors()
		for _, err := range validationErrors {
			field := formatValidationError[T](err)
			customValidationError.AddError(field, err.Tag())
		}
		panic(customValidationError)
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
