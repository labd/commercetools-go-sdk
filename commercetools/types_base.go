// Automatically generated, do not edit

package commercetools

type UpdateAction interface{}

// PagedQueryResponse is a standalone struct
type PagedQueryResponse struct {
	Total   int        `json:"total,omitempty"`
	Results []Resource `json:"results"`
	Offset  int        `json:"offset"`
	Count   int        `json:"count"`
}

// Update is a standalone struct
type Update struct {
	Version int            `json:"version"`
	Actions []UpdateAction `json:"actions"`
}
