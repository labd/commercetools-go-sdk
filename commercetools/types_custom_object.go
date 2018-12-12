// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"
)

// CustomObject is of type Resource
type CustomObject struct {
	Version        int         `json:"version"`
	LastModifiedAt time.Time   `json:"lastModifiedAt"`
	ID             string      `json:"id"`
	CreatedAt      time.Time   `json:"createdAt"`
	Value          interface{} `json:"value"`
	Key            string      `json:"key"`
	Container      string      `json:"container"`
}

// CustomObjectDraft is a standalone struct
type CustomObjectDraft struct {
	Version   int         `json:"version,omitempty"`
	Value     interface{} `json:"value"`
	Key       string      `json:"key"`
	Container string      `json:"container"`
}

// CustomObjectPagedQueryResponse is of type PagedQueryResponse
type CustomObjectPagedQueryResponse struct {
	Total   int            `json:"total,omitempty"`
	Offset  int            `json:"offset"`
	Count   int            `json:"count"`
	Results []CustomObject `json:"results"`
}

// CustomObjectReference implements the interface Reference
type CustomObjectReference struct {
	Key string        `json:"key,omitempty"`
	ID  string        `json:"id,omitempty"`
	Obj *CustomObject `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomObjectReference) MarshalJSON() ([]byte, error) {
	type Alias CustomObjectReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "key-value-document", Alias: (*Alias)(&obj)})
}
