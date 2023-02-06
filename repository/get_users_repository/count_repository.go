package get_users_repository
import(
	"example/persons/database"
	"example/persons/models"
)
func Count_repository(Totalrows *int64){
	database.DB.Model(&models.Person{}).Count(Totalrows)
}