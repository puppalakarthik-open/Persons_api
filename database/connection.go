package database
import(
	"fmt"
	"example/persons/config"
	"example/persons/models"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"log"
)
var DB *gorm.DB

func Connection() {
	dsn := "host=" + config.Db.Host +
			" user=" + config.Db.User +
			" password=" + config.Db.Password +
			" dbname=" + config.Db.DBname +
			" port=" + config.Db.Port +
			" sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	err = DB.AutoMigrate(&models.Person{})
	if err != nil {
			log.Fatal("Couldn't migrate db")
		}
}