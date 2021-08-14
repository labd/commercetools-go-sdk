// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"time"
)

type CustomObject struct {
	Id             string          `json:"id"`
	Version        int             `json:"version"`
	CreatedAt      time.Time       `json:"createdAt"`
	LastModifiedAt time.Time       `json:"lastModifiedAt"`
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitEmpty"`
	CreatedBy      *CreatedBy      `json:"createdBy,omitEmpty"`
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
	Version int         `json:"version,omitEmpty"`
}

type CustomObjectPagedQueryResponse struct {
	Limit   int            `json:"limit"`
	Count   int            `json:"count"`
	Total   int            `json:"total,omitEmpty"`
	Offset  int            `json:"offset"`
	Results []CustomObject `json:"results"`
}

type CustomObjectReference struct {
	Id  string        `json:"id"`
	Obj *CustomObject `json:"obj,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomObjectReference) MarshalJSON() ([]byte, error) {
	type Alias CustomObjectReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "key-value-document", Alias: (*Alias)(&obj)})
}
