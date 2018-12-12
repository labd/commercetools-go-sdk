// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// MoneyType is an enum type
type MoneyType string

// Enum values for MoneyType
const (
	MoneyTypeCentPrecision MoneyType = "centPrecision"
	MoneyTypeHighPrecision MoneyType = "highPrecision"
)

// ReferenceTypeID is an enum type
type ReferenceTypeID string

// Enum values for ReferenceTypeID
const (
	ReferenceTypeIDCart             ReferenceTypeID = "cart"
	ReferenceTypeIDCartDiscount     ReferenceTypeID = "cart-discount"
	ReferenceTypeIDCategory         ReferenceTypeID = "category"
	ReferenceTypeIDChannel          ReferenceTypeID = "channel"
	ReferenceTypeIDCustomer         ReferenceTypeID = "customer"
	ReferenceTypeIDCustomerGroup    ReferenceTypeID = "customer-group"
	ReferenceTypeIDDiscountCode     ReferenceTypeID = "discount-code"
	ReferenceTypeIDKeyValueDocument ReferenceTypeID = "key-value-document"
	ReferenceTypeIDPayment          ReferenceTypeID = "payment"
	ReferenceTypeIDProduct          ReferenceTypeID = "product"
	ReferenceTypeIDProductType      ReferenceTypeID = "product-type"
	ReferenceTypeIDProductDiscount  ReferenceTypeID = "product-discount"
	ReferenceTypeIDOrder            ReferenceTypeID = "order"
	ReferenceTypeIDReview           ReferenceTypeID = "review"
	ReferenceTypeIDShoppingList     ReferenceTypeID = "shopping-list"
	ReferenceTypeIDShippingMethod   ReferenceTypeID = "shipping-method"
	ReferenceTypeIDState            ReferenceTypeID = "state"
	ReferenceTypeIDTaxCategory      ReferenceTypeID = "tax-category"
	ReferenceTypeIDType             ReferenceTypeID = "type"
	ReferenceTypeIDZone             ReferenceTypeID = "zone"
	ReferenceTypeIDInventoryEntry   ReferenceTypeID = "inventory-entry"
	ReferenceTypeIDOrderEdit        ReferenceTypeID = "order-edit"
)

// CountryCode is of type string
type CountryCode string

// CurrencyCode is of type string
type CurrencyCode string

// Locale is of type string
type Locale string

// LocalizedString is a map
type LocalizedString map[string]string

// GeoJSON uses type as discriminator attribute
type GeoJSON interface{}

func mapDiscriminatorGeoJSON(input interface{}) GeoJSON {
	discriminator := input.(map[string]interface{})["type"]
	switch discriminator {
	case "Point":
		new := GeoJSONPoint{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}

// Reference uses typeId as discriminator attribute
type Reference interface{}

func mapDiscriminatorReference(input interface{}) Reference {
	discriminator := input.(map[string]interface{})["typeId"]
	switch discriminator {
	case "cart-discount":
		new := CartDiscountReference{}
		mapstructure.Decode(input, &new)
		return new
	case "cart":
		new := CartReference{}
		mapstructure.Decode(input, &new)
		return new
	case "category":
		new := CategoryReference{}
		mapstructure.Decode(input, &new)
		return new
	case "channel":
		new := ChannelReference{}
		mapstructure.Decode(input, &new)
		return new
	case "key-value-document":
		new := CustomObjectReference{}
		mapstructure.Decode(input, &new)
		return new
	case "customer-group":
		new := CustomerGroupReference{}
		mapstructure.Decode(input, &new)
		return new
	case "customer":
		new := CustomerReference{}
		mapstructure.Decode(input, &new)
		return new
	case "discount-code":
		new := DiscountCodeReference{}
		mapstructure.Decode(input, &new)
		return new
	case "inventory-entry":
		new := InventoryEntryReference{}
		mapstructure.Decode(input, &new)
		return new
	case "order-edit":
		new := OrderEditReference{}
		mapstructure.Decode(input, &new)
		return new
	case "order":
		new := OrderReference{}
		mapstructure.Decode(input, &new)
		return new
	case "payment":
		new := PaymentReference{}
		mapstructure.Decode(input, &new)
		return new
	case "product-discount":
		new := ProductDiscountReference{}
		mapstructure.Decode(input, &new)
		return new
	case "product":
		new := ProductReference{}
		mapstructure.Decode(input, &new)
		return new
	case "product-type":
		new := ProductTypeReference{}
		mapstructure.Decode(input, &new)
		return new
	case "review":
		new := ReviewReference{}
		mapstructure.Decode(input, &new)
		return new
	case "shipping-method":
		new := ShippingMethodReference{}
		mapstructure.Decode(input, &new)
		return new
	case "shopping-list":
		new := ShoppingListReference{}
		mapstructure.Decode(input, &new)
		return new
	case "state":
		new := StateReference{}
		mapstructure.Decode(input, &new)
		return new
	case "tax-category":
		new := TaxCategoryReference{}
		mapstructure.Decode(input, &new)
		return new
	case "type":
		new := TypeReference{}
		mapstructure.Decode(input, &new)
		return new
	case "zone":
		new := ZoneReference{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}

// TypedMoney uses type as discriminator attribute
type TypedMoney interface{}

func mapDiscriminatorTypedMoney(input interface{}) TypedMoney {
	discriminator := input.(map[string]interface{})["type"]
	switch discriminator {
	case "centPrecision":
		new := CentPrecisionMoney{}
		mapstructure.Decode(input, &new)
		return new
	case "highPrecision":
		new := HighPrecisionMoney{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}

// Address is a standalone struct
type Address struct {
	Title                 string      `json:"title,omitempty"`
	StreetNumber          string      `json:"streetNumber,omitempty"`
	StreetName            string      `json:"streetName,omitempty"`
	State                 string      `json:"state,omitempty"`
	Salutation            string      `json:"salutation,omitempty"`
	Region                string      `json:"region,omitempty"`
	PostalCode            string      `json:"postalCode,omitempty"`
	Phone                 string      `json:"phone,omitempty"`
	POBox                 string      `json:"pOBox,omitempty"`
	Mobile                string      `json:"mobile,omitempty"`
	LastName              string      `json:"lastName,omitempty"`
	Key                   string      `json:"key,omitempty"`
	ID                    string      `json:"id,omitempty"`
	FirstName             string      `json:"firstName,omitempty"`
	Fax                   string      `json:"fax,omitempty"`
	ExternalID            string      `json:"externalId,omitempty"`
	Email                 string      `json:"email,omitempty"`
	Department            string      `json:"department,omitempty"`
	Country               CountryCode `json:"country"`
	Company               string      `json:"company,omitempty"`
	City                  string      `json:"city,omitempty"`
	Building              string      `json:"building,omitempty"`
	Apartment             string      `json:"apartment,omitempty"`
	AdditionalStreetInfo  string      `json:"additionalStreetInfo,omitempty"`
	AdditionalAddressInfo string      `json:"additionalAddressInfo,omitempty"`
}

// Asset is a standalone struct
type Asset struct {
	Tags        []string         `json:"tags,omitempty"`
	Sources     []AssetSource    `json:"sources"`
	Name        *LocalizedString `json:"name"`
	Key         string           `json:"key,omitempty"`
	ID          string           `json:"id"`
	Description *LocalizedString `json:"description,omitempty"`
	Custom      *CustomFields    `json:"custom,omitempty"`
}

// AssetDimensions is a standalone struct
type AssetDimensions struct {
	W float64 `json:"w"`
	H float64 `json:"h"`
}

// AssetDraft is a standalone struct
type AssetDraft struct {
	Tags        []string           `json:"tags,omitempty"`
	Sources     []AssetSource      `json:"sources"`
	Name        *LocalizedString   `json:"name"`
	Key         string             `json:"key,omitempty"`
	Description *LocalizedString   `json:"description,omitempty"`
	Custom      *CustomFieldsDraft `json:"custom,omitempty"`
}

// AssetSource is a standalone struct
type AssetSource struct {
	URI         string           `json:"uri"`
	Key         string           `json:"key,omitempty"`
	Dimensions  *AssetDimensions `json:"dimensions,omitempty"`
	ContentType string           `json:"contentType,omitempty"`
}

// CentPrecisionMoney implements the interface TypedMoney
type CentPrecisionMoney struct {
	CurrencyCode   CurrencyCode `json:"currencyCode"`
	CentAmount     int          `json:"centAmount"`
	FractionDigits float64      `json:"fractionDigits"`
}

// MarshalJSON override to set the discriminator value
func (obj CentPrecisionMoney) MarshalJSON() ([]byte, error) {
	type Alias CentPrecisionMoney
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "centPrecision", Alias: (*Alias)(&obj)})
}

// DiscountedPrice is a standalone struct
type DiscountedPrice struct {
	Value    *Money                    `json:"value"`
	Discount *ProductDiscountReference `json:"discount"`
}

// GeoJSONPoint implements the interface GeoJSON
type GeoJSONPoint struct {
	Coordinates []float64 `json:"coordinates"`
}

// MarshalJSON override to set the discriminator value
func (obj GeoJSONPoint) MarshalJSON() ([]byte, error) {
	type Alias GeoJSONPoint
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "Point", Alias: (*Alias)(&obj)})
}

// HighPrecisionMoney implements the interface TypedMoney
type HighPrecisionMoney struct {
	CurrencyCode   CurrencyCode `json:"currencyCode"`
	CentAmount     int          `json:"centAmount"`
	FractionDigits float64      `json:"fractionDigits"`
	PreciseAmount  int          `json:"preciseAmount"`
}

// MarshalJSON override to set the discriminator value
func (obj HighPrecisionMoney) MarshalJSON() ([]byte, error) {
	type Alias HighPrecisionMoney
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "highPrecision", Alias: (*Alias)(&obj)})
}

// Image is a standalone struct
type Image struct {
	URL        string           `json:"url"`
	Label      string           `json:"label,omitempty"`
	Dimensions *ImageDimensions `json:"dimensions"`
}

// ImageDimensions is a standalone struct
type ImageDimensions struct {
	W float64 `json:"w"`
	H float64 `json:"h"`
}

// Money is a standalone struct
type Money struct {
	CurrencyCode CurrencyCode `json:"currencyCode"`
	CentAmount   int          `json:"centAmount"`
}

// Price is a standalone struct
type Price struct {
	Value         *Money                  `json:"value"`
	ValidUntil    time.Time               `json:"validUntil,omitempty"`
	ValidFrom     time.Time               `json:"validFrom,omitempty"`
	Tiers         []PriceTier             `json:"tiers,omitempty"`
	ID            string                  `json:"id,omitempty"`
	Discounted    *DiscountedPrice        `json:"discounted,omitempty"`
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
	Custom        *CustomFields           `json:"custom,omitempty"`
	Country       CountryCode             `json:"country,omitempty"`
	Channel       *ChannelReference       `json:"channel,omitempty"`
}

// PriceDraft is a standalone struct
type PriceDraft struct {
	Value         *Money                  `json:"value"`
	ValidUntil    time.Time               `json:"validUntil,omitempty"`
	ValidFrom     time.Time               `json:"validFrom,omitempty"`
	Tiers         []PriceTier             `json:"tiers,omitempty"`
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
	Custom        *CustomFieldsDraft      `json:"custom,omitempty"`
	Country       CountryCode             `json:"country,omitempty"`
	Channel       *ChannelReference       `json:"channel,omitempty"`
}

// PriceTier is a standalone struct
type PriceTier struct {
	Value           *Money `json:"value"`
	MinimumQuantity int    `json:"minimumQuantity"`
}

// Resource is a standalone struct
type Resource struct {
	Version        int       `json:"version"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	ID             string    `json:"id"`
	CreatedAt      time.Time `json:"createdAt"`
}

// ResourceIdentifier is a standalone struct
type ResourceIdentifier struct {
	TypeID ReferenceTypeID `json:"typeId,omitempty"`
	Key    string          `json:"key,omitempty"`
	ID     string          `json:"id,omitempty"`
}

// ScopedPrice is a standalone struct
type ScopedPrice struct {
	Value         TypedMoney              `json:"value"`
	ValidUntil    time.Time               `json:"validUntil,omitempty"`
	ValidFrom     time.Time               `json:"validFrom,omitempty"`
	ID            string                  `json:"id"`
	Discounted    *DiscountedPrice        `json:"discounted,omitempty"`
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
	Custom        *CustomFields           `json:"custom,omitempty"`
	CurrentValue  TypedMoney              `json:"currentValue"`
	Country       CountryCode             `json:"country,omitempty"`
	Channel       *ChannelReference       `json:"channel,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ScopedPrice) UnmarshalJSON(data []byte) error {
	type Alias ScopedPrice
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.CurrentValue != nil {
		obj.CurrentValue = mapDiscriminatorTypedMoney(obj.CurrentValue)
	}
	if obj.Value != nil {
		obj.Value = mapDiscriminatorTypedMoney(obj.Value)
	}

	return nil
}
