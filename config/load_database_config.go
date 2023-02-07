package config

import(
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type DbConfigStruct struct{
	
		Host   string	`json:"DB_HOST"`
		Port   string	`json:"DB_PORT"`
		User   string	`json:"DB_USER"`
		DBname string	`json:"DB_NAME"`
		Password string `json:"DB_PASSWORD"`
}
var Db *DbConfigStruct

func load_database_config(){
	Db=&DbConfigStruct{}
	err := envconfig.Process("db", Db)
	if err != nil {
		fmt.Println(err.Error())
	}
}