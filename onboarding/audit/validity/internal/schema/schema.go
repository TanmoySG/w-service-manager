package schema

import (
	"fmt"
	"io/ioutil"

	sc "github.com/xeipuuv/gojsonschema"
)

type SchemaValidator struct {
	SchemaPath   string
	DocumentPath string
}

func loadFromFile(filepath string) string {
	filebytes, _ := ioutil.ReadFile(filepath)
	return string(filebytes)
}

func (sv SchemaValidator) Validate() (bool, []error) {
	schema := sc.NewStringLoader(loadFromFile(sv.SchemaPath))
	document := sc.NewStringLoader(loadFromFile(sv.DocumentPath))

	result, err := sc.Validate(schema, document)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		return true, nil
	} else {
		var errors []error
		for _, desc := range result.Errors() {
			errors = append(errors, fmt.Errorf(desc.Description()))
		}
		return false, errors
	}
}
