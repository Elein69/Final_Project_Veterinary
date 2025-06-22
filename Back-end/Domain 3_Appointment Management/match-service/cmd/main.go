package main

import (
	"github.com/Elein69/match-service/neo4j"
	"github.com/Elein69/match-service/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	neo4j.InitDriver()

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":3002")
}
