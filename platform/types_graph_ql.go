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
	OperationName string               `json:"operationName,omitEmpty"`
	Variables     *GraphQLVariablesMap `json:"variables,omitEmpty"`
}

type GraphQLResponse struct {
	Data   interface{}    `json:"data,omitEmpty"`
	Errors []GraphQLError `json:"errors,omitEmpty"`
}

// GraphQLVariablesMap is something
type GraphQLVariablesMap map[string]interface{}
