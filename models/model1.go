package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Name      string `json:"name"`
	ID        uint    `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Gender    string `json:"gender"`
	BloodType string `json:"bloodtype"`
	Num       int32   `json:"num"`
}
type PersonRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type PersonResponse struct {
	Name      string `json:"name"`
	ID        uint    `json:"id"` 
	Gender    string `json:"gender"`
	BloodType string `json:"bloodtype"`
	Num       int32   `json:"num"`
}
type Pagination struct{
	Page	string `json:"page"`
	Limit string `json:"limit"`
}
