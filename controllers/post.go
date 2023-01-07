package controllers

import (
	"github.com/feynmaz/gormg/initializers"
	"github.com/feynmaz/gormg/models"
	"github.com/gin-gonic/gin"
)

func CreatePost(ctx *gin.Context) {
	// Get data off request body
	var body struct {
		Title string
		Body  string
	}

	ctx.Bind(&body)

	// Create a post
	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		ctx.Status(400)
		return
	}

	// Return it
	ctx.JSON(200, gin.H{
		"post": post,
	})
}

func GetPosts(ctx *gin.Context) {
	// Get posts
	var posts []models.Post
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		ctx.Status(400)
		return
	}

	ctx.JSON(200, gin.H{
		"posts": posts,
	})
}
