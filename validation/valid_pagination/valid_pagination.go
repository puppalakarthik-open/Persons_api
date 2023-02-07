package valid_pagination

import (
	"errors"
	"example/persons/models"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)
func Valid_pagination(c *gin.Context)error{
	var pages models.Pagination
	pages.Page = c.DefaultQuery("page","11")
	pages.Limit = c.DefaultQuery("limit","11")

	rules := govalidator.MapData{
		"page":        []string{"min:1", "max:2"},
		"limit":        []string{"min:1", "max:3"},
	}
	messages := govalidator.MapData{
		"page":        []string{"required:Id is must", "min:minimum 1 character", "max:Takes upto3 characters only"},
		"limit":        []string{"required:Id is must", "min:minimum 1 character", "max:Takes upto3 characters only"},
	}

	opts := govalidator.Options{
		Data:     &pages,
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