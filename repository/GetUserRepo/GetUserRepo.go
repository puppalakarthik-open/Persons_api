package GetUserRepo

import (
	"example/persons/db"
	"example/persons/models"
)
func GetUserRepo(tempperson *[]models.Person, id int)(error) {
	return db.DB.Find(&tempperson, id).Error
	
}