package get_user_repository

import (
	"example/persons/database"
	"example/persons/models"
)
func Get_user_repository(tempperson *[]models.Person, id int)(error) {
	return database.DB.Find(&tempperson, id).Error
}