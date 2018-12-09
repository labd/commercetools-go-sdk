// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

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

type ShoppingListAddLineItemAction struct {
	VariantID int                `json:"variantId,omitempty"`
	Sku       string             `json:"sku,omitempty"`
	Quantity  int                `json:"quantity,omitempty"`
	ProductID string             `json:"productId,omitempty"`
	Custom    *CustomFieldsDraft `json:"custom,omitempty"`
	AddedAt   time.Time          `json:"addedAt,omitempty"`
}

func (obj ShoppingListAddLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListAddLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addLineItem", Alias: (*Alias)(&obj)})
}

type ShoppingListAddTextLineItemAction struct {
	Quantity    int                `json:"quantity,omitempty"`
	Name        *LocalizedString   `json:"name"`
	Description *LocalizedString   `json:"description,omitempty"`
	Custom      *CustomFieldsDraft `json:"custom,omitempty"`
	AddedAt     time.Time          `json:"addedAt,omitempty"`
}

func (obj ShoppingListAddTextLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListAddTextLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addTextLineItem", Alias: (*Alias)(&obj)})
}

type ShoppingListChangeLineItemQuantityAction struct {
	Quantity   int    `json:"quantity"`
	LineItemID string `json:"lineItemId"`
}

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

func (obj ShoppingListChangeLineItemsOrderAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListChangeLineItemsOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLineItemsOrder", Alias: (*Alias)(&obj)})
}

type ShoppingListChangeNameAction struct {
	Name *LocalizedString `json:"name"`
}

func (obj ShoppingListChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type ShoppingListChangeTextLineItemNameAction struct {
	TextLineItemID string           `json:"textLineItemId"`
	Name           *LocalizedString `json:"name"`
}

func (obj ShoppingListChangeTextLineItemNameAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListChangeTextLineItemNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTextLineItemName", Alias: (*Alias)(&obj)})
}

type ShoppingListChangeTextLineItemQuantityAction struct {
	TextLineItemID string `json:"textLineItemId"`
	Quantity       int    `json:"quantity"`
}

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

func (obj ShoppingListChangeTextLineItemsOrderAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListChangeTextLineItemsOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTextLineItemsOrder", Alias: (*Alias)(&obj)})
}

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

type ShoppingListLineItemDraft struct {
	VariantID int                `json:"variantId,omitempty"`
	Sku       string             `json:"sku,omitempty"`
	Quantity  float64            `json:"quantity,omitempty"`
	ProductID string             `json:"productId,omitempty"`
	Custom    *CustomFieldsDraft `json:"custom,omitempty"`
	AddedAt   time.Time          `json:"addedAt,omitempty"`
}

type ShoppingListPagedQueryResponse struct {
	Total   int            `json:"total,omitempty"`
	Offset  int            `json:"offset"`
	Count   int            `json:"count"`
	Results []ShoppingList `json:"results"`
}

type ShoppingListReference struct {
	Key string        `json:"key,omitempty"`
	ID  string        `json:"id,omitempty"`
	Obj *ShoppingList `json:"obj,omitempty"`
}

func (obj ShoppingListReference) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "shopping-list", Alias: (*Alias)(&obj)})
}

type ShoppingListRemoveLineItemAction struct {
	Quantity   int    `json:"quantity,omitempty"`
	LineItemID string `json:"lineItemId"`
}

func (obj ShoppingListRemoveLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListRemoveLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeLineItem", Alias: (*Alias)(&obj)})
}

type ShoppingListRemoveTextLineItemAction struct {
	TextLineItemID string `json:"textLineItemId"`
	Quantity       int    `json:"quantity,omitempty"`
}

func (obj ShoppingListRemoveTextLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListRemoveTextLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeTextLineItem", Alias: (*Alias)(&obj)})
}

type ShoppingListSetAnonymousIdAction struct {
	AnonymousID string `json:"anonymousId,omitempty"`
}

func (obj ShoppingListSetAnonymousIdAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetAnonymousIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAnonymousId", Alias: (*Alias)(&obj)})
}

type ShoppingListSetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

func (obj ShoppingListSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type ShoppingListSetCustomTypeAction struct {
	Type   *TypeReference  `json:"type,omitempty"`
	Fields *FieldContainer `json:"fields,omitempty"`
}

func (obj ShoppingListSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type ShoppingListSetCustomerAction struct {
	Customer *CustomerReference `json:"customer,omitempty"`
}

func (obj ShoppingListSetCustomerAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetCustomerAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomer", Alias: (*Alias)(&obj)})
}

type ShoppingListSetDeleteDaysAfterLastModificationAction struct {
	DeleteDaysAfterLastModification int `json:"deleteDaysAfterLastModification,omitempty"`
}

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

func (obj ShoppingListSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type ShoppingListSetKeyAction struct {
	Key string `json:"key,omitempty"`
}

func (obj ShoppingListSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type ShoppingListSetLineItemCustomFieldAction struct {
	Value      interface{} `json:"value,omitempty"`
	Name       string      `json:"name"`
	LineItemID string      `json:"lineItemId"`
}

func (obj ShoppingListSetLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomField", Alias: (*Alias)(&obj)})
}

type ShoppingListSetLineItemCustomTypeAction struct {
	Type       *TypeReference  `json:"type,omitempty"`
	LineItemID string          `json:"lineItemId"`
	Fields     *FieldContainer `json:"fields,omitempty"`
}

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

func (obj ShoppingListSetSlugAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetSlugAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSlug", Alias: (*Alias)(&obj)})
}

type ShoppingListSetTextLineItemCustomFieldAction struct {
	Value          interface{} `json:"value,omitempty"`
	TextLineItemID string      `json:"textLineItemId"`
	Name           string      `json:"name"`
}

func (obj ShoppingListSetTextLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetTextLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTextLineItemCustomField", Alias: (*Alias)(&obj)})
}

type ShoppingListSetTextLineItemCustomTypeAction struct {
	Type           *TypeReference  `json:"type,omitempty"`
	TextLineItemID string          `json:"textLineItemId"`
	Fields         *FieldContainer `json:"fields,omitempty"`
}

func (obj ShoppingListSetTextLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetTextLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTextLineItemCustomType", Alias: (*Alias)(&obj)})
}

type ShoppingListSetTextLineItemDescriptionAction struct {
	TextLineItemID string           `json:"textLineItemId"`
	Description    *LocalizedString `json:"description,omitempty"`
}

func (obj ShoppingListSetTextLineItemDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetTextLineItemDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTextLineItemDescription", Alias: (*Alias)(&obj)})
}

type ShoppingListUpdate struct {
	Version int                        `json:"version"`
	Actions []ShoppingListUpdateAction `json:"actions"`
}

func (obj *ShoppingListUpdate) UnmarshalJSON(data []byte) error {
	type Alias ShoppingListUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		obj.Actions[i] = AbstractShoppingListUpdateActionDiscriminatorMapping(obj.Actions[i])
	}

	return nil
}

type ShoppingListUpdateAction interface{}
type AbstractShoppingListUpdateAction struct{}

func AbstractShoppingListUpdateActionDiscriminatorMapping(input ShoppingListUpdateAction) ShoppingListUpdateAction {
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
		new := ShoppingListSetAnonymousIdAction{}
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

type TextLineItem struct {
	Quantity    float64          `json:"quantity"`
	Name        *LocalizedString `json:"name"`
	ID          string           `json:"id"`
	Description *LocalizedString `json:"description,omitempty"`
	Custom      *CustomFields    `json:"custom,omitempty"`
	AddedAt     time.Time        `json:"addedAt"`
}

type TextLineItemDraft struct {
	Quantity    float64            `json:"quantity,omitempty"`
	Name        *LocalizedString   `json:"name"`
	Description *LocalizedString   `json:"description,omitempty"`
	Custom      *CustomFieldsDraft `json:"custom,omitempty"`
	AddedAt     time.Time          `json:"addedAt,omitempty"`
}
