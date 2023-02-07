package validate

import (
    "bytes"
    "encoding/json"
    "example/persons/config"
    "example/persons/database"
    "example/persons/models"
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func init() {
    config.Load_config()
    database.Connection()
}


func Test_valid(t *testing.T) {
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    person := models.Person{
        Name: "name",
        Gender:"M",
        BloodType:"A+",
        ID:43,
        Num:213321,
    }
    persondata, _ := json.Marshal(&person)
    c.Request = &http.Request{
        Header: make(http.Header),
        Body:   ioutil.NopCloser(bytes.NewBuffer(persondata)),
    }
    res, err := Valid(c)
    assert.Empty(t, err)
    assert.NotEmpty(t, res)
}
    
