package controllers

import (
	"net/http"
	"web-framework/auth"
	"web-framework/models/domains"
	"web-framework/models/dtos"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	auth auth.IAtuh
}

func NewUserController(auth auth.IAtuh) *UserController {
	return &UserController{
		auth: auth,
	}
}

func (h *UserController) Handler(r *gin.Engine) {
	router := r.Group("/api/Platform/Command")
	{
		authRoute := router.Group("")
		{
			authRoute.Use(h.auth.Authenticate())
			authRoute.GET("/", h.getUser)

		}

		router.POST("/", h.login)
	}

}

func (h *UserController) getUser(ctx *gin.Context) {
	userInfo := domains.Guser
	ctx.JSON(http.StatusOK, userInfo)
}

func (h *UserController) login(c *gin.Context) {
	loginDto := dtos.LoginDto{}
	err := c.ShouldBindJSON(&loginDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	user := domains.Guser
	if loginDto.Email != user.Email || loginDto.Password != user.Pass {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	token, err := h.auth.GenerateToken(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, token)

}
