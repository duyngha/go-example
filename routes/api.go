package routes

import (
	"net/http"

	"example.com/m/controllers"
	"example.com/m/middlewares"
	"example.com/m/models"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	return setupRouter()
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	models.ConnectionDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello, world!"})
	})

	r.GET("/mastheads", controllers.GetMastheads)

	r.POST("/mastheads", controllers.CreateMasthead)

	r.GET("/mastheads/:id", controllers.GetMasthead)

	r.PATCH("/mastheads/:id", controllers.UpdateMasthead)

	r.DELETE("/mastheads/:id", controllers.DeleteMasthead)

	r.POST("/token", controllers.GenerateToken)

	secured := r.Group("/logged").Use(middlewares.Auth())
	{
		secured.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"success": true})
		})
	}

	return r
}
