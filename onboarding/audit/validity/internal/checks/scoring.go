package checks

import (
	"validity/spec/validity"
)

const (
	NameValidityScore       = 25
	RepositoryValidityScore = 25
	AccessValidityScore     = 50

	MinimumThresholdScore = 75
)

func (c Client) GetServiceValidationScore(nameValidity, repositoryValidity, accessValidity validity.Validations) int {
	serviceValidationScore := 0

	if nameValidity.Valid {
		serviceValidationScore = serviceValidationScore + NameValidityScore
	}

	if repositoryValidity.Valid {
		serviceValidationScore = serviceValidationScore + RepositoryValidityScore
	}

	if accessValidity.Valid {
		serviceValidationScore = serviceValidationScore + AccessValidityScore
	}

	return serviceValidationScore
}
