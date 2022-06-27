package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type PagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/../api/general-concepts#strong-consistency).
	// This field is returned by default.
	// For improved performance, calculating this field can be deactivated by using the query parameter `withTotal=false`.
	// When the results are filtered with a [Query Predicate](/../api/predicates/query), `total` is subject to a [limit](/../api/limits#queries).
	Total   *int           `json:"total,omitempty"`
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
	// Unique identifier of the Asset.
	ID      string        `json:"id"`
	Sources []AssetSource `json:"sources"`
	// Name of the Asset.
	Name LocalizedString `json:"name"`
	// Description of the Asset.
	Description *LocalizedString `json:"description,omitempty"`
	// Keywords for categorizing and organizing Assets.
	Tags []string `json:"tags"`
	// Custom Fields defined for the Asset.
	Custom *CustomFields `json:"custom,omitempty"`
	// User-defined unique identifier of the Asset.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj Asset) MarshalJSON() ([]byte, error) {
	type Alias Asset
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["tags"] == nil {
		delete(raw, "tags")
	}

	return json.Marshal(raw)

}

/**
*	Dimensions of the Asset source specified by the number of pixels.
*
 */
type AssetDimensions struct {
	// Width of the Asset source.
	W int `json:"w"`
	// Height of the Asset source.
	H int `json:"h"`
}

type AssetDraft struct {
	Sources []AssetSource `json:"sources"`
	// Name of the Asset.
	Name LocalizedString `json:"name"`
	// Description of the Asset.
	Description *LocalizedString `json:"description,omitempty"`
	// Keywords for categorizing and organizing Assets.
	Tags []string `json:"tags"`
	// Custom Fields defined for the Asset.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
	// User-defined unique identifier for the Asset.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AssetDraft) MarshalJSON() ([]byte, error) {
	type Alias AssetDraft
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["tags"] == nil {
		delete(raw, "tags")
	}

	return json.Marshal(raw)

}

/**
*	Representation of an [Asset](#asset) in a specific format, for example a video in a certain encoding, or an image in a certain resolution.
*
 */
type AssetSource struct {
	// URI of the AssetSource.
	Uri string `json:"uri"`
	// User-defined unique identifier of the AssetSource.
	Key *string `json:"key,omitempty"`
	// Width and height of the AssetSource.
	Dimensions *AssetDimensions `json:"dimensions,omitempty"`
	// Indicates the type of content, for example `application/pdf`.
	ContentType *string `json:"contentType,omitempty"`
}

type BaseAddress struct {
	// Unique identifier of the Address.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the Address.
	Key *string `json:"key,omitempty"`
	// Two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country string `json:"country"`
	// Title of the contact, for example 'Dr.'
	Title *string `json:"title,omitempty"`
	// Salutation of the contact, for example 'Mr.' or 'Ms.'
	Salutation *string `json:"salutation,omitempty"`
	// Given name (first name) of the contact.
	FirstName *string `json:"firstName,omitempty"`
	// Family name (last name) of the contact.
	LastName *string `json:"lastName,omitempty"`
	// Name of the street.
	StreetName *string `json:"streetName,omitempty"`
	// Street number.
	StreetNumber *string `json:"streetNumber,omitempty"`
	// Further information on the street address.
	AdditionalStreetInfo *string `json:"additionalStreetInfo,omitempty"`
	// Postal code.
	PostalCode *string `json:"postalCode,omitempty"`
	// Name of the city.
	City *string `json:"city,omitempty"`
	// Name of the region.
	Region *string `json:"region,omitempty"`
	// Name of the state, for example, Colorado.
	State *string `json:"state,omitempty"`
	// Name of the company.
	Company *string `json:"company,omitempty"`
	// Name of the department.
	Department *string `json:"department,omitempty"`
	// Number or name of the building.
	Building *string `json:"building,omitempty"`
	// Number or name of the apartment.
	Apartment *string `json:"apartment,omitempty"`
	// Post office box number.
	POBox *string `json:"pOBox,omitempty"`
	// Phone number of the contact.
	Phone *string `json:"phone,omitempty"`
	// Mobile phone number of the contact.
	Mobile *string `json:"mobile,omitempty"`
	// Email address of the contact.
	Email *string `json:"email,omitempty"`
	// Fax number of the contact.
	Fax *string `json:"fax,omitempty"`
	// Further information on the Address.
	AdditionalAddressInfo *string `json:"additionalAddressInfo,omitempty"`
	// ID for the contact used in an external system.
	ExternalId *string `json:"externalId,omitempty"`
}

type Address struct {
	// Unique identifier of the Address.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the Address.
	Key *string `json:"key,omitempty"`
	// Two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country string `json:"country"`
	// Title of the contact, for example 'Dr.'
	Title *string `json:"title,omitempty"`
	// Salutation of the contact, for example 'Mr.' or 'Ms.'
	Salutation *string `json:"salutation,omitempty"`
	// Given name (first name) of the contact.
	FirstName *string `json:"firstName,omitempty"`
	// Family name (last name) of the contact.
	LastName *string `json:"lastName,omitempty"`
	// Name of the street.
	StreetName *string `json:"streetName,omitempty"`
	// Street number.
	StreetNumber *string `json:"streetNumber,omitempty"`
	// Further information on the street address.
	AdditionalStreetInfo *string `json:"additionalStreetInfo,omitempty"`
	// Postal code.
	PostalCode *string `json:"postalCode,omitempty"`
	// Name of the city.
	City *string `json:"city,omitempty"`
	// Name of the region.
	Region *string `json:"region,omitempty"`
	// Name of the state, for example, Colorado.
	State *string `json:"state,omitempty"`
	// Name of the company.
	Company *string `json:"company,omitempty"`
	// Name of the department.
	Department *string `json:"department,omitempty"`
	// Number or name of the building.
	Building *string `json:"building,omitempty"`
	// Number or name of the apartment.
	Apartment *string `json:"apartment,omitempty"`
	// Post office box number.
	POBox *string `json:"pOBox,omitempty"`
	// Phone number of the contact.
	Phone *string `json:"phone,omitempty"`
	// Mobile phone number of the contact.
	Mobile *string `json:"mobile,omitempty"`
	// Email address of the contact.
	Email *string `json:"email,omitempty"`
	// Fax number of the contact.
	Fax *string `json:"fax,omitempty"`
	// Further information on the Address.
	AdditionalAddressInfo *string `json:"additionalAddressInfo,omitempty"`
	// ID for the contact used in an external system.
	ExternalId *string `json:"externalId,omitempty"`
	// Custom Fields defined for the Address.
	Custom *CustomFields `json:"custom,omitempty"`
}

type AddressDraft struct {
	// Unique identifier for the Address. Not recommended to set it manually since the API overwrites this ID when creating an Address for a [Customer](ctp:api:type:Customer). Use `key` instead and omit this field from the request to let the API generate the ID for the Address.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier for the Address.
	Key *string `json:"key,omitempty"`
	// Two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country string `json:"country"`
	// Title of the contact, for example 'Dr.'
	Title *string `json:"title,omitempty"`
	// Salutation of the contact, for example 'Mr.' or 'Ms.'
	Salutation *string `json:"salutation,omitempty"`
	// Given name (first name) of the contact.
	FirstName *string `json:"firstName,omitempty"`
	// Family name (last name) of the contact.
	LastName *string `json:"lastName,omitempty"`
	// Name of the street.
	StreetName *string `json:"streetName,omitempty"`
	// Street number.
	StreetNumber *string `json:"streetNumber,omitempty"`
	// Further information on the street address.
	AdditionalStreetInfo *string `json:"additionalStreetInfo,omitempty"`
	// Postal code.
	PostalCode *string `json:"postalCode,omitempty"`
	// Name of the city.
	City *string `json:"city,omitempty"`
	// Name of the region.
	Region *string `json:"region,omitempty"`
	// Name of the state, for example, Colorado.
	State *string `json:"state,omitempty"`
	// Name of the company.
	Company *string `json:"company,omitempty"`
	// Name of the department.
	Department *string `json:"department,omitempty"`
	// Number or name of the building.
	Building *string `json:"building,omitempty"`
	// Number or name of the apartment.
	Apartment *string `json:"apartment,omitempty"`
	// Post office box number.
	POBox *string `json:"pOBox,omitempty"`
	// Phone number of the contact.
	Phone *string `json:"phone,omitempty"`
	// Mobile phone number of the contact.
	Mobile *string `json:"mobile,omitempty"`
	// Email address of the contact.
	Email *string `json:"email,omitempty"`
	// Fax number of the contact.
	Fax *string `json:"fax,omitempty"`
	// Further information on the Address.
	AdditionalAddressInfo *string `json:"additionalAddressInfo,omitempty"`
	// ID for the contact used in an external system.
	ExternalId *string `json:"externalId,omitempty"`
	// Custom Fields defined for the Address.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
}

type BaseResource struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
}

/**
*	These objects represent information about which [API Client](/../api/projects/api-clients) created or modified a resource. For more information, see [Client Logging](/client-logging).
*
 */
type ClientLogging struct {
	// `id` of the [APIClient](ctp:api:type:ApiClient) which created the resource.
	ClientId *string `json:"clientId,omitempty"`
	// [External user ID](/../api/client-logging#external-user-ids) provided by `X-External-User-ID` HTTP Header.
	ExternalUserId *string `json:"externalUserId,omitempty"`
	// Indicates the [Customer](ctp:api:type:Customer) who modified the resource using a token from the [password flow](/authorization#password-flow).
	Customer *CustomerReference `json:"customer,omitempty"`
	// Indicates that the resource was modified during an [anonymous session](/../api/authorization#tokens-for-anonymous-sessions) with the logged ID.
	AnonymousId *string `json:"anonymousId,omitempty"`
}

/**
*	Present on resources created after 1 February 2019 except for [events not tracked](/../api/client-logging#events-tracked).
 */
type CreatedBy struct {
	// `id` of the [APIClient](ctp:api:type:ApiClient) which created the resource.
	ClientId *string `json:"clientId,omitempty"`
	// [External user ID](/../api/client-logging#external-user-ids) provided by `X-External-User-ID` HTTP Header.
	ExternalUserId *string `json:"externalUserId,omitempty"`
	// Indicates the [Customer](ctp:api:type:Customer) who created the resource using a token from the [password flow](/authorization#password-flow).
	Customer *CustomerReference `json:"customer,omitempty"`
	// Indicates the [anonymous session](/../api/authorization#tokens-for-anonymous-sessions) during which the resource was created.
	AnonymousId *string `json:"anonymousId,omitempty"`
}

type DiscountedPrice struct {
	// Money value of the discounted price.
	Value TypedMoney `json:"value"`
	// [ProductDiscount](ctp:api:type:ProductDiscount) related to the discounted price.
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
	// Sets the money value for the discounted price.
	Value Money `json:"value"`
	// Relates the referenced [ProductDiscount](ctp:api:type:ProductDiscount) to the discounted price.
	Discount ProductDiscountReference `json:"discount"`
}

/**
*	GeoJSON Geometry represents a [Geometry Object](https://datatracker.ietf.org/doc/html/rfc7946#section-3.1) as defined in the GeoJSON standard.
*
 */
type GeoJson interface{}

func mapDiscriminatorGeoJson(input interface{}) (GeoJson, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("invalid data")
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
	// Longitude (stored on index `[0]`) and latitude (stored on index `[1]`) of the [Point](https://datatracker.ietf.org/doc/html/rfc7946#section-3.1.2).
	Coordinates []float64 `json:"coordinates"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

/**
*	A KeyReference represents a loose reference to another resource in the same Project identified by the resource's `key` field. If available, the `key` is immutable and mandatory. KeyReferences do not support [Reference Expansion](/general-concepts#reference-expansion).
*
 */
type KeyReference interface{}

func mapDiscriminatorKeyReference(input interface{}) (KeyReference, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["typeId"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'typeId'")
		}
	} else {
		return nil, errors.New("invalid data")
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

/**
*	Present on resources modified after 1 February 2019 except for [events not tracked](/../api/client-logging#events-tracked).
 */
type LastModifiedBy struct {
	// `id` of the [APIClient](ctp:api:type:ApiClient) which modified the resource.
	ClientId *string `json:"clientId,omitempty"`
	// [External user ID](/../api/client-logging#external-user-ids) provided by `X-External-User-ID` HTTP Header.
	ExternalUserId *string `json:"externalUserId,omitempty"`
	// Indicates the [Customer](ctp:api:type:Customer) who modified the resource using a token from the [password flow](/authorization#password-flow).
	Customer *CustomerReference `json:"customer,omitempty"`
	// Indicates the [anonymous session](/../api/authorization#tokens-for-anonymous-sessions) during which the resource was modified.
	AnonymousId *string `json:"anonymousId,omitempty"`
}

/**
*	JSON object where the keys are of type [Locale](ctp:api:type:Locale), and the values are the strings used for the corresponding language.
*
 */
type LocalizedString map[string]string

/**
*	Draft type that stores amounts in cent precision for the specified currency.
*
*	For storing money values in fractions of the minor unit in a currency, use [HighPrecisionMoneyDraft](ctp:api:type:HighPrecisionMoneyDraft) instead.
*
 */
type Money struct {
	// Amount in the smallest indivisible unit of a currency, such as:
	//
	// * Cents for EUR and USD, pence for GBP, or centime for CHF (5 CHF is specified as `500`).
	// * The value in the major unit for currencies without minor units, like JPY (5 JPY is specified as `5`).
	CentAmount int `json:"centAmount"`
	// Currency code compliant to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	CurrencyCode string `json:"currencyCode"`
}

/**
*	MoneyType supports two different values, one for amounts in cent precision and another one for sub-cent amounts up to 20 fraction digits.
 */
type MoneyType string

const (
	MoneyTypeCentPrecision MoneyType = "centPrecision"
	MoneyTypeHighPrecision MoneyType = "highPrecision"
)

type Price struct {
	// Unique identifier of this Price.
	ID string `json:"id"`
	// Money value of this Price.
	Value TypedMoney `json:"value"`
	// Country for which this Price is valid.
	Country *string `json:"country,omitempty"`
	// [CustomerGroup](ctp:api:type:CustomerGroup) for which this Price is valid.
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
	// `ProductDistribution` [Channel](ctp:api:type:Channel) for which this Price is valid.
	Channel *ChannelReference `json:"channel,omitempty"`
	// Date and time from which this Price is valid.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Date and time until this Price is valid.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
	// Is set if a [ProductDiscount](ctp:api:type:ProductDiscount) has been applied.
	// If set, the API uses the DiscountedPrice value for the [LineItem Price selection](/projects/carts#lineitem-price-selection).
	// When a [relative discount](/../api/projects/productDiscounts#productdiscountvaluerelative) has been applied and the fraction part of the DiscountedPrice `value` is 0.5, the `value` is rounded in favor of the customer with [half down rounding](https://en.wikipedia.org/wiki/Rounding#Round_half_down).
	Discounted *DiscountedPrice `json:"discounted,omitempty"`
	// Present if different Prices for certain [LineItem](ctp:api:type:LineItem) quantities have been specified.
	Tiers []PriceTier `json:"tiers"`
	// Custom Fields defined for the Price.
	Custom *CustomFields `json:"custom,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj Price) MarshalJSON() ([]byte, error) {
	type Alias Price
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["tiers"] == nil {
		delete(raw, "tiers")
	}

	return json.Marshal(raw)

}

type PriceDraft struct {
	// Money value of this Price.
	Value Money `json:"value"`
	// Set this field if this Price is only valid for the specified country.
	Country *string `json:"country,omitempty"`
	// Set this field if this Price is only valid for the referenced [CustomerGroup](ctp:api:type:CustomerGroup).
	CustomerGroup *CustomerGroupResourceIdentifier `json:"customerGroup,omitempty"`
	// Set this field if this Price is only valid for the referenced `ProductDistribution` [Channel](ctp:api:type:Channel).
	Channel *ChannelResourceIdentifier `json:"channel,omitempty"`
	// Set this field if this Price is valid only valid from the specified date and time.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Set this field if this Price is valid only valid until the specified date and time.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
	// Set this field to add a DiscountedPrice from an external service.
	//
	// The API sets this field automatically if at least one [ProductDiscount](ctp:api:type:ProductDiscount) applies.
	// The DiscountedPrice must reference a ProductDiscount with:
	//
	// * The `isActive` flag set to `true`.
	// * A [ProductDiscountValue](ctp:api:type:ProductDiscountValueExternal) of type `external`.
	// * A `predicate` that matches the [ProductVariant](ctp:api:type:ProductVariant) the Price is referenced from.
	Discounted *DiscountedPriceDraft `json:"discounted,omitempty"`
	// Set this field to specify different Prices for certain [LineItem](ctp:api:type:LineItem) quantities.
	Tiers []PriceTierDraft `json:"tiers"`
	// Custom Fields for the Price.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PriceDraft) MarshalJSON() ([]byte, error) {
	type Alias PriceDraft
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["tiers"] == nil {
		delete(raw, "tiers")
	}

	return json.Marshal(raw)

}

/**
*	A Price tier is selected instead of the default Price when a certain quantity of the [ProductVariant](ctp:api:type:ProductVariant) is [added to a Cart](/projects/carts#add-lineitem) and ordered.
*	_For example: the Price can be lower if more than 10 items are ordered._
*	If no Price tier is found for the Order quantity, the base Price is used.
*	A Price tier is applied for the entire quantity of a Product Variant put as [LineItem](/projects/carts#lineitem) in a Cart as soon as the minimum quantity for the Price tier is reached.
*	The Price tier is applied per Line Item of the Product Variant. If, for example, the same Product Variant appears in the same Cart as several Line Items, (what can be achieved by different values of a Custom Field on the Line Items) for each Line Item the minimum quantity must be reached to get the Price tier.
*
 */
type PriceTier struct {
	// Minimum quantity this Price tier is valid for.
	//
	// The minimum quantity is always greater than or equal to 2. The base Price is interpreted as valid for a minimum quantity equal to 1.
	MinimumQuantity int `json:"minimumQuantity"`
	// Money value that applies when the `minimumQuantity` is greater than or equal to the [LineItem](ctp:api:type:LineItem) `quantity`.
	//
	// The `currencyCode` of a Price tier is always the same as the `currencyCode` in the `value` of the related Price.
	Value TypedMoney `json:"value"`
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

/**
*	Specifies a Price tier that applies when the minimum quantity for the [LineItem](ctp:api:type:LineItem) of a ProductVariant with the related Price is reached in a Cart.
*
 */
type PriceTierDraft struct {
	// Minimum quantity this Price tier is valid for.
	//
	// The minimum quantity is always greater than or equal to 2. The base Price is interpreted as valid for a minimum quantity equal to 1.
	MinimumQuantity int `json:"minimumQuantity"`
	// Money value that applies when the `minimumQuantity` is greater than or equal to the [LineItem](ctp:api:type:LineItem) `quantity`.
	//
	// The `currencyCode` of a Price tier must be the same as the `currencyCode` in the `value` of the related Price.
	Value Money `json:"value"`
}

type QueryPrice struct {
	// Unique identifier of the given Price.
	ID *string `json:"id,omitempty"`
	// Money value of the given Price.
	Value Money `json:"value"`
	// Country for which the given Price is valid.
	Country *string `json:"country,omitempty"`
	// [CustomerGroup](ctp:api:type:CustomerGroup) for which the given Price is valid.
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
	// `ProductDistribution` [Channel](ctp:api:type:Channel) for which the given Price is valid.
	Channel *ChannelReference `json:"channel,omitempty"`
	// Date from which the given Price is valid.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Date until which the given Price is valid.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
	// [DiscountedPrice](ctp:api:type:DiscountedPrice) you specify for the given Price.
	Discounted *DiscountedPriceDraft `json:"discounted,omitempty"`
	// Custom Fields for the Price.
	Custom *CustomFields `json:"custom,omitempty"`
	// Price tier applied when the minimum quantity for the [LineItem](ctp:api:type:LineItem) of a ProductVariant with the related Price is reached in a Cart.
	Tiers []PriceTierDraft `json:"tiers"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj QueryPrice) MarshalJSON() ([]byte, error) {
	type Alias QueryPrice
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["tiers"] == nil {
		delete(raw, "tiers")
	}

	return json.Marshal(raw)

}

/**
*	A Reference represents a loose reference to another resource in the same Project identified by its `id`. The `typeId` indicates the type of the referenced resource. Each resource type has its corresponding Reference type, like [ChannelReference](ctp:api:type:ChannelReference).  A referenced resource can be embedded through [Reference Expansion](/general-concepts#reference-expansion). The expanded reference is the value of an additional `obj` field then.
*
 */
type Reference interface{}

func mapDiscriminatorReference(input interface{}) (Reference, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["typeId"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'typeId'")
		}
	} else {
		return nil, errors.New("invalid data")
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
	case "product-selection":
		obj := ProductSelectionReference{}
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
	case "standalone-price":
		obj := StandalonePriceReference{}
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

/**
*	supported resource type identifiers:
*
 */
type ReferenceTypeId string

const (
	ReferenceTypeIdCart             ReferenceTypeId = "cart"
	ReferenceTypeIdCartDiscount     ReferenceTypeId = "cart-discount"
	ReferenceTypeIdCategory         ReferenceTypeId = "category"
	ReferenceTypeIdChannel          ReferenceTypeId = "channel"
	ReferenceTypeIdCustomer         ReferenceTypeId = "customer"
	ReferenceTypeIdCustomerGroup    ReferenceTypeId = "customer-group"
	ReferenceTypeIdDiscountCode     ReferenceTypeId = "discount-code"
	ReferenceTypeIdExtension        ReferenceTypeId = "extension"
	ReferenceTypeIdInventoryEntry   ReferenceTypeId = "inventory-entry"
	ReferenceTypeIdKeyValueDocument ReferenceTypeId = "key-value-document"
	ReferenceTypeIdOrder            ReferenceTypeId = "order"
	ReferenceTypeIdOrderEdit        ReferenceTypeId = "order-edit"
	ReferenceTypeIdPayment          ReferenceTypeId = "payment"
	ReferenceTypeIdProduct          ReferenceTypeId = "product"
	ReferenceTypeIdProductDiscount  ReferenceTypeId = "product-discount"
	ReferenceTypeIdProductPrice     ReferenceTypeId = "product-price"
	ReferenceTypeIdProductSelection ReferenceTypeId = "product-selection"
	ReferenceTypeIdProductType      ReferenceTypeId = "product-type"
	ReferenceTypeIdReview           ReferenceTypeId = "review"
	ReferenceTypeIdShippingMethod   ReferenceTypeId = "shipping-method"
	ReferenceTypeIdShoppingList     ReferenceTypeId = "shopping-list"
	ReferenceTypeIdState            ReferenceTypeId = "state"
	ReferenceTypeIdStore            ReferenceTypeId = "store"
	ReferenceTypeIdSubscription     ReferenceTypeId = "subscription"
	ReferenceTypeIdTaxCategory      ReferenceTypeId = "tax-category"
	ReferenceTypeIdType             ReferenceTypeId = "type"
	ReferenceTypeIdZone             ReferenceTypeId = "zone"
)

/**
*	Draft type to create a [Reference](ctp:api:type:Reference) or a [KeyReference](ctp:api:type:KeyReference) to a resource. Provide either the `id` or (wherever supported) the `key` of the resource to reference, but depending on the API endpoint the response returns either a Reference or a KeyReference. For example, the field `parent` of a [CategoryDraft](ctp:api:type:CategoryDraft) takes a ResourceIdentifier for its value while the value of the corresponding field of a [Category](ctp:api:type:Category) is a Reference.
*
*	Each resource type has its corresponding ResourceIdentifier, like [ChannelResourceIdentifier](ctp:api:type:ChannelResourceIdentifier).
*
 */
type ResourceIdentifier interface{}

func mapDiscriminatorResourceIdentifier(input interface{}) (ResourceIdentifier, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["typeId"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'typeId'")
		}
	} else {
		return nil, errors.New("invalid data")
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
	case "product-selection":
		obj := ProductSelectionResourceIdentifier{}
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
	case "standalone-price":
		obj := StandalonePriceResourceIdentifier{}
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
	ID string `json:"id"`
	// Base polymorphic read-only Money type which is stored in cent precision or high precision. The actual type is determined by the `type` field.
	Value TypedMoney `json:"value"`
	// Base polymorphic read-only Money type which is stored in cent precision or high precision. The actual type is determined by the `type` field.
	CurrentValue TypedMoney `json:"currentValue"`
	// Two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country *string `json:"country,omitempty"`
	// [Reference](ctp:api:type:Reference) to a [CustomerGroup](ctp:api:type:CustomerGroup).
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
	// [Reference](ctp:api:type:Reference) to a [Channel](ctp:api:type:Channel).
	Channel    *ChannelReference `json:"channel,omitempty"`
	ValidFrom  *time.Time        `json:"validFrom,omitempty"`
	ValidUntil *time.Time        `json:"validUntil,omitempty"`
	Discounted *DiscountedPrice  `json:"discounted,omitempty"`
	// Serves as value of the `custom` field on a resource or data type customized with a [Type](ctp:api:type:Type).
	Custom *CustomFields `json:"custom,omitempty"`
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

/**
*	Base polymorphic read-only Money type which is stored in cent precision or high precision. The actual type is determined by the `type` field.
*
 */
type TypedMoney interface{}

func mapDiscriminatorTypedMoney(input interface{}) (TypedMoney, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("invalid data")
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

/**
*	Object that stores cent amounts in a specific currency.
*
 */
type CentPrecisionMoney struct {
	// Amount in the smallest indivisible unit of a currency, such as:
	//
	// * Cents for EUR and USD, pence for GBP, or centime for CHF (5 CHF is specified as `500`).
	// * The value in the major unit for currencies without minor units, like JPY (5 JPY is specified as `5`).
	CentAmount int `json:"centAmount"`
	// Currency code compliant to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	CurrencyCode string `json:"currencyCode"`
	// The number of default fraction digits for the given currency, like `2` for EUR or `0` for JPY.
	FractionDigits int `json:"fractionDigits"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CentPrecisionMoney) MarshalJSON() ([]byte, error) {
	type Alias CentPrecisionMoney
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "centPrecision", Alias: (*Alias)(&obj)})
}

/**
*	Money object that stores an amount of a fraction of the smallest indivisible unit of the specified currency.
 */
type HighPrecisionMoney struct {
	// Amount in the smallest indivisible unit of a currency, such as:
	//
	// * Cents for EUR and USD, pence for GBP, or centime for CHF (5 CHF is specified as `500`).
	// * The value in the major unit for currencies without minor units, like JPY (5 JPY is specified as `5`).
	CentAmount int `json:"centAmount"`
	// Currency code compliant to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	CurrencyCode string `json:"currencyCode"`
	// Number of digits after the decimal separator, greater than the default number of fraction digits for a currency.
	FractionDigits int `json:"fractionDigits"`
	// Amount in 1 / (10 ^ `fractionDigits`) of a currency.
	PreciseAmount int `json:"preciseAmount"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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
			return nil, errors.New("error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("invalid data")
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
	// Amount in the smallest indivisible unit of a currency, such as:
	//
	// * Cents for EUR and USD, pence for GBP, or centime for CHF (5 CHF is specified as `500`).
	// * The value in the major unit for currencies without minor units, like JPY (5 JPY is specified as `5`).
	CentAmount int `json:"centAmount"`
	// Currency code compliant to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	CurrencyCode string `json:"currencyCode"`
	// This field is optional for cent precision. If provided, it must be equal to the default number of fraction digits for the specified currency.
	FractionDigits *int `json:"fractionDigits,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CentPrecisionMoneyDraft) MarshalJSON() ([]byte, error) {
	type Alias CentPrecisionMoneyDraft
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "centPrecision", Alias: (*Alias)(&obj)})
}

/**
*	Money draft object to store an amount of a fraction of the smallest indivisible unit of the specified currency.
 */
type HighPrecisionMoneyDraft struct {
	// Amount in the smallest indivisible unit of a currency. This field is optional for high precision. If provided, it is checked for validity. Example:
	//
	// A Price of 1.015 USD can be rounded either to 1.01 USD or 1.02 USD. If it lies outside of this range, an error message stating that centAmount must be rounded correctly will be returned.
	//
	// If `centAmount` is not provided, the API calculates the value automatically using the default rounding mode half even.
	CentAmount *int `json:"centAmount,omitempty"`
	// Currency code compliant to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	CurrencyCode string `json:"currencyCode"`
	// Number of fraction digits for a specified high precision money. It must be greater than the default number of fraction digits for the specified currency.
	FractionDigits int `json:"fractionDigits"`
	// Amount in 1 / (10 ^ `fractionDigits`) of a currency.
	PreciseAmount int `json:"preciseAmount"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj HighPrecisionMoneyDraft) MarshalJSON() ([]byte, error) {
	type Alias HighPrecisionMoneyDraft
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "highPrecision", Alias: (*Alias)(&obj)})
}
