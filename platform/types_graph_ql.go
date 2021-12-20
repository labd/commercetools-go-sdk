// Generated file, please do not change!!!
package platform

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
	Errors []GraphQLError `json:"errors,omitempty"`
}

// GraphQLVariablesMap is something
type GraphQLVariablesMap map[string]interface{}
