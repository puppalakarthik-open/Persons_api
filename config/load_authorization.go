package config

import (
	"example/persons/models"
	"github.com/kelseyhightower/envconfig"
	"fmt"
)
var AuthorizedPerson *models.PersonRequest

func loadAuthConfig() {
	AuthorizedPerson = &models.PersonRequest{}
	err := envconfig.Process("CRED", AuthorizedPerson)
	
	if err != nil {
		fmt.Println(err.Error())
	}
}
