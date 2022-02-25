package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyShippingMethodsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyShippingMethodsKeyByKeyRequestBuilder) Get() *ByProjectKeyShippingMethodsKeyByKeyRequestMethodGet {
	return &ByProjectKeyShippingMethodsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/shipping-methods/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyShippingMethodsKeyByKeyRequestBuilder) Post(body ShippingMethodUpdate) *ByProjectKeyShippingMethodsKeyByKeyRequestMethodPost {
	return &ByProjectKeyShippingMethodsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/shipping-methods/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyShippingMethodsKeyByKeyRequestBuilder) Delete() *ByProjectKeyShippingMethodsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyShippingMethodsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/shipping-methods/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
