package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Pass string `json:"pass"`
}

func PostUser(c *gin.Context) {
	user := User{}

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	fmt.Println(user)

	if user.Name != "" && user.Age != 0 && user.Pass != "" {
		c.JSON(http.StatusAccepted, gin.H{"success": true})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"success": false})
}
