package connect

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type Deployment struct {
	// Unique identifier of the Deployment.
	ID string `json:"id"`
	// User-defined unique identifier of the Deployment.
	Key *string `json:"key,omitempty"`
	// Current version of the Deployment.
	Version int `json:"version"`
	// Details of the [Connector](ctp:connect:type:Connector) being deployed.
	Connector DeploymentConnector `json:"connector"`
	// Build reference id and outcome of the Deployment.
	Details DeploymentDetails `json:"details"`
	// Region of the Deployment.
	DeployedRegion string `json:"deployedRegion"`
	// Application deployments needed by the connector for hosting and configuration, refer to Connector configurations for details.
	Applications []ApplicationDeployment `json:"applications"`
	// Indicates the current status of the Deployment.
	Status string `json:"status"`
}

type DeploymentDraft struct {
	// User-defined unique identifier for the Deployment.
	Key *string `json:"key,omitempty"`
	// Reference to the [Connector](ctp:connect:type:Connector) being deployed.
	Connector ConnectorReference `json:"connector"`
	// Configuration values needed by the [Connector](ctp:connect:type:Connector) for hosting. Keys should match those in the Connector's `configurations` field.
	Configurations []DeploymentConfigurationApplication `json:"configurations"`
	// Region of Deployment.
	Region string `json:"region"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DeploymentDraft) MarshalJSON() ([]byte, error) {
	type Alias DeploymentDraft
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

type DeploymentUpdate struct {
	// Expected version of the Deployment on which the changes apply.
	Version int `json:"version"`
	// Update actions to be performed on the Deployment.
	Actions []DeploymentUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DeploymentUpdate) UnmarshalJSON(data []byte) error {
	type Alias DeploymentUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorDeploymentUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type DeploymentUpdateAction interface{}

func mapDiscriminatorDeploymentUpdateAction(input interface{}) (DeploymentUpdateAction, error) {
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
	case "redeploy":
		obj := DeploymentRedeployAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Update the configuration values and redeploy the Deployment.
*
*	Attempting to redeploy a Deployment that does not require configuration or a configuration key for an application returns the [DeploymentUnknownApplicationConfiguration](ctp:connect:type:DeploymentUnknownApplicationConfigurationError) or [DeploymentUnknownApplicationConfigurationKey](ctp:connect:type:DeploymentUnknownApplicationConfigurationKeyError) errors, respectively.
*
*	Attempting to redeploy a Deployment with a `Queued`, `Deploying`, or `Undeploying` [DeploymentStatus](ctp:connect:type:DeploymentStatus) returns a [DeploymentInvalidStatusTransition](ctp:connect:type:DeploymentInvalidStatusTransitionError) error.
*
 */
type DeploymentRedeployAction struct {
	// New configuration values for Deployment.
	ConfigurationValues []DeploymentConfigurationApplication `json:"configurationValues"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DeploymentRedeployAction) MarshalJSON() ([]byte, error) {
	type Alias DeploymentRedeployAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "redeploy", Alias: (*Alias)(&obj)})
}

/**
*	The data representation of a Deployment log.
*
 */
type DeploymentLog struct {
	Type string `json:"type"`
	// Unique identifier of the [Deployment](ctp:connect:type:Deployment).
	DeploymentId string `json:"deploymentId"`
	// Name of the Connect application.
	ApplicationName string `json:"applicationName"`
	// Severity of the log message.
	Severity string `json:"severity"`
	// Timestamp of the log message.
	Timestamp time.Time `json:"timestamp"`
	// Event details of the log message.
	Details interface{} `json:"details"`
}
