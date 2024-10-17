package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hazaloolu/openUp_backend/internal/auth"
	"github.com/hazaloolu/openUp_backend/internal/model"
	"github.com/hazaloolu/openUp_backend/internal/storage"
)

func Register(c *gin.Context) {

	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check if enail already exists

	var existingUser model.User

	if err := storage.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := auth.HashPassword(user.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Password = hashedPassword

	if err := storage.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User Created Succesfully"})

}
