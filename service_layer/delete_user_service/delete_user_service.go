package delete_user_service
import (
	"example/persons/repository/delete_user_repository"
	"example/persons/models"
)
func Delete_user_service(personModels *[]models.Person,id int)error{
	return delete_user_repository.Delete_user_repository(personModels,id)
}