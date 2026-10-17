package integrationlayer

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMainIdpRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMainIdpRequestBuilder) Get() *ByProjectKeyMainIdpRequestMethodGet {
	return &ByProjectKeyMainIdpRequestMethodGet{
		url:    fmt.Sprintf("/%s/main/idp", rb.projectKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyMainIdpRequestBuilder) Put(body IdpConfigInput) *ByProjectKeyMainIdpRequestMethodPut {
	return &ByProjectKeyMainIdpRequestMethodPut{
		body:   body,
		url:    fmt.Sprintf("/%s/main/idp", rb.projectKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyMainIdpRequestBuilder) Delete() *ByProjectKeyMainIdpRequestMethodDelete {
	return &ByProjectKeyMainIdpRequestMethodDelete{
		url:    fmt.Sprintf("/%s/main/idp", rb.projectKey),
		client: rb.client,
	}
}
