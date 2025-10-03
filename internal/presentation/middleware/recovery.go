package middleware

import (
	"example/hona/bootstrap"
	"example/hona/internal/domain/exceptions"
	"example/hona/internal/presentation/controller"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RecoveryMiddleware struct {
	constants *bootstrap.Constants
}

func NewRecoveryMiddleware(constants *bootstrap.Constants) *RecoveryMiddleware {
	return &RecoveryMiddleware{
		constants: constants,
	}
}

func (rm *RecoveryMiddleware) Recovery(c *gin.Context) {
	defer func() {
		if rec := recover(); rec != nil {
			if err, ok := rec.(error); ok {
				rm.handleRecoveredError(c, err)
				c.Abort()
			}
		}
	}()
	c.Next()
}

func (rm *RecoveryMiddleware) handleRecoveredError(c *gin.Context, err error) {
	if bindingError, ok := err.(*exceptions.BindingError); ok {
		rm.handleBindingError(c, bindingError)
	} else if validationErrors, ok := err.(*exceptions.ValidationErrors); ok {
		rm.handleValidationErrors(c, validationErrors)
	} else {
		rm.unhandledErrors(c, err)
	}
}

func (rm *RecoveryMiddleware) handleBindingError(c *gin.Context, bindingError *exceptions.BindingError) {
	translator := controller.GetTranslator(c, rm.constants.Context.Translator)
	message, _ := translator.T("errors.generic")
	if numError, ok := bindingError.Err.(*strconv.NumError); ok {
		message, _ = translator.T("errors.numeric", numError.Num)
	}
	controller.Response(c, 400, message, nil)
}

func (rm *RecoveryMiddleware) handleValidationErrors(c *gin.Context, validationErrors *exceptions.ValidationErrors) {
	translator := controller.GetTranslator(c, rm.constants.Context.Translator)

	messages := make(map[string]string)
	for _, validationError := range validationErrors.FieldErrors {
		translatedTagValue, _ := translator.T(validationError.Field)
		translatedTag, _ := translator.T("errors."+validationError.Tag, translatedTagValue)
		messages[validationError.Tag] = translatedTag
	}
	controller.Response(c, 422, messages, nil)
}

func (rm *RecoveryMiddleware) unhandledErrors(c *gin.Context, err error) {
	log.Println("Unhandled error occurred:", err)

	translator := controller.GetTranslator(c, rm.constants.Context.Translator)
	message, _ := translator.T("errors.generic")
	controller.Response(c, 500, message, nil)
}
