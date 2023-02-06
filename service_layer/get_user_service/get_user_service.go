package get_user_service

import (
	"example/persons/models"
	"strconv"
	"github.com/gin-gonic/gin"
	"example/persons/repository/get_user_repository"
)
func Get_user_service(c *gin.Context)([]models.Person,error) {
	var tempperson []models.Person
	idStr:=c.DefaultQuery("id","1")
	var id int
	if idStr!=""{
		id, _ = strconv.Atoi(idStr)
	}
	err:=get_user_repository.Get_user_repository(&tempperson,id)
	if err!=nil{
		return tempperson,err
	}
	return tempperson,nil
}