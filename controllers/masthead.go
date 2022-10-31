package controllers

import (
	"net/http"
	"time"

	"example.com/m/models"
	"github.com/gin-gonic/gin"
)

func GetMastheads(c *gin.Context) {
	var mastheads []models.Masthead
	models.DB.Find(&mastheads)

	c.JSON(http.StatusOK, gin.H{"data": mastheads})
}

func CreateMasthead(c *gin.Context) {
	var input models.CreateMastheadInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	masthead := models.Masthead{
		ImageURL:  input.ImageURL,
		Link:      input.Link,
		StartTime: time.Now(),
		EndTime:   time.Now(),
		Order:     input.Order,
		Status:    input.Status,
	}
	models.DB.Create(&masthead)

	c.JSON(http.StatusOK, gin.H{"data": masthead})
}

func GetMasthead(c *gin.Context) {
	var masthead models.Masthead

	if err := models.DB.Where("id = ?", c.Param("id")).First(&masthead).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": masthead})
}

func UpdateMasthead(c *gin.Context) {
	var masthead models.Masthead

	if err := models.DB.Where("id = ?", c.Param("id")).First(&masthead).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input models.UpdateMastheadInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&masthead).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": masthead})
}

func DeleteMasthead(c *gin.Context) {
	var masthead models.Masthead

	if err := models.DB.Where("id = ?", c.Param("id")).First(&masthead).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Delete(&masthead)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
