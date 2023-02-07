package delete_user_repository

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
func Test_delete_user_repository(t *testing.T){
	personModels := &[]models.Person{}
	err := Delete_user_repository(personModels,94)
	assert.Empty(t, err)
}