package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
)

/**
*	Contains an error message, the location of the code that caused the error, and other information to help you correct the error.
 */
type GraphQLError struct {
	// Detailed description of the error explaining the root cause of the problem and suggesting how to correct the error.
	Message string `json:"message"`
	// Location within your query where the error occurred.
	Locations []GraphQLErrorLocation `json:"locations"`
	// Query fields listed in order from the root of the query response up to the field in which the error occurred. `path` is displayed in the response only if an error is associated with a particular field in the query result.
	Path []interface{} `json:"path"`
	// Dictionary with additional information where applicable.
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

/**
*	Represents the location within your query where the error occurred.
 */
type GraphQLErrorLocation struct {
	// Line number of the query where the error occurred.
	Line int `json:"line"`
	// Position in `line` where the error occurred.
	Column int `json:"column"`
}

/**
*	The query, operation name, and variables that are sent to the GraphQL API.
 */
type GraphQLRequest struct {
	// String representation of the _Source Text_ of the _Document_ that is specified in the [Language section of the GraphQL specification](https://spec.graphql.org/draft/#sec-Language).
	Query string `json:"query"`
	// Name of the operation, if you defined several operations in `query`.
	OperationName *string `json:"operationName,omitempty"`
	// JSON object that contains key-value pairs in which the keys are variable names and the values are variable values.
	Variables *GraphQLVariablesMap `json:"variables,omitempty"`
}

/**
*	`error` is present in the response only if the GraphQL query was unsuccessful.
*
 */
type GraphQLResponse struct {
	// JSON object that contains the results of a GraphQL query.
	Data interface{} `json:"data,omitempty"`
	// Errors that the GraphQL query returns.
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

/**
*	The variables that the GraphQL query uses.
 */
type GraphQLVariablesMap map[string]interface{}
