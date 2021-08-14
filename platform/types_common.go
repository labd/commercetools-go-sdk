// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type PagedQueryResponse struct {
	Limit   int            `json:"limit"`
	Count   int            `json:"count"`
	Total   int            `json:"total,omitEmpty"`
	Offset  int            `json:"offset"`
	Results []BaseResource `json:"results"`
	Meta    *interface{}   `json:"meta,omitEmpty"`
}

type Update struct {
	Version int            `json:"version"`
	Actions []UpdateAction `json:"actions"`
}

type UpdateAction struct {
	Action string `json:"action"`
}

type Asset struct {
	Id          string           `json:"id"`
	Sources     []AssetSource    `json:"sources"`
	Name        LocalizedString  `json:"name"`
	Description *LocalizedString `json:"description,omitEmpty"`
	Tags        []string         `json:"tags,omitEmpty"`
	Custom      *CustomFields    `json:"custom,omitEmpty"`
	Key         string           `json:"key,omitEmpty"`
}

type AssetDimensions struct {
	W int `json:"w"`
	H int `json:"h"`
}

type AssetDraft struct {
	Sources     []AssetSource      `json:"sources"`
	Name        LocalizedString    `json:"name"`
	Description *LocalizedString   `json:"description,omitEmpty"`
	Tags        []string           `json:"tags,omitEmpty"`
	Custom      *CustomFieldsDraft `json:"custom,omitEmpty"`
	Key         string             `json:"key,omitEmpty"`
}

type AssetSource struct {
	Uri         string           `json:"uri"`
	Key         string           `json:"key,omitEmpty"`
	Dimensions  *AssetDimensions `json:"dimensions,omitEmpty"`
	ContentType string           `json:"contentType,omitEmpty"`
}

type BaseAddress struct {
	Id                   string `json:"id,omitEmpty"`
	Key                  string `json:"key,omitEmpty"`
	Title                string `json:"title,omitEmpty"`
	Salutation           string `json:"salutation,omitEmpty"`
	FirstName            string `json:"firstName,omitEmpty"`
	LastName             string `json:"lastName,omitEmpty"`
	StreetName           string `json:"streetName,omitEmpty"`
	StreetNumber         string `json:"streetNumber,omitEmpty"`
	AdditionalStreetInfo string `json:"additionalStreetInfo,omitEmpty"`
	PostalCode           string `json:"postalCode,omitEmpty"`
	City                 string `json:"city,omitEmpty"`
	Region               string `json:"region,omitEmpty"`
	State                string `json:"state,omitEmpty"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country               string `json:"country"`
	Company               string `json:"company,omitEmpty"`
	Department            string `json:"department,omitEmpty"`
	Building              string `json:"building,omitEmpty"`
	Apartment             string `json:"apartment,omitEmpty"`
	POBox                 string `json:"pOBox,omitEmpty"`
	Phone                 string `json:"phone,omitEmpty"`
	Mobile                string `json:"mobile,omitEmpty"`
	Email                 string `json:"email,omitEmpty"`
	Fax                   string `json:"fax,omitEmpty"`
	AdditionalAddressInfo string `json:"additionalAddressInfo,omitEmpty"`
	ExternalId            string `json:"externalId,omitEmpty"`
}

type Address struct {
	Id                   string `json:"id,omitEmpty"`
	Key                  string `json:"key,omitEmpty"`
	Title                string `json:"title,omitEmpty"`
	Salutation           string `json:"salutation,omitEmpty"`
	FirstName            string `json:"firstName,omitEmpty"`
	LastName             string `json:"lastName,omitEmpty"`
	StreetName           string `json:"streetName,omitEmpty"`
	StreetNumber         string `json:"streetNumber,omitEmpty"`
	AdditionalStreetInfo string `json:"additionalStreetInfo,omitEmpty"`
	PostalCode           string `json:"postalCode,omitEmpty"`
	City                 string `json:"city,omitEmpty"`
	Region               string `json:"region,omitEmpty"`
	State                string `json:"state,omitEmpty"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country               string        `json:"country"`
	Company               string        `json:"company,omitEmpty"`
	Department            string        `json:"department,omitEmpty"`
	Building              string        `json:"building,omitEmpty"`
	Apartment             string        `json:"apartment,omitEmpty"`
	POBox                 string        `json:"pOBox,omitEmpty"`
	Phone                 string        `json:"phone,omitEmpty"`
	Mobile                string        `json:"mobile,omitEmpty"`
	Email                 string        `json:"email,omitEmpty"`
	Fax                   string        `json:"fax,omitEmpty"`
	AdditionalAddressInfo string        `json:"additionalAddressInfo,omitEmpty"`
	ExternalId            string        `json:"externalId,omitEmpty"`
	Custom                *CustomFields `json:"custom,omitEmpty"`
}

type AddressDraft struct {
	Id                   string `json:"id,omitEmpty"`
	Key                  string `json:"key,omitEmpty"`
	Title                string `json:"title,omitEmpty"`
	Salutation           string `json:"salutation,omitEmpty"`
	FirstName            string `json:"firstName,omitEmpty"`
	LastName             string `json:"lastName,omitEmpty"`
	StreetName           string `json:"streetName,omitEmpty"`
	StreetNumber         string `json:"streetNumber,omitEmpty"`
	AdditionalStreetInfo string `json:"additionalStreetInfo,omitEmpty"`
	PostalCode           string `json:"postalCode,omitEmpty"`
	City                 string `json:"city,omitEmpty"`
	Region               string `json:"region,omitEmpty"`
	State                string `json:"state,omitEmpty"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country               string             `json:"country"`
	Company               string             `json:"company,omitEmpty"`
	Department            string             `json:"department,omitEmpty"`
	Building              string             `json:"building,omitEmpty"`
	Apartment             string             `json:"apartment,omitEmpty"`
	POBox                 string             `json:"pOBox,omitEmpty"`
	Phone                 string             `json:"phone,omitEmpty"`
	Mobile                string             `json:"mobile,omitEmpty"`
	Email                 string             `json:"email,omitEmpty"`
	Fax                   string             `json:"fax,omitEmpty"`
	AdditionalAddressInfo string             `json:"additionalAddressInfo,omitEmpty"`
	ExternalId            string             `json:"externalId,omitEmpty"`
	Custom                *CustomFieldsDraft `json:"custom,omitEmpty"`
}

type BaseResource struct {
	Id             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
}

type ClientLogging struct {
	ClientId       string            `json:"clientId,omitEmpty"`
	ExternalUserId string            `json:"externalUserId,omitEmpty"`
	Customer       CustomerReference `json:"customer,omitEmpty"`
	AnonymousId    string            `json:"anonymousId,omitEmpty"`
}

type CreatedBy struct {
	ClientId       string            `json:"clientId,omitEmpty"`
	ExternalUserId string            `json:"externalUserId,omitEmpty"`
	Customer       CustomerReference `json:"customer,omitEmpty"`
	AnonymousId    string            `json:"anonymousId,omitEmpty"`
}

type DiscountedPrice struct {
	Value    Money                    `json:"value"`
	Discount ProductDiscountReference `json:"discount"`
}

type GeoJson interface{}

func mapDiscriminatorGeoJson(input interface{}) (GeoJson, error) {

	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}

	switch discriminator {
	case "Point":
		new := GeoJsonPoint{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type GeoJsonPoint struct {
	Coordinates []float64 `json:"coordinates"`
}

// MarshalJSON override to set the discriminator value
func (obj GeoJsonPoint) MarshalJSON() ([]byte, error) {
	type Alias GeoJsonPoint
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "Point", Alias: (*Alias)(&obj)})
}

type Image struct {
	Url        string          `json:"url"`
	Dimensions ImageDimensions `json:"dimensions"`
	Label      string          `json:"label,omitEmpty"`
}

type ImageDimensions struct {
	W int `json:"w"`
	H int `json:"h"`
}

type KeyReference interface{}

func mapDiscriminatorKeyReference(input interface{}) (KeyReference, error) {

	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["typeId"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'typeId'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}

	switch discriminator {
	case "store":
		new := StoreKeyReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type LastModifiedBy struct {
	ClientId       string            `json:"clientId,omitEmpty"`
	ExternalUserId string            `json:"externalUserId,omitEmpty"`
	Customer       CustomerReference `json:"customer,omitEmpty"`
	AnonymousId    string            `json:"anonymousId,omitEmpty"`
}

// LocalizedString is something
type LocalizedString map[string]string
type Money struct {
	CentAmount int `json:"centAmount"`
	// The currency code compliant to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	CurrencyCode string `json:"currencyCode"`
}

type MoneyType string

const (
	MoneyTypeCentPrecision MoneyType = "centPrecision"
	MoneyTypeHighPrecision MoneyType = "highPrecision"
)

type Price struct {
	Id    string     `json:"id"`
	Value TypedMoney `json:"value"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country       string                 `json:"country,omitEmpty"`
	CustomerGroup CustomerGroupReference `json:"customerGroup,omitEmpty"`
	Channel       ChannelReference       `json:"channel,omitEmpty"`
	ValidFrom     time.Time              `json:"validFrom,omitEmpty"`
	ValidUntil    time.Time              `json:"validUntil,omitEmpty"`
	Discounted    *DiscountedPrice       `json:"discounted,omitEmpty"`
	Custom        *CustomFields          `json:"custom,omitEmpty"`
	Tiers         []PriceTier            `json:"tiers,omitEmpty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *Price) UnmarshalJSON(data []byte) error {
	type Alias Price
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Value != nil {
		var err error
		obj.Value, err = mapDiscriminatorTypedMoney(obj.Value)
		if err != nil {
			return err
		}
	}
	return nil
}

type PriceDraft struct {
	Value Money `json:"value"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country       string                          `json:"country,omitEmpty"`
	CustomerGroup CustomerGroupResourceIdentifier `json:"customerGroup,omitEmpty"`
	Channel       ChannelResourceIdentifier       `json:"channel,omitEmpty"`
	ValidFrom     time.Time                       `json:"validFrom,omitEmpty"`
	ValidUntil    time.Time                       `json:"validUntil,omitEmpty"`
	Custom        *CustomFieldsDraft              `json:"custom,omitEmpty"`
	Tiers         []PriceTierDraft                `json:"tiers,omitEmpty"`
	Discounted    *DiscountedPrice                `json:"discounted,omitEmpty"`
}

type PriceTier struct {
	MinimumQuantity int        `json:"minimumQuantity"`
	Value           TypedMoney `json:"value"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *PriceTier) UnmarshalJSON(data []byte) error {
	type Alias PriceTier
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Value != nil {
		var err error
		obj.Value, err = mapDiscriminatorTypedMoney(obj.Value)
		if err != nil {
			return err
		}
	}
	return nil
}

type PriceTierDraft struct {
	MinimumQuantity int   `json:"minimumQuantity"`
	Value           Money `json:"value"`
}

type QueryPrice struct {
	Id    string `json:"id"`
	Value Money  `json:"value"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country       string                 `json:"country,omitEmpty"`
	CustomerGroup CustomerGroupReference `json:"customerGroup,omitEmpty"`
	Channel       ChannelReference       `json:"channel,omitEmpty"`
	ValidFrom     time.Time              `json:"validFrom,omitEmpty"`
	ValidUntil    time.Time              `json:"validUntil,omitEmpty"`
	Discounted    *DiscountedPrice       `json:"discounted,omitEmpty"`
	Custom        *CustomFields          `json:"custom,omitEmpty"`
	Tiers         []PriceTierDraft       `json:"tiers,omitEmpty"`
}

type Reference interface{}

func mapDiscriminatorReference(input interface{}) (Reference, error) {

	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["typeId"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'typeId'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}

	switch discriminator {
	case "cart-discount":
		new := CartDiscountReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "cart":
		new := CartReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "category":
		new := CategoryReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "channel":
		new := ChannelReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "key-value-document":
		new := CustomObjectReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "customer-group":
		new := CustomerGroupReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "customer":
		new := CustomerReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "discount-code":
		new := DiscountCodeReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "inventory-entry":
		new := InventoryEntryReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "order-edit":
		new := OrderEditReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "order":
		new := OrderReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "payment":
		new := PaymentReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "product-discount":
		new := ProductDiscountReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "product-type":
		new := ProductTypeReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "product":
		new := ProductReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "review":
		new := ReviewReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "shipping-method":
		new := ShippingMethodReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "shopping-list":
		new := ShoppingListReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "state":
		new := StateReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "store":
		new := StoreReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "tax-category":
		new := TaxCategoryReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "type":
		new := TypeReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "zone":
		new := ZoneReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type ReferenceTypeId string

const (
	ReferenceTypeIdCart                  ReferenceTypeId = "cart"
	ReferenceTypeIdCartDiscount          ReferenceTypeId = "cart-discount"
	ReferenceTypeIdCategory              ReferenceTypeId = "category"
	ReferenceTypeIdChannel               ReferenceTypeId = "channel"
	ReferenceTypeIdCustomer              ReferenceTypeId = "customer"
	ReferenceTypeIdCustomerEmailToken    ReferenceTypeId = "customer-email-token"
	ReferenceTypeIdCustomerGroup         ReferenceTypeId = "customer-group"
	ReferenceTypeIdCustomerPasswordToken ReferenceTypeId = "customer-password-token"
	ReferenceTypeIdDiscountCode          ReferenceTypeId = "discount-code"
	ReferenceTypeIdExtension             ReferenceTypeId = "extension"
	ReferenceTypeIdInventoryEntry        ReferenceTypeId = "inventory-entry"
	ReferenceTypeIdKeyValueDocument      ReferenceTypeId = "key-value-document"
	ReferenceTypeIdOrder                 ReferenceTypeId = "order"
	ReferenceTypeIdOrderEdit             ReferenceTypeId = "order-edit"
	ReferenceTypeIdPayment               ReferenceTypeId = "payment"
	ReferenceTypeIdProduct               ReferenceTypeId = "product"
	ReferenceTypeIdProductDiscount       ReferenceTypeId = "product-discount"
	ReferenceTypeIdProductType           ReferenceTypeId = "product-type"
	ReferenceTypeIdReview                ReferenceTypeId = "review"
	ReferenceTypeIdShippingMethod        ReferenceTypeId = "shipping-method"
	ReferenceTypeIdShoppingList          ReferenceTypeId = "shopping-list"
	ReferenceTypeIdState                 ReferenceTypeId = "state"
	ReferenceTypeIdStore                 ReferenceTypeId = "store"
	ReferenceTypeIdSubscription          ReferenceTypeId = "subscription"
	ReferenceTypeIdTaxCategory           ReferenceTypeId = "tax-category"
	ReferenceTypeIdType                  ReferenceTypeId = "type"
	ReferenceTypeIdZone                  ReferenceTypeId = "zone"
)

type ResourceIdentifier interface{}

func mapDiscriminatorResourceIdentifier(input interface{}) (ResourceIdentifier, error) {

	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["typeId"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'typeId'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}

	switch discriminator {
	case "cart-discount":
		new := CartDiscountResourceIdentifier{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "cart":
		new := CartResourceIdentifier{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "category":
		new := CategoryResourceIdentifier{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "channel":
		new := ChannelResourceIdentifier{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "customer-group":
		new := CustomerGroupResourceIdentifier{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "customer":
		new := CustomerResourceIdentifier{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "discount-code":
		new := DiscountCodeResourceIdentifier{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "inventory-entry":
		new := InventoryEntryResourceIdentifier{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "order-edit":
		new := OrderEditResourceIdentifier{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "order":
		new := OrderResourceIdentifier{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "payment":
		new := PaymentResourceIdentifier{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "product-discount":
		new := ProductDiscountResourceIdentifier{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "product-type":
		new := ProductTypeResourceIdentifier{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "product":
		new := ProductResourceIdentifier{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "review":
		new := ReviewResourceIdentifier{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "shipping-method":
		new := ShippingMethodResourceIdentifier{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "shopping-list":
		new := ShoppingListResourceIdentifier{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "state":
		new := StateResourceIdentifier{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "store":
		new := StoreResourceIdentifier{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "tax-category":
		new := TaxCategoryResourceIdentifier{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "type":
		new := TypeResourceIdentifier{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "zone":
		new := ZoneResourceIdentifier{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type ScopedPrice struct {
	Id           string     `json:"id"`
	Value        TypedMoney `json:"value"`
	CurrentValue TypedMoney `json:"currentValue"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country       string                 `json:"country,omitEmpty"`
	CustomerGroup CustomerGroupReference `json:"customerGroup,omitEmpty"`
	Channel       ChannelReference       `json:"channel,omitEmpty"`
	ValidFrom     time.Time              `json:"validFrom,omitEmpty"`
	ValidUntil    time.Time              `json:"validUntil,omitEmpty"`
	Discounted    *DiscountedPrice       `json:"discounted,omitEmpty"`
	Custom        *CustomFields          `json:"custom,omitEmpty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ScopedPrice) UnmarshalJSON(data []byte) error {
	type Alias ScopedPrice
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Value != nil {
		var err error
		obj.Value, err = mapDiscriminatorTypedMoney(obj.Value)
		if err != nil {
			return err
		}
	}
	if obj.CurrentValue != nil {
		var err error
		obj.CurrentValue, err = mapDiscriminatorTypedMoney(obj.CurrentValue)
		if err != nil {
			return err
		}
	}
	return nil
}

type TypedMoney interface{}

func mapDiscriminatorTypedMoney(input interface{}) (TypedMoney, error) {

	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}

	switch discriminator {
	case "centPrecision":
		new := CentPrecisionMoney{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "highPrecision":
		new := HighPrecisionMoney{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type CentPrecisionMoney struct {
	FractionDigits int `json:"fractionDigits"`
	CentAmount     int `json:"centAmount"`
	// The currency code compliant to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	CurrencyCode string `json:"currencyCode"`
}

// MarshalJSON override to set the discriminator value
func (obj CentPrecisionMoney) MarshalJSON() ([]byte, error) {
	type Alias CentPrecisionMoney
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "centPrecision", Alias: (*Alias)(&obj)})
}

type HighPrecisionMoney struct {
	FractionDigits int `json:"fractionDigits"`
	CentAmount     int `json:"centAmount"`
	// The currency code compliant to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	CurrencyCode  string `json:"currencyCode"`
	PreciseAmount int    `json:"preciseAmount"`
}

// MarshalJSON override to set the discriminator value
func (obj HighPrecisionMoney) MarshalJSON() ([]byte, error) {
	type Alias HighPrecisionMoney
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "highPrecision", Alias: (*Alias)(&obj)})
}

type TypedMoneyDraft interface{}

func mapDiscriminatorTypedMoneyDraft(input interface{}) (TypedMoneyDraft, error) {

	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}

	switch discriminator {
	case "centPrecision":
		new := CentPrecisionMoneyDraft{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "highPrecision":
		new := HighPrecisionMoneyDraft{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type CentPrecisionMoneyDraft struct {
	CentAmount int `json:"centAmount"`
	// The currency code compliant to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	CurrencyCode   string `json:"currencyCode"`
	FractionDigits int    `json:"fractionDigits,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj CentPrecisionMoneyDraft) MarshalJSON() ([]byte, error) {
	type Alias CentPrecisionMoneyDraft
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "centPrecision", Alias: (*Alias)(&obj)})
}

type HighPrecisionMoneyDraft struct {
	CentAmount int `json:"centAmount"`
	// The currency code compliant to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	CurrencyCode   string `json:"currencyCode"`
	FractionDigits int    `json:"fractionDigits,omitEmpty"`
	PreciseAmount  int    `json:"preciseAmount"`
}

// MarshalJSON override to set the discriminator value
func (obj HighPrecisionMoneyDraft) MarshalJSON() ([]byte, error) {
	type Alias HighPrecisionMoneyDraft
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "highPrecision", Alias: (*Alias)(&obj)})
}
