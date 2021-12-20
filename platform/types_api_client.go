// Generated file, please do not change!!!
package platform

import (
	"time"
)

type ApiClient struct {
	// Unique ID of the API client.
	// This is the OAuth2 `client_id` that can be used to [obtain an access token](/../api/authorization#requesting-an-access-token-using-commercetools-oauth-20-server).
	ID string `json:"id"`
	// Name of the API Client.
	Name string `json:"name"`
	// Whitespace-separated list of [OAuth scopes](/../api/scopes) that can be used when [obtaining an access token](/../api/authorization#requesting-an-access-token-using-commercetools-oauth-20-server).
	Scope string `json:"scope"`
	// Only shown once in the response of creating the API Client.
	// This is the OAuth2 `client_secret` that can be used to [obtain an access token](/../api/authorization#requesting-an-access-token-using-commercetools-oauth-20-server).
	Secret *string `json:"secret,omitempty"`
	// Date of the last day this API Client was used to [obtain an access token](/../api/authorization#requesting-an-access-token-using-commercetools-oauth-20-server).
	LastUsedAt *Date `json:"lastUsedAt,omitempty"`
	// If set, the client will be deleted on (or shortly after) this point in time.
	DeleteAt *time.Time `json:"deleteAt,omitempty"`
	// Date and time (UTC) the API Client was initially created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
}

type ApiClientDraft struct {
	// Name of the API Client.
	Name string `json:"name"`
	// Whitespace-separated list of [OAuth scopes](/../api/scopes) that can be used when [obtaining an access token](/../api/authorization#requesting-an-access-token-using-commercetools-oauth-20-server).
	Scope string `json:"scope"`
	// If set, the client will be deleted after the specified amount of days.
	DeleteDaysAfterCreation *int `json:"deleteDaysAfterCreation,omitempty"`
}

/**
*	[PagedQueryResult](/general-concepts#pagedqueryresult) with `results` containing an array of [APIClient](ctp:api:type:ApiClient).
*
 */
type ApiClientPagedQueryResponse struct {
	// Number of results requested in the query request.
	Limit int `json:"limit"`
	// Offset supplied by the client or server default.
	// It is the number of elements skipped, not a page number.
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/../api/general-concepts#strong-consistency).
	// This field is returned by default.
	// For improved performance, calculating this field can be deactivated by using the query parameter `withTotal=false`.
	// When the results are filtered with a [Query Predicate](/../api/predicates/query), `total` is subject to a [limit](/../api/limits#queries).
	Total *int `json:"total,omitempty"`
	// API Clients matching the query.
	Results []ApiClient `json:"results"`
}
