package connect

// Generated file, please do not change!!!

import (
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
type ConfigurationKeyStandard struct {
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
type ConfigurationKeySecured struct {
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
	StandardConfiguration []ConfigurationKeyStandard `json:"standardConfiguration"`
	// Contains the name and description of secret environment variables.
	SecuredConfiguration []ConfigurationKeySecured `json:"securedConfiguration"`
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
	// Contains values of secret keys. Cannot be accessed after being set.
	SecuredConfiguration []ConfigurationValue `json:"securedConfiguration"`
	// URL generated after deployment of service applications.
	Url *string `json:"url,omitempty"`
	// Pub/Sub Topic generated after deployment of event applications.
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
*	Reference to a Connector. Either `id` or `key` is required.
 */
type ConnectorReference struct {
	// Unique identifier of the referenced Connector.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced Connector.
	Key *string `json:"key,omitempty"`
	// If `true`, the previewable [ConnectorStaged](ctp:connect:type:ConnectorStaged) will be referenced instead of the published Connector. The `isPreviewable` field of the [ConnectorStaged](ctp:connect:type:ConnectorStaged) must be `true` to reference a previewable ConnectorStaged.
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
	// If `true`, the Connector is certified.
	Certified bool `json:"certified"`
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
	// Contains informations, errors and warnings about the Connector deployment.
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
*	The host Region of a commercetools Composable Commerce Project. For more information, see [Hosts](hosts-and-authorization#hosts).
*
 */
type Region string

const (
	RegionEuropeWest1Gcp         Region = "europe-west1.gcp"
	RegionUsCentral1Gcp          Region = "us-central1.gcp"
	RegionAustraliaSoutheast1Gcp Region = "australia-southeast1.gcp"
)

/**
*	The type of integration provided by a Connector.
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
	IntegrationTypeOther       IntegrationType = "other"
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
