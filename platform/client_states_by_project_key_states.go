package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyStatesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyStatesRequestBuilder) WithKey(key string) *ByProjectKeyStatesKeyByKeyRequestBuilder {
	return &ByProjectKeyStatesKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyStatesRequestBuilder) WithId(id string) *ByProjectKeyStatesByIDRequestBuilder {
	return &ByProjectKeyStatesByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyStatesRequestBuilder) Get() *ByProjectKeyStatesRequestMethodGet {
	return &ByProjectKeyStatesRequestMethodGet{
		url:    fmt.Sprintf("/%s/states", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if a State exists for a given Query Predicate. Returns a `200 OK` status if any States match the Query Predicate or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyStatesRequestBuilder) Head() *ByProjectKeyStatesRequestMethodHead {
	return &ByProjectKeyStatesRequestMethodHead{
		url:    fmt.Sprintf("/%s/states", rb.projectKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyStatesRequestBuilder) Post(body StateDraft) *ByProjectKeyStatesRequestMethodPost {
	return &ByProjectKeyStatesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/states", rb.projectKey),
		client: rb.client,
	}
}
