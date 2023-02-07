package get_users_repository
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

func Test_get_users_repository(t *testing.T) {
	var tempPerson []models.Person
	err := Get_users_repository(1,2,&tempPerson)
	assert.Empty(t, err)
}