package controllers

import (
	"net/http"

	"github.com/ddenizakpinar/wendi-go/models"
	"github.com/gin-gonic/gin"
)

type CreateMonsterInput struct {
	Name           string `json:"name" binding:"required"`
	Classification string `json:"classification" binding:"required"`
}

type UpdateMonsterInput struct {
	Name           string `json:"name"`
	Classification string `json:"classification"`
}

// GET /Monsters
func FindMonsters(c *gin.Context) {
	var Monsters []models.Monster
	models.DB.Find(&Monsters)

	c.JSON(http.StatusOK, gin.H{"data": Monsters})
}

// POST /Monsters
func CreateMonster(c *gin.Context) {
	// Validate input
	var input CreateMonsterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Monster
	Monster := models.Monster{Name: input.Name, Classification: input.Classification}
	models.DB.Create(&Monster)

	c.JSON(http.StatusOK, gin.H{"data": Monster})
}

// GET /Monsters/:id
func FindMonster(c *gin.Context) { // Get model if exist
	var Monster models.Monster

	if err := models.DB.Where("id = ?", c.Param("id")).First(&Monster).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Monster not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Monster})
}

// PATCH /Monsters/:id
func UpdateMonster(c *gin.Context) {
	// Get model if exist
	var Monster models.Monster
	if err := models.DB.Where("id = ?", c.Param("id")).First(&Monster).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Monster not found!"})
		return
	}

	// Validate input
	var input UpdateMonsterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&Monster).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": Monster})
}

// DELETE /Monsters/:id
func DeleteMonster(c *gin.Context) {
	var Monster models.Monster
	if err := models.DB.Where("id = ?", c.Param("id")).First(&Monster).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Monster not found!"})
		return
	}

	models.DB.Delete(&Monster)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
