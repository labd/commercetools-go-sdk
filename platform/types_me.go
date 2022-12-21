package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type MyBusinessUnitAssociateDraft struct {
	// Expected version of the BusinessUnit on which the changes should be applied. If the expected version does not match the actual version, a [409 Conflict](/../api/errors#409-conflict) error will be returned.
	Version int `json:"version"`
	// [Customer](ctp:api:type:Customer) to create and assign to the Business Unit.
	Customer MyCustomerDraft `json:"customer"`
}

type MyBusinessUnitDraft interface{}

func mapDiscriminatorMyBusinessUnitDraft(input interface{}) (MyBusinessUnitDraft, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["unitType"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'unitType'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "Company":
		obj := MyCompanyDraft{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "Division":
		obj := MyDivisionDraft{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type MyBusinessUnitUpdate struct {
	// Expected version of the BusinessUnit on which the changes should be applied.
	// If the expected version does not match the actual version, a [409 Conflict](/../api/errors#409-conflict) error will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the BusinessUnit.
	Actions []BusinessUnitUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *MyBusinessUnitUpdate) UnmarshalJSON(data []byte) error {
	type Alias MyBusinessUnitUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorBusinessUnitUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type MyBusinessUnitUpdateAction interface{}

func mapDiscriminatorMyBusinessUnitUpdateAction(input interface{}) (MyBusinessUnitUpdateAction, error) {
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
	case "addAddress":
		obj := MyBusinessUnitAddAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addBillingAddressId":
		obj := MyBusinessUnitAddBillingAddressIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addShippingAddressId":
		obj := MyBusinessUnitAddShippingAddressIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeAddress":
		obj := MyBusinessUnitChangeAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeAssociate":
		obj := MyBusinessUnitChangeAssociateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeName":
		obj := MyBusinessUnitChangeNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeParentUnit":
		obj := MyBusinessUnitChangeParentUnitAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeAddress":
		obj := MyBusinessUnitRemoveAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeAssociate":
		obj := MyBusinessUnitRemoveAssociateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeBillingAddressId":
		obj := MyBusinessUnitRemoveBillingAddressIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeShippingAddressId":
		obj := MyBusinessUnitRemoveShippingAddressIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAddressCustomField":
		obj := MyBusinessUnitSetAddressCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAddressCustomType":
		obj := MyBusinessUnitSetAddressCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setContactEmail":
		obj := MyBusinessUnitSetContactEmailAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomField":
		obj := MyBusinessUnitSetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomType":
		obj := MyBusinessUnitSetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDefaultBillingAddress":
		obj := MyBusinessUnitSetDefaultBillingAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDefaultShippingAddress":
		obj := MyBusinessUnitSetDefaultShippingAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type MyCartDraft struct {
	// A three-digit currency code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Currency      string  `json:"currency"`
	CustomerEmail *string `json:"customerEmail,omitempty"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country *string `json:"country,omitempty"`
	// Default inventory mode is `None`.
	InventoryMode   *InventoryMode                    `json:"inventoryMode,omitempty"`
	LineItems       []MyLineItemDraft                 `json:"lineItems"`
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
	ItemShippingAddresses []BaseAddress `json:"itemShippingAddresses"`
	// The BusinessUnit the cart will belong to.
	BusinessUnit *BusinessUnitKeyReference `json:"businessUnit,omitempty"`
	// [Reference](/../api/types#reference) to a [Store](ctp:api:type:Store) by its key.
	Store *StoreKeyReference `json:"store,omitempty"`
	// The code of existing DiscountCodes.
	DiscountCodes []string `json:"discountCodes"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCartDraft) MarshalJSON() ([]byte, error) {
	type Alias MyCartDraft
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

	if raw["itemShippingAddresses"] == nil {
		delete(raw, "itemShippingAddresses")
	}

	if raw["discountCodes"] == nil {
		delete(raw, "discountCodes")
	}

	return json.Marshal(raw)

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
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorMyCartUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type MyCartUpdateAction interface{}

func mapDiscriminatorMyCartUpdateAction(input interface{}) (MyCartUpdateAction, error) {
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

/**
*	Draft type to represent the top level of a business.
*	Contains the fields and values of the generic [MyBusinessUnitDraft](ctp:api:type:BusinessUnitDraft) that are used specifically for creating a [Company](ctp:api:type:Company).
*
 */
type MyCompanyDraft struct {
	// User-defined unique identifier for the BusinessUnit.
	Key string `json:"key"`
	// Name of the Business Unit.
	Name string `json:"name"`
	// Email address of the Business Unit.
	ContactEmail *string `json:"contactEmail,omitempty"`
	// Custom Fields for the Business Unit.
	Custom *CustomFields `json:"custom,omitempty"`
	// Addresses used by the Business Unit.
	Addresses []BaseAddress `json:"addresses"`
	// Indexes of entries in `addresses` to set as shipping addresses.
	// The `shippingAddressIds` of the [Customer](ctp:api:type:Customer) will be replaced by these addresses.
	ShippingAddresses []int `json:"shippingAddresses"`
	// Index of the entry in `addresses` to set as the default shipping address.
	DefaultShipingAddress *int `json:"defaultShipingAddress,omitempty"`
	// Indexes of entries in `addresses` to set as billing addresses.
	// The `billingAddressIds` of the [Customer](ctp:api:type:Customer) will be replaced by these addresses.
	BillingAddresses []int `json:"billingAddresses"`
	// Index of the entry in `addresses` to set as the default billing address.
	DefaultBillingAddress *int `json:"defaultBillingAddress,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCompanyDraft) MarshalJSON() ([]byte, error) {
	type Alias MyCompanyDraft
	data, err := json.Marshal(struct {
		Action string `json:"unitType"`
		*Alias
	}{Action: "Company", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["addresses"] == nil {
		delete(raw, "addresses")
	}

	if raw["shippingAddresses"] == nil {
		delete(raw, "shippingAddresses")
	}

	if raw["billingAddresses"] == nil {
		delete(raw, "billingAddresses")
	}

	return json.Marshal(raw)

}

type MyCustomerDraft struct {
	// Email address of the Customer that is [unique](/../api/customers-overview#customer-uniqueness) for an entire Project or Store the Customer is assigned to.
	// It is the mandatory unique identifier of a Customer.
	Email string `json:"email"`
	// Password of the Customer.
	Password string `json:"password"`
	// Given name (first name) of the Customer.
	FirstName *string `json:"firstName,omitempty"`
	// Family name (last name) of the Customer.
	LastName *string `json:"lastName,omitempty"`
	// Middle name of the Customer.
	MiddleName *string `json:"middleName,omitempty"`
	// Title of the Customer, for example, 'Dr.'.
	Title *string `json:"title,omitempty"`
	// Salutation of the Customer, for example, 'Mr.' or 'Mrs.'.
	Salutation *string `json:"salutation,omitempty"`
	// Date of birth of the Customer.
	DateOfBirth *Date `json:"dateOfBirth,omitempty"`
	// Company name of the Customer.
	CompanyName *string `json:"companyName,omitempty"`
	// Unique VAT ID of the Customer.
	VatId *string `json:"vatId,omitempty"`
	// Addresses of the Customer.
	Addresses []BaseAddress `json:"addresses"`
	// Index of the address in the `addresses` array to use as the default shipping address.
	// The `defaultShippingAddressId` of the Customer will be set to the `id` of that address.
	DefaultShippingAddress *int `json:"defaultShippingAddress,omitempty"`
	// Index of the address in the `addresses` array to use as the default billing address.
	// The `defaultBillingAddressId` of the Customer will be set to the `id` of that address.
	DefaultBillingAddress *int `json:"defaultBillingAddress,omitempty"`
	// Custom Fields for the Customer.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
	// Preferred language of the Customer. Must be one of the languages supported by the [Project](ctp:api:type:Project).
	Locale *string `json:"locale,omitempty"`
	// Sets the [Stores](ctp:api:type:Store) for the Customer.
	Stores []StoreResourceIdentifier `json:"stores"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCustomerDraft) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerDraft
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

	if raw["addresses"] == nil {
		delete(raw, "addresses")
	}

	if raw["stores"] == nil {
		delete(raw, "stores")
	}

	return json.Marshal(raw)

}

type MyCustomerUpdate struct {
	// Expected version of the Customer on which the changes should be applied. If the expected version does not match the actual version, a [409 Conflict](/../api/errors#409-conflict) error will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the Customer.
	Actions []MyCustomerUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *MyCustomerUpdate) UnmarshalJSON(data []byte) error {
	type Alias MyCustomerUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorMyCustomerUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type MyCustomerUpdateAction interface{}

func mapDiscriminatorMyCustomerUpdateAction(input interface{}) (MyCustomerUpdateAction, error) {
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

/**
*	Draft type to model divisions that are part of the [Company](ctp:api:type:Company) or a higher order [Division](ctp:api:type:Division).
*	Contains the fields and values of the generic [MyBusinessUnitDraft](ctp:api:type:MyBusinessUnitDraft) that are used specifically for creating a Division.
*
 */
type MyDivisionDraft struct {
	// User-defined unique identifier for the BusinessUnit.
	Key string `json:"key"`
	// Name of the Business Unit.
	Name string `json:"name"`
	// Email address of the Business Unit.
	ContactEmail *string `json:"contactEmail,omitempty"`
	// Custom Fields for the Business Unit.
	Custom *CustomFields `json:"custom,omitempty"`
	// Addresses used by the Business Unit.
	Addresses []BaseAddress `json:"addresses"`
	// Indexes of entries in `addresses` to set as shipping addresses.
	// The `shippingAddressIds` of the [Customer](ctp:api:type:Customer) will be replaced by these addresses.
	ShippingAddresses []int `json:"shippingAddresses"`
	// Index of the entry in `addresses` to set as the default shipping address.
	DefaultShipingAddress *int `json:"defaultShipingAddress,omitempty"`
	// Indexes of entries in `addresses` to set as billing addresses.
	// The `billingAddressIds` of the [Customer](ctp:api:type:Customer) will be replaced by these addresses.
	BillingAddresses []int `json:"billingAddresses"`
	// Index of the entry in `addresses` to set as the default billing address.
	DefaultBillingAddress *int `json:"defaultBillingAddress,omitempty"`
	// The parent unit of this Division. Can be a Company or a Division.
	ParentUnit BusinessUnitResourceIdentifier `json:"parentUnit"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyDivisionDraft) MarshalJSON() ([]byte, error) {
	type Alias MyDivisionDraft
	data, err := json.Marshal(struct {
		Action string `json:"unitType"`
		*Alias
	}{Action: "Division", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["addresses"] == nil {
		delete(raw, "addresses")
	}

	if raw["shippingAddresses"] == nil {
		delete(raw, "shippingAddresses")
	}

	if raw["billingAddresses"] == nil {
		delete(raw, "billingAddresses")
	}

	return json.Marshal(raw)

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
	// Unique identifier of the Cart that initiates an Order creation.
	ID      string `json:"id"`
	Version int    `json:"version"`
}

type MyPayment struct {
	// Unique identifier of the Payment.
	ID string `json:"id"`
	// Current version of the Payment.
	Version int `json:"version"`
	// Reference to a [Customer](ctp:api:type:Customer) associated with the Payment. Set automatically with a [password flow token](/../api/authorization#password-flow). Either `customer` or `anonymousId` is present.
	Customer *CustomerReference `json:"customer,omitempty"`
	// [Anonymous session](/../api/authorization#tokens-for-anonymous-sessions) associated with the Payment. Set automatically with a [token for an anonymous session](/../api/authorization#tokens-for-anonymous-sessions). Either `customer` or `anonymousId` is present.
	AnonymousId *string `json:"anonymousId,omitempty"`
	// Money value the Payment intends to receive from the customer.
	// The value typically matches the [Cart](ctp:api:type:Cart) or [Order](ctp:api:type:Order) gross total.
	AmountPlanned CentPrecisionMoney `json:"amountPlanned"`
	// Information regarding the payment interface (for example, a PSP), and the specific payment method used.
	PaymentMethodInfo PaymentMethodInfo `json:"paymentMethodInfo"`
	// Financial transactions of the Payment. Each Transaction has a [TransactionType](ctp:api:type:TransactionType) and a [TransactionState](ctp:api:type:TransactionState).
	Transactions []Transaction `json:"transactions"`
	// Custom Fields defined for the Payment.
	Custom *CustomFields `json:"custom,omitempty"`
}

type MyPaymentDraft struct {
	// Money value the Payment intends to receive from the customer.
	// The value usually matches the [Cart](ctp:api:type:Cart) or [Order](ctp:api:type:Order) gross total.
	AmountPlanned Money `json:"amountPlanned"`
	// Information regarding the payment interface (for example, a PSP), and the specific payment method used.
	PaymentMethodInfo *PaymentMethodInfo `json:"paymentMethodInfo,omitempty"`
	// Custom Fields for the Payment.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
	// Financial transactions of the [TransactionTypes](ctp:api:type:TransactionType) `Authorization` or `Charge`.
	Transaction *MyTransactionDraft `json:"transaction,omitempty"`
}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with `results` containing an array of [MyPayment](ctp:api:type:MyPayment).
*
 */
type MyPaymentPagedQueryResponse struct {
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
	// [MyPayments](ctp:api:type:MyPayment) matching the query.
	Results []MyPayment `json:"results"`
}

type MyPaymentUpdate struct {
	// Expected version of the Payment on which the changes should be applied. If the expected version does not match the actual version, a [409 Conflict](/../api/errors#409-conflict) will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the Payment.
	Actions []MyPaymentUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *MyPaymentUpdate) UnmarshalJSON(data []byte) error {
	type Alias MyPaymentUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorMyPaymentUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type MyPaymentUpdateAction interface{}

func mapDiscriminatorMyPaymentUpdateAction(input interface{}) (MyPaymentUpdateAction, error) {
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
	case "setTransactionCustomField":
		obj := MyPaymentSetTransactionCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type MyQuoteRequestDraft struct {
	// `id` of the Cart from which the Quote Request is created.
	CartId string `json:"cartId"`
	// Current version of the Cart.
	CartVersion int `json:"cartVersion"`
	// Message from the Buyer included in the Quote Request.
	Comment string `json:"comment"`
}

type MyQuoteRequestUpdate struct {
	Version int                          `json:"version"`
	Actions []MyQuoteRequestUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *MyQuoteRequestUpdate) UnmarshalJSON(data []byte) error {
	type Alias MyQuoteRequestUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorMyQuoteRequestUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type MyQuoteRequestUpdateAction interface{}

func mapDiscriminatorMyQuoteRequestUpdateAction(input interface{}) (MyQuoteRequestUpdateAction, error) {
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
	case "cancelQuoteRequest":
		obj := MyQuoteRequestCancelAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	[QuoteStates](ctp:api:type:QuoteState) that can be set using the [Change My Quote State](ctp:api:type:MyQuoteChangeMyQuoteStateAction) update action.
*
 */
type MyQuoteState string

const (
	MyQuoteStateDeclined MyQuoteState = "Declined"
	MyQuoteStateAccepted MyQuoteState = "Accepted"
)

type MyQuoteUpdate struct {
	// Expected version of the [Quote](ctp:api:type:Quote) to which the changes should be applied.
	// If the expected version does not match the actual version, a [409 Conflict](/../api/errors#409-conflict) error will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the [Quote](ctp:api:type:Quote).
	Actions []MyQuoteUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *MyQuoteUpdate) UnmarshalJSON(data []byte) error {
	type Alias MyQuoteUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorMyQuoteUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type MyQuoteUpdateAction interface{}

func mapDiscriminatorMyQuoteUpdateAction(input interface{}) (MyQuoteUpdateAction, error) {
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
	case "changeMyQuoteState":
		obj := MyQuoteChangeMyQuoteStateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	A [MyShoppingListDraft](ctp:api:type:MyShoppingListDraft) is the object submitted as payload to the [Create MyShoppingList request](#create-shoppinglist).
*	The `customer` field of [ShoppingList](ctp:api:type:ShoppingList) is automatically set with
*	a [password flow token](/authorization#password-flow).
*	The `anonymousId` is automatically set with a [token for an anonymous session](/authorization#tokens-for-anonymous-sessions).
*	The `key` and `slug` fields can not be set.
*
 */
type MyShoppingListDraft struct {
	// Name of the [ShoppingList](ctp:api:type:ShoppingList).
	Name LocalizedString `json:"name"`
	// Description of the ShoppingList.
	Description *LocalizedString `json:"description,omitempty"`
	// [Line Items](ctp:api:type:ShoppingListLineItem) (containing Products) to add to the ShoppingList.
	LineItems []ShoppingListLineItemDraft `json:"lineItems"`
	// [Line Items](ctp:api:type:TextLineItem) (containing text values) to add to the ShoppingList.
	TextLineItems []TextLineItemDraft `json:"textLineItems"`
	// Custom Fields defined for the ShoppingList.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
	// Number of days after which the ShoppingList will be automatically deleted if it has not been modified. If not set, the [default value](ctp:api:type:ShoppingListsConfiguration) configured in the [Project](ctp:api:type:Project) is used.
	DeleteDaysAfterLastModification *int `json:"deleteDaysAfterLastModification,omitempty"`
	// Assigns the new ShoppingList to the [Store](ctp:api:type:Store). The Store assignment can not be modified.
	Store *StoreResourceIdentifier `json:"store,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyShoppingListDraft) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListDraft
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

type MyShoppingListUpdate struct {
	// Expected version of the ShoppingList on which the changes should be applied. If the expected version does not match the actual version, a [409 Conflict](/../api/errors#409-conflict) will be returned.
	Version int `json:"version"`
	// List of update actions to be performed on the ShoppingList.
	Actions []MyShoppingListUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *MyShoppingListUpdate) UnmarshalJSON(data []byte) error {
	type Alias MyShoppingListUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorMyShoppingListUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type MyShoppingListUpdateAction interface{}

func mapDiscriminatorMyShoppingListUpdateAction(input interface{}) (MyShoppingListUpdateAction, error) {
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
	// Date and time (UTC) the Transaction took place.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// Type of the Transaction.
	// Only `Authorization` or `Charge` is allowed.
	Type TransactionType `json:"type"`
	// Money value for the Transaction.
	Amount Money `json:"amount"`
	// Identifier used by the payment service that manages the Transaction.
	// Can be used to correlate the Transaction to an interface interaction.
	InteractionId *string `json:"interactionId,omitempty"`
	// Custom Fields of the Transaction.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
}

type ReplicaMyCartDraft struct {
	Reference interface{} `json:"reference"`
}

/**
*	Adding an address to a [Business Unit](ctp:api:type:BusinessUnit) generates a [BusinessUnitAddressAdded](ctp:api:type:BusinessUnitAddressAddedMessage) Message.
*
 */
type MyBusinessUnitAddAddressAction struct {
	// The address to add to `addresses`.
	Address BaseAddress `json:"address"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyBusinessUnitAddAddressAction) MarshalJSON() ([]byte, error) {
	type Alias MyBusinessUnitAddAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addAddress", Alias: (*Alias)(&obj)})
}

/**
*	Adding a billing address to a [Business Unit](ctp:api:type:BusinessUnit) generates a [BusinessUnitBillingAddressAdded](ctp:api:type:BusinessUnitBillingAddressAddedMessage) Message.
*
 */
type MyBusinessUnitAddBillingAddressIdAction struct {
	// ID of the address to add as a billing address. Either `addressId` or `addressKey` is required.
	AddressId *string `json:"addressId,omitempty"`
	// Key of the address to add as a billing address. Either `addressId` or `addressKey` is required.
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyBusinessUnitAddBillingAddressIdAction) MarshalJSON() ([]byte, error) {
	type Alias MyBusinessUnitAddBillingAddressIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addBillingAddressId", Alias: (*Alias)(&obj)})
}

/**
*	Adding a shipping address to a [Business Unit](ctp:api:type:BusinessUnit) generates a [BusinessUnitShippingAddressAdded](ctp:api:type:BusinessUnitShippingAddressAddedMessage) Message.
*
 */
type MyBusinessUnitAddShippingAddressIdAction struct {
	// ID of the address to add as a shipping address. Either `addressId` or `addressKey` is required.
	AddressId *string `json:"addressId,omitempty"`
	// Key of the address to add as a shipping address. Either `addressId` or `addressKey` is required.
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyBusinessUnitAddShippingAddressIdAction) MarshalJSON() ([]byte, error) {
	type Alias MyBusinessUnitAddShippingAddressIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addShippingAddressId", Alias: (*Alias)(&obj)})
}

/**
*	Changing the address on a Business Unit generates the [BusinessUnitAddressChanged](ctp:api:type:BusinessUnitAddressChangedMessage) Message.
*
 */
type MyBusinessUnitChangeAddressAction struct {
	// ID of the address to change. Either `addressId` or `addressKey` is required.
	AddressId *string `json:"addressId,omitempty"`
	// Key of the address to change. Either `addressId` or `addressKey` is required.
	AddressKey *string `json:"addressKey,omitempty"`
	// New address to set.
	Address BaseAddress `json:"address"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyBusinessUnitChangeAddressAction) MarshalJSON() ([]byte, error) {
	type Alias MyBusinessUnitChangeAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAddress", Alias: (*Alias)(&obj)})
}

/**
*	Updating the [Associate](ctp:api:type:Associate) on a [Business Unit](ctp:api:type:BusinessUnit) generates the [BusinessUnitAssociateChanged](ctp:api:type:BusinessUnitAssociateChangedMessage) Message.
*
 */
type MyBusinessUnitChangeAssociateAction struct {
	// The Associate to add.
	Associate AssociateDraft `json:"associate"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyBusinessUnitChangeAssociateAction) MarshalJSON() ([]byte, error) {
	type Alias MyBusinessUnitChangeAssociateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAssociate", Alias: (*Alias)(&obj)})
}

/**
*	Updating the name on a [Business Unit](ctp:api:type:BusinessUnit) generates a [BusinessUnitNameChanged](ctp:api:type:BusinessUnitNameChangedMessage) Message.
*
 */
type MyBusinessUnitChangeNameAction struct {
	// New name to set.
	Name string `json:"name"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyBusinessUnitChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias MyBusinessUnitChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

/**
*	Changing the parent of a [Business Unit](ctp:api:type:BusinessUnit) generates a [BusinessUnitParentUnitChanged](ctp:api:type:BusinessUnitParentUnitChangedMessage) Message. The user must be an Associate with the `Admin` role in the new parent unit.
*
 */
type MyBusinessUnitChangeParentUnitAction struct {
	// New parent unit of the [Business Unit](ctp:api:type:BusinessUnit).
	ParentUnit BusinessUnitResourceIdentifier `json:"parentUnit"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyBusinessUnitChangeParentUnitAction) MarshalJSON() ([]byte, error) {
	type Alias MyBusinessUnitChangeParentUnitAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeParentUnit", Alias: (*Alias)(&obj)})
}

/**
*	Removing the address from a [Business Unit](ctp:api:type:BusinessUnit) generates the [BusinessUnitAddressRemoved](ctp:api:type:BusinessUnitAddressRemovedMessage) Message.
*
 */
type MyBusinessUnitRemoveAddressAction struct {
	// ID of the address to be removed. Either `addressId` or `addressKey` is required.
	AddressId *string `json:"addressId,omitempty"`
	// Key of the address to be removed. Either `addressId` or `addressKey` is required.
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyBusinessUnitRemoveAddressAction) MarshalJSON() ([]byte, error) {
	type Alias MyBusinessUnitRemoveAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeAddress", Alias: (*Alias)(&obj)})
}

/**
*	Removing an [Associate](ctp:api:type:Associate) from a [Business Unit](ctp:api:type:BusinessUnit) generates a [BusinessUnitAssociateRemoved](ctp:api:type:BusinessUnitAssociateRemovedMessage) Message.
*
 */
type MyBusinessUnitRemoveAssociateAction struct {
	// [Associate](ctp:api:type:Associate) to remove.
	Customer CustomerResourceIdentifier `json:"customer"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyBusinessUnitRemoveAssociateAction) MarshalJSON() ([]byte, error) {
	type Alias MyBusinessUnitRemoveAssociateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeAssociate", Alias: (*Alias)(&obj)})
}

/**
*	Removing a billing address from a [Business Unit](ctp:api:type:BusinessUnit) generates a [BusinessUnitBillingAddressRemoved](ctp:api:type:BusinessUnitBillingAddressRemovedMessage) Message.
*
 */
type MyBusinessUnitRemoveBillingAddressIdAction struct {
	// ID of the billing address to be removed. Either `addressId` or `addressKey` is required.
	AddressId *string `json:"addressId,omitempty"`
	// Key of the billing address to be removed. Either `addressId` or `addressKey` is required.
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyBusinessUnitRemoveBillingAddressIdAction) MarshalJSON() ([]byte, error) {
	type Alias MyBusinessUnitRemoveBillingAddressIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeBillingAddressId", Alias: (*Alias)(&obj)})
}

/**
*	Removing a shipping address from a [Business Unit](ctp:api:type:BusinessUnit) generates a [BusinessUnitShippingAddressRemoved](ctp:api:type:BusinessUnitShippingAddressRemovedMessage) Message.
*
 */
type MyBusinessUnitRemoveShippingAddressIdAction struct {
	// ID of the shipping address to be removed. Either `addressId` or `addressKey` is required.
	AddressId *string `json:"addressId,omitempty"`
	// Key of the shipping address to be removed. Either `addressId` or `addressKey` is required.
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyBusinessUnitRemoveShippingAddressIdAction) MarshalJSON() ([]byte, error) {
	type Alias MyBusinessUnitRemoveShippingAddressIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeShippingAddressId", Alias: (*Alias)(&obj)})
}

type MyBusinessUnitSetAddressCustomFieldAction struct {
	// ID of the `address` to be extended.
	AddressId string `json:"addressId"`
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Trying to remove a field that does not exist will fail with an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyBusinessUnitSetAddressCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias MyBusinessUnitSetAddressCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAddressCustomField", Alias: (*Alias)(&obj)})
}

type MyBusinessUnitSetAddressCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the `address` with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the `address`.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the `address`.
	Fields *FieldContainer `json:"fields,omitempty"`
	// ID of the `address` to be extended.
	AddressId string `json:"addressId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyBusinessUnitSetAddressCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias MyBusinessUnitSetAddressCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAddressCustomType", Alias: (*Alias)(&obj)})
}

/**
*	Setting the contact email on a [Business Unit](ctp:api:type:BusinessUnit) generates a [BusinessUnitContactEmailSet](ctp:api:type:BusinessUnitContactEmailSetMessage) Message.
*
 */
type MyBusinessUnitSetContactEmailAction struct {
	// Email to set.
	// If `contactEmail` is absent or `null`, the existing contact email, if any, will be removed.
	ContactEmail *string `json:"contactEmail,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyBusinessUnitSetContactEmailAction) MarshalJSON() ([]byte, error) {
	type Alias MyBusinessUnitSetContactEmailAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setContactEmail", Alias: (*Alias)(&obj)})
}

type MyBusinessUnitSetCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Trying to remove a field that does not exist will fail with an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyBusinessUnitSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias MyBusinessUnitSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type MyBusinessUnitSetCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the BusinessUnit with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the BusinessUnit.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) for the BusinessUnit.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyBusinessUnitSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias MyBusinessUnitSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

/**
*	Setting the default billing address on a [Business Unit](ctp:api:type:BusinessUnit) generates the [BusinessUnitDefaultBillingAddressSet](ctp:api:type:BusinessUnitDefaultBillingAddressSetMessage) Message.
*
 */
type MyBusinessUnitSetDefaultBillingAddressAction struct {
	// ID of the address to add as a billing address. Either `addressId` or `addressKey` is required.
	AddressId *string `json:"addressId,omitempty"`
	// Key of the address to add as a billing address. Either `addressId` or `addressKey` is required.
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyBusinessUnitSetDefaultBillingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias MyBusinessUnitSetDefaultBillingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDefaultBillingAddress", Alias: (*Alias)(&obj)})
}

/**
*	Setting the default shipping address on a [Business Unit](ctp:api:type:BusinessUnit) generates a [BusinessUnitDefaultShippingAddressSet](ctp:api:type:BusinessUnitDefaultShippingAddressSetMessage) Message.
*
 */
type MyBusinessUnitSetDefaultShippingAddressAction struct {
	// ID of the address to add as a shipping address. Either `addressId` or `addressKey` is required.
	AddressId *string `json:"addressId,omitempty"`
	// Key of the address to add as a shipping address. Either `addressId` or `addressKey` is required.
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyBusinessUnitSetDefaultShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias MyBusinessUnitSetDefaultShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDefaultShippingAddress", Alias: (*Alias)(&obj)})
}

type MyCartAddDiscountCodeAction struct {
	Code string `json:"code"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCartAddItemShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartAddItemShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addItemShippingAddress", Alias: (*Alias)(&obj)})
}

type MyCartAddLineItemAction struct {
	// The representation used when creating or updating a [customizable data type](/../api/projects/types#list-of-customizable-data-types) with Custom Fields.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [Channel](ctp:api:type:Channel).
	DistributionChannel *ChannelResourceIdentifier `json:"distributionChannel,omitempty"`
	ExternalTaxRate     *ExternalTaxRateDraft      `json:"externalTaxRate,omitempty"`
	ProductId           *string                    `json:"productId,omitempty"`
	VariantId           *int                       `json:"variantId,omitempty"`
	Sku                 *string                    `json:"sku,omitempty"`
	Quantity            *int                       `json:"quantity,omitempty"`
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [Channel](ctp:api:type:Channel).
	SupplyChannel *ChannelResourceIdentifier `json:"supplyChannel,omitempty"`
	// Draft type that stores amounts in cent precision for the specified currency.
	//
	// For storing money values in fractions of the minor unit in a currency, use [HighPrecisionMoneyDraft](ctp:api:type:HighPrecisionMoneyDraft) instead.
	ExternalPrice      *Money                      `json:"externalPrice,omitempty"`
	ExternalTotalPrice *ExternalLineItemTotalPrice `json:"externalTotalPrice,omitempty"`
	ShippingDetails    *ItemShippingDetailsDraft   `json:"shippingDetails,omitempty"`
	AddedAt            *time.Time                  `json:"addedAt,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCartAddLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartAddLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addLineItem", Alias: (*Alias)(&obj)})
}

type MyCartAddPaymentAction struct {
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) of a [Payment](ctp:api:type:Payment).
	Payment PaymentResourceIdentifier `json:"payment"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCartApplyDeltaToLineItemShippingDetailsTargetsAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartApplyDeltaToLineItemShippingDetailsTargetsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "applyDeltaToLineItemShippingDetailsTargets", Alias: (*Alias)(&obj)})
}

type MyCartChangeLineItemQuantityAction struct {
	LineItemId string `json:"lineItemId"`
	Quantity   int    `json:"quantity"`
	// Draft type that stores amounts in cent precision for the specified currency.
	//
	// For storing money values in fractions of the minor unit in a currency, use [HighPrecisionMoneyDraft](ctp:api:type:HighPrecisionMoneyDraft) instead.
	ExternalPrice      *Money                      `json:"externalPrice,omitempty"`
	ExternalTotalPrice *ExternalLineItemTotalPrice `json:"externalTotalPrice,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCartRecalculateAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartRecalculateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "recalculate", Alias: (*Alias)(&obj)})
}

type MyCartRemoveDiscountCodeAction struct {
	// [Reference](ctp:api:type:Reference) to a [DiscountCode](ctp:api:type:DiscountCode).
	DiscountCode DiscountCodeReference `json:"discountCode"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCartRemoveItemShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartRemoveItemShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeItemShippingAddress", Alias: (*Alias)(&obj)})
}

type MyCartRemoveLineItemAction struct {
	LineItemId string `json:"lineItemId"`
	Quantity   *int   `json:"quantity,omitempty"`
	// Draft type that stores amounts in cent precision for the specified currency.
	//
	// For storing money values in fractions of the minor unit in a currency, use [HighPrecisionMoneyDraft](ctp:api:type:HighPrecisionMoneyDraft) instead.
	ExternalPrice           *Money                      `json:"externalPrice,omitempty"`
	ExternalTotalPrice      *ExternalLineItemTotalPrice `json:"externalTotalPrice,omitempty"`
	ShippingDetailsToRemove *ItemShippingDetailsDraft   `json:"shippingDetailsToRemove,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCartRemoveLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartRemoveLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeLineItem", Alias: (*Alias)(&obj)})
}

type MyCartRemovePaymentAction struct {
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) of a [Payment](ctp:api:type:Payment).
	Payment PaymentResourceIdentifier `json:"payment"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCartSetBillingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartSetBillingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setBillingAddress", Alias: (*Alias)(&obj)})
}

type MyCartSetCountryAction struct {
	// Two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country *string `json:"country,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCartSetCountryAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartSetCountryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCountry", Alias: (*Alias)(&obj)})
}

type MyCartSetCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCartSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type MyCartSetCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the MyCart with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the MyCart.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the MyCart.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCartSetDeleteDaysAfterLastModificationAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartSetDeleteDaysAfterLastModificationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeleteDaysAfterLastModification", Alias: (*Alias)(&obj)})
}

type MyCartSetLineItemCustomFieldAction struct {
	LineItemId string `json:"lineItemId"`
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCartSetLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartSetLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomField", Alias: (*Alias)(&obj)})
}

type MyCartSetLineItemCustomTypeAction struct {
	LineItemId string `json:"lineItemId"`
	// Defines the [Type](ctp:api:type:Type) that extends the LineItem with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the LineItem.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the LineItem.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCartSetLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartSetLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomType", Alias: (*Alias)(&obj)})
}

type MyCartSetLineItemDistributionChannelAction struct {
	LineItemId string `json:"lineItemId"`
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [Channel](ctp:api:type:Channel).
	DistributionChannel *ChannelResourceIdentifier `json:"distributionChannel,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCartSetLineItemShippingDetailsAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartSetLineItemShippingDetailsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemShippingDetails", Alias: (*Alias)(&obj)})
}

type MyCartSetLineItemSupplyChannelAction struct {
	LineItemId string `json:"lineItemId"`
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [Channel](ctp:api:type:Channel).
	SupplyChannel *ChannelResourceIdentifier `json:"supplyChannel,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCartSetShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartSetShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingAddress", Alias: (*Alias)(&obj)})
}

type MyCartSetShippingMethodAction struct {
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [ShippingMethod](ctp:api:type:ShippingMethod).
	ShippingMethod  *ShippingMethodResourceIdentifier `json:"shippingMethod,omitempty"`
	ExternalTaxRate *ExternalTaxRateDraft             `json:"externalTaxRate,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCartUpdateItemShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartUpdateItemShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "updateItemShippingAddress", Alias: (*Alias)(&obj)})
}

/**
*	Adding an address to the Customer produces the [CustomerAddressAdded](ctp:api:type:CustomerAddressAddedMessage) Message.
*
 */
type MyCustomerAddAddressAction struct {
	// Value to append to the `addresses` array.
	Address BaseAddress `json:"address"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCustomerAddAddressAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerAddAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addAddress", Alias: (*Alias)(&obj)})
}

/**
*	Adds an address from the `addresses` array to `billingAddressIds`. Either `addressId` or `addressKey` is required.
*
 */
type MyCustomerAddBillingAddressIdAction struct {
	// `id` of the [Address](ctp:api:type:Address) to become a billing address.
	AddressId *string `json:"addressId,omitempty"`
	// `key` of the [Address](ctp:api:type:Address) to become a billing address.
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCustomerAddBillingAddressIdAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerAddBillingAddressIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addBillingAddressId", Alias: (*Alias)(&obj)})
}

/**
*	Adds an address from the `addresses` array to `shippingAddressIds`. Either `addressId` or `addressKey` is required.
*
 */
type MyCustomerAddShippingAddressIdAction struct {
	// `id` of the [Address](ctp:api:type:Address) to become a shipping address.
	AddressId *string `json:"addressId,omitempty"`
	// `key` of the [Address](ctp:api:type:Address) to become a shipping address.
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCustomerAddShippingAddressIdAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerAddShippingAddressIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addShippingAddressId", Alias: (*Alias)(&obj)})
}

/**
*	Changing an address of the Customer produces the [CustomerAddressChanged](ctp:api:type:CustomerAddressChangedMessage) Message.
*
*	Either `addressId` or `addressKey` is required.
*
 */
type MyCustomerChangeAddressAction struct {
	// `id` of the [Address](ctp:api:type:Address) to change.
	AddressId *string `json:"addressId,omitempty"`
	// `key` of the [Address](ctp:api:type:Address) to change.
	AddressKey *string `json:"addressKey,omitempty"`
	// Value to set.
	Address BaseAddress `json:"address"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCustomerChangeAddressAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerChangeAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAddress", Alias: (*Alias)(&obj)})
}

/**
*	Changing the email of the Customer produces the [CustomerEmailChanged](ctp:api:type:CustomerEmailChangedMessage) Message.
*
 */
type MyCustomerChangeEmailAction struct {
	// New value to set.
	Email string `json:"email"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCustomerChangeEmailAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerChangeEmailAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeEmail", Alias: (*Alias)(&obj)})
}

/**
*	Removing an address of the Customer produces the [CustomerAddressRemoved](ctp:api:type:CustomerAddressRemovedMessage) Message.
*
*	Either `addressId` or `addressKey` is required.
*
 */
type MyCustomerRemoveAddressAction struct {
	// `id` of the [Address](ctp:api:type:Address) to remove.
	AddressId *string `json:"addressId,omitempty"`
	// `key` of the [Address](ctp:api:type:Address) to remove.
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCustomerRemoveAddressAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerRemoveAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeAddress", Alias: (*Alias)(&obj)})
}

/**
*	Removes an existing billing address from `billingAddressesIds`.
*	If the billing address is the default billing address, the `defaultBillingAddressId` is unset. Either `addressId` or `addressKey` is required.
*
 */
type MyCustomerRemoveBillingAddressIdAction struct {
	// `id` of the [Address](ctp:api:type:Address) to remove from `billingAddressesIds`.
	AddressId *string `json:"addressId,omitempty"`
	// `key` of the [Address](ctp:api:type:Address) to remove from `billingAddressesIds`.
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCustomerRemoveBillingAddressIdAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerRemoveBillingAddressIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeBillingAddressId", Alias: (*Alias)(&obj)})
}

/**
*	Removes an existing shipping address from `shippingAddressesIds`.
*	If the shipping address is the default shipping address, the `defaultShippingAddressId` is unset. Either `addressId` or `addressKey` is required.
*
 */
type MyCustomerRemoveShippingAddressIdAction struct {
	// `id` of the [Address](ctp:api:type:Address) to remove from `shippingAddressesIds`.
	AddressId *string `json:"addressId,omitempty"`
	// `key` of the [Address](ctp:api:type:Address) to remove from `shippingAddressesIds`.
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCustomerRemoveShippingAddressIdAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerRemoveShippingAddressIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeShippingAddressId", Alias: (*Alias)(&obj)})
}

/**
*	Setting the `companyName` field on the Customer produces the [CustomerCompanyNameSet](ctp:api:type:CustomerCompanyNameSetMessage) Message.
*
 */
type MyCustomerSetCompanyNameAction struct {
	// Value to set.
	// If empty, any existing value is removed.
	CompanyName *string `json:"companyName,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCustomerSetCompanyNameAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetCompanyNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCompanyName", Alias: (*Alias)(&obj)})
}

type MyCustomerSetCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// If `value` is provided, it is set for the field defined by `name`.
	// Trying to remove a field that does not exist will fail with an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCustomerSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type MyCustomerSetCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the MyCustomer with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the MyCustomer.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the MyCustomer.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCustomerSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

/**
*	Setting the date of birth of the Customer produces the [CustomerDateOfBirthSet](ctp:api:type:CustomerDateOfBirthSetMessage) Message.
*
 */
type MyCustomerSetDateOfBirthAction struct {
	// Value to set.
	// If empty, any existing value is removed.
	DateOfBirth *Date `json:"dateOfBirth,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCustomerSetDateOfBirthAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetDateOfBirthAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDateOfBirth", Alias: (*Alias)(&obj)})
}

/**
*	Sets the default billing address from `addresses`.
*	If the address is not currently a billing address, it is added to `billingAddressIds`. Either `addressId` or `addressKey` is required.
*
 */
type MyCustomerSetDefaultBillingAddressAction struct {
	// `id` of the [Address](ctp:api:type:Address) to become the default billing address.
	AddressId *string `json:"addressId,omitempty"`
	// `key` of the [Address](ctp:api:type:Address) to become the default billing address.
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCustomerSetDefaultBillingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetDefaultBillingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDefaultBillingAddress", Alias: (*Alias)(&obj)})
}

/**
*	Sets the default shipping address from `addresses`.
*	If the address is not currently a shipping address, it is added to `shippingAddressIds`. Either `addressId` or `addressKey` is required.
*
 */
type MyCustomerSetDefaultShippingAddressAction struct {
	// `id` of the [Address](ctp:api:type:Address) to become the default shipping address.
	AddressId *string `json:"addressId,omitempty"`
	// `key` of the [Address](ctp:api:type:Address) to become the default shipping address.
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCustomerSetDefaultShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetDefaultShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDefaultShippingAddress", Alias: (*Alias)(&obj)})
}

/**
*	Setting the first name of the Customer produces the [CustomerFirstNameSetMessage](ctp:api:type:CustomerFirstNameSetMessage).
*
 */
type MyCustomerSetFirstNameAction struct {
	// Value to set.
	// If empty, any existing value is removed.
	FirstName *string `json:"firstName,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCustomerSetFirstNameAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetFirstNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setFirstName", Alias: (*Alias)(&obj)})
}

/**
*	Setting the last name of the Customer produces the [CustomerLastNameSetMessage](ctp:api:type:CustomerLastNameSetMessage).
*
 */
type MyCustomerSetLastNameAction struct {
	// Value to set.
	// If empty, any existing value is removed.
	LastName *string `json:"lastName,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCustomerSetLastNameAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetLastNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLastName", Alias: (*Alias)(&obj)})
}

type MyCustomerSetLocaleAction struct {
	// Value to set.
	// Must be one of the languages supported by the [Project](ctp:api:type:Project).
	Locale *string `json:"locale,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCustomerSetLocaleAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetLocaleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLocale", Alias: (*Alias)(&obj)})
}

type MyCustomerSetMiddleNameAction struct {
	// Value to set.
	// If empty, any existing value is removed.
	MiddleName *string `json:"middleName,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCustomerSetMiddleNameAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetMiddleNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMiddleName", Alias: (*Alias)(&obj)})
}

type MyCustomerSetSalutationAction struct {
	// Value to set.
	// If empty, any existing value is removed.
	Salutation *string `json:"salutation,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCustomerSetSalutationAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetSalutationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSalutation", Alias: (*Alias)(&obj)})
}

/**
*	Setting the title of the Customer produces the [CustomerTitleSetMessage](ctp:api:type:CustomerTitleSetMessage).
*
 */
type MyCustomerSetTitleAction struct {
	// Value to set.
	// If empty, any existing value is removed.
	Title *string `json:"title,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCustomerSetTitleAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetTitleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTitle", Alias: (*Alias)(&obj)})
}

type MyCustomerSetVatIdAction struct {
	// Value to set.
	// If empty, any existing value is removed.
	VatId *string `json:"vatId,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyCustomerSetVatIdAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetVatIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setVatId", Alias: (*Alias)(&obj)})
}

/**
*	Adding a Transaction to a Payment generates the [PaymentTransactionAdded](ctp:api:type:PaymentTransactionAddedMessage) Message.
*	Once a Transaction is added to the Payment, it can no longer be updated or deleted using the My Payments API.
*
 */
type MyPaymentAddTransactionAction struct {
	// Transaction to add to the Payment.
	Transaction TransactionDraft `json:"transaction"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyPaymentAddTransactionAction) MarshalJSON() ([]byte, error) {
	type Alias MyPaymentAddTransactionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addTransaction", Alias: (*Alias)(&obj)})
}

/**
*	Can be used to update the Payment if a customer changes the [Cart](ctp:api:type:Cart), or adds or removes a [CartDiscount](ctp:api:type:CartDiscount) during checkout.
*
 */
type MyPaymentChangeAmountPlannedAction struct {
	// New value to set.
	Amount Money `json:"amount"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyPaymentChangeAmountPlannedAction) MarshalJSON() ([]byte, error) {
	type Alias MyPaymentChangeAmountPlannedAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAmountPlanned", Alias: (*Alias)(&obj)})
}

type MyPaymentSetCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyPaymentSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias MyPaymentSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type MyPaymentSetMethodInfoInterfaceAction struct {
	// Value to set.
	// Once set, the `paymentInterface` of the `paymentMethodInfo` cannot be changed.
	Interface string `json:"interface"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyPaymentSetMethodInfoInterfaceAction) MarshalJSON() ([]byte, error) {
	type Alias MyPaymentSetMethodInfoInterfaceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMethodInfoInterface", Alias: (*Alias)(&obj)})
}

type MyPaymentSetMethodInfoMethodAction struct {
	// Value to set.
	// If empty, any existing value will be removed.
	Method *string `json:"method,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyPaymentSetMethodInfoMethodAction) MarshalJSON() ([]byte, error) {
	type Alias MyPaymentSetMethodInfoMethodAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMethodInfoMethod", Alias: (*Alias)(&obj)})
}

type MyPaymentSetMethodInfoNameAction struct {
	// Value to set.
	// If empty, any existing value will be removed.
	Name *LocalizedString `json:"name,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyPaymentSetMethodInfoNameAction) MarshalJSON() ([]byte, error) {
	type Alias MyPaymentSetMethodInfoNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMethodInfoName", Alias: (*Alias)(&obj)})
}

type MyPaymentSetTransactionCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyPaymentSetTransactionCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias MyPaymentSetTransactionCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTransactionCustomField", Alias: (*Alias)(&obj)})
}

type MyQuoteChangeMyQuoteStateAction struct {
	// New state to be set for the Quote.
	QuoteState MyQuoteState `json:"quoteState"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyQuoteChangeMyQuoteStateAction) MarshalJSON() ([]byte, error) {
	type Alias MyQuoteChangeMyQuoteStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeMyQuoteState", Alias: (*Alias)(&obj)})
}

/**
*	Transitions the `quoteRequestState` of the Quote Request to `Cancelled`. Can only be used when the Quote Request is in state `Submitted`.
*
 */
type MyQuoteRequestCancelAction struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyQuoteRequestCancelAction) MarshalJSON() ([]byte, error) {
	type Alias MyQuoteRequestCancelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "cancelQuoteRequest", Alias: (*Alias)(&obj)})
}

type MyShoppingListAddLineItemAction struct {
	// `sku` of the [ProductVariant](ctp:api:type:ProductVariant).
	Sku *string `json:"sku,omitempty"`
	// Unique identifier of a [Product](ctp:api:type:Product).
	ProductId *string `json:"productId,omitempty"`
	// `id` of the [ProductVariant](ctp:api:type:ProductVariant). If not set, the ShoppingListLineItem refers to the Master Variant.
	VariantId *int `json:"variantId,omitempty"`
	// Number of Products in the [ShoppingListLineItem](ctp:api:type:ShoppingListLineItem).
	Quantity *int `json:"quantity,omitempty"`
	// Date and time the TextLineItem is added to the [ShoppingList](ctp:api:type:ShoppingList). If not set, the current date and time (UTC) is used.
	AddedAt *time.Time `json:"addedAt,omitempty"`
	// Custom Fields defined for the ShoppingListLineItem.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyShoppingListAddLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListAddLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addLineItem", Alias: (*Alias)(&obj)})
}

type MyShoppingListAddTextLineItemAction struct {
	// Name of the [TextLineItem](ctp:api:type:TextLineItem).
	Name LocalizedString `json:"name"`
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
func (obj MyShoppingListAddTextLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListAddTextLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addTextLineItem", Alias: (*Alias)(&obj)})
}

type MyShoppingListChangeLineItemQuantityAction struct {
	// The `id` of the [ShoppingListLineItem](ctp:api:type:ShoppingListLineItem) to update.
	LineItemId string `json:"lineItemId"`
	// New value to set. If `0`, the ShoppingListLineItem is removed from the ShoppingList.
	Quantity int `json:"quantity"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyShoppingListChangeLineItemQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListChangeLineItemQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLineItemQuantity", Alias: (*Alias)(&obj)})
}

type MyShoppingListChangeLineItemsOrderAction struct {
	// All existing [ShoppingListLineItem](ctp:api:type:ShoppingListLineItem) `id`s of the [ShoppingList](ctp:api:type:ShoppingList) in the desired new order.
	LineItemOrder []string `json:"lineItemOrder"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyShoppingListChangeLineItemsOrderAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListChangeLineItemsOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLineItemsOrder", Alias: (*Alias)(&obj)})
}

type MyShoppingListChangeNameAction struct {
	// New value to set. Must not be empty.
	Name LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyShoppingListChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type MyShoppingListChangeTextLineItemNameAction struct {
	// The `id` of the [TextLineItem](ctp:api:type:TextLineItem) to update.
	TextLineItemId string `json:"textLineItemId"`
	// New value to set. Must not be empty.
	Name LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyShoppingListChangeTextLineItemNameAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListChangeTextLineItemNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTextLineItemName", Alias: (*Alias)(&obj)})
}

type MyShoppingListChangeTextLineItemQuantityAction struct {
	// The `id` of the [TextLineItem](ctp:api:type:TextLineItem) to update.
	TextLineItemId string `json:"textLineItemId"`
	// New value to set. If `0`, the TextLineItem is removed from the ShoppingList.
	Quantity int `json:"quantity"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyShoppingListChangeTextLineItemQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListChangeTextLineItemQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTextLineItemQuantity", Alias: (*Alias)(&obj)})
}

type MyShoppingListChangeTextLineItemsOrderAction struct {
	// All existing [TextLineItem](ctp:api:type:TextLineItem) `id`s in the desired new order.
	TextLineItemOrder []string `json:"textLineItemOrder"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyShoppingListChangeTextLineItemsOrderAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListChangeTextLineItemsOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTextLineItemsOrder", Alias: (*Alias)(&obj)})
}

type MyShoppingListRemoveLineItemAction struct {
	// The `id` of the [ShoppingListLineItem](ctp:api:type:ShoppingListLineItem) to update.
	LineItemId string `json:"lineItemId"`
	// Amount to remove from the `quantity` of the ShoppingListLineItem. If not set, the ShoppingListLineItem is removed from the ShoppingList. If this value matches or exceeds the current `quantity` of the ShoppingListLineItem, the ShoppingListLineItem is removed from the ShoppingList.
	Quantity *int `json:"quantity,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyShoppingListRemoveLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListRemoveLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeLineItem", Alias: (*Alias)(&obj)})
}

type MyShoppingListRemoveTextLineItemAction struct {
	// The `id` of the [TextLineItem](ctp:api:type:TextLineItem) to update.
	TextLineItemId string `json:"textLineItemId"`
	// Amount to remove from the `quantity` of the TextLineItem. If not set, the TextLineItem is removed from the ShoppingList. If this value matches or exceeds the current `quantity` of the TextLineItem, the TextLineItem is removed from the ShoppingList.
	Quantity *int `json:"quantity,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyShoppingListRemoveTextLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListRemoveTextLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeTextLineItem", Alias: (*Alias)(&obj)})
}

type MyShoppingListSetCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyShoppingListSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type MyShoppingListSetCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the MyShoppingList with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the MyShoppingList.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the MyShoppingList.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyShoppingListSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type MyShoppingListSetDeleteDaysAfterLastModificationAction struct {
	// Value to set. If empty, any existing value will be removed.
	DeleteDaysAfterLastModification *int `json:"deleteDaysAfterLastModification,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyShoppingListSetDeleteDaysAfterLastModificationAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListSetDeleteDaysAfterLastModificationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeleteDaysAfterLastModification", Alias: (*Alias)(&obj)})
}

type MyShoppingListSetDescriptionAction struct {
	// Value to set. If empty, any existing value will be removed.
	Description *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyShoppingListSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type MyShoppingListSetLineItemCustomFieldAction struct {
	// Unique identifier of an existing [ShoppingListLineItem](ctp:api:type:ShoppingListLineItem) in the [ShoppingList](ctp:api:type:ShoppingList).
	LineItemId string `json:"lineItemId"`
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyShoppingListSetLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListSetLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomField", Alias: (*Alias)(&obj)})
}

type MyShoppingListSetLineItemCustomTypeAction struct {
	// Unique identifier of an existing [ShoppingListLineItem](ctp:api:type:ShoppingListLineItem) in the [ShoppingList](ctp:api:type:ShoppingList).
	LineItemId string `json:"lineItemId"`
	// Defines the [Type](ctp:api:type:Type) that extends the ShoppingListLineItem with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the ShoppingListLineItem.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the ShoppingListLineItem.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyShoppingListSetLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListSetLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomType", Alias: (*Alias)(&obj)})
}

type MyShoppingListSetTextLineItemCustomFieldAction struct {
	// The `id` of the [TextLineItem](ctp:api:type:TextLineItem) to update.
	TextLineItemId string `json:"textLineItemId"`
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyShoppingListSetTextLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListSetTextLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTextLineItemCustomField", Alias: (*Alias)(&obj)})
}

type MyShoppingListSetTextLineItemCustomTypeAction struct {
	// The `id` of the [TextLineItem](ctp:api:type:TextLineItem) to update.
	TextLineItemId string `json:"textLineItemId"`
	// Defines the [Type](ctp:api:type:Type) that extends the TextLineItem with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the TextLineItem.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the TextLineItem.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyShoppingListSetTextLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListSetTextLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTextLineItemCustomType", Alias: (*Alias)(&obj)})
}

type MyShoppingListSetTextLineItemDescriptionAction struct {
	// The `id` of the [TextLineItem](ctp:api:type:TextLineItem) to update.
	TextLineItemId string `json:"textLineItemId"`
	// Value to set. If empty, any existing value will be removed.
	Description *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MyShoppingListSetTextLineItemDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListSetTextLineItemDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTextLineItemDescription", Alias: (*Alias)(&obj)})
}
