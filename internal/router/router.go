package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hazaloolu/openUp_backend/internal/handler"
	"github.com/hazaloolu/openUp_backend/internal/middleware"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/Register", handler.Register)
	r.POST("/login", handler.Login)

	// Protected Routes
	autheticated := r.Group("/")
	autheticated.Use(middleware.AuthMiddleware())
	{
		autheticated.POST("create-post", handler.Create_post)
		autheticated.GET("Feed", handler.GetAllPosts)
	}

	return r

}
