package connect

// Generated file, please do not change!!!

import (
	"fmt"
)

type ConnectorsDraftsRequestBuilder struct {
	client *Client
}

func (rb *ConnectorsDraftsRequestBuilder) WithId(id string) *ConnectorsDraftsByIDRequestBuilder {
	return &ConnectorsDraftsByIDRequestBuilder{
		id:     id,
		client: rb.client,
	}
}
func (rb *ConnectorsDraftsRequestBuilder) WithKey(key string) *ConnectorsDraftsKeyByKeyRequestBuilder {
	return &ConnectorsDraftsKeyByKeyRequestBuilder{
		key:    key,
		client: rb.client,
	}
}
func (rb *ConnectorsDraftsRequestBuilder) Post(body ConnectorDraft) *ConnectorsDraftsRequestMethodPost {
	return &ConnectorsDraftsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/connectors/drafts"),
		client: rb.client,
	}
}

func (rb *ConnectorsDraftsRequestBuilder) Get() *ConnectorsDraftsRequestMethodGet {
	return &ConnectorsDraftsRequestMethodGet{
		url:    fmt.Sprintf("/connectors/drafts"),
		client: rb.client,
	}
}
