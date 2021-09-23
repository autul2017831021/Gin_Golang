package main

import (
	"fmt"
	"web-framework/auth"
	"web-framework/configs"
	"web-framework/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	//loading config from json
	config := configs.LoadConfig()

	router := gin.Default()

	appconfig := configs.LoadConfig()
	auth := auth.NewAuth(appconfig)
	controllers.NewUserController(auth).Handler(router)

	router.Run(fmt.Sprintf(":%d", config.Port))
}
