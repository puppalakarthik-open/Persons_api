package config

import (
	"example/persons/models"
	"github.com/kelseyhightower/envconfig"
	"log"
)
var AuthorizedPerson *models.PersonRequest

func load_authorization_config() {
	AuthorizedPerson = &models.PersonRequest{}
	err := envconfig.Process("CRED", AuthorizedPerson)
	
	if err != nil {
		log.Fatal(err.Error())
	}
}
