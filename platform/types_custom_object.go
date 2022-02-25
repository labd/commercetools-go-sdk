package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"time"
)

type CustomObject struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 2019-02-01 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 2019-02-01 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// A namespace to group custom objects.
	Container string      `json:"container"`
	Key       string      `json:"key"`
	Value     interface{} `json:"value"`
}

type CustomObjectDraft struct {
	// A namespace to group custom objects.
	Container string `json:"container"`
	// A user-defined key that is unique within the given container.
	Key     string      `json:"key"`
	Value   interface{} `json:"value"`
	Version *int        `json:"version,omitempty"`
}

type CustomObjectPagedQueryResponse struct {
	Limit   int            `json:"limit"`
	Count   int            `json:"count"`
	Total   *int           `json:"total,omitempty"`
	Offset  int            `json:"offset"`
	Results []CustomObject `json:"results"`
}

type CustomObjectReference struct {
	// Unique ID of the referenced resource.
	ID  string        `json:"id"`
	Obj *CustomObject `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomObjectReference) MarshalJSON() ([]byte, error) {
	type Alias CustomObjectReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "key-value-document", Alias: (*Alias)(&obj)})
}
