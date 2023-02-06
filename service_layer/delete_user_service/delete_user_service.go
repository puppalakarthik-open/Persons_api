package delete_user_service
import (
	"example/persons/repository/delete_user_repository"
	"strconv"
	"example/persons/models"
	"github.com/gin-gonic/gin"
)
func Delete_user_service(c *gin.Context)error{
	personModels := &[]models.Person{}
	idStr:=c.DefaultQuery("id","1")
	var id int
	if idStr!=""{
		id, _ = strconv.Atoi(idStr)
	}
	return delete_user_repository.Delete_user_repository(personModels,id)
}