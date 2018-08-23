package customfields

import (
	"github.com/labd/commercetools-go-sdk/commercetools"
)

// Custom fields give you the possibility to add fields to predefined resources,
// similar to AttributeDefinitions on ProductTypes.
type Custom struct {
	Custom CustomFields `json:"custom"`
}

// CustomFields allow you to enhance the resources as you need.
type CustomFields struct {
	Type commercetools.Reference `json:"type"`
	// A valid JSON object, based on FieldDefinition.
	Fields map[string]string `json:"fields"`
}

// CustomFieldsDraft is the representation to be sent to the server
// when creating a resource with custom fields.
type CustomFieldsDraft struct {
	// The id or the key of the type to use.
	Type commercetools.ResourceIdentifier `json:"type"`
	// A valid JSON object, based on the FieldDefinitions of the Type.
	Fields map[string]string `json:"fields,omitempty"`
}
