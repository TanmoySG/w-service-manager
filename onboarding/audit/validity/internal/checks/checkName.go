package checks

import (
	// "fmt"
	"validity/internal/service-directory"
)

func (c Client) CheckNameExistenceInServicesDirectory(name string) (bool, error) {
	sd := serviceDirectory.MockClient{
		WDBClient: c.WDBClient ,
		Collection: "servicesDirectory",
	}
	
	sd.GetServiceNameList()

	return true, nil
}
