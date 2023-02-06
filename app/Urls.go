package app

import (
	"example/persons/Control"
	"example/persons/Control/CreateUser"
	"example/persons/Control/DeleteUser"
	"example/persons/Control/GetUser"
	"example/persons/Control/GetUsers"
	"example/persons/middle"
)

func Urls() {
	router.POST("/CreateUsers",middle.Authenticate, CreateUser.CreateUserHandler)
	router.DELETE("/DeleteUser",middle.Authenticate, DeleteUser.DeleteUserHandler)
	router.GET("/GetUser",middle.Authenticate, GetUser.GetUserHandler)
	router.GET("/GetUsers",middle.Authenticate, GetUsers.GetUsersHandler)
	router.POST("/login", Control.Login)
}
