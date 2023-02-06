package Control

import (

	"example/persons/jwt"
	"example/persons/models"
	"strings"

	"encoding/base64"
	"example/persons/config"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	encoded_things := strings.SplitN(c.GetHeader("Authorization"), " ", 2)[1]
	b, _ := base64.StdEncoding.DecodeString(encoded_things)
	credentials := strings.SplitN(string(b), ":", 2)
	validperson := models.PersonRequest{UserName:credentials[0],Password:credentials[1]}
	if validperson == *config.AuthorizedPerson {
			c.Next()
	} else {
		c.JSON(401, gin.H{"error": "Wrong Username or password"})
		c.Abort()
		return
	}
	alpha,err:=jwt.GenerateToken(validperson)
	if err!=nil{
		c.JSON(401,gin.H{"Message":"Unauthorized"})
		c.Abort()
	}
	c.JSON(200,gin.H{"Message":alpha})
}
