package http

import (
	"example/hona/bootstrap"
	"example/hona/internal/application/service"
	"example/hona/internal/presentation/controller/v1/customer"

	"github.com/gin-gonic/gin"
)

func SetupCustomerRoutes(router *gin.RouterGroup) {
	config := bootstrap.Run()
	mathService := service.NewMathService()
	customerController := customer.NewCustomerController(mathService, config.Constants)

	salamGroup := router.Group("/salam")
	{
		salamGroup.GET("/kian", customerController.Salam)
		salamGroup.GET("/mmd", customerController.Salam)
		salamGroup.GET("/mobina", customerController.Salam)
	}

	router.GET("/add/:num1/:num2", customerController.Adder)
	router.GET("/hello", customerController.SayHello)
}
