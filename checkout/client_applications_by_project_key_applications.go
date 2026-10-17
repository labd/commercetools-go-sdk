package checkout

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyApplicationsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyApplicationsRequestBuilder) WithId(id string) *ByProjectKeyApplicationsByIdRequestBuilder {
	return &ByProjectKeyApplicationsByIdRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyApplicationsRequestBuilder) WithKey(key string) *ByProjectKeyApplicationsKeyByKeyRequestBuilder {
	return &ByProjectKeyApplicationsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Retrieves Applications
*
 */
func (rb *ByProjectKeyApplicationsRequestBuilder) Get() *ByProjectKeyApplicationsRequestMethodGet {
	return &ByProjectKeyApplicationsRequestMethodGet{
		url:    fmt.Sprintf("/%s/applications", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Creates a new [Application](ctp:checkout:type:Application).
*
 */
func (rb *ByProjectKeyApplicationsRequestBuilder) Post(body ApplicationDraft) *ByProjectKeyApplicationsRequestMethodPost {
	return &ByProjectKeyApplicationsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/applications", rb.projectKey),
		client: rb.client,
	}
}
