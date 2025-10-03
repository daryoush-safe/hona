package customer

import (
	"example/hona/bootstrap"
	"example/hona/internal/application/dto/math"
	"example/hona/internal/application/service"
	"example/hona/internal/presentation/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	mathService *service.MathService
	constants   *bootstrap.Constants
}

func NewCustomerController(mathService *service.MathService, constants *bootstrap.Constants) *CustomerController {
	return &CustomerController{
		mathService: mathService,
		constants:   constants,
	}
}

func (cc *CustomerController) Salam(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "aleyk",
	})
}

func (cc *CustomerController) Adder(c *gin.Context) {
	type addParams struct {
		Num1 int `uri:"num1" validate:"required"`
		Num2 int `uri:"num2" validate:"required"`
	}
	params := controller.Validator[addParams](c, cc.constants)

	p := math.AddRequest{
		Num1: params.Num1,
		Num2: params.Num2,
	}
	res := cc.mathService.Adder(p)

	translator := controller.GetTranslator(c, cc.constants.Context.Translator)
	message, _ := translator.T("success.add")
	controller.Response(c, 200, message, res)
}

func (cc *CustomerController) SayHello(c *gin.Context) {
	type helloParams struct {
		Name string `form:"name" validate:"required"`
	}
	params := controller.Validator[helloParams](c, cc.constants)

	translator := controller.GetTranslator(c, cc.constants.Context.Translator)
	message, _ := translator.T("success.hello", params.Name)
	controller.Response(c, 200, message, nil)
}
