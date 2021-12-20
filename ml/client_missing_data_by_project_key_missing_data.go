// Generated file, please do not change!!!
package ml

import ()

type ByProjectKeyMissingDataRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMissingDataRequestBuilder) Attributes() *ByProjectKeyMissingDataAttributesRequestBuilder {
	return &ByProjectKeyMissingDataAttributesRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMissingDataRequestBuilder) Images() *ByProjectKeyMissingDataImagesRequestBuilder {
	return &ByProjectKeyMissingDataImagesRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMissingDataRequestBuilder) Prices() *ByProjectKeyMissingDataPricesRequestBuilder {
	return &ByProjectKeyMissingDataPricesRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
