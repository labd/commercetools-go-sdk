// Generated file, please do not change!!!
package importapi

import (
	"encoding/json"
)

/**
*	Different from Address in that `key` is required and `id` is not supported.
*
 */
type CustomerAddress struct {
	// User-defined identifier for the address.
	// Must follow the pattern `[a-zA-Z0-9_\-]{2,256}` and must be unique per customer.
	Key                  string `json:"key"`
	Title                string `json:"title,omitEmpty"`
	Salutation           string `json:"salutation,omitEmpty"`
	FirstName            string `json:"firstName,omitEmpty"`
	LastName             string `json:"lastName,omitEmpty"`
	StreetName           string `json:"streetName,omitEmpty"`
	StreetNumber         string `json:"streetNumber,omitEmpty"`
	AdditionalStreetInfo string `json:"additionalStreetInfo,omitEmpty"`
	PostalCode           string `json:"postalCode,omitEmpty"`
	City                 string `json:"city,omitEmpty"`
	Region               string `json:"region,omitEmpty"`
	State                string `json:"state,omitEmpty"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country               string `json:"country"`
	Company               string `json:"company,omitEmpty"`
	Department            string `json:"department,omitEmpty"`
	Building              string `json:"building,omitEmpty"`
	Apartment             string `json:"apartment,omitEmpty"`
	POBox                 string `json:"pOBox,omitEmpty"`
	Phone                 string `json:"phone,omitEmpty"`
	Mobile                string `json:"mobile,omitEmpty"`
	Email                 string `json:"email,omitEmpty"`
	Fax                   string `json:"fax,omitEmpty"`
	AdditionalAddressInfo string `json:"additionalAddressInfo,omitEmpty"`
	ExternalId            string `json:"externalId,omitEmpty"`
}

/**
*	The data representation for a Customer to be imported that is persisted as a [Customer](/../api/projects/customers#top) in the Project.
*
 */
type CustomerImport struct {
	Key string `json:"key"`
	// Maps to `Customer.customerNumber`.
	CustomerNumber string `json:"customerNumber,omitEmpty"`
	// Maps to `Customer.email`.
	Email string `json:"email"`
	// Maps to `Customer.password`.
	Password string `json:"password"`
	// The References to the Stores with which the Customer is associated. If referenced Stores do not exist, the `state` of the [ImportOperation](/import-operation#importoperation) will be set to `Unresolved` until the necessary Stores are created.
	Stores []StoreKeyReference `json:"stores,omitEmpty"`
	// Maps to `Customer.firstName`.
	FirstName string `json:"firstName,omitEmpty"`
	// Maps to `Customer.lastName`.
	LastName string `json:"lastName,omitEmpty"`
	// Maps to `Customer.middleName`.
	MiddleName string `json:"middleName,omitEmpty"`
	// Maps to `Customer.title`.
	Title string `json:"title,omitEmpty"`
	// Maps to `Customer.salutation`.
	Salutation string `json:"salutation,omitEmpty"`
	// Maps to `Customer.externalId`.
	ExternalId string `json:"externalId,omitEmpty"`
	// Maps to `Customer.dateOfBirth`.
	DateOfBirth Date `json:"dateOfBirth,omitEmpty"`
	// Maps to `Customer.companyName`.
	CompanyName string `json:"companyName,omitEmpty"`
	// Maps to `Customer.vatId`.
	VatId string `json:"vatId,omitEmpty"`
	// Maps to `Customer.isEmailVerified`.
	IsEmailVerified bool `json:"isEmailVerified,omitEmpty"`
	// The Reference to the [CustomerGroup](/../api/projects/customerGroups#customergroup) with which the Customer is associated.
	// If referenced CustomerGroup does not exist, the `state` of the [ImportOperation](/import-operation#importoperation) will be set to `Unresolved` until the necessary CustomerGroup is created.
	CustomerGroup CustomerGroupKeyReference `json:"customerGroup,omitEmpty"`
	// Maps to `Customer.addresses`.
	Addresses []CustomerAddress `json:"addresses"`
	// The index of the address in the addresses array. The `defaultBillingAddressId` of the customer will be set to the ID of that address.
	DefaultBillingAddress int `json:"defaultBillingAddress,omitEmpty"`
	// The indices of the billing addresses in the addresses array. The `billingAddressIds` of the customer will be set to the IDs of that addresses.
	BillingAddresses []int `json:"billingAddresses,omitEmpty"`
	// The index of the address in the addresses array. The `defaultShippingAddressId` of the customer will be set to the ID of that address.
	DefaultShippingAddress int `json:"defaultShippingAddress,omitEmpty"`
	// The indices of the shipping addresses in the addresses array. The `shippingAddressIds` of the customer will be set to the IDs of that addresses.
	ShippingAddresses []int `json:"shippingAddresses,omitEmpty"`
	// Maps to `Customer.locale`.
	Locale string `json:"locale,omitEmpty"`
	// The custom fields for this Customer.
	Custom *Custom `json:"custom,omitEmpty"`
}
