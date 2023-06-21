package frontend

// Generated file, please do not change!!!

import (
	"fmt"
)

type FrontasticDataSourceByIdentifierRequestBuilder struct {
	identifier string
	client     *Client
}

/**
*	Returns data and preview from the data source.
 */
func (rb *FrontasticDataSourceByIdentifierRequestBuilder) Get(body interface{}) *FrontasticDataSourceByIdentifierRequestMethodGet {
	return &FrontasticDataSourceByIdentifierRequestMethodGet{
		body:   body,
		url:    fmt.Sprintf("/frontastic/data-source/%s", rb.identifier),
		client: rb.client,
	}
}
