package routes

import (
	"example/hona/internal/presentation/routes/http/v1"

	"github.com/gin-gonic/gin"
)

func Run(ginEngine *gin.Engine) {
	v1 := ginEngine.Group("/v1")
	{
		http.SetupGeneralRoutes(v1)
		http.SetupCustomerRoutes(v1)
	}
}
