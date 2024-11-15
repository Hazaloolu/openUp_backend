package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hazaloolu/openUp_backend/internal/model"
	"github.com/hazaloolu/openUp_backend/internal/storage"
)

func Create_post(c *gin.Context) {
	var post model.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("UserID").(uint)
	post.AuthorID = userID

	if err := storage.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var createdPost model.Post
	if err := storage.DB.Preload("Author").First(&createdPost, post.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user information", "details": err.Error()})
		return
	}

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

func GetAllPosts(c *gin.Context) {
	var posts []model.Post
	var total int64

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	offset := (page - 1) * limit

	if err := storage.DB.Offset(offset).Limit(limit).Preload("User").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve posts"})
		return
	}

	storage.DB.Model(&model.Post{}).Count(&total)

	c.JSON(http.StatusOK, gin.H{
		"posts":      posts,
		"page":       page,
		"limit":      limit,
		"total":      total,
		"totalPages": (total + int64(limit) - 1) / int64(limit),
	})
}
