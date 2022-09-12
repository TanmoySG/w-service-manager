// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    validity, err := UnmarshalValidity(bytes)
//    bytes, err = validity.Marshal()

package validity

import (
	"encoding/json"
	"validity/spec/contract"
)

func UnmarshalValidity(data []byte) (Validity, error) {
	var r Validity
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Validity) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Contract Validation Schema for Audit
type Validity struct {
	Checks      Checks            `json:"checks"`
	Contract    contract.Contract `json:"contract"`
	ContractID  string            `json:"contract_id"` // Service Contract ID
	Kind        string            `json:"kind"`
	RequestID   string            `json:"request_id"`   // Service Onboarding Request ID
	ServiceName string            `json:"service_name"` // Name of the Service. Should be URL/Computer Friendly - words seperated by dot (.),; hyphens(-), underscore (_) and in smaller case.
}

// Validity Checks
type Checks struct {
	DataAccess  *Validations `json:"data_access,omitempty"`
	Repository  *Validations `json:"repository,omitempty"`
	ServiceName *Validations `json:"service_name,omitempty"`
	Valid       *bool        `json:"valid,omitempty"`
}

type Validations struct {
	Error              string               `json:"error"` // Error, if valid : false, otherwise null
	FieldLevelValidity []FieldLevelValidity `json:"field_level_validity,omitempty"`
	Valid              bool                 `json:"valid"`
}

type FieldLevelValidity struct {
	Error *string `json:"error,omitempty"`
	Field *string `json:"field,omitempty"`
	Valid *bool   `json:"valid,omitempty"`
}
