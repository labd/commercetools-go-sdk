package connect

// Generated file, please do not change!!!

import (
	"fmt"
)

type ConnectorsDraftsKeyByKeyRequestBuilder struct {
	key    string
	client *Client
}

func (rb *ConnectorsDraftsKeyByKeyRequestBuilder) Get() *ConnectorsDraftsKeyByKeyRequestMethodGet {
	return &ConnectorsDraftsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/connectors/drafts/key=%s", rb.key),
		client: rb.client,
	}
}

func (rb *ConnectorsDraftsKeyByKeyRequestBuilder) Post(body ConnectorUpdate) *ConnectorsDraftsKeyByKeyRequestMethodPost {
	return &ConnectorsDraftsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/connectors/drafts/key=%s", rb.key),
		client: rb.client,
	}
}
