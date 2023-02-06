package create_user_repository
import (
	"example/persons/models"
	"example/persons/database"
	)
func Create_user_repository(alpha models.Person)error{
	err1 := database.DB.Create(&alpha).Error
	if err1 != nil {
		return err1
	}
	return nil
}