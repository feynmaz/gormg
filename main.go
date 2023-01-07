package main

import (
	"github.com/feynmaz/gormg/controllers"
	"github.com/feynmaz/gormg/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/posts", controllers.CreatePost)
	r.PUT("/posts/:id", controllers.UpdatePost)

	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPost)

	r.DELETE("/posts/:id", controllers.DeletePost)



	r.Run()
}
