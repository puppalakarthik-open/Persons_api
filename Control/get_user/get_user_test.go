package get_user
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
	config.Load_config()
	database.Connection()
}


func TestGetUserHandler(t *testing.T) {
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.GET("/",Get_user_handler)
	c.Request, _ = http.NewRequest(http.MethodGet, "/", bytes.NewBuffer([]byte(``)))
	r.ServeHTTP(w, c.Request)
	assert.Equal(t, 200, w.Code)
}