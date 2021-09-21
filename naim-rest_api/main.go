package main

import (
	"fmt"
	"web-framework/configs"
	"web-framework/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	//loading config from json
	config := configs.LoadConfig()

	router := gin.Default()

	controllers.NewUserController().Handler(router)

	router.Run(fmt.Sprintf(":%d", config.Port))
}
