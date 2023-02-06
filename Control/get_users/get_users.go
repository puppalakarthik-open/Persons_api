package get_users

import (
	"example/persons/transformer"
	"example/persons/service_layer/get_users_service"
	"github.com/gin-gonic/gin"
)
func Get_users_handler(c *gin.Context) {
	//service
	tempdata,err:=get_users_service.Get_users_service(c)
	if err != nil {
			c.JSON(400, gin.H{"Message": "User fetch unsuccessfully"})
			return
		}
	//transformer
	tempdata1:=transformer.Transform(tempdata)
	c.JSON(200, gin.H{"Message": "User fetched successfully", "data": tempdata1})
}
