package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hazaloolu/openUp_backend/internal/model"
	"github.com/hazaloolu/openUp_backend/internal/storage"
)

func Create_post(c *gin.Context) {
	var post model.Post

	// Bind JSON request to post struct
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the authenticated user's ID from context and set as AuthorID
	userID := c.MustGet("UserID").(uint)
	post.AuthorID = userID

	// Create the post in the database
	if err := storage.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with success message
	var createdPost model.Post
	if err := storage.DB.Preload("Author").First(&createdPost, post.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user information", "details": err.Error()})
		return
	}

	// Respond with success message, post data, and user's name/username
	c.JSON(http.StatusCreated, gin.H{
		"message": "Post created successfully",
		"post": map[string]interface{}{
			"id":         createdPost.ID,
			"title":      createdPost.Title,
			"content":    createdPost.Content,
			"author_id":  createdPost.AuthorID,
			"username":   createdPost.Author.Username, // Include the username of the user
			"created_at": createdPost.CreatedAt,
			"updated_at": createdPost.UpdatedAt,
		},
	})
}
