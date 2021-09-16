package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {

	userInfo := userGlobal

	c.JSON(http.StatusOK, userInfo)
}
