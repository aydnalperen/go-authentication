package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &Handler{
		DB: db,
	}

	routes := r.Group("/")
	routes.POST("register", h.Registeration)
	routes.POST("login", h.Login)
	routes.POST("logout", h.Logout)
}
