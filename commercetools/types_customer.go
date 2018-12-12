// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// AnonymousCartSignInMode is an enum type
type AnonymousCartSignInMode string

// Enum values for AnonymousCartSignInMode
const (
	AnonymousCartSignInModeMergeWithExistingCustomerCart AnonymousCartSignInMode = "MergeWithExistingCustomerCart"
	AnonymousCartSignInModeUseAsNewActiveCustomerCart    AnonymousCartSignInMode = "UseAsNewActiveCustomerCart"
)

// CustomerUpdateAction uses action as discriminator attribute
type CustomerUpdateAction interface{}

func mapDiscriminatorCustomerUpdateAction(input CustomerUpdateAction) CustomerUpdateAction {
	discriminator := input.(map[string]interface{})["action"]
	switch discriminator {
	case "addAddress":
		new := CustomerAddAddressAction{}
		mapstructure.Decode(input, &new)
		return new
	case "addBillingAddressId":
		new := CustomerAddBillingAddressIDAction{}
		mapstructure.Decode(input, &new)
		return new
	case "addShippingAddressId":
		new := CustomerAddShippingAddressIDAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeAddress":
		new := CustomerChangeAddressAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeEmail":
		new := CustomerChangeEmailAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeAddress":
		new := CustomerRemoveAddressAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeBillingAddressId":
		new := CustomerRemoveBillingAddressIDAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeShippingAddressId":
		new := CustomerRemoveShippingAddressIDAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCompanyName":
		new := CustomerSetCompanyNameAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomField":
		new := CustomerSetCustomFieldAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomType":
		new := CustomerSetCustomTypeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomerGroup":
		new := CustomerSetCustomerGroupAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomerNumber":
		new := CustomerSetCustomerNumberAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setDateOfBirth":
		new := CustomerSetDateOfBirthAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setDefaultBillingAddress":
		new := CustomerSetDefaultBillingAddressAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setDefaultShippingAddress":
		new := CustomerSetDefaultShippingAddressAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setExternalId":
		new := CustomerSetExternalIDAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setFirstName":
		new := CustomerSetFirstNameAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setKey":
		new := CustomerSetKeyAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setLastName":
		new := CustomerSetLastNameAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setLocale":
		new := CustomerSetLocaleAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setMiddleName":
		new := CustomerSetMiddleNameAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setSalutation":
		new := CustomerSetSalutationAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setTitle":
		new := CustomerSetTitleAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setVatId":
		new := CustomerSetVatIDAction{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}

// Customer is of type Resource
type Customer struct {
	Version                  int                     `json:"version"`
	LastModifiedAt           time.Time               `json:"lastModifiedAt"`
	ID                       string                  `json:"id"`
	CreatedAt                time.Time               `json:"createdAt"`
	VatID                    string                  `json:"vatId,omitempty"`
	Title                    string                  `json:"title,omitempty"`
	ShippingAddressIds       []string                `json:"shippingAddressIds,omitempty"`
	Salutation               string                  `json:"salutation,omitempty"`
	Password                 string                  `json:"password"`
	MiddleName               string                  `json:"middleName,omitempty"`
	Locale                   string                  `json:"locale,omitempty"`
	LastName                 string                  `json:"lastName,omitempty"`
	Key                      string                  `json:"key,omitempty"`
	IsEmailVerified          bool                    `json:"isEmailVerified"`
	FirstName                string                  `json:"firstName,omitempty"`
	ExternalID               string                  `json:"externalId,omitempty"`
	Email                    string                  `json:"email"`
	DefaultShippingAddressID string                  `json:"defaultShippingAddressId,omitempty"`
	DefaultBillingAddressID  string                  `json:"defaultBillingAddressId,omitempty"`
	DateOfBirth              interface{}             `json:"dateOfBirth,omitempty"`
	CustomerNumber           string                  `json:"customerNumber,omitempty"`
	CustomerGroup            *CustomerGroupReference `json:"customerGroup,omitempty"`
	Custom                   *CustomFields           `json:"custom,omitempty"`
	CompanyName              string                  `json:"companyName,omitempty"`
	BillingAddressIds        []string                `json:"billingAddressIds,omitempty"`
	Addresses                []Address               `json:"addresses"`
}

// CustomerAddAddressAction implements the interface CustomerUpdateAction
type CustomerAddAddressAction struct {
	Address *Address `json:"address"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerAddAddressAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addAddress", Alias: (*Alias)(&obj)})
}

// CustomerAddBillingAddressIDAction implements the interface CustomerUpdateAction
type CustomerAddBillingAddressIDAction struct {
	AddressID string `json:"addressId"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerAddBillingAddressIDAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddBillingAddressIDAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addBillingAddressId", Alias: (*Alias)(&obj)})
}

// CustomerAddShippingAddressIDAction implements the interface CustomerUpdateAction
type CustomerAddShippingAddressIDAction struct {
	AddressID string `json:"addressId"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerAddShippingAddressIDAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddShippingAddressIDAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addShippingAddressId", Alias: (*Alias)(&obj)})
}

// CustomerChangeAddressAction implements the interface CustomerUpdateAction
type CustomerChangeAddressAction struct {
	AddressID string   `json:"addressId"`
	Address   *Address `json:"address"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerChangeAddressAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerChangeAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAddress", Alias: (*Alias)(&obj)})
}

// CustomerChangeEmailAction implements the interface CustomerUpdateAction
type CustomerChangeEmailAction struct {
	Email string `json:"email"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerChangeEmailAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerChangeEmailAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeEmail", Alias: (*Alias)(&obj)})
}

// CustomerChangePassword is a standalone struct
type CustomerChangePassword struct {
	Version         int    `json:"version"`
	NewPassword     string `json:"newPassword"`
	ID              string `json:"id"`
	CurrentPassword string `json:"currentPassword"`
}

// CustomerCreateEmailToken is a standalone struct
type CustomerCreateEmailToken struct {
	Version    int    `json:"version,omitempty"`
	TTLMinutes int    `json:"ttlMinutes"`
	ID         string `json:"id"`
}

// CustomerCreatePasswordResetToken is a standalone struct
type CustomerCreatePasswordResetToken struct {
	Email string `json:"email"`
}

// CustomerDraft is a standalone struct
type CustomerDraft struct {
	VatID                  string                  `json:"vatId,omitempty"`
	Title                  string                  `json:"title,omitempty"`
	ShippingAddresses      []int                   `json:"shippingAddresses,omitempty"`
	Salutation             string                  `json:"salutation,omitempty"`
	Password               string                  `json:"password"`
	MiddleName             string                  `json:"middleName,omitempty"`
	Locale                 string                  `json:"locale,omitempty"`
	LastName               string                  `json:"lastName,omitempty"`
	Key                    string                  `json:"key,omitempty"`
	IsEmailVerified        bool                    `json:"isEmailVerified,omitempty"`
	FirstName              string                  `json:"firstName,omitempty"`
	ExternalID             string                  `json:"externalId,omitempty"`
	Email                  string                  `json:"email"`
	DefaultShippingAddress int                     `json:"defaultShippingAddress,omitempty"`
	DefaultBillingAddress  int                     `json:"defaultBillingAddress,omitempty"`
	DateOfBirth            interface{}             `json:"dateOfBirth,omitempty"`
	CustomerNumber         string                  `json:"customerNumber,omitempty"`
	CustomerGroup          *CustomerGroupReference `json:"customerGroup,omitempty"`
	Custom                 *CustomFieldsDraft      `json:"custom,omitempty"`
	CompanyName            string                  `json:"companyName,omitempty"`
	BillingAddresses       []int                   `json:"billingAddresses,omitempty"`
	AnonymousID            string                  `json:"anonymousId,omitempty"`
	AnonymousCartID        string                  `json:"anonymousCartId,omitempty"`
	Addresses              []Address               `json:"addresses,omitempty"`
}

// CustomerEmailVerify is a standalone struct
type CustomerEmailVerify struct {
	Version    int    `json:"version,omitempty"`
	TokenValue string `json:"tokenValue"`
}

// CustomerPagedQueryResponse is of type PagedQueryResponse
type CustomerPagedQueryResponse struct {
	Total   int        `json:"total,omitempty"`
	Offset  int        `json:"offset"`
	Count   int        `json:"count"`
	Results []Customer `json:"results"`
}

// CustomerReference implements the interface Reference
type CustomerReference struct {
	Key string    `json:"key,omitempty"`
	ID  string    `json:"id,omitempty"`
	Obj *Customer `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerReference) MarshalJSON() ([]byte, error) {
	type Alias CustomerReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "customer", Alias: (*Alias)(&obj)})
}

// CustomerRemoveAddressAction implements the interface CustomerUpdateAction
type CustomerRemoveAddressAction struct {
	AddressID string `json:"addressId"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerRemoveAddressAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerRemoveAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeAddress", Alias: (*Alias)(&obj)})
}

// CustomerRemoveBillingAddressIDAction implements the interface CustomerUpdateAction
type CustomerRemoveBillingAddressIDAction struct {
	AddressID string `json:"addressId"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerRemoveBillingAddressIDAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerRemoveBillingAddressIDAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeBillingAddressId", Alias: (*Alias)(&obj)})
}

// CustomerRemoveShippingAddressIDAction implements the interface CustomerUpdateAction
type CustomerRemoveShippingAddressIDAction struct {
	AddressID string `json:"addressId"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerRemoveShippingAddressIDAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerRemoveShippingAddressIDAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeShippingAddressId", Alias: (*Alias)(&obj)})
}

// CustomerResetPassword is a standalone struct
type CustomerResetPassword struct {
	Version     int    `json:"version,omitempty"`
	TokenValue  string `json:"tokenValue"`
	NewPassword string `json:"newPassword"`
}

// CustomerSetCompanyNameAction implements the interface CustomerUpdateAction
type CustomerSetCompanyNameAction struct {
	CompanyName string `json:"companyName,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerSetCompanyNameAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetCompanyNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCompanyName", Alias: (*Alias)(&obj)})
}

// CustomerSetCustomFieldAction implements the interface CustomerUpdateAction
type CustomerSetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

// CustomerSetCustomTypeAction implements the interface CustomerUpdateAction
type CustomerSetCustomTypeAction struct {
	Type   *TypeReference  `json:"type,omitempty"`
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

// CustomerSetCustomerGroupAction implements the interface CustomerUpdateAction
type CustomerSetCustomerGroupAction struct {
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerSetCustomerGroupAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetCustomerGroupAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomerGroup", Alias: (*Alias)(&obj)})
}

// CustomerSetCustomerNumberAction implements the interface CustomerUpdateAction
type CustomerSetCustomerNumberAction struct {
	CustomerNumber string `json:"customerNumber,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerSetCustomerNumberAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetCustomerNumberAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomerNumber", Alias: (*Alias)(&obj)})
}

// CustomerSetDateOfBirthAction implements the interface CustomerUpdateAction
type CustomerSetDateOfBirthAction struct {
	DateOfBirth interface{} `json:"dateOfBirth,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerSetDateOfBirthAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetDateOfBirthAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDateOfBirth", Alias: (*Alias)(&obj)})
}

// CustomerSetDefaultBillingAddressAction implements the interface CustomerUpdateAction
type CustomerSetDefaultBillingAddressAction struct {
	AddressID string `json:"addressId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerSetDefaultBillingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetDefaultBillingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDefaultBillingAddress", Alias: (*Alias)(&obj)})
}

// CustomerSetDefaultShippingAddressAction implements the interface CustomerUpdateAction
type CustomerSetDefaultShippingAddressAction struct {
	AddressID string `json:"addressId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerSetDefaultShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetDefaultShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDefaultShippingAddress", Alias: (*Alias)(&obj)})
}

// CustomerSetExternalIDAction implements the interface CustomerUpdateAction
type CustomerSetExternalIDAction struct {
	ExternalID string `json:"externalId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerSetExternalIDAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetExternalIDAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setExternalId", Alias: (*Alias)(&obj)})
}

// CustomerSetFirstNameAction implements the interface CustomerUpdateAction
type CustomerSetFirstNameAction struct {
	FirstName string `json:"firstName,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerSetFirstNameAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetFirstNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setFirstName", Alias: (*Alias)(&obj)})
}

// CustomerSetKeyAction implements the interface CustomerUpdateAction
type CustomerSetKeyAction struct {
	Key string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

// CustomerSetLastNameAction implements the interface CustomerUpdateAction
type CustomerSetLastNameAction struct {
	LastName string `json:"lastName,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerSetLastNameAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetLastNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLastName", Alias: (*Alias)(&obj)})
}

// CustomerSetLocaleAction implements the interface CustomerUpdateAction
type CustomerSetLocaleAction struct {
	Locale string `json:"locale,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerSetLocaleAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetLocaleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLocale", Alias: (*Alias)(&obj)})
}

// CustomerSetMiddleNameAction implements the interface CustomerUpdateAction
type CustomerSetMiddleNameAction struct {
	MiddleName string `json:"middleName,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerSetMiddleNameAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetMiddleNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMiddleName", Alias: (*Alias)(&obj)})
}

// CustomerSetSalutationAction implements the interface CustomerUpdateAction
type CustomerSetSalutationAction struct {
	Salutation string `json:"salutation,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerSetSalutationAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetSalutationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSalutation", Alias: (*Alias)(&obj)})
}

// CustomerSetTitleAction implements the interface CustomerUpdateAction
type CustomerSetTitleAction struct {
	Title string `json:"title,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerSetTitleAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetTitleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTitle", Alias: (*Alias)(&obj)})
}

// CustomerSetVatIDAction implements the interface CustomerUpdateAction
type CustomerSetVatIDAction struct {
	VatID string `json:"vatId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerSetVatIDAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetVatIDAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setVatId", Alias: (*Alias)(&obj)})
}

// CustomerSignInResult is a standalone struct
type CustomerSignInResult struct {
	Customer *Customer   `json:"customer"`
	Cart     interface{} `json:"cart,omitempty"`
}

// CustomerSignin is a standalone struct
type CustomerSignin struct {
	Password                string                  `json:"password"`
	Email                   string                  `json:"email"`
	AnonymousID             string                  `json:"anonymousId,omitempty"`
	AnonymousCartSignInMode AnonymousCartSignInMode `json:"anonymousCartSignInMode,omitempty"`
	AnonymousCartID         string                  `json:"anonymousCartId,omitempty"`
}

// CustomerToken is a standalone struct
type CustomerToken struct {
	Value          string    `json:"value"`
	LastModifiedAt time.Time `json:"lastModifiedAt,omitempty"`
	ID             string    `json:"id"`
	ExpiresAt      time.Time `json:"expiresAt"`
	CustomerID     string    `json:"customerId"`
	CreatedAt      time.Time `json:"createdAt"`
}

// CustomerUpdate is of type Update
type CustomerUpdate struct {
	Version int                    `json:"version"`
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
		obj.Actions[i] = mapDiscriminatorCustomerUpdateAction(obj.Actions[i])
	}

	return nil
}
