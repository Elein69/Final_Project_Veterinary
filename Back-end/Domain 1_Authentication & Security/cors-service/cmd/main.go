package main

import (
	"github.com/Elein69/cors-service/internal/config"
	"github.com/Elein69/cors-service/internal/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	r := gin.Default()

	r.Use(middleware.CORSMiddleware(cfg.AllowedOrigins))

	r.GET("/validate-cors", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "CORS origin allowed"})
	})

	r.Run(":" + cfg.Port)
}
