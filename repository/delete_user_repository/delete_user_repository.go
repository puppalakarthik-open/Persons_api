package delete_user_repository
import(
	"example/persons/database"
	"example/persons/models"
)
func Delete_user_repository(alpha *[]models.Person,id int)error{
	return database.DB.Delete(alpha, id).Error
}
