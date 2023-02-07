package config
import(
	"log"
	"github.com/kelseyhightower/envconfig"
)
type AppconfigStruct struct{
	Env			 string 		`json:"APP_ENV"`
	Name 		 string 		`json:"APP_NAME"`
	Debug 		 bool 			`json:"APP_DEBUG"`
	Port 		 string 		`json:"APP_PORT"`
	Maintenance  bool			`json:"APP_MAINTENANCE"`
}
var AppConfig *AppconfigStruct

func load_application_config() {
	AppConfig = &AppconfigStruct{}
	err := envconfig.Process("app", AppConfig)
	if err != nil {
		log.Fatal(err.Error())
	}
}



