package connect

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
)

type Connector struct {
	// Unique identifier of the Connector.
	ID string `json:"id"`
	// User-defined unique identifier of the Connector.
	Key *string `json:"key,omitempty"`
	// Current version of the Connector.
	Version int `json:"version"`
	// Name of the Connector.
	Name string `json:"name"`
	// Description of the Connector.
	Description *string `json:"description,omitempty"`
	// Integration types of the Connector. Can be used to filter search and query results.
	IntegrationTypes []IntegrationType `json:"integrationTypes"`
	// Owner of the Connector.
	Creator Creator `json:"creator"`
	// GitHub repository details of the Connector.
	Repository Repository `json:"repository"`
	// Configurations needed by Connectors for hosting. Loaded as environment variables in the workload.
	Configurations []ConnectorConfigurationApplication `json:"configurations"`
	// If `true`, only Composable Commerce Projects specified in `privateProjects` can access the Connector.
	Private bool `json:"private"`
	// If not empty, Connectors can only be deployed in these Regions. If empty, Connectors can be deployed in any [supported Region](hosts-and-authorization#hosts).
	SupportedRegions []Region `json:"supportedRegions"`
	// If `true`, the Connector is certified.
	Certified bool `json:"certified"`
	// URL to the documentation of the Connector.
	DocumentationUrl *string `json:"documentationUrl,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj Connector) MarshalJSON() ([]byte, error) {
	type Alias Connector
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

	if raw["configurations"] == nil {
		delete(raw, "configurations")
	}

	return json.Marshal(raw)

}

type ConnectorDraft struct {
	// User-defined unique identifier of the Connector.
	Key *string `json:"key,omitempty"`
	// Name of the Connector.
	Name string `json:"name"`
	// Description of the Connector.
	Description *string `json:"description,omitempty"`
	// Integration types of the Connector. Can be used to filter search and query results.
	IntegrationTypes []IntegrationType `json:"integrationTypes"`
	// Owner of the Connector.
	Creator Creator `json:"creator"`
	// GitHub repository details of the Connector.
	Repository Repository `json:"repository"`
	// If provided, Connectors can only be deployed in these Regions. If not provided, Connectors can be deployed in any [supported Region](hosts-and-authorization#hosts).
	SupportedRegions []Region `json:"supportedRegions"`
	// Composable Commerce Projects that can access the Connector. If empty, only the creator can access this Connector.
	PrivateProjects []string `json:"privateProjects"`
	// URL to the documentation of the Connector.
	DocumentationUrl *string `json:"documentationUrl,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConnectorDraft) MarshalJSON() ([]byte, error) {
	type Alias ConnectorDraft
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

	if raw["integrationTypes"] == nil {
		delete(raw, "integrationTypes")
	}

	if raw["supportedRegions"] == nil {
		delete(raw, "supportedRegions")
	}

	if raw["privateProjects"] == nil {
		delete(raw, "privateProjects")
	}

	return json.Marshal(raw)

}

type ConnectorStaged struct {
	// Unique identifier of the Connector.
	ID string `json:"id"`
	// User-defined unique identifier of the Connector.
	Key *string `json:"key,omitempty"`
	// Current version of the Connector.
	Version int `json:"version"`
	// Name of the Connector.
	Name string `json:"name"`
	// Description of the Connector.
	Description string `json:"description"`
	// Integration types of the Connector. Can be used to filter search and query results.
	IntegrationTypes []IntegrationType `json:"integrationTypes"`
	// Owner of the Connector.
	Creator Creator `json:"creator"`
	// GitHub repository details of the Connector.
	Repository Repository `json:"repository"`
	// Configurations needed by Connectors for hosting. Loaded as environment variables in the workload.
	Configurations []ConnectorConfigurationApplication `json:"configurations"`
	// If `true`, only Composable Commerce Projects specified in `privateProjects` can access the Connector.
	Private bool `json:"private"`
	// If `private` is true, only these Composable Commerce Projects can access the Connector.
	PrivateProjects []string `json:"privateProjects"`
	// If not empty, Connectors can only be deployed in these Regions. If empty, Connectors can be deployed in any [supported Region](hosts-and-authorization#hosts).
	SupportedRegions []Region `json:"supportedRegions"`
	// Comments made during the certification process.
	CertificationInfo *CertificationInfo `json:"certificationInfo,omitempty"`
	// Current status of the Connector.
	Status string `json:"status"`
	// The publishing request report of the Connector.
	PublishingReport *ConnectorReport `json:"publishingReport,omitempty"`
	// If `true`, the Connector is published and ready for use.
	AlreadyListed bool `json:"alreadyListed"`
	// If `true`, the ConnectorStaged data is different from the production [Connector](ctp:connect:type:Connector) data.
	HasChanges bool `json:"hasChanges"`
	// The previewable status of the ConnectorStaged.
	IsPreviewable string `json:"isPreviewable"`
	// The previewable request report.
	PreviewableReport *ConnectorReport `json:"previewableReport,omitempty"`
	// URL to the documentation of the Connector.
	DocumentationUrl *string `json:"documentationUrl,omitempty"`
}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with results containing an array of [ConnectorStaged](ctp:connect:type:ConnectorStaged).
*
 */
type ConnectorStagedPagedQueryResponse struct {
	// The maximum number of the ConnectorStaged returned.
	Limit int `json:"limit"`
	// The offset of the ConnectorStaged returned.
	Offset int `json:"offset"`
	// The number of ConnectorStaged returned.
	Count int `json:"count"`
	// The total number of ConnectorStaged matching the query.
	Total int `json:"total"`
	// ConnectorStaged matching the query.
	Results []ConnectorStaged `json:"results"`
}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with results containing an array of [Connector](ctp:connect:type:Connector).
*
 */
type ConnectorSearchPagedQueryResponse struct {
	// The maximum number of Connectors returned.
	Limit int `json:"limit"`
	// The offset of the Connectors returned.
	Offset int `json:"offset"`
	// The number of Connectors returned.
	Count int `json:"count"`
	// The total number of Connectors matching the query.
	Total int `json:"total"`
	// Connectors matching the query.
	Results []Connector `json:"results"`
}

type ConnectorUpdate struct {
	// Expected version of the Connector on which the changes apply.
	Version int `json:"version"`
	// Update actions to be performed on the Connector.
	Actions []ConnectorUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ConnectorUpdate) UnmarshalJSON(data []byte) error {
	type Alias ConnectorUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorConnectorUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type ConnectorUpdateAction interface{}

func mapDiscriminatorConnectorUpdateAction(input interface{}) (ConnectorUpdateAction, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["action"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'action'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "setCreatorCompany":
		obj := ConnectorSetCreatorCompanyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCreatorEmail":
		obj := ConnectorSetCreatorEmailAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCreatorName":
		obj := ConnectorSetCreatorNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCreatorTitle":
		obj := ConnectorSetCreatorTitleAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDescription":
		obj := ConnectorSetDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCreatorSupportUrl":
		obj := ConnectorSetCreatorSupportUrlAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setName":
		obj := ConnectorSetNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCreatorLogo":
		obj := ConnectorSetCreatorLogoAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCreatorNoOfContributors":
		obj := ConnectorSetCreatorNoOfContributorsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addPrivateProject":
		obj := ConnectorAddPrivateProjectAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removePrivateProject":
		obj := ConnectorRemovePrivateProjectAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setRepository":
		obj := ConnectorSetRepositoryAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setSupportedRegions":
		obj := ConnectorSetSupportedRegionsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDocumentationUrl":
		obj := ConnectorSetDocumentationUrlAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setIntegrationTypes":
		obj := ConnectorSetIntegrationTypesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "unlist":
		obj := ConnectorUnlistAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addCertificationComment":
		obj := ConnectorAddCertificationCommentAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "updatePreviewable":
		obj := ConnectorUpdatePreviewableAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "publish":
		obj := ConnectorPublishAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Updates the company of the creator.
*
 */
type ConnectorSetCreatorCompanyAction struct {
	// Value to set.
	CreatorCompany string `json:"creatorCompany"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConnectorSetCreatorCompanyAction) MarshalJSON() ([]byte, error) {
	type Alias ConnectorSetCreatorCompanyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCreatorCompany", Alias: (*Alias)(&obj)})
}

/**
*	Updates the email address of the creator.
*
 */
type ConnectorSetCreatorEmailAction struct {
	// Value to set.
	CreatorEmail string `json:"creatorEmail"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConnectorSetCreatorEmailAction) MarshalJSON() ([]byte, error) {
	type Alias ConnectorSetCreatorEmailAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCreatorEmail", Alias: (*Alias)(&obj)})
}

/**
*	Updates the name of the creator.
*
 */
type ConnectorSetCreatorNameAction struct {
	// Value to set.
	CreatorName string `json:"creatorName"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConnectorSetCreatorNameAction) MarshalJSON() ([]byte, error) {
	type Alias ConnectorSetCreatorNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCreatorName", Alias: (*Alias)(&obj)})
}

/**
*	Updates the title of the creator.
*
 */
type ConnectorSetCreatorTitleAction struct {
	// Value to set.
	CreatorTitle string `json:"creatorTitle"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConnectorSetCreatorTitleAction) MarshalJSON() ([]byte, error) {
	type Alias ConnectorSetCreatorTitleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCreatorTitle", Alias: (*Alias)(&obj)})
}

/**
*	Updates the description of the Connector.
*
 */
type ConnectorSetDescriptionAction struct {
	// Value to set.
	Description string `json:"description"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConnectorSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ConnectorSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

/**
*	Updates the support URL of the creator.
*
 */
type ConnectorSetCreatorSupportUrlAction struct {
	// Value to set.
	CreatorSupportUrl string `json:"creatorSupportUrl"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConnectorSetCreatorSupportUrlAction) MarshalJSON() ([]byte, error) {
	type Alias ConnectorSetCreatorSupportUrlAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCreatorSupportUrl", Alias: (*Alias)(&obj)})
}

/**
*	Updates the name of the Connector.
*
 */
type ConnectorSetNameAction struct {
	// Value to set.
	Name string `json:"name"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConnectorSetNameAction) MarshalJSON() ([]byte, error) {
	type Alias ConnectorSetNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setName", Alias: (*Alias)(&obj)})
}

/**
*	Updates the logo of the creator.
*
 */
type ConnectorSetCreatorLogoAction struct {
	// Value to set.
	LogoUrl string `json:"logoUrl"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConnectorSetCreatorLogoAction) MarshalJSON() ([]byte, error) {
	type Alias ConnectorSetCreatorLogoAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCreatorLogo", Alias: (*Alias)(&obj)})
}

/**
*	Updates the number of contributors of the creator.
*
 */
type ConnectorSetCreatorNoOfContributorsAction struct {
	// Value to set.
	CreatorNoOfContributors int `json:"creatorNoOfContributors"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConnectorSetCreatorNoOfContributorsAction) MarshalJSON() ([]byte, error) {
	type Alias ConnectorSetCreatorNoOfContributorsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCreatorNoOfContributors", Alias: (*Alias)(&obj)})
}

/**
*	Allow a Composable Commerce Project to access a private Connector.
*
*	Attempting to add a `privateProject` to a non-private ConnectorStaged returns the [ConnectorStagedNotPrivate](ctp:connect:type:ConnectorStagedNotPrivateError) error.
*
 */
type ConnectorAddPrivateProjectAction struct {
	// The Composable Commerce Project to add to `privateProjects`.
	PrivateProject string `json:"privateProject"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConnectorAddPrivateProjectAction) MarshalJSON() ([]byte, error) {
	type Alias ConnectorAddPrivateProjectAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addPrivateProject", Alias: (*Alias)(&obj)})
}

/**
*	Remove a Composable Commerce Project's access to a private Connector.
*
*	Attempting to remove a `privateProject` from a non-private ConnectorStaged returns the [ConnectorStagedNotPrivate](ctp:connect:type:ConnectorStagedNotPrivateError) error.
*
 */
type ConnectorRemovePrivateProjectAction struct {
	// The Composable Commerce Project to remove from `privateProjects`.
	PrivateProject string `json:"privateProject"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConnectorRemovePrivateProjectAction) MarshalJSON() ([]byte, error) {
	type Alias ConnectorRemovePrivateProjectAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removePrivateProject", Alias: (*Alias)(&obj)})
}

/**
*	Updates the GitHub repository details of the Connector.
*
 */
type ConnectorSetRepositoryAction struct {
	// New HTTPS or SSH GitHub URL to assign to the Connector.
	Url string `json:"url"`
	// New Git tag to assign to the Connector.
	Tag string `json:"tag"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConnectorSetRepositoryAction) MarshalJSON() ([]byte, error) {
	type Alias ConnectorSetRepositoryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setRepository", Alias: (*Alias)(&obj)})
}

/**
*	Updates the regions that the Connector can be deployed in.
*
 */
type ConnectorSetSupportedRegionsAction struct {
	// New value to set.
	Regions []Region `json:"regions"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConnectorSetSupportedRegionsAction) MarshalJSON() ([]byte, error) {
	type Alias ConnectorSetSupportedRegionsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSupportedRegions", Alias: (*Alias)(&obj)})
}

/**
*	Updates the documentation URL of the Connector.
*
 */
type ConnectorSetDocumentationUrlAction struct {
	// Value to set.
	DocumentationUrl string `json:"documentationUrl"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConnectorSetDocumentationUrlAction) MarshalJSON() ([]byte, error) {
	type Alias ConnectorSetDocumentationUrlAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDocumentationUrl", Alias: (*Alias)(&obj)})
}

/**
*	Updates the integration types of the Connector.
*
 */
type ConnectorSetIntegrationTypesAction struct {
	// New value to set.
	IntegrationTypes []IntegrationType `json:"integrationTypes"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConnectorSetIntegrationTypesAction) MarshalJSON() ([]byte, error) {
	type Alias ConnectorSetIntegrationTypesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setIntegrationTypes", Alias: (*Alias)(&obj)})
}

/**
*	Removes a certified and listed Connector from search results and listings. This update action does not affect deployed instances of the Connector.
*
 */
type ConnectorUnlistAction struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConnectorUnlistAction) MarshalJSON() ([]byte, error) {
	type Alias ConnectorUnlistAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "unlist", Alias: (*Alias)(&obj)})
}

/**
*	Add a comment during the certification process of the Connector.
*
 */
type ConnectorAddCertificationCommentAction struct {
	// Comment to add.
	Comment string `json:"comment"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConnectorAddCertificationCommentAction) MarshalJSON() ([]byte, error) {
	type Alias ConnectorAddCertificationCommentAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addCertificationComment", Alias: (*Alias)(&obj)})
}

/**
*	Requests the previewable status of a ConnectorStaged. Previewable ConnectorStaged should be deployed for testing and preview purposes. You can only deploy previewable ConnectorStaged to Projects listed in `privateProjects`.
*
*	After using this update action, the status of `isPreviewable` will change to `pending`. Following validation, the status of `isPreviewable` will change to `true` if the previewable status is granted, or `false` if it is rejected.
*	In the case of a `false` status, contact the [Connect support team](https://support.commercetools.com/) regarding any issues raised during the validation process.
*
*	Requesting the previewable status for a ConnectorStaged that is currently being reviewed returns the [ConnectorStagedPreviewRequestUnderProcess](ctp:connect:type:ConnectorStagedPreviewRequestUnderProcessError) error.
*
 */
type ConnectorUpdatePreviewableAction struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConnectorUpdatePreviewableAction) MarshalJSON() ([]byte, error) {
	type Alias ConnectorUpdatePreviewableAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "updatePreviewable", Alias: (*Alias)(&obj)})
}

/**
*	Starts the Connector publishing process. You will be unable to update the Connector until the process completes.
*
*	`certification` should only be `true` if you want to make the Connector publicly available.
*
 */
type ConnectorPublishAction struct {
	// - Set to `false` to make the Connector private and skip the [certification process](/certification). The published Connector will only be deployable on Projects listed in `ConnectorStaged.privateProjects`.
	// - Set to `true` to make the Connector public. This will submit the ConnectorStaged to the [certification process](/certification). After completing the certification process, the Connector will become publicly available.
	Certification bool `json:"certification"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConnectorPublishAction) MarshalJSON() ([]byte, error) {
	type Alias ConnectorPublishAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "publish", Alias: (*Alias)(&obj)})
}
