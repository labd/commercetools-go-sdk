package commercetools

import (
	"encoding/json"

	"github.com/mitchellh/mapstructure"
)

// LocalizedString is a JSON object where the keys are of IETF language tag, and
// the values the corresponding strings used for that language.
type LocalizedString map[string]string

// Reference is a JSON object representing a (loose) reference to another
// resource on the commercetools platform.
type Reference struct {
	// The typeId designates the type of the referenced resource.
	TypeID string `json:"typeId"`

	// The unique ID of the referenced resource.
	ID string `json:"id"`
}

// ResourceIdentifier is used to identify a resource. A reference to a resource
// can be created by providing the ID of the resource. Some resources also use
// the key as a unique identifier and a reference can be created by providing
// the key instead of the ID. In this case, the server will find the resource
// with the given key and use the id of the found resource to create a
// reference.
type ResourceIdentifier struct {
	// The unique ID of the referenced resource.
	ID string `json:"id,omitempty"`

	// The unique key of the referenced resource.
	Key string `json:"key,omitempty"`

	// The typeId is optional, but if given, it must match the expected
	// reference type of the referenced resource.
	TypeID string `json:"typeId,omitempty"`
}

// UpdateAction contains the payload for an update.
type UpdateAction interface {
	MarshalJSON() ([]byte, error)
}

// UpdateActions bundles multiple update actions for a single request.
type UpdateActions []UpdateAction

// Money is a JSON object combining a currency and an amount (in cents).
type Money struct {
	// When creating values of Money it’s not necessary to specify this field
	// since it’s the default.
	Type string `type:"type"`

	// The currency code compliant to ISO 4217.
	CurrencyCode string `type:"currencyCode"`

	// The amount in cents (the smallest indivisible unit of the currency).
	CentAmount int `type:"centAmount"`

	// For money type it’s equal to the number of default fraction digits for a
	// currency, can be omitted since it’s always equal to currency fraction
	// digits.
	FractionDigits int `type:"fractionDigits"`
}

// HighPrecisionMoney is a JSON object combining a currency and an amount below
// the smallest indivisible unit of the currency.
type HighPrecisionMoney struct {
	Money

	// The amount in 1 / (10 * fractionDigits) of a currency.
	PreciseAmount int `type:"preciseAmount"`
}

// Asset can be used to represent media assets, such as images, videos or PDFs.
type Asset struct {
	ID          string          `type:"id"`
	Key         string          `type:"key"`
	Sources     []AssetSource   `type:"sources"`
	Name        LocalizedString `json:"name"`
	Description LocalizedString `json:"description,omitempty"`
	Tags        []string        `json:"tags,omitempty"`
	// Custom
}

// AssetDraft is given as payload.
type AssetDraft struct {
	// User-defined identifier for the asset.
	Key string `type:"key"`

	// Requires at least one entry
	Sources []AssetSource `type:"sources"`

	Name        LocalizedString `json:"name"`
	Description LocalizedString `json:"description,omitempty"`
	Tags        []string        `json:"tags,omitempty"`
	// Custom
}

// AssetSource is a representation of an Asset in a specific format, e.g. a
// video in a certain encoding, or an image in a certain resolution.
type AssetSource struct {
	URI        string          `type:"uri"`
	Key        string          `type:"key"`
	Dimensions AssetDimensions `type:"dimensions,omitempty"`
}

// AssetDimensions indicates the width and height of the asset source.
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

// TextInputHint provides a visual representation type for a field. It is only
// relevant for string-based field types like StringType and
// LocalizedStringType.
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
