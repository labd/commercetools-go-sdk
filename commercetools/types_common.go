// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

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

type Asset struct {
	Tags        []string         `json:"tags,omitempty"`
	Sources     []AssetSource    `json:"sources"`
	Name        *LocalizedString `json:"name"`
	Key         string           `json:"key,omitempty"`
	ID          string           `json:"id"`
	Description *LocalizedString `json:"description,omitempty"`
	Custom      *CustomFields    `json:"custom,omitempty"`
}

type AssetDimensions struct {
	W float64 `json:"w"`
	H float64 `json:"h"`
}

type AssetDraft struct {
	Tags        []string           `json:"tags,omitempty"`
	Sources     []AssetSource      `json:"sources"`
	Name        *LocalizedString   `json:"name"`
	Key         string             `json:"key,omitempty"`
	Description *LocalizedString   `json:"description,omitempty"`
	Custom      *CustomFieldsDraft `json:"custom,omitempty"`
}

type AssetSource struct {
	Uri         string           `json:"uri"`
	Key         string           `json:"key,omitempty"`
	Dimensions  *AssetDimensions `json:"dimensions,omitempty"`
	ContentType string           `json:"contentType,omitempty"`
}

type CentPrecisionMoney struct {
	CurrencyCode   CurrencyCode `json:"currencyCode"`
	CentAmount     int          `json:"centAmount"`
	FractionDigits float64      `json:"fractionDigits"`
}

func (obj CentPrecisionMoney) MarshalJSON() ([]byte, error) {
	type Alias CentPrecisionMoney
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "centPrecision", Alias: (*Alias)(&obj)})
}

type CountryCode string
type CurrencyCode string

type DiscountedPrice struct {
	Value    *Money                    `json:"value"`
	Discount *ProductDiscountReference `json:"discount"`
}

type GeoJson interface{}
type AbstractGeoJson struct {
	Coordinates []float64 `json:"coordinates"`
}

func AbstractGeoJsonDiscriminatorMapping(input GeoJson) GeoJson {
	discriminator := input.(map[string]interface{})["type"]
	switch discriminator {
	case "Point":
		new := GeoJsonPoint{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}

type GeoJsonPoint struct {
	Coordinates []float64 `json:"coordinates"`
}

func (obj GeoJsonPoint) MarshalJSON() ([]byte, error) {
	type Alias GeoJsonPoint
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "Point", Alias: (*Alias)(&obj)})
}

type HighPrecisionMoney struct {
	CurrencyCode   CurrencyCode `json:"currencyCode"`
	CentAmount     int          `json:"centAmount"`
	FractionDigits float64      `json:"fractionDigits"`
	PreciseAmount  int          `json:"preciseAmount"`
}

func (obj HighPrecisionMoney) MarshalJSON() ([]byte, error) {
	type Alias HighPrecisionMoney
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "highPrecision", Alias: (*Alias)(&obj)})
}

type Image struct {
	URL        string           `json:"url"`
	Label      string           `json:"label,omitempty"`
	Dimensions *ImageDimensions `json:"dimensions"`
}

type ImageDimensions struct {
	W float64 `json:"w"`
	H float64 `json:"h"`
}
type Locale string
type LocalizedString map[string]string

type Money struct {
	CurrencyCode CurrencyCode `json:"currencyCode"`
	CentAmount   int          `json:"centAmount"`
}
type MoneyType string

const (
	MoneyTypeCentPrecision MoneyType = "centPrecision"
	MoneyTypeHighPrecision MoneyType = "highPrecision"
)

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

type PriceTier struct {
	Value           *Money `json:"value"`
	MinimumQuantity int    `json:"minimumQuantity"`
}

type Reference interface{}
type AbstractReference struct {
	Key string `json:"key,omitempty"`
	ID  string `json:"id,omitempty"`
}

func AbstractReferenceDiscriminatorMapping(input Reference) Reference {
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

type ReferenceTypeID string

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

type Resource struct {
	Version        int       `json:"version"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	ID             string    `json:"id"`
	CreatedAt      time.Time `json:"createdAt"`
}

type ResourceIdentifier struct {
	TypeID ReferenceTypeID `json:"typeId,omitempty"`
	Key    string          `json:"key,omitempty"`
	ID     string          `json:"id,omitempty"`
}

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

func (obj *ScopedPrice) UnmarshalJSON(data []byte) error {
	type Alias ScopedPrice
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.CurrentValue != nil {
		obj.CurrentValue = AbstractTypedMoneyDiscriminatorMapping(obj.CurrentValue)
	}
	if obj.Value != nil {
		obj.Value = AbstractTypedMoneyDiscriminatorMapping(obj.Value)
	}

	return nil
}

type TypedMoney interface{}
type AbstractTypedMoney struct {
	CurrencyCode   CurrencyCode `json:"currencyCode"`
	CentAmount     int          `json:"centAmount"`
	FractionDigits float64      `json:"fractionDigits"`
}

func AbstractTypedMoneyDiscriminatorMapping(input TypedMoney) TypedMoney {
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
