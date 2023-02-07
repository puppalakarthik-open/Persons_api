package get_user_service

import (
	"example/persons/models"
	"example/persons/repository/get_user_repository"
)
func Get_user_service(tempPerson []models.Person,id int)([]models.Person,error) {
	err:=get_user_repository.Get_user_repository(&tempPerson,id)
	if err!=nil{
		return tempPerson,err
	}
	return tempPerson,nil
}