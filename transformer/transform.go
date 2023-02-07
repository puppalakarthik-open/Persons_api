package transformer

import (
	"example/persons/models"
)

func Transform(alpha []models.Person) []models.PersonResponse {
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
