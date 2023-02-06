package CreateUser

import (
	"example/persons/validation"
	"example/persons/Servicelayer/CreateUserSer"
	"github.com/gin-gonic/gin"
)

func CreateUserHandler(c *gin.Context) {
	//Validation
	d, err := validation.Valid(c)
	if err != nil {
		c.JSON(400, gin.H{"Message": "entry isn't valid"})
		return
	}
	//service
	err = CreateUserSer.CreateUserSer(d)

	if err != nil {
		c.JSON(400, gin.H{"Message": "entry isn't posted"})
		return
	}
	c.JSON(200, gin.H{"Message": "User created successfully"})
	// transformer

}
