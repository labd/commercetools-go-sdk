// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"
)

type CustomObject struct {
	Version        int         `json:"version"`
	LastModifiedAt time.Time   `json:"lastModifiedAt"`
	ID             string      `json:"id"`
	CreatedAt      time.Time   `json:"createdAt"`
	Value          interface{} `json:"value"`
	Key            string      `json:"key"`
	Container      string      `json:"container"`
}

type CustomObjectDraft struct {
	Version   int         `json:"version,omitempty"`
	Value     interface{} `json:"value"`
	Key       string      `json:"key"`
	Container string      `json:"container"`
}

type CustomObjectPagedQueryResponse struct {
	Total   int            `json:"total,omitempty"`
	Offset  int            `json:"offset"`
	Count   int            `json:"count"`
	Results []CustomObject `json:"results"`
}

type CustomObjectReference struct {
	Key string        `json:"key,omitempty"`
	ID  string        `json:"id,omitempty"`
	Obj *CustomObject `json:"obj,omitempty"`
}

func (obj CustomObjectReference) MarshalJSON() ([]byte, error) {
	type Alias CustomObjectReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "key-value-document", Alias: (*Alias)(&obj)})
}
