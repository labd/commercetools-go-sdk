package history

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

/**
*	JSON object where the keys are of type [Locale](ctp:api:type:Locale), and the values are the strings used for the corresponding language.
*
 */
type LocalizedString map[string]string
type Asset struct {
	// Unique identifier of the Asset. Not required when importing Assets using the [Import API](/import-export/import-resources).
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
	// User-defined identifier of the Asset. It is unique per [Category](ctp:api:type:Category) or [ProductVariant](ctp:api:type:ProductVariant).
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

/**
*	Representation of an [Asset](#asset) in a specific format, for example a video in a certain encoding, or an image in a certain resolution.
*
 */
type AssetSource struct {
	// URI of the AssetSource.
	Uri string `json:"uri"`
	// User-defined identifier of the AssetSource. Must be unique per [Asset](ctp:api:type:Asset).
	Key *string `json:"key,omitempty"`
	// Width and height of the AssetSource.
	Dimensions *AssetDimensions `json:"dimensions,omitempty"`
	// Indicates the type of content, for example `application/pdf`.
	ContentType *string `json:"contentType,omitempty"`
}

type Associate struct {
	// Roles assigned to the Associate within a Business Unit.
	AssociateRoleAssignments []AssociateRoleAssignment `json:"associateRoleAssignments"`
	// Deprecated type. Use `associateRoleAssignments` instead.
	Roles []AssociateRoleDeprecated `json:"roles"`
	// The [Customer](ctp:api:type:Customer) that acts as an Associate in the Business Unit.
	Customer CustomerReference `json:"customer"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj Associate) MarshalJSON() ([]byte, error) {
	type Alias Associate
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

	if raw["roles"] == nil {
		delete(raw, "roles")
	}

	return json.Marshal(raw)

}

type AssociateRoleAssignment struct {
	// Role the Associate holds within a Business Unit.
	AssociateRole AssociateRoleKeyReference `json:"associateRole"`
	// Determines whether the AssociateRoleAssignment can be inherited by child Business Units.
	Inheritance AssociateRoleInheritanceMode `json:"inheritance"`
}

/**
*	Determines whether an [AssociateRoleAssignment](ctp:api:type:AssociateRoleAssignment) can be inherited by child Business Units.
*
 */
type AssociateRoleInheritanceMode string

const (
	AssociateRoleInheritanceModeEnabled  AssociateRoleInheritanceMode = "Enabled"
	AssociateRoleInheritanceModeDisabled AssociateRoleInheritanceMode = "Disabled"
)

/**
*	Specifies how an Attribute (or a set of Attributes) should be validated across all variants of a Product:
*
 */
type AttributeConstraintEnum string

const (
	AttributeConstraintEnumNone              AttributeConstraintEnum = "None"
	AttributeConstraintEnumUnique            AttributeConstraintEnum = "Unique"
	AttributeConstraintEnumCombinationUnique AttributeConstraintEnum = "CombinationUnique"
	AttributeConstraintEnumSameForAll        AttributeConstraintEnum = "SameForAll"
)

type AttributeDefinition struct {
	// Describes the Type of the Attribute.
	Type AttributeType `json:"type"`
	// User-defined name of the Attribute that is unique within the [Project](ctp:api:type:Project).
	Name string `json:"name"`
	// Human-readable label for the Attribute.
	Label LocalizedString `json:"label"`
	// If `true`, the Attribute must have a value on a [ProductVariant](ctp:api:type:ProductVariant).
	IsRequired bool `json:"isRequired"`
	// Specifies whether the Attribute is defined at the Product or Variant level.
	Level AttributeLevelEnum `json:"level"`
	// Specifies how Attributes are validated across all variants of a Product.
	AttributeConstraint AttributeConstraintEnum `json:"attributeConstraint"`
	// Provides additional Attribute information to aid content managers configure Product details.
	InputTip *LocalizedString `json:"inputTip,omitempty"`
	// Provides a visual representation directive for values of this Attribute (only relevant for [AttributeTextType](ctp:api:type:AttributeTextType) and [AttributeLocalizableTextType](ctp:api:type:AttributeLocalizableTextType)).
	InputHint TextInputHint `json:"inputHint"`
	// If `true`, the Attribute's values are available in the [Product Search](/../api/projects/product-search) or the [Product Projection Search](/../api/projects/product-projection-search) API for use in full-text search queries, filters, and facets.
	// However, if an Attribute's `level` is set as `Product`, then Product Projection Search does **not support** the Attribute.
	//
	// The exact features that are available with this flag depend on the specific [AttributeType](ctp:api:type:AttributeType).
	// The maximum size of a searchable field is **restricted** by the [Field content size limit](/../api/limits#field-content-size).
	// This constraint is enforced at both [Product creation](ctp:api:endpoint:/{projectKey}/products:POST) and [Product update](/../api/projects/products#update-product).
	// If the length of the input exceeds the maximum size, an [InvalidField](ctp:api:type:InvalidFieldError) error is returned.
	IsSearchable bool `json:"isSearchable"`
}

/**
*	A localized enum value must be unique within the enum, else a [DuplicateEnumValues](ctp:api:type:DuplicateEnumValuesError) error is returned.
*
 */
type AttributeLocalizedEnumValue struct {
	// Key of the value used as a programmatic identifier, for example in facets & filters.
	Key string `json:"key"`
	// Descriptive, localized label of the value.
	Label LocalizedString `json:"label"`
}

/**
*	A plain enum value must be unique within the enum, else a [DuplicateEnumValues](ctp:api:type:DuplicateEnumValuesError) error is returned.
*
 */
type AttributePlainEnumValue struct {
	// Key of the value used as a programmatic identifier, for example in facets & filters.
	Key string `json:"key"`
	// Descriptive label of the value.
	Label string `json:"label"`
}

/**
*	Umbrella type for specific attribute types discriminated by property `name`.
 */
type AttributeType struct {
	Name string `json:"name"`
}

type AuthenticationMode string

const (
	AuthenticationModePassword     AuthenticationMode = "Password"
	AuthenticationModeExternalAuth AuthenticationMode = "ExternalAuth"
)

/**
*	Determines whether a Business Unit can inherit Associates from a parent.
*
 */
type BusinessUnitAssociateMode string

const (
	BusinessUnitAssociateModeExplicit              BusinessUnitAssociateMode = "Explicit"
	BusinessUnitAssociateModeExplicitAndFromParent BusinessUnitAssociateMode = "ExplicitAndFromParent"
)

/**
*	Indicates whether the Business Unit can be edited and used in [Carts](ctp:api:type:Cart), [Orders](ctp:api:type:Order), [Quote Requests](ctp:api:type:QuoteRequest), or [Quotes](ctp:api:type:Quote).
*
 */
type BusinessUnitStatus string

const (
	BusinessUnitStatusActive   BusinessUnitStatus = "Active"
	BusinessUnitStatusInactive BusinessUnitStatus = "Inactive"
)

/**
*	Defines whether the Stores of the Business Unit are set directly on the Business Unit or are inherited from its parent unit.
*
 */
type BusinessUnitStoreMode string

const (
	BusinessUnitStoreModeExplicit   BusinessUnitStoreMode = "Explicit"
	BusinessUnitStoreModeFromParent BusinessUnitStoreMode = "FromParent"
)

/**
*	JSON object where the keys are [Category](ctp:api:type:Category) `id`, and the values are order hint values: strings representing a number between `0` and `1`, but not ending in `0`. Order hints allow controlling the order of Products and how they appear in Categories. Products without order hints have an order score below `0`. Order hints are not unique. If a subset of Products have the same value for order hint in a specific category, the behavior is undetermined.
 */
type CategoryOrderHints map[string]string

/**
*	Describes the purpose and type of the Channel. A Channel can have one or more roles.
*
 */
type ChannelRoleEnum string

const (
	ChannelRoleEnumInventorySupply     ChannelRoleEnum = "InventorySupply"
	ChannelRoleEnumProductDistribution ChannelRoleEnum = "ProductDistribution"
	ChannelRoleEnumOrderExport         ChannelRoleEnum = "OrderExport"
	ChannelRoleEnumOrderImport         ChannelRoleEnum = "OrderImport"
	ChannelRoleEnumPrimary             ChannelRoleEnum = "Primary"
)

/**
*	Serves as value of the `custom` field on a resource or data type customized with a [Type](ctp:api:type:Type).
*
 */
type CustomFields struct {
	// Reference to the [Type](ctp:api:type:Type) that holds the [FieldDefinitions](ctp:api:type:FieldDefinition) for the Custom Fields.
	Type TypeReference `json:"type"`
	// Object containing the Custom Fields of the [customized resource or data type](/../api/projects/types#resourcetypeid).
	Fields FieldContainer `json:"fields"`
}

/**
*	A generic item that can be added to the Cart but is not bound to a Product that can be used for discounts (negative money), vouchers, complex cart rules, additional services, or fees.
*	You control the lifecycle of this item.
*
 */
type CustomLineItem struct {
	// Unique identifier of the Custom Line Item.
	ID string `json:"id"`
	// User-defined unique identifier of the Custom Line Item.
	Key *string `json:"key,omitempty"`
	// Name of the Custom Line Item.
	Name LocalizedString `json:"name"`
	// Money value of the Custom Line Item.
	Money TypedMoney `json:"money"`
	// Automatically set after the `taxRate` is set.
	TaxedPrice *TaxedItemPrice `json:"taxedPrice,omitempty"`
	// Total taxed prices based on the quantity of the Custom Line Item assigned to each [Shipping Method](ctp:api:type:ShippingMethod). Only applicable for Carts with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	// Automatically set after `perMethodTaxRate` is set.
	TaxedPricePortions []MethodTaxedPrice `json:"taxedPricePortions"`
	// Total price of the Custom Line Item (`money` multiplied by `quantity`).
	// If the Custom Line Item is discounted, the total price is `discountedPricePerQuantity` multiplied by `quantity`.
	//
	// Includes taxes if the [TaxRate](ctp:api:type:TaxRate) `includedInPrice` is `true`.
	TotalPrice CentPrecisionMoney `json:"totalPrice"`
	// User-defined identifier used in a deep-link URL for the Custom Line Item.
	// It matches the pattern `[a-zA-Z0-9_-]{2,256}`.
	Slug string `json:"slug"`
	// Number of Custom Line Items in the [Cart](ctp:api:type:Cart) or [Order](ctp:api:type:Order).
	Quantity int `json:"quantity"`
	// Tracks specific quantities of the Custom Line Item within a given State. When a Custom Line Item is added to a Cart, its full quantity is set to the built-in "Initial" state. State transitions for Custom Line Items are managed on the [Order](ctp:api:type:Order).
	State []ItemState `json:"state"`
	// Used to select a Tax Rate when a Cart has the `Platform` [TaxMode](ctp:api:type:TaxMode).
	TaxCategory *TaxCategoryReference `json:"taxCategory,omitempty"`
	// - For a Cart with `Platform` [TaxMode](ctp:api:type:TaxMode), the `taxRate` of Custom Line Items is set automatically once a shipping address is set. The rate is based on the [TaxCategory](ctp:api:type:TaxCategory) that applies for the shipping address.
	// - For a Cart with `External` TaxMode, the `taxRate` of Custom Line Items can be set using [ExternalTaxRateDraft](ctp:api:type:ExternalTaxRateDraft).
	TaxRate *TaxRate `json:"taxRate,omitempty"`
	// Tax Rate per Shipping Method for a Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode). For a Cart with `Platform` [TaxMode](ctp:api:type:TaxMode) it is automatically set after the [Shipping Method is added](ctp:api:type:CartAddShippingMethodAction).
	// For a Cart with `External` [TaxMode](ctp:api:type:TaxMode), the Tax Rate must be set with [ExternalTaxRateDraft](ctp:api:type:ExternalTaxRateDraft).
	PerMethodTaxRate []MethodTaxRate `json:"perMethodTaxRate"`
	// Discounted price of a single quantity of the Custom Line Item.
	DiscountedPricePerQuantity []DiscountedLineItemPriceForQuantity `json:"discountedPricePerQuantity"`
	// Custom Fields of the Custom Line Item.
	Custom *CustomFields `json:"custom,omitempty"`
	// Container for Custom Line Item-specific addresses.
	ShippingDetails *ItemShippingDetails `json:"shippingDetails,omitempty"`
	// Indicates whether Cart Discounts with a matching [CartDiscountCustomLineItemsTarget](ctp:api:type:CartDiscountCustomLineItemsTarget), [MultiBuyCustomLineItemsTarget](ctp:api:type:MultiBuyCustomLineItemsTarget), or [CartDiscountPatternTarget](ctp:api:type:CartDiscountPatternTarget) are applied to the Custom Line Item.
	PriceMode CustomLineItemPriceMode `json:"priceMode"`
	// Recurring Order and frequency data.
	RecurrenceInfo *CustomLineItemRecurrenceInfo `json:"recurrenceInfo,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CustomLineItem) UnmarshalJSON(data []byte) error {
	type Alias CustomLineItem
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Money != nil {
		var err error
		obj.Money, err = mapDiscriminatorTypedMoney(obj.Money)
		if err != nil {
			return err
		}
	}

	return nil
}

/**
*	Contains information on how items are shipped to Customers, for example, a delivery note.
*
 */
type Delivery struct {
	// Unique identifier of the Delivery.
	ID string `json:"id"`
	// User-defined unique identifier of the Delivery.
	Key *string `json:"key,omitempty"`
	// Date and time (UTC) the Delivery was created.
	CreatedAt time.Time `json:"createdAt"`
	// Line Items or Custom Line Items that are delivered.
	Items []DeliveryItem `json:"items"`
	// Information regarding the appearance, content, and shipment of a Parcel.
	Parcels []Parcel `json:"parcels"`
	// Address to which Parcels are delivered.
	Address *Address `json:"address,omitempty"`
	// Custom Fields of the Delivery.
	Custom *CustomFields `json:"custom,omitempty"`
}

type DeliveryItem struct {
	// `id` of the [LineItem](ctp:api:type:LineItem) or [CustomLineItem](ctp:api:type:CustomLineItem) delivered.
	ID string `json:"id"`
	// Number of Line Items or Custom Line Items delivered.
	Quantity int `json:"quantity"`
}

/**
*	Represents a [CartDiscount](ctp:api:type:CartDiscount) that is only associated with a single Cart or Order.
*
 */
type DirectDiscount struct {
	// Unique identifier of the Direct Discount.
	ID string `json:"id"`
	// Effect of the Discount on the Cart.
	Value CartDiscountValue `json:"value"`
	// Segment of the Cart that is discounted.
	//
	// Empty when the `value` is set to `giftLineItem`.
	Target *CartDiscountTarget `json:"target,omitempty"`
}

type DiscountCodeInfo struct {
	// Discount Code associated with the Cart or Order.
	DiscountCode DiscountCodeReference `json:"discountCode"`
	// Indicates the state of the Discount Code applied to the Cart or Order.
	State DiscountCodeState `json:"state"`
}

/**
*	Indicates the state of a Discount Code in a Cart.
*
*	If an Order is created from a Cart with a state other than `MatchesCart` or `ApplicationStoppedByGroupBestDeal`, a [DiscountCodeNonApplicable](ctp:api:type:DiscountCodeNonApplicableError) error is returned.
*
*	For Orders created from a Cart with a `ApplicationStoppedByGroupBestDeal` state, the discount code is not applied.
*
 */
type DiscountCodeState string

const (
	DiscountCodeStateNotActive                            DiscountCodeState = "NotActive"
	DiscountCodeStateNotValid                             DiscountCodeState = "NotValid"
	DiscountCodeStateDoesNotMatchCart                     DiscountCodeState = "DoesNotMatchCart"
	DiscountCodeStateMatchesCart                          DiscountCodeState = "MatchesCart"
	DiscountCodeStateMaxApplicationReached                DiscountCodeState = "MaxApplicationReached"
	DiscountCodeStateApplicationStoppedByPreviousDiscount DiscountCodeState = "ApplicationStoppedByPreviousDiscount"
	DiscountCodeStateApplicationStoppedByGroupBestDeal    DiscountCodeState = "ApplicationStoppedByGroupBestDeal"
)

type DiscountedLineItemPortion struct {
	// A [CartDiscountReference](ctp:api:type:CartDiscountReference) or [DirectDiscountReference](ctp:api:type:DirectDiscountReference) of the applicable discount on the Line Item.
	Discount Reference `json:"discount"`
	// Money value of the applicable discount.
	DiscountedAmount TypedMoney `json:"discountedAmount"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DiscountedLineItemPortion) UnmarshalJSON(data []byte) error {
	type Alias DiscountedLineItemPortion
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Discount != nil {
		var err error
		obj.Discount, err = mapDiscriminatorReference(obj.Discount)
		if err != nil {
			return err
		}
	}
	if obj.DiscountedAmount != nil {
		var err error
		obj.DiscountedAmount, err = mapDiscriminatorTypedMoney(obj.DiscountedAmount)
		if err != nil {
			return err
		}
	}

	return nil
}

type DiscountedLineItemPrice struct {
	// Money value of the discounted Line Item or Custom Line Item.
	Value TypedMoney `json:"value"`
	// Discount applicable on the Line Item or Custom Line Item.
	IncludedDiscounts []DiscountedLineItemPortion `json:"includedDiscounts"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DiscountedLineItemPrice) UnmarshalJSON(data []byte) error {
	type Alias DiscountedLineItemPrice
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

type DiscountedLineItemPriceForQuantity struct {
	// Number of Line Items or Custom Line Items in the Cart.
	Quantity int `json:"quantity"`
	// Discounted price of the Line Item or Custom Line Item.
	DiscountedPrice DiscountedLineItemPrice `json:"discountedPrice"`
}

type DiscountTypeCombination interface{}

func mapDiscriminatorDiscountTypeCombination(input interface{}) (DiscountTypeCombination, error) {
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
	case "BestDeal":
		obj := BestDeal{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "Stacking":
		obj := Stacking{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Indicates if a Product Discount or Cart Discount offers the best deal for a Cart or Order.
*
 */
type BestDeal struct {
	// Discount type that offers the best deal; the value can be `ProductDiscount` or `CartDiscount`.
	ChosenDiscountType string `json:"chosenDiscountType"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BestDeal) MarshalJSON() ([]byte, error) {
	type Alias BestDeal
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BestDeal", Alias: (*Alias)(&obj)})
}

/**
*	Indicates both Product Discounts and Cart Discounts apply to a Cart and Order.
*
 */
type Stacking struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj Stacking) MarshalJSON() ([]byte, error) {
	type Alias Stacking
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "Stacking", Alias: (*Alias)(&obj)})
}

/**
*	Defines a [Custom Field](/../api/projects/custom-fields) and its meta-information.
*	This FieldDefinition is similar to an [AttributeDefinition](ctp:api:type:AttributeDefinition) of [Product Types](/../api/projects/productTypes).
*
 */
type FieldDefinition struct {
	// Data type of the Custom Field to define.
	Type FieldType `json:"type"`
	// Name of the Custom Field to define.
	// Must be unique for a given [ResourceTypeId](ctp:api:type:ResourceTypeId).
	// In case there is a FieldDefinition with the same `name` in another [Type](ctp:api:type:Type), both FieldDefinitions must have the same `type`.
	Name string `json:"name"`
	// A human-readable label for the field.
	Label LocalizedString `json:"label"`
	// Defines whether the field is required to have a value.
	Required bool `json:"required"`
	// Defines the visual representation of the field in user interfaces like the Merchant Center.
	// It is only relevant for string-based [FieldTypes](ctp:api:type:FieldType) like [CustomFieldStringType](ctp:api:type:CustomFieldStringType) and [CustomFieldLocalizedStringType](ctp:api:type:CustomFieldLocalizedStringType).
	InputHint *TypeTextInputHint `json:"inputHint,omitempty"`
}

type FieldType struct {
	Name string `json:"name"`
}

type Image struct {
	// URL of the image in its original size that must be unique within a single [ProductVariant](ctp:api:type:ProductVariant).
	Url string `json:"url"`
	// Dimensions of the original image.
	Dimensions ImageDimensions `json:"dimensions"`
	// Custom label for the image.
	Label *string `json:"label,omitempty"`
}

type ImageDimensions struct {
	// Width of the image.
	W int `json:"w"`
	// Height of the image.
	H int `json:"h"`
}

type InheritedAssociate struct {
	// Inherited roles of the Associate within a Business Unit.
	AssociateRoleAssignments []InheritedAssociateRoleAssignment `json:"associateRoleAssignments"`
	// The [Customer](ctp:api:type:Customer) that acts as an Associate in the Business Unit.
	Customer CustomerReference `json:"customer"`
}

type InheritedAssociateRoleAssignment struct {
	// Inherited role the Associate holds within a Business Unit.
	AssociateRole AssociateRoleKeyReference `json:"associateRole"`
	// Reference to the parent Business Unit where the assignment is defined explicitly.
	Source BusinessUnitKeyReference `json:"source"`
}

type ItemShippingDetails struct {
	// Holds information on the quantity of Line Items or Custom Line Items and the address it is shipped.
	Targets []ItemShippingTarget `json:"targets"`
	// - `true` if the quantity of Line Items or Custom Line Items is equal to the sum of sub-quantities defined in `targets`.
	// - `false` if the quantity of Line Items or Custom Line Items is not equal to the sum of sub-quantities defined in `targets`.
	//   Ordering a Cart when the value is `false` returns an [InvalidItemShippingDetails](ctp:api:type:InvalidItemShippingDetailsError) error.
	Valid bool `json:"valid"`
}

/**
*	Determines the address (as a reference to an address in `itemShippingAddresses`) and the quantity shipped to the address.
*
*	If multiple shipping addresses are present for a Line Item or Custom Line Item, sub-quantities must be specified.
*	An array of addresses and sub-quantities is stored per Line Item or Custom Line Item.
*
 */
type ItemShippingTarget struct {
	// Key of the address in the [Cart](ctp:api:type:Cart) `itemShippingAddresses`.
	// Duplicate address keys are not allowed.
	AddressKey string `json:"addressKey"`
	// Quantity of Line Items or Custom Line Items shipped to the address with the specified `addressKey`.
	//
	// If a quantity is updated to `0` when defining [ItemShippingDetailsDraft](ctp:api:type:ItemShippingDetailsDraft), the `targets` are removed from a Line Item or Custom Line Item in the resulting [ItemShippingDetails](ctp:api:type:ItemShippingDetails).
	Quantity int `json:"quantity"`
	// User-defined unique identifier of the Shipping Method in a Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	//
	// It connects Line Item or Custom Line Item quantities with individual Shipping Methods.
	ShippingMethodKey *string `json:"shippingMethodKey,omitempty"`
}

type ItemState struct {
	// Number of Line Items or Custom Line Items in this State.
	Quantity int `json:"quantity"`
	// State of the Line Items or Custom Line Items in a custom workflow.
	State StateReference `json:"state"`
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
	case "associate-role":
		obj := AssociateRoleKeyReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "business-unit":
		obj := BusinessUnitKeyReference{}
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
	}
	return nil, nil
}

/**
*	The representation of a [Line Item](/../api/carts-orders-overview#line-items) in a [Cart](ctp:api:type:Cart) or in an [Order](ctp:api:type:Order).
*
 */
type LineItem struct {
	// Unique identifier of the LineItem.
	ID string `json:"id"`
	// User-defined unique identifier of the LineItem.
	Key *string `json:"key,omitempty"`
	// `id` of the [Product](ctp:api:type:Product) the Line Item is based on.
	ProductId string `json:"productId"`
	// `key` of the [Product](ctp:api:type:Product).
	//
	// This field is only present on:
	//
	// - Line Items in a [Cart](ctp:api:type:Cart) when the `key` is available on that specific Product at the time the LineItem was created or updated on the Cart.
	// - Line Items in an [Order](ctp:api:type:Order) when the `key` is available on the specific Product at the time the Order was created from the Cart.
	//
	// Present on resources created or updated after 3 December 2021.
	ProductKey *string `json:"productKey,omitempty"`
	// Name of the Product.
	Name LocalizedString `json:"name"`
	// `slug` of the current version of the Product. Updated automatically if the `slug` changes. Empty if the Product has been deleted.
	// The `productSlug` field of LineItem is not expanded when using [Reference Expansion](/../api/general-concepts#reference-expansion).
	ProductSlug *LocalizedString `json:"productSlug,omitempty"`
	// Product Type of the Product.
	ProductType ProductTypeReference `json:"productType"`
	// Holds the data of the Product Variant added to the Cart.
	//
	// The data is saved at the time the Product Variant is added to the Cart and is not updated automatically when Product Variant data changes.
	// Must be updated using the [Recalculate](ctp:api:type:CartRecalculateAction) update action.
	Variant ProductVariant `json:"variant"`
	// Price of a Line Item selected from the Product Variant according to the [Product](ctp:api:type:Product) `priceMode`. If the `priceMode` is `Embedded` [ProductPriceMode](ctp:api:type:ProductPriceModeEnum) and the `variant` field hasn't been updated, the price may not correspond to a price in `variant.prices`.
	Price Price `json:"price"`
	// Number of Line Items of the given Product Variant present in the [Cart](ctp:api:type:Cart) or [Order](ctp:api:type:Order).
	Quantity int `json:"quantity"`
	// Total price of this Line Item equalling `price` multiplied by `quantity`. If the Line Item is discounted, the total price is the `discountedPricePerQuantity` multiplied by `quantity`.
	// Includes taxes if the [TaxRate](ctp:api:type:TaxRate) `includedInPrice` is `true`.
	//
	// If `ExternalPrice` [LineItemPriceMode](#ctp:api:type:LineItemPriceMode) is used with high-precision money, then the total price is rounded by using the `HalfEven` rounding mode.
	TotalPrice CentPrecisionMoney `json:"totalPrice"`
	// Discounted price of a single quantity of the Line Item.
	DiscountedPricePerQuantity []DiscountedLineItemPriceForQuantity `json:"discountedPricePerQuantity"`
	// Automatically set after `taxRate` is set.
	TaxedPrice *TaxedItemPrice `json:"taxedPrice,omitempty"`
	// Total taxed prices based on the quantity of Line Item assigned to each [Shipping Method](ctp:api:type:ShippingMethod). Only applicable for Carts with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	// Automatically set after `perMethodTaxRate` is set.
	TaxedPricePortions []MethodTaxedPrice `json:"taxedPricePortions"`
	// Tracks specific quantities of the Line Item within a given State. When a Line Item is added to a Cart, its full quantity is set to the built-in "Initial" state. State transitions for Line Items are managed on the [Order](ctp:api:type:Order).
	State []ItemState `json:"state"`
	// - For a Cart with `Platform` [TaxMode](ctp:api:type:TaxMode), the `taxRate` of Line Items is set automatically once a shipping address is set. The rate is based on the [TaxCategory](ctp:api:type:TaxCategory) that applies for the shipping address.
	// - For a Cart with `External` TaxMode, the `taxRate` of Line Items can be set using [ExternalTaxRateDraft](ctp:api:type:ExternalTaxRateDraft).
	TaxRate *TaxRate `json:"taxRate,omitempty"`
	// Tax Rate per Shipping Method for a Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode). For a Cart with `Platform` [TaxMode](ctp:api:type:TaxMode) it is automatically set after the [Shipping Method is added](ctp:api:type:CartAddShippingMethodAction).
	// For a Cart with `External` [TaxMode](ctp:api:type:TaxMode), the Tax Rate must be set with [ExternalTaxRateDraft](ctp:api:type:ExternalTaxRateDraft).
	PerMethodTaxRate []MethodTaxRate `json:"perMethodTaxRate"`
	// Identifies [Inventory entries](/../api/projects/inventory) that are reserved. The referenced Channel has the `InventorySupply` [ChannelRoleEnum](ctp:api:type:ChannelRoleEnum).
	SupplyChannel *ChannelReference `json:"supplyChannel,omitempty"`
	// Used to [select](/../api/pricing-and-discounts-overview#line-item-price-selection) a Product Price. The referenced Channel has the `ProductDistribution` [ChannelRoleEnum](ctp:api:type:ChannelRoleEnum).
	DistributionChannel *ChannelReference `json:"distributionChannel,omitempty"`
	// Indicates how the Price for the Line Item is set.
	PriceMode LineItemPriceMode `json:"priceMode"`
	// Indicates how the Line Item is added to the Cart.
	LineItemMode LineItemMode `json:"lineItemMode"`
	// Inventory mode specific to this Line Item only, and valid for the entire `quantity` of the Line Item.
	// Only present if the inventory mode is different from the `inventoryMode` specified on the [Cart](ctp:api:type:Cart).
	InventoryMode *InventoryMode `json:"inventoryMode,omitempty"`
	// Container for Line Item-specific addresses.
	ShippingDetails *ItemShippingDetails `json:"shippingDetails,omitempty"`
	// Custom Fields of the Line Item.
	Custom *CustomFields `json:"custom,omitempty"`
	// Date and time (UTC) the Line Item was added to the Cart.
	AddedAt *time.Time `json:"addedAt,omitempty"`
	// Date and time (UTC) the Line Item was last updated.
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`
	// Recurring Order and frequency data.
	RecurrenceInfo *LineItemRecurrenceInfo `json:"recurrenceInfo,omitempty"`
}

/**
*	A geographical location representing a country and optionally a state within this country.  A location can only be assigned to one Zone.
 */
type Location struct {
	// Country code of the geographic location.
	Country string `json:"country"`
	// State within the country.
	State *string `json:"state,omitempty"`
}

/**
*	Draft object to store money in cent amounts for a specific currency.
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
*	Determines the type of money used.
 */
type MoneyType string

const (
	MoneyTypeCentPrecision MoneyType = "centPrecision"
	MoneyTypeHighPrecision MoneyType = "highPrecision"
)

/**
*	Indicates the state of the Order.
*
 */
type OrderState string

const (
	OrderStateOpen      OrderState = "Open"
	OrderStateConfirmed OrderState = "Confirmed"
	OrderStateComplete  OrderState = "Complete"
	OrderStateCancelled OrderState = "Cancelled"
)

/**
*	Information regarding the appearance, content, and shipment of a Parcel.
*
 */
type Parcel struct {
	// Unique identifier of the Parcel.
	ID string `json:"id"`
	// User-defined unique identifier of the Parcel.
	Key *string `json:"key,omitempty"`
	// Date and time (UTC) the Parcel was created.
	CreatedAt time.Time `json:"createdAt"`
	// Information about the dimensions of the Parcel.
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
	// Shipment tracking information of the Parcel.
	TrackingData *TrackingData `json:"trackingData,omitempty"`
	// Line Items or Custom Line Items delivered in this Parcel.
	Items []DeliveryItem `json:"items"`
	// Custom Fields of the Parcel.
	Custom *CustomFields `json:"custom,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj Parcel) MarshalJSON() ([]byte, error) {
	type Alias Parcel
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

	if raw["items"] == nil {
		delete(raw, "items")
	}

	return json.Marshal(raw)

}

type ParcelMeasurements struct {
	// Height of the Parcel.
	HeightInMillimeter *int `json:"heightInMillimeter,omitempty"`
	// Length of the Parcel.
	LengthInMillimeter *int `json:"lengthInMillimeter,omitempty"`
	// Width of the Parcel.
	WidthInMillimeter *int `json:"widthInMillimeter,omitempty"`
	// Weight of the Parcel.
	WeightInGram *int `json:"weightInGram,omitempty"`
}

type PaymentInfo struct {
	// [References](ctp:api:type:Reference) to the Payments associated with the Order.
	Payments []PaymentReference `json:"payments"`
}

/**
*	Indicates the payment status for the Order.
*
 */
type PaymentState string

const (
	PaymentStateBalanceDue PaymentState = "BalanceDue"
	PaymentStateFailed     PaymentState = "Failed"
	PaymentStatePending    PaymentState = "Pending"
	PaymentStateCreditOwed PaymentState = "CreditOwed"
	PaymentStatePaid       PaymentState = "Paid"
)

/**
*	Permissions grant granular access to [Approval Rules](ctp:api:type:ApprovalRule), [Approval Flows](ctp:api:type:ApprovalFlow), [Business Units](ctp:api:type:BusinessUnit), [Carts](ctp:api:type:Cart), [Orders](ctp:api:type:Order), [Quotes](ctp:api:type:Quote), [Quote Requests](ctp:api:type:QuoteRequest), and [Shopping Lists](ctp:api:type:ShoppingList).
*
 */
type Permission string

const (
	PermissionAddChildUnits                      Permission = "AddChildUnits"
	PermissionUpdateAssociates                   Permission = "UpdateAssociates"
	PermissionUpdateBusinessUnitDetails          Permission = "UpdateBusinessUnitDetails"
	PermissionUpdateParentUnit                   Permission = "UpdateParentUnit"
	PermissionViewMyCarts                        Permission = "ViewMyCarts"
	PermissionViewOthersCarts                    Permission = "ViewOthersCarts"
	PermissionUpdateMyCarts                      Permission = "UpdateMyCarts"
	PermissionUpdateOthersCarts                  Permission = "UpdateOthersCarts"
	PermissionCreateMyCarts                      Permission = "CreateMyCarts"
	PermissionCreateOthersCarts                  Permission = "CreateOthersCarts"
	PermissionDeleteMyCarts                      Permission = "DeleteMyCarts"
	PermissionDeleteOthersCarts                  Permission = "DeleteOthersCarts"
	PermissionViewMyOrders                       Permission = "ViewMyOrders"
	PermissionViewOthersOrders                   Permission = "ViewOthersOrders"
	PermissionUpdateMyOrders                     Permission = "UpdateMyOrders"
	PermissionUpdateOthersOrders                 Permission = "UpdateOthersOrders"
	PermissionCreateMyOrdersFromMyCarts          Permission = "CreateMyOrdersFromMyCarts"
	PermissionCreateMyOrdersFromMyQuotes         Permission = "CreateMyOrdersFromMyQuotes"
	PermissionCreateOrdersFromOthersCarts        Permission = "CreateOrdersFromOthersCarts"
	PermissionCreateOrdersFromOthersQuotes       Permission = "CreateOrdersFromOthersQuotes"
	PermissionViewMyQuotes                       Permission = "ViewMyQuotes"
	PermissionViewOthersQuotes                   Permission = "ViewOthersQuotes"
	PermissionAcceptMyQuotes                     Permission = "AcceptMyQuotes"
	PermissionAcceptOthersQuotes                 Permission = "AcceptOthersQuotes"
	PermissionDeclineMyQuotes                    Permission = "DeclineMyQuotes"
	PermissionDeclineOthersQuotes                Permission = "DeclineOthersQuotes"
	PermissionRenegotiateMyQuotes                Permission = "RenegotiateMyQuotes"
	PermissionRenegotiateOthersQuotes            Permission = "RenegotiateOthersQuotes"
	PermissionReassignMyQuotes                   Permission = "ReassignMyQuotes"
	PermissionReassignOthersQuotes               Permission = "ReassignOthersQuotes"
	PermissionViewMyQuoteRequests                Permission = "ViewMyQuoteRequests"
	PermissionViewOthersQuoteRequests            Permission = "ViewOthersQuoteRequests"
	PermissionUpdateMyQuoteRequests              Permission = "UpdateMyQuoteRequests"
	PermissionUpdateOthersQuoteRequests          Permission = "UpdateOthersQuoteRequests"
	PermissionCreateMyQuoteRequestsFromMyCarts   Permission = "CreateMyQuoteRequestsFromMyCarts"
	PermissionCreateQuoteRequestsFromOthersCarts Permission = "CreateQuoteRequestsFromOthersCarts"
	PermissionCreateApprovalRules                Permission = "CreateApprovalRules"
	PermissionUpdateApprovalRules                Permission = "UpdateApprovalRules"
	PermissionUpdateApprovalFlows                Permission = "UpdateApprovalFlows"
	PermissionViewMyShoppingLists                Permission = "ViewMyShoppingLists"
	PermissionViewOthersShoppingLists            Permission = "ViewOthersShoppingLists"
	PermissionUpdateMyShoppingLists              Permission = "UpdateMyShoppingLists"
	PermissionUpdateOthersShoppingLists          Permission = "UpdateOthersShoppingLists"
	PermissionCreateMyShoppingLists              Permission = "CreateMyShoppingLists"
	PermissionCreateOthersShoppingLists          Permission = "CreateOthersShoppingLists"
	PermissionDeleteMyShoppingLists              Permission = "DeleteMyShoppingLists"
	PermissionDeleteOthersShoppingLists          Permission = "DeleteOthersShoppingLists"
)

/**
*	The representation for prices embedded in [LineItems](ctp:api:type:LineItem) and in [ProductVariants](ctp:api:type:ProductVariant) when the [ProductPriceMode](ctp:api:type:ProductPriceModeEnum) is `Embedded`.
*	For the `Standalone` ProductPriceMode refer to [StandalonePrice](ctp:api:type:StandalonePrice).
 */
type Price struct {
	// Unique identifier of this Price.
	ID string `json:"id"`
	// User-defined identifier of the Price. It is unique per [ProductVariant](ctp:api:type:ProductVariant).
	Key *string `json:"key,omitempty"`
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
	// Date and time until this Price is valid. Prices that are no longer valid are not automatically removed, but they can be [removed](ctp:api:type:ProductRemovePriceAction) if necessary.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
	// Is set if a [ProductDiscount](ctp:api:type:ProductDiscount) has been applied.
	// If set, the API uses the DiscountedPrice value for the [Line Item price selection](/../api/pricing-and-discounts-overview#line-item-price-selection).
	// When a [relative discount](ctp:api:type:ProductDiscountValueRelative) has been applied and the fraction part of the DiscountedPrice `value` is 0.5, the `value` is rounded in favor of the customer with [half-down rounding](https://en.wikipedia.org/wiki/Rounding#Rounding_half_down).
	Discounted *DiscountedPrice `json:"discounted,omitempty"`
	// Present if different Prices for certain [LineItem](ctp:api:type:LineItem) quantities have been specified.
	//
	// If `discounted` is present, the tiered Price is ignored for a Product Variant.
	Tiers []PriceTier `json:"tiers"`
	// Custom Fields defined for the Price.
	Custom *CustomFields `json:"custom,omitempty"`
	// [Recurrence Policy](ctp:api:type:RecurrencePolicy) for which this Price is valid.
	RecurrencePolicy *RecurrencePolicyReference `json:"recurrencePolicy,omitempty"`
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

type ProductSelectionSetting struct {
	// Reference to a ProductSelection.
	ProductSelection ProductSelectionReference `json:"productSelection"`
	// If `true`, all Products assigned to this Product Selection are part of the Store's assortment.
	Active bool `json:"active"`
}

/**
*	The [InventoryEntry](ctp:api:type:InventoryEntry) information of the Product Variant. If there is a supply [Channel](ctp:api:type:Channel) for the InventoryEntry, then `channels` is returned. If not, then `isOnStock`, `restockableInDays`, and `availableQuantity` are returned.
*
 */
type ProductVariantAvailability struct {
	// For each [InventoryEntry](ctp:api:type:InventoryEntry) with a supply Channel, an entry is added to `channels`.
	Channels *ProductVariantChannelAvailabilityMap `json:"channels,omitempty"`
	// Indicates whether a Product Variant is in stock.
	IsOnStock *bool `json:"isOnStock,omitempty"`
	// Number of days to restock a Product Variant once it is out of stock.
	RestockableInDays *int `json:"restockableInDays,omitempty"`
	// Number of items of the Product Variant that are in stock.
	AvailableQuantity *int `json:"availableQuantity,omitempty"`
	// Unique identifier of the [InventoryEntry](ctp:api:type:InventoryEntry).
	ID *string `json:"id,omitempty"`
	// Current version of the [InventoryEntry](ctp:api:type:InventoryEntry).
	Version *int `json:"version,omitempty"`
}

type ProductVariantChannelAvailability struct {
	// Indicates whether a Product Variant is in stock in a specified [Channel](ctp:api:type:Channel).
	IsOnStock *bool `json:"isOnStock,omitempty"`
	// Number of days to restock a Product Variant once it is out of stock in a specified [Channel](ctp:api:type:Channel).
	RestockableInDays *int `json:"restockableInDays,omitempty"`
	// Number of items of this Product Variant that are in stock in a specified [Channel](ctp:api:type:Channel).
	AvailableQuantity *int `json:"availableQuantity,omitempty"`
	// Unique identifier of the [InventoryEntry](ctp:api:type:InventoryEntry).
	ID string `json:"id"`
	// Current version of the [InventoryEntry](ctp:api:type:InventoryEntry).
	Version int `json:"version"`
}

/**
*	JSON object where the keys are supply [Channel](/projects/channels) `id`, and the values are [ProductVariantChannelAvailability](/projects/products#productvariantchannelavailability).
*
 */
type ProductVariantChannelAvailabilityMap map[string]ProductVariantChannelAvailability

/**
*	Polymorphic base type for Product Variant Selections. The actual type is determined by the `type` field.
*
 */
type ProductVariantSelection struct {
	// Determines whether the SKUs are to be included in, or excluded from, the Product Selection.
	Type ProductVariantSelectionTypeEnum `json:"type"`
}

type ProductVariantSelectionTypeEnum string

const (
	ProductVariantSelectionTypeEnumInclusion        ProductVariantSelectionTypeEnum = "inclusion"
	ProductVariantSelectionTypeEnumExclusion        ProductVariantSelectionTypeEnum = "exclusion"
	ProductVariantSelectionTypeEnumIncludeOnly      ProductVariantSelectionTypeEnum = "includeOnly"
	ProductVariantSelectionTypeEnumIncludeAllExcept ProductVariantSelectionTypeEnum = "includeAllExcept"
)

/**
*	Predefined states tracking the status of the Quote Request in the negotiation process.
*
 */
type QuoteRequestState string

const (
	QuoteRequestStateSubmitted QuoteRequestState = "Submitted"
	QuoteRequestStateAccepted  QuoteRequestState = "Accepted"
	QuoteRequestStateClosed    QuoteRequestState = "Closed"
	QuoteRequestStateRejected  QuoteRequestState = "Rejected"
	QuoteRequestStateCancelled QuoteRequestState = "Cancelled"
)

/**
*	Predefined states tracking the status of the Quote.
*
 */
type QuoteState string

const (
	QuoteStatePending                  QuoteState = "Pending"
	QuoteStateDeclined                 QuoteState = "Declined"
	QuoteStateDeclinedForRenegotiation QuoteState = "DeclinedForRenegotiation"
	QuoteStateRenegotiationAddressed   QuoteState = "RenegotiationAddressed"
	QuoteStateAccepted                 QuoteState = "Accepted"
	QuoteStateWithdrawn                QuoteState = "Withdrawn"
)

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
	case "product":
		obj := ProductReference{}
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
	case "recurrence-policy":
		obj := RecurrencePolicyReference{}
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
	case "state":
		obj := StateReference{}
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
	}
	return nil, nil
}

/**
*	Type of resource the value should reference. Supported resource type identifiers are:
*
 */
type ReferenceTypeId string

const (
	ReferenceTypeIdApprovalFlow          ReferenceTypeId = "approval-flow"
	ReferenceTypeIdApprovalRule          ReferenceTypeId = "approval-rule"
	ReferenceTypeIdAssociateRole         ReferenceTypeId = "associate-role"
	ReferenceTypeIdAttributeGroup        ReferenceTypeId = "attribute-group"
	ReferenceTypeIdBusinessUnit          ReferenceTypeId = "business-unit"
	ReferenceTypeIdCart                  ReferenceTypeId = "cart"
	ReferenceTypeIdCartDiscount          ReferenceTypeId = "cart-discount"
	ReferenceTypeIdCategory              ReferenceTypeId = "category"
	ReferenceTypeIdChannel               ReferenceTypeId = "channel"
	ReferenceTypeIdCustomer              ReferenceTypeId = "customer"
	ReferenceTypeIdCustomerEmailToken    ReferenceTypeId = "customer-email-token"
	ReferenceTypeIdCustomerGroup         ReferenceTypeId = "customer-group"
	ReferenceTypeIdCustomerPasswordToken ReferenceTypeId = "customer-password-token"
	ReferenceTypeIdDirectDiscount        ReferenceTypeId = "direct-discount"
	ReferenceTypeIdDiscountCode          ReferenceTypeId = "discount-code"
	ReferenceTypeIdDiscountGroup         ReferenceTypeId = "discount-group"
	ReferenceTypeIdExtension             ReferenceTypeId = "extension"
	ReferenceTypeIdInventoryEntry        ReferenceTypeId = "inventory-entry"
	ReferenceTypeIdKeyValueDocument      ReferenceTypeId = "key-value-document"
	ReferenceTypeIdOrder                 ReferenceTypeId = "order"
	ReferenceTypeIdOrderEdit             ReferenceTypeId = "order-edit"
	ReferenceTypeIdPaymentMethod         ReferenceTypeId = "payment-method"
	ReferenceTypeIdPayment               ReferenceTypeId = "payment"
	ReferenceTypeIdProduct               ReferenceTypeId = "product"
	ReferenceTypeIdProductDiscount       ReferenceTypeId = "product-discount"
	ReferenceTypeIdProductPrice          ReferenceTypeId = "product-price"
	ReferenceTypeIdProductSelection      ReferenceTypeId = "product-selection"
	ReferenceTypeIdProductTailoring      ReferenceTypeId = "product-tailoring"
	ReferenceTypeIdProductType           ReferenceTypeId = "product-type"
	ReferenceTypeIdQuote                 ReferenceTypeId = "quote"
	ReferenceTypeIdQuoteRequest          ReferenceTypeId = "quote-request"
	ReferenceTypeIdRecurrencePolicy      ReferenceTypeId = "recurrence-policy"
	ReferenceTypeIdRecurringOrder        ReferenceTypeId = "recurring-order"
	ReferenceTypeIdReview                ReferenceTypeId = "review"
	ReferenceTypeIdShippingMethod        ReferenceTypeId = "shipping-method"
	ReferenceTypeIdShoppingList          ReferenceTypeId = "shopping-list"
	ReferenceTypeIdStagedQuote           ReferenceTypeId = "staged-quote"
	ReferenceTypeIdStandalonePrice       ReferenceTypeId = "standalone-price"
	ReferenceTypeIdState                 ReferenceTypeId = "state"
	ReferenceTypeIdStore                 ReferenceTypeId = "store"
	ReferenceTypeIdSubscription          ReferenceTypeId = "subscription"
	ReferenceTypeIdTaxCategory           ReferenceTypeId = "tax-category"
	ReferenceTypeIdType                  ReferenceTypeId = "type"
	ReferenceTypeIdZone                  ReferenceTypeId = "zone"
)

type Reservation struct {
	Quantity int `json:"quantity"`
	// A Reference represents a loose reference to another resource in the same Project identified by its `id`. The `typeId` indicates the type of the referenced resource. Each resource type has its corresponding Reference type, like [ChannelReference](ctp:api:type:ChannelReference).  A referenced resource can be embedded through [Reference Expansion](/general-concepts#reference-expansion). The expanded reference is the value of an additional `obj` field then.
	Owner             Reference `json:"owner"`
	CreatedAt         string    `json:"createdAt"`
	CheckoutStartedAt string    `json:"checkoutStartedAt"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *Reservation) UnmarshalJSON(data []byte) error {
	type Alias Reservation
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Owner != nil {
		var err error
		obj.Owner, err = mapDiscriminatorReference(obj.Owner)
		if err != nil {
			return err
		}
	}

	return nil
}

type ResourceIdentifier struct {
	ID  *string `json:"id,omitempty"`
	Key *string `json:"key,omitempty"`
	// Type of resource the value should reference. Supported resource type identifiers are:
	TypeId ReferenceTypeId `json:"typeId"`
}

/**
*	Stores information about returns connected to an Order.
*
 */
type ReturnInfo struct {
	// Information on the Line Items or Custom Line Items returned.
	Items []ReturnItem `json:"items"`
	// User-defined identifier to track the return.
	ReturnTrackingId *string `json:"returnTrackingId,omitempty"`
	// Date and time (UTC) the return is initiated.
	ReturnDate *time.Time `json:"returnDate,omitempty"`
}

type ReturnItem struct {
	// Unique identifier of the Return Item.
	ID string `json:"id"`
	// User-defined unique identifier of the Return Item.
	Key *string `json:"key,omitempty"`
	// Number of Line Items or Custom Line Items returned.
	Quantity int    `json:"quantity"`
	Type     string `json:"type"`
	// User-defined description for the return.
	Comment *string `json:"comment,omitempty"`
	// Shipment status of the Return Item.
	ShipmentState ReturnShipmentState `json:"shipmentState"`
	// Payment status of the Return Item:
	//
	// - `NonRefundable`, for items in the `Advised` [ReturnShipmentState](ctp:api:type:ReturnShipmentState)
	// - `Initial`, for items in the `Returned` [ReturnShipmentState](ctp:api:type:ReturnShipmentState)
	PaymentState ReturnPaymentState `json:"paymentState"`
	// Custom Fields of the Return Item.
	Custom *CustomFields `json:"custom,omitempty"`
	// Date and time (UTC) the Return Item was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Date and time (UTC) the Return Item was initially created.
	CreatedAt time.Time `json:"createdAt"`
}

type ReturnPaymentState string

const (
	ReturnPaymentStateNonRefundable ReturnPaymentState = "NonRefundable"
	ReturnPaymentStateInitial       ReturnPaymentState = "Initial"
	ReturnPaymentStateRefunded      ReturnPaymentState = "Refunded"
	ReturnPaymentStateNotRefunded   ReturnPaymentState = "NotRefunded"
)

type ReturnShipmentState string

const (
	ReturnShipmentStateAdvised     ReturnShipmentState = "Advised"
	ReturnShipmentStateReturned    ReturnShipmentState = "Returned"
	ReturnShipmentStateBackInStock ReturnShipmentState = "BackInStock"
	ReturnShipmentStateUnusable    ReturnShipmentState = "Unusable"
)

type ReviewRatingStatistics struct {
	// Average rating of one target
	// This number is rounded with 5 decimals.
	AverageRating float64 `json:"averageRating"`
	// Highest rating of one target
	HighestRating float64 `json:"highestRating"`
	// Lowest rating of one target
	LowestRating float64 `json:"lowestRating"`
	// Number of ratings taken into account
	Count int `json:"count"`
	// Full distribution of the ratings.
	// The keys are the different ratings and the values are the count of reviews having this rating.
	// Only the used ratings appear in this object.
	RatingsDistribution interface{} `json:"ratingsDistribution"`
}

/**
*	Determines how monetary values are rounded.
*
 */
type RoundingMode string

const (
	RoundingModeHalfEven RoundingMode = "HalfEven"
	RoundingModeHalfUp   RoundingMode = "HalfUp"
	RoundingModeHalfDown RoundingMode = "HalfDown"
)

type SearchKeyword struct {
	// Text to return in the [SuggestionResult](ctp:api:type:SuggestionResult).
	Text string `json:"text"`
	// If no tokenizer is defined, the `text` is used as a single token.
	SuggestTokenizer *SuggestTokenizer `json:"suggestTokenizer,omitempty"`
}

/**
*	Search keywords are JSON objects primarily used by [Search Term Suggestions](/projects/search-term-suggestions), but are also considered for a [full text search](/projects/product-projection-search#full-text-search) in the Product Projection Search API.
*	The keys are of type [Locale](ctp:api:type:Locale), and the values are an array of [SearchKeyword](ctp:api:type:SearchKeyword).
*
 */
type SearchKeywords map[string][]SearchKeyword

/**
*	Defines which matching items are to be discounted.
*
 */
type SelectionMode string

const (
	SelectionModeCheapest      SelectionMode = "Cheapest"
	SelectionModeMostExpensive SelectionMode = "MostExpensive"
)

/**
*	Indicates the shipment status of the Order.
*
 */
type ShipmentState string

const (
	ShipmentStateShipped   ShipmentState = "Shipped"
	ShipmentStateDelivered ShipmentState = "Delivered"
	ShipmentStateReady     ShipmentState = "Ready"
	ShipmentStatePending   ShipmentState = "Pending"
	ShipmentStateDelayed   ShipmentState = "Delayed"
	ShipmentStatePartial   ShipmentState = "Partial"
	ShipmentStateBackorder ShipmentState = "Backorder"
	ShipmentStateCanceled  ShipmentState = "Canceled"
)

type ShippingRate struct {
	// Currency amount of the ShippingRate.
	Price CentPrecisionMoney `json:"price"`
	// [Free shipping](/../api/shipping-delivery-overview#free-shipping) is applied if the sum of the (Custom) Line Item Prices reaches the specified value.
	FreeAbove *CentPrecisionMoney `json:"freeAbove,omitempty"`
	// `true` if the ShippingRate matches given [Cart](ctp:api:type:Cart) or [Location](ctp:api:type:Location).
	// Only appears in response to requests for [Get ShippingMethods for a Cart](ctp:api:endpoint:/{projectKey}/shipping-methods/matching-cart:GET) or
	// [Get ShippingMethods for a Location](ctp:api:endpoint:/{projectKey}/shipping-methods/matching-location:GET).
	IsMatching *bool `json:"isMatching,omitempty"`
	// Price tiers for the ShippingRate.
	Tiers []ShippingRatePriceTier `json:"tiers"`
}

type ShippingRatePriceTier struct {
	Type ShippingRateTierType `json:"type"`
}

type ShippingRateTierType string

const (
	ShippingRateTierTypeCartValue          ShippingRateTierType = "CartValue"
	ShippingRateTierTypeCartClassification ShippingRateTierType = "CartClassification"
	ShippingRateTierTypeCartScore          ShippingRateTierType = "CartScore"
)

/**
*	Describes how the Cart Discount interacts with other Discounts.
*
 */
type StackingMode string

const (
	StackingModeStacking              StackingMode = "Stacking"
	StackingModeStopAfterThisDiscount StackingMode = "StopAfterThisDiscount"
)

/**
*	Predefined states tracking the status of the Staged Quote.
*
 */
type StagedQuoteState string

const (
	StagedQuoteStateInProgress StagedQuoteState = "InProgress"
	StagedQuoteStateSent       StagedQuoteState = "Sent"
	StagedQuoteStateClosed     StagedQuoteState = "Closed"
)

/**
*	For some resource types, a State can fulfill the following predefined roles:
*
 */
type StateRoleEnum string

const (
	StateRoleEnumReviewIncludedInStatistics StateRoleEnum = "ReviewIncludedInStatistics"
	StateRoleEnumReturn                     StateRoleEnum = "Return"
)

/**
*	Resource or object type the State can be assigned to.
*
 */
type StateTypeEnum string

const (
	StateTypeEnumOrderState          StateTypeEnum = "OrderState"
	StateTypeEnumRecurringOrderState StateTypeEnum = "RecurringOrderState"
	StateTypeEnumLineItemState       StateTypeEnum = "LineItemState"
	StateTypeEnumProductState        StateTypeEnum = "ProductState"
	StateTypeEnumReviewState         StateTypeEnum = "ReviewState"
	StateTypeEnumPaymentState        StateTypeEnum = "PaymentState"
	StateTypeEnumQuoteRequestState   StateTypeEnum = "QuoteRequestState"
	StateTypeEnumStagedQuoteState    StateTypeEnum = "StagedQuoteState"
	StateTypeEnumQuoteState          StateTypeEnum = "QuoteState"
)

type StoreCountry struct {
	// Two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Code string `json:"code"`
}

/**
*	It is used to calculate the [taxPortions](/../api/projects/carts#taxedprice) field in a Cart or Order.
 */
type SubRate struct {
	// Name of the SubRate.
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}

type SuggestTokenizer struct {
	Type string `json:"type"`
}

/**
*	Contains synchronization activity information of the Order (like export or import).
*
 */
type SyncInfo struct {
	// Connection to a synchronization destination.
	Channel ChannelReference `json:"channel"`
	// Identifier of an external order instance, file, or other resource.
	ExternalId *string `json:"externalId,omitempty"`
	// Date and time (UTC) the information was synced.
	SyncedAt time.Time `json:"syncedAt"`
}

/**
*	Determines in which [Tax calculation mode](/carts-orders-overview#tax-calculation-mode) taxed prices are calculated.
*
 */
type TaxCalculationMode string

const (
	TaxCalculationModeLineItemLevel  TaxCalculationMode = "LineItemLevel"
	TaxCalculationModeUnitPriceLevel TaxCalculationMode = "UnitPriceLevel"
)

/**
*	Indicates how taxes are set on the Cart.
*
 */
type TaxMode string

const (
	TaxModePlatform       TaxMode = "Platform"
	TaxModeExternal       TaxMode = "External"
	TaxModeExternalAmount TaxMode = "ExternalAmount"
	TaxModeDisabled       TaxMode = "Disabled"
)

type TaxRate struct {
	// Present if the TaxRate is part of a [TaxCategory](ctp:api:type:TaxCategory).
	// Absent for external TaxRates in [LineItem](ctp:api:type:LineItem), [CustomLineItem](ctp:api:type:CustomLineItem), and [ShippingInfo](ctp:api:type:ShippingInfo).
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the TaxRate.
	// Present when set using [TaxRateDraft](ctp:api:type:TaxRateDraft). Not available for external TaxRates created using [ExternalTaxRateDraft](ctp:api:type:ExternalTaxRateDraft).
	Key *string `json:"key,omitempty"`
	// Name of the TaxRate.
	Name string `json:"name"`
	// Tax rate. If subrates are used, the amount is the sum of all rates in `subRates`.
	Amount float64 `json:"amount"`
	// If `true`, tax is included in [Embedded Prices](ctp:api:type:Price) or [Standalone Prices](ctp:api:type:StandalonePrice), and the `taxedPrice` is present on [LineItems](ctp:api:type:LineItem). In this case, the `totalNet` price on [TaxedPrice](ctp:api:type:TaxedPrice) includes the TaxRate.
	IncludedInPrice bool `json:"includedInPrice"`
	// Country in which the tax rate is applied in [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) format.
	Country string `json:"country"`
	// State within the country, such as Texas in the United States.
	State *string `json:"state,omitempty"`
	// Used when the total tax is a combination of multiple taxes (for example, local, state/provincial, and/or federal taxes). The total of all subrates must equal the TaxRate `amount`.
	// These subrates are used to calculate the `taxPortions` field of a [Cart](ctp:api:type:Cart) or [Order](ctp:api:type:Order) and the `taxedPrice` field of [LineItems](ctp:api:type:LineItem), [CustomLineItems](ctp:api:type:CustomLineItem), and [ShippingInfos](ctp:api:type:ShippingInfo).
	SubRates []SubRate `json:"subRates"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TaxRate) MarshalJSON() ([]byte, error) {
	type Alias TaxRate
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

	if raw["subRates"] == nil {
		delete(raw, "subRates")
	}

	return json.Marshal(raw)

}

type TaxedItemPrice struct {
	// Total net amount of the Line Item or Custom Line Item.
	TotalNet CentPrecisionMoney `json:"totalNet"`
	// Total gross amount of the Line Item or Custom Line Item.
	TotalGross CentPrecisionMoney `json:"totalGross"`
	// Taxable portions added to the total net price.
	//
	// Calculated from the [TaxRates](ctp:api:type:TaxRate).
	TaxPortions []TaxPortion `json:"taxPortions"`
	// Total tax applicable for the Line Item or Custom Line Item.
	// Automatically calculated as the difference between the `totalGross` and `totalNet` values.
	TotalTax *CentPrecisionMoney `json:"totalTax,omitempty"`
}

type TaxedPrice struct {
	// Total net price of the Cart or Order.
	TotalNet CentPrecisionMoney `json:"totalNet"`
	// Total gross price of the Cart or Order.
	TotalGross CentPrecisionMoney `json:"totalGross"`
	// Taxable portions added to the total net price.
	//
	// Calculated from the [TaxRates](ctp:api:type:TaxRate).
	TaxPortions []TaxPortion `json:"taxPortions"`
	// Total tax applicable for the Cart or Order.
	//
	// Automatically calculated as the difference between the `totalGross` and `totalNet` values.
	TotalTax *CentPrecisionMoney `json:"totalTax,omitempty"`
}

/**
*	A text input hint is a string with one of the following values:
*
 */
type TextInputHint string

const (
	TextInputHintSingleLine TextInputHint = "SingleLine"
	TextInputHintMultiLine  TextInputHint = "MultiLine"
)

/**
*	TextLineItems are Line Items that use text values instead of references to Products.
*
 */
type TextLineItem struct {
	// Date and time (UTC) the TextLineItem was added to the [ShoppingList](ctp:api:type:ShoppingList).
	AddedAt time.Time `json:"addedAt"`
	// Custom Fields of the TextLineItem.
	Custom *CustomFields `json:"custom,omitempty"`
	// Description of the TextLineItem.
	Description *LocalizedString `json:"description,omitempty"`
	// Unique identifier of the TextLineItem.
	ID string `json:"id"`
	// User-defined identifier of the TextLineItem. It is unique per [ShoppingList](ctp:api:type:ShoppingList).
	Key *string `json:"key,omitempty"`
	// Name of the TextLineItem.
	Name LocalizedString `json:"name"`
	// Number of entries in the TextLineItem.
	Quantity int `json:"quantity"`
}

/**
*	Information that helps track a Parcel.
*
 */
type TrackingData struct {
	// Identifier to track the Parcel.
	TrackingId *string `json:"trackingId,omitempty"`
	// Name of the carrier that delivers the Parcel.
	Carrier *string `json:"carrier,omitempty"`
	// Name of the provider that serves as facade to several carriers.
	Provider *string `json:"provider,omitempty"`
	// Transaction identifier with the `provider`.
	ProviderTransaction *string `json:"providerTransaction,omitempty"`
	// - If `true`, the Parcel is being returned.
	// - If `false`, the Parcel is being delivered to the customer.
	IsReturn *bool `json:"isReturn,omitempty"`
}

/**
*	Represents a financial transaction typically created as a result of a notification from the payment service.
*
 */
type Transaction struct {
	// Unique identifier of the Transaction.
	ID string `json:"id"`
	// Date and time (UTC) the Transaction took place.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// Type of the Transaction. For example, `Authorization`.
	Type TransactionType `json:"type"`
	// Money value of the Transaction.
	Amount CentPrecisionMoney `json:"amount"`
	// Identifier used by the interface that manages the Transaction (usually the PSP).
	// If a matching interaction was logged in the `interfaceInteractions` array, the corresponding interaction can be found with this ID.
	InteractionId *string `json:"interactionId,omitempty"`
	// State of the Transaction.
	State TransactionState `json:"state"`
	// Custom Fields defined for the Transaction.
	Custom *CustomFields `json:"custom,omitempty"`
}

/**
*	Transactions can be in one of the following States:
*
 */
type TransactionState string

const (
	TransactionStateInitial TransactionState = "Initial"
	TransactionStatePending TransactionState = "Pending"
	TransactionStateSuccess TransactionState = "Success"
	TransactionStateFailure TransactionState = "Failure"
)

type TransactionType string

const (
	TransactionTypeAuthorization       TransactionType = "Authorization"
	TransactionTypeCancelAuthorization TransactionType = "CancelAuthorization"
	TransactionTypeCharge              TransactionType = "Charge"
	TransactionTypeRefund              TransactionType = "Refund"
	TransactionTypeChargeback          TransactionType = "Chargeback"
)

type Variant struct {
	ID  int    `json:"id"`
	Sku string `json:"sku"`
	Key string `json:"key"`
}

type Attribute struct {
	// Name of the Attribute.
	Name string `json:"name"`
	// The [AttributeType](ctp:api:type:AttributeType) determines the format of the Attribute `value` to be provided:
	//
	// - For [Enum Type](ctp:api:type:AttributeEnumType) and [Localized Enum Type](ctp:api:type:AttributeLocalizedEnumType),
	//   use the `key` of the [Plain Enum Value](ctp:api:type:AttributePlainEnumValue) or [Localized Enum Value](ctp:api:type:AttributeLocalizedEnumValue) objects,
	//   or the complete objects as `value`.
	// - For [Localizable Text Type](ctp:api:type:AttributeLocalizableTextType), use the [LocalizedString](ctp:api:type:LocalizedString) object as `value`.
	// - For [Money Type](ctp:api:type:AttributeMoneyType) Attributes, use the [Money](ctp:api:type:Money) object as `value`.
	// - For [Set Type](ctp:api:type:AttributeSetType) Attributes, use the entire `set` object  as `value`.
	// - For [Nested Type](ctp:api:type:AttributeNestedType) Attributes, use the list of values of all Attributes of the nested Product as `value`.
	// - For [Reference Type](ctp:api:type:AttributeReferenceType) Attributes, use the [Reference](ctp:api:type:Reference) object as `value`.
	Value interface{} `json:"value"`
}

/**
*	Determines whether a Business Unit can inherit [Approval Rules](/projects/approval-rules) from a parent. Only Business Units of type `Division` can use `ExplicitAndFromParent`.
*
 */
type BusinessUnitApprovalRuleMode string

const (
	BusinessUnitApprovalRuleModeExplicit              BusinessUnitApprovalRuleMode = "Explicit"
	BusinessUnitApprovalRuleModeExplicitAndFromParent BusinessUnitApprovalRuleMode = "ExplicitAndFromParent"
)

/**
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [BusinessUnit](ctp:api:type:BusinessUnit). Either `id` or `key` is required. If both are set, an [InvalidJsonInput](/../api/errors#invalidjsoninput) error is returned.
*
 */
type BusinessUnitResourceIdentifier struct {
	// Unique identifier of the referenced [BusinessUnit](ctp:api:type:BusinessUnit). Required if `key` is absent.
	ID *string `json:"id,omitempty"`
	// Unique key of the referenced [BusinessUnit](ctp:api:type:BusinessUnit). Required if `id` is absent.
	Key *string `json:"key,omitempty"`
	// Type of resource the value should reference. Supported resource type identifiers are:
	TypeId ReferenceTypeId `json:"typeId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"null"`
		*Alias
	}{Action: "business-unit", Alias: (*Alias)(&obj)})
}

/**
*	The type of the Business Unit indicating its position in a hierarchy.
*
 */
type BusinessUnitType string

const (
	BusinessUnitTypeCompany  BusinessUnitType = "Company"
	BusinessUnitTypeDivision BusinessUnitType = "Division"
)

/**
*	Defines an allowed value of a [CustomFieldEnumType](ctp:api:type:CustomFieldEnumType) field.
*
 */
type CustomFieldEnumValue struct {
	// Key of the value used as a programmatic identifier.
	Key string `json:"key"`
	// Descriptive label of the value.
	Label string `json:"label"`
}

/**
*	Defines an allowed value of a [CustomFieldLocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType) field.
*
 */
type CustomFieldLocalizedEnumValue struct {
	// Key of the value used as a programmatic identifier.
	Key string `json:"key"`
	// Descriptive localized label of the value.
	Label LocalizedString `json:"label"`
}

type CustomerGroupAssignment struct {
	// Reference to a Customer Group.
	CustomerGroup CustomerGroupReference `json:"customerGroup"`
}

type DiscountOnTotalPrice struct {
	// Money value of the discount on the total price of the Cart or Order.
	DiscountedAmount TypedMoney `json:"discountedAmount"`
	// Discounts that impact the total price of the Cart or Order.
	IncludedDiscounts []DiscountedTotalPricePortion `json:"includedDiscounts"`
	// Money value of the discount on the total net price of the Cart or Order.
	//
	// The same percentage of discount applies as on the `discountedAmount`.
	// Present only when `taxedPrice` of the Cart or Order exists.
	DiscountedNetAmount TypedMoney `json:"discountedNetAmount,omitempty"`
	// Money value of the discount on the total gross price of the Cart or Order.
	//
	// The same percentage of discount applies as on the `discountedAmount`.
	// Present only when `taxedPrice` of the Cart or Order exists.
	DiscountedGrossAmount TypedMoney `json:"discountedGrossAmount,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DiscountOnTotalPrice) UnmarshalJSON(data []byte) error {
	type Alias DiscountOnTotalPrice
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.DiscountedAmount != nil {
		var err error
		obj.DiscountedAmount, err = mapDiscriminatorTypedMoney(obj.DiscountedAmount)
		if err != nil {
			return err
		}
	}
	if obj.DiscountedNetAmount != nil {
		var err error
		obj.DiscountedNetAmount, err = mapDiscriminatorTypedMoney(obj.DiscountedNetAmount)
		if err != nil {
			return err
		}
	}
	if obj.DiscountedGrossAmount != nil {
		var err error
		obj.DiscountedGrossAmount, err = mapDiscriminatorTypedMoney(obj.DiscountedGrossAmount)
		if err != nil {
			return err
		}
	}

	return nil
}

type DiscountedTotalPricePortion struct {
	// A [CartDiscountReference](ctp:api:type:CartDiscountReference) or [DirectDiscountReference](ctp:api:type:DirectDiscountReference) to the discount applied on the Cart `totalPrice`.
	Discount Reference `json:"discount"`
	// Money value of the discount.
	DiscountedAmount TypedMoney `json:"discountedAmount"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DiscountedTotalPricePortion) UnmarshalJSON(data []byte) error {
	type Alias DiscountedTotalPricePortion
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Discount != nil {
		var err error
		obj.Discount, err = mapDiscriminatorReference(obj.Discount)
		if err != nil {
			return err
		}
	}
	if obj.DiscountedAmount != nil {
		var err error
		obj.DiscountedAmount, err = mapDiscriminatorTypedMoney(obj.DiscountedAmount)
		if err != nil {
			return err
		}
	}

	return nil
}

/**
*	This mode determines the type of Prices used for [price selection](/../api/pricing-and-discounts-overview#price-selection) by Line Items and Products.
*	For more information about the difference between the Prices, see [Pricing](/../api/pricing-and-discounts-overview).
*
 */
type ProductPriceModeEnum string

const (
	ProductPriceModeEnumEmbedded   ProductPriceModeEnum = "Embedded"
	ProductPriceModeEnumStandalone ProductPriceModeEnum = "Standalone"
)

/**
*	Only Product Variants with the explicitly listed SKUs are part of a Product Selection with `IndividualExclusion` [ProductSelectionMode](ctp:api:type:ProductSelectionMode).
*
 */
type ProductVariantExclusion struct {
	// Non-empty array of SKUs representing Product Variants to be included in the Product Selection with `IndividualExclusion` [ProductSelectionMode](ctp:api:type:ProductSelectionMode).
	Skus []string `json:"skus"`
}

type Shipping struct {
	// User-defined unique identifier of the Shipping in a Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey string `json:"shippingKey"`
	// Automatically set when the [Shipping Method is added](ctp:api:type:CartAddShippingMethodAction).
	ShippingInfo ShippingInfo `json:"shippingInfo"`
	// Determines the shipping rates and Tax Rates of associated Line Items.
	ShippingAddress Address `json:"shippingAddress"`
	// Used as an input to select a [ShippingRatePriceTier](ctp:api:type:ShippingRatePriceTier).
	// The data type of this field depends on the `shippingRateInputType.type` configured in the [Project](ctp:api:type:Project):
	//
	// - If `CartClassification`, it is [ClassificationShippingRateInput](ctp:api:type:ClassificationShippingRateInput).
	// - If `CartScore`, it is [ScoreShippingRateInput](ctp:api:type:ScoreShippingRateInput).
	// - If `CartValue`, it cannot be used.
	ShippingRateInput *ShippingRateInput `json:"shippingRateInput,omitempty"`
	// Custom Fields of Shipping with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingCustomFields *CustomFields `json:"shippingCustomFields,omitempty"`
}

type ShippingInfo struct {
	// Name of the Shipping Method.
	ShippingMethodName string `json:"shippingMethodName"`
	// Determined based on the [ShippingRate](ctp:api:type:ShippingRate) and its tiered prices, and either the sum of [LineItem](ctp:api:type:LineItem) prices or the `shippingRateInput` field.
	Price CentPrecisionMoney `json:"price"`
	// Used to determine the price.
	ShippingRate ShippingRate `json:"shippingRate"`
	// Automatically set after the `taxRate` is set.
	TaxedPrice *TaxedItemPrice `json:"taxedPrice,omitempty"`
	// Automatically set in the `Platform` [TaxMode](ctp:api:type:TaxMode) after the [shipping address is set](ctp:api:type:CartSetShippingAddressAction).
	//
	// For the `External` [TaxMode](ctp:api:type:TaxMode) the Tax Rate must be set explicitly with the [ExternalTaxRateDraft](ctp:api:type:ExternalTaxRateDraft).
	TaxRate *TaxRate `json:"taxRate,omitempty"`
	// Used to select a Tax Rate when a Cart has the `Platform` [TaxMode](ctp:api:type:TaxMode).
	TaxCategory *TaxCategoryReference `json:"taxCategory,omitempty"`
	// Not set if a custom Shipping Method is used.
	ShippingMethod *ShippingMethodReference `json:"shippingMethod,omitempty"`
	// Information on how items are delivered to customers.
	Deliveries []Delivery `json:"deliveries"`
	// Discounted price of the Shipping Method.
	DiscountedPrice *DiscountedLineItemPrice `json:"discountedPrice,omitempty"`
	// Indicates whether the [ShippingMethod](ctp:api:type:ShippingMethod) referenced in this ShippingInfo is allowed for the Cart.
	ShippingMethodState ShippingMethodState `json:"shippingMethodState"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShippingInfo) MarshalJSON() ([]byte, error) {
	type Alias ShippingInfo
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

	if raw["deliveries"] == nil {
		delete(raw, "deliveries")
	}

	return json.Marshal(raw)

}

type ShippingRateInput struct {
	Type string `json:"type"`
}

/**
*	ShoppingListLineItems are Line Items that contain references to [ProductVariants](ctp:api:type:ProductVariant) in a [Product](ctp:api:type:Product).
*
*	In addition to standard [Reference Expansion](/general-concepts#reference-expansion), a ShoppingListLineItem offers expansion on `productSlug` and `variant`, defined with the query parameter `expand`.
*
 */
type ShoppingListLineItem struct {
	// Date and time (UTC) the ShoppingListLineItem was added to the ShoppingList.
	AddedAt time.Time `json:"addedAt"`
	// Custom Fields of the ShoppingListLineItem.
	Custom *CustomFields `json:"custom,omitempty"`
	// If the Product or Product Variant is deleted, `deactivatedAt` is the date and time (UTC) of deletion.
	//
	// This data is updated in an [eventual consistent manner](/general-concepts#eventual-consistency) when the Product Variant cannot be ordered anymore.
	DeactivatedAt *time.Time `json:"deactivatedAt,omitempty"`
	// Unique identifier of the ShoppingListLineItem.
	ID string `json:"id"`
	// User-defined identifier of the ShoppingListLineItem. It is unique per [ShoppingList](ctp:api:type:ShoppingList).
	Key *string `json:"key,omitempty"`
	// Name of the Product.
	//
	// This data is updated in an [eventual consistent manner](/general-concepts#eventual-consistency) when the Product's name changes.
	Name LocalizedString `json:"name"`
	// Unique identifier of a [Product](ctp:api:type:Product).
	ProductId string `json:"productId"`
	// The Product Type defining the Attributes of the [Product](ctp:api:type:Product).
	ProductType ProductTypeReference `json:"productType"`
	// Whether the related [Product](ctp:api:type:Product) is published or not.
	//
	// This data is updated in an [eventual consistent manner](/general-concepts#eventual-consistency) when the Product's published status changes.
	Published bool `json:"published"`
	// Number of Products in the ShoppingListLineItem.
	Quantity int `json:"quantity"`
	// `id` of the [ProductVariant](ctp:api:type:ProductVariant) the ShoppingListLineItem refers to. If not set, the ShoppingListLineItem refers to the Master Variant.
	VariantId *int `json:"variantId,omitempty"`
	// Data of the [ProductVariant](ctp:api:type:ProductVariant).  This data includes all the Product Attributes and Variant Attributes to ensure the full Attribute context of the Product Variant.
	//
	// Returned when expanded using `expand=lineItems[*].variant`. You cannot expand only a single element of the array.
	Variant *ProductVariant `json:"variant,omitempty"`
	// Slug of the current [ProductData](ctp:api:type:ProductData).
	//
	// Returned when expanded using `expand=lineItems[*].productSlug`. You cannot expand only a single element of the array.
	ProductSlug *LocalizedString `json:"productSlug,omitempty"`
}

/**
*	Provides a visual representation type for this field. It is only relevant for string-based field types like [CustomFieldStringType](ctp:api:type:CustomFieldStringType) and [CustomFieldLocalizedStringType](ctp:api:type:CustomFieldLocalizedStringType). Following values are supported:
*
 */
type TypeTextInputHint string

const (
	TypeTextInputHintSingleLine TypeTextInputHint = "SingleLine"
	TypeTextInputHintMultiLine  TypeTextInputHint = "MultiLine"
)

/**
*	Base polymorphic read-only money type that stores currency in cent precision or high precision, that is in sub-cents.
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
	}
	return nil, nil
}

/**
*	Object that stores money in cent amounts of a specific currency.
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

/**
*	Determines whether the selected [ShippingMethod](ctp:api:type:ShippingMethod) is allowed for the Cart. For more information, see [Predicates](/shipping-delivery-overview#predicates).
*
 */
type ShippingMethodState string

const (
	ShippingMethodStateDoesNotMatchCart ShippingMethodState = "DoesNotMatchCart"
	ShippingMethodStateMatchesCart      ShippingMethodState = "MatchesCart"
)

/**
*	Roles defining how an [Associate](ctp:api:type:Associate) can interact with a Business Unit.
*
 */
type AssociateRoleDeprecated string

const (
	AssociateRoleDeprecatedAdmin AssociateRoleDeprecated = "Admin"
	AssociateRoleDeprecatedBuyer AssociateRoleDeprecated = "Buyer"
)

/**
*	[KeyReference](ctp:api:type:KeyReference) to an [AssociateRole](ctp:api:type:AssociateRole).
*
 */
type AssociateRoleKeyReference struct {
	// Unique and immutable key of the referenced [AssociateRole](ctp:api:type:AssociateRole).
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AssociateRoleKeyReference) MarshalJSON() ([]byte, error) {
	type Alias AssociateRoleKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "associate-role", Alias: (*Alias)(&obj)})
}

type AttributeLevelEnum string

const (
	AttributeLevelEnumProduct AttributeLevelEnum = "Product"
	AttributeLevelEnumVariant AttributeLevelEnum = "Variant"
)

/**
*	Polymorphic base type that represents a postal address and contact details.
*	Depending on the read or write action, it can be either [Address](ctp:api:type:Address) or [AddressDraft](ctp:api:type:AddressDraft) that
*	only differ in the data type for the optional `custom` field.
*
 */
type BaseAddress struct {
	// Unique identifier of the Address.
	//
	// It is not recommended to set it manually since the API overwrites this ID when creating an Address for a [Customer](ctp:api:type:Customer).
	// Use `key` instead and omit this field from the request to let the API generate the ID for the Address.
	ID *string `json:"id,omitempty"`
	// User-defined identifier of the Address that must be unique when multiple addresses are referenced in [BusinessUnits](ctp:api:type:BusinessUnit), [Customers](ctp:api:type:Customer), and `itemShippingAddresses` (LineItem-specific addresses) of a [Cart](ctp:api:type:Cart), [Order](ctp:api:type:Order), [QuoteRequest](ctp:api:type:QuoteRequest), or [Quote](ctp:api:type:Quote).
	Key *string `json:"key,omitempty"`
	// Name of the country.
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

/**
*	Address type returned by read methods.
*	Optionally, the `custom` field can be present in addition to the fields of a [BaseAddress](ctp:api:type:BaseAddress).
*
 */
type Address struct {
	// Unique identifier of the Address.
	//
	// It is not recommended to set it manually since the API overwrites this ID when creating an Address for a [Customer](ctp:api:type:Customer).
	// Use `key` instead and omit this field from the request to let the API generate the ID for the Address.
	ID *string `json:"id,omitempty"`
	// User-defined identifier of the Address that must be unique when multiple addresses are referenced in [BusinessUnits](ctp:api:type:BusinessUnit), [Customers](ctp:api:type:Customer), and `itemShippingAddresses` (LineItem-specific addresses) of a [Cart](ctp:api:type:Cart), [Order](ctp:api:type:Order), [QuoteRequest](ctp:api:type:QuoteRequest), or [Quote](ctp:api:type:Quote).
	Key *string `json:"key,omitempty"`
	// Name of the country.
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

/**
*	[KeyReference](ctp:api:type:KeyReference) to a [BusinessUnit](ctp:api:type:BusinessUnit).
*
 */
type BusinessUnitKeyReference struct {
	// Unique and immutable key of the referenced [BusinessUnit](ctp:api:type:BusinessUnit).
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitKeyReference) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "business-unit", Alias: (*Alias)(&obj)})
}

/**
*	[Reference](ctp:api:type:Reference) to a [CartDiscount](ctp:api:type:CartDiscount).
*
 */
type CartDiscountReference struct {
	// Unique identifier of the referenced [CartDiscount](ctp:api:type:CartDiscount).
	ID string `json:"id"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountReference) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "cart-discount", Alias: (*Alias)(&obj)})
}

type CartDiscountTarget struct {
	Type string `json:"type"`
}

type CartDiscountValue struct {
	Type string `json:"type"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [Category](ctp:api:type:Category).
*
 */
type CategoryReference struct {
	// Unique identifier of the referenced [Category](ctp:api:type:Category).
	ID string `json:"id"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CategoryReference) MarshalJSON() ([]byte, error) {
	type Alias CategoryReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "category", Alias: (*Alias)(&obj)})
}

/**
*	[Reference](ctp:api:type:Reference) to a [Channel](ctp:api:type:Channel).
*
 */
type ChannelReference struct {
	// Unique identifier of the referenced [Channel](ctp:api:type:Channel).
	ID string `json:"id"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChannelReference) MarshalJSON() ([]byte, error) {
	type Alias ChannelReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "channel", Alias: (*Alias)(&obj)})
}

/**
*	Determines if Cart Discounts can be applied to a Custom Line Item in the Cart.
*
 */
type CustomLineItemPriceMode string

const (
	CustomLineItemPriceModeStandard CustomLineItemPriceMode = "Standard"
	CustomLineItemPriceModeExternal CustomLineItemPriceMode = "External"
)

/**
*	Information about recurring orders and frequencies.
*
 */
type CustomLineItemRecurrenceInfo struct {
	// [Reference](ctp:api:type:Reference) to a RecurrencePolicy.
	RecurrencePolicy RecurrencePolicyReference `json:"recurrencePolicy"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [CustomerGroup](ctp:api:type:CustomerGroup).
*
 */
type CustomerGroupReference struct {
	// Unique identifier of the referenced [CustomerGroup](ctp:api:type:CustomerGroup).
	ID string `json:"id"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerGroupReference) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "customer-group", Alias: (*Alias)(&obj)})
}

/**
*	[Reference](ctp:api:type:Reference) to a [Customer](ctp:api:type:Customer).
*
 */
type CustomerReference struct {
	// Unique identifier of the referenced [Customer](ctp:api:type:Customer).
	ID string `json:"id"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerReference) MarshalJSON() ([]byte, error) {
	type Alias CustomerReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "customer", Alias: (*Alias)(&obj)})
}

/**
*	[Reference](ctp:api:type:Reference) to a [DiscountCode](ctp:api:type:DiscountCode).
*
 */
type DiscountCodeReference struct {
	// Unique identifier of the referenced [DiscountCode](ctp:api:type:DiscountCode).
	ID string `json:"id"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeReference) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "discount-code", Alias: (*Alias)(&obj)})
}

type FieldContainer map[string]interface{}

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
		obj := GeoLocation{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type GeoLocation struct {
	Coordinates []int `json:"coordinates"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj GeoLocation) MarshalJSON() ([]byte, error) {
	type Alias GeoLocation
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "Point", Alias: (*Alias)(&obj)})
}

/**
*	Indicates how Line Items in a Cart are tracked.
*
 */
type InventoryMode string

const (
	InventoryModeNone           InventoryMode = "None"
	InventoryModeTrackOnly      InventoryMode = "TrackOnly"
	InventoryModeReserveOnOrder InventoryMode = "ReserveOnOrder"
)

/**
*	Indicates how a Line Item was added to a Cart.
*
 */
type LineItemMode string

const (
	LineItemModeStandard     LineItemMode = "Standard"
	LineItemModeGiftLineItem LineItemMode = "GiftLineItem"
)

/**
*	This mode indicates how the price is set for the Line Item.
*
 */
type LineItemPriceMode string

const (
	LineItemPriceModePlatform      LineItemPriceMode = "Platform"
	LineItemPriceModeExternalPrice LineItemPriceMode = "ExternalPrice"
	LineItemPriceModeExternalTotal LineItemPriceMode = "ExternalTotal"
)

/**
*	Information about recurring orders and frequencies.
*
 */
type LineItemRecurrenceInfo struct {
	// [Reference](ctp:api:type:Reference) to a RecurrencePolicy.
	RecurrencePolicy RecurrencePolicyReference `json:"recurrencePolicy"`
	// Indicates how the price of a line item will be selected during order creation.
	PriceSelectionMode PriceSelectionMode `json:"priceSelectionMode"`
}

type MethodTaxRate struct {
	// User-defined unique identifier of the Shipping Method in a Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingMethodKey string `json:"shippingMethodKey"`
	// Tax Rate for the Shipping Method.
	TaxRate *TaxRate `json:"taxRate,omitempty"`
}

type MethodTaxedPrice struct {
	// User-defined unique identifier of the [Shipping Method](ctp:api:type:ShippingMethod) in a Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingMethodKey string `json:"shippingMethodKey"`
	// Total taxed price based on the quantity of the Line Item or Custom Line Item assigned to the Shipping Method identified by `shippingMethodKey`.
	TaxedPrice *TaxedItemPrice `json:"taxedPrice,omitempty"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [Payment](ctp:api:type:Payment).
*
 */
type PaymentReference struct {
	// Unique identifier of the referenced [Payment](ctp:api:type:Payment).
	ID string `json:"id"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentReference) MarshalJSON() ([]byte, error) {
	type Alias PaymentReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "payment", Alias: (*Alias)(&obj)})
}

/**
*	Indicates how the price of a [Line Item](ctp:api:type:LineItem) or [Custom Line Item](ctp:api:type:CustomLineItem) is selected during Order creation.
*
 */
type PriceSelectionMode string

const (
	PriceSelectionModeFixed   PriceSelectionMode = "Fixed"
	PriceSelectionModeDynamic PriceSelectionMode = "Dynamic"
)

/**
*	A Price tier is selected instead of the default Price when a certain quantity of the [ProductVariant](ctp:api:type:ProductVariant) is [added to a Cart](/projects/carts#add-lineitem) and ordered.
*	_For   If no Price tier is found for the Order quantity, the base Price is used.
*	A Price tier is applied for the entire quantity of a Product Variant put as [LineItem](/projects/carts#lineitem) in a Cart as soon as the minimum quantity for the Price tier is reached.
*	The Price tier is applied per Line Item of the Product Variant. If, for example, the same Product Variant appears in the same Cart as several Line Items, (what can be achieved by different values of a Custom Field on the Line Items) for each Line Item the minimum quantity must be reached to get the Price tier.
*
 */
type PriceTier struct {
	// Minimum quantity this Price tier is valid for.
	//
	// The minimum quantity is always greater than or equal to 2. The base Price is interpreted as valid for a minimum quantity equal to 1.
	// A [Price](ctp:api:type:Price) or [StandalonePrice](ctp:api:type:StandalonePrice) cannot contain more than one tier with the same `minimumQuantity`.
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
*	[Reference](ctp:api:type:Reference) to a [ProductDiscount](ctp:api:type:ProductDiscount).
*
 */
type ProductDiscountReference struct {
	// Unique identifier of the referenced [ProductDiscount](ctp:api:type:ProductDiscount).
	ID string `json:"id"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountReference) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "product-discount", Alias: (*Alias)(&obj)})
}

type ProductDiscountValue struct {
	Type string `json:"type"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [Product](ctp:api:type:Product).
*
 */
type ProductReference struct {
	// Unique identifier of the referenced [Product](ctp:api:type:Product).
	ID string `json:"id"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductReference) MarshalJSON() ([]byte, error) {
	type Alias ProductReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "product", Alias: (*Alias)(&obj)})
}

/**
*	[Reference](ctp:api:type:Reference) to a [ProductSelection](ctp:api:type:ProductSelection).
*
 */
type ProductSelectionReference struct {
	// Unique identifier of the referenced [ProductSelection](ctp:api:type:ProductSelection).
	ID string `json:"id"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSelectionReference) MarshalJSON() ([]byte, error) {
	type Alias ProductSelectionReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "product-selection", Alias: (*Alias)(&obj)})
}

/**
*	[Reference](ctp:api:type:Reference) to a [ProductType](ctp:api:type:ProductType).
*
 */
type ProductTypeReference struct {
	// Unique identifier of the referenced [ProductType](ctp:api:type:ProductType).
	ID string `json:"id"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTypeReference) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "product-type", Alias: (*Alias)(&obj)})
}

/**
*	A concrete sellable good for which inventory can be tracked. Product Variants are generally mapped to specific SKUs.
*
 */
type ProductVariant struct {
	// A unique, sequential identifier of the Product Variant within the Product.
	ID int `json:"id"`
	// User-defined unique SKU of the Product Variant.
	Sku *string `json:"sku,omitempty"`
	// User-defined unique identifier of the ProductVariant.
	//
	// This is different from [Product](ctp:api:type:Product) `key`.
	Key *string `json:"key,omitempty"`
	// The Embedded Prices of the Product Variant.
	// Cannot contain two Prices of the same Price scope (with same currency, country, Customer Group, Channel, `validFrom` and `validUntil`).
	Prices []Price `json:"prices"`
	// Variant Attributes according to the respective [AttributeDefinition](ctp:api:type:AttributeDefinition).
	Attributes []Attribute `json:"attributes"`
	// Only available when [price selection](/../api/pricing-and-discounts-overview#price-selection) is used.
	// Cannot be used in a [Query Predicate](ctp:api:type:QueryPredicate).
	Price *Price `json:"price,omitempty"`
	// Images of the Product Variant.
	Images []Image `json:"images"`
	// Media assets of the Product Variant.
	Assets []Asset `json:"assets"`
	// Set if the Product Variant is tracked by [Inventory](ctp:api:type:InventoryEntry).
	// Can be used as an optimization to reduce calls to the Inventory service.
	// May not contain the latest Inventory State (it is [eventually consistent](/general-concepts#eventual-consistency)).
	Availability *ProductVariantAvailability `json:"availability,omitempty"`
	// `true` if the Product Variant matches the search query.
	// Only available in response to a [Product Projection Search](ctp:api:type:ProductProjectionSearch) request.
	IsMatchingVariant *bool `json:"isMatchingVariant,omitempty"`
	// Only available in response to a [Product Projection Search](ctp:api:type:ProductProjectionSearch) request
	// with [Product price selection](/../api/pricing-and-discounts-overview#product-price-selection).
	// Can be used to sort, [filter](ctp:api:type:ProductProjectionSearchFilterScopedPrice), and facet.
	ScopedPrice *ScopedPrice `json:"scopedPrice,omitempty"`
	// Only available in response to a [Product Projection Search](ctp:api:type:ProductProjectionSearchFilterScopedPrice) request
	// with [Product price selection](/../api/pricing-and-discounts-overview#product-price-selection).
	ScopedPriceDiscounted *bool `json:"scopedPriceDiscounted,omitempty"`
	// Only available when [Product price selection](/../api/pricing-and-discounts-overview#product-price-selection) is used.
	// Cannot be used in a [Query Predicate](ctp:api:type:QueryPredicate).
	RecurrencePrices []Price `json:"recurrencePrices"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductVariant) MarshalJSON() ([]byte, error) {
	type Alias ProductVariant
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

	if raw["prices"] == nil {
		delete(raw, "prices")
	}

	if raw["attributes"] == nil {
		delete(raw, "attributes")
	}

	if raw["images"] == nil {
		delete(raw, "images")
	}

	if raw["assets"] == nil {
		delete(raw, "assets")
	}

	if raw["recurrencePrices"] == nil {
		delete(raw, "recurrencePrices")
	}

	return json.Marshal(raw)

}

/**
*	[Reference](ctp:api:type:Reference) to a [RecurrencePolicy](ctp:api:type:RecurrencePolicy).
*
 */
type RecurrencePolicyReference struct {
	// Unique identifier of the referenced [RecurrencePolicy](ctp:api:type:RecurrencePolicy).
	ID string `json:"id"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RecurrencePolicyReference) MarshalJSON() ([]byte, error) {
	type Alias RecurrencePolicyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "recurrence-policy", Alias: (*Alias)(&obj)})
}

/**
*	With Types, you can model your own Custom Fields on the following resources and data types.
*
 */
type ResourceTypeId string

const (
	ResourceTypeIdAddress                     ResourceTypeId = "address"
	ResourceTypeIdAsset                       ResourceTypeId = "asset"
	ResourceTypeIdApprovalFlow                ResourceTypeId = "approval-flow"
	ResourceTypeIdApprovalRule                ResourceTypeId = "approval-rule"
	ResourceTypeIdAssociateRole               ResourceTypeId = "associate-role"
	ResourceTypeIdBusinessUnit                ResourceTypeId = "business-unit"
	ResourceTypeIdCartDiscount                ResourceTypeId = "cart-discount"
	ResourceTypeIdCategory                    ResourceTypeId = "category"
	ResourceTypeIdChannel                     ResourceTypeId = "channel"
	ResourceTypeIdCustomer                    ResourceTypeId = "customer"
	ResourceTypeIdCustomerGroup               ResourceTypeId = "customer-group"
	ResourceTypeIdCustomLineItem              ResourceTypeId = "custom-line-item"
	ResourceTypeIdDiscountCode                ResourceTypeId = "discount-code"
	ResourceTypeIdInventoryEntry              ResourceTypeId = "inventory-entry"
	ResourceTypeIdLineItem                    ResourceTypeId = "line-item"
	ResourceTypeIdOrder                       ResourceTypeId = "order"
	ResourceTypeIdOrderEdit                   ResourceTypeId = "order-edit"
	ResourceTypeIdOrderDelivery               ResourceTypeId = "order-delivery"
	ResourceTypeIdOrderParcel                 ResourceTypeId = "order-parcel"
	ResourceTypeIdOrderReturnItem             ResourceTypeId = "order-return-item"
	ResourceTypeIdPayment                     ResourceTypeId = "payment"
	ResourceTypeIdPaymentInterfaceInteraction ResourceTypeId = "payment-interface-interaction"
	ResourceTypeIdProductPrice                ResourceTypeId = "product-price"
	ResourceTypeIdProductSelection            ResourceTypeId = "product-selection"
	ResourceTypeIdProductTailoring            ResourceTypeId = "product-tailoring"
	ResourceTypeIdQuote                       ResourceTypeId = "quote"
	ResourceTypeIdReview                      ResourceTypeId = "review"
	ResourceTypeIdRecurringOrder              ResourceTypeId = "recurring-order"
	ResourceTypeIdShipping                    ResourceTypeId = "shipping"
	ResourceTypeIdShippingMethod              ResourceTypeId = "shipping-method"
	ResourceTypeIdShoppingList                ResourceTypeId = "shopping-list"
	ResourceTypeIdShoppingListTextLineItem    ResourceTypeId = "shopping-list-text-line-item"
	ResourceTypeIdStandalonePrice             ResourceTypeId = "standalone-price"
	ResourceTypeIdStore                       ResourceTypeId = "store"
	ResourceTypeIdTransaction                 ResourceTypeId = "transaction"
)

/**
*	Scoped Price is contained in a [ProductVariant](ctp:api:type:ProductVariant) which is returned in response to a
*	[Product Projection Search](ctp:api:type:ProductProjectionSearchFilterScopedPrice) request when [Scoped Price Search](/../api/pricing-and-discounts-overview#scoped-price-search) is used.
*
 */
type ScopedPrice struct {
	// Platform-generated unique identifier of the Price.
	ID string `json:"id"`
	// Original value of the Price.
	Value TypedMoney `json:"value"`
	// If available, either the original price `value` or `discounted` value.
	CurrentValue TypedMoney `json:"currentValue"`
	// Country code of the geographic location.
	Country *string `json:"country,omitempty"`
	// Reference to a CustomerGroup.
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
	// Reference to a Channel.
	Channel *ChannelReference `json:"channel,omitempty"`
	// Date and time from which the Price is valid.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Date and time until which the Price is valid.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
	// Is set when a matching [ProductDiscount](ctp:api:type:ProductDiscount) exists. If set, the [Cart](ctp:api:type:Cart) uses the discounted value for the [Cart Price calculation](ctp:api:type:CartAddLineItemAction).
	//
	// When a [relative Product Discount](ctp:api:type:ProductDiscountValueRelative) is applied and the fractional part of the discounted Price is 0.5, the discounted Price is [rounded half down](https://en.wikipedia.org/wiki/Rounding#Rounding_half_down) in favor of the Customer.
	Discounted *DiscountedPrice `json:"discounted,omitempty"`
	// Custom Fields of the Price.
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
*	[Reference](ctp:api:type:Reference) to a [ShippingMethod](ctp:api:type:ShippingMethod).
*
 */
type ShippingMethodReference struct {
	// Unique identifier of the referenced [ShippingMethod](ctp:api:type:ShippingMethod).
	ID string `json:"id"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShippingMethodReference) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "shipping-method", Alias: (*Alias)(&obj)})
}

/**
*	[Reference](ctp:api:type:Reference) to a [State](ctp:api:type:State).
*
 */
type StateReference struct {
	// Unique identifier of the referenced [State](ctp:api:type:State).
	ID string `json:"id"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StateReference) MarshalJSON() ([]byte, error) {
	type Alias StateReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "state", Alias: (*Alias)(&obj)})
}

/**
*	[KeyReference](ctp:api:type:KeyReference) to a [Store](ctp:api:type:Store).
*
 */
type StoreKeyReference struct {
	// Unique and immutable key of the referenced [Store](ctp:api:type:Store).
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
*	[Reference](ctp:api:type:Reference) to a [TaxCategory](ctp:api:type:TaxCategory).
*
 */
type TaxCategoryReference struct {
	// Unique identifier of the referenced [TaxCategory](ctp:api:type:TaxCategory).
	ID string `json:"id"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TaxCategoryReference) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "tax-category", Alias: (*Alias)(&obj)})
}

/**
*	The tax portions are calculated from the [TaxRates](ctp:api:type:TaxRate).
*	If a Tax Rate has [SubRates](ctp:api:type:SubRate), they are used and can be identified by name.
*	Tax portions from Line Items with the same `rate` and `name` are accumulated to the same tax portion.
*
 */
type TaxPortion struct {
	// Name of the tax portion.
	Name *string `json:"name,omitempty"`
	// A number in the range 0-1.
	Rate float64 `json:"rate"`
	// Money value of the tax portion.
	Amount CentPrecisionMoney `json:"amount"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [Type](ctp:api:type:Type).
*
 */
type TypeReference struct {
	// Unique identifier of the referenced [Type](ctp:api:type:Type).
	ID string `json:"id"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TypeReference) MarshalJSON() ([]byte, error) {
	type Alias TypeReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "type", Alias: (*Alias)(&obj)})
}

/**
*	These objects represent information about which [API Client](/../api/projects/api-clients) created or modified a resource. For more information, see [Client Logging](/../api/general-concepts#client-logging).
*
 */
type ClientLogging struct {
	// `id` of the [API Client](ctp:api:type:ApiClient) which created the resource.
	ClientId *string `json:"clientId,omitempty"`
	// [External user ID](/../api/general-concepts#external-user-ids) provided by `X-External-User-ID` HTTP Header.
	ExternalUserId *string `json:"externalUserId,omitempty"`
	// Indicates the [Customer](ctp:api:type:Customer) who modified the resource using a token from the [password flow](/authorization#password-flow).
	Customer *CustomerReference `json:"customer,omitempty"`
	// Indicates that the resource was modified during an [anonymous session](ctp:api:type:AnonymousSession) with the logged ID.
	AnonymousId *string `json:"anonymousId,omitempty"`
	// Indicates the [Customer](ctp:api:type:Customer) who created or modified the resource in the context of a [Business Unit](ctp:api:type:BusinessUnit). Only available for [B2B](/../offering/composable-commerce#composable-commerce-for-b2b)-enabled Projects when an Associate acts on behalf of a company using the [associate endpoints](/associates-overview#on-the-associate-endpoints).
	Associate *CustomerReference `json:"associate,omitempty"`
}

/**
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [Zone](ctp:api:type:Zone). Either `id` or `key` is required. If both are set, an [InvalidJsonInput](/../api/errors#invalidjsoninput) error is returned.
*
 */
type ZoneResourceIdentifier struct {
	// Unique identifier of the referenced [Zone](ctp:api:type:Zone). Required if `key` is absent.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced [Zone](ctp:api:type:Zone). Required if `id` is absent.
	Key *string `json:"key,omitempty"`
	// Type of resource the value should reference. Supported resource type identifiers are:
	TypeId ReferenceTypeId `json:"typeId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ZoneResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias ZoneResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"null"`
		*Alias
	}{Action: "zone", Alias: (*Alias)(&obj)})
}
