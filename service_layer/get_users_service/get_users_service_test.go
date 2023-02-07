package get_users_service
import(
	"testing"
	"example/persons/config"
	"example/persons/database"
	"github.com/stretchr/testify/assert"
)

func init() {
	config.Load_config()
	database.Connection()
}

func Test_get_users_service(t *testing.T) {
	temPerson,err := Get_users_service(3,2)
	assert.NotEmpty(t,temPerson)
	assert.Empty(t, err)
}