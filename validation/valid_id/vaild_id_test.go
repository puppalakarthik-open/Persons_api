package valid_id

import(
	"testing"
	"example/persons/config"
	"example/persons/database"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
)

func init() {
	config.Load_config()
	database.Connection()
}


func Test_valid_id(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	err:=Valid_id(c)
	assert.Empty(t,err)
}