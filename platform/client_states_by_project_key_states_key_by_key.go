package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyStatesKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyStatesKeyByKeyRequestBuilder) Get() *ByProjectKeyStatesKeyByKeyRequestMethodGet {
	return &ByProjectKeyStatesKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/states/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a State exists with the provided `key`. Returns a `200 OK` status if the State exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyStatesKeyByKeyRequestBuilder) Head() *ByProjectKeyStatesKeyByKeyRequestMethodHead {
	return &ByProjectKeyStatesKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/states/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyStatesKeyByKeyRequestBuilder) Post(body StateUpdate) *ByProjectKeyStatesKeyByKeyRequestMethodPost {
	return &ByProjectKeyStatesKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/states/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyStatesKeyByKeyRequestBuilder) Delete() *ByProjectKeyStatesKeyByKeyRequestMethodDelete {
	return &ByProjectKeyStatesKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/states/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
