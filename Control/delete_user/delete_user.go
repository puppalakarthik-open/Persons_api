package delete_user

import (
	"example/persons/service_layer/delete_user_service"
	"github.com/gin-gonic/gin"
)
func Delete_user_handler(c *gin.Context) {
	err:=delete_user_service.Delete_user_service(c)
	if err != nil {
		c.JSON(400, gin.H{"Message": "Can't delete"})
		return
	}
	c.JSON(200, gin.H{"Message": "User deleted successfully"})
}