package transformer

import (
	"example/persons/models"
)

func Transform(person []models.Person) []models.PersonResponse {
	user := make([]models.PersonResponse, len(person))
	for i, person := range person {
		user[i] = models.PersonResponse{
			Name:      person.Name,
			BloodType: person.BloodType,
			Gender:    person.Gender,
			ID:        person.ID,
			Num:       person.Num,
		}
	}
	return user

}
