package connect

// Generated file, please do not change!!!

import (
	"fmt"
)

type ConnectorsSearchRequestBuilder struct {
	client *Client
}

/**
*	Retrieves all available Connectors.
 */
func (rb *ConnectorsSearchRequestBuilder) Get() *ConnectorsSearchRequestMethodGet {
	return &ConnectorsSearchRequestMethodGet{
		url:    fmt.Sprintf("/connectors/search"),
		client: rb.client,
	}
}
