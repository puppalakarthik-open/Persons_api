package app

import (
	"example/persons/config"
	"example/persons/database"
	"example/persons/control"
	"fmt"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func App() {
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	config.Load_config()
	database.Connection()
	fmt.Println("Working...")
	router.POST("/login", control.Login)
	Urls()
	router.Run(":" + config.AppConfig.Port)
}
