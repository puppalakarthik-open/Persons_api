package GetUsers
import(
	"testing"
	"bytes"
	"net/http"
	"example/persons/config"
	"example/persons/db"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadConfig()
	db.Connection()
}

func TestGetUsersHandler(t *testing.T) {
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.GET("/",GetUsersHandler)
	c.Request, _ = http.NewRequest(http.MethodGet, "/", bytes.NewBuffer([]byte(``)))
	r.ServeHTTP(w, c.Request)
	assert.Equal(t, 200, w.Code)
}
