package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
)

type GraphQLError struct {
	Message   string                 `json:"message"`
	Locations []GraphQLErrorLocation `json:"locations"`
	Path      []interface{}          `json:"path"`
	// Represents a single error.
	Extensions GraphQLErrorObject `json:"extensions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *GraphQLError) UnmarshalJSON(data []byte) error {
	type Alias GraphQLError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Extensions != nil {
		var err error
		obj.Extensions, err = mapDiscriminatorGraphQLErrorObject(obj.Extensions)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj GraphQLError) MarshalJSON() ([]byte, error) {
	type Alias GraphQLError
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["path"] == nil {
		delete(raw, "path")
	}

	return json.Marshal(raw)

}

func (obj GraphQLError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown GraphQLError: failed to parse error response"
}

type GraphQLErrorLocation struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

type GraphQLRequest struct {
	Query         string               `json:"query"`
	OperationName *string              `json:"operationName,omitempty"`
	Variables     *GraphQLVariablesMap `json:"variables,omitempty"`
}

type GraphQLResponse struct {
	Data   interface{}    `json:"data,omitempty"`
	Errors []GraphQLError `json:"errors"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj GraphQLResponse) MarshalJSON() ([]byte, error) {
	type Alias GraphQLResponse
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["errors"] == nil {
		delete(raw, "errors")
	}

	return json.Marshal(raw)

}

type GraphQLVariablesMap map[string]interface{}
