package get_users

import (
	"example/persons/service_layer/get_users_service"
	"example/persons/transformer"
	"example/persons/validation/valid_pagination"
	"example/persons/pagination"
	"github.com/gin-gonic/gin"
)
func Get_users_handler(c *gin.Context) {
	//Validation
	err:=valid_pagination.Valid_pagination(c)
	if err != nil {
		c.JSON(400, gin.H{"Message": "User fetch unsuccessfully"})
		return
	}
	//service
	perPage,offset,err:=pagination.Pagination(c)
	if err!=nil{
		return 
	}
	tempdata,err:=get_users_service.Get_users_service(perPage,offset)
	if err != nil {
			c.JSON(400, gin.H{"Message": "User fetch unsuccessfully"})
			return
		}
	//transformer
	tempdata1:=transformer.Transform(tempdata)
	c.JSON(200, gin.H{"Message": "User fetched successfully", "data": tempdata1})
}
