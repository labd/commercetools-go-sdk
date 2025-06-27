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
	// Deployment type. If a [ConnectorStaged](ctp:connect:type:ConnectorStaged) is referenced in the `connector` field, then the field can only be `preview`. If a [Connector](ctp:connect:type:Connector) is referenced in the `connector` field, then the field can either be `sandbox` or `production`.
	Type string `json:"type"`
	// Details of the [Connector](ctp:connect:type:Connector) being deployed.
	Connector DeploymentConnector `json:"connector"`
	// Global configuration applied to all applications in the deployment.
	GlobalConfiguration *DeploymentGlobalConfiguration `json:"globalConfiguration,omitempty"`
	// API Client credentials used for deployment.
	ApiClient *DeploymentApiClient `json:"apiClient,omitempty"`
	// Build reference id and outcome of the Deployment.
	Details DeploymentDetails `json:"details"`
	// Region of the Deployment.
	DeployedRegion string `json:"deployedRegion"`
	// Application deployments needed by the connector for hosting and configuration, refer to Connector configurations for details.
	Applications []ApplicationDeployment `json:"applications"`
	// If `true`, the Deployment is a preview.
	Preview bool `json:"preview"`
	// The current status of the Deployment.
	Status string `json:"status"`
}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with results containing an array of [Deployment](ctp:connect:type:Deployment).
*
 */
type DeploymentPagedQueryResponse struct {
	// The limit of the Deployments returned.
	Limit int `json:"limit"`
	// The offset of the Deployments returned.
	Offset int `json:"offset"`
	// The number of Deployments returned.
	Count int `json:"count"`
	// The total number of Deployments available.
	Total int `json:"total"`
	// Deployments matching the query.
	Results []Deployment `json:"results"`
}

type DeploymentDraft struct {
	// User-defined unique identifier for the Deployment.
	Key *string `json:"key,omitempty"`
	// Deployment type. If a [ConnectorStaged](ctp:connect:type:ConnectorStaged) is referenced in the `connector` field, you can send only `preview`. If a [Connector](ctp:connect:type:Connector) is referenced in the `connector` field, the value defaults to `sandbox`. However, you can send `production` when you are ready to deploy the Connector to your production environment.
	Type *string `json:"type,omitempty"`
	// Reference to the [Connector](ctp:connect:type:Connector) or [ConnectorStaged](ctp:connect:type:ConnectorStaged) being deployed.
	Connector ConnectorReference `json:"connector"`
	// Configuration values needed by the [Connector](ctp:connect:type:Connector) for hosting. Keys should match those in the Connector's `configurations` field.
	Configurations []DeploymentConfigurationApplication `json:"configurations"`
	// Global configuration values needed by the [Connector](ctp:connect:type:Connector) for hosting. Keys should match those in the Connector's `globalConfiguration` field.
	GlobalConfiguration *DeploymentGlobalConfiguration `json:"globalConfiguration,omitempty"`
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
*	Updates configuration values and redeploys a Deployment in `Draft`, `Deployed`, `Failed`, or `UndeployFailed` states. The new configuration values must include at least one valid application. Ensure that all deployment configuration keys and values match those specified in the application's `connect.yaml` file.
*
*	Specific error codes:
*	- [DeploymentUnknownApplicationConfiguration](ctp:connect:type:DeploymentUnknownApplicationConfigurationError)
*	- [DeploymentUnknownApplicationConfigurationKey](ctp:connect:type:DeploymentUnknownApplicationConfigurationKeyError)
*	- [DeploymentInvalidStatusTransition](ctp:connect:type:DeploymentInvalidStatusTransitionError)
*	- [DeploymentApplicationDoNotBelong](ctp:connect:type:DeploymentApplicationDoNotBelongError)
*	- [DeploymentApplicationRequiredError](ctp:connect:type:DeploymentApplicationRequiredError)
*	- [DeploymentConnectorUpdateFailure](ctp:connect:type:DeploymentConnectorUpdateFailureError)
*	- [DeploymentUnknownGlobalConfigurationKeyError](ctp:connect:type:DeploymentUnknownGlobalConfigurationKeyError)
*	- [DeploymentEmptyRequiredGlobalConfigurationKeyError](ctp:connect:type:DeploymentEmptyRequiredGlobalConfigurationKeyError)
*
 */
type DeploymentRedeployAction struct {
	// Whether the scripts execution should be skipped during redeployment.
	SkipScripts *bool `json:"skipScripts,omitempty"`
	// New configuration values for Deployment. If empty, any existing value is unchanged.
	ConfigurationValues []DeploymentConfigurationApplication `json:"configurationValues"`
	// Whether the Deployment Connector should be updated to its latest state during redeployment.
	UpdateConnector *bool `json:"updateConnector,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DeploymentRedeployAction) MarshalJSON() ([]byte, error) {
	type Alias DeploymentRedeployAction
	data, err := json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "redeploy", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["configurationValues"] == nil {
		delete(raw, "configurationValues")
	}

	return json.Marshal(raw)

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

/**
*	A cursor paged query result containing an array of [DeploymentLog](ctp:connect:type:DeploymentLog).
*
 */
type DeploymentLogCursorPagedQueryResponse struct {
	// Array of DeploymentLog objects.
	Data []DeploymentLog `json:"data"`
	// The next page token.
	Next string `json:"next"`
}
