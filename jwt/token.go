package jwt

import (

	"example/persons/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(validperson models.PersonRequest) (string, error) {
	type MyCustomClaims struct {
		user string 
		jwt.StandardClaims
	}
	claims := MyCustomClaims{
		validperson.UserName,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret"))
}

func Validatetoken(tokenString string) (bool, string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return false, err.Error()
	}
	if token.Valid{
		return true,""
	}
	return false,"Token not valid"
}
