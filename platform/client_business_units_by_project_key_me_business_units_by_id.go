package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeBusinessUnitsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyMeBusinessUnitsByIDRequestBuilder) Get() *ByProjectKeyMeBusinessUnitsByIDRequestMethodGet {
	return &ByProjectKeyMeBusinessUnitsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/business-units/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyMeBusinessUnitsByIDRequestBuilder) Post(body MyBusinessUnitUpdate) *ByProjectKeyMeBusinessUnitsByIDRequestMethodPost {
	return &ByProjectKeyMeBusinessUnitsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/business-units/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyMeBusinessUnitsByIDRequestBuilder) Delete() *ByProjectKeyMeBusinessUnitsByIDRequestMethodDelete {
	return &ByProjectKeyMeBusinessUnitsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/me/business-units/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
