package platform

// Generated file, please do not change!!!

import (
	"time"
)

type ApiClient struct {
	// The OAuth2 `client_id` that can be used to [obtain an access token](/../api/authorization#request-an-access-token-using-the-composable-commerce-oauth-20-service).
	ID string `json:"id"`
	// Name of the APIClient.
	Name string `json:"name"`
	// Whitespace-separated list of [OAuth scopes](/../api/scopes) that can be used when [obtaining an access token](/../api/authorization#request-an-access-token-using-the-composable-commerce-oauth-20-service).
	Scope string `json:"scope"`
	// Only shown once in the response of creating the APIClient.
	// This is the OAuth2 `client_secret` that can be used to [obtain an access token](/../api/authorization#request-an-access-token-using-the-composable-commerce-oauth-20-service).
	Secret *string `json:"secret,omitempty"`
	// Date of the last day this APIClient was used to [obtain an access token](/../api/authorization#request-an-access-token-using-the-composable-commerce-oauth-20-service).
	LastUsedAt *Date `json:"lastUsedAt,omitempty"`
	// If set, the Client will be deleted on (or shortly after) this point in time.
	DeleteAt *time.Time `json:"deleteAt,omitempty"`
	// Date and time (UTC) the APIClient was initially created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Expiration time in seconds for each access token obtained by the APIClient. Only present when set with the [APIClientDraft](ctp:api:type:ApiClientDraft). If not present the default value applies.
	AccessTokenValiditySeconds *int `json:"accessTokenValiditySeconds,omitempty"`
	// Inactivity expiration time in seconds for each refresh token obtained by the APIClient. Only present when set with the [APIClientDraft](ctp:api:type:ApiClientDraft). If not present the default value applies.
	RefreshTokenValiditySeconds *int `json:"refreshTokenValiditySeconds,omitempty"`
}

type ApiClientDraft struct {
	// Name of the APIClient.
	Name string `json:"name"`
	// Whitespace-separated list of [OAuth scopes](/../api/scopes) that can be used when [obtaining an access token](/../api/authorization#request-an-access-token-using-the-composable-commerce-oauth-20-service).
	Scope string `json:"scope"`
	// If set, the Client will be deleted after the specified amount of days.
	DeleteDaysAfterCreation *int `json:"deleteDaysAfterCreation,omitempty"`
	// Expiration time in seconds for each access token obtained by the APIClient. If not set the default value applies.
	AccessTokenValiditySeconds *int `json:"accessTokenValiditySeconds,omitempty"`
	// Inactivity expiration time in seconds for each refresh token obtained by the APIClient. The expiration time for refresh tokens is restarted each time the token is used. If not set the default value applies.
	RefreshTokenValiditySeconds *int `json:"refreshTokenValiditySeconds,omitempty"`
}

/**
*	[PagedQueryResult](/general-concepts#pagedqueryresult) with `results` containing an array of [APIClient](ctp:api:type:ApiClient).
*
 */
type ApiClientPagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/../api/general-concepts#strong-consistency).
	// This field is returned by default.
	// For improved performance, calculating this field can be deactivated by using the query parameter `withTotal=false`.
	// When the results are filtered with a [Query Predicate](/../api/predicates/query), `total` is subject to a [limit](/../api/limits#queries).
	Total *int `json:"total,omitempty"`
	// APIClients matching the query.
	Results []ApiClient `json:"results"`
}
