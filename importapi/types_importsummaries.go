package importapi

// Generated file, please do not change!!!

/**
*	The current status of [ImportOperations](ctp:import:type:ImportOperation) in an [ImportContainer](ctp:import:type:ImportContainer).
*
 */
type ImportSummary struct {
	// The current [ProcessingStates](ctp:import:type:ProcessingState) of ImportOperations in an ImportContainer.
	States OperationStates `json:"states"`
	// The total number of ImportOperations in `states`.
	Total int `json:"total"`
}

/**
*	The number of [ImportOperations](ctp:import:type:ImportOperation) in each [ProcessingState](ctp:import:type:ProcessingState).
 */
type OperationStates struct {
	// The number of ImportOperations in the `processing` state.
	Processing int `json:"processing"`
	// The number of ImportOperations in the `validationFailed` state.
	ValidationFailed int `json:"validationFailed"`
	// The number of ImportOperations in the `unresolved` state.
	Unresolved int `json:"unresolved"`
	// The number of ImportOperations in the `waitForMasterVariant` state.
	WaitForMasterVariant int `json:"waitForMasterVariant"`
	// The number of ImportOperations in the `imported` state.
	Imported int `json:"imported"`
	// The number of ImportOperations in the `rejected` state.
	Rejected int `json:"rejected"`
	// The number of ImportOperations in the `canceled` state.
	Canceled int `json:"canceled"`
}
