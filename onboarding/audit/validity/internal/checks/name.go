package checks

import (
	"validity/internal/service-directory"

	"golang.org/x/exp/slices"
)

func (c Client) CheckNameExistenceInServicesDirectory(name string) bool {
	sd := serviceDirectory.MockClient{
		WDBClient:  c.WDBClient,
		Collection: "servicesDirectory",
	}

	servicesList := sd.GetServiceNameList()
	return !slices.Contains(*servicesList, name)
}
