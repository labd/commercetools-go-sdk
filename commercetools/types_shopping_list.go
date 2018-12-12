// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// ShoppingListUpdateAction uses action as discriminator attribute
type ShoppingListUpdateAction interface{}

func mapDiscriminatorShoppingListUpdateAction(input interface{}) ShoppingListUpdateAction {
	discriminator := input.(map[string]interface{})["action"]
	switch discriminator {
	case "addLineItem":
		new := ShoppingListAddLineItemAction{}
		mapstructure.Decode(input, &new)
		return new
	case "addTextLineItem":
		new := ShoppingListAddTextLineItemAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeLineItemQuantity":
		new := ShoppingListChangeLineItemQuantityAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeLineItemsOrder":
		new := ShoppingListChangeLineItemsOrderAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeName":
		new := ShoppingListChangeNameAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeTextLineItemName":
		new := ShoppingListChangeTextLineItemNameAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeTextLineItemQuantity":
		new := ShoppingListChangeTextLineItemQuantityAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeTextLineItemsOrder":
		new := ShoppingListChangeTextLineItemsOrderAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeLineItem":
		new := ShoppingListRemoveLineItemAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeTextLineItem":
		new := ShoppingListRemoveTextLineItemAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setAnonymousId":
		new := ShoppingListSetAnonymousIDAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomField":
		new := ShoppingListSetCustomFieldAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomType":
		new := ShoppingListSetCustomTypeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomer":
		new := ShoppingListSetCustomerAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setDeleteDaysAfterLastModification":
		new := ShoppingListSetDeleteDaysAfterLastModificationAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setDescription":
		new := ShoppingListSetDescriptionAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setKey":
		new := ShoppingListSetKeyAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setLineItemCustomField":
		new := ShoppingListSetLineItemCustomFieldAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setLineItemCustomType":
		new := ShoppingListSetLineItemCustomTypeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setSlug":
		new := ShoppingListSetSlugAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setTextLineItemCustomField":
		new := ShoppingListSetTextLineItemCustomFieldAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setTextLineItemCustomType":
		new := ShoppingListSetTextLineItemCustomTypeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setTextLineItemDescription":
		new := ShoppingListSetTextLineItemDescriptionAction{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}

// ShoppingList is of type Resource
type ShoppingList struct {
	Version                         int                    `json:"version"`
	LastModifiedAt                  time.Time              `json:"lastModifiedAt"`
	ID                              string                 `json:"id"`
	CreatedAt                       time.Time              `json:"createdAt"`
	TextLineItems                   []TextLineItem         `json:"textLineItems,omitempty"`
	Slug                            *LocalizedString       `json:"slug,omitempty"`
	Name                            *LocalizedString       `json:"name"`
	LineItems                       []ShoppingListLineItem `json:"lineItems,omitempty"`
	Key                             string                 `json:"key,omitempty"`
	Description                     *LocalizedString       `json:"description,omitempty"`
	DeleteDaysAfterLastModification int                    `json:"deleteDaysAfterLastModification,omitempty"`
	Customer                        *CustomerReference     `json:"customer,omitempty"`
	Custom                          *CustomFields          `json:"custom,omitempty"`
	AnonymousID                     string                 `json:"anonymousId,omitempty"`
}

// ShoppingListAddLineItemAction implements the interface ShoppingListUpdateAction
type ShoppingListAddLineItemAction struct {
	VariantID int                `json:"variantId,omitempty"`
	SKU       string             `json:"sku,omitempty"`
	Quantity  int                `json:"quantity,omitempty"`
	ProductID string             `json:"productId,omitempty"`
	Custom    *CustomFieldsDraft `json:"custom,omitempty"`
	AddedAt   time.Time          `json:"addedAt,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListAddLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListAddLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addLineItem", Alias: (*Alias)(&obj)})
}

// ShoppingListAddTextLineItemAction implements the interface ShoppingListUpdateAction
type ShoppingListAddTextLineItemAction struct {
	Quantity    int                `json:"quantity,omitempty"`
	Name        *LocalizedString   `json:"name"`
	Description *LocalizedString   `json:"description,omitempty"`
	Custom      *CustomFieldsDraft `json:"custom,omitempty"`
	AddedAt     time.Time          `json:"addedAt,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListAddTextLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListAddTextLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addTextLineItem", Alias: (*Alias)(&obj)})
}

// ShoppingListChangeLineItemQuantityAction implements the interface ShoppingListUpdateAction
type ShoppingListChangeLineItemQuantityAction struct {
	Quantity   int    `json:"quantity"`
	LineItemID string `json:"lineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListChangeLineItemQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListChangeLineItemQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLineItemQuantity", Alias: (*Alias)(&obj)})
}

// ShoppingListChangeLineItemsOrderAction implements the interface ShoppingListUpdateAction
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

// ShoppingListChangeNameAction implements the interface ShoppingListUpdateAction
type ShoppingListChangeNameAction struct {
	Name *LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

// ShoppingListChangeTextLineItemNameAction implements the interface ShoppingListUpdateAction
type ShoppingListChangeTextLineItemNameAction struct {
	TextLineItemID string           `json:"textLineItemId"`
	Name           *LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListChangeTextLineItemNameAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListChangeTextLineItemNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTextLineItemName", Alias: (*Alias)(&obj)})
}

// ShoppingListChangeTextLineItemQuantityAction implements the interface ShoppingListUpdateAction
type ShoppingListChangeTextLineItemQuantityAction struct {
	TextLineItemID string `json:"textLineItemId"`
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

// ShoppingListChangeTextLineItemsOrderAction implements the interface ShoppingListUpdateAction
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

// ShoppingListDraft is a standalone struct
type ShoppingListDraft struct {
	TextLineItems                   []TextLineItemDraft         `json:"textLineItems,omitempty"`
	Slug                            *LocalizedString            `json:"slug,omitempty"`
	Name                            *LocalizedString            `json:"name"`
	LineItems                       []ShoppingListLineItemDraft `json:"lineItems,omitempty"`
	Key                             string                      `json:"key,omitempty"`
	Description                     *LocalizedString            `json:"description,omitempty"`
	DeleteDaysAfterLastModification int                         `json:"deleteDaysAfterLastModification,omitempty"`
	Customer                        *CustomerReference          `json:"customer,omitempty"`
	Custom                          *CustomFieldsDraft          `json:"custom,omitempty"`
	AnonymousID                     string                      `json:"anonymousId,omitempty"`
}

// ShoppingListLineItem is a standalone struct
type ShoppingListLineItem struct {
	VariantID     int                   `json:"variantId,omitempty"`
	Variant       *ProductVariant       `json:"variant,omitempty"`
	Quantity      float64               `json:"quantity"`
	ProductType   *ProductTypeReference `json:"productType"`
	ProductSlug   *LocalizedString      `json:"productSlug,omitempty"`
	ProductID     string                `json:"productId"`
	Name          *LocalizedString      `json:"name"`
	ID            string                `json:"id"`
	DeactivatedAt time.Time             `json:"deactivatedAt,omitempty"`
	Custom        *CustomFields         `json:"custom,omitempty"`
	AddedAt       time.Time             `json:"addedAt"`
}

// ShoppingListLineItemDraft is a standalone struct
type ShoppingListLineItemDraft struct {
	VariantID int                `json:"variantId,omitempty"`
	SKU       string             `json:"sku,omitempty"`
	Quantity  float64            `json:"quantity,omitempty"`
	ProductID string             `json:"productId,omitempty"`
	Custom    *CustomFieldsDraft `json:"custom,omitempty"`
	AddedAt   time.Time          `json:"addedAt,omitempty"`
}

// ShoppingListPagedQueryResponse is of type PagedQueryResponse
type ShoppingListPagedQueryResponse struct {
	Total   int            `json:"total,omitempty"`
	Offset  int            `json:"offset"`
	Count   int            `json:"count"`
	Results []ShoppingList `json:"results"`
}

// ShoppingListReference implements the interface Reference
type ShoppingListReference struct {
	Key string        `json:"key,omitempty"`
	ID  string        `json:"id,omitempty"`
	Obj *ShoppingList `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListReference) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "shopping-list", Alias: (*Alias)(&obj)})
}

// ShoppingListRemoveLineItemAction implements the interface ShoppingListUpdateAction
type ShoppingListRemoveLineItemAction struct {
	Quantity   int    `json:"quantity,omitempty"`
	LineItemID string `json:"lineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListRemoveLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListRemoveLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeLineItem", Alias: (*Alias)(&obj)})
}

// ShoppingListRemoveTextLineItemAction implements the interface ShoppingListUpdateAction
type ShoppingListRemoveTextLineItemAction struct {
	TextLineItemID string `json:"textLineItemId"`
	Quantity       int    `json:"quantity,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListRemoveTextLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListRemoveTextLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeTextLineItem", Alias: (*Alias)(&obj)})
}

// ShoppingListSetAnonymousIDAction implements the interface ShoppingListUpdateAction
type ShoppingListSetAnonymousIDAction struct {
	AnonymousID string `json:"anonymousId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetAnonymousIDAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetAnonymousIDAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAnonymousId", Alias: (*Alias)(&obj)})
}

// ShoppingListSetCustomFieldAction implements the interface ShoppingListUpdateAction
type ShoppingListSetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

// ShoppingListSetCustomTypeAction implements the interface ShoppingListUpdateAction
type ShoppingListSetCustomTypeAction struct {
	Type   *TypeReference  `json:"type,omitempty"`
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

// ShoppingListSetCustomerAction implements the interface ShoppingListUpdateAction
type ShoppingListSetCustomerAction struct {
	Customer *CustomerReference `json:"customer,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetCustomerAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetCustomerAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomer", Alias: (*Alias)(&obj)})
}

// ShoppingListSetDeleteDaysAfterLastModificationAction implements the interface ShoppingListUpdateAction
type ShoppingListSetDeleteDaysAfterLastModificationAction struct {
	DeleteDaysAfterLastModification int `json:"deleteDaysAfterLastModification,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetDeleteDaysAfterLastModificationAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetDeleteDaysAfterLastModificationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeleteDaysAfterLastModification", Alias: (*Alias)(&obj)})
}

// ShoppingListSetDescriptionAction implements the interface ShoppingListUpdateAction
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

// ShoppingListSetKeyAction implements the interface ShoppingListUpdateAction
type ShoppingListSetKeyAction struct {
	Key string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

// ShoppingListSetLineItemCustomFieldAction implements the interface ShoppingListUpdateAction
type ShoppingListSetLineItemCustomFieldAction struct {
	Value      interface{} `json:"value,omitempty"`
	Name       string      `json:"name"`
	LineItemID string      `json:"lineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomField", Alias: (*Alias)(&obj)})
}

// ShoppingListSetLineItemCustomTypeAction implements the interface ShoppingListUpdateAction
type ShoppingListSetLineItemCustomTypeAction struct {
	Type       *TypeReference  `json:"type,omitempty"`
	LineItemID string          `json:"lineItemId"`
	Fields     *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomType", Alias: (*Alias)(&obj)})
}

// ShoppingListSetSlugAction implements the interface ShoppingListUpdateAction
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

// ShoppingListSetTextLineItemCustomFieldAction implements the interface ShoppingListUpdateAction
type ShoppingListSetTextLineItemCustomFieldAction struct {
	Value          interface{} `json:"value,omitempty"`
	TextLineItemID string      `json:"textLineItemId"`
	Name           string      `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetTextLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetTextLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTextLineItemCustomField", Alias: (*Alias)(&obj)})
}

// ShoppingListSetTextLineItemCustomTypeAction implements the interface ShoppingListUpdateAction
type ShoppingListSetTextLineItemCustomTypeAction struct {
	Type           *TypeReference  `json:"type,omitempty"`
	TextLineItemID string          `json:"textLineItemId"`
	Fields         *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetTextLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetTextLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTextLineItemCustomType", Alias: (*Alias)(&obj)})
}

// ShoppingListSetTextLineItemDescriptionAction implements the interface ShoppingListUpdateAction
type ShoppingListSetTextLineItemDescriptionAction struct {
	TextLineItemID string           `json:"textLineItemId"`
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

// ShoppingListUpdate is of type Update
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
	for i := range obj.Actions {
		obj.Actions[i] = mapDiscriminatorShoppingListUpdateAction(obj.Actions[i])
	}

	return nil
}

// TextLineItem is a standalone struct
type TextLineItem struct {
	Quantity    float64          `json:"quantity"`
	Name        *LocalizedString `json:"name"`
	ID          string           `json:"id"`
	Description *LocalizedString `json:"description,omitempty"`
	Custom      *CustomFields    `json:"custom,omitempty"`
	AddedAt     time.Time        `json:"addedAt"`
}

// TextLineItemDraft is a standalone struct
type TextLineItemDraft struct {
	Quantity    float64            `json:"quantity,omitempty"`
	Name        *LocalizedString   `json:"name"`
	Description *LocalizedString   `json:"description,omitempty"`
	Custom      *CustomFieldsDraft `json:"custom,omitempty"`
	AddedAt     time.Time          `json:"addedAt,omitempty"`
}
