package create_user_service

import (
	"example/persons/config"
	"example/persons/database"
	"example/persons/models"
	"testing"
	// "github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	config.LoadConfig()
	database.Connection()
}

func TestCreateUsersSer(t *testing.T) {
person_req := models.Person{
	Name: "Kuradu",
	ID: 37,
	Gender: "M",
	BloodType: "C+",
	Num: 896050,	
}
err := Create_user_service(person_req)
assert.Empty(t, err)
}
