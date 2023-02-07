package config
import(
	"log"
	"path/filepath"
	"runtime"
	"github.com/joho/godotenv"
	
)
func Load_config(){
	_,b,_,_:=runtime.Caller(0)
	ProjectRootPath:=filepath.Join(filepath.Dir(b),"../")
	err := godotenv.Load(ProjectRootPath+"/.env")
	if err != nil {
		log.Fatal(err)
	}
	load_application_config()
	load_authorization_config()
	load_database_config()
}