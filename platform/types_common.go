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
	Total   *int           `json:"total,omitempty"`
	Offset  int            `json:"offset"`
	Results []BaseResource `json:"results"`
	Meta    *interface{}   `json:"meta,omitempty"`
}

type Update struct {
	Version int            `json:"version"`
	Actions []UpdateAction `json:"actions"`
}

type UpdateAction struct {
	Action string `json:"action"`
}

type Asset struct {
	ID          string           `json:"id"`
	Sources     []AssetSource    `json:"sources"`
	Name        LocalizedString  `json:"name"`
	Description *LocalizedString `json:"description,omitempty"`
	Tags        []string         `json:"tags,omitempty"`
	Custom      *CustomFields    `json:"custom,omitempty"`
	Key         *string          `json:"key,omitempty"`
}

type AssetDimensions struct {
	W int `json:"w"`
	H int `json:"h"`
}

type AssetDraft struct {
	Sources     []AssetSource      `json:"sources"`
	Name        LocalizedString    `json:"name"`
	Description *LocalizedString   `json:"description,omitempty"`
	Tags        []string           `json:"tags,omitempty"`
	Custom      *CustomFieldsDraft `json:"custom,omitempty"`
	Key         *string            `json:"key,omitempty"`
}

type AssetSource struct {
	Uri         string           `json:"uri"`
	Key         *string          `json:"key,omitempty"`
	Dimensions  *AssetDimensions `json:"dimensions,omitempty"`
	ContentType *string          `json:"contentType,omitempty"`
}

type BaseAddress struct {
	ID                   *string `json:"id,omitempty"`
	Key                  *string `json:"key,omitempty"`
	Title                *string `json:"title,omitempty"`
	Salutation           *string `json:"salutation,omitempty"`
	FirstName            *string `json:"firstName,omitempty"`
	LastName             *string `json:"lastName,omitempty"`
	StreetName           *string `json:"streetName,omitempty"`
	StreetNumber         *string `json:"streetNumber,omitempty"`
	AdditionalStreetInfo *string `json:"additionalStreetInfo,omitempty"`
	PostalCode           *string `json:"postalCode,omitempty"`
	City                 *string `json:"city,omitempty"`
	Region               *string `json:"region,omitempty"`
	State                *string `json:"state,omitempty"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country               string  `json:"country"`
	Company               *string `json:"company,omitempty"`
	Department            *string `json:"department,omitempty"`
	Building              *string `json:"building,omitempty"`
	Apartment             *string `json:"apartment,omitempty"`
	POBox                 *string `json:"pOBox,omitempty"`
	Phone                 *string `json:"phone,omitempty"`
	Mobile                *string `json:"mobile,omitempty"`
	Email                 *string `json:"email,omitempty"`
	Fax                   *string `json:"fax,omitempty"`
	AdditionalAddressInfo *string `json:"additionalAddressInfo,omitempty"`
	ExternalId            *string `json:"externalId,omitempty"`
}

type Address struct {
	ID                   *string `json:"id,omitempty"`
	Key                  *string `json:"key,omitempty"`
	Title                *string `json:"title,omitempty"`
	Salutation           *string `json:"salutation,omitempty"`
	FirstName            *string `json:"firstName,omitempty"`
	LastName             *string `json:"lastName,omitempty"`
	StreetName           *string `json:"streetName,omitempty"`
	StreetNumber         *string `json:"streetNumber,omitempty"`
	AdditionalStreetInfo *string `json:"additionalStreetInfo,omitempty"`
	PostalCode           *string `json:"postalCode,omitempty"`
	City                 *string `json:"city,omitempty"`
	Region               *string `json:"region,omitempty"`
	State                *string `json:"state,omitempty"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country               string        `json:"country"`
	Company               *string       `json:"company,omitempty"`
	Department            *string       `json:"department,omitempty"`
	Building              *string       `json:"building,omitempty"`
	Apartment             *string       `json:"apartment,omitempty"`
	POBox                 *string       `json:"pOBox,omitempty"`
	Phone                 *string       `json:"phone,omitempty"`
	Mobile                *string       `json:"mobile,omitempty"`
	Email                 *string       `json:"email,omitempty"`
	Fax                   *string       `json:"fax,omitempty"`
	AdditionalAddressInfo *string       `json:"additionalAddressInfo,omitempty"`
	ExternalId            *string       `json:"externalId,omitempty"`
	Custom                *CustomFields `json:"custom,omitempty"`
}

type AddressDraft struct {
	ID                   *string `json:"id,omitempty"`
	Key                  *string `json:"key,omitempty"`
	Title                *string `json:"title,omitempty"`
	Salutation           *string `json:"salutation,omitempty"`
	FirstName            *string `json:"firstName,omitempty"`
	LastName             *string `json:"lastName,omitempty"`
	StreetName           *string `json:"streetName,omitempty"`
	StreetNumber         *string `json:"streetNumber,omitempty"`
	AdditionalStreetInfo *string `json:"additionalStreetInfo,omitempty"`
	PostalCode           *string `json:"postalCode,omitempty"`
	City                 *string `json:"city,omitempty"`
	Region               *string `json:"region,omitempty"`
	State                *string `json:"state,omitempty"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country               string             `json:"country"`
	Company               *string            `json:"company,omitempty"`
	Department            *string            `json:"department,omitempty"`
	Building              *string            `json:"building,omitempty"`
	Apartment             *string            `json:"apartment,omitempty"`
	POBox                 *string            `json:"pOBox,omitempty"`
	Phone                 *string            `json:"phone,omitempty"`
	Mobile                *string            `json:"mobile,omitempty"`
	Email                 *string            `json:"email,omitempty"`
	Fax                   *string            `json:"fax,omitempty"`
	AdditionalAddressInfo *string            `json:"additionalAddressInfo,omitempty"`
	ExternalId            *string            `json:"externalId,omitempty"`
	Custom                *CustomFieldsDraft `json:"custom,omitempty"`
}

type BaseResource struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
}

type ClientLogging struct {
	ClientId       *string            `json:"clientId,omitempty"`
	ExternalUserId *string            `json:"externalUserId,omitempty"`
	Customer       *CustomerReference `json:"customer,omitempty"`
	AnonymousId    *string            `json:"anonymousId,omitempty"`
}

type CreatedBy struct {
	ClientId       *string            `json:"clientId,omitempty"`
	ExternalUserId *string            `json:"externalUserId,omitempty"`
	Customer       *CustomerReference `json:"customer,omitempty"`
	AnonymousId    *string            `json:"anonymousId,omitempty"`
}

type DiscountedPrice struct {
	Value    TypedMoney               `json:"value"`
	Discount ProductDiscountReference `json:"discount"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DiscountedPrice) UnmarshalJSON(data []byte) error {
	type Alias DiscountedPrice
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

type DiscountedPriceDraft struct {
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
		obj := GeoJsonPoint{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
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
	Label      *string         `json:"label,omitempty"`
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
		obj := StoreKeyReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type LastModifiedBy struct {
	ClientId       *string            `json:"clientId,omitempty"`
	ExternalUserId *string            `json:"externalUserId,omitempty"`
	Customer       *CustomerReference `json:"customer,omitempty"`
	AnonymousId    *string            `json:"anonymousId,omitempty"`
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
	ID    string     `json:"id"`
	Value TypedMoney `json:"value"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country *string `json:"country,omitempty"`
	// [Reference](/types#reference) to a [CustomerGroup](ctp:api:type:CustomerGroup).
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
	Channel       *ChannelReference       `json:"channel,omitempty"`
	ValidFrom     *time.Time              `json:"validFrom,omitempty"`
	ValidUntil    *time.Time              `json:"validUntil,omitempty"`
	Discounted    *DiscountedPrice        `json:"discounted,omitempty"`
	Custom        *CustomFields           `json:"custom,omitempty"`
	Tiers         []PriceTier             `json:"tiers,omitempty"`
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
	Country *string `json:"country,omitempty"`
	// [ResourceIdentifier](/../api/types#resourceidentifier) to a [CustomerGroup](ctp:api:type:CustomerGroup).
	CustomerGroup *CustomerGroupResourceIdentifier `json:"customerGroup,omitempty"`
	Channel       *ChannelResourceIdentifier       `json:"channel,omitempty"`
	ValidFrom     *time.Time                       `json:"validFrom,omitempty"`
	ValidUntil    *time.Time                       `json:"validUntil,omitempty"`
	Custom        *CustomFieldsDraft               `json:"custom,omitempty"`
	Tiers         []PriceTierDraft                 `json:"tiers,omitempty"`
	Discounted    *DiscountedPriceDraft            `json:"discounted,omitempty"`
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
	ID    string `json:"id"`
	Value Money  `json:"value"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country *string `json:"country,omitempty"`
	// [Reference](/types#reference) to a [CustomerGroup](ctp:api:type:CustomerGroup).
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
	Channel       *ChannelReference       `json:"channel,omitempty"`
	ValidFrom     *time.Time              `json:"validFrom,omitempty"`
	ValidUntil    *time.Time              `json:"validUntil,omitempty"`
	Discounted    *DiscountedPriceDraft   `json:"discounted,omitempty"`
	Custom        *CustomFields           `json:"custom,omitempty"`
	Tiers         []PriceTierDraft        `json:"tiers,omitempty"`
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
		obj := CartDiscountReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "cart":
		obj := CartReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "category":
		obj := CategoryReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "channel":
		obj := ChannelReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "key-value-document":
		obj := CustomObjectReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "customer-group":
		obj := CustomerGroupReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "customer":
		obj := CustomerReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "discount-code":
		obj := DiscountCodeReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "inventory-entry":
		obj := InventoryEntryReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "order-edit":
		obj := OrderEditReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "order":
		obj := OrderReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "payment":
		obj := PaymentReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "product-discount":
		obj := ProductDiscountReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "product-type":
		obj := ProductTypeReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "product":
		obj := ProductReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "review":
		obj := ReviewReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "shipping-method":
		obj := ShippingMethodReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "shopping-list":
		obj := ShoppingListReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "state":
		obj := StateReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "store":
		obj := StoreReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "tax-category":
		obj := TaxCategoryReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "type":
		obj := TypeReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "zone":
		obj := ZoneReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
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
		obj := CartDiscountResourceIdentifier{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "cart":
		obj := CartResourceIdentifier{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "category":
		obj := CategoryResourceIdentifier{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "channel":
		obj := ChannelResourceIdentifier{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "customer-group":
		obj := CustomerGroupResourceIdentifier{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "customer":
		obj := CustomerResourceIdentifier{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "discount-code":
		obj := DiscountCodeResourceIdentifier{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "inventory-entry":
		obj := InventoryEntryResourceIdentifier{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "order-edit":
		obj := OrderEditResourceIdentifier{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "order":
		obj := OrderResourceIdentifier{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "payment":
		obj := PaymentResourceIdentifier{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "product-discount":
		obj := ProductDiscountResourceIdentifier{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "product-type":
		obj := ProductTypeResourceIdentifier{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "product":
		obj := ProductResourceIdentifier{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "review":
		obj := ReviewResourceIdentifier{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "shipping-method":
		obj := ShippingMethodResourceIdentifier{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "shopping-list":
		obj := ShoppingListResourceIdentifier{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "state":
		obj := StateResourceIdentifier{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "store":
		obj := StoreResourceIdentifier{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "tax-category":
		obj := TaxCategoryResourceIdentifier{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "type":
		obj := TypeResourceIdentifier{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "zone":
		obj := ZoneResourceIdentifier{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type ScopedPrice struct {
	ID           string     `json:"id"`
	Value        TypedMoney `json:"value"`
	CurrentValue TypedMoney `json:"currentValue"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country *string `json:"country,omitempty"`
	// [Reference](/types#reference) to a [CustomerGroup](ctp:api:type:CustomerGroup).
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
	Channel       *ChannelReference       `json:"channel,omitempty"`
	ValidFrom     *time.Time              `json:"validFrom,omitempty"`
	ValidUntil    *time.Time              `json:"validUntil,omitempty"`
	Discounted    *DiscountedPrice        `json:"discounted,omitempty"`
	Custom        *CustomFields           `json:"custom,omitempty"`
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
		obj := CentPrecisionMoney{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "highPrecision":
		obj := HighPrecisionMoney{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
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
		obj := CentPrecisionMoneyDraft{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "highPrecision":
		obj := HighPrecisionMoneyDraft{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type CentPrecisionMoneyDraft struct {
	CentAmount int `json:"centAmount"`
	// The currency code compliant to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	CurrencyCode   string `json:"currencyCode"`
	FractionDigits *int   `json:"fractionDigits,omitempty"`
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
	FractionDigits *int   `json:"fractionDigits,omitempty"`
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
