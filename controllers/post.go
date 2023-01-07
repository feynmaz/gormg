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

	// respond
	ctx.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetPost(ctx *gin.Context) {
	// Get id off url
	id := ctx.Param("id")

	// Get post
	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		ctx.Status(400)
		return
	}

	// respond
	ctx.JSON(200, gin.H{
		"id":   id,
		"post": post,
	})
}

func UpdatePost(ctx *gin.Context) {
	// Get id off url
	id := ctx.Param("id")

	// Get the data off request body
	var body struct {
		Title string
		Body  string
	}
	ctx.Bind(&body)

	// Find post to update
	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		ctx.Status(400)
		return
	}

	// Update it
	result = initializers.DB.Model(&post).Updates(
		models.Post{
			Title: body.Title,
            Body:  body.Body,
		},
	)
	if result.Error != nil {
		ctx.Status(400)
		return
	}

	// respond
	ctx.JSON(200, gin.H{
		"post": post,
	})
}
