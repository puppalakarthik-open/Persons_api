package get_user_repository
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

func Test_get_user_repository(t *testing.T) {
	var tempPerson []models.Person
	err := Get_user_repository(&tempPerson,11)
	assert.Empty(t, err)
}