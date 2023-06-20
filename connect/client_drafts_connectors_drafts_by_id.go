package connect

// Generated file, please do not change!!!

import (
	"fmt"
)

type ConnectorsDraftsByIDRequestBuilder struct {
	id     string
	client *Client
}

func (rb *ConnectorsDraftsByIDRequestBuilder) Get() *ConnectorsDraftsByIDRequestMethodGet {
	return &ConnectorsDraftsByIDRequestMethodGet{
		url:    fmt.Sprintf("/connectors/drafts/%s", rb.id),
		client: rb.client,
	}
}

func (rb *ConnectorsDraftsByIDRequestBuilder) Post(body ConnectorUpdate) *ConnectorsDraftsByIDRequestMethodPost {
	return &ConnectorsDraftsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/connectors/drafts/%s", rb.id),
		client: rb.client,
	}
}
