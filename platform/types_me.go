// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type MyCartDraft struct {
	// A three-digit currency code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Currency      string  `json:"currency"`
	CustomerEmail *string `json:"customerEmail,omitempty"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country *string `json:"country,omitempty"`
	// Default inventory mode is `None`.
	InventoryMode   *InventoryMode                    `json:"inventoryMode,omitempty"`
	LineItems       []MyLineItemDraft                 `json:"lineItems,omitempty"`
	ShippingAddress *BaseAddress                      `json:"shippingAddress,omitempty"`
	BillingAddress  *BaseAddress                      `json:"billingAddress,omitempty"`
	ShippingMethod  *ShippingMethodResourceIdentifier `json:"shippingMethod,omitempty"`
	// The custom fields.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
	Locale *string            `json:"locale,omitempty"`
	// The `TaxMode` `Disabled` can not be set on the My Carts endpoint.
	TaxMode *TaxMode `json:"taxMode,omitempty"`
	// The cart will be deleted automatically if it hasn't been modified for the specified amount of days and it is in the `Active` CartState.
	// If a ChangeSubscription for carts exists, a `ResourceDeleted` notification will be sent.
	DeleteDaysAfterLastModification *int `json:"deleteDaysAfterLastModification,omitempty"`
	// Contains addresses for orders with multiple shipping addresses.
	// Each address must contain a key which is unique in this cart.
	ItemShippingAddresses []BaseAddress      `json:"itemShippingAddresses,omitempty"`
	Store                 *StoreKeyReference `json:"store,omitempty"`
	DiscountCodes         []DiscountCodeInfo `json:"discountCodes,omitempty"`
}

type MyCartUpdate struct {
	Version int                  `json:"version"`
	Actions []MyCartUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *MyCartUpdate) UnmarshalJSON(data []byte) error {
	type Alias MyCartUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	return nil
}

type MyCartUpdateAction interface{}

func mapDiscriminatorMyCartUpdateAction(input interface{}) (MyCartUpdateAction, error) {

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
	case "addDiscountCode":
		obj := MyCartAddDiscountCodeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addItemShippingAddress":
		obj := MyCartAddItemShippingAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addLineItem":
		obj := MyCartAddLineItemAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addPayment":
		obj := MyCartAddPaymentAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "applyDeltaToLineItemShippingDetailsTargets":
		obj := MyCartApplyDeltaToLineItemShippingDetailsTargetsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeLineItemQuantity":
		obj := MyCartChangeLineItemQuantityAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeTaxMode":
		obj := MyCartChangeTaxModeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "recalculate":
		obj := MyCartRecalculateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeDiscountCode":
		obj := MyCartRemoveDiscountCodeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeItemShippingAddress":
		obj := MyCartRemoveItemShippingAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeLineItem":
		obj := MyCartRemoveLineItemAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removePayment":
		obj := MyCartRemovePaymentAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setBillingAddress":
		obj := MyCartSetBillingAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCountry":
		obj := MyCartSetCountryAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomField":
		obj := MyCartSetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomType":
		obj := MyCartSetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomerEmail":
		obj := MyCartSetCustomerEmailAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDeleteDaysAfterLastModification":
		obj := MyCartSetDeleteDaysAfterLastModificationAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLineItemCustomField":
		obj := MyCartSetLineItemCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLineItemCustomType":
		obj := MyCartSetLineItemCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLineItemDistributionChannel":
		obj := MyCartSetLineItemDistributionChannelAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLineItemShippingDetails":
		obj := MyCartSetLineItemShippingDetailsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLineItemSupplyChannel":
		obj := MyCartSetLineItemSupplyChannelAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLocale":
		obj := MyCartSetLocaleAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setShippingAddress":
		obj := MyCartSetShippingAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setShippingMethod":
		obj := MyCartSetShippingMethodAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "updateItemShippingAddress":
		obj := MyCartUpdateItemShippingAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type MyCustomerDraft struct {
	Email       string  `json:"email"`
	Password    string  `json:"password"`
	FirstName   *string `json:"firstName,omitempty"`
	LastName    *string `json:"lastName,omitempty"`
	MiddleName  *string `json:"middleName,omitempty"`
	Title       *string `json:"title,omitempty"`
	DateOfBirth *Date   `json:"dateOfBirth,omitempty"`
	CompanyName *string `json:"companyName,omitempty"`
	VatId       *string `json:"vatId,omitempty"`
	// Sets the ID of each address to be unique in the addresses list.
	Addresses []BaseAddress `json:"addresses,omitempty"`
	// The index of the address in the addresses array.
	// The `defaultShippingAddressId` of the customer will be set to the ID of that address.
	DefaultShippingAddress *int `json:"defaultShippingAddress,omitempty"`
	// The index of the address in the addresses array.
	// The `defaultBillingAddressId` of the customer will be set to the ID of that address.
	DefaultBillingAddress *int `json:"defaultBillingAddress,omitempty"`
	// The custom fields.
	Custom *CustomFields             `json:"custom,omitempty"`
	Locale *string                   `json:"locale,omitempty"`
	Stores []StoreResourceIdentifier `json:"stores,omitempty"`
}

type MyCustomerUpdate struct {
	Version int                      `json:"version"`
	Actions []MyCustomerUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *MyCustomerUpdate) UnmarshalJSON(data []byte) error {
	type Alias MyCustomerUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	return nil
}

type MyCustomerUpdateAction interface{}

func mapDiscriminatorMyCustomerUpdateAction(input interface{}) (MyCustomerUpdateAction, error) {

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
	case "addAddress":
		obj := MyCustomerAddAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addBillingAddressId":
		obj := MyCustomerAddBillingAddressIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addShippingAddressId":
		obj := MyCustomerAddShippingAddressIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeAddress":
		obj := MyCustomerChangeAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeEmail":
		obj := MyCustomerChangeEmailAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeAddress":
		obj := MyCustomerRemoveAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeBillingAddressId":
		obj := MyCustomerRemoveBillingAddressIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeShippingAddressId":
		obj := MyCustomerRemoveShippingAddressIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCompanyName":
		obj := MyCustomerSetCompanyNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomField":
		obj := MyCustomerSetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomType":
		obj := MyCustomerSetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDateOfBirth":
		obj := MyCustomerSetDateOfBirthAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDefaultBillingAddress":
		obj := MyCustomerSetDefaultBillingAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDefaultShippingAddress":
		obj := MyCustomerSetDefaultShippingAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setFirstName":
		obj := MyCustomerSetFirstNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLastName":
		obj := MyCustomerSetLastNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLocale":
		obj := MyCustomerSetLocaleAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMiddleName":
		obj := MyCustomerSetMiddleNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setSalutation":
		obj := MyCustomerSetSalutationAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setTitle":
		obj := MyCustomerSetTitleAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setVatId":
		obj := MyCustomerSetVatIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type MyLineItemDraft struct {
	ProductId *string `json:"productId,omitempty"`
	VariantId *int    `json:"variantId,omitempty"`
	Quantity  int     `json:"quantity"`
	// When the line item was added to the cart. Optional for backwards
	// compatibility reasons only.
	AddedAt *time.Time `json:"addedAt,omitempty"`
	// By providing supply channel information, you can unique identify
	// inventory entries that should be reserved.
	// The provided channel should have the InventorySupply role.
	SupplyChannel *ChannelResourceIdentifier `json:"supplyChannel,omitempty"`
	// The channel is used to select a ProductPrice.
	// The provided channel should have the ProductDistribution role.
	DistributionChannel *ChannelResourceIdentifier `json:"distributionChannel,omitempty"`
	// The custom fields.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
	// Container for line item specific address(es).
	ShippingDetails *ItemShippingDetailsDraft `json:"shippingDetails,omitempty"`
	Sku             *string                   `json:"sku,omitempty"`
}

type MyOrderFromCartDraft struct {
	// The unique ID of the cart from which an order is created.
	ID      string `json:"id"`
	Version int    `json:"version"`
}

type MyPayment struct {
	ID      string `json:"id"`
	Version int    `json:"version"`
	// A reference to the customer this payment belongs to.
	Customer *CustomerReference `json:"customer,omitempty"`
	// Identifies payments belonging to an anonymous session (the customer has not signed up/in yet).
	AnonymousId *string `json:"anonymousId,omitempty"`
	// How much money this payment intends to receive from the customer.
	// The value usually matches the cart or order gross total.
	AmountPlanned     TypedMoney        `json:"amountPlanned"`
	PaymentMethodInfo PaymentMethodInfo `json:"paymentMethodInfo"`
	// A list of financial transactions of different TransactionTypes
	// with different TransactionStates.
	Transactions []Transaction `json:"transactions"`
	Custom       *CustomFields `json:"custom,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *MyPayment) UnmarshalJSON(data []byte) error {
	type Alias MyPayment
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.AmountPlanned != nil {
		var err error
		obj.AmountPlanned, err = mapDiscriminatorTypedMoney(obj.AmountPlanned)
		if err != nil {
			return err
		}
	}
	return nil
}

type MyPaymentDraft struct {
	// How much money this payment intends to receive from the customer.
	// The value usually matches the cart or order gross total.
	AmountPlanned     Money              `json:"amountPlanned"`
	PaymentMethodInfo *PaymentMethodInfo `json:"paymentMethodInfo,omitempty"`
	Custom            *CustomFieldsDraft `json:"custom,omitempty"`
	// A list of financial transactions of the `Authorization` or `Charge`
	// TransactionTypes.
	Transaction *MyTransactionDraft `json:"transaction,omitempty"`
}

type MyPaymentPagedQueryResponse struct {
	Limit   int         `json:"limit"`
	Count   int         `json:"count"`
	Total   *int        `json:"total,omitempty"`
	Offset  int         `json:"offset"`
	Results []MyPayment `json:"results"`
}

type MyPaymentUpdate struct {
	Version int                     `json:"version"`
	Actions []MyPaymentUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *MyPaymentUpdate) UnmarshalJSON(data []byte) error {
	type Alias MyPaymentUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	return nil
}

type MyPaymentUpdateAction interface{}

func mapDiscriminatorMyPaymentUpdateAction(input interface{}) (MyPaymentUpdateAction, error) {

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
	case "addTransaction":
		obj := MyPaymentAddTransactionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeAmountPlanned":
		obj := MyPaymentChangeAmountPlannedAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomField":
		obj := MyPaymentSetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMethodInfoInterface":
		obj := MyPaymentSetMethodInfoInterfaceAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMethodInfoMethod":
		obj := MyPaymentSetMethodInfoMethodAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMethodInfoName":
		obj := MyPaymentSetMethodInfoNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type MyShoppingListDraft struct {
	Name          LocalizedString             `json:"name"`
	Description   *LocalizedString            `json:"description,omitempty"`
	LineItems     []ShoppingListLineItemDraft `json:"lineItems,omitempty"`
	TextLineItems []TextLineItemDraft         `json:"textLineItems,omitempty"`
	// The custom fields.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
	// The shopping list will be deleted automatically if it hasn't been modified for the specified amount of days.
	DeleteDaysAfterLastModification *int                     `json:"deleteDaysAfterLastModification,omitempty"`
	Store                           *StoreResourceIdentifier `json:"store,omitempty"`
}

type MyShoppingListUpdate struct {
	Version int                          `json:"version"`
	Actions []MyShoppingListUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *MyShoppingListUpdate) UnmarshalJSON(data []byte) error {
	type Alias MyShoppingListUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	return nil
}

type MyShoppingListUpdateAction interface{}

func mapDiscriminatorMyShoppingListUpdateAction(input interface{}) (MyShoppingListUpdateAction, error) {

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
		obj := MyShoppingListAddLineItemAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addTextLineItem":
		obj := MyShoppingListAddTextLineItemAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeLineItemQuantity":
		obj := MyShoppingListChangeLineItemQuantityAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeLineItemsOrder":
		obj := MyShoppingListChangeLineItemsOrderAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeName":
		obj := MyShoppingListChangeNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeTextLineItemName":
		obj := MyShoppingListChangeTextLineItemNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeTextLineItemQuantity":
		obj := MyShoppingListChangeTextLineItemQuantityAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeTextLineItemsOrder":
		obj := MyShoppingListChangeTextLineItemsOrderAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeLineItem":
		obj := MyShoppingListRemoveLineItemAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeTextLineItem":
		obj := MyShoppingListRemoveTextLineItemAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomField":
		obj := MyShoppingListSetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomType":
		obj := MyShoppingListSetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDeleteDaysAfterLastModification":
		obj := MyShoppingListSetDeleteDaysAfterLastModificationAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDescription":
		obj := MyShoppingListSetDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLineItemCustomField":
		obj := MyShoppingListSetLineItemCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLineItemCustomType":
		obj := MyShoppingListSetLineItemCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setTextLineItemCustomField":
		obj := MyShoppingListSetTextLineItemCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setTextLineItemCustomType":
		obj := MyShoppingListSetTextLineItemCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setTextLineItemDescription":
		obj := MyShoppingListSetTextLineItemDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type MyTransactionDraft struct {
	// The time at which the transaction took place.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// The type of this transaction.
	// Only the `Authorization` or `Charge`
	// TransactionTypes are allowed here.
	Type   TransactionType `json:"type"`
	Amount Money           `json:"amount"`
	// The identifier that is used by the interface that managed the transaction (usually the PSP).
	// If a matching interaction was logged in the interfaceInteractions array,
	// the corresponding interaction should be findable with this ID.
	// The `state` is set to the `Initial` TransactionState.
	InteractionId *string `json:"interactionId,omitempty"`
}

type MyCartAddDiscountCodeAction struct {
	Code string `json:"code"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartAddDiscountCodeAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartAddDiscountCodeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addDiscountCode", Alias: (*Alias)(&obj)})
}

type MyCartAddItemShippingAddressAction struct {
	Address BaseAddress `json:"address"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartAddItemShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartAddItemShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addItemShippingAddress", Alias: (*Alias)(&obj)})
}

type MyCartAddLineItemAction struct {
	Custom              *CustomFieldsDraft          `json:"custom,omitempty"`
	DistributionChannel *ChannelResourceIdentifier  `json:"distributionChannel,omitempty"`
	ExternalTaxRate     *ExternalTaxRateDraft       `json:"externalTaxRate,omitempty"`
	ProductId           *string                     `json:"productId,omitempty"`
	VariantId           *int                        `json:"variantId,omitempty"`
	Sku                 *string                     `json:"sku,omitempty"`
	Quantity            *int                        `json:"quantity,omitempty"`
	SupplyChannel       *ChannelResourceIdentifier  `json:"supplyChannel,omitempty"`
	ExternalPrice       *Money                      `json:"externalPrice,omitempty"`
	ExternalTotalPrice  *ExternalLineItemTotalPrice `json:"externalTotalPrice,omitempty"`
	ShippingDetails     *ItemShippingDetailsDraft   `json:"shippingDetails,omitempty"`
	AddedAt             *time.Time                  `json:"addedAt,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartAddLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartAddLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addLineItem", Alias: (*Alias)(&obj)})
}

type MyCartAddPaymentAction struct {
	Payment PaymentResourceIdentifier `json:"payment"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartAddPaymentAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartAddPaymentAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addPayment", Alias: (*Alias)(&obj)})
}

type MyCartApplyDeltaToLineItemShippingDetailsTargetsAction struct {
	LineItemId   string               `json:"lineItemId"`
	TargetsDelta []ItemShippingTarget `json:"targetsDelta"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartApplyDeltaToLineItemShippingDetailsTargetsAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartApplyDeltaToLineItemShippingDetailsTargetsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "applyDeltaToLineItemShippingDetailsTargets", Alias: (*Alias)(&obj)})
}

type MyCartChangeLineItemQuantityAction struct {
	LineItemId         string                      `json:"lineItemId"`
	Quantity           int                         `json:"quantity"`
	ExternalPrice      *Money                      `json:"externalPrice,omitempty"`
	ExternalTotalPrice *ExternalLineItemTotalPrice `json:"externalTotalPrice,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartChangeLineItemQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartChangeLineItemQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLineItemQuantity", Alias: (*Alias)(&obj)})
}

type MyCartChangeTaxModeAction struct {
	TaxMode TaxMode `json:"taxMode"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartChangeTaxModeAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartChangeTaxModeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTaxMode", Alias: (*Alias)(&obj)})
}

type MyCartRecalculateAction struct {
	UpdateProductData *bool `json:"updateProductData,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartRecalculateAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartRecalculateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "recalculate", Alias: (*Alias)(&obj)})
}

type MyCartRemoveDiscountCodeAction struct {
	DiscountCode DiscountCodeReference `json:"discountCode"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartRemoveDiscountCodeAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartRemoveDiscountCodeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeDiscountCode", Alias: (*Alias)(&obj)})
}

type MyCartRemoveItemShippingAddressAction struct {
	AddressKey string `json:"addressKey"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartRemoveItemShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartRemoveItemShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeItemShippingAddress", Alias: (*Alias)(&obj)})
}

type MyCartRemoveLineItemAction struct {
	LineItemId              string                      `json:"lineItemId"`
	Quantity                *int                        `json:"quantity,omitempty"`
	ExternalPrice           *Money                      `json:"externalPrice,omitempty"`
	ExternalTotalPrice      *ExternalLineItemTotalPrice `json:"externalTotalPrice,omitempty"`
	ShippingDetailsToRemove *ItemShippingDetailsDraft   `json:"shippingDetailsToRemove,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartRemoveLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartRemoveLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeLineItem", Alias: (*Alias)(&obj)})
}

type MyCartRemovePaymentAction struct {
	Payment PaymentResourceIdentifier `json:"payment"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartRemovePaymentAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartRemovePaymentAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removePayment", Alias: (*Alias)(&obj)})
}

type MyCartSetBillingAddressAction struct {
	Address *BaseAddress `json:"address,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartSetBillingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartSetBillingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setBillingAddress", Alias: (*Alias)(&obj)})
}

type MyCartSetCountryAction struct {
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country *string `json:"country,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartSetCountryAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartSetCountryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCountry", Alias: (*Alias)(&obj)})
}

type MyCartSetCustomFieldAction struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type MyCartSetCustomTypeAction struct {
	Type   *TypeResourceIdentifier `json:"type,omitempty"`
	Fields *FieldContainer         `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type MyCartSetCustomerEmailAction struct {
	Email *string `json:"email,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartSetCustomerEmailAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartSetCustomerEmailAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomerEmail", Alias: (*Alias)(&obj)})
}

type MyCartSetDeleteDaysAfterLastModificationAction struct {
	DeleteDaysAfterLastModification *int `json:"deleteDaysAfterLastModification,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartSetDeleteDaysAfterLastModificationAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartSetDeleteDaysAfterLastModificationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeleteDaysAfterLastModification", Alias: (*Alias)(&obj)})
}

type MyCartSetLineItemCustomFieldAction struct {
	LineItemId string      `json:"lineItemId"`
	Name       string      `json:"name"`
	Value      interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartSetLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartSetLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomField", Alias: (*Alias)(&obj)})
}

type MyCartSetLineItemCustomTypeAction struct {
	LineItemId string                  `json:"lineItemId"`
	Type       *TypeResourceIdentifier `json:"type,omitempty"`
	Fields     *FieldContainer         `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartSetLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartSetLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomType", Alias: (*Alias)(&obj)})
}

type MyCartSetLineItemDistributionChannelAction struct {
	LineItemId          string                     `json:"lineItemId"`
	DistributionChannel *ChannelResourceIdentifier `json:"distributionChannel,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartSetLineItemDistributionChannelAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartSetLineItemDistributionChannelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemDistributionChannel", Alias: (*Alias)(&obj)})
}

type MyCartSetLineItemShippingDetailsAction struct {
	LineItemId      string                    `json:"lineItemId"`
	ShippingDetails *ItemShippingDetailsDraft `json:"shippingDetails,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartSetLineItemShippingDetailsAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartSetLineItemShippingDetailsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemShippingDetails", Alias: (*Alias)(&obj)})
}

type MyCartSetLineItemSupplyChannelAction struct {
	LineItemId          string                     `json:"lineItemId"`
	DistributionChannel *ChannelResourceIdentifier `json:"distributionChannel,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartSetLineItemSupplyChannelAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartSetLineItemSupplyChannelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemSupplyChannel", Alias: (*Alias)(&obj)})
}

type MyCartSetLocaleAction struct {
	Locale *string `json:"locale,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartSetLocaleAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartSetLocaleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLocale", Alias: (*Alias)(&obj)})
}

type MyCartSetShippingAddressAction struct {
	Address *BaseAddress `json:"address,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartSetShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartSetShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingAddress", Alias: (*Alias)(&obj)})
}

type MyCartSetShippingMethodAction struct {
	ShippingMethod  *ShippingMethodResourceIdentifier `json:"shippingMethod,omitempty"`
	ExternalTaxRate *ExternalTaxRateDraft             `json:"externalTaxRate,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartSetShippingMethodAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartSetShippingMethodAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingMethod", Alias: (*Alias)(&obj)})
}

type MyCartUpdateItemShippingAddressAction struct {
	Address BaseAddress `json:"address"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartUpdateItemShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartUpdateItemShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "updateItemShippingAddress", Alias: (*Alias)(&obj)})
}

type MyCustomerAddAddressAction struct {
	Address BaseAddress `json:"address"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerAddAddressAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerAddAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addAddress", Alias: (*Alias)(&obj)})
}

type MyCustomerAddBillingAddressIdAction struct {
	AddressId  *string `json:"addressId,omitempty"`
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerAddBillingAddressIdAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerAddBillingAddressIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addBillingAddressId", Alias: (*Alias)(&obj)})
}

type MyCustomerAddShippingAddressIdAction struct {
	AddressId  *string `json:"addressId,omitempty"`
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerAddShippingAddressIdAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerAddShippingAddressIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addShippingAddressId", Alias: (*Alias)(&obj)})
}

type MyCustomerChangeAddressAction struct {
	AddressId  *string     `json:"addressId,omitempty"`
	AddressKey *string     `json:"addressKey,omitempty"`
	Address    BaseAddress `json:"address"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerChangeAddressAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerChangeAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAddress", Alias: (*Alias)(&obj)})
}

type MyCustomerChangeEmailAction struct {
	Email string `json:"email"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerChangeEmailAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerChangeEmailAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeEmail", Alias: (*Alias)(&obj)})
}

type MyCustomerRemoveAddressAction struct {
	AddressId  *string `json:"addressId,omitempty"`
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerRemoveAddressAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerRemoveAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeAddress", Alias: (*Alias)(&obj)})
}

type MyCustomerRemoveBillingAddressIdAction struct {
	AddressId  *string `json:"addressId,omitempty"`
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerRemoveBillingAddressIdAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerRemoveBillingAddressIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeBillingAddressId", Alias: (*Alias)(&obj)})
}

type MyCustomerRemoveShippingAddressIdAction struct {
	AddressId  *string `json:"addressId,omitempty"`
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerRemoveShippingAddressIdAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerRemoveShippingAddressIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeShippingAddressId", Alias: (*Alias)(&obj)})
}

type MyCustomerSetCompanyNameAction struct {
	CompanyName *string `json:"companyName,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerSetCompanyNameAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetCompanyNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCompanyName", Alias: (*Alias)(&obj)})
}

type MyCustomerSetCustomFieldAction struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type MyCustomerSetCustomTypeAction struct {
	Type   *TypeResourceIdentifier `json:"type,omitempty"`
	Fields *FieldContainer         `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type MyCustomerSetDateOfBirthAction struct {
	DateOfBirth *Date `json:"dateOfBirth,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerSetDateOfBirthAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetDateOfBirthAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDateOfBirth", Alias: (*Alias)(&obj)})
}

type MyCustomerSetDefaultBillingAddressAction struct {
	AddressId  *string `json:"addressId,omitempty"`
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerSetDefaultBillingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetDefaultBillingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDefaultBillingAddress", Alias: (*Alias)(&obj)})
}

type MyCustomerSetDefaultShippingAddressAction struct {
	AddressId  *string `json:"addressId,omitempty"`
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerSetDefaultShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetDefaultShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDefaultShippingAddress", Alias: (*Alias)(&obj)})
}

type MyCustomerSetFirstNameAction struct {
	FirstName *string `json:"firstName,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerSetFirstNameAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetFirstNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setFirstName", Alias: (*Alias)(&obj)})
}

type MyCustomerSetLastNameAction struct {
	LastName *string `json:"lastName,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerSetLastNameAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetLastNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLastName", Alias: (*Alias)(&obj)})
}

type MyCustomerSetLocaleAction struct {
	Locale *string `json:"locale,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerSetLocaleAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetLocaleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLocale", Alias: (*Alias)(&obj)})
}

type MyCustomerSetMiddleNameAction struct {
	MiddleName *string `json:"middleName,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerSetMiddleNameAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetMiddleNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMiddleName", Alias: (*Alias)(&obj)})
}

type MyCustomerSetSalutationAction struct {
	Salutation *string `json:"salutation,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerSetSalutationAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetSalutationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSalutation", Alias: (*Alias)(&obj)})
}

type MyCustomerSetTitleAction struct {
	Title *string `json:"title,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerSetTitleAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetTitleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTitle", Alias: (*Alias)(&obj)})
}

type MyCustomerSetVatIdAction struct {
	VatId *string `json:"vatId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerSetVatIdAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetVatIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setVatId", Alias: (*Alias)(&obj)})
}

type MyPaymentAddTransactionAction struct {
	Transaction TransactionDraft `json:"transaction"`
}

// MarshalJSON override to set the discriminator value
func (obj MyPaymentAddTransactionAction) MarshalJSON() ([]byte, error) {
	type Alias MyPaymentAddTransactionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addTransaction", Alias: (*Alias)(&obj)})
}

type MyPaymentChangeAmountPlannedAction struct {
	Amount Money `json:"amount"`
}

// MarshalJSON override to set the discriminator value
func (obj MyPaymentChangeAmountPlannedAction) MarshalJSON() ([]byte, error) {
	type Alias MyPaymentChangeAmountPlannedAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAmountPlanned", Alias: (*Alias)(&obj)})
}

type MyPaymentSetCustomFieldAction struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyPaymentSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias MyPaymentSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type MyPaymentSetMethodInfoInterfaceAction struct {
	Interface string `json:"interface"`
}

// MarshalJSON override to set the discriminator value
func (obj MyPaymentSetMethodInfoInterfaceAction) MarshalJSON() ([]byte, error) {
	type Alias MyPaymentSetMethodInfoInterfaceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMethodInfoInterface", Alias: (*Alias)(&obj)})
}

type MyPaymentSetMethodInfoMethodAction struct {
	Method *string `json:"method,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyPaymentSetMethodInfoMethodAction) MarshalJSON() ([]byte, error) {
	type Alias MyPaymentSetMethodInfoMethodAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMethodInfoMethod", Alias: (*Alias)(&obj)})
}

type MyPaymentSetMethodInfoNameAction struct {
	Name *LocalizedString `json:"name,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyPaymentSetMethodInfoNameAction) MarshalJSON() ([]byte, error) {
	type Alias MyPaymentSetMethodInfoNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMethodInfoName", Alias: (*Alias)(&obj)})
}

type MyShoppingListAddLineItemAction struct {
	Sku       *string            `json:"sku,omitempty"`
	ProductId *string            `json:"productId,omitempty"`
	VariantId *int               `json:"variantId,omitempty"`
	Quantity  *int               `json:"quantity,omitempty"`
	AddedAt   *time.Time         `json:"addedAt,omitempty"`
	Custom    *CustomFieldsDraft `json:"custom,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListAddLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListAddLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addLineItem", Alias: (*Alias)(&obj)})
}

type MyShoppingListAddTextLineItemAction struct {
	Name        LocalizedString    `json:"name"`
	Description *LocalizedString   `json:"description,omitempty"`
	Quantity    *int               `json:"quantity,omitempty"`
	AddedAt     *time.Time         `json:"addedAt,omitempty"`
	Custom      *CustomFieldsDraft `json:"custom,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListAddTextLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListAddTextLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addTextLineItem", Alias: (*Alias)(&obj)})
}

type MyShoppingListChangeLineItemQuantityAction struct {
	LineItemId string `json:"lineItemId"`
	Quantity   int    `json:"quantity"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListChangeLineItemQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListChangeLineItemQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLineItemQuantity", Alias: (*Alias)(&obj)})
}

type MyShoppingListChangeLineItemsOrderAction struct {
	LineItemOrder []string `json:"lineItemOrder"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListChangeLineItemsOrderAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListChangeLineItemsOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLineItemsOrder", Alias: (*Alias)(&obj)})
}

type MyShoppingListChangeNameAction struct {
	Name LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type MyShoppingListChangeTextLineItemNameAction struct {
	TextLineItemId string          `json:"textLineItemId"`
	Name           LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListChangeTextLineItemNameAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListChangeTextLineItemNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTextLineItemName", Alias: (*Alias)(&obj)})
}

type MyShoppingListChangeTextLineItemQuantityAction struct {
	TextLineItemId string `json:"textLineItemId"`
	Quantity       int    `json:"quantity"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListChangeTextLineItemQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListChangeTextLineItemQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTextLineItemQuantity", Alias: (*Alias)(&obj)})
}

type MyShoppingListChangeTextLineItemsOrderAction struct {
	TextLineItemOrder []string `json:"textLineItemOrder"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListChangeTextLineItemsOrderAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListChangeTextLineItemsOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTextLineItemsOrder", Alias: (*Alias)(&obj)})
}

type MyShoppingListRemoveLineItemAction struct {
	LineItemId string `json:"lineItemId"`
	Quantity   *int   `json:"quantity,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListRemoveLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListRemoveLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeLineItem", Alias: (*Alias)(&obj)})
}

type MyShoppingListRemoveTextLineItemAction struct {
	TextLineItemId string `json:"textLineItemId"`
	Quantity       *int   `json:"quantity,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListRemoveTextLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListRemoveTextLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeTextLineItem", Alias: (*Alias)(&obj)})
}

type MyShoppingListSetCustomFieldAction struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type MyShoppingListSetCustomTypeAction struct {
	Type   *TypeResourceIdentifier `json:"type,omitempty"`
	Fields *FieldContainer         `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type MyShoppingListSetDeleteDaysAfterLastModificationAction struct {
	DeleteDaysAfterLastModification *int `json:"deleteDaysAfterLastModification,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListSetDeleteDaysAfterLastModificationAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListSetDeleteDaysAfterLastModificationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeleteDaysAfterLastModification", Alias: (*Alias)(&obj)})
}

type MyShoppingListSetDescriptionAction struct {
	Description *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type MyShoppingListSetLineItemCustomFieldAction struct {
	LineItemId string      `json:"lineItemId"`
	Name       string      `json:"name"`
	Value      interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListSetLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListSetLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomField", Alias: (*Alias)(&obj)})
}

type MyShoppingListSetLineItemCustomTypeAction struct {
	LineItemId string                  `json:"lineItemId"`
	Type       *TypeResourceIdentifier `json:"type,omitempty"`
	Fields     *FieldContainer         `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListSetLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListSetLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomType", Alias: (*Alias)(&obj)})
}

type MyShoppingListSetTextLineItemCustomFieldAction struct {
	TextLineItemId string      `json:"textLineItemId"`
	Name           string      `json:"name"`
	Value          interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListSetTextLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListSetTextLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTextLineItemCustomField", Alias: (*Alias)(&obj)})
}

type MyShoppingListSetTextLineItemCustomTypeAction struct {
	TextLineItemId string                  `json:"textLineItemId"`
	Type           *TypeResourceIdentifier `json:"type,omitempty"`
	Fields         *FieldContainer         `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListSetTextLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListSetTextLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTextLineItemCustomType", Alias: (*Alias)(&obj)})
}

type MyShoppingListSetTextLineItemDescriptionAction struct {
	TextLineItemId string           `json:"textLineItemId"`
	Description    *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListSetTextLineItemDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListSetTextLineItemDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTextLineItemDescription", Alias: (*Alias)(&obj)})
}
