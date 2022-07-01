package controller

import (
	"app/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(ctx *gin.Context) {
	body := models.User_no_gorm{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user = models.User{}
	user.PairBody(&body)
	if result := h.DB.First(&user, "email=?", user.email); result != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
}
