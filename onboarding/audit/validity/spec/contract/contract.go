// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    contract, err := UnmarshalContract(bytes)
//    bytes, err = contract.Marshal()

package contract

import "encoding/json"

func UnmarshalContract(data []byte) (Contract, error) {
	var r Contract
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Contract) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Service Onboarding Contract Schema
type Contract struct {
	Data      []Datum   `json:"data"`      
	Developer Developer `json:"developer"` 
	Kind      Kind      `json:"kind"`      // Kind of Resource
	RequestID string    `json:"request_id"`// Service Onboarding
	Service   Service   `json:"service"`   
}

// Information about Data Required and Usage
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

// Kind of Resource
type Kind string
const (
	ContractIntakeServiceOnboarding Kind = "contract.intake.service-onboarding"
)
