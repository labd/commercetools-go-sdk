// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type MyCartDraft struct {
	// A three-digit currency code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Currency      string `json:"currency"`
	CustomerEmail string `json:"customerEmail,omitEmpty"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country string `json:"country,omitEmpty"`
	// Default inventory mode is `None`.
	InventoryMode   InventoryMode                    `json:"inventoryMode,omitEmpty"`
	LineItems       []MyLineItemDraft                `json:"lineItems,omitEmpty"`
	ShippingAddress *BaseAddress                     `json:"shippingAddress,omitEmpty"`
	BillingAddress  *BaseAddress                     `json:"billingAddress,omitEmpty"`
	ShippingMethod  ShippingMethodResourceIdentifier `json:"shippingMethod,omitEmpty"`
	// The custom fields.
	Custom *CustomFieldsDraft `json:"custom,omitEmpty"`
	Locale string             `json:"locale,omitEmpty"`
	// The `TaxMode` `Disabled` can not be set on the My Carts endpoint.
	TaxMode TaxMode `json:"taxMode,omitEmpty"`
	// The cart will be deleted automatically if it hasn't been modified for the specified amount of days and it is in the `Active` CartState.
	// If a ChangeSubscription for carts exists, a `ResourceDeleted` notification will be sent.
	DeleteDaysAfterLastModification int `json:"deleteDaysAfterLastModification,omitEmpty"`
	// Contains addresses for orders with multiple shipping addresses.
	// Each address must contain a key which is unique in this cart.
	ItemShippingAddresses []BaseAddress      `json:"itemShippingAddresses,omitEmpty"`
	Store                 StoreKeyReference  `json:"store,omitEmpty"`
	DiscountCodes         []DiscountCodeInfo `json:"discountCodes,omitEmpty"`
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
		new := MyCartAddDiscountCodeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addItemShippingAddress":
		new := MyCartAddItemShippingAddressAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addLineItem":
		new := MyCartAddLineItemAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addPayment":
		new := MyCartAddPaymentAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "applyDeltaToLineItemShippingDetailsTargets":
		new := MyCartApplyDeltaToLineItemShippingDetailsTargetsAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeLineItemQuantity":
		new := MyCartChangeLineItemQuantityAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeTaxMode":
		new := MyCartChangeTaxModeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "recalculate":
		new := MyCartRecalculateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeDiscountCode":
		new := MyCartRemoveDiscountCodeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeItemShippingAddress":
		new := MyCartRemoveItemShippingAddressAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeLineItem":
		new := MyCartRemoveLineItemAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removePayment":
		new := MyCartRemovePaymentAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setBillingAddress":
		new := MyCartSetBillingAddressAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCountry":
		new := MyCartSetCountryAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomField":
		new := MyCartSetCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomType":
		new := MyCartSetCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomerEmail":
		new := MyCartSetCustomerEmailAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDeleteDaysAfterLastModification":
		new := MyCartSetDeleteDaysAfterLastModificationAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemCustomField":
		new := MyCartSetLineItemCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemCustomType":
		new := MyCartSetLineItemCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemDistributionChannel":
		new := MyCartSetLineItemDistributionChannelAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemShippingDetails":
		new := MyCartSetLineItemShippingDetailsAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLocale":
		new := MyCartSetLocaleAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setShippingAddress":
		new := MyCartSetShippingAddressAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setShippingMethod":
		new := MyCartSetShippingMethodAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "updateItemShippingAddress":
		new := MyCartUpdateItemShippingAddressAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type MyCustomerDraft struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	FirstName   string `json:"firstName,omitEmpty"`
	LastName    string `json:"lastName,omitEmpty"`
	MiddleName  string `json:"middleName,omitEmpty"`
	Title       string `json:"title,omitEmpty"`
	DateOfBirth Date   `json:"dateOfBirth,omitEmpty"`
	CompanyName string `json:"companyName,omitEmpty"`
	VatId       string `json:"vatId,omitEmpty"`
	// Sets the ID of each address to be unique in the addresses list.
	Addresses []BaseAddress `json:"addresses,omitEmpty"`
	// The index of the address in the addresses array.
	// The `defaultShippingAddressId` of the customer will be set to the ID of that address.
	DefaultShippingAddress int `json:"defaultShippingAddress,omitEmpty"`
	// The index of the address in the addresses array.
	// The `defaultBillingAddressId` of the customer will be set to the ID of that address.
	DefaultBillingAddress int `json:"defaultBillingAddress,omitEmpty"`
	// The custom fields.
	Custom *CustomFields             `json:"custom,omitEmpty"`
	Locale string                    `json:"locale,omitEmpty"`
	Stores []StoreResourceIdentifier `json:"stores,omitEmpty"`
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
		new := MyCustomerAddAddressAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addBillingAddressId":
		new := MyCustomerAddBillingAddressIdAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addShippingAddressId":
		new := MyCustomerAddShippingAddressIdAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeAddress":
		new := MyCustomerChangeAddressAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeEmail":
		new := MyCustomerChangeEmailAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeAddress":
		new := MyCustomerRemoveAddressAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeBillingAddressId":
		new := MyCustomerRemoveBillingAddressIdAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeShippingAddressId":
		new := MyCustomerRemoveShippingAddressIdAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCompanyName":
		new := MyCustomerSetCompanyNameAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomField":
		new := MyCustomerSetCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomType":
		new := MyCustomerSetCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDateOfBirth":
		new := MyCustomerSetDateOfBirthAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDefaultBillingAddress":
		new := MyCustomerSetDefaultBillingAddressAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDefaultShippingAddress":
		new := MyCustomerSetDefaultShippingAddressAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setFirstName":
		new := MyCustomerSetFirstNameAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLastName":
		new := MyCustomerSetLastNameAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLocale":
		new := MyCustomerSetLocaleAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setMiddleName":
		new := MyCustomerSetMiddleNameAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setSalutation":
		new := MyCustomerSetSalutationAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setTitle":
		new := MyCustomerSetTitleAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setVatId":
		new := MyCustomerSetVatIdAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type MyLineItemDraft struct {
	ProductId string `json:"productId,omitEmpty"`
	VariantId int    `json:"variantId,omitEmpty"`
	Quantity  int    `json:"quantity"`
	// When the line item was added to the cart. Optional for backwards
	// compatibility reasons only.
	AddedAt time.Time `json:"addedAt,omitEmpty"`
	// By providing supply channel information, you can unique identify
	// inventory entries that should be reserved.
	// The provided channel should have the InventorySupply role.
	SupplyChannel ChannelResourceIdentifier `json:"supplyChannel,omitEmpty"`
	// The channel is used to select a ProductPrice.
	// The provided channel should have the ProductDistribution role.
	DistributionChannel ChannelResourceIdentifier `json:"distributionChannel,omitEmpty"`
	// The custom fields.
	Custom *CustomFieldsDraft `json:"custom,omitEmpty"`
	// Container for line item specific address(es).
	ShippingDetails *ItemShippingDetailsDraft `json:"shippingDetails,omitEmpty"`
	Sku             string                    `json:"sku,omitEmpty"`
}

type MyOrderFromCartDraft struct {
	// The unique ID of the cart from which an order is created.
	Id      string `json:"id"`
	Version int    `json:"version"`
}

type MyPayment struct {
	Id      string `json:"id"`
	Version int    `json:"version"`
	// A reference to the customer this payment belongs to.
	Customer CustomerReference `json:"customer,omitEmpty"`
	// Identifies payments belonging to an anonymous session (the customer has not signed up/in yet).
	AnonymousId string `json:"anonymousId,omitEmpty"`
	// How much money this payment intends to receive from the customer.
	// The value usually matches the cart or order gross total.
	AmountPlanned     TypedMoney        `json:"amountPlanned"`
	PaymentMethodInfo PaymentMethodInfo `json:"paymentMethodInfo"`
	// A list of financial transactions of different TransactionTypes
	// with different TransactionStates.
	Transactions []Transaction `json:"transactions"`
	Custom       *CustomFields `json:"custom,omitEmpty"`
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
	PaymentMethodInfo *PaymentMethodInfo `json:"paymentMethodInfo,omitEmpty"`
	Custom            *CustomFieldsDraft `json:"custom,omitEmpty"`
	// A list of financial transactions of the `Authorization` or `Charge`
	// TransactionTypes.
	Transaction *MyTransactionDraft `json:"transaction,omitEmpty"`
}

type MyPaymentPagedQueryResponse struct {
	Limit   int         `json:"limit"`
	Count   int         `json:"count"`
	Total   int         `json:"total,omitEmpty"`
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
		new := MyPaymentAddTransactionAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeAmountPlanned":
		new := MyPaymentChangeAmountPlannedAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomField":
		new := MyPaymentSetCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setMethodInfoInterface":
		new := MyPaymentSetMethodInfoInterfaceAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setMethodInfoMethod":
		new := MyPaymentSetMethodInfoMethodAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setMethodInfoName":
		new := MyPaymentSetMethodInfoNameAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type MyShoppingListDraft struct {
	Name          LocalizedString             `json:"name"`
	Description   *LocalizedString            `json:"description,omitEmpty"`
	LineItems     []ShoppingListLineItemDraft `json:"lineItems,omitEmpty"`
	TextLineItems []TextLineItemDraft         `json:"textLineItems,omitEmpty"`
	// The custom fields.
	Custom *CustomFieldsDraft `json:"custom,omitEmpty"`
	// The shopping list will be deleted automatically if it hasn't been modified for the specified amount of days.
	DeleteDaysAfterLastModification int                     `json:"deleteDaysAfterLastModification,omitEmpty"`
	Store                           StoreResourceIdentifier `json:"store,omitEmpty"`
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
		new := MyShoppingListAddLineItemAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addTextLineItem":
		new := MyShoppingListAddTextLineItemAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeLineItemQuantity":
		new := MyShoppingListChangeLineItemQuantityAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeLineItemsOrder":
		new := MyShoppingListChangeLineItemsOrderAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeName":
		new := MyShoppingListChangeNameAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeTextLineItemName":
		new := MyShoppingListChangeTextLineItemNameAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeTextLineItemQuantity":
		new := MyShoppingListChangeTextLineItemQuantityAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeTextLineItemsOrder":
		new := MyShoppingListChangeTextLineItemsOrderAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeLineItem":
		new := MyShoppingListRemoveLineItemAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeTextLineItem":
		new := MyShoppingListRemoveTextLineItemAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomField":
		new := MyShoppingListSetCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomType":
		new := MyShoppingListSetCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDeleteDaysAfterLastModification":
		new := MyShoppingListSetDeleteDaysAfterLastModificationAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDescription":
		new := MyShoppingListSetDescriptionAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemCustomField":
		new := MyShoppingListSetLineItemCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemCustomType":
		new := MyShoppingListSetLineItemCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setTextLineItemCustomField":
		new := MyShoppingListSetTextLineItemCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setTextLineItemCustomType":
		new := MyShoppingListSetTextLineItemCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setTextLineItemDescription":
		new := MyShoppingListSetTextLineItemDescriptionAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type MyTransactionDraft struct {
	// The time at which the transaction took place.
	Timestamp time.Time `json:"timestamp,omitEmpty"`
	// The type of this transaction.
	// Only the `Authorization` or `Charge`
	// TransactionTypes are allowed here.
	Type   TransactionType `json:"type"`
	Amount Money           `json:"amount"`
	// The identifier that is used by the interface that managed the transaction (usually the PSP).
	// If a matching interaction was logged in the interfaceInteractions array,
	// the corresponding interaction should be findable with this ID.
	// The `state` is set to the `Initial` TransactionState.
	InteractionId string `json:"interactionId,omitEmpty"`
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
	Custom              *CustomFieldsDraft          `json:"custom,omitEmpty"`
	DistributionChannel ChannelResourceIdentifier   `json:"distributionChannel,omitEmpty"`
	ExternalTaxRate     *ExternalTaxRateDraft       `json:"externalTaxRate,omitEmpty"`
	ProductId           string                      `json:"productId,omitEmpty"`
	VariantId           int                         `json:"variantId,omitEmpty"`
	Sku                 string                      `json:"sku,omitEmpty"`
	Quantity            int                         `json:"quantity,omitEmpty"`
	SupplyChannel       ChannelResourceIdentifier   `json:"supplyChannel,omitEmpty"`
	ExternalPrice       *Money                      `json:"externalPrice,omitEmpty"`
	ExternalTotalPrice  *ExternalLineItemTotalPrice `json:"externalTotalPrice,omitEmpty"`
	ShippingDetails     *ItemShippingDetailsDraft   `json:"shippingDetails,omitEmpty"`
	AddedAt             time.Time                   `json:"addedAt,omitEmpty"`
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
	ExternalPrice      *Money                      `json:"externalPrice,omitEmpty"`
	ExternalTotalPrice *ExternalLineItemTotalPrice `json:"externalTotalPrice,omitEmpty"`
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
	UpdateProductData bool `json:"updateProductData,omitEmpty"`
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
	Quantity                int                         `json:"quantity,omitEmpty"`
	ExternalPrice           *Money                      `json:"externalPrice,omitEmpty"`
	ExternalTotalPrice      *ExternalLineItemTotalPrice `json:"externalTotalPrice,omitEmpty"`
	ShippingDetailsToRemove *ItemShippingDetailsDraft   `json:"shippingDetailsToRemove,omitEmpty"`
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
	Address *BaseAddress `json:"address,omitEmpty"`
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
	Country string `json:"country,omitEmpty"`
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
	Value interface{} `json:"value,omitEmpty"`
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
	Type   TypeResourceIdentifier `json:"type,omitEmpty"`
	Fields *FieldContainer        `json:"fields,omitEmpty"`
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
	Email string `json:"email,omitEmpty"`
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
	DeleteDaysAfterLastModification int `json:"deleteDaysAfterLastModification,omitEmpty"`
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
	Value      interface{} `json:"value,omitEmpty"`
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
	LineItemId string                 `json:"lineItemId"`
	Type       TypeResourceIdentifier `json:"type,omitEmpty"`
	Fields     *FieldContainer        `json:"fields,omitEmpty"`
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
	LineItemId          string                    `json:"lineItemId"`
	DistributionChannel ChannelResourceIdentifier `json:"distributionChannel,omitEmpty"`
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
	ShippingDetails *ItemShippingDetailsDraft `json:"shippingDetails,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartSetLineItemShippingDetailsAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartSetLineItemShippingDetailsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemShippingDetails", Alias: (*Alias)(&obj)})
}

type MyCartSetLocaleAction struct {
	Locale string `json:"locale,omitEmpty"`
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
	Address *BaseAddress `json:"address,omitEmpty"`
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
	ShippingMethod  ShippingMethodResourceIdentifier `json:"shippingMethod,omitEmpty"`
	ExternalTaxRate *ExternalTaxRateDraft            `json:"externalTaxRate,omitEmpty"`
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
	AddressId  string `json:"addressId,omitEmpty"`
	AddressKey string `json:"addressKey,omitEmpty"`
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
	AddressId  string `json:"addressId,omitEmpty"`
	AddressKey string `json:"addressKey,omitEmpty"`
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
	AddressId  string      `json:"addressId,omitEmpty"`
	AddressKey string      `json:"addressKey,omitEmpty"`
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
	AddressId  string `json:"addressId,omitEmpty"`
	AddressKey string `json:"addressKey,omitEmpty"`
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
	AddressId  string `json:"addressId,omitEmpty"`
	AddressKey string `json:"addressKey,omitEmpty"`
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
	AddressId  string `json:"addressId,omitEmpty"`
	AddressKey string `json:"addressKey,omitEmpty"`
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
	CompanyName string `json:"companyName,omitEmpty"`
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
	Value interface{} `json:"value,omitEmpty"`
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
	Type   TypeResourceIdentifier `json:"type,omitEmpty"`
	Fields *FieldContainer        `json:"fields,omitEmpty"`
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
	DateOfBirth Date `json:"dateOfBirth,omitEmpty"`
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
	AddressId  string `json:"addressId,omitEmpty"`
	AddressKey string `json:"addressKey,omitEmpty"`
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
	AddressId  string `json:"addressId,omitEmpty"`
	AddressKey string `json:"addressKey,omitEmpty"`
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
	FirstName string `json:"firstName,omitEmpty"`
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
	LastName string `json:"lastName,omitEmpty"`
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
	Locale string `json:"locale,omitEmpty"`
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
	MiddleName string `json:"middleName,omitEmpty"`
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
	Salutation string `json:"salutation,omitEmpty"`
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
	Title string `json:"title,omitEmpty"`
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
	VatId string `json:"vatId,omitEmpty"`
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
	Value interface{} `json:"value,omitEmpty"`
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
	Method string `json:"method,omitEmpty"`
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
	Name *LocalizedString `json:"name,omitEmpty"`
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
	Sku       string             `json:"sku,omitEmpty"`
	ProductId string             `json:"productId,omitEmpty"`
	VariantId int                `json:"variantId,omitEmpty"`
	Quantity  int                `json:"quantity,omitEmpty"`
	AddedAt   time.Time          `json:"addedAt,omitEmpty"`
	Custom    *CustomFieldsDraft `json:"custom,omitEmpty"`
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
	Description *LocalizedString   `json:"description,omitEmpty"`
	Quantity    int                `json:"quantity,omitEmpty"`
	AddedAt     time.Time          `json:"addedAt,omitEmpty"`
	Custom      *CustomFieldsDraft `json:"custom,omitEmpty"`
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
	Quantity   int    `json:"quantity,omitEmpty"`
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
	Quantity       int    `json:"quantity,omitEmpty"`
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
	Value interface{} `json:"value,omitEmpty"`
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
	Type   TypeResourceIdentifier `json:"type,omitEmpty"`
	Fields *FieldContainer        `json:"fields,omitEmpty"`
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
	DeleteDaysAfterLastModification int `json:"deleteDaysAfterLastModification,omitEmpty"`
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
	Description *LocalizedString `json:"description,omitEmpty"`
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
	Value      interface{} `json:"value,omitEmpty"`
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
	LineItemId string                 `json:"lineItemId"`
	Type       TypeResourceIdentifier `json:"type,omitEmpty"`
	Fields     *FieldContainer        `json:"fields,omitEmpty"`
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
	Value          interface{} `json:"value,omitEmpty"`
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
	TextLineItemId string                 `json:"textLineItemId"`
	Type           TypeResourceIdentifier `json:"type,omitEmpty"`
	Fields         *FieldContainer        `json:"fields,omitEmpty"`
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
	Description    *LocalizedString `json:"description,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListSetTextLineItemDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListSetTextLineItemDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTextLineItemDescription", Alias: (*Alias)(&obj)})
}
