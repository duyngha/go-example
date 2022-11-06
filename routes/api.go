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

	r.POST("/auth/token", controllers.GenerateToken)

	secured := r.Group("/api").Use(middlewares.Auth())
	{
		secured.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"success": true})
		})

		secured.GET("/mastheads", controllers.GetMastheads)

		secured.POST("/mastheads", controllers.CreateMasthead)

		secured.GET("/mastheads/:id", controllers.GetMasthead)

		secured.PATCH("/mastheads/:id", controllers.UpdateMasthead)

		secured.DELETE("/mastheads/:id", controllers.DeleteMasthead)
	}

	return r
}
