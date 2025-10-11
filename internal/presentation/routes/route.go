package routes

import (
	"example/hona/bootstrap"
	"example/hona/internal/presentation/middleware"
	"example/hona/internal/presentation/routes/http/v1"

	"github.com/gin-gonic/gin"
)

func Run(ginEngine *gin.Engine) {
	config := bootstrap.Run()
	recovery := middleware.NewRecoveryMiddleware(config.Constants)
	localization := middleware.NewLocalizationMiddleware()

	ginEngine.Use(localization.Localization)
	ginEngine.Use(recovery.Recovery)

	v1 := ginEngine.Group("/v1")
	{
		http.SetupGeneralRoutes(v1)
		http.SetupCustomerRoutes(v1)
	}
}
