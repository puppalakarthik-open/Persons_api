package DeleteUserSer
import (
	"example/persons/repository/DeleteUserRepo"
	"strconv"
	"example/persons/models"
	"github.com/gin-gonic/gin"
)
func DeleteUserSer(c *gin.Context)error{
	personModels := &[]models.Person{}
	idStr:=c.DefaultQuery("id","1")
	var id int
	if idStr!=""{
		id, _ = strconv.Atoi(idStr)
	}
	return DeleteUserRepo.DeleteUserRepo(personModels,id)
}