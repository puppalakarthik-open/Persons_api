package valid_id

import (
	"errors"
	"example/persons/models"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)
func Valid_id(c *gin.Context)error{
	var id models.Id
	id.Id = c.DefaultQuery("id","11")
	

	rules := govalidator.MapData{
		"id":        []string{"min:1", "max:2"},
	}
	messages := govalidator.MapData{
		"id":        []string{"required:Id is must", "min:minimum 1 character", "max:Takes upto3 characters only"},
		}

	opts := govalidator.Options{
		Data:     &id,
		Messages: messages,
		Rules:    rules,
	}
	v := govalidator.New(opts)
	e := v.ValidateStruct()

	if len(e) > 0 {
		c.JSON(400, e)
		return errors.New("invalid struct")
	}
	return nil
}