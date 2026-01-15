package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyTypesKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyTypesKeyByKeyRequestBuilder) Get() *ByProjectKeyTypesKeyByKeyRequestMethodGet {
	return &ByProjectKeyTypesKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/types/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a Type exists with the provided `key`. Returns a `200` status if the Type exists, or a `404` status otherwise.
 */
func (rb *ByProjectKeyTypesKeyByKeyRequestBuilder) Head() *ByProjectKeyTypesKeyByKeyRequestMethodHead {
	return &ByProjectKeyTypesKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/types/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyTypesKeyByKeyRequestBuilder) Post(body TypeUpdate) *ByProjectKeyTypesKeyByKeyRequestMethodPost {
	return &ByProjectKeyTypesKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/types/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyTypesKeyByKeyRequestBuilder) Delete() *ByProjectKeyTypesKeyByKeyRequestMethodDelete {
	return &ByProjectKeyTypesKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/types/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
