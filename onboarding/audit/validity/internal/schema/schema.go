package schema

import (
	"fmt"
	"io/ioutil"

	sc "github.com/xeipuuv/gojsonschema"
)

type SchemaValidator struct {
	Schema   string
	Document string
}

func LoadFromFile(filepath string) string {
	filebytes, _ := ioutil.ReadFile(filepath)
	return string(filebytes)
}

func (sv SchemaValidator) Validate() (bool, []error) {
	schema := sc.NewStringLoader(sv.Schema)
	document := sc.NewStringLoader(sv.Document)

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
