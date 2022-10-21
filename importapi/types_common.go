package importapi

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
)

type Asset struct {
	// User-defined identifier for the asset.
	// Asset keys are unique inside their container (a product variant or a category).
	Key     string        `json:"key"`
	Sources []AssetSource `json:"sources"`
	// A localized string is a JSON object where the keys are of [IETF language tag](https://en.wikipedia.org/wiki/IETF_language_tag), and the values the corresponding strings used for that language.
	// ```json
	// {
	//   "de": "Hundefutter",
	//   "en": "dog food"
	// }
	// ```
	Name LocalizedString `json:"name"`
	// A localized string is a JSON object where the keys are of [IETF language tag](https://en.wikipedia.org/wiki/IETF_language_tag), and the values the corresponding strings used for that language.
	// ```json
	// {
	//   "de": "Hundefutter",
	//   "en": "dog food"
	// }
	// ```
	Description *LocalizedString `json:"description,omitempty"`
	Tags        []string         `json:"tags"`
	// The representation to be sent to the server when creating a resource with custom fields.
	Custom *Custom `json:"custom,omitempty"`
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
*	The width and height of the Asset Source.
 */
type AssetDimensions struct {
	// The width of the asset source.
	W int `json:"w"`
	// The height of the asset source.
	H int `json:"h"`
}

/**
*	An AssetSource is a representation of an Asset in a specific format, for example, a video in a certain encoding or an image in a certain resolution.
 */
type AssetSource struct {
	Uri string  `json:"uri"`
	Key *string `json:"key,omitempty"`
	// The width and height of the Asset Source.
	Dimensions  *AssetDimensions `json:"dimensions,omitempty"`
	ContentType *string          `json:"contentType,omitempty"`
}

/**
*	An Image uploaded to commercetools Composable Commerce is stored in a Content Delivery Network and it's available in several pre-defined sizes. If you already have an image stored on an external service, you can save the URL when creating a new product or adding a variant, or you can add it later.
 */
type Image struct {
	// URL of the image in its original size. The URL must be unique within a single variant. It can be used to obtain the image in different sizes.
	Url string `json:"url"`
	// Dimensions of the original image. This can be used by your application, for example, to determine whether the image is large enough to display a zoom view.
	Dimensions AssetDimensions `json:"dimensions"`
	// Custom label that can be used, for example, as an image description.
	Label *string `json:"label,omitempty"`
}

type EnumValue struct {
	Key   string `json:"key"`
	Label string `json:"label"`
}

type LocalizedEnumValue struct {
	Key string `json:"key"`
	// A localized string is a JSON object where the keys are of [IETF language tag](https://en.wikipedia.org/wiki/IETF_language_tag), and the values the corresponding strings used for that language.
	// ```json
	// {
	//   "de": "Hundefutter",
	//   "en": "dog food"
	// }
	// ```
	Label LocalizedString `json:"label"`
}

/**
*	A localized string is a JSON object where the keys are of [IETF language tag](https://en.wikipedia.org/wiki/IETF_language_tag), and the values the corresponding strings used for that language.
*	```json
*	{
*	  "de": "Hundefutter",
*	  "en": "dog food"
*	}
*	```
*
 */
type LocalizedString map[string]string

/**
*	A representation of the resource to import.
*	Import resources are similar to draft types, but they only support key references.
*	In general, import resources are more granular then regular resources.
*	They are optimized for incremental updates and therefore have a slightly different structure.
*
 */
type ImportResource struct {
	// User-defined unique identifier.
	Key string `json:"key"`
}

/**
*	References a resource by key.
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
	case "cart":
		obj := CartKeyReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "cart-discount":
		obj := CartDiscountKeyReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "category":
		obj := CategoryKeyReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "channel":
		obj := ChannelKeyReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "customer":
		obj := CustomerKeyReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "customer-group":
		obj := CustomerGroupKeyReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "discount-code":
		obj := DiscountCodeKeyReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "order":
		obj := OrderKeyReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "payment":
		obj := PaymentKeyReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "price":
		obj := PriceKeyReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "product":
		obj := ProductKeyReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "product-discount":
		obj := ProductDiscountKeyReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "product-type":
		obj := ProductTypeKeyReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "product-variant":
		obj := ProductVariantKeyReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "shipping-method":
		obj := ShippingMethodKeyReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "state":
		obj := StateKeyReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "store":
		obj := StoreKeyReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "tax-category":
		obj := TaxCategoryKeyReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "type":
		obj := TypeKeyReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "key-value-document":
		obj := CustomObjectKeyReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	References a cart by key.
 */
type CartKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartKeyReference) MarshalJSON() ([]byte, error) {
	type Alias CartKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "cart", Alias: (*Alias)(&obj)})
}

/**
*	References a cart discount by key.
 */
type CartDiscountKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountKeyReference) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "cart-discount", Alias: (*Alias)(&obj)})
}

/**
*	References a category by key.
 */
type CategoryKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CategoryKeyReference) MarshalJSON() ([]byte, error) {
	type Alias CategoryKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "category", Alias: (*Alias)(&obj)})
}

/**
*	References a channel by key.
 */
type ChannelKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChannelKeyReference) MarshalJSON() ([]byte, error) {
	type Alias ChannelKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "channel", Alias: (*Alias)(&obj)})
}

/**
*	References a customer by key.
 */
type CustomerKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerKeyReference) MarshalJSON() ([]byte, error) {
	type Alias CustomerKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "customer", Alias: (*Alias)(&obj)})
}

/**
*	References a customer group by key.
 */
type CustomerGroupKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerGroupKeyReference) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "customer-group", Alias: (*Alias)(&obj)})
}

/**
*	References a discount code by key.
 */
type DiscountCodeKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeKeyReference) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "discount-code", Alias: (*Alias)(&obj)})
}

/**
*	References an order by key.
 */
type OrderKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderKeyReference) MarshalJSON() ([]byte, error) {
	type Alias OrderKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "order", Alias: (*Alias)(&obj)})
}

/**
*	References a payment by key.
 */
type PaymentKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentKeyReference) MarshalJSON() ([]byte, error) {
	type Alias PaymentKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "payment", Alias: (*Alias)(&obj)})
}

/**
*	References a price by key.
 */
type PriceKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PriceKeyReference) MarshalJSON() ([]byte, error) {
	type Alias PriceKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "price", Alias: (*Alias)(&obj)})
}

/**
*	References a product by key.
 */
type ProductKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductKeyReference) MarshalJSON() ([]byte, error) {
	type Alias ProductKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "product", Alias: (*Alias)(&obj)})
}

/**
*	References a product discount by key.
 */
type ProductDiscountKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountKeyReference) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "product-discount", Alias: (*Alias)(&obj)})
}

/**
*	References a product type by key.
 */
type ProductTypeKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTypeKeyReference) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "product-type", Alias: (*Alias)(&obj)})
}

/**
*	References a product variant by key.
 */
type ProductVariantKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductVariantKeyReference) MarshalJSON() ([]byte, error) {
	type Alias ProductVariantKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "product-variant", Alias: (*Alias)(&obj)})
}

/**
*	References a shipping method by key.
 */
type ShippingMethodKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShippingMethodKeyReference) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "shipping-method", Alias: (*Alias)(&obj)})
}

/**
*	References a state by key.
 */
type StateKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StateKeyReference) MarshalJSON() ([]byte, error) {
	type Alias StateKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "state", Alias: (*Alias)(&obj)})
}

/**
*	References a store by key.
 */
type StoreKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StoreKeyReference) MarshalJSON() ([]byte, error) {
	type Alias StoreKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "store", Alias: (*Alias)(&obj)})
}

/**
*	References a tax category by key.
 */
type TaxCategoryKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TaxCategoryKeyReference) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "tax-category", Alias: (*Alias)(&obj)})
}

/**
*	References a type by key.
 */
type TypeKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TypeKeyReference) MarshalJSON() ([]byte, error) {
	type Alias TypeKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "type", Alias: (*Alias)(&obj)})
}

/**
*	References a key value document by key.
 */
type CustomObjectKeyReference struct {
	Key       string `json:"key"`
	Container string `json:"container"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomObjectKeyReference) MarshalJSON() ([]byte, error) {
	type Alias CustomObjectKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "key-value-document", Alias: (*Alias)(&obj)})
}

type UnresolvedReferences struct {
	Key string `json:"key"`
	// The type of the referenced resource.
	TypeId ReferenceType `json:"typeId"`
}

type MoneyType string

const (
	MoneyTypeCentPrecision MoneyType = "centPrecision"
	MoneyTypeHighPrecision MoneyType = "highPrecision"
)

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
	case "highPrecision":
		obj := HighPrecisionMoney{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "centPrecision":
		obj := Money{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type HighPrecisionMoney struct {
	FractionDigits *int `json:"fractionDigits,omitempty"`
	CentAmount     int  `json:"centAmount"`
	// The currency code compliant to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	CurrencyCode  string `json:"currencyCode"`
	PreciseAmount int    `json:"preciseAmount"`
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

type Money struct {
	FractionDigits *int `json:"fractionDigits,omitempty"`
	CentAmount     int  `json:"centAmount"`
	// The currency code compliant to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	CurrencyCode string `json:"currencyCode"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj Money) MarshalJSON() ([]byte, error) {
	type Alias Money
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "centPrecision", Alias: (*Alias)(&obj)})
}

type DiscountedPrice struct {
	Value TypedMoney `json:"value"`
	// Reference to a ProductDiscount.
	Discount ProductDiscountKeyReference `json:"discount"`
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

/**
*	A price tier is selected instead of the default price when a certain quantity of the ProductVariant is added to a cart and ordered.
*
 */
type PriceTier struct {
	// The minimum quantity this price tier is valid for.
	MinimumQuantity int `json:"minimumQuantity"`
	// The currency of a price tier is always the same as the currency of the base Price.
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
*	The type of the import resource.
*
 */
type ImportResourceType string

const (
	ImportResourceTypeCategory            ImportResourceType = "category"
	ImportResourceTypeOrder               ImportResourceType = "order"
	ImportResourceTypeOrderPatch          ImportResourceType = "order-patch"
	ImportResourceTypePrice               ImportResourceType = "price"
	ImportResourceTypeProduct             ImportResourceType = "product"
	ImportResourceTypeProductDraft        ImportResourceType = "product-draft"
	ImportResourceTypeProductType         ImportResourceType = "product-type"
	ImportResourceTypeProductVariant      ImportResourceType = "product-variant"
	ImportResourceTypeProductVariantPatch ImportResourceType = "product-variant-patch"
	ImportResourceTypeCustomer            ImportResourceType = "customer"
	ImportResourceTypeInventory           ImportResourceType = "inventory"
	ImportResourceTypeStandalonePrice     ImportResourceType = "standalone-price"
)

/**
*	The type of the referenced resource.
*
 */
type ReferenceType string

const (
	ReferenceTypeCart             ReferenceType = "cart"
	ReferenceTypeCartDiscount     ReferenceType = "cart-discount"
	ReferenceTypeCategory         ReferenceType = "category"
	ReferenceTypeChannel          ReferenceType = "channel"
	ReferenceTypeCustomer         ReferenceType = "customer"
	ReferenceTypeCustomerGroup    ReferenceType = "customer-group"
	ReferenceTypeDiscountCode     ReferenceType = "discount-code"
	ReferenceTypeOrder            ReferenceType = "order"
	ReferenceTypePayment          ReferenceType = "payment"
	ReferenceTypePrice            ReferenceType = "price"
	ReferenceTypeProduct          ReferenceType = "product"
	ReferenceTypeProductDiscount  ReferenceType = "product-discount"
	ReferenceTypeProductType      ReferenceType = "product-type"
	ReferenceTypeProductVariant   ReferenceType = "product-variant"
	ReferenceTypeShippingMethod   ReferenceType = "shipping-method"
	ReferenceTypeState            ReferenceType = "state"
	ReferenceTypeStore            ReferenceType = "store"
	ReferenceTypeTaxCategory      ReferenceType = "tax-category"
	ReferenceTypeType             ReferenceType = "type"
	ReferenceTypeKeyValueDocument ReferenceType = "key-value-document"
)

/**
*	Every [Import Operation](/import-operation) is assigned with one of the following states.
*
 */
type ProcessingState string

const (
	ProcessingStateProcessing           ProcessingState = "processing"
	ProcessingStateValidationFailed     ProcessingState = "validationFailed"
	ProcessingStateUnresolved           ProcessingState = "unresolved"
	ProcessingStateWaitForMasterVariant ProcessingState = "waitForMasterVariant"
	ProcessingStateImported             ProcessingState = "imported"
	ProcessingStateRejected             ProcessingState = "rejected"
)

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

type ProductPriceModeEnum string

const (
	ProductPriceModeEnumEmbedded   ProductPriceModeEnum = "Embedded"
	ProductPriceModeEnumStandalone ProductPriceModeEnum = "Standalone"
)
