package controller

import (
	"app/pkg/common/hash"
	"app/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(ctx *gin.Context) {
	var body = models.User_login{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.User
	result := h.DB.First(&user, "email=?", body.email)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	hashedInput, _ := hash.HashPassword(body.password)
	err := hash.CheckPasswordHash(hashedInput, user.passwordhash)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	ctx.JSON(http.StatusAccepted, &result)
}
