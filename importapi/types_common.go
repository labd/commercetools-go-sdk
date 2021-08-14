// Generated file, please do not change!!!
package importapi

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
	Description *LocalizedString `json:"description,omitEmpty"`
	Tags        []string         `json:"tags,omitEmpty"`
	// The representation to be sent to the server when creating a resource with custom fields.
	Custom *Custom `json:"custom,omitEmpty"`
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
	Uri string `json:"uri"`
	Key string `json:"key,omitEmpty"`
	// The width and height of the Asset Source.
	Dimensions  *AssetDimensions `json:"dimensions,omitEmpty"`
	ContentType string           `json:"contentType,omitEmpty"`
}

/**
*	An Image uploaded to the commercetools platform is stored in a Content Delivery Network and it's available in several pre-defined sizes. If you already have an image stored on an external service, you can save the URL when creating a new product or adding a variant, or you can add it later.
 */
type Image struct {
	// URL of the image in its original size. The URL must be unique within a single variant. It can be used to obtain the image in different sizes.
	Url string `json:"url"`
	// Dimensions of the original image. This can be used by your application, for example, to determine whether the image is large enough to display a zoom view.
	Dimensions AssetDimensions `json:"dimensions"`
	// Custom label that can be used, for example, as an image description.
	Label string `json:"label,omitEmpty"`
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

// LocalizedString is something
type LocalizedString map[string]string

/**
*	A representation of the resource to import.
*	Import resources are similar to commercetools draft types, but they only support key references.
*	In general, import resources are more granular then the normal commercetools resource.
*	They are optimized for incremental updates and therefore have a slightly different structure.
*
 */
type ImportResource struct {
	Key string `json:"key"`
}

/**
*	References a resource by its key.
 */
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
	case "cart":
		new := CartKeyReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "cart-discount":
		new := CartDiscountKeyReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "category":
		new := CategoryKeyReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "channel":
		new := ChannelKeyReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "customer":
		new := CustomerKeyReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "customer-group":
		new := CustomerGroupKeyReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "discount-code":
		new := DiscountCodeKeyReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "order":
		new := OrderKeyReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "payment":
		new := PaymentKeyReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "price":
		new := PriceKeyReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "product":
		new := ProductKeyReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "product-discount":
		new := ProductDiscountKeyReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "product-type":
		new := ProductTypeKeyReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "product-variant":
		new := ProductVariantKeyReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "shipping-method":
		new := ShippingMethodKeyReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "state":
		new := StateKeyReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "store":
		new := StoreKeyReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "tax-category":
		new := TaxCategoryKeyReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "type":
		new := TypeKeyReference{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

/**
*	References a cart by its key.
 */
type CartKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value
func (obj CartKeyReference) MarshalJSON() ([]byte, error) {
	type Alias CartKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "cart", Alias: (*Alias)(&obj)})
}

/**
*	References a cart discount by its key.
 */
type CartDiscountKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountKeyReference) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "cart-discount", Alias: (*Alias)(&obj)})
}

/**
*	References a category by its key.
 */
type CategoryKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value
func (obj CategoryKeyReference) MarshalJSON() ([]byte, error) {
	type Alias CategoryKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "category", Alias: (*Alias)(&obj)})
}

/**
*	References a channel by its key.
 */
type ChannelKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value
func (obj ChannelKeyReference) MarshalJSON() ([]byte, error) {
	type Alias ChannelKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "channel", Alias: (*Alias)(&obj)})
}

/**
*	References a customer by its key.
 */
type CustomerKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerKeyReference) MarshalJSON() ([]byte, error) {
	type Alias CustomerKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "customer", Alias: (*Alias)(&obj)})
}

/**
*	References a customer group by its key.
 */
type CustomerGroupKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerGroupKeyReference) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "customer-group", Alias: (*Alias)(&obj)})
}

/**
*	References a discount code by its key.
 */
type DiscountCodeKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value
func (obj DiscountCodeKeyReference) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "discount-code", Alias: (*Alias)(&obj)})
}

/**
*	References an order by its key.
 */
type OrderKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderKeyReference) MarshalJSON() ([]byte, error) {
	type Alias OrderKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "order", Alias: (*Alias)(&obj)})
}

/**
*	References a payment by its key.
 */
type PaymentKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentKeyReference) MarshalJSON() ([]byte, error) {
	type Alias PaymentKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "payment", Alias: (*Alias)(&obj)})
}

/**
*	References a price by its key.
 */
type PriceKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value
func (obj PriceKeyReference) MarshalJSON() ([]byte, error) {
	type Alias PriceKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "price", Alias: (*Alias)(&obj)})
}

/**
*	References a product by its key.
 */
type ProductKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductKeyReference) MarshalJSON() ([]byte, error) {
	type Alias ProductKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "product", Alias: (*Alias)(&obj)})
}

/**
*	References a product discount by its key.
 */
type ProductDiscountKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductDiscountKeyReference) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "product-discount", Alias: (*Alias)(&obj)})
}

/**
*	References a product type by its key.
 */
type ProductTypeKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductTypeKeyReference) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "product-type", Alias: (*Alias)(&obj)})
}

/**
*	References a product variant by its key.
 */
type ProductVariantKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductVariantKeyReference) MarshalJSON() ([]byte, error) {
	type Alias ProductVariantKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "product-variant", Alias: (*Alias)(&obj)})
}

/**
*	References a shipping method by its key.
 */
type ShippingMethodKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodKeyReference) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "shipping-method", Alias: (*Alias)(&obj)})
}

/**
*	References a state by its key.
 */
type StateKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value
func (obj StateKeyReference) MarshalJSON() ([]byte, error) {
	type Alias StateKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "state", Alias: (*Alias)(&obj)})
}

/**
*	References a store by its key.
 */
type StoreKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value
func (obj StoreKeyReference) MarshalJSON() ([]byte, error) {
	type Alias StoreKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "store", Alias: (*Alias)(&obj)})
}

/**
*	References a tax category by its key.
 */
type TaxCategoryKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value
func (obj TaxCategoryKeyReference) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "tax-category", Alias: (*Alias)(&obj)})
}

/**
*	References a type by its key.
 */
type TypeKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeKeyReference) MarshalJSON() ([]byte, error) {
	type Alias TypeKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "type", Alias: (*Alias)(&obj)})
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
			return nil, errors.New("Error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}

	switch discriminator {
	case "highPrecision":
		new := HighPrecisionMoney{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "centPrecision":
		new := Money{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type HighPrecisionMoney struct {
	FractionDigits int `json:"fractionDigits,omitEmpty"`
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

type Money struct {
	FractionDigits int `json:"fractionDigits,omitEmpty"`
	CentAmount     int `json:"centAmount"`
	// The currency code compliant to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	CurrencyCode string `json:"currencyCode"`
}

// MarshalJSON override to set the discriminator value
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
)

/**
*	The type of the referenced resource.
*
 */
type ReferenceType string

const (
	ReferenceTypeCart            ReferenceType = "cart"
	ReferenceTypeCartDiscount    ReferenceType = "cart-discount"
	ReferenceTypeCategory        ReferenceType = "category"
	ReferenceTypeChannel         ReferenceType = "channel"
	ReferenceTypeCustomer        ReferenceType = "customer"
	ReferenceTypeCustomerGroup   ReferenceType = "customer-group"
	ReferenceTypeDiscountCode    ReferenceType = "discount-code"
	ReferenceTypeOrder           ReferenceType = "order"
	ReferenceTypePayment         ReferenceType = "payment"
	ReferenceTypePrice           ReferenceType = "price"
	ReferenceTypeProduct         ReferenceType = "product"
	ReferenceTypeProductDiscount ReferenceType = "product-discount"
	ReferenceTypeProductType     ReferenceType = "product-type"
	ReferenceTypeProductVariant  ReferenceType = "product-variant"
	ReferenceTypeShippingMethod  ReferenceType = "shipping-method"
	ReferenceTypeState           ReferenceType = "state"
	ReferenceTypeStore           ReferenceType = "store"
	ReferenceTypeTaxCategory     ReferenceType = "tax-category"
	ReferenceTypeType            ReferenceType = "type"
)

/**
*	Represents the status of a resource under an import process. Every resource has the initial state `Unresolved`.
*
 */
type ProcessingState string

const (
	ProcessingStateValidationFailed     ProcessingState = "ValidationFailed"
	ProcessingStateUnresolved           ProcessingState = "Unresolved"
	ProcessingStateWaitForMasterVariant ProcessingState = "WaitForMasterVariant"
	ProcessingStateImported             ProcessingState = "Imported"
	ProcessingStateRejected             ProcessingState = "Rejected"
)

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
