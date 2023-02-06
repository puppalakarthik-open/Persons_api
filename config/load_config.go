package config
import(
	// "os"
	"log"
	"path/filepath"
	"runtime"
	"github.com/joho/godotenv"
	
)
func LoadConfig(){
	_,b,_,_:=runtime.Caller(0)
	ProjectRootPath:=filepath.Join(filepath.Dir(b),"../")
	err := godotenv.Load(ProjectRootPath+"/.env")
	// curPath, _ := os.Getwd()
	// err := godotenv.Load(curPath+"/.env")
	if err != nil {
		log.Fatal(err)
	}
	loadAppConfig()
	loadAuthConfig()
	loadDbConfig()
}