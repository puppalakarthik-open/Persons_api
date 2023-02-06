package GetUser

import (
	"example/persons/Servicelayer/GetUserSer"
	"example/persons/transformer"
	"github.com/gin-gonic/gin"
	
)
func GetUserHandler(c *gin.Context) {
	//service
	tempdata, err1:=GetUserSer.GetUserSer(c)
	if err1 != nil {
		c.JSON(400, gin.H{"Message": "User fetch unsuccessfully"})
		return
	}
	//transformer
	tempdata1:=transformer.Transform(tempdata)
	c.JSON(200, gin.H{"Message": "User fetched successfully", "data": tempdata1})
	}
