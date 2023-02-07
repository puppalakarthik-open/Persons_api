package get_user_service
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

func Test_get_user_service(t *testing.T) {
	var tempPerson []models.Person
	temPerson,err := Get_user_service(tempPerson,11)
	assert.NotEmpty(t,temPerson)
	assert.Empty(t, err)
}