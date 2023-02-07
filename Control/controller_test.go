package control
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


func Test_Login(t *testing.T){
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.GET("/",Login)
	c.Request, _ = http.NewRequest(http.MethodGet, "/", bytes.NewBuffer([]byte(`
	{
	"admin",
	"password"
	}
	`)))
	r.ServeHTTP(w, c.Request)
	assert.Equal(t, 200, w.Code)
}