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
	LineItems []ShoppingListLineItem `json:"lineItems"`
	Name      LocalizedString        `json:"name"`
	// Human-readable identifiers usually used as deep-link URL to the related shopping list.
	// Each slug is unique across a project, but a shopping list can have the same slug for different languages.
	// The slug must match the pattern [a-zA-Z0-9_-]{2,256}.
	Slug          *LocalizedString `json:"slug,omitempty"`
	TextLineItems []TextLineItem   `json:"textLineItems"`
	// Identifies shopping lists belonging to an anonymous session (the customer has not signed up/in yet).
	AnonymousId *string            `json:"anonymousId,omitempty"`
	Store       *StoreKeyReference `json:"store,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShoppingList) MarshalJSON() ([]byte, error) {
	type Alias ShoppingList
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["lineItems"] == nil {
		delete(target, "lineItems")
	}

	if target["textLineItems"] == nil {
		delete(target, "textLineItems")
	}

	return json.Marshal(target)
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
	LineItems []ShoppingListLineItemDraft `json:"lineItems"`
	Name      LocalizedString             `json:"name"`
	// Human-readable identifiers usually used as deep-link URL to the related shopping list.
	// Each slug is unique across a project, but a shopping list can have the same slug for different languages.
	// The slug must match the pattern [a-zA-Z0-9_-]{2,256}.
	Slug          *LocalizedString    `json:"slug,omitempty"`
	TextLineItems []TextLineItemDraft `json:"textLineItems"`
	// Identifies shopping lists belonging to an anonymous session (the customer has not signed up/in yet).
	AnonymousId *string                  `json:"anonymousId,omitempty"`
	Store       *StoreResourceIdentifier `json:"store,omitempty"`
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

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["lineItems"] == nil {
		delete(target, "lineItems")
	}

	if target["textLineItems"] == nil {
		delete(target, "textLineItems")
	}

	return json.Marshal(target)
}

type ShoppingListLineItem struct {
	AddedAt time.Time `json:"addedAt"`
	// Serves as value of the `custom` field on a resource or data type customized with a [Type](ctp:api:type:Type).
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
	AddedAt *time.Time `json:"addedAt,omitempty"`
	// The representation used when creating or updating a [customizable data type](/../api/projects/types#list-of-customizable-data-types) with Custom Fields.
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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
	Sku       *string    `json:"sku,omitempty"`
	ProductId *string    `json:"productId,omitempty"`
	VariantId *int       `json:"variantId,omitempty"`
	Quantity  *int       `json:"quantity,omitempty"`
	AddedAt   *time.Time `json:"addedAt,omitempty"`
	// The representation used when creating or updating a [customizable data type](/../api/projects/types#list-of-customizable-data-types) with Custom Fields.
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
	Name        LocalizedString  `json:"name"`
	Description *LocalizedString `json:"description,omitempty"`
	Quantity    *int             `json:"quantity,omitempty"`
	AddedAt     *time.Time       `json:"addedAt,omitempty"`
	// The representation used when creating or updating a [customizable data type](/../api/projects/types#list-of-customizable-data-types) with Custom Fields.
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
	LineItemId string `json:"lineItemId"`
	Quantity   int    `json:"quantity"`
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
	TextLineItemId string          `json:"textLineItemId"`
	Name           LocalizedString `json:"name"`
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
	TextLineItemId string `json:"textLineItemId"`
	Quantity       int    `json:"quantity"`
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

type ShoppingListRemoveLineItemAction struct {
	LineItemId string `json:"lineItemId"`
	Quantity   *int   `json:"quantity,omitempty"`
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
	TextLineItemId string `json:"textLineItemId"`
	Quantity       *int   `json:"quantity,omitempty"`
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

type ShoppingListSetAnonymousIdAction struct {
	// Anonymous ID of the anonymous customer that this shopping list belongs to.
	// If this field is not set any existing `anonymousId` is removed.
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

type ShoppingListSetCustomFieldAction struct {
	Name string `json:"name"`
	// The value of a Custom Field.
	// The data type of the value depends on the specific [FieldType](/projects/types#fieldtype) given in the `type` field of the [FieldDefinition](/ctp:api:type:FieldDefinition) for a Custom Field.
	// It can be any of the following:
	//
	//  Field type                                                 | Data type                                                                                                                                                                 |
	//  ---------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
	//  [CustomFieldBooleanType](ctp:api:type:CustomFieldBooleanType)                 | Boolean (`true` or `false`)                                                                                                                                                     |
	//  [CustomFieldStringType](ctp:api:type:CustomFieldStringType)                   | String                                                                                                                                                                |
	//  [CustomFieldLocalizedStringType](ctp:api:type:CustomFieldLocalizedStringType) | [LocalizedString](ctp:api:type:LocalizedString)                                                                                                                             |
	//  [CustomFieldEnumType](ctp:api:type:CustomFieldEnumType)                       | String. Must be a `key` of one of the [EnumValues](ctp:api:type:CustomFieldEnumValue) defined in the [EnumType](ctp:api:type:CustomFiedEnumType)                                     |
	//  [CustomFieldLocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType)     | String. Must be a `key` of one of the [LocalizedEnumValues](ctp:api:type:CustomFieldLocalizedEnumValue) defined in the [LocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType) |
	//  [CustomFieldNumberType](ctp:api:type:CustomFieldNumberType)                   | Number                                                                                                                                                                |
	//  [CustomFieldMoneyType](ctp:api:type:CustomFieldMoneyType)                     | [CentPrecisionMoney](/../api/types#centprecisionmoney)                                                                                                                                         |
	//  [CustomFieldDateType](ctp:api:type:CustomFieldDateType)                       | [Date](ctp:api:type:Date)                                                                                                                                                   |
	//  [CustomFieldTimeType](ctp:api:type:CustomFieldTimeType)                       | [Time](ctp:api:type:Time)                                                                                                                                                   |
	//  [CustomFieldDateTimeType](ctp:api:type:CustomFieldDateTimeType)               | [DateTime](ctp:api:type:DateTime)                                                                                                                                           |
	//  [CustomFieldReferenceType](ctp:api:type:CustomFieldReferenceType)             | [Reference](/../api/types#reference)                                                                                                                                         |
	//  [CustomFieldSetType](ctp:api:type:CustomFieldSetType)                         | JSON array without duplicates consisting of [CustomFieldValues](ctp:api:type:CustomFieldValue) of a single [FieldType](ctp:api:type:FieldType). For example, a Custom Field of SetType of DateType takes a JSON array of mutually different Dates for its value. The order of items in the array is not fixed. For more examples, see the [example FieldContainer](ctp:api:type:FieldContainer).|
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
	// If set, the custom type is set to this new value.
	// If absent, the custom type and any existing custom fields are removed.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// If set, the custom fields are set to this new value.
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

type ShoppingListSetDeleteDaysAfterLastModificationAction struct {
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
	// User-specific unique identifier for the shopping list.
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
	LineItemId string `json:"lineItemId"`
	Name       string `json:"name"`
	// The value of a Custom Field.
	// The data type of the value depends on the specific [FieldType](/projects/types#fieldtype) given in the `type` field of the [FieldDefinition](/ctp:api:type:FieldDefinition) for a Custom Field.
	// It can be any of the following:
	//
	//  Field type                                                 | Data type                                                                                                                                                                 |
	//  ---------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
	//  [CustomFieldBooleanType](ctp:api:type:CustomFieldBooleanType)                 | Boolean (`true` or `false`)                                                                                                                                                     |
	//  [CustomFieldStringType](ctp:api:type:CustomFieldStringType)                   | String                                                                                                                                                                |
	//  [CustomFieldLocalizedStringType](ctp:api:type:CustomFieldLocalizedStringType) | [LocalizedString](ctp:api:type:LocalizedString)                                                                                                                             |
	//  [CustomFieldEnumType](ctp:api:type:CustomFieldEnumType)                       | String. Must be a `key` of one of the [EnumValues](ctp:api:type:CustomFieldEnumValue) defined in the [EnumType](ctp:api:type:CustomFiedEnumType)                                     |
	//  [CustomFieldLocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType)     | String. Must be a `key` of one of the [LocalizedEnumValues](ctp:api:type:CustomFieldLocalizedEnumValue) defined in the [LocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType) |
	//  [CustomFieldNumberType](ctp:api:type:CustomFieldNumberType)                   | Number                                                                                                                                                                |
	//  [CustomFieldMoneyType](ctp:api:type:CustomFieldMoneyType)                     | [CentPrecisionMoney](/../api/types#centprecisionmoney)                                                                                                                                         |
	//  [CustomFieldDateType](ctp:api:type:CustomFieldDateType)                       | [Date](ctp:api:type:Date)                                                                                                                                                   |
	//  [CustomFieldTimeType](ctp:api:type:CustomFieldTimeType)                       | [Time](ctp:api:type:Time)                                                                                                                                                   |
	//  [CustomFieldDateTimeType](ctp:api:type:CustomFieldDateTimeType)               | [DateTime](ctp:api:type:DateTime)                                                                                                                                           |
	//  [CustomFieldReferenceType](ctp:api:type:CustomFieldReferenceType)             | [Reference](/../api/types#reference)                                                                                                                                         |
	//  [CustomFieldSetType](ctp:api:type:CustomFieldSetType)                         | JSON array without duplicates consisting of [CustomFieldValues](ctp:api:type:CustomFieldValue) of a single [FieldType](ctp:api:type:FieldType). For example, a Custom Field of SetType of DateType takes a JSON array of mutually different Dates for its value. The order of items in the array is not fixed. For more examples, see the [example FieldContainer](ctp:api:type:FieldContainer).|
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
	LineItemId string `json:"lineItemId"`
	// [ResourceIdentifier](/../api/types#resourceidentifier) of a [Type](ctp:api:type:Type).
	Type   *TypeResourceIdentifier `json:"type,omitempty"`
	Fields *FieldContainer         `json:"fields,omitempty"`
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
	TextLineItemId string `json:"textLineItemId"`
	Name           string `json:"name"`
	// The value of a Custom Field.
	// The data type of the value depends on the specific [FieldType](/projects/types#fieldtype) given in the `type` field of the [FieldDefinition](/ctp:api:type:FieldDefinition) for a Custom Field.
	// It can be any of the following:
	//
	//  Field type                                                 | Data type                                                                                                                                                                 |
	//  ---------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
	//  [CustomFieldBooleanType](ctp:api:type:CustomFieldBooleanType)                 | Boolean (`true` or `false`)                                                                                                                                                     |
	//  [CustomFieldStringType](ctp:api:type:CustomFieldStringType)                   | String                                                                                                                                                                |
	//  [CustomFieldLocalizedStringType](ctp:api:type:CustomFieldLocalizedStringType) | [LocalizedString](ctp:api:type:LocalizedString)                                                                                                                             |
	//  [CustomFieldEnumType](ctp:api:type:CustomFieldEnumType)                       | String. Must be a `key` of one of the [EnumValues](ctp:api:type:CustomFieldEnumValue) defined in the [EnumType](ctp:api:type:CustomFiedEnumType)                                     |
	//  [CustomFieldLocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType)     | String. Must be a `key` of one of the [LocalizedEnumValues](ctp:api:type:CustomFieldLocalizedEnumValue) defined in the [LocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType) |
	//  [CustomFieldNumberType](ctp:api:type:CustomFieldNumberType)                   | Number                                                                                                                                                                |
	//  [CustomFieldMoneyType](ctp:api:type:CustomFieldMoneyType)                     | [CentPrecisionMoney](/../api/types#centprecisionmoney)                                                                                                                                         |
	//  [CustomFieldDateType](ctp:api:type:CustomFieldDateType)                       | [Date](ctp:api:type:Date)                                                                                                                                                   |
	//  [CustomFieldTimeType](ctp:api:type:CustomFieldTimeType)                       | [Time](ctp:api:type:Time)                                                                                                                                                   |
	//  [CustomFieldDateTimeType](ctp:api:type:CustomFieldDateTimeType)               | [DateTime](ctp:api:type:DateTime)                                                                                                                                           |
	//  [CustomFieldReferenceType](ctp:api:type:CustomFieldReferenceType)             | [Reference](/../api/types#reference)                                                                                                                                         |
	//  [CustomFieldSetType](ctp:api:type:CustomFieldSetType)                         | JSON array without duplicates consisting of [CustomFieldValues](ctp:api:type:CustomFieldValue) of a single [FieldType](ctp:api:type:FieldType). For example, a Custom Field of SetType of DateType takes a JSON array of mutually different Dates for its value. The order of items in the array is not fixed. For more examples, see the [example FieldContainer](ctp:api:type:FieldContainer).|
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
	TextLineItemId string `json:"textLineItemId"`
	// [ResourceIdentifier](/../api/types#resourceidentifier) of a [Type](ctp:api:type:Type).
	Type   *TypeResourceIdentifier `json:"type,omitempty"`
	Fields *FieldContainer         `json:"fields,omitempty"`
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
	TextLineItemId string           `json:"textLineItemId"`
	Description    *LocalizedString `json:"description,omitempty"`
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
