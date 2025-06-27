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
	// Name of the Asset.
	Name LocalizedString `json:"name"`
	// Description of the Asset.
	Description *LocalizedString `json:"description,omitempty"`
	// Keywords for categorizing and organizing Assets.
	Tags []string `json:"tags"`
	// Custom Fields defined for the Asset.
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
	case "key-value-document":
		obj := CustomObjectKeyReference{}
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
	}
	return nil, nil
}

/**
*	Used by the [Import API](/import-export/overview) to identify a Cart
 */
type CartKeyReference struct {
	// User-defined unique identifier of the referenced Cart.
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
*	Used by the [Import API](/import-export/overview) to identify a CartDiscount.
 */
type CartDiscountKeyReference struct {
	// User-defined unique identifier of the referenced CartDiscount.
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
*	Used by the [Import API](/import-export/overview) to identify a Category.
 */
type CategoryKeyReference struct {
	// User-defined unique identifier of the referenced Category.
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
*	Used by the [Import API](/import-export/overview) to identify a Channel.
 */
type ChannelKeyReference struct {
	// User-defined unique identifier of the referenced Channel.
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
*	Used by the [Import API](/import-export/overview) to identify a Customer.
 */
type CustomerKeyReference struct {
	// User-defined unique identifier of the referenced Customer.
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
*	Used by the [Import API](/import-export/overview) to identify a CustomerGroup.
 */
type CustomerGroupKeyReference struct {
	// User-defined unique identifier of the referenced CustomerGroup.
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
*	Used by the [Import API](/import-export/overview) to identify a CustomObject.
 */
type CustomObjectKeyReference struct {
	// User-defined unique identifier of the referenced CustomObject.
	Key string `json:"key"`
	// The `container` of the referenced CustomObject.
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

/**
*	Used by the [Import API](/import-export/overview) to identify a DiscountCode.
 */
type DiscountCodeKeyReference struct {
	// User-defined unique identifier of the referenced DiscountCode.
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
	// User-defined unique identifier of the referenced resource.
	// If the referenced resource does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced resource is created.
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
*	Used by the [Import API](/import-export/overview) to identify a Payment.
 */
type PaymentKeyReference struct {
	// User-defined unique identifier of the referenced Payment.
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
*	Used by the [Import API](/import-export/overview) to identify an Embedded Price.
 */
type PriceKeyReference struct {
	// User-defined unique identifier of the referenced Embedded Price.
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
*	Used by the [Import API](/import-export/overview) to identify a Product.
 */
type ProductKeyReference struct {
	// User-defined unique identifier of the referenced Product.
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
*	Used by the [Import API](/import-export/overview) to identify a ProductDiscount.
 */
type ProductDiscountKeyReference struct {
	// User-defined unique identifier of the referenced ProductDiscount.
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
*	Used by the [Import API](/import-export/overview) to identify a ProductType.
 */
type ProductTypeKeyReference struct {
	// User-defined unique identifier of the referenced ProductType.
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
*	Used by the [Import API](/import-export/overview) to identify a ProductVariant.
 */
type ProductVariantKeyReference struct {
	// User-defined unique identifier of the referenced ProductVariant.
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
*	Used by the [Import API](/import-export/overview) to identify a ShippingMethod.
 */
type ShippingMethodKeyReference struct {
	// User-defined unique identifier of the referenced ShippingMethod.
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
*	Used by the [Import API](/import-export/overview) to identify a State.
 */
type StateKeyReference struct {
	// User-defined unique identifier of the referenced State.
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
*	Used by the [Import API](/import-export/overview) to identify a Store.
 */
type StoreKeyReference struct {
	// User-defined unique identifier of the referenced Store.
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
*	Used by the [Import API](/import-export/overview) to identify a TaxCategory.
 */
type TaxCategoryKeyReference struct {
	// User-defined unique identifier of the referenced TaxCategory.
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
*	Used by the [Import API](/import-export/overview) to identify a Type.
 */
type TypeKeyReference struct {
	// User-defined unique identifier of the referenced Type.
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
*	Contains a reference to a resource which does not exist. For example, if a Category is imported with a parent Category that does not exist, the reference to the parent Category is an unresolved reference.
 */
type UnresolvedReferences struct {
	// `key` of the unresolved resource.
	Key string `json:"key"`
	// Type of the unresolved resource.
	TypeId ReferenceType `json:"typeId"`
}

/**
*	The type of money.
*	The `centPrecision` type is used for currencies with minor units, such as EUR and USD.
*	The `highPrecision` type is used for currencies without minor units, such as JPY.
*
 */
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
	// The number of fraction digits of the money value.
	// This is used to determine how many digits are after the decimal point.
	// For example, for EUR and USD, this is `2`, and for JPY, this is `0`.
	FractionDigits *int `json:"fractionDigits,omitempty"`
	// Amount in the smallest indivisible unit of a currency, such as:
	//
	// * Cents for EUR and USD, pence for GBP, or centime for CHF (5 CHF is specified as `500`).
	// * The value in the major unit for currencies without minor units, like JPY (5 JPY is specified as `5`).
	CentAmount int `json:"centAmount"`
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
	// The number of fraction digits of the money value.
	// This is used to determine how many digits are after the decimal point.
	// For example, for EUR and USD, this is `2`, and for JPY, this is `0`.
	FractionDigits *int `json:"fractionDigits,omitempty"`
	// Amount in the smallest indivisible unit of a currency, such as:
	//
	// * Cents for EUR and USD, pence for GBP, or centime for CHF (5 CHF is specified as `500`).
	// * The value in the major unit for currencies without minor units, like JPY (5 JPY is specified as `5`).
	CentAmount int `json:"centAmount"`
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
	// Money value of the discounted price.
	Value TypedMoney `json:"value"`
	// Reference to a ProductDiscount. If the referenced [ProductDiscount](ctp:api:type:ProductDiscount) does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced ProductDiscount is created.
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
*	The resource types that can be imported.
*
 */
type ImportResourceType string

const (
	ImportResourceTypeCategory            ImportResourceType = "category"
	ImportResourceTypeCustomer            ImportResourceType = "customer"
	ImportResourceTypeDiscountCode        ImportResourceType = "discount-code"
	ImportResourceTypeInventory           ImportResourceType = "inventory"
	ImportResourceTypeOrder               ImportResourceType = "order"
	ImportResourceTypeOrderPatch          ImportResourceType = "order-patch"
	ImportResourceTypePrice               ImportResourceType = "price"
	ImportResourceTypeProduct             ImportResourceType = "product"
	ImportResourceTypeProductDraft        ImportResourceType = "product-draft"
	ImportResourceTypeProductSelection    ImportResourceType = "product-selection"
	ImportResourceTypeProductType         ImportResourceType = "product-type"
	ImportResourceTypeProductVariant      ImportResourceType = "product-variant"
	ImportResourceTypeProductVariantPatch ImportResourceType = "product-variant-patch"
	ImportResourceTypeStandalonePrice     ImportResourceType = "standalone-price"
	ImportResourceTypeType                ImportResourceType = "type"
)

/**
*	Type of referenced resource.
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
	ReferenceTypeKeyValueDocument ReferenceType = "key-value-document"
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
)

/**
*	Every [Import Operation](ctp:import:type:ImportOperation) is assigned one of the following states.
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
	ProcessingStateCanceled             ProcessingState = "canceled"
)

type Address struct {
	// Unique identifier of the Address.
	//
	// It is not recommended to set it manually since the API overwrites this ID when creating an Address for a [Customer](ctp:api:type:Customer).
	// Use `key` instead and omit this field from the request to let the API generate the ID for the Address.
	ID *string `json:"id,omitempty"`
	// User-defined identifier of the Address that must be unique when multiple addresses are referenced in [BusinessUnits](ctp:api:type:BusinessUnit), [Customers](ctp:api:type:Customer), and `itemShippingAddresses` (LineItem-specific addresses) of a [Cart](ctp:api:type:Cart), [Order](ctp:api:type:Order), [QuoteRequest](ctp:api:type:QuoteRequest), or [Quote](ctp:api:type:Quote).
	Key *string `json:"key,omitempty"`
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
	// Name of the country.
	Country string `json:"country"`
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
	Custom *Custom `json:"custom,omitempty"`
}

type ProductPriceModeEnum string

const (
	ProductPriceModeEnumEmbedded   ProductPriceModeEnum = "Embedded"
	ProductPriceModeEnumStandalone ProductPriceModeEnum = "Standalone"
)
