package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllInfrastructuresHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"infrastructures": db.GetAllInfrastructures()})
}

func setInfrastructuresHandler(c *gin.Context) {
	var req struct {
		Name    string `json:"name"`
		Type    string `json:"type"`
		Details string `json:"details"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	newInfra := db.AddInfrastructure(req.Name, req.Type, req.Details)
	c.JSON(http.StatusOK, gin.H{"message": "Infrastructure added successfully", "infrastructure": newInfra})
}

func getInfrastructuresHandler(c *gin.Context) {
	var req struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	infrastructure := db.GetInfrastructure(req.Name)
	if infrastructure == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Infrastructure not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"infrastructure": infrastructure})
}
