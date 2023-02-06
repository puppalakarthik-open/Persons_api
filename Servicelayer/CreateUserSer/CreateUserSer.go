package CreateUserSer
import (
"example/persons/models"
"example/persons/repository/CreateUserRepo"
)
func CreateUserSer(alpha models.Person)error{
	return CreateUserRepo.CreateUserRepo(alpha)
}