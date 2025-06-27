package connect

// Generated file, please do not change!!!

import (
	"encoding/json"
	"time"
)

/**
*	A connector internal build artifact (ex: docker image).
*
 */
type Artifact struct {
	ApplicationName string `json:"applicationName"`
	ArtifactPath    string `json:"artifactPath"`
}

type CertificationInfo struct {
	// Comments made during the certification process.
	Comments []CertificationInfoComment `json:"comments"`
}

/**
*	Contains metadata and body text of a comment made during the certification process.
*
 */
type CertificationInfoComment struct {
	// The commercetools Composable Commerce Project key associated with the person making the comment.
	UserId string `json:"userId"`
	// Date and time (UTC) the comment was added.
	Datetime time.Time `json:"datetime"`
	// The body text of the comment.
	Comment string `json:"comment"`
}

/**
*	The name, description, and default value of a standard environment variable.
*
 */
type StandardConfigurationKey struct {
	// Name of the environment variable.
	Key string `json:"key"`
	// Description of the environment variable.
	Description string `json:"description"`
	// Default value of the environment variable.
	Default *string `json:"default,omitempty"`
	// Indicates if the environment variable is required.
	Required *bool `json:"required,omitempty"`
}

/**
*	The name and description of a secret environment variable.
*
 */
type SecuredConfigurationKey struct {
	// Name of the environment variable.
	Key string `json:"key"`
	// Description of the environment variable.
	Description string `json:"description"`
	// Indicates if the environment variable is required.
	Required *bool `json:"required,omitempty"`
}

type ConfigurationValue struct {
	// Name of the environment variable.
	Key string `json:"key"`
	// Value of the environment variable.
	Value string `json:"value"`
}

/**
*	The configuration of a Connect application. These values are automatically obtained from the connect.yaml file in the GitHub repository.
*
 */
type ConnectorConfigurationApplication struct {
	// Name of the Connect application.
	ApplicationName string `json:"applicationName"`
	// The Connect application type.
	ApplicationType string `json:"applicationType"`
	// Contains the name, description, and default values of standard environment variables.
	StandardConfiguration []StandardConfigurationKey `json:"standardConfiguration"`
	// Contains the name and description of secret environment variables.
	SecuredConfiguration []SecuredConfigurationKey `json:"securedConfiguration"`
}

/**
*	Global configuration applied to all applications in the deployment.
 */
type ConnectorGlobalConfiguration struct {
	// Contains the name and description of standard environment variables.
	StandardConfiguration []StandardConfigurationKey `json:"standardConfiguration"`
	// Contains the name and description of secret environment variables.
	SecuredConfiguration []SecuredConfigurationKey `json:"securedConfiguration"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConnectorGlobalConfiguration) MarshalJSON() ([]byte, error) {
	type Alias ConnectorGlobalConfiguration
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

	if raw["standardConfiguration"] == nil {
		delete(raw, "standardConfiguration")
	}

	if raw["securedConfiguration"] == nil {
		delete(raw, "securedConfiguration")
	}

	return json.Marshal(raw)

}

/**
*	Configuration for generating API Client credentials.
*
 */
type ConnectorApiClient struct {
	// List of [scopes](/hosts-and-authorization#authorization) assigned to the API Client.
	Scopes []string `json:"scopes"`
}

/**
*	API Client credentials used for deployment.
*
 */
type DeploymentApiClient struct {
	// Client ID of the API Client.
	Name string `json:"name"`
	// List of [scopes](/hosts-and-authorization#authorization) assigned to the API Client.
	Scopes []string `json:"scopes"`
}

/**
*	Describes the configuration set of a given application.
*
 */
type DeploymentConfigurationApplication struct {
	// Name of the application. Should match the value of `name` within the connect.yaml file of the Connect application.
	ApplicationName string `json:"applicationName"`
	// Contains values of keys that are saved in plain text. Can be accessed after being set.
	StandardConfiguration []ConfigurationValue `json:"standardConfiguration"`
	// Contains values of secret keys. Cannot be accessed after being set.
	SecuredConfiguration []ConfigurationValue `json:"securedConfiguration"`
}

/**
*	Describes an application deployment of the Connector.
*
 */
type ApplicationDeployment struct {
	// Unique identifier of the application deployment.
	ID string `json:"id"`
	// Name of the application. Should match the value of `name` within the connect.yaml file of the Connector.
	ApplicationName string `json:"applicationName"`
	// Contains values of keys that are saved in plain text. Can be accessed after being set.
	StandardConfiguration []ConfigurationValue `json:"standardConfiguration"`
	// Contains values of secret keys. After being set, these values are encrypted and cannot be accessed.
	SecuredConfiguration []ConfigurationValue `json:"securedConfiguration"`
	// URL generated after deployment of service applications.
	Url *string `json:"url,omitempty"`
	// Google Cloud Pub/Sub Topic generated after deployment of event applications.
	Topic *string `json:"topic,omitempty"`
	// Cron schedule for job applications.
	Schedule *string `json:"schedule,omitempty"`
}

/**
*	The type of message being reported.
*
 */
type ConnectorReportEntryType string

const (
	ConnectorReportEntryTypeInformation ConnectorReportEntryType = "Information"
	ConnectorReportEntryTypeWarning     ConnectorReportEntryType = "Warning"
	ConnectorReportEntryTypeError       ConnectorReportEntryType = "Error"
)

/**
*	Describes an information, error, or warning notice.
 */
type ConnectorReportEntry struct {
	// The report entry type.
	Type ConnectorReportEntryType `json:"type"`
	// The title of the report entry.
	Title string `json:"title"`
	// The message related to the report entry.
	Message *string `json:"message,omitempty"`
	// When the report entry was created.
	CreatedAt time.Time `json:"createdAt"`
}

/**
*	Contains report entries for publish/preview requests.
 */
type ConnectorReport struct {
	// Contains information, error, and warning notices.
	Entries []ConnectorReportEntry `json:"entries"`
}

/**
*	Indicates the current status of the ConnectorStaged.
*
 */
type ConnectorStatus string

const (
	ConnectorStatusDraft                 ConnectorStatus = "Draft"
	ConnectorStatusProcessing            ConnectorStatus = "Processing"
	ConnectorStatusReadyForCertification ConnectorStatus = "ReadyForCertification"
	ConnectorStatusInCertification       ConnectorStatus = "InCertification"
	ConnectorStatusPublished             ConnectorStatus = "Published"
	ConnectorStatusFailed                ConnectorStatus = "Failed"
)

/**
*	Reference to a [Connector](ctp:connect:type:Connector) or [ConnectorStaged](ctp:connect:type:ConnectorStaged). Either `id` or `key` is required.
 */
type ConnectorReference struct {
	// Unique identifier of the referenced Connector.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced Connector.
	Key *string `json:"key,omitempty"`
	// If `true`, a previewable [ConnectorStaged](ctp:connect:type:ConnectorStaged) must be referenced in `id` or `key` instead of the published Connector. The `isPreviewable` field of the [ConnectorStaged](ctp:connect:type:ConnectorStaged) must be `true` to reference a previewable ConnectorStaged.
	Staged *bool `json:"staged,omitempty"`
	// Version of the referenced Connector.
	Version *int `json:"version,omitempty"`
}

/**
*	Details of the individual or organization who developed the Connector.
*
 */
type Creator struct {
	// Name of the person who owns the Connector.
	Name *string `json:"name,omitempty"`
	// Title of the person who owns the Connector.
	Title *string `json:"title,omitempty"`
	// Contact email address of the creator.
	Email string `json:"email"`
	// Name of the company that owns the Connector.
	Company *string `json:"company,omitempty"`
	// URL to a logo image used to represent the creator.
	LogoUrl *string `json:"logoUrl,omitempty"`
	// Number of contributors currently working for the company.
	NoOfContributors *int `json:"noOfContributors,omitempty"`
	// URL to the support website of the Connector.
	SupportUrl *string `json:"supportUrl,omitempty"`
}

/**
*	The details of the deployed [Connector](ctp:connect:type:Connector).
 */
type DeploymentConnector struct {
	// `id` of the Connector.
	ID string `json:"id"`
	// User-defined unique identifier of the Connector.
	Key *string `json:"key,omitempty"`
	// Version of the Connector.
	Version int `json:"version"`
	// Name of the Connector.
	Name string `json:"name"`
	// Description of the Connector.
	Description *string `json:"description,omitempty"`
	// Connector integration type list.
	IntegrationTypes []IntegrationType `json:"integrationTypes"`
	// Owner of the Connector.
	Creator DeploymentCreator `json:"creator"`
	// Details of the GitHub repository that contains the Connect applications.
	Repository Repository `json:"repository"`
	// Configurations needed by Connectors for hosting. Loaded as environment variables in the application.
	Configurations []ConnectorConfigurationApplication `json:"configurations"`
	// Global configuration applied to all applications in the deployment.
	GlobalConfiguration *ConnectorGlobalConfiguration `json:"globalConfiguration,omitempty"`
	// If `true`, the Connector is certified.
	Certified bool `json:"certified"`
	// If not empty, Connectors can only be deployed in these Regions. If empty, Connectors can be deployed in any [supported Region](hosts-and-authorization#hosts).
	SupportedRegions []Region `json:"supportedRegions"`
	// URL to the documentation of the Connector.
	DocumentationUrl *string `json:"documentationUrl,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DeploymentConnector) MarshalJSON() ([]byte, error) {
	type Alias DeploymentConnector
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

	if raw["supportedRegions"] == nil {
		delete(raw, "supportedRegions")
	}

	return json.Marshal(raw)

}

/**
*	Global configuration applied to all applications in the deployment.
 */
type DeploymentGlobalConfiguration struct {
	// List of standard environment variables.
	StandardConfiguration []ConfigurationValue `json:"standardConfiguration"`
	// List of secured environment variables.
	SecuredConfiguration []ConfigurationValue `json:"securedConfiguration"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DeploymentGlobalConfiguration) MarshalJSON() ([]byte, error) {
	type Alias DeploymentGlobalConfiguration
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

	if raw["standardConfiguration"] == nil {
		delete(raw, "standardConfiguration")
	}

	if raw["securedConfiguration"] == nil {
		delete(raw, "securedConfiguration")
	}

	return json.Marshal(raw)

}

/**
*	The details of the deployed [DeploymentConnector](ctp:connect:type:DeploymentConnector) creator.
 */
type DeploymentCreator struct {
	// Name of the person who owns the Connector.
	Name *string `json:"name,omitempty"`
	// Title of the person who owns the Connector.
	Title *string `json:"title,omitempty"`
	// Contact email address of the creator.
	Email string `json:"email"`
	// Name of the company that owns the Connector.
	Company *string `json:"company,omitempty"`
	// Number of contributors currently working for the company.
	NoOfContributors *int `json:"noOfContributors,omitempty"`
	// URL to an image used as the company logo.
	LogoUrl *string `json:"logoUrl,omitempty"`
	// Url to support website of the company.
	SupportUrl *string `json:"supportUrl,omitempty"`
}

/**
*	The type of message being reported.
*
 */
type DeploymentReportEntryType string

const (
	DeploymentReportEntryTypeInformation DeploymentReportEntryType = "Information"
	DeploymentReportEntryTypeWarning     DeploymentReportEntryType = "Warning"
	DeploymentReportEntryTypeError       DeploymentReportEntryType = "Error"
)

/**
*	Describes an information, error, or warning in the deployment report.
 */
type DeploymentReportEntry struct {
	// The report entry type.
	Type DeploymentReportEntryType `json:"type"`
	// The title of the report entry.
	Title string `json:"title"`
	// The message related to the report entry.
	Message *string `json:"message,omitempty"`
	// The name of the Connect application related to the entry.
	Application *string `json:"application,omitempty"`
	// When the report entry was created.
	CreatedAt time.Time `json:"createdAt"`
}

/**
*	Describes a report of the Connector deployment process.
 */
type DeploymentReport struct {
	// Contains information, errors and warnings about the Connector deployment.
	Entries []DeploymentReportEntry `json:"entries"`
}

type DeploymentDetailsBuild struct {
	// The build execution id of the deployed application.
	ID string `json:"id"`
	// The build report of the deployed Connector.
	Report *DeploymentReport `json:"report,omitempty"`
}

/**
*	Additional details about the deployed Connector.
 */
type DeploymentDetails struct {
	// The build details of deployed applications.
	Build DeploymentDetailsBuild `json:"build"`
}

/**
*	Indicates the current status of the Deployment.
*
 */
type DeploymentStatus string

const (
	DeploymentStatusDraft          DeploymentStatus = "Draft"
	DeploymentStatusQueued         DeploymentStatus = "Queued"
	DeploymentStatusDeploying      DeploymentStatus = "Deploying"
	DeploymentStatusDeployed       DeploymentStatus = "Deployed"
	DeploymentStatusFailed         DeploymentStatus = "Failed"
	DeploymentStatusUndeploying    DeploymentStatus = "Undeploying"
	DeploymentStatusUndeployFailed DeploymentStatus = "UndeployFailed"
)

/**
*	The severity of the log message.
 */
type DeploymentLogSeverity string

const (
	DeploymentLogSeverityDEFAULT   DeploymentLogSeverity = "DEFAULT"
	DeploymentLogSeverityDEBUG     DeploymentLogSeverity = "DEBUG"
	DeploymentLogSeverityINFO      DeploymentLogSeverity = "INFO"
	DeploymentLogSeverityNOTICE    DeploymentLogSeverity = "NOTICE"
	DeploymentLogSeverityWARNING   DeploymentLogSeverity = "WARNING"
	DeploymentLogSeverityERROR     DeploymentLogSeverity = "ERROR"
	DeploymentLogSeverityCRITICAL  DeploymentLogSeverity = "CRITICAL"
	DeploymentLogSeverityALERT     DeploymentLogSeverity = "ALERT"
	DeploymentLogSeverityEMERGENCY DeploymentLogSeverity = "EMERGENCY"
)

/**
*	The Region of a commercetools Composable Commerce Project or Deployment. For more information, see [Hosts](/hosts-and-authorization#hosts).
*
 */
type Region string

const (
	RegionEuropeWest1Gcp         Region = "europe-west1.gcp"
	RegionUsCentral1Gcp          Region = "us-central1.gcp"
	RegionAustraliaSoutheast1Gcp Region = "australia-southeast1.gcp"
)

/**
*	Integration type of the Connector.
*
 */
type IntegrationType string

const (
	IntegrationTypeTax         IntegrationType = "tax"
	IntegrationTypeMarketplace IntegrationType = "marketplace"
	IntegrationTypeOms         IntegrationType = "oms"
	IntegrationTypePsp         IntegrationType = "psp"
	IntegrationTypePim         IntegrationType = "pim"
	IntegrationTypePromotion   IntegrationType = "promotion"
	IntegrationTypeSearch      IntegrationType = "search"
	IntegrationTypeErp         IntegrationType = "erp"
	IntegrationTypeCrm         IntegrationType = "crm"
	IntegrationTypeEmail       IntegrationType = "email"
	IntegrationTypeAnalytics   IntegrationType = "analytics"
	IntegrationTypeShipping    IntegrationType = "shipping"
	IntegrationTypeGiftcard    IntegrationType = "giftcard"
	IntegrationTypeOther       IntegrationType = "other"
)

/**
*	Deployment type of the Connector.
*
 */
type DeploymentType string

const (
	DeploymentTypePreview    DeploymentType = "preview"
	DeploymentTypeSandbox    DeploymentType = "sandbox"
	DeploymentTypeProduction DeploymentType = "production"
)

/**
*	Details of the GitHub repository that contains the Connect applications.
*
 */
type Repository struct {
	// HTTPS or SSH GitHub URL of the GitHub repository. Private repositories must use an SSH URL.
	Url string `json:"url"`
	// Git tag of the release to use.
	Tag string `json:"tag"`
}

/**
*	The previewable status of the ConnectorStaged.
*
 */
type IsPreviewable string

const (
	IsPreviewableTrue    IsPreviewable = "true"
	IsPreviewableFalse   IsPreviewable = "false"
	IsPreviewablePending IsPreviewable = "pending"
	IsPreviewableNone    IsPreviewable = "none"
)
