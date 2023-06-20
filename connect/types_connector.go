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
	// Owner of the Connector.
	Creator Creator `json:"creator"`
	// GitHub repository details of the Connector.
	Repository Repository `json:"repository"`
	// If provided, Connectors can only be deployed in these Regions. If not provided, Connectors can be deployed in any [supported Region](hosts-and-authorization#hosts).
	SupportedRegions []Region `json:"supportedRegions"`
	// Contains Composable Commerce Project keys listed using the format "{region}:{projectKey}". If `private` is true, only these Composable Commerce Projects can access the Connector.
	PrivateProjects []string `json:"privateProjects"`
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
	// Owner of the Connector.
	Creator Creator `json:"creator"`
	// GitHub repository details of the Connector.
	Repository Repository `json:"repository"`
	// Configurations needed by Connectors for hosting. Loaded as environment variables in the workload.
	Configurations []ConnectorConfigurationApplication `json:"configurations"`
	// If `true`, only Composable Commerce Projects specified in `privateProjects` can access the Connector.
	Private bool `json:"private"`
	// Contains Composable Commerce Project keys listed using the format "{region}:{projectKey}". If `private` is true, only these Composable Commerce Projects can access the Connector once listed.
	PrivateProjects []string `json:"privateProjects"`
	// If not empty, Connectors can only be deployed in these Regions. If empty, Connectors can be deployed in any [supported Region](hosts-and-authorization#hosts).
	SupportedRegions []Region `json:"supportedRegions"`
	// Comments made during the certification process.
	CertificationInfo *CertificationInfo `json:"certificationInfo,omitempty"`
	// Current status of the Connector.
	Status string `json:"status"`
	// If `true`, the Connector is published and ready for use.
	AlreadyListed bool `json:"alreadyListed"`
	// If `true`, the ConnectorStaged data is different from the production [Connector](ctp:connect:type:Connector) data.
	HasChanges bool `json:"hasChanges"`
	// The previewable status of the ConnectorStaged.
	IsPreviewable *string `json:"isPreviewable,omitempty"`
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
	case "setConfigurations":
		obj := ConnectorSetConfigurationsAction{}
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
	case "setName":
		obj := ConnectorSetNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCreatorNoOfDevelopers":
		obj := ConnectorSetCreatorNoOfDevelopersAction{}
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
	case "triggerCertification":
		obj := ConnectorTriggerCertificationAction{}
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
	}
	return nil, nil
}

/**
*	Updates the creator's company.
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
*	Set the configuration for the Connector.
*
 */
type ConnectorSetConfigurationsAction struct {
	// Value to set. If null or empty, any existing value is removed.
	Configurations []ConnectorConfigurationApplication `json:"configurations"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConnectorSetConfigurationsAction) MarshalJSON() ([]byte, error) {
	type Alias ConnectorSetConfigurationsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setConfigurations", Alias: (*Alias)(&obj)})
}

/**
*	Updates the creator's contact email address.
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
*	Updates the creator's name.
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
*	Updates the creator's title.
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
*	Updates the description in Connector.
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
*	Updates the number of developers within the creator's company.
*
 */
type ConnectorSetCreatorNoOfDevelopersAction struct {
	// Value to set.
	CreatorNoOfDevelopers int `json:"creatorNoOfDevelopers"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConnectorSetCreatorNoOfDevelopersAction) MarshalJSON() ([]byte, error) {
	type Alias ConnectorSetCreatorNoOfDevelopersAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCreatorNoOfDevelopers", Alias: (*Alias)(&obj)})
}

/**
*	Allow a Composable Commerce Project to access a private Connector.
*
*	Attempting to add a `privateProject` to a non-private ConnectorStaged returns the [ConnectorStagedNotPrivate](ctp:connect:type:ConnectorStagedNotPrivateError) error.
*
 */
type ConnectorAddPrivateProjectAction struct {
	// The Composable Commerce Project to add to `privateProjects`. Must use the format "{region}:{projectKey}".
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
	// The Composable Commerce Project to remove from `privateProjects`. Must use the format "{region}:{projectKey}".
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
*	Update the GitHub repository details for the Connector.
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
*	Updates the Regions that the Connector can be deployed in.
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
*	After triggering the certification process you cannot update the Connector until the review process is complete.
*
*	Triggering the certification process when the Connector is already certified, or when the Connector is in the certification process, returns the [ConnectorAlreadyCertified](ctp:connect:type:ConnectorAlreadyCertifiedError) and [ConnectorStagedInCertification](ctp:connect:type:ConnectorStagedInCertificationError) errors, respectively.
*
 */
type ConnectorTriggerCertificationAction struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConnectorTriggerCertificationAction) MarshalJSON() ([]byte, error) {
	type Alias ConnectorTriggerCertificationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "triggerCertification", Alias: (*Alias)(&obj)})
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
*	Requests previewable status for the ConnectorStaged. A previewable ConnectorStaged can used in a Deployment for testing/preview purposes.
*
*	After using this update action the status of `isPreviewable` will change to `pending` and your request will be reviewed by the Connect team. After being reviewed, the status of `isPreviewable` will change to `true` if the previewable status has been granted or `false` if rejected.
*
*	If `false`, the Connect team will contact you regarding any issues raised during the review process.
*
*	Requesting previewable status for a ConnectorStaged that is currently being reviewed for previewable status returns the [ConnectorStagedPreviewRequestUnderProcess](ctp:connect:type:ConnectorStagedPreviewRequestUnderProcessError) error.
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
