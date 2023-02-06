package CreateUserRepo
import (
	"example/persons/models"
	"example/persons/db"
	)
func CreateUserRepo(alpha models.Person)error{
	err1 := db.DB.Create(&alpha).Error
	if err1 != nil {
		return err1
	}
	return nil
}