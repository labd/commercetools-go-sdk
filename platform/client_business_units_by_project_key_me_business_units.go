package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeBusinessUnitsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMeBusinessUnitsRequestBuilder) WithId(id string) *ByProjectKeyMeBusinessUnitsByIDRequestBuilder {
	return &ByProjectKeyMeBusinessUnitsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMeBusinessUnitsRequestBuilder) WithKey(key string) *ByProjectKeyMeBusinessUnitsKeyByKeyRequestBuilder {
	return &ByProjectKeyMeBusinessUnitsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMeBusinessUnitsRequestBuilder) Get() *ByProjectKeyMeBusinessUnitsRequestMethodGet {
	return &ByProjectKeyMeBusinessUnitsRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/business-units", rb.projectKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyMeBusinessUnitsRequestBuilder) Post(body MyBusinessUnitDraft) *ByProjectKeyMeBusinessUnitsRequestMethodPost {
	return &ByProjectKeyMeBusinessUnitsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/business-units", rb.projectKey),
		client: rb.client,
	}
}
