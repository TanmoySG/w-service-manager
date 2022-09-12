package checks

import (
	"validity/internal/control-list"
	"validity/spec/contract"
	"validity/spec/validity"
)

var AccessInvalidityReason = "Unauthorized Access Assigned"

func (c Client) CheckAssignedDataAccess(assignedAccess []contract.Datum) validity.Validations {

	cl, _ := controlList.LoadControlList(c.ControlList)

	var fieldLevelAccessValidity []validity.FieldLevelValidity

	var overallDataAccessValidity bool = Valid.Flag
	var overallDataAccessValidityReason string = *Valid.Message

	for _, data := range assignedAccess {
		if cl.ValidateAccessForField(data.Data, data.Access) {
			fieldAccess := validity.FieldLevelValidity{
				Error: nil,
				Field: &data.Data,
				Valid: &Valid.Flag,
			}
			fieldLevelAccessValidity = append(fieldLevelAccessValidity, fieldAccess)
		} else {
			fieldAccess := validity.FieldLevelValidity{
				Error: &AccessInvalidityReason,
				Field: &data.Data,
				Valid: &Invalid.Flag,
			}
			fieldLevelAccessValidity = append(fieldLevelAccessValidity, fieldAccess)
			overallDataAccessValidity = Invalid.Flag
			overallDataAccessValidityReason = AccessInvalidityReason
		}
	}

	dataAccessValidity := validity.Validations{
		Error:              overallDataAccessValidityReason,
		FieldLevelValidity: fieldLevelAccessValidity,
		Valid:              overallDataAccessValidity,
	}

	return dataAccessValidity
}
