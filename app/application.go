package app

import (
	"example/persons/config"
	"example/persons/database"
	"fmt"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func App() {
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	config.LoadConfig()
	database.Connection()
	fmt.Println("Working...")
	Urls()
	router.Run(":" + config.AppConfig.Port)
}
