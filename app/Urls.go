package app

import (
	"example/persons/control/create_user"
	"example/persons/control/delete_user"
	"example/persons/control/get_user"
	"example/persons/control/get_users"
	"example/persons/middleware"
)

func Urls() {
	router.Use(middleware.Authenticate)
	router.POST("/CreateUsers",create_user.Create_user_handler)
	router.DELETE("/DeleteUser",delete_user.Delete_user_handler)
	router.GET("/GetUser", get_user.Get_user_handler)
	router.GET("/GetUsers", get_users.Get_users_handler)
}
