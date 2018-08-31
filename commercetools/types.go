package commercetools

import (
	"encoding/json"

	"github.com/mitchellh/mapstructure"
)

type LocalizedString map[string]string

type Reference struct {
	TypeID string `json:"typeId"`
	ID     string `json:"id"`
}

type ResourceIdentifier struct {
	ID     string `json:"id,omitempty"`
	Key    string `json:"key,omitempty"`
	TypeID string `json:"typeId,omitempty"`
}

type UpdateAction interface {
	MarshalJSON() ([]byte, error)
}

type UpdateActions []UpdateAction

type Money struct {
	Type           string `type:"type"`
	CurrencyCode   string `type:"currencyCode"`
	CentAmount     int    `type:"centAmount"`
	FractionDigits int    `type:"fractionDigits"`
}

type HighPrecisionMoney struct {
	Money
	PreciseAmount int `type:"preciseAmount"`
}

type Asset struct {
	ID          string          `type:"id"`
	Key         string          `type:"key"`
	Sources     []AssetSource   `type:"sources"`
	Name        LocalizedString `json:"name"`
	Description LocalizedString `json:"description,omitempty"`
	Tags        []string        `json:"tags,omitempty"`
	// Custom
}

type AssetDraft struct {
	Key         string          `type:"key"`
	Sources     []AssetSource   `type:"sources"`
	Name        LocalizedString `json:"name"`
	Description LocalizedString `json:"description,omitempty"`
	Tags        []string        `json:"tags,omitempty"`
	// Custom
}

type AssetSource struct {
	URI        string          `type:"uri"`
	Key        string          `type:"key"`
	Dimensions AssetDimensions `type:"dimensions,omitempty"`
}

type AssetDimensions struct {
	H int `json:"h"`
	W int `json:"w"`
}

// EnumValue stores enums in field types.
type EnumValue struct {
	// The key of the value used as a programmatic identifier.
	Key string `json:"key"`
	// A descriptive label of the value.
	Label string `json:"label"`
}

// LocalizedEnumValue stores localized enums in field types.
type LocalizedEnumValue struct {
	// The key of the value used as a programmatic identifier.
	Key string `json:"key"`
	// A descriptive, localized label of the value.
	Label LocalizedString `json:"label"`
}

// TextInputHint provides a visual representation type for a field.
// It is only relevant for string-based field types like StringType and LocalizedStringType.
type TextInputHint string

const (
	// SingleLineTextInputHint allows a single line hint.
	SingleLineTextInputHint TextInputHint = "SingleLine"
	// MultiLineTextInputHint allows a multi line hint.
	MultiLineTextInputHint TextInputHint = "MultiLine"
)

// Address is a JSON object representation of a postal address.
type Address struct {
	ID                   string `json:"id,omitempty"`
	Key                  string `json:"key,omitempty"`
	Title                string `json:"title,omitempty"`
	Salutation           string `json:"salutation,omitempty"`
	FirstName            string `json:"firstName,omitempty"`
	LastName             string `json:"lastName,omitempty"`
	StreetName           string `json:"streetName,omitempty"`
	StreetNumber         string `json:"streetNumber,omitempty"`
	AdditionalStreetInfo string `json:"additionalStreetInfo,omitempty"`
	PostalCode           string `json:"postalCode,omitempty"`
	City                 string `json:"city,omitempty"`
	Region               string `json:"region,omitempty"`
	State                string `json:"state,omitempty"`
	// A two-digit country code as per ISO 3166-1 alpha-2
	Country               string `json:"country,omitempty"`
	Company               string `json:"company,omitempty"`
	Department            string `json:"department,omitempty"`
	Building              string `json:"building,omitempty"`
	Apartment             string `json:"apartment,omitempty"`
	POBox                 string `json:"pOBox,omitempty"`
	Phone                 string `json:"phone,omitempty"`
	Mobile                string `json:"mobile,omitempty"`
	Email                 string `json:"email,omitempty"`
	Fax                   string `json:"fax,omitempty"`
	AdditionalAddressInfo string `json:"additionalAddressInfo,omitempty"`
	ExternalID            string `json:"externalId,omitempty"`
}

// Point indicates a single position.
type Point struct {
	Coordinates []float64 `json:"coordinates"`
}

// MarshalJSON override to add the type value.
func (t Point) MarshalJSON() ([]byte, error) {
	type Alias Point

	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "Point",
		Alias: (*Alias)(&t),
	})
}

// GeoJSONGeometry represents a Geometry Object as defined in the GeoJSON standard.
type GeoJSONGeometry interface{}

func GeoJSONGeometryMapping(input GeoJSONGeometry) GeoJSONGeometry {
	FieldType := input.(map[string]interface{})["type"]
	switch FieldType {
	case "Point":
		newType := Point{}
		mapstructure.Decode(input, &newType)
		return newType
	}
	return nil
}

// PagedQueryResult for query responses of requests supporting paging via limit
// and offset.
type PagedQueryResult struct {
	// The offset supplied by the client or the server default. It is the number
	// of elements skipped, not a page number.
	Offset int `json:"offset"`

	Limit int `json:"limit"`

	// The actual number of results returned in results.
	Count int `json:"count"`

	// The total number of results matching the query. This field is returned by
	// default. For improved performance, calculating this field can be
	// deactivated by using the query parameter withTotal=false.
	Total int `json:"total,omitempty"`

	Results []interface{} `json:"results"`

	Facets map[string]FacetResult `json:"facets,omitempty"`
}

type FacetResult struct {
	Type  string `json:"type"`
	Count int    `json:"count"`
}
