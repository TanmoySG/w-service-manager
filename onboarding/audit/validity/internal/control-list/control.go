package controlList

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	READ   = "read"
	WRITE  = "write"
	UPDATE = "update"
	DELETE = "delete"
)

type ControList map[string][]string

func isSubset(first, second []string) bool {
	set := make(map[string]int)
	for _, value := range second {
		set[value] += 1
	}

	for _, value := range first {
		if count, found := set[value]; !found {
			return false
		} else if count < 1 {
			return false
		} else {
			set[value] = count - 1
		}
	}

	return true
}

func LoadControlList(filepath string) (ControList, error) {
	controlListBytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("error parsing file : %s", err)
	}

	var controList ControList
	errr := json.Unmarshal(controlListBytes, &controList)
	if errr != nil {
		return nil, fmt.Errorf("error parsing file : %s", errr)
	}

	return controList, nil
}

func (cl ControList) GetAllowedAccessForField(fieldName string) []string {
	return cl[fieldName]
}

func (cl ControList) ValidateAccessForField(fieldName string, assignedAccessToField []string) bool {
	allowedAccessForField := cl.GetAllowedAccessForField(fieldName)
	return isSubset(assignedAccessToField, allowedAccessForField)
}

func (cl ControList) ValidateAccessForAll(assignedAccessToField ControList) bool {
	for field, access := range assignedAccessToField {
		if !cl.ValidateAccessForField(field, access) {
			return false
		}
	}
	return true
}
