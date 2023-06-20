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

type ConfigurationKey struct {
	// Name of the environment variable.
	Key string `json:"key"`
	// Description of the environment variable.
	Description string `json:"description"`
}

type ConfigurationValue struct {
	// Name of the environment variable.
	Key string `json:"key"`
	// Value of the environment variable.
	Value string `json:"value"`
}

/**
*	Describes the required configuration for a Connector.
*
 */
type ConnectorConfigurationApplication struct {
	// Name of the application. Should match the value of `name` within the connect.yaml file of the Connect application.
	ApplicationName string `json:"applicationName"`
	// Contains the name and description of keys saved as plain text.
	StandardConfiguration []ConfigurationKey `json:"standardConfiguration"`
	// Contains the name and description of secret keys.
	SecuredConfiguration []ConfigurationKey `json:"securedConfiguration"`
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
*	Indicates the current status of the ConnectorStaged.
*
 */
type ConnectorStatus string

const (
	ConnectorStatusDraft                 ConnectorStatus = "Draft"
	ConnectorStatusReadyForCertification ConnectorStatus = "ReadyForCertification"
	ConnectorStatusInCertification       ConnectorStatus = "InCertification"
	ConnectorStatusFeedbackRequired      ConnectorStatus = "FeedbackRequired"
	ConnectorStatusCertified             ConnectorStatus = "Certified"
)

/**
*	Reference to a Connector. Either `id` or `key` is required.
 */
type ConnectorReference struct {
	// Unique identifier of the referenced Connector.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced Connector.
	Key *string `json:"key,omitempty"`
	// If `true`, the previewable [ConnectorStaged](ctp:connect:type:ConnectorStaged) will be referenced instead of the certified Connector. The `isPreviewable` field of the [ConnectorStaged](ctp:connect:type:ConnectorStaged) must be `true` to reference a previewable ConnectorStaged.
	Staged *bool `json:"staged,omitempty"`
	// Version of the referenced Connector.
	Version int `json:"version"`
}

/**
*	Contains details of the creator of the Connector.
*	    These details are used to contact the creator during the certification process.
*
 */
type Creator struct {
	// Name of the person who owns the Connector.
	Name *string `json:"name,omitempty"`
	// Title of the person who owns the Connector.
	Title *string `json:"title,omitempty"`
	// Contact email address of the Creator.
	Email string `json:"email"`
	// Name of the company that owns the Connector.
	Company *string `json:"company,omitempty"`
	// Number of developers currently working for the company.
	NoOfDevelopers *int `json:"noOfDevelopers,omitempty"`
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
	// Owner of the Connector.
	Creator Creator `json:"creator"`
	// GitHub repository details of the Connector.
	Repository Repository `json:"repository"`
}

type DeploymentDetailsBuild struct {
	// The build execution id of the deployed application.
	ID string `json:"id"`
}

/**
*	The build and runtime details of deployed applications.
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
*	Details of the GitHub repository that contains the Connect application.
*
 */
type Repository struct {
	// HTTPS or SSH GitHub URL of the GitHub repository. Private repositories must use an SSH URL.
	Url string `json:"url"`
	// Git tag of the release to use.
	Tag string `json:"tag"`
}

/**
*	A Paged Result.
*
 */
type Paged struct {
	Limit   float64     `json:"limit"`
	Offset  float64     `json:"offset"`
	Count   float64     `json:"count"`
	Total   float64     `json:"total"`
	Results interface{} `json:"results"`
}

/**
*	A Cursor Paged Result.
*
 */
type CursorPaged struct {
	Data interface{} `json:"data"`
	Next string      `json:"next"`
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
