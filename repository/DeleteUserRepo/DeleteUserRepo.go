package DeleteUserRepo
import(
	"example/persons/db"
	"example/persons/models"
)
func DeleteUserRepo(alpha *[]models.Person,id int)error{
	return db.DB.Delete(alpha, id).Error
}
