package connect

// Generated file, please do not change!!!

import (
	"fmt"
)

type ConnectorsKeyByKeyRequestBuilder struct {
	key    string
	client *Client
}

func (rb *ConnectorsKeyByKeyRequestBuilder) Get() *ConnectorsKeyByKeyRequestMethodGet {
	return &ConnectorsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/connectors/key=%s", rb.key),
		client: rb.client,
	}
}
