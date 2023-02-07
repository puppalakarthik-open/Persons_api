package get_user

import (
	"example/persons/service_layer/get_user_service"
	"example/persons/transformer"
	"example/persons/models"
	"example/persons/validation/valid_id"
	"github.com/gin-gonic/gin"
	"strconv"
)
func Get_user_handler(c *gin.Context) {
	//validation
	err:=valid_id.Valid_id(c)
	if err != nil {
		c.JSON(400, gin.H{"Message": "User fetch unsuccessfully"})
		return
	}
	//service
	var tempPerson []models.Person
	idStr:=c.DefaultQuery("id","1")
	var id int
	if idStr!=""{
		id, _ = strconv.Atoi(idStr)
	}
	tempdata, err1:=get_user_service.Get_user_service(tempPerson,id)
	if err1 != nil {
		c.JSON(400, gin.H{"Message": "User fetch unsuccessfully"})
		return
	}
	//transformer
	tempdata1:=transformer.Transform(tempdata)
	c.JSON(200, gin.H{"Message": "User fetched successfully", "data": tempdata1})
	}
