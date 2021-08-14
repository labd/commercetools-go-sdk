// Generated file, please do not change!!!
package importapi

/**
*	Describes the status of an [ImportSink](/import-sink#importsink) by the number of resources in each [Processing State](/processing-state#the-list-of-processing-states).
*	Can be used to monitor the import progress per [Import Sink](/import-sink).
*
 */
type ImportSummary struct {
	// The import status of an [ImportSink](/import-sink#importsink) given by the number of resources in each [Processing State](/processing-state#the-list-of-processing-states).
	States OperationStates `json:"states"`
	// The total number of [ImportOperations](/import-operation#importoperation) received for this Import Summary.
	Total int `json:"total"`
}

type OperationStates struct {
	// The number of resources in the `ValidationFailed` state.
	ValidationFailed int `json:"ValidationFailed"`
	// The number of resources in the `Unresolved` state.
	Unresolved int `json:"Unresolved"`
	// The number of resources in the `WaitForMasterVariant` state.
	WaitForMasterVariant int `json:"WaitForMasterVariant"`
	// The number of resources in the `Imported` state.
	Imported int `json:"Imported"`
	// The number of resources in the `Rejected` state.
	Rejected int `json:"Rejected"`
}
