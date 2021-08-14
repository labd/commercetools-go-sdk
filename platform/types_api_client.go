// Generated file, please do not change!!!
package platform

import (
	"time"
)

type ApiClient struct {
	// The unique ID of the API client.
	// This is the OAuth2 `client_id` and can be used to obtain a token.
	Id   string `json:"id"`
	Name string `json:"name"`
	// A whitespace separated list of the OAuth scopes.
	// This is the OAuth2 `scope` and can be used to obtain a token.
	Scope     string    `json:"scope"`
	CreatedAt time.Time `json:"createdAt,omitEmpty"`
	// The last day this API Client was used to obtain a token.
	LastUsedAt Date `json:"lastUsedAt,omitEmpty"`
	// If set, the client will be deleted on (or shortly after) this point in time.
	DeleteAt time.Time `json:"deleteAt,omitEmpty"`
	// The secret is only shown once in the response of creating the API Client.
	// This is the OAuth2 `client_secret` and can be used to obtain a token.
	Secret string `json:"secret,omitEmpty"`
}

type ApiClientDraft struct {
	Name  string `json:"name"`
	Scope string `json:"scope"`
	// If set, the client will be deleted after the specified amount of days.
	DeleteDaysAfterCreation int `json:"deleteDaysAfterCreation,omitEmpty"`
}

type ApiClientPagedQueryResponse struct {
	Limit   int         `json:"limit"`
	Count   int         `json:"count"`
	Total   int         `json:"total,omitEmpty"`
	Offset  int         `json:"offset"`
	Results []ApiClient `json:"results"`
}
