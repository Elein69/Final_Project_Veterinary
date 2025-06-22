package controllers

import (
	"net/http"

	"github.com/Elein69/match-service/models"
	"github.com/Elein69/match-service/services"
	"github.com/gin-gonic/gin"
)

func FindMatch(c *gin.Context) {
	var req models.MatchRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	match, err := services.FindBestVeterinarian(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if match == nil {
		c.JSON(http.StatusOK, gin.H{"match_found": false, "message": "No match found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"match_found":  true,
		"veterinarian": match,
	})
}
