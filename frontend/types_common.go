package frontend

// Generated file, please do not change!!!

import (
	"io"
)

type DataSourceConfiguration struct {
	DataSourceId  string      `json:"dataSourceId"`
	Type          string      `json:"type"`
	Name          string      `json:"name"`
	Configuration interface{} `json:"configuration"`
}

type DataSourceResponse struct {
	DataSourcePayload interface{} `json:"dataSourcePayload"`
	PreviewPayload    interface{} `json:"previewPayload"`
}

/**
*	Object containing the data sources configured for the page folder and referenced in the page.
 */
type DataSources map[string]interface{}
type Environment string

const (
	EnvironmentProduction  Environment = "production"
	EnvironmentDevelopment Environment = "development"
)

type Error struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func (obj Error) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown Error: failed to parse error response"
}

type Footer struct {
	LayoutElements []LayoutElement `json:"layoutElements"`
}

type Head struct {
	LayoutElements []LayoutElement `json:"layoutElements"`
}

type LayoutElement struct {
	LayoutElementId string                     `json:"layoutElementId"`
	Configuration   LayoutElementConfiguration `json:"configuration"`
	Tastics         []Tastic                   `json:"tastics"`
}

type LayoutElementConfiguration struct {
	Size    float64 `json:"size"`
	Mobile  bool    `json:"mobile"`
	Tablet  bool    `json:"tablet"`
	Desktop bool    `json:"desktop"`
}

type Main struct {
	LayoutElements []LayoutElement `json:"layoutElements"`
}

/**
*	Page represents the structure and data of the published page displayed through commercetools Frontend.
*
 */
type Page struct {
	PageId   string   `json:"pageId"`
	Sections Sections `json:"sections"`
	State    string   `json:"state"`
}

type PageDataResponse struct {
	PageFolder PageFolder `json:"pageFolder"`
	// Page represents the structure and data of the published page displayed through commercetools Frontend.
	Page Page     `json:"page"`
	Data ViewData `json:"data"`
}

type PageFolder struct {
	PageFolderId                string                    `json:"pageFolderId"`
	IsDynamic                   bool                      `json:"isDynamic"`
	PageFolderType              string                    `json:"pageFolderType"`
	Configuration               interface{}               `json:"configuration"`
	DataSourceConfigurations    []DataSourceConfiguration `json:"dataSourceConfigurations"`
	Name                        *string                   `json:"name,omitempty"`
	AncestorIdsMaterializedPath string                    `json:"ancestorIdsMaterializedPath"`
	Depth                       float64                   `json:"depth"`
	Sort                        float64                   `json:"sort"`
}

type PagePreviewContext struct {
	CustomerName string `json:"customerName"`
}

type PagePreviewDataResponse struct {
	PageFolder PageFolder `json:"pageFolder"`
	// Page represents the structure and data of the published page displayed through commercetools Frontend.
	Page           Page               `json:"page"`
	Data           ViewData           `json:"data"`
	PreviewId      string             `json:"previewId"`
	PreviewContext PagePreviewContext `json:"previewContext"`
}

type PathConfiguration struct {
	Path             string   `json:"path"`
	PathTranslations []string `json:"pathTranslations"`
}

type ProjectContext struct {
	// Indicates whether the project is intended as production or development environment.
	Environment Environment `json:"environment"`
	// Locales that can be used in the project.
	Locales []string `json:"locales"`
	// Locale used by default in the project.
	DefaultLocale string `json:"defaultLocale"`
}

type RedirectResponse struct {
	StatusCode float64 `json:"statusCode"`
	Reason     string  `json:"reason"`
	TargetType string  `json:"targetType"`
	Target     string  `json:"target"`
}

type Sections struct {
	Head   Head   `json:"head"`
	Main   Main   `json:"main"`
	Footer Footer `json:"footer"`
}

type Tastic struct {
	TasticId      string              `json:"tasticId"`
	TasticType    string              `json:"tasticType"`
	Configuration TasticConfiguration `json:"configuration"`
}

type TasticConfiguration struct {
	Desktop bool `json:"desktop"`
	Mobile  bool `json:"mobile"`
	Tablet  bool `json:"tablet"`
}

type ViewData struct {
	// Data sources configured for the page folder and referenced in the page. They are indexed by data source identifier.
	DataSources DataSources `json:"dataSources"`
}

type BuildUpload struct {
	// Metadata about the build.
	Metadata BuildMetadata `json:"metadata"`
	// Javascript bundle containing all code of the backend extensions.
	Extension io.Reader `json:"extension"`
	// Map for the build file. When provided, exceptions will be logged with readable stack traces.
	ExtensionMap *io.Reader `json:"extension_map,omitempty"`
}

type BuildMetadata struct {
	// Name of the Git branch to deploy.
	Branch string `json:"branch"`
	// Full commit ID hash of the current HEAD.
	Revision string `json:"revision"`
	// Current time represented in the format `YYYY.MM.DD.HH.mm`. This value should be unique and sortable.
	VersionNumber string `json:"versionNumber"`
	// Node.js version for the [Extension runner](/../frontend-getting-started/development-concepts#extension-runner).
	NodeJsVersion *int `json:"nodeJsVersion,omitempty"`
	// If `true`, the build is marked as successful and [displayed in the Studio](/../frontend-studio/using-deployment-in-the-studio).
	BuildSuccessful bool `json:"buildSuccessful"`
	// Date and time of the build in the [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) format. For example, `2023-11-21T08:14:31.830Z`.
	BuildTime string `json:"buildTime"`
	// Build duration in seconds.
	BuildDuration int `json:"buildDuration"`
	// CI execution logs.
	BuildLog string `json:"buildLog"`
	// If `true`, the build is automatically deployed to the staging environment, if `buildSuccessful` is `true`.
	Deploy *bool `json:"deploy,omitempty"`
}

type BuildUploadResult struct {
	// Status of the build.
	Status string `json:"status"`
	// Description of the upload result.
	Message string `json:"message"`
}
