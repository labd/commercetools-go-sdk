package platform

// Generated file, please do not change!!!

/**
*	A lightweight representation of a [Product's](ctp:api:type:Product) Variants with only the requested Attributes and minimal availability data.
*	Designed for building attribute selectors on product detail pages (PDPs).
*
 */
type VariantAttributes struct {
	// Unique identifier of the [Product](ctp:api:type:Product).
	ProductId string `json:"productId"`
	// User-defined unique identifier of the [Product](ctp:api:type:Product).
	ProductKey *string `json:"productKey,omitempty"`
	// Metadata for the requested Attributes, derived from the [ProductType](ctp:api:type:ProductType).
	// Attributes not found in the ProductType are silently omitted.
	Attributes []VariantAttributesAttributeMetadata `json:"attributes"`
	// All Variants of the Product with their requested Attributes and availability.
	Variants []VariantAttributesVariant `json:"variants"`
}

/**
*	Metadata about a requested [Attribute](ctp:api:type:AttributeDefinition), derived from the [ProductType](ctp:api:type:ProductType).
*
 */
type VariantAttributesAttributeMetadata struct {
	// Name of the Attribute as defined in the [ProductType](ctp:api:type:ProductType).
	Name string `json:"name"`
	// Localized label of the Attribute as defined in the [ProductType](ctp:api:type:ProductType).
	// Subject to [locale projection](/api/projects/productProjections#locales) if `localeProjection` is specified.
	Label LocalizedString `json:"label"`
	// Name of the [AttributeType](ctp:api:type:AttributeType), for example `boolean`, `text`, `ltext`, `enum`, `lenum`, `number`, `money`, `date`, `time`, `datetime`, `reference`, `set`, or `nested`.
	Type string `json:"type"`
}

/**
*	Lightweight availability information for a Variant. Present only if [InventoryEntries](ctp:api:type:InventoryEntry) exist for the Variant's SKU.
*
 */
type VariantAttributesAvailability struct {
	// Whether the Variant is in stock.
	IsOnStock *bool `json:"isOnStock,omitempty"`
	// Number of days to restock the Variant once it is out of stock.
	RestockableInDays *int `json:"restockableInDays,omitempty"`
	// Number of items of this Variant that are in stock.
	AvailableQuantity *int `json:"availableQuantity,omitempty"`
	// Per-channel availability for the requested supply [Channels](ctp:api:type:Channel).
	// Present only when the `filter[supplyChannels]` query parameter is provided.
	Channels *VariantAttributesChannelAvailabilityMap `json:"channels,omitempty"`
}

/**
*	Availability of a Variant in a specific supply [Channel](ctp:api:type:Channel).
*
 */
type VariantAttributesChannelAvailability struct {
	// Whether the Variant is in stock in the specified [Channel](ctp:api:type:Channel).
	IsOnStock *bool `json:"isOnStock,omitempty"`
	// Number of days to restock the Variant once it is out of stock in the specified [Channel](ctp:api:type:Channel).
	RestockableInDays *int `json:"restockableInDays,omitempty"`
	// Number of items of this Variant that are in stock in the specified [Channel](ctp:api:type:Channel).
	AvailableQuantity *int `json:"availableQuantity,omitempty"`
}

/**
*	JSON object where the keys are supply [Channel](ctp:api:type:Channel) `id`, and the values are [VariantAttributesChannelAvailability](ctp:api:type:VariantAttributesChannelAvailability).
*
 */
type VariantAttributesChannelAvailabilityMap map[string]VariantAttributesChannelAvailability

/**
*	A Variant with its requested Attributes and lightweight availability information.
*
 */
type VariantAttributesVariant struct {
	// Unique identifier of the [Variant](ctp:api:type:Variant).
	ID string `json:"id"`
	// SKU of the [Variant](ctp:api:type:Variant).
	Sku *string `json:"sku,omitempty"`
	// Key of the [Variant](ctp:api:type:Variant).
	Key *string `json:"key,omitempty"`
	// Availability information for this variant.
	// Present only if [InventoryEntries](ctp:api:type:InventoryEntry) exist for the variant's SKU.
	Availability *VariantAttributesAvailability `json:"availability,omitempty"`
	// Requested [Attributes](ctp:api:type:Attribute) of the Variant.
	// Only Variant-level Attributes that exist on this Variant and were requested via `filter[attributes]` are included.
	Attributes []Attribute `json:"attributes"`
}
