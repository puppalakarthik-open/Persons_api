package create_user

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"example/persons/config"
	"example/persons/database"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	config.Load_config()
	database.Connection()
}

func TestCreateUserHandler(t *testing.T) {
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.POST("/", Create_user_handler)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(`
	{
		"Name": "Kuradu",
		"ID": 97,
		"Gender": "M",
		"Blood_Type": "C+",
		"Num": 896050
	}
	`)))
	r.ServeHTTP(w, c.Request)
	assert.Equal(t, 200, w.Code)
}
