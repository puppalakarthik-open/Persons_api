package create_user_repository
import(
	"testing"
	"example/persons/config"
	"example/persons/database"
	"example/persons/models"
	"github.com/stretchr/testify/assert"
)

func init() {
	config.Load_config()
	database.Connection()
}

func Test_create_user_repository(t *testing.T) {
	person := models.Person{
		Name: "Kuradu",
		ID: 94,
		Gender: "M",
		BloodType: "C+",
		Num: 896050,	
	}
	err := Create_user_repository(person)
	assert.Empty(t, err)
}