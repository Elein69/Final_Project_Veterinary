package routes

import (
	"github.com/Elein69/match-service/controllers"
	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all API endpoints for the match service.
func SetupRoutes(r *gin.Engine) {
	r.POST("/match", controllers.FindMatch)
}
