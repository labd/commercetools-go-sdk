// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
)

type GraphQLError struct {
	Message   string                 `json:"message"`
	Locations []GraphQLErrorLocation `json:"locations"`
	Path      []interface{}          `json:"path"`
}

func (obj GraphQLError) Error() string {
	return obj.Message
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

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["errors"] == nil {
		delete(target, "errors")
	}

	return json.Marshal(target)
}

// GraphQLVariablesMap is something
type GraphQLVariablesMap map[string]interface{}
