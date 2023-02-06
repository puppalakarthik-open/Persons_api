package CreateUser

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"example/persons/config"
	"example/persons/db"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	config.LoadConfig()
	db.Connection()
}

func TestCreateUserHandler(t *testing.T) {
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.POST("/", CreateUserHandler)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(`
	{
		"Name": "Kuradu",
		"ID": 37,
		"Gender": "M",
		"BloodType": "C+",
		"Num": 896050
	}
	`)))
	r.ServeHTTP(w, c.Request)
	assert.Equal(t, 200, w.Code)
}
