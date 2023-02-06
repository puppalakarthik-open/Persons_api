package GetUsersRepo

import (
	"example/persons/db"
	"example/persons/models"
	"fmt"
)
func GetUsersRepo(perPage int,offset int,personModels *[]models.Person)(error){
	fmt.Println("In repo")
	return db.DB.Limit(perPage).Offset(offset).Find(personModels).Error
}