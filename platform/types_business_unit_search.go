package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"time"
)

type BusinessUnitIndexingProgress struct {
	// The number of Business Units successfully indexed.
	Indexed int `json:"indexed"`
	// The number of Business Units that failed to be indexed.
	Failed int `json:"failed"`
	// The estimated total number of Business Units to be indexed.
	EstimatedTotal int `json:"estimatedTotal"`
}

type BusinessUnitPagedSearchResponse struct {
	// Total number of results matching the query.
	Total int `json:"total"`
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// Search result containing the Business Units matching the search query.
	Results []BusinessUnitSearchResult `json:"results"`
}

type BusinessUnitSearchIndexingStatusResponse struct {
	// Current status of indexing the Business Unit Search.
	Status BusinessUnitIndexingStatus `json:"status"`
	// Progress of indexing. Only available when indexing is in progress.
	States *BusinessUnitIndexingProgress `json:"states,omitempty"`
	// Date and time (UTC) when the last indexing started.
	StartedAt *time.Time `json:"startedAt,omitempty"`
	// Time when the status was last modified.
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`
	// Indicates how many times the system tried to start indexing after failed attempts. The counter is set to null after an indexing finished successfully.
	RetryCount *int `json:"retryCount,omitempty"`
}

type BusinessUnitSearchRequest struct {
	// The Business Unit Search query.
	Query *SearchQuery `json:"query,omitempty"`
	// Controls how results to your query are sorted. If not provided, the results are sorted by relevance in descending order.
	Sort []SearchSorting `json:"sort"`
	// The maximum number of search results to be returned.
	Limit *int `json:"limit,omitempty"`
	// The number of search results to be skipped in the response for pagination.
	Offset *int `json:"offset,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitSearchRequest) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitSearchRequest
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["sort"] == nil {
		delete(raw, "sort")
	}

	return json.Marshal(raw)

}

type BusinessUnitSearchResult struct {
	// `id` of the [BusinessUnit](ctp:api:type:BusinessUnit) matching the search query.
	ID string `json:"id"`
	// How closely this customer matches the search query.
	Relevance float64 `json:"relevance"`
}
