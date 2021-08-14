// Generated file, please do not change!!!
package importapi

/**
*	Different from Address in that `key` is required and `id` is not supported.
*
 */
type CustomerAddress struct {
	// User-defined identifier for the address.
	// Must follow the pattern `[a-zA-Z0-9_\-]{2,256}` and must be unique per customer.
	Key                  string  `json:"key"`
	Title                *string `json:"title,omitempty"`
	Salutation           *string `json:"salutation,omitempty"`
	FirstName            *string `json:"firstName,omitempty"`
	LastName             *string `json:"lastName,omitempty"`
	StreetName           *string `json:"streetName,omitempty"`
	StreetNumber         *string `json:"streetNumber,omitempty"`
	AdditionalStreetInfo *string `json:"additionalStreetInfo,omitempty"`
	PostalCode           *string `json:"postalCode,omitempty"`
	City                 *string `json:"city,omitempty"`
	Region               *string `json:"region,omitempty"`
	State                *string `json:"state,omitempty"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country               string  `json:"country"`
	Company               *string `json:"company,omitempty"`
	Department            *string `json:"department,omitempty"`
	Building              *string `json:"building,omitempty"`
	Apartment             *string `json:"apartment,omitempty"`
	POBox                 *string `json:"pOBox,omitempty"`
	Phone                 *string `json:"phone,omitempty"`
	Mobile                *string `json:"mobile,omitempty"`
	Email                 *string `json:"email,omitempty"`
	Fax                   *string `json:"fax,omitempty"`
	AdditionalAddressInfo *string `json:"additionalAddressInfo,omitempty"`
	ExternalId            *string `json:"externalId,omitempty"`
}

/**
*	The data representation for a Customer to be imported that is persisted as a [Customer](/../api/projects/customers#top) in the Project.
*
 */
type CustomerImport struct {
	Key string `json:"key"`
	// Maps to `Customer.customerNumber`.
	CustomerNumber *string `json:"customerNumber,omitempty"`
	// Maps to `Customer.email`.
	Email string `json:"email"`
	// Maps to `Customer.password`.
	Password string `json:"password"`
	// The References to the Stores with which the Customer is associated. If referenced Stores do not exist, the `state` of the [ImportOperation](/import-operation#importoperation) will be set to `unresolved` until the necessary Stores are created.
	Stores []StoreKeyReference `json:"stores,omitempty"`
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
	// The Reference to the [CustomerGroup](/../api/projects/customerGroups#customergroup) with which the Customer is associated.
	// If referenced CustomerGroup does not exist, the `state` of the [ImportOperation](/import-operation#importoperation) will be set to `unresolved` until the necessary CustomerGroup is created.
	CustomerGroup *CustomerGroupKeyReference `json:"customerGroup,omitempty"`
	// Maps to `Customer.addresses`.
	Addresses []CustomerAddress `json:"addresses"`
	// The index of the address in the addresses array. The `defaultBillingAddressId` of the customer will be set to the ID of that address.
	DefaultBillingAddress *int `json:"defaultBillingAddress,omitempty"`
	// The indices of the billing addresses in the addresses array. The `billingAddressIds` of the customer will be set to the IDs of that addresses.
	BillingAddresses []int `json:"billingAddresses,omitempty"`
	// The index of the address in the addresses array. The `defaultShippingAddressId` of the customer will be set to the ID of that address.
	DefaultShippingAddress *int `json:"defaultShippingAddress,omitempty"`
	// The indices of the shipping addresses in the addresses array. The `shippingAddressIds` of the customer will be set to the IDs of that addresses.
	ShippingAddresses []int `json:"shippingAddresses,omitempty"`
	// Maps to `Customer.locale`.
	Locale *string `json:"locale,omitempty"`
	// The custom fields for this Customer.
	Custom *Custom `json:"custom,omitempty"`
}
