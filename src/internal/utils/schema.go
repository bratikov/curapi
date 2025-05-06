package utils

import (
	"encoding/json"
	"fmt"

	"github.com/invopop/jsonschema"
	"github.com/xeipuuv/gojsonschema"
)

func ValidateRequest(request string, anyStruct interface{}) error {
	schema := jsonschema.Reflect(anyStruct)

	sch, err := schema.MarshalJSON()
	if err != nil {
		return err
	}

	schemaLoader := gojsonschema.NewStringLoader(string(sch))
	documentLoader := gojsonschema.NewStringLoader(request)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return err
	}

	if !result.Valid() {
		errors := ""
		for _, desc := range result.Errors() {
			errors += fmt.Sprintf("%s, ", desc)
		}
		return fmt.Errorf("%s", errors)
	}

	return nil
}

func SelfValidate(anyStruct interface{}) error {
	schema := jsonschema.Reflect(anyStruct)

	sch, err := schema.MarshalJSON()
	if err != nil {
		return err
	}

	req, err := json.Marshal(anyStruct)
	if err != nil {
		return err
	}

	schemaLoader := gojsonschema.NewStringLoader(string(sch))
	documentLoader := gojsonschema.NewStringLoader(string(req))

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return err
	}

	if !result.Valid() {
		errors := ""
		for _, desc := range result.Errors() {
			errors += fmt.Sprintf("%s, ", desc)
		}
		return fmt.Errorf("%s", errors)
	}

	return nil
}
