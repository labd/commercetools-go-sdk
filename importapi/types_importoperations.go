// Generated file, please do not change!!!
package importapi

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
	// The version of the impmorted resource when the import was successful.
	ResourceVersion *int `json:"resourceVersion,omitempty"`
	// Contains an error if the import of the resource was not successful. See [Errors](/error).
	Errors []ErrorObject `json:"errors,omitempty"`
	// In case of unresolved status this array will show the unresolved references
	UnresolvedReferences []UnresolvedReferences `json:"unresolvedReferences,omitempty"`
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

	return nil
}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) for Import Operations.
*
 */
type ImportOperationPagedResponse struct {
	// The number of results requested in the query request.
	Limit int `json:"limit"`
	// The number of elements skipped, not a page number.
	// Supplied by the client or the server default.
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
	Errors []ErrorObject `json:"errors,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ImportOperationStatus) UnmarshalJSON(data []byte) error {
	type Alias ImportOperationStatus
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	return nil
}
