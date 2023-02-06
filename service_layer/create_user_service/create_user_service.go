package create_user_service
import (
"example/persons/models"
"example/persons/repository/create_user_repository"
)
func Create_user_service(alpha models.Person)error{
	return create_user_repository.Create_user_repository(alpha)
}