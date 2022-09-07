package checks

import (
	"validity/pkg/ping"
	"validity/spec/validity"
)

var RepositoryInvalidityReason = "Repository Doesn't exist or is private."

func (c Client) CheckRepositoryStatus(repositoryURL string) validity.Validations {
	var overallValidity bool = Valid.Flag
	var overallValidityReason string = *Valid.Message

	statusCode, _ := ping.Ping(repositoryURL, ping.GET, nil)

	statusOk := statusCode >= 200 && statusCode < 300

	if !statusOk {
		overallValidity = Invalid.Flag
		overallValidityReason = RepositoryInvalidityReason
	}

	repoValidity := validity.Validations{
		Error: overallValidityReason,
		Valid: overallValidity,
	}

	return repoValidity
}
