package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type AnonymousCartSignInMode string

const (
	AnonymousCartSignInModeMergeWithExistingCustomerCart AnonymousCartSignInMode = "MergeWithExistingCustomerCart"
	AnonymousCartSignInModeUseAsNewActiveCustomerCart    AnonymousCartSignInMode = "UseAsNewActiveCustomerCart"
)

type AuthenticationMode string

const (
	AuthenticationModePassword     AuthenticationMode = "Password"
	AuthenticationModeExternalAuth AuthenticationMode = "ExternalAuth"
)

/**
*	If `stores` is not empty, the Customer is specific to those Stores.
*
 */
type Customer struct {
	// Unique identifier of the Customer.
	ID string `json:"id"`
	// Current version of the Customer.
	Version int `json:"version"`
	// Date and time (UTC) the Customer was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the Customer was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// User-defined unique identifier of the Customer.
	Key *string `json:"key,omitempty"`
	// User-defined unique identifier of the Customer.
	//
	// Can be used to refer to a Customer in a human-readable way (in emails, invoices, and other correspondence).
	CustomerNumber *string `json:"customerNumber,omitempty"`
	// Optional identifier for use in external systems like Customer Relationship Management (CRM) or Enterprise Resource Planning (ERP).
	ExternalId *string `json:"externalId,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/../api/general-concepts#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/../api/general-concepts#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Email address of the Customer that is [unique](/../api/customers-overview#customer-uniqueness) for an entire Project or to a Store the Customer is assigned to.
	// It is the mandatory unique identifier of a Customer.
	Email string `json:"email"`
	// Present only when `authenticationMode` is set to `Password`.
	Password *string `json:"password,omitempty"`
	// Given name (first name) of the Customer.
	FirstName *string `json:"firstName,omitempty"`
	// Family name (last name) of the Customer.
	LastName *string `json:"lastName,omitempty"`
	// Middle name of the Customer.
	MiddleName *string `json:"middleName,omitempty"`
	// Title of the Customer, for example, 'Dr.'.
	Title *string `json:"title,omitempty"`
	// Date of birth of the Customer.
	DateOfBirth *Date `json:"dateOfBirth,omitempty"`
	// Company name of the Customer.
	CompanyName *string `json:"companyName,omitempty"`
	// Individual VAT ID of the Customer.
	VatId *string `json:"vatId,omitempty"`
	// Addresses used by the Customer.
	Addresses []Address `json:"addresses"`
	// ID of the address in `addresses` used as the default shipping address.
	DefaultShippingAddressId *string `json:"defaultShippingAddressId,omitempty"`
	// IDs of addresses in `addresses` used as shipping addresses.
	ShippingAddressIds []string `json:"shippingAddressIds"`
	// ID of the address in `addresses` used as the default billing address.
	DefaultBillingAddressId *string `json:"defaultBillingAddressId,omitempty"`
	// IDs of addresses in `addresses` used as billing addresses.
	BillingAddressIds []string `json:"billingAddressIds"`
	// Indicates whether the email address of the Customer is [verified](#email-verification-of-customer).
	IsEmailVerified bool `json:"isEmailVerified"`
	// [CustomerGroup](ctp:api:type:CustomerGroup) to which the Customer belongs.
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
	// Custom Fields for the Customer.
	Custom *CustomFields `json:"custom,omitempty"`
	// Preferred language of the Customer.
	Locale *string `json:"locale,omitempty"`
	// Salutation of the Customer, for example, 'Mr.' or 'Mrs.'.
	Salutation *string `json:"salutation,omitempty"`
	// [Stores](ctp:api:type:Store) to which the Customer is assigned to.
	//
	// - If no Stores are specified, the Customer is a global customer, and can log in using the [Password Flow for global Customers](/../api/authorization#password-flow-for-global-customers).
	// - If any Stores are specified, the Customer can only log in using the [Password Flow for Customers in a Store](/../api/authorization#password-flow-for-customers-in-a-store) for those specific Stores.
	Stores []StoreKeyReference `json:"stores"`
	// Indicates whether the `password` is required for the Customer.
	AuthenticationMode AuthenticationMode `json:"authenticationMode"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj Customer) MarshalJSON() ([]byte, error) {
	type Alias Customer
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

	if raw["shippingAddressIds"] == nil {
		delete(raw, "shippingAddressIds")
	}

	if raw["billingAddressIds"] == nil {
		delete(raw, "billingAddressIds")
	}

	if raw["stores"] == nil {
		delete(raw, "stores")
	}

	return json.Marshal(raw)

}

type CustomerChangePassword struct {
	// Unique identifier of the Customer.
	ID string `json:"id"`
	// Expected version of the Customer on which the changes should be applied.
	Version int `json:"version"`
	// Current password of the Customer.
	//
	// If the current password does not match, an [InvalidCurrentPassword](ctp:api:type:InvalidCurrentPasswordError) error is returned.
	CurrentPassword string `json:"currentPassword"`
	// New password to be set.
	NewPassword string `json:"newPassword"`
}

type CustomerCreateEmailToken struct {
	// Unique identifier of the Customer.
	ID string `json:"id"`
	// Expected version of the Customer.
	Version *int `json:"version,omitempty"`
	// Validity period of the generated token in minutes.
	TtlMinutes int `json:"ttlMinutes"`
}

type CustomerCreatePasswordResetToken struct {
	// Email address of the Customer treated as [case-insensitive](/../api/customers-overview#email-case-insensitivity).
	Email string `json:"email"`
	// Validity period of the generated token in minutes.
	TtlMinutes *int `json:"ttlMinutes,omitempty"`
}

type CustomerDraft struct {
	// User-defined unique identifier for the Customer.
	// The `key` field is preferred over `customerNumber` as it is mutable and provides more flexibility.
	Key *string `json:"key,omitempty"`
	// User-defined unique identifier for a Customer.
	// Once set, it cannot be changed.
	//
	// Can be used to refer to a Customer in a human-readable way (in emails, invoices, and other correspondence).
	CustomerNumber *string `json:"customerNumber,omitempty"`
	// Optional identifier for use in external systems like Customer Relationship Management (CRM) or Enterprise Resource Planning (ERP).
	ExternalId *string `json:"externalId,omitempty"`
	// Email address of the Customer that must be [unique](/../api/customers-overview#customer-uniqueness) for an entire Project or to a Store the Customer is assigned to.
	// It is the mandatory unique identifier of a Customer.
	Email string `json:"email"`
	// Required when `authenticationMode` is set to `Password`.
	// Provide the Customer's password in plain text. The API stores passwords in an encrypted format.
	Password *string `json:"password,omitempty"`
	// Given name (first name) of the Customer.
	FirstName *string `json:"firstName,omitempty"`
	// Family name (last name) of the Customer.
	LastName *string `json:"lastName,omitempty"`
	// Middle name of the Customer.
	MiddleName *string `json:"middleName,omitempty"`
	// Title of the Customer, for example, 'Dr.'.
	Title *string `json:"title,omitempty"`
	// Deprecated since an anonymous [Cart](ctp:api:type:Cart) can be identified by its `id` or external `key`.
	AnonymousCartId *string `json:"anonymousCartId,omitempty"`
	// Identifies a [Cart](ctp:api:type:Cart) that will be assigned to the new Customer.
	AnonymousCart *CartResourceIdentifier `json:"anonymousCart,omitempty"`
	// Identifies Carts and Orders belonging to an anonymous session that will be assigned to the new Customer.
	AnonymousId *string `json:"anonymousId,omitempty"`
	// Date of birth of the Customer.
	DateOfBirth *Date `json:"dateOfBirth,omitempty"`
	// Company name of the Customer. When representing a company as a Customer, [Business Units](ctp:api:type:BusinessUnit) provide extended funtionality.
	CompanyName *string `json:"companyName,omitempty"`
	// Individual VAT ID of the Customer.
	VatId *string `json:"vatId,omitempty"`
	// Addresses of the Customer.
	Addresses []BaseAddress `json:"addresses"`
	// Index of the address in the `addresses` array to use as the default shipping address.
	// The `defaultShippingAddressId` of the Customer will be set to the `id` of that address.
	DefaultShippingAddress *int `json:"defaultShippingAddress,omitempty"`
	// Indices of the shipping addresses in the `addresses` array.
	// The `shippingAddressIds` of the Customer will be set to the IDs of these addresses.
	ShippingAddresses []int `json:"shippingAddresses"`
	// Index of the address in the `addresses` array to use as the default billing address.
	// The `defaultBillingAddressId` of the Customer will be set to the `id` of that address.
	DefaultBillingAddress *int `json:"defaultBillingAddress,omitempty"`
	// Indices of the billing addresses in the `addresses` array.
	// The `billingAddressIds` of the Customer will be set to the IDs of these addresses.
	BillingAddresses []int `json:"billingAddresses"`
	// Set to `true` if the email address of the Customer has been verified already.
	// The intended use is to leave this field unset upon sign-up of the Customer and initiate the [email verification](#email-verification-of-customer) afterwards.
	IsEmailVerified *bool `json:"isEmailVerified,omitempty"`
	// Sets the [CustomerGroup](ctp:api:type:CustomerGroup) for the Customer.
	CustomerGroup *CustomerGroupResourceIdentifier `json:"customerGroup,omitempty"`
	// Custom Fields for the Customer.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
	// Preferred language of the Customer.
	// Must be one of the languages supported by the [Project](ctp:api:type:Project).
	Locale *string `json:"locale,omitempty"`
	// Salutation of the Customer, for example, 'Mr.' or 'Mrs.'.
	Salutation *string `json:"salutation,omitempty"`
	// Sets the [Stores](ctp:api:type:Store) for the Customer.
	//
	// - If no Stores are specified, the Customer is a global customer, and can log in using the [Password Flow for global Customers](/../api/authorization#password-flow-for-global-customers).
	// - If any Stores are specified, the Customer can only log in using the [Password Flow for Customers in a Store](/../api/authorization#password-flow-for-customers-in-a-store) for those specific Stores.
	Stores []StoreResourceIdentifier `json:"stores"`
	// - Set to `Password` to make the `password` field required for the Customer.
	// - Set to `ExternalAuth` when the password is not required for the Customer.
	AuthenticationMode *AuthenticationMode `json:"authenticationMode,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerDraft) MarshalJSON() ([]byte, error) {
	type Alias CustomerDraft
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

	if raw["shippingAddresses"] == nil {
		delete(raw, "shippingAddresses")
	}

	if raw["billingAddresses"] == nil {
		delete(raw, "billingAddresses")
	}

	if raw["stores"] == nil {
		delete(raw, "stores")
	}

	return json.Marshal(raw)

}

/**
*	[Reference](ctp:api:type:Reference) to a [CustomerToken](ctp:api:type:CustomerToken) for email verification.
*
 */
type CustomerEmailTokenReference struct {
	// Unique identifier of the referenced [CustomerToken](ctp:api:type:CustomerToken).
	ID string `json:"id"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerEmailTokenReference) MarshalJSON() ([]byte, error) {
	type Alias CustomerEmailTokenReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "customer-email-token", Alias: (*Alias)(&obj)})
}

type CustomerEmailVerify struct {
	// Expected version of the Customer.
	Version *int `json:"version,omitempty"`
	// Value of the token to verify Customer email.
	TokenValue string `json:"tokenValue"`
}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with results containing an array of [Customer](ctp:api:type:Customer).
*
 */
type CustomerPagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/../api/general-concepts#strong-consistency).
	// This field is returned by default.
	// For improved performance, calculating this field can be deactivated by using the query parameter `withTotal=false`.
	// When the results are filtered with a [Query Predicate](/../api/predicates/query), `total` is subject to a [limit](/../api/limits#queries).
	Total *int `json:"total,omitempty"`
	// [Customers](ctp:api:type:Customer) matching the query.
	Results []Customer `json:"results"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [CustomerToken](ctp:api:type:CustomerToken) for password reset.
*
 */
type CustomerPasswordTokenReference struct {
	// Unique identifier of the referenced [CustomerToken](ctp:api:type:CustomerToken).
	ID string `json:"id"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerPasswordTokenReference) MarshalJSON() ([]byte, error) {
	type Alias CustomerPasswordTokenReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "customer-password-token", Alias: (*Alias)(&obj)})
}

/**
*	[Reference](ctp:api:type:Reference) to a [Customer](ctp:api:type:Customer).
*
 */
type CustomerReference struct {
	// Unique identifier of the referenced [Customer](ctp:api:type:Customer).
	ID string `json:"id"`
	// Contains the representation of the expanded Customer. Only present in responses to requests with [Reference Expansion](ctp:api:type:Expansion) for Customers.
	Obj *Customer `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerReference) MarshalJSON() ([]byte, error) {
	type Alias CustomerReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "customer", Alias: (*Alias)(&obj)})
}

type CustomerResetPassword struct {
	// Value of the token to reset the Customer password.
	TokenValue string `json:"tokenValue"`
	// New password to be set.
	NewPassword string `json:"newPassword"`
	// Expected version of the Customer.
	Version *int `json:"version,omitempty"`
}

/**
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [Customer](ctp:api:type:Customer). Either `id` or `key` is required. If both are set, an [InvalidJsonInput](/../api/errors#invalidjsoninput) error is returned.
*
 */
type CustomerResourceIdentifier struct {
	// Unique identifier of the referenced [Customer](ctp:api:type:Customer). Required if `key` is absent.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced [Customer](ctp:api:type:Customer). Required if `id` is absent.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias CustomerResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "customer", Alias: (*Alias)(&obj)})
}

type CustomerSignInResult struct {
	// Customer [signed up](ctp:api:endpoint:/{projectKey}/customers:POST) or [signed in](ctp:api:endpoint:/{projectKey}/login:POST) after authentication.
	Customer Customer `json:"customer"`
	// Cart associated with the Customer.
	// If empty, the Customer does not have a Cart assigned.
	Cart *Cart `json:"cart,omitempty"`
}

type CustomerSignin struct {
	// Email address of the Customer treated as [case-insensitive](/../api/customers-overview#email-case-insensitivity).
	Email string `json:"email"`
	// Password of the Customer.
	Password string `json:"password"`
	// Deprecated since it is now possible to identify an anonymous cart by using its `id` or external `key`.
	AnonymousCartId *string `json:"anonymousCartId,omitempty"`
	// Identifies a [Cart](ctp:api:type:Cart) that will be assigned to the Customer.
	AnonymousCart *CartResourceIdentifier `json:"anonymousCart,omitempty"`
	// - Set to `MergeWithExistingCustomerCart` if [LineItems](ctp:api:type:LineItem) of the anonymous Cart should be merged with the active Customer Cart that has been modified most recently.
	// - Set to `UseAsNewActiveCustomerCart` if the anonymous Cart should be used as the new active Customer Cart and no [LineItems](ctp:api:type:LineItem) are to be merged.
	AnonymousCartSignInMode *AnonymousCartSignInMode `json:"anonymousCartSignInMode,omitempty"`
	// If both `anonymousCart` and `anonymousId` are provided, the `anonymousId` on the CustomerSignin must match that of the anonymous [Cart](ctp:api:type:Cart).
	// Otherwise a [400 Bad Request](ctp:api:type:InvalidOperationError) `Invalid Operation` error is returned with the message:
	// "Cart with the ID cart-id does not have the expected anonymousId.".
	AnonymousId *string `json:"anonymousId,omitempty"`
	// - If `true`, the [LineItem](ctp:api:type:LineItem) Product data (`name`, `variant`, and `productType`) of the returned Cart will be updated.
	// - If `false`, only the prices, discounts, and tax rates will be updated.
	UpdateProductData *bool `json:"updateProductData,omitempty"`
}

type CustomerToken struct {
	// Unique identifier of the token.
	ID string `json:"id"`
	// The `id` of the Customer.
	CustomerId string `json:"customerId"`
	// Value of the token.
	Value string `json:"value"`
	// Date and time (UTC) the token expires.
	ExpiresAt time.Time `json:"expiresAt"`
	// Date and time (UTC) the token was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// When the token is created, `lastModifiedAt` is set to `createdAt`.
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`
}

type CustomerUpdate struct {
	// Expected version of the Customer on which the changes should be applied.
	// If the expected version does not match the actual version, a [ConcurrentModification](ctp:api:type:ConcurrentModificationError) error will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the Customer.
	Actions []CustomerUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CustomerUpdate) UnmarshalJSON(data []byte) error {
	type Alias CustomerUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorCustomerUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type CustomerUpdateAction interface{}

func mapDiscriminatorCustomerUpdateAction(input interface{}) (CustomerUpdateAction, error) {
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
		obj := CustomerAddAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addBillingAddressId":
		obj := CustomerAddBillingAddressIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addShippingAddressId":
		obj := CustomerAddShippingAddressIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addStore":
		obj := CustomerAddStoreAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeAddress":
		obj := CustomerChangeAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeEmail":
		obj := CustomerChangeEmailAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeAddress":
		obj := CustomerRemoveAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeBillingAddressId":
		obj := CustomerRemoveBillingAddressIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeShippingAddressId":
		obj := CustomerRemoveShippingAddressIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeStore":
		obj := CustomerRemoveStoreAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAddressCustomField":
		obj := CustomerSetAddressCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAddressCustomType":
		obj := CustomerSetAddressCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAuthenticationMode":
		obj := CustomerSetAuthenticationModeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCompanyName":
		obj := CustomerSetCompanyNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomField":
		obj := CustomerSetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomType":
		obj := CustomerSetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomerGroup":
		obj := CustomerSetCustomerGroupAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomerNumber":
		obj := CustomerSetCustomerNumberAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDateOfBirth":
		obj := CustomerSetDateOfBirthAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDefaultBillingAddress":
		obj := CustomerSetDefaultBillingAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDefaultShippingAddress":
		obj := CustomerSetDefaultShippingAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setExternalId":
		obj := CustomerSetExternalIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setFirstName":
		obj := CustomerSetFirstNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := CustomerSetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLastName":
		obj := CustomerSetLastNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLocale":
		obj := CustomerSetLocaleAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMiddleName":
		obj := CustomerSetMiddleNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setSalutation":
		obj := CustomerSetSalutationAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setStores":
		obj := CustomerSetStoresAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setTitle":
		obj := CustomerSetTitleAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setVatId":
		obj := CustomerSetVatIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type MyCustomerChangePassword struct {
	// Expected version of the Customer on which the changes should be applied.
	Version int `json:"version"`
	// Current password of the Customer.
	//
	// If the current password does not match, an [InvalidCurrentPassword](ctp:api:type:InvalidCurrentPasswordError) error is returned.
	CurrentPassword string `json:"currentPassword"`
	// New password to be set.
	NewPassword string `json:"newPassword"`
}

type MyCustomerEmailVerify struct {
	// Value of the token to verify Customer email.
	TokenValue string `json:"tokenValue"`
}

type MyCustomerResetPassword struct {
	// Value of the token to reset the Customer password.
	TokenValue string `json:"tokenValue"`
	// New password to be set.
	NewPassword string `json:"newPassword"`
}

type MyCustomerSignin struct {
	// Email address of the Customer treated as [case-insensitive](/../api/customers-overview#email-case-insensitivity).
	Email string `json:"email"`
	// Password of the Customer.
	Password string `json:"password"`
	// - If `MergeWithExistingCustomerCart`, [LineItems](ctp:api:type:LineItem) of the anonymous Cart are merged with the recently modified active Customer Cart.
	// - If `UseAsNewActiveCustomerCart`, the anonymous Cart is used as the new active Customer Cart, and no [LineItems](ctp:api:type:LineItem) are merged.
	ActiveCartSignInMode *AnonymousCartSignInMode `json:"activeCartSignInMode,omitempty"`
	// - If `true`, the [LineItem](ctp:api:type:LineItem) Product data (`name`, `variant`, and `productType`) of the returned Cart is updated.
	// - If `false`, only the prices, discounts, and tax rates are updated.
	UpdateProductData *bool `json:"updateProductData,omitempty"`
}

/**
*	Adding an address to the Customer produces the [CustomerAddressAdded](ctp:api:type:CustomerAddressAddedMessage) Message.
*
 */
type CustomerAddAddressAction struct {
	// Value to append to the `addresses` array.
	Address BaseAddress `json:"address"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerAddAddressAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addAddress", Alias: (*Alias)(&obj)})
}

/**
*	Adds an Address from the `addresses` array to `billingAddressIds`. Either `addressId` or `addressKey` is required.
*
 */
type CustomerAddBillingAddressIdAction struct {
	// `id` of the [Address](ctp:api:type:Address) to become a billing address.
	AddressId *string `json:"addressId,omitempty"`
	// `key` of the [Address](ctp:api:type:Address) to become a billing address.
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerAddBillingAddressIdAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddBillingAddressIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addBillingAddressId", Alias: (*Alias)(&obj)})
}

/**
*	Adds an Address from the `addresses` array to `shippingAddressIds`. Either `addressId` or `addressKey` is required.
*
 */
type CustomerAddShippingAddressIdAction struct {
	// `id` of the [Address](ctp:api:type:Address) to become a shipping address.
	AddressId *string `json:"addressId,omitempty"`
	// `key` of the [Address](ctp:api:type:Address) to become a shipping address.
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerAddShippingAddressIdAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddShippingAddressIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addShippingAddressId", Alias: (*Alias)(&obj)})
}

/**
*	Associates the Customer with a Store.
*
 */
type CustomerAddStoreAction struct {
	// ResourceIdentifier of the Store to add.
	Store StoreResourceIdentifier `json:"store"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerAddStoreAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddStoreAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addStore", Alias: (*Alias)(&obj)})
}

/**
*	Changing an address of the Customer produces the [CustomerAddressChanged](ctp:api:type:CustomerAddressChangedMessage) Message.
*
*	Either `addressId` or `addressKey` is required.
*
 */
type CustomerChangeAddressAction struct {
	// `id` of the [Address](ctp:api:type:Address) to change.
	AddressId *string `json:"addressId,omitempty"`
	// `key` of the [Address](ctp:api:type:Address) to change.
	AddressKey *string `json:"addressKey,omitempty"`
	// Value to set.
	Address BaseAddress `json:"address"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerChangeAddressAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerChangeAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAddress", Alias: (*Alias)(&obj)})
}

/**
*	Changes the `email` of the Customer and sets the `isEmailVerified` property to `false`. This update action generates a [CustomerEmailChanged](ctp:api:type:CustomerEmailChangedMessage) Message.
*
 */
type CustomerChangeEmailAction struct {
	// Value to set.
	Email string `json:"email"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerChangeEmailAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerChangeEmailAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeEmail", Alias: (*Alias)(&obj)})
}

/**
*	Removing an address from the Customer produces the [CustomerAddressRemoved](ctp:api:type:CustomerAddressRemovedMessage) Message.
*
*	Either `addressId` or `addressKey` is required.
*
 */
type CustomerRemoveAddressAction struct {
	// `id` of the [Address](ctp:api:type:Address) to remove.
	AddressId *string `json:"addressId,omitempty"`
	// `key` of the [Address](ctp:api:type:Address) to remove.
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerRemoveAddressAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerRemoveAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeAddress", Alias: (*Alias)(&obj)})
}

/**
*	Removes a billing address from `billingAddressesIds`.
*	If the billing address is the default billing address, the `defaultBillingAddressId` is unset. Either `addressId` or `addressKey` is required.
*
 */
type CustomerRemoveBillingAddressIdAction struct {
	// `id` of the [Address](ctp:api:type:Address) to remove from `billingAddressesIds`.
	AddressId *string `json:"addressId,omitempty"`
	// `key` of the [Address](ctp:api:type:Address) to remove from `billingAddressesIds`.
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerRemoveBillingAddressIdAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerRemoveBillingAddressIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeBillingAddressId", Alias: (*Alias)(&obj)})
}

/**
*	Removes a shipping address from `shippingAddressesIds`.
*	If the shipping address is the default shipping address, the `defaultShippingAddressId` is unset. Either `addressId` or `addressKey` is required.
*
 */
type CustomerRemoveShippingAddressIdAction struct {
	// `id` of the [Address](ctp:api:type:Address) to remove from `shippingAddressesIds`.
	AddressId *string `json:"addressId,omitempty"`
	// `key` of the [Address](ctp:api:type:Address) to remove from `shippingAddressesIds`.
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerRemoveShippingAddressIdAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerRemoveShippingAddressIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeShippingAddressId", Alias: (*Alias)(&obj)})
}

/**
*	Removes the association to a Store from the Customer.
*	If no more Stores are assigned, the Customer becomes a [global Customer](/../api/customers-overview#global-versus-store-specific-customers).
*
 */
type CustomerRemoveStoreAction struct {
	// ResourceIdentifier of the Store to remove.
	Store StoreResourceIdentifier `json:"store"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerRemoveStoreAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerRemoveStoreAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeStore", Alias: (*Alias)(&obj)})
}

/**
*	Adding a Custom Field to an Address of a Customer generates the [CustomerAddressCustomFieldAdded](ctp:api:type:CustomerAddressCustomFieldAddedMessage) Message, removing one generates the [CustomerAddressCustomFieldRemoved](ctp:api:type:CustomerAddressCustomFieldRemovedMessage) Message, and updating an existing one generates the [CustomerAddressCustomFieldChanged](ctp:api:type:CustomerAddressCustomFieldChangedMessage) Message.
*
 */
type CustomerSetAddressCustomFieldAction struct {
	// User-defined unique identifier of the [Address](ctp:api:type:Address) to be updated.
	AddressId string `json:"addressId"`
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// If `value` is provided, it is set for the field defined by `name`.
	// Trying to remove a field that does not exist will fail with an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetAddressCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetAddressCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAddressCustomField", Alias: (*Alias)(&obj)})
}

/**
*	Adding or updating a Custom Type on an Address of a Customer generates the [CustomerAddressCustomTypeSet](ctp:api:type:CustomerAddressCustomTypeSetMessage) Message, and removing one generates the [CustomerAddressCustomTypeRemoved](ctp:api:type:CustomerAddressCustomTypeRemovedMessage) Message.
*
 */
type CustomerSetAddressCustomTypeAction struct {
	// User-defined unique identifier of the [Address](ctp:api:type:Address) to be updated.
	AddressId string `json:"addressId"`
	// Defines the [Type](ctp:api:type:Type) that extends the `address` with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the `address`.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the `address`.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetAddressCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetAddressCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAddressCustomType", Alias: (*Alias)(&obj)})
}

type CustomerSetAuthenticationModeAction struct {
	// Value to set.
	// Changing a Customer's `authMode` from `Password` to `ExternalAuth` deletes the Customer's password.
	AuthMode AuthenticationMode `json:"authMode"`
	// Required when `authMode` is `Password`.
	Password *string `json:"password,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetAuthenticationModeAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetAuthenticationModeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAuthenticationMode", Alias: (*Alias)(&obj)})
}

/**
*	Setting a company name produces the [CustomerCompanyNameSet](ctp:api:type:CustomerCompanyNameSetMessage) Message.
*
 */
type CustomerSetCompanyNameAction struct {
	// Value to set.
	// If empty, any existing value is removed.
	CompanyName *string `json:"companyName,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetCompanyNameAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetCompanyNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCompanyName", Alias: (*Alias)(&obj)})
}

/**
*	Adding a Custom Field to a Customer generates the [CustomerCustomFieldAdded](ctp:api:type:CustomerCustomFieldAddedMessage) Message, removing one generates the [CustomerCustomFieldRemoved](ctp:api:type:CustomerCustomFieldRemovedMessage) Message, and updating an existing one generates the [CustomerCustomFieldChanged](ctp:api:type:CustomerCustomFieldChangedMessage) Message.
*
 */
type CustomerSetCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// If `value` is provided, it is set for the field defined by `name`.
	// Trying to remove a field that does not exist will fail with an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

/**
*	Adding or updating a Custom Type on a Customer generates the [CustomerCustomTypeSet](ctp:api:type:CustomerCustomTypeSetMessage) Message, removing one generates the [CustomerCustomTypeRemoved](ctp:api:type:CustomerCustomTypeRemovedMessage) Message.
*
 */
type CustomerSetCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the Customer with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the Customer.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the Customer.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

/**
*	Setting the Customer Group of the Customer produces the [CustomerGroupSet](ctp:api:type:CustomerGroupSetMessage) Message.
*
 */
type CustomerSetCustomerGroupAction struct {
	// Value to set.
	// If empty, any existing value is removed.
	CustomerGroup *CustomerGroupResourceIdentifier `json:"customerGroup,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetCustomerGroupAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetCustomerGroupAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomerGroup", Alias: (*Alias)(&obj)})
}

/**
*	Sets a new ID that can be used to refer to a Customer in a human-reabable way (for use in emails, invoices, etc).
*
 */
type CustomerSetCustomerNumberAction struct {
	// Value to set.
	// Once set, it cannot be changed.
	CustomerNumber *string `json:"customerNumber,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetCustomerNumberAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetCustomerNumberAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomerNumber", Alias: (*Alias)(&obj)})
}

/**
*	Setting the date of birth of the Customer produces the [CustomerDateOfBirthSet](ctp:api:type:CustomerDateOfBirthSetMessage) Message.
*
 */
type CustomerSetDateOfBirthAction struct {
	// Value to set.
	// If empty, any existing value is removed.
	DateOfBirth *Date `json:"dateOfBirth,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetDateOfBirthAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetDateOfBirthAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDateOfBirth", Alias: (*Alias)(&obj)})
}

/**
*	Sets the default billing address from `addresses`.
*	The action adds the `id` of the specified Address to the `billingAddressIds` if not contained already. Either `addressId` or `addressKey` is required.
*
 */
type CustomerSetDefaultBillingAddressAction struct {
	// `id` of the [Address](ctp:api:type:Address) to become the default billing address.
	AddressId *string `json:"addressId,omitempty"`
	// `key` of the [Address](ctp:api:type:Address) to become the default billing address.
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetDefaultBillingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetDefaultBillingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDefaultBillingAddress", Alias: (*Alias)(&obj)})
}

/**
*	Sets the default shipping address from `addresses`.
*	The action adds the `id` of the specified address to the `shippingAddressIds` if not contained already. Either `addressId` or `addressKey` is required.
*
*	If the Tax Category of the Cart [ShippingInfo](ctp:api:type:ShippingInfo) is missing the TaxRate matching country and state given in the `shippingAddress` of that Cart, a [MissingTaxRateForCountry](ctp:api:type:MissingTaxRateForCountryError) error is returned.
*
 */
type CustomerSetDefaultShippingAddressAction struct {
	// `id` of the [Address](ctp:api:type:Address) to become the default shipping address.
	AddressId *string `json:"addressId,omitempty"`
	// `key` of the [Address](ctp:api:type:Address) to become the default shipping address.
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetDefaultShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetDefaultShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDefaultShippingAddress", Alias: (*Alias)(&obj)})
}

type CustomerSetExternalIdAction struct {
	// Value to set.
	// If empty, any existing value is removed.
	ExternalId *string `json:"externalId,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetExternalIdAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetExternalIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setExternalId", Alias: (*Alias)(&obj)})
}

/**
*	Setting the first name of the Customer produces the [CustomeFirstNameSet](ctp:api:type:CustomerFirstNameSetMessage) Message.
*
 */
type CustomerSetFirstNameAction struct {
	// Value to set. If empty, any existing value is removed.
	FirstName *string `json:"firstName,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetFirstNameAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetFirstNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setFirstName", Alias: (*Alias)(&obj)})
}

type CustomerSetKeyAction struct {
	// If `key` is absent or `null`, the existing key, if any, will be removed.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

/**
*	Setting the last name of the Customer produces the [CustomerLastNameSet](ctp:api:type:CustomerLastNameSetMessage) Message.
*
 */
type CustomerSetLastNameAction struct {
	// Value to set. If empty, any existing value is removed.
	LastName *string `json:"lastName,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetLastNameAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetLastNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLastName", Alias: (*Alias)(&obj)})
}

type CustomerSetLocaleAction struct {
	// Value to set.
	// Must be one of the languages supported by the [Project](ctp:api:type:Project).
	Locale *string `json:"locale,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetLocaleAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetLocaleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLocale", Alias: (*Alias)(&obj)})
}

type CustomerSetMiddleNameAction struct {
	// Value to set. If empty, any existing value is removed.
	MiddleName *string `json:"middleName,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetMiddleNameAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetMiddleNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMiddleName", Alias: (*Alias)(&obj)})
}

type CustomerSetSalutationAction struct {
	// Value to set. If empty, any existing value is removed.
	Salutation *string `json:"salutation,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetSalutationAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetSalutationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSalutation", Alias: (*Alias)(&obj)})
}

/**
*	Sets the Stores the Customer account is associated with.
*	If no Stores are specified, the Customer becomes a [global Customer](/../api/customers-overview#global-versus-store-specific-customers).
*
 */
type CustomerSetStoresAction struct {
	// ResourceIdentifier of the Stores to set.
	Stores []StoreResourceIdentifier `json:"stores"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetStoresAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetStoresAction
	data, err := json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setStores", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["stores"] == nil {
		delete(raw, "stores")
	}

	return json.Marshal(raw)

}

/**
*	Setting the title of the Customer produces the [CustomerTitleSet](ctp:api:type:CustomerTitleSetMessage) Message.
*
 */
type CustomerSetTitleAction struct {
	// Value to set. If empty, any existing value is removed.
	Title *string `json:"title,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetTitleAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetTitleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTitle", Alias: (*Alias)(&obj)})
}

type CustomerSetVatIdAction struct {
	// Value to set.
	// If empty, any existing value is removed.
	VatId *string `json:"vatId,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerSetVatIdAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetVatIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setVatId", Alias: (*Alias)(&obj)})
}
