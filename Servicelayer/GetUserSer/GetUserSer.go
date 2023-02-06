package GetUserSer

import (
	"example/persons/models"
	"strconv"
	"github.com/gin-gonic/gin"
	"example/persons/repository/GetUserRepo"
)
func GetUserSer(c *gin.Context)([]models.Person,error) {
	var tempperson []models.Person
	idStr:=c.DefaultQuery("id","1")
	var id int
	if idStr!=""{
		id, _ = strconv.Atoi(idStr)
	}
	err:=GetUserRepo.GetUserRepo(&tempperson,id)
	if err!=nil{
		return tempperson,err
	}
	return tempperson,nil
}