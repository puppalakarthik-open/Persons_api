package GetUsers

import (
	"example/persons/transformer"
	"example/persons/Servicelayer/GetUsersSer"
	"github.com/gin-gonic/gin"
)
func GetUsersHandler(c *gin.Context) {
	//service
	tempdata,err:=GetUsersSer.GetUsersSer(c)
	if err != nil {
			c.JSON(400, gin.H{"Message": "User fetch unsuccessfully"})
			return
		}
	//transformer
	tempdata1:=transformer.Transform(tempdata)
	c.JSON(200, gin.H{"Message": "User fetched successfully", "data": tempdata1})
}
