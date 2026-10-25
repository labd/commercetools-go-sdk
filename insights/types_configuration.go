package insights

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

/**
*	Draft to create a Platform Insights configuration for your Project.
*
 */
type ProjectConfigurationDraft struct {
	// Whether the Platform Insights configuration for your Project is active.
	Active *bool `json:"active,omitempty"`
	// List of supported APM providers.
	Providers []ProviderDraft `json:"providers"`
	// List of Attributes to add to a metric.
	AdditionalAttributes []AttributeDraft `json:"additionalAttributes"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProjectConfigurationDraft) UnmarshalJSON(data []byte) error {
	type Alias ProjectConfigurationDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Providers {
		var err error
		obj.Providers[i], err = mapDiscriminatorProviderDraft(obj.Providers[i])
		if err != nil {
			return err
		}
	}
	for i := range obj.AdditionalAttributes {
		var err error
		obj.AdditionalAttributes[i], err = mapDiscriminatorAttributeDraft(obj.AdditionalAttributes[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProjectConfigurationDraft) MarshalJSON() ([]byte, error) {
	type Alias ProjectConfigurationDraft
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

	if raw["additionalAttributes"] == nil {
		delete(raw, "additionalAttributes")
	}

	return json.Marshal(raw)

}

/**
*	Update actions to be performed on the Platform Insights configuration for your Project.
*
 */
type ConfigurationUpdate struct {
	Actions []ConfigurationUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ConfigurationUpdate) UnmarshalJSON(data []byte) error {
	type Alias ConfigurationUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorConfigurationUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type ConfigurationUpdateAction interface{}

func mapDiscriminatorConfigurationUpdateAction(input interface{}) (ConfigurationUpdateAction, error) {
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
	case "activate":
		obj := ConfigurationActivateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "deactivate":
		obj := ConfigurationDeactivateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setActive":
		obj := ConfigurationSetActiveAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeProvider":
		obj := ConfigurationRemoveProviderAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setEventTypes":
		obj := ConfigurationSetEventTypesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setProvider":
		obj := ConfigurationSetProviderAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Provider != nil {
			var err error
			obj.Provider, err = mapDiscriminatorProviderDraft(obj.Provider)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "setAttributes":
		obj := ConfigurationSetAttributesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		for i := range obj.Attributes {
			var err error
			obj.Attributes[i], err = mapDiscriminatorAttributeDraft(obj.Attributes[i])
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "addAttribute":
		obj := ConfigurationAddAttributeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Attribute != nil {
			var err error
			obj.Attribute, err = mapDiscriminatorAttributeDraft(obj.Attribute)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "removeAttribute":
		obj := ConfigurationRemoveAttributeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Activates a Platform Insights configuration for your Project.
*
 */
type ConfigurationActivateAction struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConfigurationActivateAction) MarshalJSON() ([]byte, error) {
	type Alias ConfigurationActivateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "activate", Alias: (*Alias)(&obj)})
}

/**
*	Deactivates a Platform Insights configuration for your Project.
*
 */
type ConfigurationDeactivateAction struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConfigurationDeactivateAction) MarshalJSON() ([]byte, error) {
	type Alias ConfigurationDeactivateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "deactivate", Alias: (*Alias)(&obj)})
}

/**
*	Activates a provider Platform Insights configuration for your Project.
*
 */
type ConfigurationSetActiveAction struct {
	// List of supported APM providers for your Platform Insights configuration.
	Provider ProviderTypeId `json:"provider"`
	Active   bool           `json:"active"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConfigurationSetActiveAction) MarshalJSON() ([]byte, error) {
	type Alias ConfigurationSetActiveAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setActive", Alias: (*Alias)(&obj)})
}

/**
*	Removes an APM provider from a Platform Insights configuration for your Project.
*
 */
type ConfigurationRemoveProviderAction struct {
	// List of supported APM providers for your Platform Insights configuration.
	Provider ProviderTypeId `json:"provider"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConfigurationRemoveProviderAction) MarshalJSON() ([]byte, error) {
	type Alias ConfigurationRemoveProviderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeProvider", Alias: (*Alias)(&obj)})
}

type ConfigurationSetEventTypesAction struct {
	// Provider to be updated in the Platform Insights configuration for your Project.
	Provider ProviderTypeId `json:"provider"`
	// List of eventTypes to be used for this provider.
	EventTypes []EventType `json:"eventTypes"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConfigurationSetEventTypesAction) MarshalJSON() ([]byte, error) {
	type Alias ConfigurationSetEventTypesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setEventTypes", Alias: (*Alias)(&obj)})
}

/**
*	Sets an APM provider to a Platform Insights configuration for your Project.
*
 */
type ConfigurationSetProviderAction struct {
	// New APM provider to be set for your Platform Insights configuration.
	Provider ProviderDraft `json:"provider"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ConfigurationSetProviderAction) UnmarshalJSON(data []byte) error {
	type Alias ConfigurationSetProviderAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Provider != nil {
		var err error
		obj.Provider, err = mapDiscriminatorProviderDraft(obj.Provider)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConfigurationSetProviderAction) MarshalJSON() ([]byte, error) {
	type Alias ConfigurationSetProviderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setProvider", Alias: (*Alias)(&obj)})
}

/**
*	Sets the Attributes that are to be added to the events of your Platform Insights configuration.
*
 */
type ConfigurationSetAttributesAction struct {
	// List of Attributes to be set for the events of your Platform Insights configuration.
	Attributes []AttributeDraft `json:"attributes"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ConfigurationSetAttributesAction) UnmarshalJSON(data []byte) error {
	type Alias ConfigurationSetAttributesAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Attributes {
		var err error
		obj.Attributes[i], err = mapDiscriminatorAttributeDraft(obj.Attributes[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConfigurationSetAttributesAction) MarshalJSON() ([]byte, error) {
	type Alias ConfigurationSetAttributesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAttributes", Alias: (*Alias)(&obj)})
}

/**
*	Adds Attributes that are to be included in the events of your Platform Insights configuration.
*
 */
type ConfigurationAddAttributeAction struct {
	// Attribute to be set for your Platform Insights configuration.
	Attribute AttributeDraft `json:"attribute"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ConfigurationAddAttributeAction) UnmarshalJSON(data []byte) error {
	type Alias ConfigurationAddAttributeAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Attribute != nil {
		var err error
		obj.Attribute, err = mapDiscriminatorAttributeDraft(obj.Attribute)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConfigurationAddAttributeAction) MarshalJSON() ([]byte, error) {
	type Alias ConfigurationAddAttributeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addAttribute", Alias: (*Alias)(&obj)})
}

/**
*	Removes Attributes from your Platform Insights configuration.
*
 */
type ConfigurationRemoveAttributeAction struct {
	// Attribute to be removed from your Platform Insights configuration.
	Attribute string `json:"attribute"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConfigurationRemoveAttributeAction) MarshalJSON() ([]byte, error) {
	type Alias ConfigurationRemoveAttributeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeAttribute", Alias: (*Alias)(&obj)})
}

/**
*	List of event types to be exported to a supported APM provider.
*
 */
type EventType string

const (
	EventTypeLogs    EventType = "Logs"
	EventTypeMetrics EventType = "Metrics"
)

/**
*	List of supported APM providers for your Platform Insights configuration.
*
 */
type ProviderTypeId string

const (
	ProviderTypeIdNewRelic            ProviderTypeId = "NewRelic"
	ProviderTypeIdDynatraceSaaS       ProviderTypeId = "DynatraceSaaS"
	ProviderTypeIdDynatraceActiveGate ProviderTypeId = "DynatraceActiveGate"
	ProviderTypeIdDatadog             ProviderTypeId = "Datadog"
	ProviderTypeIdOtlp                ProviderTypeId = "Otlp"
	ProviderTypeIdOtlpHttp            ProviderTypeId = "OtlpHttp"
)

/**
*	APM provider details for your Platform Insights configuration.
*
 */
type ProviderDraft interface{}

func mapDiscriminatorProviderDraft(input interface{}) (ProviderDraft, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "NewRelic":
		obj := NewRelicProviderDraft{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "DynatraceSaaS":
		obj := DynatraceSaaSProviderDraft{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "DynatraceActiveGate":
		obj := DynatraceActiveGateProviderDraft{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "Datadog":
		obj := DatadogProviderDraft{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "Otlp":
		obj := OtlpProviderDraft{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Auth != nil {
			var err error
			obj.Auth, err = mapDiscriminatorOtlpAuthDraft(obj.Auth)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "OtlpHttp":
		obj := OtlpHttpProviderDraft{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Auth != nil {
			var err error
			obj.Auth, err = mapDiscriminatorOtlpAuthDraft(obj.Auth)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	}
	return nil, nil
}

type NewRelicProviderDraft struct {
	// Whether this provider configuration is active.
	Active *bool `json:"active,omitempty"`
	// List of event types to be exported to the APM provider.
	EventTypes []EventType `json:"eventTypes"`
	// Region in which your NewRelic account is based. For example, EU.
	Region NewRelicRegion `json:"region"`
	// API key (type `Ingest - License`) that's used to authenticate your NewRelic account. If you don't specify the API key, an existing key will
	// be kept.
	ApiKey *string `json:"apiKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj NewRelicProviderDraft) MarshalJSON() ([]byte, error) {
	type Alias NewRelicProviderDraft
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "NewRelic", Alias: (*Alias)(&obj)})
}

type DynatraceSaaSProviderDraft struct {
	// Whether this provider configuration is active.
	Active *bool `json:"active,omitempty"`
	// List of event types to be exported to the APM provider.
	EventTypes []EventType `json:"eventTypes"`
	// Unique user-defined environment identifier of your DynatraceSaaS account.
	EnvironmentId string `json:"environmentId"`
	// API key used to authenticate your DynatraceSaaS account. The required scopes are `Ingest-Logs` and `Ingest-Metrics`.
	// If you don't specify the API key, an existing key will be kept.
	ApiKey *string `json:"apiKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DynatraceSaaSProviderDraft) MarshalJSON() ([]byte, error) {
	type Alias DynatraceSaaSProviderDraft
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "DynatraceSaaS", Alias: (*Alias)(&obj)})
}

type DynatraceActiveGateProviderDraft struct {
	// Whether this provider configuration is active.
	Active *bool `json:"active,omitempty"`
	// List of event types to be exported to the APM provider.
	EventTypes []EventType `json:"eventTypes"`
	// Your ActiveGate domain.
	ActiveGateDomain string `json:"activeGateDomain"`
	// Your ActiveGate domain port.
	ActiveGatePort *float64 `json:"activeGatePort,omitempty"`
	// Unique user-defined environment identifier of your DynatraceActiveGate account.
	EnvironmentId string `json:"environmentId"`
	// API key used to authenticate your DynatraceActiveGate account. The required scopes are `Ingest-Logs` and `Ingest-Metrics`
	// If you don't specify the API key, an existing key will be kept.
	ApiKey *string `json:"apiKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DynatraceActiveGateProviderDraft) MarshalJSON() ([]byte, error) {
	type Alias DynatraceActiveGateProviderDraft
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "DynatraceActiveGate", Alias: (*Alias)(&obj)})
}

type DatadogProviderDraft struct {
	// Whether this provider configuration is active.
	Active *bool `json:"active,omitempty"`
	// List of event types to be exported to the APM provider.
	EventTypes []EventType `json:"eventTypes"`
	// Website in which your Datadog account is hosted.
	Site DatadogSite `json:"site"`
	// API key used to authenticate your Datadog account. If you don't specify the API key, an existing key will be kept.
	ApiKey *string `json:"apiKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DatadogProviderDraft) MarshalJSON() ([]byte, error) {
	type Alias DatadogProviderDraft
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "Datadog", Alias: (*Alias)(&obj)})
}

/**
*	Authentication details used to connect to an OTLP provider.
*
 */
type OtlpAuth interface{}

func mapDiscriminatorOtlpAuth(input interface{}) (OtlpAuth, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "OAuth2":
		obj := OtlpAuthOAuth2{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Supported authentication methods for connecting to an OTLP provider.
*
 */
type OtlpAuthType string

const (
	OtlpAuthTypeOAuth2 OtlpAuthType = "OAuth2"
)

/**
*	Authenticates with an OTLP provider using the OAuth 2.0 client credentials flow.
*	The `clientSecret` used to obtain the access token is write-only and isn't returned in the response.
*
 */
type OtlpAuthOAuth2 struct {
	// Client ID used to authenticate with your OTLP provider using the OAuth 2.0 client credentials flow.
	ClientId string `json:"clientId"`
	// URL of the OAuth 2.0 token endpoint used to obtain an access token for your OTLP provider.
	TokenUrl string `json:"tokenUrl"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OtlpAuthOAuth2) MarshalJSON() ([]byte, error) {
	type Alias OtlpAuthOAuth2
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OAuth2", Alias: (*Alias)(&obj)})
}

type OtlpProviderHeader struct {
	// Name of the header.
	Name string `json:"name"`
}

type OtlpProviderDraft struct {
	// Whether this provider configuration is active.
	Active *bool `json:"active,omitempty"`
	// List of event types to be exported to the APM provider.
	EventTypes []EventType `json:"eventTypes"`
	// Destination for your OpenTelemetry Protocol (OTLP) provider.
	Endpoint string `json:"endpoint"`
	// Headers configured for connecting to your OTLP provider.
	Headers []OtlpProviderHeaderDraft `json:"headers"`
	// Authentication details used to connect to an OTLP provider.
	Auth OtlpAuthDraft `json:"auth,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OtlpProviderDraft) UnmarshalJSON(data []byte) error {
	type Alias OtlpProviderDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Auth != nil {
		var err error
		obj.Auth, err = mapDiscriminatorOtlpAuthDraft(obj.Auth)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OtlpProviderDraft) MarshalJSON() ([]byte, error) {
	type Alias OtlpProviderDraft
	data, err := json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "Otlp", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["headers"] == nil {
		delete(raw, "headers")
	}

	return json.Marshal(raw)

}

/**
*	Authentication details used to connect to an OTLP provider.
*
 */
type OtlpAuthDraft interface{}

func mapDiscriminatorOtlpAuthDraft(input interface{}) (OtlpAuthDraft, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "OAuth2":
		obj := OtlpAuthOAuth2Draft{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Sets up OAuth 2.0 client credentials authentication for your OTLP provider.
*
 */
type OtlpAuthOAuth2Draft struct {
	// Client ID used to authenticate with your OTLP provider using the OAuth 2.0 client credentials flow.
	ClientId string `json:"clientId"`
	// Client secret used to authenticate with your OTLP provider using the OAuth 2.0 client credentials flow.
	ClientSecret string `json:"clientSecret"`
	// URL of the OAuth 2.0 token endpoint used to obtain an access token for your OTLP provider.
	TokenUrl string `json:"tokenUrl"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OtlpAuthOAuth2Draft) MarshalJSON() ([]byte, error) {
	type Alias OtlpAuthOAuth2Draft
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OAuth2", Alias: (*Alias)(&obj)})
}

type OtlpHttpProviderDraft struct {
	// Whether this provider configuration is active.
	Active *bool `json:"active,omitempty"`
	// List of event types to be exported to the APM provider.
	EventTypes []EventType `json:"eventTypes"`
	// Destination for your OpenTelemetry Protocol (OTLP) provider.
	Endpoint string `json:"endpoint"`
	// Headers configured for connecting to your OTLP provider.
	Headers []OtlpProviderHeaderDraft `json:"headers"`
	// Authentication details used to connect to your OTLP provider.
	Auth OtlpAuthDraft `json:"auth,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OtlpHttpProviderDraft) UnmarshalJSON(data []byte) error {
	type Alias OtlpHttpProviderDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Auth != nil {
		var err error
		obj.Auth, err = mapDiscriminatorOtlpAuthDraft(obj.Auth)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OtlpHttpProviderDraft) MarshalJSON() ([]byte, error) {
	type Alias OtlpHttpProviderDraft
	data, err := json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OtlpHttp", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["headers"] == nil {
		delete(raw, "headers")
	}

	return json.Marshal(raw)

}

type OtlpProviderHeaderDraft struct {
	// Name of the header.
	Name string `json:"name"`
	// Value of the header.
	Value string `json:"value"`
}

/**
*	Region where your Platform Insights configuration is being run.
*
 */
type ServiceRegion string

const (
	ServiceRegionLocalDev              ServiceRegion = "local-dev"
	ServiceRegionCtpGcpProductionUsV1  ServiceRegion = "ctp-gcp-production-us-v1"
	ServiceRegionCtpGcpProductionEuV1  ServiceRegion = "ctp-gcp-production-eu-v1"
	ServiceRegionCtpProductionAu       ServiceRegion = "ctp-production-au"
	ServiceRegionCtpAwsProductionFra   ServiceRegion = "ctp-aws-production-fra"
	ServiceRegionCtpAwsProductionOhio  ServiceRegion = "ctp-aws-production-ohio"
	ServiceRegionCtpGcpPerformanceTest ServiceRegion = "ctp-gcp-performance-test"
	ServiceRegionCtpStaging            ServiceRegion = "ctp-staging"
)

/**
*	Platform Insights configuration of your Project.
*
 */
type ProjectConfiguration struct {
	// Whether the Platform Insights configuration is active.
	Active bool `json:"active"`
	// User-defined unique identifier of the Project.
	ProjectKey string `json:"projectKey"`
	// List of supported APM providers.
	Providers []Provider `json:"providers"`
	// Date and time (UTC) on which the Platform Insights configuration was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// User IDs and references who last modified the Platform Insights configuration.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// List of Attributes to add to a metric.
	AdditionalAttributes []Attribute `json:"additionalAttributes"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProjectConfiguration) UnmarshalJSON(data []byte) error {
	type Alias ProjectConfiguration
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Providers {
		var err error
		obj.Providers[i], err = mapDiscriminatorProvider(obj.Providers[i])
		if err != nil {
			return err
		}
	}
	for i := range obj.AdditionalAttributes {
		var err error
		obj.AdditionalAttributes[i], err = mapDiscriminatorAttribute(obj.AdditionalAttributes[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProjectConfiguration) MarshalJSON() ([]byte, error) {
	type Alias ProjectConfiguration
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

	if raw["additionalAttributes"] == nil {
		delete(raw, "additionalAttributes")
	}

	return json.Marshal(raw)

}

type LastModifiedBy struct {
	// `id` of the [API Client](ctp:api:type:ApiClient) that created the resource.
	ClientId *string `json:"clientId,omitempty"`
}

type InternalProjectConfiguration struct {
	Active     bool   `json:"active"`
	ProjectKey string `json:"projectKey"`
	// Region where your Platform Insights configuration is being run.
	ServiceRegion ServiceRegion   `json:"serviceRegion"`
	Providers     []ProviderDraft `json:"providers"`
	// List of Attributes to add to a metric.
	AdditionalAttributes []Attribute `json:"additionalAttributes"`
	// Date and time (UTC) on which the Platform Insights configuration was last updated.
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`
	// User IDs and references who last modified the Platform Insights configuration.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *InternalProjectConfiguration) UnmarshalJSON(data []byte) error {
	type Alias InternalProjectConfiguration
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Providers {
		var err error
		obj.Providers[i], err = mapDiscriminatorProviderDraft(obj.Providers[i])
		if err != nil {
			return err
		}
	}
	for i := range obj.AdditionalAttributes {
		var err error
		obj.AdditionalAttributes[i], err = mapDiscriminatorAttribute(obj.AdditionalAttributes[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InternalProjectConfiguration) MarshalJSON() ([]byte, error) {
	type Alias InternalProjectConfiguration
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

	if raw["additionalAttributes"] == nil {
		delete(raw, "additionalAttributes")
	}

	return json.Marshal(raw)

}

type Provider interface{}

func mapDiscriminatorProvider(input interface{}) (Provider, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "Otlp":
		obj := OtlpProvider{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Auth != nil {
			var err error
			obj.Auth, err = mapDiscriminatorOtlpAuth(obj.Auth)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "OtlpHttp":
		obj := OtlpHttpProvider{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Auth != nil {
			var err error
			obj.Auth, err = mapDiscriminatorOtlpAuth(obj.Auth)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "NewRelic":
		obj := NewRelicProvider{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "DynatraceSaaS":
		obj := DynatraceSaaSProvider{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "DynatraceActiveGate":
		obj := DynatraceActiveGateProvider{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "Datadog":
		obj := DatadogProvider{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type OtlpProvider struct {
	// Whether this provider configuration is active.
	Active bool `json:"active"`
	// Whether it's possible to send metrics to the APM provider.
	Valid *bool `json:"valid,omitempty"`
	// List of event types to be exported to a supported APM provider.
	EventTypes []EventType `json:"eventTypes"`
	// Destination for your OpenTelemetry Protocol (OTLP) provider.
	Endpoint string `json:"endpoint"`
	// Headers configured for connecting to your OTLP provider.
	Headers []OtlpProviderHeader `json:"headers"`
	// Authentication details used to connect to your OTLP provider.
	Auth OtlpAuth `json:"auth,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OtlpProvider) UnmarshalJSON(data []byte) error {
	type Alias OtlpProvider
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Auth != nil {
		var err error
		obj.Auth, err = mapDiscriminatorOtlpAuth(obj.Auth)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OtlpProvider) MarshalJSON() ([]byte, error) {
	type Alias OtlpProvider
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "Otlp", Alias: (*Alias)(&obj)})
}

type OtlpHttpProvider struct {
	// Whether this provider configuration is active.
	Active bool `json:"active"`
	// Whether it's possible to send metrics to the APM provider.
	Valid *bool `json:"valid,omitempty"`
	// List of event types to be exported to a supported APM provider.
	EventTypes []EventType `json:"eventTypes"`
	// Destination for your OpenTelemetry Protocol (OTLP) provider.
	Endpoint string `json:"endpoint"`
	// Headers configured for connecting to your OTLP provider.
	Headers []OtlpProviderHeader `json:"headers"`
	// Authentication details used to connect to your OTLP provider.
	Auth OtlpAuth `json:"auth,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OtlpHttpProvider) UnmarshalJSON(data []byte) error {
	type Alias OtlpHttpProvider
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Auth != nil {
		var err error
		obj.Auth, err = mapDiscriminatorOtlpAuth(obj.Auth)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OtlpHttpProvider) MarshalJSON() ([]byte, error) {
	type Alias OtlpHttpProvider
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OtlpHttp", Alias: (*Alias)(&obj)})
}

/**
*	List of regions where NewRelic is hosted.
*
 */
type NewRelicRegion string

const (
	NewRelicRegionEu NewRelicRegion = "eu"
	NewRelicRegionUs NewRelicRegion = "us"
)

type NewRelicProvider struct {
	// Whether this provider configuration is active.
	Active bool `json:"active"`
	// Whether it's possible to send metrics to the APM provider.
	Valid *bool `json:"valid,omitempty"`
	// List of event types to be exported to a supported APM provider.
	EventTypes []EventType `json:"eventTypes"`
	// Region where your NewRelic account is hosted. For example, EU.
	Region NewRelicRegion `json:"region"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj NewRelicProvider) MarshalJSON() ([]byte, error) {
	type Alias NewRelicProvider
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "NewRelic", Alias: (*Alias)(&obj)})
}

type DynatraceSaaSProvider struct {
	// Whether this provider configuration is active.
	Active bool `json:"active"`
	// Whether it's possible to send metrics to the APM provider.
	Valid *bool `json:"valid,omitempty"`
	// List of event types to be exported to a supported APM provider.
	EventTypes []EventType `json:"eventTypes"`
	// Unique user-defined environment identifier of your DynatraceSaaS account.
	EnvironmentId string `json:"environmentId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DynatraceSaaSProvider) MarshalJSON() ([]byte, error) {
	type Alias DynatraceSaaSProvider
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "DynatraceSaaS", Alias: (*Alias)(&obj)})
}

type DynatraceActiveGateProvider struct {
	// Whether this provider configuration is active.
	Active bool `json:"active"`
	// Whether it's possible to send metrics to the APM provider.
	Valid *bool `json:"valid,omitempty"`
	// List of event types to be exported to a supported APM provider.
	EventTypes []EventType `json:"eventTypes"`
	// Your ActiveGate domain.
	ActiveGateDomain string `json:"activeGateDomain"`
	// Your ActiveGate domain port.
	ActiveGatePort *float64 `json:"activeGatePort,omitempty"`
	// Unique user-defined environment identifier of your DynatraceSaaS account.
	EnvironmentId string `json:"environmentId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DynatraceActiveGateProvider) MarshalJSON() ([]byte, error) {
	type Alias DynatraceActiveGateProvider
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "DynatraceActiveGate", Alias: (*Alias)(&obj)})
}

type DatadogSite string

const (
	DatadogSiteUS1    DatadogSite = "US1"
	DatadogSiteUS3    DatadogSite = "US3"
	DatadogSiteUS5    DatadogSite = "US5"
	DatadogSiteEU1    DatadogSite = "EU1"
	DatadogSiteUS1FED DatadogSite = "US1-FED"
	DatadogSiteAP1    DatadogSite = "AP1"
)

type DatadogProvider struct {
	// Whether this provider configuration is active.
	Active bool `json:"active"`
	// Whether it's possible to send metrics to the APM provider.
	Valid *bool `json:"valid,omitempty"`
	// List of event types to be exported to a supported APM provider.
	EventTypes []EventType `json:"eventTypes"`
	// Website where your Datadog account is hosted.
	Site DatadogSite `json:"site"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DatadogProvider) MarshalJSON() ([]byte, error) {
	type Alias DatadogProvider
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "Datadog", Alias: (*Alias)(&obj)})
}

type AttributeTypeId string

const (
	AttributeTypeIdString AttributeTypeId = "String"
)

type Attribute interface{}

func mapDiscriminatorAttribute(input interface{}) (Attribute, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "String":
		obj := StringAttribute{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type AttributeDraft interface{}

func mapDiscriminatorAttributeDraft(input interface{}) (AttributeDraft, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "String":
		obj := StringAttributeDraft{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type StringAttributeDraft struct {
	// List of event types to be exported to a supported APM provider.
	EventType EventType `json:"eventType"`
	Name      string    `json:"name"`
	Value     string    `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StringAttributeDraft) MarshalJSON() ([]byte, error) {
	type Alias StringAttributeDraft
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "String", Alias: (*Alias)(&obj)})
}

type StringAttribute struct {
	// List of event types to be exported to a supported APM provider.
	EventType EventType `json:"eventType"`
	Name      string    `json:"name"`
	Value     string    `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StringAttribute) MarshalJSON() ([]byte, error) {
	type Alias StringAttribute
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "String", Alias: (*Alias)(&obj)})
}
