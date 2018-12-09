// Automatically generated, do not edit

package commercetools

type PagedQueryResponse struct {
	Total   int        `json:"total,omitempty"`
	Results []Resource `json:"results"`
	Offset  int        `json:"offset"`
	Count   int        `json:"count"`
}

type Update struct {
	Version int                    `json:"version"`
	Actions []AbstractUpdateAction `json:"actions"`
}

type UpdateAction interface{}
type AbstractUpdateAction struct {
	Action string `json:"action"`
}
