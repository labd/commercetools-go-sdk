package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"time"
)

type CustomerIndexingProgress struct {
	// The number of Customers successfully indexed.
	Indexed int `json:"indexed"`
	// The number of Customers that failed to be indexed.
	Failed int `json:"failed"`
	// The estimated total number of Customers to be indexed.
	EstimatedTotal int `json:"estimatedTotal"`
}

type CustomerPagedSearchResponse struct {
	// Total number of results matching the query.
	Total int `json:"total"`
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// Search result containing the Customers matching the search query.
	Results []CustomerSearchResult `json:"results"`
}

type CustomerSearchIndexingStatusResponse struct {
	// Current status of indexing the Customer Search.
	Status CustomerIndexingStatus `json:"status"`
	// Progress of indexing. Only available when indexing is in progress.
	States *CustomerIndexingProgress `json:"states,omitempty"`
	// Date and time (UTC) when the last indexing started.
	StartedAt *time.Time `json:"startedAt,omitempty"`
	// Time when the status was last modified.
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`
	// Indicates how many times the system tried to start indexing after failed attempts. The counter is set to null after an indexing finished successfully.
	RetryCount *int `json:"retryCount,omitempty"`
}

type CustomerSearchRequest struct {
	// The Customer search query.
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
func (obj CustomerSearchRequest) MarshalJSON() ([]byte, error) {
	type Alias CustomerSearchRequest
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

type CustomerSearchResult struct {
	// `id` of the [Customer](ctp:api:type:Customer) matching the search query.
	ID string `json:"id"`
	// How closely this customer matches the search query.
	Relevance float64 `json:"relevance"`
}
