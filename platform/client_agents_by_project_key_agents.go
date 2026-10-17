package platform

// Generated file, please do not change!!!

type ByProjectKeyAgentsRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Converts unstructured customer input into a [Cart](ctp:api:type:Cart), [QuoteRequest](ctp:api:type:QuoteRequest), or [ShoppingList](ctp:api:type:ShoppingList).
*
 */
func (rb *ByProjectKeyAgentsRequestBuilder) Intake() *ByProjectKeyAgentsIntakeRequestBuilder {
	return &ByProjectKeyAgentsIntakeRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
