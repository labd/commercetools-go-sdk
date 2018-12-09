// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

type AnonymousCartSignInMode string

const (
	AnonymousCartSignInModeMergeWithExistingCustomerCart AnonymousCartSignInMode = "MergeWithExistingCustomerCart"
	AnonymousCartSignInModeUseAsNewActiveCustomerCart    AnonymousCartSignInMode = "UseAsNewActiveCustomerCart"
)

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

type CustomerAddAddressAction struct {
	Address *Address `json:"address"`
}

func (obj CustomerAddAddressAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addAddress", Alias: (*Alias)(&obj)})
}

type CustomerAddBillingAddressIdAction struct {
	AddressID string `json:"addressId"`
}

func (obj CustomerAddBillingAddressIdAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddBillingAddressIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addBillingAddressId", Alias: (*Alias)(&obj)})
}

type CustomerAddShippingAddressIdAction struct {
	AddressID string `json:"addressId"`
}

func (obj CustomerAddShippingAddressIdAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddShippingAddressIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addShippingAddressId", Alias: (*Alias)(&obj)})
}

type CustomerChangeAddressAction struct {
	AddressID string   `json:"addressId"`
	Address   *Address `json:"address"`
}

func (obj CustomerChangeAddressAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerChangeAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAddress", Alias: (*Alias)(&obj)})
}

type CustomerChangeEmailAction struct {
	Email string `json:"email"`
}

func (obj CustomerChangeEmailAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerChangeEmailAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeEmail", Alias: (*Alias)(&obj)})
}

type CustomerChangePassword struct {
	Version         int    `json:"version"`
	NewPassword     string `json:"newPassword"`
	ID              string `json:"id"`
	CurrentPassword string `json:"currentPassword"`
}

type CustomerCreateEmailToken struct {
	Version    int    `json:"version,omitempty"`
	TtlMinutes int    `json:"ttlMinutes"`
	ID         string `json:"id"`
}

type CustomerCreatePasswordResetToken struct {
	Email string `json:"email"`
}

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

type CustomerEmailVerify struct {
	Version    int    `json:"version,omitempty"`
	TokenValue string `json:"tokenValue"`
}

type CustomerPagedQueryResponse struct {
	Total   int        `json:"total,omitempty"`
	Offset  int        `json:"offset"`
	Count   int        `json:"count"`
	Results []Customer `json:"results"`
}

type CustomerReference struct {
	Key string    `json:"key,omitempty"`
	ID  string    `json:"id,omitempty"`
	Obj *Customer `json:"obj,omitempty"`
}

func (obj CustomerReference) MarshalJSON() ([]byte, error) {
	type Alias CustomerReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "customer", Alias: (*Alias)(&obj)})
}

type CustomerRemoveAddressAction struct {
	AddressID string `json:"addressId"`
}

func (obj CustomerRemoveAddressAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerRemoveAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeAddress", Alias: (*Alias)(&obj)})
}

type CustomerRemoveBillingAddressIdAction struct {
	AddressID string `json:"addressId"`
}

func (obj CustomerRemoveBillingAddressIdAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerRemoveBillingAddressIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeBillingAddressId", Alias: (*Alias)(&obj)})
}

type CustomerRemoveShippingAddressIdAction struct {
	AddressID string `json:"addressId"`
}

func (obj CustomerRemoveShippingAddressIdAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerRemoveShippingAddressIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeShippingAddressId", Alias: (*Alias)(&obj)})
}

type CustomerResetPassword struct {
	Version     int    `json:"version,omitempty"`
	TokenValue  string `json:"tokenValue"`
	NewPassword string `json:"newPassword"`
}

type CustomerSetCompanyNameAction struct {
	CompanyName string `json:"companyName,omitempty"`
}

func (obj CustomerSetCompanyNameAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetCompanyNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCompanyName", Alias: (*Alias)(&obj)})
}

type CustomerSetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

func (obj CustomerSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type CustomerSetCustomTypeAction struct {
	Type   *TypeReference  `json:"type,omitempty"`
	Fields *FieldContainer `json:"fields,omitempty"`
}

func (obj CustomerSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type CustomerSetCustomerGroupAction struct {
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
}

func (obj CustomerSetCustomerGroupAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetCustomerGroupAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomerGroup", Alias: (*Alias)(&obj)})
}

type CustomerSetCustomerNumberAction struct {
	CustomerNumber string `json:"customerNumber,omitempty"`
}

func (obj CustomerSetCustomerNumberAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetCustomerNumberAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomerNumber", Alias: (*Alias)(&obj)})
}

type CustomerSetDateOfBirthAction struct {
	DateOfBirth interface{} `json:"dateOfBirth,omitempty"`
}

func (obj CustomerSetDateOfBirthAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetDateOfBirthAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDateOfBirth", Alias: (*Alias)(&obj)})
}

type CustomerSetDefaultBillingAddressAction struct {
	AddressID string `json:"addressId,omitempty"`
}

func (obj CustomerSetDefaultBillingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetDefaultBillingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDefaultBillingAddress", Alias: (*Alias)(&obj)})
}

type CustomerSetDefaultShippingAddressAction struct {
	AddressID string `json:"addressId,omitempty"`
}

func (obj CustomerSetDefaultShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetDefaultShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDefaultShippingAddress", Alias: (*Alias)(&obj)})
}

type CustomerSetExternalIdAction struct {
	ExternalID string `json:"externalId,omitempty"`
}

func (obj CustomerSetExternalIdAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetExternalIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setExternalId", Alias: (*Alias)(&obj)})
}

type CustomerSetFirstNameAction struct {
	FirstName string `json:"firstName,omitempty"`
}

func (obj CustomerSetFirstNameAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetFirstNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setFirstName", Alias: (*Alias)(&obj)})
}

type CustomerSetKeyAction struct {
	Key string `json:"key,omitempty"`
}

func (obj CustomerSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type CustomerSetLastNameAction struct {
	LastName string `json:"lastName,omitempty"`
}

func (obj CustomerSetLastNameAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetLastNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLastName", Alias: (*Alias)(&obj)})
}

type CustomerSetLocaleAction struct {
	Locale string `json:"locale,omitempty"`
}

func (obj CustomerSetLocaleAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetLocaleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLocale", Alias: (*Alias)(&obj)})
}

type CustomerSetMiddleNameAction struct {
	MiddleName string `json:"middleName,omitempty"`
}

func (obj CustomerSetMiddleNameAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetMiddleNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMiddleName", Alias: (*Alias)(&obj)})
}

type CustomerSetSalutationAction struct {
	Salutation string `json:"salutation,omitempty"`
}

func (obj CustomerSetSalutationAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetSalutationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSalutation", Alias: (*Alias)(&obj)})
}

type CustomerSetTitleAction struct {
	Title string `json:"title,omitempty"`
}

func (obj CustomerSetTitleAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetTitleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTitle", Alias: (*Alias)(&obj)})
}

type CustomerSetVatIdAction struct {
	VatID string `json:"vatId,omitempty"`
}

func (obj CustomerSetVatIdAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetVatIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setVatId", Alias: (*Alias)(&obj)})
}

type CustomerSignInResult struct {
	Customer *Customer   `json:"customer"`
	Cart     interface{} `json:"cart,omitempty"`
}

type CustomerSignin struct {
	Password                string                  `json:"password"`
	Email                   string                  `json:"email"`
	AnonymousID             string                  `json:"anonymousId,omitempty"`
	AnonymousCartSignInMode AnonymousCartSignInMode `json:"anonymousCartSignInMode,omitempty"`
	AnonymousCartID         string                  `json:"anonymousCartId,omitempty"`
}

type CustomerToken struct {
	Value          string    `json:"value"`
	LastModifiedAt time.Time `json:"lastModifiedAt,omitempty"`
	ID             string    `json:"id"`
	ExpiresAt      time.Time `json:"expiresAt"`
	CustomerID     string    `json:"customerId"`
	CreatedAt      time.Time `json:"createdAt"`
}

type CustomerUpdate struct {
	Version int                    `json:"version"`
	Actions []CustomerUpdateAction `json:"actions"`
}

func (obj *CustomerUpdate) UnmarshalJSON(data []byte) error {
	type Alias CustomerUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		obj.Actions[i] = AbstractCustomerUpdateActionDiscriminatorMapping(obj.Actions[i])
	}

	return nil
}

type CustomerUpdateAction interface{}
type AbstractCustomerUpdateAction struct{}

func AbstractCustomerUpdateActionDiscriminatorMapping(input CustomerUpdateAction) CustomerUpdateAction {
	discriminator := input.(map[string]interface{})["action"]
	switch discriminator {
	case "addAddress":
		new := CustomerAddAddressAction{}
		mapstructure.Decode(input, &new)
		return new
	case "addBillingAddressId":
		new := CustomerAddBillingAddressIdAction{}
		mapstructure.Decode(input, &new)
		return new
	case "addShippingAddressId":
		new := CustomerAddShippingAddressIdAction{}
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
		new := CustomerRemoveBillingAddressIdAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeShippingAddressId":
		new := CustomerRemoveShippingAddressIdAction{}
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
		new := CustomerSetExternalIdAction{}
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
		new := CustomerSetVatIdAction{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}
