package app

import (
	"example/persons/config"
	"example/persons/db"
	"fmt"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func App() {
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	config.LoadConfig()
	db.Connection()
	fmt.Println("Working...")
	Urls()
	router.Run(":" + config.AppConfig.Port)
}
