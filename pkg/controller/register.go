package controller

import (
	"app/pkg/common/hash"
	"app/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Registeration(ctx *gin.Context) {
	body := models.User_no_gorm{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body.passwordhash = hash.HashPassword(body.passwordhash)
	var user models.User

	user.PairBody(&body)
	if control := h.DB.First(&user, "email=?", user.email); control.Error != nil {
		ctx.AbortWithError(http.StatusConflict, control.Error)
		return
	}
	if result := h.DB.Create(&user); result != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	ctx.JSON(http.StatusCreated, &user)
}
