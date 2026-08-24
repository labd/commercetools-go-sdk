package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyVariantsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyVariantsByIDRequestBuilder) Images() *ByProjectKeyVariantsByIDImagesRequestBuilder {
	return &ByProjectKeyVariantsByIDImagesRequestBuilder{
		projectKey: rb.projectKey,
		id:         rb.id,
		client:     rb.client,
	}
}

/**
*	Retrieves a Variant by its ID.
*
 */
func (rb *ByProjectKeyVariantsByIDRequestBuilder) Get() *ByProjectKeyVariantsByIDRequestMethodGet {
	return &ByProjectKeyVariantsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/variants/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a Variant exists by its ID.
*	Returns a `200` status if the Variant exists, or a `404` status otherwise.
*
 */
func (rb *ByProjectKeyVariantsByIDRequestBuilder) Head() *ByProjectKeyVariantsByIDRequestMethodHead {
	return &ByProjectKeyVariantsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/variants/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Updates a Variant identified by its ID.
*
 */
func (rb *ByProjectKeyVariantsByIDRequestBuilder) Post(body VariantUpdate) *ByProjectKeyVariantsByIDRequestMethodPost {
	return &ByProjectKeyVariantsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/variants/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Deletes a Variant identified by its ID.
*
 */
func (rb *ByProjectKeyVariantsByIDRequestBuilder) Delete() *ByProjectKeyVariantsByIDRequestMethodDelete {
	return &ByProjectKeyVariantsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/variants/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
