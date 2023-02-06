package validation

import (
	"errors"
	"example/persons/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

func Valid(c *gin.Context) (models.Person,error) {
	var per models.Person
	err := c.BindJSON(&per)
	fmt.Println(per)
	if err != nil {
		c.JSON(400, gin.H{"message": "binding in validator"})
		return per, err
	}

	rules := govalidator.MapData{
		"name":      []string{"required", "between:3,10"},
		"id":        []string{"min:1", "max:100"},
		"gender":    []string{"required", "in:M,F"},
		"bloodtype": []string{"required", "min:1", "max:10"},
		"num":       []string{"min:10", "max:10000000000"},
	}
	messages := govalidator.MapData{
		"name":      []string{"required:Name is required", "between: name should be between 3 and 5 characters"},
		"id":        []string{"required:Id is must", "min:minimum 1 character", "max:Takes upto3 characters only"},
		"gender":    []string{"required:Gender must be filled", "min:minimum 1 character required"},
		"bloodtype": []string{"required:please enter the blood type", "min:minimum 2 character required", "max:Takes upto3 characters only"},
		"num":       []string{"min:Takes 10 characters only", "max:Takes 10 characters only"},
	}

	opts := govalidator.Options{
		Data:     &per,
		Messages: messages,
		Rules:    rules,
	}
	fmt.Println("in validator function")
	v := govalidator.New(opts)
	e := v.ValidateStruct()
	fmt.Println(e)

	if len(e) > 0 {
		c.JSON(400, e)
		return per, errors.New("invalid struct")
	}
	return per, nil
}