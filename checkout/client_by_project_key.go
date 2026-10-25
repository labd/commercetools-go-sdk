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
func (rb *ByProjectKeyRequestBuilder) Transactions() *ByProjectKeyTransactionsRequestBuilder {
	return &ByProjectKeyTransactionsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyRequestBuilder) RecurringPaymentJobs() *ByProjectKeyRecurringPaymentJobsRequestBuilder {
	return &ByProjectKeyRecurringPaymentJobsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyRequestBuilder) RecurringPayments() *ByProjectKeyRecurringPaymentsRequestBuilder {
	return &ByProjectKeyRecurringPaymentsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyRequestBuilder) PaymentIntegrations() *ByProjectKeyPaymentIntegrationsRequestBuilder {
	return &ByProjectKeyPaymentIntegrationsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyRequestBuilder) Applications() *ByProjectKeyApplicationsRequestBuilder {
	return &ByProjectKeyApplicationsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
