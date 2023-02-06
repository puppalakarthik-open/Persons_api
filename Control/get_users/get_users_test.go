package get_users
import(
	"testing"
	"bytes"
	"net/http"
	"example/persons/config"
	"example/persons/database"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadConfig()
	database.Connection()
}

func Test_Get_users_handler(t *testing.T) {
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.GET("/",Get_users_handler)
	c.Request, _ = http.NewRequest(http.MethodGet, "/", bytes.NewBuffer([]byte(``)))
	r.ServeHTTP(w, c.Request)
	assert.Equal(t, 200, w.Code)
}
