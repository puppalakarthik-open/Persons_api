package CreateUserSer

import (
	"example/persons/config"
	"example/persons/db"
	"example/persons/models"
	"testing"
	// "github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	config.LoadConfig()
	db.Connection()
}

func TestCreateUsersSer(t *testing.T) {
person_req := models.Person{
	Name: "Kuradu",
	ID: 37,
	Gender: "M",
	BloodType: "C+",
	Num: 896050,
}
err := CreateUserSer(person_req)
assert.Empty(t, err)
}
