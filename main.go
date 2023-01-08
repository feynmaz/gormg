package main

import (
	"net/http"
	"os"

	"github.com/feynmaz/gormg/controllers"
	"github.com/feynmaz/gormg/initializers"
	"github.com/labstack/echo/v4"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	e := echo.New()

	e.GET("/ping", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "pong")
	})

	e.POST("/posts", controllers.CreatePost)
	e.PUT("/posts/:id", controllers.UpdatePost)

	e.GET("/posts", controllers.GetPosts)
	e.GET("/posts/:id", controllers.GetPost)

	e.DELETE("/posts/:id", controllers.DeletePost)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
