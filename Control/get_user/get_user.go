package get_user

import (
	"example/persons/service_layer/get_user_service"
	"example/persons/transformer"
	"github.com/gin-gonic/gin"
	
)
func Get_user_handler(c *gin.Context) {
	//validation
	//service
	tempdata, err1:=get_user_service.Get_user_service(c)
	if err1 != nil {
		c.JSON(400, gin.H{"Message": "User fetch unsuccessfully"})
		return
	}
	//transformer
	tempdata1:=transformer.Transform(tempdata)
	c.JSON(200, gin.H{"Message": "User fetched successfully", "data": tempdata1})
	}
