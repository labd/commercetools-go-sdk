package importapi

// Generated file, please do not change!!!

/**
*	Describes the status of an [ImportContainer](ctp:import:type:ImportContainer) by the number of resources in each [ProcessingState](ctp:import:type:ProcessingState).
*	Can be used to monitor the import progress per [Import Container](ctp:import:type:ImportContainer).
*
 */
type ImportSummary struct {
	// The import status of an [ImportContainer](ctp:import:type:ImportContainer) given by the number of resources in each [ProcessingState](ctp:import:type:ProcessingState).
	States OperationStates `json:"states"`
	// The total number of [ImportOperations](ctp:import:type:ImportOperation) received for this Import Summary.
	Total int `json:"total"`
}

/**
*	The number of resources in each [ProcessingState](ctp:import:type:ProcessingState).
 */
type OperationStates struct {
	// The number of resources in the `processing` state.
	Processing int `json:"processing"`
	// The number of resources in the `validationFailed` state.
	ValidationFailed int `json:"validationFailed"`
	// The number of resources in the `unresolved` state.
	Unresolved int `json:"unresolved"`
	// The number of resources in the `waitForMasterVariant` state.
	WaitForMasterVariant int `json:"waitForMasterVariant"`
	// The number of resources in the `imported` state.
	Imported int `json:"imported"`
	// The number of resources in the `rejected` state.
	Rejected int `json:"rejected"`
	// The number of resources in the `canceled` state.
	Canceled int `json:"canceled"`
}
