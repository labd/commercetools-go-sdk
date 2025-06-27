package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type ShoppingList struct {
	// Unique identifier of the ShoppingList.
	ID string `json:"id"`
	// Current version of the ShoppingList.
	Version int `json:"version"`
	// Date and time (UTC) the ShoppingList was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the ShoppingList was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Name of the ShoppingList.
	Name LocalizedString `json:"name"`
	// User-defined unique identifier of the ShoppingList.
	Key *string `json:"key,omitempty"`
	// Reference to a [Customer](ctp:api:type:Customer) associated with the ShoppingList.
	Customer *CustomerReference `json:"customer,omitempty"`
	// Human-readable identifiers usually used as deep-link URL to the related ShoppingList.
	// Each slug is unique across a Project, but a ShoppingList can have the same slug for different languages.
	// The slug must match the pattern `[a-zA-Z0-9_-]{2,256}`. For [good performance](/predicates/query#performance-considerations), indexes are provided for the first 15 `languages` set on the [Project](ctp:api:type:Project).
	Slug *LocalizedString `json:"slug,omitempty"`
	// Description of the ShoppingList.
	Description *LocalizedString `json:"description,omitempty"`
	// Line Items (containing Products) of the ShoppingList.
	LineItems []ShoppingListLineItem `json:"lineItems"`
	// Line Items (containing text values) of the ShoppingList.
	TextLineItems []TextLineItem `json:"textLineItems"`
	// Number of days after the last modification before a ShoppingList is deleted. If not set, the [default value](ctp:api:type:ShoppingListsConfiguration) configured in the [Project](ctp:api:type:Project) is used.
	DeleteDaysAfterLastModification *int `json:"deleteDaysAfterLastModification,omitempty"`
	// Identifies ShoppingLists belonging to an [anonymous session](ctp:api:type:AnonymousSession).
	AnonymousId *string `json:"anonymousId,omitempty"`
	// Store to which the ShoppingList is assigned.
	Store *StoreKeyReference `json:"store,omitempty"`
	// [Reference](ctp:api:type:Reference) to the Business Unit the Shopping List belongs to. Only available for [B2B](/../offering/composable-commerce#composable-commerce-for-b2b)-enabled Projects.
	BusinessUnit *BusinessUnitKeyReference `json:"businessUnit,omitempty"`
	// Custom Fields defined for the ShoppingList.
	Custom *CustomFields `json:"custom,omitempty"`
	// IDs and references that last modified the ShoppingList.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// IDs and references that created the ShoppingList.
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
}

type ShoppingListDraft struct {
	// Name of the ShoppingList.
	Name LocalizedString `json:"name"`
	// Human-readable identifiers usually used as deep-link URL to the related ShoppingList.
	// Each slug is unique across a Project, but a ShoppingList can have the same slug for different languages.
	// The slug must match the pattern `[a-zA-Z0-9_-]{2,256}`.
	Slug *LocalizedString `json:"slug,omitempty"`
	// The [Customer](ctp:api:type:Customer) the ShoppingList should be associated to.
	Customer *CustomerResourceIdentifier `json:"customer,omitempty"`
	// User-defined unique identifier for the ShoppingList.
	Key *string `json:"key,omitempty"`
	// Description of the ShoppingList.
	Description *LocalizedString `json:"description,omitempty"`
	// Identifies ShoppingLists belonging to an [anonymous session](ctp:api:type:AnonymousSession).
	AnonymousId *string `json:"anonymousId,omitempty"`
	// Number of days after the last modification before a ShoppingList is deleted. If not set, the [default value](ctp:api:type:ShoppingListsConfiguration) configured in the [Project](ctp:api:type:Project) is used.
	DeleteDaysAfterLastModification *int `json:"deleteDaysAfterLastModification,omitempty"`
	// Line Items (containing Products) to add to the ShoppingList.
	LineItems []ShoppingListLineItemDraft `json:"lineItems"`
	// Line Items (containing text values) to add to the ShoppingList.
	TextLineItems []TextLineItemDraft `json:"textLineItems"`
	// Assigns the new ShoppingList to the [Store](ctp:api:type:Store).
	Store *StoreResourceIdentifier `json:"store,omitempty"`
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) of the Business Unit the Shopping List should belong to. When the `customer` of the Shopping List is set, the [Customer](ctp:api:type:Customer) must be an [Associate](ctp:api:type:Associate) of the Business Unit. Only available for [B2B](/../offering/composable-commerce#composable-commerce-for-b2b)-enabled Projects.
	BusinessUnit *BusinessUnitResourceIdentifier `json:"businessUnit,omitempty"`
	// Custom Fields defined for the ShoppingList.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShoppingListDraft) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListDraft
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

	if raw["lineItems"] == nil {
		delete(raw, "lineItems")
	}

	if raw["textLineItems"] == nil {
		delete(raw, "textLineItems")
	}

	return json.Marshal(raw)

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
*	The [ProductVariant](ctp:api:type:ProductVariant) to be included in the ShoppingListLineItem must be specified using the `productID` and `variantID`, or by the `sku`.
*
 */
type ShoppingListLineItemDraft struct {
	// User-defined identifier of the ShoppingListLineItem. Must be unique per [ShoppingList](ctp:api:type:ShoppingList).
	Key *string `json:"key,omitempty"`
	// Unique identifier of a [Product](ctp:api:type:Product).
	ProductId *string `json:"productId,omitempty"`
	// `id` of the [ProductVariant](ctp:api:type:ProductVariant). If not set, the ShoppingListLineItem refers to the Master Variant.
	VariantId *int `json:"variantId,omitempty"`
	// `sku` of the [ProductVariant](ctp:api:type:ProductVariant).
	Sku *string `json:"sku,omitempty"`
	// Date and time the ShoppingListLineItem is added to the [ShoppingList](ctp:api:type:ShoppingList). If not set, the current date and time (UTC) is used.
	AddedAt *time.Time `json:"addedAt,omitempty"`
	// Custom Fields of the ShoppingListLineItem.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
	// Number of Products in the ShoppingListLineItem.
	Quantity *int `json:"quantity,omitempty"`
}

type ShoppingListPagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/../api/general-concepts#strong-consistency).
	// This field is returned by default.
	// For improved performance, calculating this field can be deactivated by using the query parameter `withTotal=false`.
	// When the results are filtered with a [Query Predicate](/../api/predicates/query), `total` is subject to a [limit](/../api/limits#queries).
	Total *int `json:"total,omitempty"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// [ShoppingLists](ctp:api:type:ShoppingList) matching the query.
	Results []ShoppingList `json:"results"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [ShoppingList](ctp:api:type:ShoppingList).
*
 */
type ShoppingListReference struct {
	// Unique identifier of the referenced [ShoppingList](ctp:api:type:ShoppingList).
	ID string `json:"id"`
	// Contains the representation of the expanded ShoppingList. Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for ShoppingLists.
	Obj *ShoppingList `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShoppingListReference) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "shopping-list", Alias: (*Alias)(&obj)})
}

/**
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [ShoppingList](ctp:api:type:ShoppingList). Either `id` or `key` is required. If both are set, an [InvalidJsonInput](/../api/errors#invalidjsoninput) error is returned.
*
 */
type ShoppingListResourceIdentifier struct {
	// Unique identifier of the referenced [ShoppingList](ctp:api:type:ShoppingList). Required if `key` is absent.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced [ShoppingList](ctp:api:type:ShoppingList). Required if `id` is absent.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShoppingListResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "shopping-list", Alias: (*Alias)(&obj)})
}

type ShoppingListUpdate struct {
	// Expected version of the ShoppingList on which the changes should be applied.
	// If the expected version does not match the actual version, a [ConcurrentModification](ctp:api:type:ConcurrentModificationError) error will be returned.
	Version int `json:"version"`
	// List of update actions to be performed on the ShoppingList.
	Actions []ShoppingListUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ShoppingListUpdate) UnmarshalJSON(data []byte) error {
	type Alias ShoppingListUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorShoppingListUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type ShoppingListUpdateAction interface{}

func mapDiscriminatorShoppingListUpdateAction(input interface{}) (ShoppingListUpdateAction, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["action"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'action'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "addLineItem":
		obj := ShoppingListAddLineItemAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addTextLineItem":
		obj := ShoppingListAddTextLineItemAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeLineItemQuantity":
		obj := ShoppingListChangeLineItemQuantityAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeLineItemsOrder":
		obj := ShoppingListChangeLineItemsOrderAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeName":
		obj := ShoppingListChangeNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeTextLineItemName":
		obj := ShoppingListChangeTextLineItemNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeTextLineItemQuantity":
		obj := ShoppingListChangeTextLineItemQuantityAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeTextLineItemsOrder":
		obj := ShoppingListChangeTextLineItemsOrderAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeLineItem":
		obj := ShoppingListRemoveLineItemAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeTextLineItem":
		obj := ShoppingListRemoveTextLineItemAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAnonymousId":
		obj := ShoppingListSetAnonymousIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setBusinessUnit":
		obj := ShoppingListSetBusinessUnitAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomField":
		obj := ShoppingListSetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomType":
		obj := ShoppingListSetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomer":
		obj := ShoppingListSetCustomerAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDeleteDaysAfterLastModification":
		obj := ShoppingListSetDeleteDaysAfterLastModificationAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDescription":
		obj := ShoppingListSetDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := ShoppingListSetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLineItemCustomField":
		obj := ShoppingListSetLineItemCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLineItemCustomType":
		obj := ShoppingListSetLineItemCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setSlug":
		obj := ShoppingListSetSlugAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setStore":
		obj := ShoppingListSetStoreAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setTextLineItemCustomField":
		obj := ShoppingListSetTextLineItemCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setTextLineItemCustomType":
		obj := ShoppingListSetTextLineItemCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setTextLineItemDescription":
		obj := ShoppingListSetTextLineItemDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

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

type TextLineItemDraft struct {
	// User-defined unique identifier of the TextLineItem. Must be unique per [ShoppingList](ctp:api:type:ShoppingList).
	Key *string `json:"key,omitempty"`
	// Date and time the TextLineItem is added to the [ShoppingList](ctp:api:type:ShoppingList). If not set, the current date and time (UTC) is used.
	AddedAt *time.Time `json:"addedAt,omitempty"`
	// Custom Fields for the TextLineItem.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
	// Description of the TextLineItem.
	Description *LocalizedString `json:"description,omitempty"`
	// Name of the TextLineItem.
	Name LocalizedString `json:"name"`
	// Number of entries in the TextLineItem.
	Quantity *int `json:"quantity,omitempty"`
}

/**
*	The [ProductVariant](ctp:api:type:ProductVariant) to be included in the ShoppingListLineItem must be specified using the `productID` and `variantID`, or by the `sku`.
*	If the ShoppingList already contains a ShoppingListLineItem for the same Product Variant with the same Custom Fields, then only the quantity of the existing ShoppingListLineItem is increased.
*	A ShoppingListLineItem with an empty `variantId` is not considered the same as a ShoppingListLineItem with a `variantId` currently referring to the Master Variant.
*
*	Product Attributes are merged with Variant Attributes to ensure the full Attribute context of the Product Variant.
*
*	Produces the [Shopping List Line Item Added](ctp:api:type:ShoppingListLineItemAddedMessage) Message.
*
 */
type ShoppingListAddLineItemAction struct {
	// User-defined identifier of the ShoppingListLineItem. Must be unique per [ShoppingList](ctp:api:type:ShoppingList).
	Key *string `json:"key,omitempty"`
	// `sku` of the [ProductVariant](ctp:api:type:ProductVariant).
	Sku *string `json:"sku,omitempty"`
	// Unique identifier of a [Product](ctp:api:type:Product).
	ProductId *string `json:"productId,omitempty"`
	// `id` of the [ProductVariant](ctp:api:type:ProductVariant). If not set, the ShoppingListLineItem refers to the Master Variant.
	VariantId *int `json:"variantId,omitempty"`
	// Number of Products in the ShoppingListLineItem.
	Quantity *int `json:"quantity,omitempty"`
	// Date and time the ShoppingListLineItem is added to the [ShoppingList](ctp:api:type:ShoppingList). If not set, the current date and time (UTC) is used.
	AddedAt *time.Time `json:"addedAt,omitempty"`
	// Custom Fields defined for the ShoppingListLineItem.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShoppingListAddLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListAddLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addLineItem", Alias: (*Alias)(&obj)})
}

type ShoppingListAddTextLineItemAction struct {
	// Name of the TextLineItem.
	Name LocalizedString `json:"name"`
	// User-defined identifier of the TextLineItem. Must be unique per [ShoppingList](ctp:api:type:ShoppingList).
	Key *string `json:"key,omitempty"`
	// Description of the TextLineItem.
	Description *LocalizedString `json:"description,omitempty"`
	// Number of entries in the TextLineItem.
	Quantity *int `json:"quantity,omitempty"`
	// Date and time the TextLineItem is added to the [ShoppingList](ctp:api:type:ShoppingList). If not set, the current date and time (UTC) is used.
	AddedAt *time.Time `json:"addedAt,omitempty"`
	// Custom Fields defined for the TextLineItem.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShoppingListAddTextLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListAddTextLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addTextLineItem", Alias: (*Alias)(&obj)})
}

type ShoppingListChangeLineItemQuantityAction struct {
	// The `id` of the [ShoppingListLineItem](ctp:api:type:ShoppingListLineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemId *string `json:"lineItemId,omitempty"`
	// The `key` of the [ShoppingListLineItem](ctp:api:type:ShoppingListLineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemKey *string `json:"lineItemKey,omitempty"`
	// New value to set. If `0`, the ShoppingListLineItem is removed from the ShoppingList.
	Quantity int `json:"quantity"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShoppingListChangeLineItemQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListChangeLineItemQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLineItemQuantity", Alias: (*Alias)(&obj)})
}

type ShoppingListChangeLineItemsOrderAction struct {
	// All existing ShoppingListLineItem `id`s in the desired new order.
	LineItemOrder []string `json:"lineItemOrder"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShoppingListChangeLineItemsOrderAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListChangeLineItemsOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLineItemsOrder", Alias: (*Alias)(&obj)})
}

type ShoppingListChangeNameAction struct {
	// New value to set. Must not be empty.
	Name LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShoppingListChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type ShoppingListChangeTextLineItemNameAction struct {
	// The `id` of the [TextLineItem](ctp:api:type:TextLineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	TextLineItemId *string `json:"textLineItemId,omitempty"`
	// The `key` of the [TextLineItem](ctp:api:type:TextLineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	TextLineItemKey *string `json:"textLineItemKey,omitempty"`
	// New value to set. Must not be empty.
	Name LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShoppingListChangeTextLineItemNameAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListChangeTextLineItemNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTextLineItemName", Alias: (*Alias)(&obj)})
}

type ShoppingListChangeTextLineItemQuantityAction struct {
	// The `id` of the [TextLineItem](ctp:api:type:TextLineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	TextLineItemId *string `json:"textLineItemId,omitempty"`
	// The `key` of the [TextLineItem](ctp:api:type:TextLineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	TextLineItemKey *string `json:"textLineItemKey,omitempty"`
	// New value to set. If `0`, the TextLineItem is removed from the ShoppingList.
	Quantity int `json:"quantity"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShoppingListChangeTextLineItemQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListChangeTextLineItemQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTextLineItemQuantity", Alias: (*Alias)(&obj)})
}

type ShoppingListChangeTextLineItemsOrderAction struct {
	// Must contain all existing [TextLineItem](ctp:api:type:TextLineItem) `id`s in the desired new order.
	TextLineItemOrder []string `json:"textLineItemOrder"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShoppingListChangeTextLineItemsOrderAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListChangeTextLineItemsOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTextLineItemsOrder", Alias: (*Alias)(&obj)})
}

/**
*	Produces the [Shopping List Line Item Removed](ctp:api:type:ShoppingListLineItemRemovedMessage) Message.
*
 */
type ShoppingListRemoveLineItemAction struct {
	// The `id` of the [ShoppingListLineItem](ctp:api:type:ShoppingListLineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemId *string `json:"lineItemId,omitempty"`
	// The `key` of the [ShoppingListLineItem](ctp:api:type:ShoppingListLineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemKey *string `json:"lineItemKey,omitempty"`
	// Amount to remove from the `quantity` of the ShoppingListLineItem. If not set, the ShoppingListLineItem is removed from the ShoppingList. If this value matches or exceeds the current `quantity` of the ShoppingListLineItem, the ShoppingListLineItem is removed from the ShoppingList.
	Quantity *int `json:"quantity,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShoppingListRemoveLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListRemoveLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeLineItem", Alias: (*Alias)(&obj)})
}

type ShoppingListRemoveTextLineItemAction struct {
	// The `id` of the [TextLineItem](ctp:api:type:TextLineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	TextLineItemId *string `json:"textLineItemId,omitempty"`
	// The `key` of the [TextLineItem](ctp:api:type:TextLineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	TextLineItemKey *string `json:"textLineItemKey,omitempty"`
	// Amount to remove from the `quantity` of the TextLineItem. If not set, the TextLineItem is removed from the ShoppingList. If this value matches or exceeds the current `quantity` of the TextLineItem, the TextLineItem is removed from the ShoppingList.
	Quantity *int `json:"quantity,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShoppingListRemoveTextLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListRemoveTextLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeTextLineItem", Alias: (*Alias)(&obj)})
}

/**
*	If the Shopping List is already associated with a Customer, an [InvalidOperation](ctp:api:type:InvalidOperationError) error is returned.
*
 */
type ShoppingListSetAnonymousIdAction struct {
	// Value to set. If empty, any existing value will be removed.
	AnonymousId *string `json:"anonymousId,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShoppingListSetAnonymousIdAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetAnonymousIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAnonymousId", Alias: (*Alias)(&obj)})
}

/**
*	Updates the Business Unit on the Shopping List. The Shopping List must have an existing Business Unit assigned already.
*
 */
type ShoppingListSetBusinessUnitAction struct {
	// The Business Unit to assign to the Shopping List, which must have access to the [Store](/../api/projects/stores) that is set on the Shopping List.
	BusinessUnit BusinessUnitResourceIdentifier `json:"businessUnit"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShoppingListSetBusinessUnitAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetBusinessUnitAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setBusinessUnit", Alias: (*Alias)(&obj)})
}

type ShoppingListSetCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShoppingListSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type ShoppingListSetCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the ShoppingList with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the ShoppingList.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the ShoppingList.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShoppingListSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type ShoppingListSetCustomerAction struct {
	// The [Customer](ctp:api:type:Customer) the ShoppingList should be associated to. If empty, any existing value will be removed.
	Customer *CustomerResourceIdentifier `json:"customer,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShoppingListSetCustomerAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetCustomerAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomer", Alias: (*Alias)(&obj)})
}

/**
*	Number of days after the last modification before a Shopping List is deleted.
*
 */
type ShoppingListSetDeleteDaysAfterLastModificationAction struct {
	// Value to set. If not provided, the default value for this field configured in [Project settings](ctp:api:type:ShoppingListsConfiguration) is assigned.
	DeleteDaysAfterLastModification *int `json:"deleteDaysAfterLastModification,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShoppingListSetDeleteDaysAfterLastModificationAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetDeleteDaysAfterLastModificationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeleteDaysAfterLastModification", Alias: (*Alias)(&obj)})
}

type ShoppingListSetDescriptionAction struct {
	// Value to set. If empty, any existing value will be removed.
	Description *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShoppingListSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type ShoppingListSetKeyAction struct {
	// Value to set. If empty, any existing value will be removed.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShoppingListSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type ShoppingListSetLineItemCustomFieldAction struct {
	// The `id` of the [ShoppingListLineItem](ctp:api:type:ShoppingListLineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemId *string `json:"lineItemId,omitempty"`
	// The `key` of the [ShoppingListLineItem](ctp:api:type:ShoppingListLineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemKey *string `json:"lineItemKey,omitempty"`
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShoppingListSetLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomField", Alias: (*Alias)(&obj)})
}

type ShoppingListSetLineItemCustomTypeAction struct {
	// The `id` of the [ShoppingListLineItem](ctp:api:type:ShoppingListLineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemId *string `json:"lineItemId,omitempty"`
	// The `key` of the [ShoppingListLineItem](ctp:api:type:ShoppingListLineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemKey *string `json:"lineItemKey,omitempty"`
	// Defines the [Type](ctp:api:type:Type) that extends the ShoppingListLineItem with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the ShoppingListLineItem.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the ShoppingListLineItem.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShoppingListSetLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomType", Alias: (*Alias)(&obj)})
}

type ShoppingListSetSlugAction struct {
	// Value to set. If empty, any existing value will be removed. Each slug is unique across a Project, but a ShoppingList can have the same slug for different languages. Must match the pattern `^[A-Za-z0-9_-]{2,256}+$`
	Slug *LocalizedString `json:"slug,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShoppingListSetSlugAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetSlugAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSlug", Alias: (*Alias)(&obj)})
}

type ShoppingListSetStoreAction struct {
	// The [Store](ctp:api:type:Store) the ShoppingList should be assigned to. If empty, any existing value will be removed.
	Store *StoreResourceIdentifier `json:"store,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShoppingListSetStoreAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetStoreAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setStore", Alias: (*Alias)(&obj)})
}

type ShoppingListSetTextLineItemCustomFieldAction struct {
	// The `id` of the [TextLineItem](ctp:api:type:TextLineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	TextLineItemId *string `json:"textLineItemId,omitempty"`
	// The `key` of the [TextLineItem](ctp:api:type:TextLineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	TextLineItemKey *string `json:"textLineItemKey,omitempty"`
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShoppingListSetTextLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetTextLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTextLineItemCustomField", Alias: (*Alias)(&obj)})
}

type ShoppingListSetTextLineItemCustomTypeAction struct {
	// The `id` of the [TextLineItem](ctp:api:type:TextLineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	TextLineItemId *string `json:"textLineItemId,omitempty"`
	// The `key` of the [TextLineItem](ctp:api:type:TextLineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	TextLineItemKey *string `json:"textLineItemKey,omitempty"`
	// Defines the [Type](ctp:api:type:Type) that extends the TextLineItem with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the TextLineItem.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the TextLineItem.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShoppingListSetTextLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetTextLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTextLineItemCustomType", Alias: (*Alias)(&obj)})
}

type ShoppingListSetTextLineItemDescriptionAction struct {
	// The `id` of the [TextLineItem](ctp:api:type:TextLineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	TextLineItemId *string `json:"textLineItemId,omitempty"`
	// The `key` of the [TextLineItem](ctp:api:type:TextLineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	TextLineItemKey *string `json:"textLineItemKey,omitempty"`
	// Value to set. If empty, any existing value will be removed.
	Description *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShoppingListSetTextLineItemDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetTextLineItemDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTextLineItemDescription", Alias: (*Alias)(&obj)})
}
