package integrationlayer

// Generated file, please do not change!!!

/**
*	An error, as every management route answers one.
 */
type ErrorResponse struct {
	Message          *string `json:"error,omitempty"`
	ErrorDescription *string `json:"error_description,omitempty"`
}
