package checkout

// Generated file, please do not change!!!

type PaymentIntegrationDraft struct {
	// User-defined unique identifier of the PaymentIntegration. `MinLength: 2 ​MaxLength: 256 ​Pattern: ^[A-Za-z0-9_-]+$`
	Key *string `json:"key,omitempty"`
	// Reference to an Application associated with the PaymentIntegration.
	Application ApplicationReference `json:"application"`
	// Indicates whether the Payment Integration is active or not. By default in Inactive.
	Status *PaymentIntegrationStatus `json:"status,omitempty"`
	// Type of the PaymentIntegration. ie. card, paypal, applepay.
	Type string `json:"type"`
	// Name of the PaymentIntegration in the Merchant Center.
	Name string `json:"name"`
	// Predicate in JSONata format. Applies only when `componentType` is `Component`.
	Predicate *string `json:"predicate,omitempty"`
	// Indicates the component type of payment integration.
	ComponentType PaymentComponentType `json:"componentType"`
	// Connector deployment info for the payment integration.
	ConnectorDeployment ConnectorDeploymentReference `json:"connectorDeployment"`
	// Customization for the information that is displayed in the UI. Applies only when `componentType` is `Component`
	DisplayInfo *DisplayInfo `json:"displayInfo,omitempty"`
	// Sorting info of the payment integration.
	SortingInfo *SortingInfo `json:"sortingInfo,omitempty"`
	// Configuration of the payment integration automatic reversal.
	AutomatedReversalConfiguration *AutomatedReversalConfiguration `json:"automatedReversalConfiguration,omitempty"`
}
