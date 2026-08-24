package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyVariantsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Retrieves a Variant by its key.
*
 */
func (rb *ByProjectKeyVariantsKeyByKeyRequestBuilder) Get() *ByProjectKeyVariantsKeyByKeyRequestMethodGet {
	return &ByProjectKeyVariantsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/variants/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a Variant exists by its key.
*	Returns a `200` status if the Variant exists, or a `404` status otherwise.
*
 */
func (rb *ByProjectKeyVariantsKeyByKeyRequestBuilder) Head() *ByProjectKeyVariantsKeyByKeyRequestMethodHead {
	return &ByProjectKeyVariantsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/variants/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Updates a Variant identified by its key.
*
 */
func (rb *ByProjectKeyVariantsKeyByKeyRequestBuilder) Post(body VariantUpdate) *ByProjectKeyVariantsKeyByKeyRequestMethodPost {
	return &ByProjectKeyVariantsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/variants/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Deletes a Variant identified by its key.
*
 */
func (rb *ByProjectKeyVariantsKeyByKeyRequestBuilder) Delete() *ByProjectKeyVariantsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyVariantsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/variants/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
