package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hazaloolu/openUp_backend/internal/handler"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/Register", handler.Register)
	r.POST("/Login", handler.Login)

	return r
}
