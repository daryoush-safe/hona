package customer

import (
	"example/hona/internal/application/dto/math"
	"example/hona/internal/application/service"
	"example/hona/internal/presentation/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerController struct{
	mathService *service.MathService
}

func NewCustomerController(mathService *service.MathService) *CustomerController {
	return &CustomerController{
		mathService: mathService,
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
	params := controller.Validator[addParams](c)

	p := math.AddRequest{
		Num1: params.Num1,
		Num2: params.Num2,
	}
	res := cc.mathService.Adder(p)

	c.JSON(http.StatusOK, gin.H{
		"your num1:":  params.Num1,
		"your num2:":  params.Num2,
		"num1 + num2": res,
	})
}

func (cc *CustomerController) SayHello(c *gin.Context) {
	type helloParams struct {
		Name string `form:"name" validate:"required"`
	}
	params := controller.Validator[helloParams](c)

	translator := controller.GetTranslator(c, "translator")
	message, _ := translator.T("success.hello", params.Name)
	controller.Response(c, 200, message, nil)
}
