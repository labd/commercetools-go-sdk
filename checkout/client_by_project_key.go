package checkout

// Generated file, please do not change!!!

type ByProjectKeyRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyRequestBuilder) PaymentIntents() *ByProjectKeyPaymentIntentsRequestBuilder {
	return &ByProjectKeyPaymentIntentsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
