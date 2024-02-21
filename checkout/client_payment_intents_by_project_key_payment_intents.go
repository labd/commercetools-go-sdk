package checkout

// Generated file, please do not change!!!

type ByProjectKeyPaymentIntentsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyPaymentIntentsRequestBuilder) WithPaymentId(paymentId string) *ByProjectKeyPaymentIntentsByPaymentIdRequestBuilder {
	return &ByProjectKeyPaymentIntentsByPaymentIdRequestBuilder{
		paymentId:  paymentId,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
