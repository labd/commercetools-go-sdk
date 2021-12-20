// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type ShoppingList struct {
	// The unique ID of the shopping list.
	ID string `json:"id"`
	// The current version of the shopping list.
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy         `json:"createdBy,omitempty"`
	Custom    *CustomFields      `json:"custom,omitempty"`
	Customer  *CustomerReference `json:"customer,omitempty"`
	// The shopping list will be deleted automatically if it hasn't been modified for the specified amount of days.
	DeleteDaysAfterLastModification *int             `json:"deleteDaysAfterLastModification,omitempty"`
	Description                     *LocalizedString `json:"description,omitempty"`
	// User-specific unique identifier for the shopping list.
	Key       *string                `json:"key,omitempty"`
	LineItems []ShoppingListLineItem `json:"lineItems,omitempty"`
	Name      LocalizedString        `json:"name"`
	// Human-readable identifiers usually used as deep-link URL to the related shopping list.
	// Each slug is unique across a project, but a shopping list can have the same slug for different languages.
	// The slug must match the pattern [a-zA-Z0-9_-]{2,256}.
	Slug          *LocalizedString `json:"slug,omitempty"`
	TextLineItems []TextLineItem   `json:"textLineItems,omitempty"`
	// Identifies shopping lists belonging to an anonymous session (the customer has not signed up/in yet).
	AnonymousId *string            `json:"anonymousId,omitempty"`
	Store       *StoreKeyReference `json:"store,omitempty"`
}

type ShoppingListDraft struct {
	// The custom fields.
	Custom   *CustomFieldsDraft          `json:"custom,omitempty"`
	Customer *CustomerResourceIdentifier `json:"customer,omitempty"`
	// The shopping list will be deleted automatically if it hasn't been modified for the specified amount of days.
	DeleteDaysAfterLastModification *int             `json:"deleteDaysAfterLastModification,omitempty"`
	Description                     *LocalizedString `json:"description,omitempty"`
	// User-specific unique identifier for the shopping list.
	Key       *string                     `json:"key,omitempty"`
	LineItems []ShoppingListLineItemDraft `json:"lineItems,omitempty"`
	Name      LocalizedString             `json:"name"`
	// Human-readable identifiers usually used as deep-link URL to the related shopping list.
	// Each slug is unique across a project, but a shopping list can have the same slug for different languages.
	// The slug must match the pattern [a-zA-Z0-9_-]{2,256}.
	Slug          *LocalizedString    `json:"slug,omitempty"`
	TextLineItems []TextLineItemDraft `json:"textLineItems,omitempty"`
	// Identifies shopping lists belonging to an anonymous session (the customer has not signed up/in yet).
	AnonymousId *string                  `json:"anonymousId,omitempty"`
	Store       *StoreResourceIdentifier `json:"store,omitempty"`
}

type ShoppingListLineItem struct {
	AddedAt       time.Time            `json:"addedAt"`
	Custom        *CustomFields        `json:"custom,omitempty"`
	DeactivatedAt *time.Time           `json:"deactivatedAt,omitempty"`
	ID            string               `json:"id"`
	Name          LocalizedString      `json:"name"`
	ProductId     string               `json:"productId"`
	ProductSlug   *LocalizedString     `json:"productSlug,omitempty"`
	ProductType   ProductTypeReference `json:"productType"`
	Quantity      int                  `json:"quantity"`
	Variant       *ProductVariant      `json:"variant,omitempty"`
	VariantId     *int                 `json:"variantId,omitempty"`
}

type ShoppingListLineItemDraft struct {
	AddedAt   *time.Time         `json:"addedAt,omitempty"`
	Custom    *CustomFieldsDraft `json:"custom,omitempty"`
	Sku       *string            `json:"sku,omitempty"`
	ProductId *string            `json:"productId,omitempty"`
	Quantity  *int               `json:"quantity,omitempty"`
	VariantId *int               `json:"variantId,omitempty"`
}

type ShoppingListPagedQueryResponse struct {
	Limit   int            `json:"limit"`
	Count   int            `json:"count"`
	Total   *int           `json:"total,omitempty"`
	Offset  int            `json:"offset"`
	Results []ShoppingList `json:"results"`
}

type ShoppingListReference struct {
	// Unique ID of the referenced resource.
	ID  string        `json:"id"`
	Obj *ShoppingList `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListReference) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "shopping-list", Alias: (*Alias)(&obj)})
}

type ShoppingListResourceIdentifier struct {
	// Unique ID of the referenced resource. Either `id` or `key` is required.
	ID *string `json:"id,omitempty"`
	// Unique key of the referenced resource. Either `id` or `key` is required.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "shopping-list", Alias: (*Alias)(&obj)})
}

type ShoppingListUpdate struct {
	Version int                        `json:"version"`
	Actions []ShoppingListUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ShoppingListUpdate) UnmarshalJSON(data []byte) error {
	type Alias ShoppingListUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	return nil
}

type ShoppingListUpdateAction interface{}

func mapDiscriminatorShoppingListUpdateAction(input interface{}) (ShoppingListUpdateAction, error) {

	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["action"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'action'")
		}
	} else {
		return nil, errors.New("Invalid data")
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

type TextLineItem struct {
	// When the text line item was added to the shopping list.
	AddedAt     time.Time        `json:"addedAt"`
	Custom      *CustomFields    `json:"custom,omitempty"`
	Description *LocalizedString `json:"description,omitempty"`
	// The unique ID of this TextLineItem.
	ID       string          `json:"id"`
	Name     LocalizedString `json:"name"`
	Quantity int             `json:"quantity"`
}

type TextLineItemDraft struct {
	// Defaults to the current date and time.
	AddedAt *time.Time `json:"addedAt,omitempty"`
	// The custom fields.
	Custom      *CustomFieldsDraft `json:"custom,omitempty"`
	Description *LocalizedString   `json:"description,omitempty"`
	Name        LocalizedString    `json:"name"`
	// Defaults to `1`.
	Quantity *int `json:"quantity,omitempty"`
}

type ShoppingListAddLineItemAction struct {
	Sku       *string            `json:"sku,omitempty"`
	ProductId *string            `json:"productId,omitempty"`
	VariantId *int               `json:"variantId,omitempty"`
	Quantity  *int               `json:"quantity,omitempty"`
	AddedAt   *time.Time         `json:"addedAt,omitempty"`
	Custom    *CustomFieldsDraft `json:"custom,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListAddLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListAddLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addLineItem", Alias: (*Alias)(&obj)})
}

type ShoppingListAddTextLineItemAction struct {
	Name        LocalizedString    `json:"name"`
	Description *LocalizedString   `json:"description,omitempty"`
	Quantity    *int               `json:"quantity,omitempty"`
	AddedAt     *time.Time         `json:"addedAt,omitempty"`
	Custom      *CustomFieldsDraft `json:"custom,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListAddTextLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListAddTextLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addTextLineItem", Alias: (*Alias)(&obj)})
}

type ShoppingListChangeLineItemQuantityAction struct {
	LineItemId string `json:"lineItemId"`
	Quantity   int    `json:"quantity"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListChangeLineItemQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListChangeLineItemQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLineItemQuantity", Alias: (*Alias)(&obj)})
}

type ShoppingListChangeLineItemsOrderAction struct {
	LineItemOrder []string `json:"lineItemOrder"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListChangeLineItemsOrderAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListChangeLineItemsOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLineItemsOrder", Alias: (*Alias)(&obj)})
}

type ShoppingListChangeNameAction struct {
	Name LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type ShoppingListChangeTextLineItemNameAction struct {
	TextLineItemId string          `json:"textLineItemId"`
	Name           LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListChangeTextLineItemNameAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListChangeTextLineItemNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTextLineItemName", Alias: (*Alias)(&obj)})
}

type ShoppingListChangeTextLineItemQuantityAction struct {
	TextLineItemId string `json:"textLineItemId"`
	Quantity       int    `json:"quantity"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListChangeTextLineItemQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListChangeTextLineItemQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTextLineItemQuantity", Alias: (*Alias)(&obj)})
}

type ShoppingListChangeTextLineItemsOrderAction struct {
	TextLineItemOrder []string `json:"textLineItemOrder"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListChangeTextLineItemsOrderAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListChangeTextLineItemsOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTextLineItemsOrder", Alias: (*Alias)(&obj)})
}

type ShoppingListRemoveLineItemAction struct {
	LineItemId string `json:"lineItemId"`
	Quantity   *int   `json:"quantity,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListRemoveLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListRemoveLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeLineItem", Alias: (*Alias)(&obj)})
}

type ShoppingListRemoveTextLineItemAction struct {
	TextLineItemId string `json:"textLineItemId"`
	Quantity       *int   `json:"quantity,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListRemoveTextLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListRemoveTextLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeTextLineItem", Alias: (*Alias)(&obj)})
}

type ShoppingListSetAnonymousIdAction struct {
	// Anonymous ID of the anonymous customer that this shopping list belongs to.
	// If this field is not set any existing `anonymousId` is removed.
	AnonymousId *string `json:"anonymousId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetAnonymousIdAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetAnonymousIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAnonymousId", Alias: (*Alias)(&obj)})
}

type ShoppingListSetCustomFieldAction struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type ShoppingListSetCustomTypeAction struct {
	// If set, the custom type is set to this new value.
	// If absent, the custom type and any existing custom fields are removed.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// If set, the custom fields are set to this new value.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type ShoppingListSetCustomerAction struct {
	Customer *CustomerResourceIdentifier `json:"customer,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetCustomerAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetCustomerAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomer", Alias: (*Alias)(&obj)})
}

type ShoppingListSetDeleteDaysAfterLastModificationAction struct {
	DeleteDaysAfterLastModification *int `json:"deleteDaysAfterLastModification,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetDeleteDaysAfterLastModificationAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetDeleteDaysAfterLastModificationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeleteDaysAfterLastModification", Alias: (*Alias)(&obj)})
}

type ShoppingListSetDescriptionAction struct {
	Description *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type ShoppingListSetKeyAction struct {
	// User-specific unique identifier for the shopping list.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type ShoppingListSetLineItemCustomFieldAction struct {
	LineItemId string      `json:"lineItemId"`
	Name       string      `json:"name"`
	Value      interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomField", Alias: (*Alias)(&obj)})
}

type ShoppingListSetLineItemCustomTypeAction struct {
	LineItemId string                  `json:"lineItemId"`
	Type       *TypeResourceIdentifier `json:"type,omitempty"`
	Fields     *FieldContainer         `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomType", Alias: (*Alias)(&obj)})
}

type ShoppingListSetSlugAction struct {
	Slug *LocalizedString `json:"slug,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetSlugAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetSlugAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSlug", Alias: (*Alias)(&obj)})
}

type ShoppingListSetStoreAction struct {
	Store *StoreResourceIdentifier `json:"store,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetStoreAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetStoreAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setStore", Alias: (*Alias)(&obj)})
}

type ShoppingListSetTextLineItemCustomFieldAction struct {
	TextLineItemId string      `json:"textLineItemId"`
	Name           string      `json:"name"`
	Value          interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetTextLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetTextLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTextLineItemCustomField", Alias: (*Alias)(&obj)})
}

type ShoppingListSetTextLineItemCustomTypeAction struct {
	TextLineItemId string                  `json:"textLineItemId"`
	Type           *TypeResourceIdentifier `json:"type,omitempty"`
	Fields         *FieldContainer         `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetTextLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetTextLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTextLineItemCustomType", Alias: (*Alias)(&obj)})
}

type ShoppingListSetTextLineItemDescriptionAction struct {
	TextLineItemId string           `json:"textLineItemId"`
	Description    *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetTextLineItemDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetTextLineItemDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTextLineItemDescription", Alias: (*Alias)(&obj)})
}
