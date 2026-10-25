package integrationlayer

// Generated file, please do not change!!!

type SourceCoordinateSpace string

const (
	SourceCoordinateSpaceCustomField SourceCoordinateSpace = "customField"
	SourceCoordinateSpaceAttribute   SourceCoordinateSpace = "attribute"
)

type SourceCoordinate struct {
	Space     SourceCoordinateSpace `json:"space"`
	TypeKey   string                `json:"typeKey"`
	FieldName string                `json:"fieldName"`
}

/**
*	A field the Project has or a stored rule names. A rule exists for it when it is active or has a reason.
 */
type FieldVisibilityRow struct {
	Coordinate SourceCoordinate `json:"coordinate"`
	FieldType  *string          `json:"fieldType,omitempty"`
	Active     bool             `json:"active"`
	Reason     *string          `json:"reason,omitempty"`
}

type FieldVisibilityView struct {
	Fields  []FieldVisibilityRow `json:"fields"`
	Version int                  `json:"version"`
}

type FieldVisibilityReplaceResult struct {
	HiddenFields []SourceCoordinate `json:"hiddenFields"`
	Version      int                `json:"version"`
	// The rules are stored, but republishing the schema failed.
	PublishPending bool `json:"publishPending"`
}
