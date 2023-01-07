package main

import (
	"github.com/feynmaz/gormg/initializers"
	"github.com/feynmaz/gormg/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
