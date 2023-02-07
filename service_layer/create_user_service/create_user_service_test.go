package create_user_service

import (
	"example/persons/config"
	"example/persons/database"
	"example/persons/models"
	"testing"
	"github.com/stretchr/testify/assert"
)

func init() {
	config.Load_config()
	database.Connection()
}

func TestCreateUsersSer(t *testing.T) {
person:= models.Person{
	Name: "Kuradu",
	ID: 93,
	Gender: "M",
	BloodType: "C+",
	Num: 896050,	
}
err := Create_user_service(person)
assert.Empty(t, err)
}
