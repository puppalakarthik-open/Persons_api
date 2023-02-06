package create_user

import (
	"example/persons/validation"
	"example/persons/service_layer/create_user_service"
	"github.com/gin-gonic/gin"
)

func Create_user_handler(c *gin.Context) {
	//Validation
	d, err := validation.Valid(c)
	if err != nil {
		c.JSON(400, gin.H{"Message": "entry isn't valid"})
		return
	}
	//service
	err = create_user_service.Create_user_service(d)

	if err != nil {
		c.JSON(400, gin.H{"Message": "entry isn't posted"})
		return
	}
	c.JSON(200, gin.H{"Message": "User created successfully"})
	// transformer

}
