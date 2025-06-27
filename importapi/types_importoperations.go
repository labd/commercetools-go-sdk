package importapi

// Generated file, please do not change!!!

import (
	"encoding/json"
	"time"
)

/**
*	Represents the import status of a resource.
*
 */
type ImportOperation struct {
	// Current version of the ImportOperation.
	Version int `json:"version"`
	// `key` of the [ImportContainer](ctp:import:type:ImportContainer).
	ImportContainerKey string `json:"importContainerKey"`
	// `key` of the resource being imported.
	ResourceKey string `json:"resourceKey"`
	// Unique identifier of the ImportOperation.
	ID string `json:"id"`
	// The import status of the resource. If `rejected` or `validationFailed`, the import was unsuccessful.
	State ProcessingState `json:"state"`
	// The `version` of the imported resource when the import was successful.
	ResourceVersion *int `json:"resourceVersion,omitempty"`
	// Contains errors if the import was unsuccessful. See [Errors](/import-export/error).
	Errors []ErrorObject `json:"errors"`
	// If the resource being imported contains references to resources which do not exist, these references are contained within this array.
	UnresolvedReferences []UnresolvedReferences `json:"unresolvedReferences"`
	// Date and time (UTC) the ImportOperation was created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the ImportOperation was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Date and time (UTC) the ImportOperation will be deleted.
	ExpiresAt time.Time `json:"expiresAt"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ImportOperation) UnmarshalJSON(data []byte) error {
	type Alias ImportOperation
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Errors {
		var err error
		obj.Errors[i], err = mapDiscriminatorErrorObject(obj.Errors[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ImportOperation) MarshalJSON() ([]byte, error) {
	type Alias ImportOperation
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

	if raw["errors"] == nil {
		delete(raw, "errors")
	}

	if raw["unresolvedReferences"] == nil {
		delete(raw, "unresolvedReferences")
	}

	return json.Marshal(raw)

}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) for Import Operations.
*
 */
type ImportOperationPagedResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// The actual number of results returned.
	Count int `json:"count"`
	// The total number of import operations matching the query.
	Total int `json:"total"`
	// The array of Import Operations matching the query.
	Results []ImportOperation `json:"results"`
}

/**
*	Describes the validation state of a newly created [ImportOperation](#importoperation).
*
 */
type ImportOperationState string

const (
	ImportOperationStateProcessing       ImportOperationState = "processing"
	ImportOperationStateValidationFailed ImportOperationState = "validationFailed"
)

/**
*	The status of a new [ImportOperation](#importoperation).
 */
type ImportOperationStatus struct {
	// `id` of the [ImportOperation](#importoperation).
	OperationId *string `json:"operationId,omitempty"`
	// Validation state of the [ImportOperation](#importoperation).
	State ImportOperationState `json:"state"`
	// [Errors](/import-export/error) for the [ImportOperation](#importoperation).
	Errors []ErrorObject `json:"errors"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ImportOperationStatus) UnmarshalJSON(data []byte) error {
	type Alias ImportOperationStatus
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Errors {
		var err error
		obj.Errors[i], err = mapDiscriminatorErrorObject(obj.Errors[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ImportOperationStatus) MarshalJSON() ([]byte, error) {
	type Alias ImportOperationStatus
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

	if raw["errors"] == nil {
		delete(raw, "errors")
	}

	return json.Marshal(raw)

}
