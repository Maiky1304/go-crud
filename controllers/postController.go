package controllers

import (
	"github.com/Maiky1304/go-crud/initializers"
	"github.com/Maiky1304/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	// Get data off request body
	var body struct {
		Title string
		Body  string
	}
	err := c.Bind(&body)
	if err != nil {
		c.Status(400)
		return
	}

	// Create post
	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return the post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostFind(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	if post.ID == 0 {
		c.Status(404)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	var post models.Post
	initializers.DB.First(&post, id)

	if post.ID == 0 {
		c.Status(404)
		return
	}

	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	if post.ID == 0 {
		c.Status(404)
		return
	}

	initializers.DB.Delete(&post)
	c.JSON(200, gin.H{
		"deleted": true,
	})
}
