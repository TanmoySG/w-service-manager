package checks

import (
	"validity/pkg/ping"
)

func (c Client) CheckRepositoryStatus(repositoryURL string) bool {
	statusCode, _ := ping.Ping(repositoryURL, ping.GET, nil)

	statusOk := statusCode >= 200 && statusCode < 300

	if !statusOk {
		return false
	} else {
		return true
	}
}
