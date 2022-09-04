package checks

import (
	"validity/pkg/ping"
	"validity/spec/validity"
)

var RepositoryInvalidityReason = "Repository Doesn't exist or is private."

func (c Client) CheckRepositoryStatus(repositoryURL string) validity.Validations {
	var overallValidity bool = Valid
	var overallValidityReason string = "Valid"

	statusCode, _ := ping.Ping(repositoryURL, ping.GET, nil)

	statusOk := statusCode >= 200 && statusCode < 300

	if !statusOk {
		overallValidity = Invalid
		overallValidityReason = RepositoryInvalidityReason
	}

	repoValidity := validity.Validations{
		Error: overallValidityReason,
		Valid: overallValidity,
	}

	return repoValidity
}
