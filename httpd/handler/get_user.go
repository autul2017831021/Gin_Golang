package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {

	userInfo := User{
		Name: "Autul",
		Age:  23,
		Pass: "abcd",
	}

	c.JSON(http.StatusOK, userInfo)
}
