// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type ShoppingList struct {
	// The unique ID of the shopping list.
	Id string `json:"id"`
	// The current version of the shopping list.
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources updated after 1/02/2019 except for events not tracked.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitEmpty"`
	// Present on resources created after 1/02/2019 except for events not tracked.
	CreatedBy *CreatedBy        `json:"createdBy,omitEmpty"`
	Custom    *CustomFields     `json:"custom,omitEmpty"`
	Customer  CustomerReference `json:"customer,omitEmpty"`
	// The shopping list will be deleted automatically if it hasn't been modified for the specified amount of days.
	DeleteDaysAfterLastModification int              `json:"deleteDaysAfterLastModification,omitEmpty"`
	Description                     *LocalizedString `json:"description,omitEmpty"`
	// User-specific unique identifier for the shopping list.
	Key       string                 `json:"key,omitEmpty"`
	LineItems []ShoppingListLineItem `json:"lineItems,omitEmpty"`
	Name      LocalizedString        `json:"name"`
	// Human-readable identifiers usually used as deep-link URL to the related shopping list.
	// Each slug is unique across a project, but a shopping list can have the same slug for different languages.
	// The slug must match the pattern [a-zA-Z0-9_-]{2,256}.
	Slug          *LocalizedString `json:"slug,omitEmpty"`
	TextLineItems []TextLineItem   `json:"textLineItems,omitEmpty"`
	// Identifies shopping lists belonging to an anonymous session (the customer has not signed up/in yet).
	AnonymousId string            `json:"anonymousId,omitEmpty"`
	Store       StoreKeyReference `json:"store,omitEmpty"`
}

type ShoppingListDraft struct {
	// The custom fields.
	Custom   *CustomFieldsDraft         `json:"custom,omitEmpty"`
	Customer CustomerResourceIdentifier `json:"customer,omitEmpty"`
	// The shopping list will be deleted automatically if it hasn't been modified for the specified amount of days.
	DeleteDaysAfterLastModification int              `json:"deleteDaysAfterLastModification,omitEmpty"`
	Description                     *LocalizedString `json:"description,omitEmpty"`
	// User-specific unique identifier for the shopping list.
	Key       string                      `json:"key,omitEmpty"`
	LineItems []ShoppingListLineItemDraft `json:"lineItems,omitEmpty"`
	Name      LocalizedString             `json:"name"`
	// Human-readable identifiers usually used as deep-link URL to the related shopping list.
	// Each slug is unique across a project, but a shopping list can have the same slug for different languages.
	// The slug must match the pattern [a-zA-Z0-9_-]{2,256}.
	Slug          *LocalizedString    `json:"slug,omitEmpty"`
	TextLineItems []TextLineItemDraft `json:"textLineItems,omitEmpty"`
	// Identifies shopping lists belonging to an anonymous session (the customer has not signed up/in yet).
	AnonymousId string                  `json:"anonymousId,omitEmpty"`
	Store       StoreResourceIdentifier `json:"store,omitEmpty"`
}

type ShoppingListLineItem struct {
	AddedAt       time.Time            `json:"addedAt"`
	Custom        *CustomFields        `json:"custom,omitEmpty"`
	DeactivatedAt time.Time            `json:"deactivatedAt,omitEmpty"`
	Id            string               `json:"id"`
	Name          LocalizedString      `json:"name"`
	ProductId     string               `json:"productId"`
	ProductSlug   *LocalizedString     `json:"productSlug,omitEmpty"`
	ProductType   ProductTypeReference `json:"productType"`
	Quantity      int                  `json:"quantity"`
	Variant       *ProductVariant      `json:"variant,omitEmpty"`
	VariantId     int                  `json:"variantId,omitEmpty"`
}

type ShoppingListLineItemDraft struct {
	AddedAt   time.Time          `json:"addedAt,omitEmpty"`
	Custom    *CustomFieldsDraft `json:"custom,omitEmpty"`
	Sku       string             `json:"sku,omitEmpty"`
	ProductId string             `json:"productId,omitEmpty"`
	Quantity  int                `json:"quantity,omitEmpty"`
	VariantId int                `json:"variantId,omitEmpty"`
}

type ShoppingListPagedQueryResponse struct {
	Limit   int            `json:"limit"`
	Count   int            `json:"count"`
	Total   int            `json:"total,omitEmpty"`
	Offset  int            `json:"offset"`
	Results []ShoppingList `json:"results"`
}

type ShoppingListReference struct {
	Id  string        `json:"id"`
	Obj *ShoppingList `json:"obj,omitEmpty"`
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
	Id  string `json:"id,omitEmpty"`
	Key string `json:"key,omitEmpty"`
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
		new := ShoppingListAddLineItemAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addTextLineItem":
		new := ShoppingListAddTextLineItemAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeLineItemQuantity":
		new := ShoppingListChangeLineItemQuantityAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeLineItemsOrder":
		new := ShoppingListChangeLineItemsOrderAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeName":
		new := ShoppingListChangeNameAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeTextLineItemName":
		new := ShoppingListChangeTextLineItemNameAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeTextLineItemQuantity":
		new := ShoppingListChangeTextLineItemQuantityAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeTextLineItemsOrder":
		new := ShoppingListChangeTextLineItemsOrderAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeLineItem":
		new := ShoppingListRemoveLineItemAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeTextLineItem":
		new := ShoppingListRemoveTextLineItemAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setAnonymousId":
		new := ShoppingListSetAnonymousIdAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomField":
		new := ShoppingListSetCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomType":
		new := ShoppingListSetCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomer":
		new := ShoppingListSetCustomerAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDeleteDaysAfterLastModification":
		new := ShoppingListSetDeleteDaysAfterLastModificationAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDescription":
		new := ShoppingListSetDescriptionAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setKey":
		new := ShoppingListSetKeyAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemCustomField":
		new := ShoppingListSetLineItemCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemCustomType":
		new := ShoppingListSetLineItemCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setSlug":
		new := ShoppingListSetSlugAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setStore":
		new := ShoppingListSetStoreAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setTextLineItemCustomField":
		new := ShoppingListSetTextLineItemCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setTextLineItemCustomType":
		new := ShoppingListSetTextLineItemCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setTextLineItemDescription":
		new := ShoppingListSetTextLineItemDescriptionAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type TextLineItem struct {
	// When the text line item was added to the shopping list.
	AddedAt     time.Time        `json:"addedAt"`
	Custom      *CustomFields    `json:"custom,omitEmpty"`
	Description *LocalizedString `json:"description,omitEmpty"`
	// The unique ID of this TextLineItem.
	Id       string          `json:"id"`
	Name     LocalizedString `json:"name"`
	Quantity int             `json:"quantity"`
}

type TextLineItemDraft struct {
	// Defaults to the current date and time.
	AddedAt time.Time `json:"addedAt,omitEmpty"`
	// The custom fields.
	Custom      *CustomFieldsDraft `json:"custom,omitEmpty"`
	Description *LocalizedString   `json:"description,omitEmpty"`
	Name        LocalizedString    `json:"name"`
	// Defaults to `1`.
	Quantity int `json:"quantity,omitEmpty"`
}

type ShoppingListAddLineItemAction struct {
	Sku       string             `json:"sku,omitEmpty"`
	ProductId string             `json:"productId,omitEmpty"`
	VariantId int                `json:"variantId,omitEmpty"`
	Quantity  int                `json:"quantity,omitEmpty"`
	AddedAt   time.Time          `json:"addedAt,omitEmpty"`
	Custom    *CustomFieldsDraft `json:"custom,omitEmpty"`
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
	Description *LocalizedString   `json:"description,omitEmpty"`
	Quantity    int                `json:"quantity,omitEmpty"`
	AddedAt     time.Time          `json:"addedAt,omitEmpty"`
	Custom      *CustomFieldsDraft `json:"custom,omitEmpty"`
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
	Quantity   int    `json:"quantity,omitEmpty"`
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
	Quantity       int    `json:"quantity,omitEmpty"`
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
	AnonymousId string `json:"anonymousId,omitEmpty"`
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
	Value interface{} `json:"value,omitEmpty"`
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
	Type TypeResourceIdentifier `json:"type,omitEmpty"`
	// If set, the custom fields are set to this new value.
	Fields *FieldContainer `json:"fields,omitEmpty"`
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
	Customer CustomerResourceIdentifier `json:"customer,omitEmpty"`
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
	DeleteDaysAfterLastModification int `json:"deleteDaysAfterLastModification,omitEmpty"`
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
	Description *LocalizedString `json:"description,omitEmpty"`
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
	Key string `json:"key,omitEmpty"`
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
	Value      interface{} `json:"value,omitEmpty"`
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
	LineItemId string                 `json:"lineItemId"`
	Type       TypeResourceIdentifier `json:"type,omitEmpty"`
	Fields     *FieldContainer        `json:"fields,omitEmpty"`
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
	Slug *LocalizedString `json:"slug,omitEmpty"`
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
	Store StoreResourceIdentifier `json:"store,omitEmpty"`
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
	Value          interface{} `json:"value,omitEmpty"`
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
	TextLineItemId string                 `json:"textLineItemId"`
	Type           TypeResourceIdentifier `json:"type,omitEmpty"`
	Fields         *FieldContainer        `json:"fields,omitEmpty"`
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
	Description    *LocalizedString `json:"description,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetTextLineItemDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetTextLineItemDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTextLineItemDescription", Alias: (*Alias)(&obj)})
}
