package user

import (
	"github.com/gin-gonic/gin"
	"github.com/mdtajulislammt/zvonsystem_backend/internal/middleware"
)

func RegisterRoutes(r *gin.Engine, c *UserController) {
	routes := r.Group("/api/users")
	routes.Use(middleware.AuthMiddleware())
	routes.POST("/", c.Create)
	routes.GET("/", c.GetAll)
}
