package importapi

// Generated file, please do not change!!!

import (
	"encoding/json"
)

type AuthenticationMode string

const (
	AuthenticationModePassword     AuthenticationMode = "Password"
	AuthenticationModeExternalAuth AuthenticationMode = "ExternalAuth"
)

/**
*	Different from [Address](/types#address) in that `key` is required and `id` is not supported.
*
 */
type CustomerAddress struct {
	// User-defined identifier for the address.
	// Must follow the pattern `[a-zA-Z0-9_\-]{2,256}` and must be unique per customer.
	Key string `json:"key"`
	// Name of the country.
	Country string `json:"country"`
	// Title of the contact, for example 'Dr.'
	Title *string `json:"title,omitempty"`
	// Salutation of the contact, for example 'Mr.' or 'Ms.'
	Salutation *string `json:"salutation,omitempty"`
	// Given name (first name) of the contact.
	FirstName *string `json:"firstName,omitempty"`
	// Family name (last name) of the contact.
	LastName *string `json:"lastName,omitempty"`
	// Name of the street.
	StreetName *string `json:"streetName,omitempty"`
	// Street number.
	StreetNumber *string `json:"streetNumber,omitempty"`
	// Further information on the street address.
	AdditionalStreetInfo *string `json:"additionalStreetInfo,omitempty"`
	// Postal code.
	PostalCode *string `json:"postalCode,omitempty"`
	// Name of the city.
	City *string `json:"city,omitempty"`
	// Name of the region.
	Region *string `json:"region,omitempty"`
	// Name of the state, for example, Colorado.
	State *string `json:"state,omitempty"`
	// Name of the company.
	Company *string `json:"company,omitempty"`
	// Name of the department.
	Department *string `json:"department,omitempty"`
	// Number or name of the building.
	Building *string `json:"building,omitempty"`
	// Number or name of the apartment.
	Apartment *string `json:"apartment,omitempty"`
	// Post office box number.
	POBox *string `json:"pOBox,omitempty"`
	// Phone number of the contact.
	Phone *string `json:"phone,omitempty"`
	// Mobile phone number of the contact.
	Mobile *string `json:"mobile,omitempty"`
	// Email address of the contact.
	Email *string `json:"email,omitempty"`
	// Fax number of the contact.
	Fax *string `json:"fax,omitempty"`
	// Further information on the Address.
	AdditionalAddressInfo *string `json:"additionalAddressInfo,omitempty"`
	// ID for the contact used in an external system.
	ExternalId *string `json:"externalId,omitempty"`
	// Custom Fields for the address.
	Custom *Custom `json:"custom,omitempty"`
}

/**
*	The data representation for a Customer to be imported that is persisted as a [Customer](ctp:api:type:Customer) in the Project.
*
 */
type CustomerImport struct {
	// User-defined unique identifier. If a [Customer](ctp:api:type:Customer) with this `key` exists, it will be updated with the imported data.
	Key string `json:"key"`
	// Maps to `Customer.customerNumber`.
	CustomerNumber *string `json:"customerNumber,omitempty"`
	// Maps to `Customer.email`.
	Email string `json:"email"`
	// Required when `authenticationMode` is set to `Password`. Maps to `Customer.password`.
	Password *string `json:"password,omitempty"`
	// The References to the Stores with which the Customer is associated. If referenced Stores do not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the necessary Stores are created.
	Stores []StoreKeyReference `json:"stores"`
	// Maps to `Customer.firstName`.
	FirstName *string `json:"firstName,omitempty"`
	// Maps to `Customer.lastName`.
	LastName *string `json:"lastName,omitempty"`
	// Maps to `Customer.middleName`.
	MiddleName *string `json:"middleName,omitempty"`
	// Maps to `Customer.title`.
	Title *string `json:"title,omitempty"`
	// Maps to `Customer.salutation`.
	Salutation *string `json:"salutation,omitempty"`
	// Maps to `Customer.externalId`.
	ExternalId *string `json:"externalId,omitempty"`
	// Maps to `Customer.dateOfBirth`.
	DateOfBirth *Date `json:"dateOfBirth,omitempty"`
	// Maps to `Customer.companyName`.
	CompanyName *string `json:"companyName,omitempty"`
	// Maps to `Customer.vatId`.
	VatId *string `json:"vatId,omitempty"`
	// Maps to `Customer.isEmailVerified`.
	IsEmailVerified *bool `json:"isEmailVerified,omitempty"`
	// The Reference to the [CustomerGroup](ctp:api:type:CustomerGroup) with which the Customer is associated.
	// If referenced CustomerGroup does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the necessary CustomerGroup is created.
	CustomerGroup *CustomerGroupKeyReference `json:"customerGroup,omitempty"`
	// Maps to `Customer.addresses`.
	Addresses []CustomerAddress `json:"addresses"`
	// The index of the address in the addresses array. The `defaultBillingAddressId` of the customer will be set to the ID of that address.
	DefaultBillingAddress *int `json:"defaultBillingAddress,omitempty"`
	// The indices of the billing addresses in the addresses array. The `billingAddressIds` of the customer will be set to the IDs of that addresses.
	BillingAddresses []int `json:"billingAddresses"`
	// The index of the address in the addresses array. The `defaultShippingAddressId` of the customer will be set to the ID of that address.
	DefaultShippingAddress *int `json:"defaultShippingAddress,omitempty"`
	// The indices of the shipping addresses in the addresses array. The `shippingAddressIds` of the customer will be set to the IDs of that addresses.
	ShippingAddresses []int `json:"shippingAddresses"`
	// Maps to `Customer.locale`.
	Locale *string `json:"locale,omitempty"`
	// The Custom Fields for this Customer.
	Custom *Custom `json:"custom,omitempty"`
	// - Set to `Password` to make the `password` field required for the Customer.
	// - Set to `ExternalAuth` when the password is not required for the Customer.
	AuthenticationMode *AuthenticationMode `json:"authenticationMode,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerImport) MarshalJSON() ([]byte, error) {
	type Alias CustomerImport
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

	if raw["stores"] == nil {
		delete(raw, "stores")
	}

	if raw["addresses"] == nil {
		delete(raw, "addresses")
	}

	if raw["billingAddresses"] == nil {
		delete(raw, "billingAddresses")
	}

	if raw["shippingAddresses"] == nil {
		delete(raw, "shippingAddresses")
	}

	return json.Marshal(raw)

}
