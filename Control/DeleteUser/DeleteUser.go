package DeleteUser

import (
	"example/persons/Servicelayer/DeleteUserSer"
	"github.com/gin-gonic/gin"
)
func DeleteUserHandler(c *gin.Context) {
	err:=DeleteUserSer.DeleteUserSer(c)
	if err != nil {
		c.JSON(400, gin.H{"Message": "Can't delete"})
		return
	}
	c.JSON(200, gin.H{"Message": "User deleted successfully"})
}