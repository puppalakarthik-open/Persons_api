package valid_pagination

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


func Test_valid_pagination(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	err:=Valid_pagination(c)
	assert.Empty(t,err)
}