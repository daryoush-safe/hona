package main

import (
	"example/hona/internal/presentation/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.DisableConsoleColor()

	ginEngine := gin.Default()

	routes.Run(ginEngine)

	ginEngine.Run()
}
