package checks

import (
	serviceDirectory "validity/internal/service-directory"
	"validity/spec/validity"

	"golang.org/x/exp/slices"
)

var NameInvalidityReason = "Name Already Exists."

func (c Client) CheckNameExistenceInServicesDirectory(name string) validity.Validations {
	var overallValidity bool = Valid.Flag
	var overallValidityReason string = *Valid.Message

	sd := serviceDirectory.MockClient{
		WDBClient:  c.WDBClient,
		Collection: "servicesDirectory",
	}
	servicesList := sd.GetServiceNameList()

	if slices.Contains(*servicesList, name) {
		overallValidity = Invalid.Flag
		overallValidityReason = NameInvalidityReason
	}

	nameValidity := validity.Validations{
		Error: overallValidityReason,
		Valid: overallValidity,
	}

	return nameValidity
}
