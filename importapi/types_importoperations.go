package importapi

// Generated file, please do not change!!!

import (
	"encoding/json"
	"time"
)

/**
*	Import Operation describes the import status of a specific resource.
*
 */
type ImportOperation struct {
	// The version of the ImportOperation.
	Version int `json:"version"`
	// The key of the [importContainer](/import-container#importcontainer).
	ImportContainerKey string `json:"importContainerKey"`
	// The key of the resource.
	ResourceKey string `json:"resourceKey"`
	// The ID of the ImportOperation.
	ID string `json:"id"`
	// The import status of the resource. Set to `rejected` or `validationFailed` if the import of the resource was not successful.
	State ProcessingState `json:"state"`
	// The version of the imported resource when the import was successful.
	ResourceVersion *int `json:"resourceVersion,omitempty"`
	// Contains an error if the import of the resource was not successful. See [Errors](/error).
	Errors []ErrorObject `json:"errors"`
	// In case of unresolved status this array will show the unresolved references
	UnresolvedReferences []UnresolvedReferences `json:"unresolvedReferences"`
	// The time when the ImportOperation was created.
	CreatedAt time.Time `json:"createdAt"`
	// The last time When the ImportOperation was modified.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// The expiration time of the ImportOperation.
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

type ImportOperationStatus struct {
	// The ID of the [ImportOperation](#importoperation).
	OperationId *string `json:"operationId,omitempty"`
	// The validation state of the [ImportOperation](#importoperation).
	State ImportOperationState `json:"state"`
	// The validation errors for the [ImportOperation](#importoperation).
	// See [Errors](/error).
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
