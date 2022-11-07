package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"example.com/m/internal/models"
	"example.com/m/internal/requests"
	"github.com/gin-gonic/gin"
)

func GetMastheads(c *gin.Context) {
	var mastheads []models.Masthead
	models.DB.Find(&mastheads)

	c.JSON(http.StatusOK, gin.H{"data": mastheads})
}

func CreateMasthead(c *gin.Context) {
	var input requests.CreateMastheadInput

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
	log.Printf("%v", c.Request.Header)
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

	var input requests.UpdateMastheadInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Dùng Marshal để chuyển những key không có giá trị, hoặc không tồn tại trên response về nil
	// Sau đó dùng Unmarshal để remove những key có giá trị nil ra khỏi interface
	var data map[string]interface{}
	tmpData, _ := json.Marshal(input)
	json.Unmarshal(tmpData, &data)

	models.DB.Model(&masthead).Updates(data)

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