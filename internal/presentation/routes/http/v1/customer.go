package http

import (
	"example/hona/internal/presentation/controller/v1/customer"

	"github.com/gin-gonic/gin"
)

func SetupCustomerRoutes(router *gin.RouterGroup) {
	salamGroup := router.Group("/salam")
	{
		salamGroup.GET("/kian", customer.Salam)
		salamGroup.GET("/mmd", customer.Salam)
		salamGroup.GET("/mobina", customer.Salam)
	}

	router.GET("/add/:num1/:num2")
}
