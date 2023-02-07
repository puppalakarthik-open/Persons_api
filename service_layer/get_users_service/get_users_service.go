package get_users_service

import (
	"example/persons/models"
	"example/persons/repository/get_users_repository"
)
func Get_users_service(perPage int,offset int)([]models.Person,error){

	personModels := []models.Person{}
	err := get_users_repository.Get_users_repository(perPage,offset,&personModels)
	if err != nil {
		return personModels,err
	}
	return personModels,nil
}