// Automatically generated, do not edit

package commercetools

import "time"

type ApiClient struct {
	Secret     string      `json:"secret,omitempty"`
	Scope      string      `json:"scope"`
	Name       string      `json:"name"`
	LastUsedAt interface{} `json:"lastUsedAt"`
	ID         string      `json:"id"`
	CreatedAt  time.Time   `json:"createdAt"`
}

type ApiClientDraft struct {
	Scope string `json:"scope"`
	Name  string `json:"name"`
}

type ApiClientPagedQueryResponse struct {
	Total   int         `json:"total,omitempty"`
	Results []ApiClient `json:"results"`
	Offset  int         `json:"offset"`
	Count   int         `json:"count"`
}
