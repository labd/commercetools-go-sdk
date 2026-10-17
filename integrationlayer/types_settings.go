package integrationlayer

// Generated file, please do not change!!!

/**
*	The language, currency and country a new session starts with, and the Project's label.
 */
type ProjectSettings struct {
	Label    string `json:"label"`
	Language string `json:"language"`
	Currency string `json:"currency"`
	Country  string `json:"country"`
}

/**
*	The project section as read. A Project created before a setting existed may lack it.
 */
type ProjectSection struct {
	Label    *string         `json:"label,omitempty"`
	Language *string         `json:"language,omitempty"`
	Currency *string         `json:"currency,omitempty"`
	Country  *string         `json:"country,omitempty"`
	Checkout *CheckoutConfig `json:"checkout,omitempty"`
}

type PlatformSettings struct {
	RegionId string `json:"regionId"`
}

/**
*	The values the commercetools Project offers.
 */
type ProjectPresentment struct {
	Languages  []string `json:"languages"`
	Countries  []string `json:"countries"`
	Currencies []string `json:"currencies"`
}

type ProjectSettingsView struct {
	// The project section as read. A Project created before a setting existed may lack it.
	Project  ProjectSection   `json:"project"`
	Platform PlatformSettings `json:"platform"`
	// The values the commercetools Project offers.
	Presentment *ProjectPresentment `json:"presentment,omitempty"`
	Version     int                 `json:"version"`
}

type ProjectSettingsReplaceResult struct {
	// The language, currency and country a new session starts with, and the Project's label.
	Project ProjectSettings `json:"project"`
	Version int             `json:"version"`
}
