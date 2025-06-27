package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"time"
)

type CustomObject struct {
	// Unique identifier of the CustomObject.
	ID string `json:"id"`
	// Current version of the CustomObject.
	Version int `json:"version"`
	// Date and time (UTC) the CustomObject was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the CustomObject was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// IDs and references that last modified the CustomObject.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// IDs and references that created the CustomObject.
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Namespace to group CustomObjects.
	Container string `json:"container"`
	// User-defined unique identifier of the CustomObject within the defined `container`.
	Key string `json:"key"`
	// Can be any JSON standard type, such as number, string, boolean, array, object, or a [common API data type](/../api/types).
	//
	// - For values of type [Reference](ctp:api:type:Reference) the integrity of the data is not guaranteed. If the referenced object is deleted, the API does not delete the corresponding reference to it and the `value` points to a non-existing object in such case.
	Value interface{} `json:"value"`
}

type CustomObjectDraft struct {
	// Namespace to group CustomObjects.
	Container string `json:"container"`
	// User-defined unique identifier of the CustomObject within the defined `container`.
	Key string `json:"key"`
	// Can be any JSON standard type, such as number, string, boolean, array, object, or a [common API data type](/../api/types).
	//
	// - Fields within `value` that have `null` values **are not saved**.
	// - For values of type [Reference](ctp:api:type:Reference) the integrity of the data is not guaranteed. If the referenced object is deleted, the API does not delete the corresponding reference to it and the `value` points to a non-existing object in such case.
	Value interface{} `json:"value"`
	// Current version of the CustomObject.
	Version *int `json:"version,omitempty"`
}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with `results` containing an array of [CustomObject](ctp:api:type:CustomObject).
*
 */
type CustomObjectPagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// The total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/../api/general-concepts#strong-consistency).
	// This field is returned by default.
	// For improved performance, calculating this field can be deactivated by using the query parameter `withTotal=false`.
	// When the results are filtered with a [Query Predicate](/../api/predicates/query), `total` is subject to a [limit](/../api/limits#queries).
	Total *int `json:"total,omitempty"`
	// [CustomObjects](ctp:api:type:CustomObject) matching the query.
	Results []CustomObject `json:"results"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [CustomObject](ctp:api:type:CustomObject).
*
 */
type CustomObjectReference struct {
	// Unique identifier of the referenced [CustomObject](ctp:api:type:CustomObject).
	ID string `json:"id"`
	// Contains the representation of the expanded CustomObject. Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for CustomObjects.
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
