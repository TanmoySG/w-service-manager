// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    validity, err := UnmarshalValidity(bytes)
//    bytes, err = validity.Marshal()

package main

import "encoding/json"

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
	Checks      Checks   `json:"checks"`      
	Contract    Contract `json:"contract"`    
	ContractID  string   `json:"contract_id"` // Service Contract ID
	Kind        string   `json:"kind"`        
	RequestID   string   `json:"request_id"`  // Service Onboarding Request ID
	ServiceName string   `json:"service_name"`// Name of the Service. Should be URL/Computer Friendly - words seperated by dot (.),; hyphens(-), underscore (_) and in smaller case.
}

// Validity Checks
type Checks struct {
	DataAccess  *Validations `json:"data_access,omitempty"` 
	Repository  *Validations `json:"repository,omitempty"`  
	ServiceName *Validations `json:"service_name,omitempty"`
}

type Validations struct {
	Error              string               `json:"error"`                         // Error, if valid : false, otherwise null
	FieldLevelValidity []FieldLevelValidity `json:"field_level_validity,omitempty"`
	Valid              bool                 `json:"valid"`                         
}

type FieldLevelValidity struct {
	Error *string `json:"error,omitempty"`
	Field *string `json:"field,omitempty"`
	Valid *bool   `json:"valid,omitempty"`
}

// Service Onboarding Contract.
type Contract struct {
	Data      []Datum   `json:"data"`      // Information about Data Required and Usage
	Developer Developer `json:"developer"` // Information about the App/Servie Owner/Representative
	Kind      string    `json:"kind"`      
	RequestID string    `json:"request_id"`// Service Onboarding
	Service   Service   `json:"service"`   // Information about Service
}

type Datum struct {
	Access []Access `json:"access"`// Access type
	Data   string   `json:"data"`  // Name of the Data
	Use    string   `json:"use"`   // Usage Information
}

// Information about the App/Servie Owner/Representative
type Developer struct {
	Admin       []string `json:"admin"`                // Service Admin - An Overall Service Admin Mail
	Contributor []string `json:"contributor,omitempty"`// Service Contributor - A group Mail ID or DL for Contibutors
}

// Information about Service
type Service struct {
	Details    []string `json:"details"`             // Information about what the Service does.
	Name       string   `json:"name"`                // Name of Service
	Repository *string  `json:"repository,omitempty"`// Open Source Repository for Service
}

type Access string
const (
	Read Access = "read"
)
