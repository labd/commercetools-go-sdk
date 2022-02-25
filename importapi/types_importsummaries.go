package importapi

// Generated file, please do not change!!!

/**
*	Describes the status of an [ImportContainer](/import-container#importcontainer) by the number of resources in each [Processing State](/processing-state#processingstate).
*	Can be used to monitor the import progress per [Import Container](/import-container).
*
 */
type ImportSummary struct {
	// The import status of an [ImportContainer](/import-container#importcontainer) given by the number of resources in each [Processing State](/processing-state#processingstate).
	States OperationStates `json:"states"`
	// The total number of [ImportOperations](/import-operation#importoperation) received for this Import Summary.
	Total int `json:"total"`
}

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
}
