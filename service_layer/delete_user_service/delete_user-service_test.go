package delete_user_service

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
func Test_delete_user_service(t *testing.T){
	personModels := &[]models.Person{}
	err := Delete_user_service(personModels,92)
	assert.Empty(t, err)
}