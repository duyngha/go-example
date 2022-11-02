package controllers

import (
	"fmt"
	"net/http"
	"time"

	"example.com/m/models"
	"example.com/m/requests"
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
	var masthead models.Masthead

	if err := models.DB.Where("id = ?", c.Param("id")).First(&masthead).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": masthead})
}

func UpdateMasthead(c *gin.Context) {
	var masthead models.Masthead

	// tmp := c.Request.FormValue("order")
	// fmt.Println(reflect.TypeOf(tmp))
	// if tmp == "" {
	// 	fmt.Print("hey!")
	// }

	if err := models.DB.Where("id = ?", c.Param("id")).First(&masthead).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input requests.UpdateMastheadInput

	// if err := c.ShouldBindJSON(&input); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// models.DB.Model(&masthead).Updates(map[string]interface{}{
	// 	"ImageURL": input.ImageURL,
	// 	"Link":     input.Link,
	// 	"Order":    input.Order,
	// 	"Status":   input.Status,
	// })

	// fmt.Println(input)

	var m = map[string]interface{}{
		"ImageURL": input.ImageURL,
		"Link":     input.Link,
		"Order":    input.Order,
		"Status":   input.Status,
	}

	tmp := c.Request.ParseForm()
	fmt.Println(tmp)

	for k, v := range m {
		if v == "" || v == nil {
			delete(m, k)
		}
	}

	// fmt.Println(m)
	models.DB.Model(&masthead).Updates(m)

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
