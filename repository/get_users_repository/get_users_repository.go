package get_users_repository

import (
	"example/persons/database"
	"example/persons/models"
)
func Get_users_repository(perPage int,offset int,personModels *[]models.Person)(error){
	return database.DB.Limit(perPage).Offset(offset).Find(personModels).Error
}