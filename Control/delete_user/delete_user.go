package delete_user

import (
	"example/persons/service_layer/delete_user_service"
	"github.com/gin-gonic/gin"
	"example/persons/models"
	"strconv"
)
func Delete_user_handler(c *gin.Context) {
	personModels := &[]models.Person{}
	idStr:=c.DefaultQuery("id","1")
	var id int
	if idStr!=""{
		id, _ = strconv.Atoi(idStr)
	}
	err:=delete_user_service.Delete_user_service(personModels,id)
	if err != nil {
		c.JSON(400, gin.H{"Message": "Can't delete"})
		return
	}
	c.JSON(200, gin.H{"Message": "User deleted successfully"})
}