package app

import (
	"example/persons/control"
	"example/persons/control/create_user"
	"example/persons/control/delete_user"
	"example/persons/control/get_user"
	"example/persons/control/get_users"
	"example/persons/middle"
)

func Urls() {
	router.POST("/CreateUsers",middle.Authenticate, create_user.Create_user_handler)
	router.DELETE("/DeleteUser",middle.Authenticate, delete_user.Delete_user_handler)
	router.GET("/GetUser",middle.Authenticate, get_user.Get_user_handler)
	router.GET("/GetUsers",middle.Authenticate, get_users.Get_users_handler)
	router.POST("/login", control.Login)
}
