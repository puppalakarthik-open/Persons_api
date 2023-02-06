package middle

import (
	"example/persons/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	k:=c.GetHeader("Authorization")
	clientToken:=strings.TrimPrefix(k,"Bearer ")
	if _, err := jwt.Validatetoken(clientToken); err != "" {  			
		c.JSON(401, gin.H{"message": err,"Msg":"error in validation"})
		c.Abort()
	}
	c.Next()
}
