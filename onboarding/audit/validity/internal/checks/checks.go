package checks

import (
	"validity/pkg/wdb"
	"validity/spec/contract"
	"validity/spec/validity"
)

type Client struct {
	WDBClient   wdb.Client
	ControlList string
}

type CheckFlag struct {
	Flag    bool
	Message *string
}

var (
	validityReason = "Valid" // Need to be a bit more creative :)

	Invalid = CheckFlag{
		Flag: false,
	}

	Valid = CheckFlag{
		Flag:    true,
		Message: &validityReason,
	}
)

func NewChecksClient(client wdb.Client, serviceDirectory string, controlList string) Client {
	wdbClient := wdb.Client{
		Cluster:  client.Cluster,
		Token:    client.Token,
		Database: &serviceDirectory,
	}

	return Client{
		WDBClient:   wdbClient,
		ControlList: controlList,
	}
}

func (c Client) GetContractValidity(contract contract.Contract) validity.Checks {

	overallContractValidity := Valid.Flag

	nameValidation := c.CheckNameExistenceInServicesDirectory(contract.Service.Name)
	repositoryValidation := c.CheckRepositoryStatus(*contract.Service.Repository)
	accessValidation := c.CheckAssignedDataAccess(contract.Data)

	serviceValidationScore := c.GetServiceValidationScore(nameValidation, repositoryValidation, accessValidation)

	if serviceValidationScore < MinimumThresholdScore {
		overallContractValidity = Invalid.Flag
	}

	checks := validity.Checks{
		Valid:       &overallContractValidity,
		DataAccess:  &accessValidation,
		ServiceName: &nameValidation,
		Repository:  &repositoryValidation,
	}

	return checks
}
