package GetUsersRepo
import(
	"example/persons/db"
	"example/persons/models"
)
func CountRepo(Totalrows *int64){
	db.DB.Model(&models.Person{}).Count(Totalrows)
}