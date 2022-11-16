package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"example.com/m/internal/aws"
	"example.com/m/internal/databases"
	"example.com/m/internal/models"
	"example.com/m/internal/requests"
	"github.com/gin-gonic/gin"
)

func GetMastheads(c *gin.Context) {
	var mastheads []models.Masthead
	databases.DB.Find(&mastheads)

	c.JSON(http.StatusOK, gin.H{"data": mastheads})
}

func CreateMasthead(c *gin.Context) {
	var input requests.CreateMastheadInput

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var imageUrl, err = aws.Upload(c, "banners", "image_url")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	masthead := models.Masthead{
		ImageURL:  imageUrl,
		Link:      input.Link,
		StartTime: time.Now(),
		EndTime:   time.Now(),
		Order:     input.Order,
		Status:    input.Status,
	}
	databases.DB.Create(&masthead)

	c.JSON(http.StatusOK, gin.H{"data": masthead})
}

func GetMasthead(c *gin.Context) {
	var masthead models.Masthead

	if err := databases.DB.Where("id = ?", c.Param("id")).First(&masthead).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	masthead.ImageURL = aws.GetFileURLWithCDN(masthead.ImageURL)

	c.JSON(http.StatusOK, gin.H{"data": masthead})
}

func UpdateMasthead(c *gin.Context) {
	var masthead models.Masthead

	if err := databases.DB.Where("id = ?", c.Param("id")).First(&masthead).Error; err != nil {
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

	databases.DB.Model(&masthead).Updates(data)

	c.JSON(http.StatusOK, gin.H{"data": masthead})
}

func DeleteMasthead(c *gin.Context) {
	var masthead models.Masthead

	if err := databases.DB.Where("id = ?", c.Param("id")).First(&masthead).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	databases.DB.Delete(&masthead)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func UploadImage(c *gin.Context) {
	url, err := aws.Upload(c, "banners/", "file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"success": url})
	}
}
