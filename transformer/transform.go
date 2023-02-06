package transformer

import (
	"example/persons/models"
	"fmt"
)

func Transform(alpha []models.Person) []models.PersonResponse {
	fmt.Println("in tranform layer")
	beta := make([]models.PersonResponse, len(alpha))
	for i, alpha := range alpha {
		beta[i] = models.PersonResponse{
			Name:      alpha.Name,
			BloodType: alpha.BloodType,
			Gender:    alpha.Gender,
			ID:        alpha.ID,
			Num:       alpha.Num,
		}
	}
	return beta

}
