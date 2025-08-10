package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getAllInfrastructuresHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"infrastructures": db.GetAllInfrastructures()})
}

func setInfrastructuresHandler(c *gin.Context) {
	name := uuid.New()
	var req struct {
		// Name    string `json:"name"`
		Type    string `json:"type"`
		Details string `json:"details"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "message": err})
		return
	}
	newInfra := db.AddInfrastructure(req.Type+"-"+name.String(), req.Type, req.Details)
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

func deleteInfrastructuresHandler(c *gin.Context) {
	var req struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := db.DeleteInfrastructure(req.Name); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Infrastructure not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Infrastructure deleted successfully"})
}

func updateInfrastructuresHandler(c *gin.Context) {
	var req struct {
		Name    string `json:"name"`
		Type    string `json:"type"`
		Details string `json:"details"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	infrastructure := db.UpdateInfrastructure(req.Name, req.Type, req.Details)
	if infrastructure == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Infrastructure not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Infrastructure updated successfully", "infrastructure": infrastructure})
}
