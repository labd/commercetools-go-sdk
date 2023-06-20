package connect

// Generated file, please do not change!!!

import (
	"fmt"
)

type ConnectorsByIDRequestBuilder struct {
	id     string
	client *Client
}

func (rb *ConnectorsByIDRequestBuilder) Get() *ConnectorsByIDRequestMethodGet {
	return &ConnectorsByIDRequestMethodGet{
		url:    fmt.Sprintf("/connectors/%s", rb.id),
		client: rb.client,
	}
}
