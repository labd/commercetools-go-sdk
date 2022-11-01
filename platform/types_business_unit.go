package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type Associate struct {
	// Roles the Associate holds within the Business Unit.
	Roles []AssociateRole `json:"roles"`
	// The [Customer](ctp:api:type:Customer) that is part of the Business Unit.
	Customer CustomerReference `json:"customer"`
}

type AssociateDraft struct {
	// Roles the Associate should hold within the Business Unit.
	Roles []AssociateRole `json:"roles"`
	// The [Customer](ctp:api:type:Customer) to be part of the Business Unit.
	Customer CustomerResourceIdentifier `json:"customer"`
}

/**
*	Roles defining how an [Associate](ctp:api:type:Associate) can interact with a Business Unit.
*
 */
type AssociateRole string

const (
	AssociateRoleAdmin AssociateRole = "Admin"
	AssociateRoleBuyer AssociateRole = "Buyer"
)

/**
*	Generic type to model those fields all Business Units have in common.
*
 */
type BusinessUnit interface{}

func mapDiscriminatorBusinessUnit(input interface{}) (BusinessUnit, error) {
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
		obj := Company{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "Division":
		obj := Division{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Generic draft type to model those fields all Business Units have in common.
*
 */
type BusinessUnitDraft interface{}

func mapDiscriminatorBusinessUnitDraft(input interface{}) (BusinessUnitDraft, error) {
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
		obj := CompanyDraft{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "Division":
		obj := DivisionDraft{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	[Reference](/../api/types#reference) to a [BusinessUnit](ctp:api:type:BusinessUnit) by its key.
*
 */
type BusinessUnitKeyReference struct {
	// Unique and immutable key of the referenced [BusinessUnit](ctp:api:type:BusinessUnit).
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitKeyReference) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "business-unit", Alias: (*Alias)(&obj)})
}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with results containing an array of [BusinessUnit](ctp:api:type:BusinessUnit).
*
 */
type BusinessUnitPagedQueryResponse struct {
	// Number of requested [results](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Number of elements [skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/../api/general-concepts#strong-consistency).
	// This field is returned by default.
	// For improved performance, calculating this field can be deactivated by using the query parameter `withTotal=false`.
	// When the results are filtered with a [Query Predicate](/../api/predicates/query), `total` is subject to a [limit](/../api/limits#queries).
	Total *int `json:"total,omitempty"`
	// [BusinessUnits](ctp:api:type:BusinessUnit) matching the query.
	Results []BusinessUnit `json:"results"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *BusinessUnitPagedQueryResponse) UnmarshalJSON(data []byte) error {
	type Alias BusinessUnitPagedQueryResponse
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Results {
		var err error
		obj.Results[i], err = mapDiscriminatorBusinessUnit(obj.Results[i])
		if err != nil {
			return err
		}
	}

	return nil
}

/**
*	[Reference](/../api/types#reference) to a [BusinessUnit](ctp:api:type:BusinessUnit).
*
 */
type BusinessUnitReference struct {
	// Unique identifier of the referenced [BusinessUnit](ctp:api:type:BusinessUnit).
	ID string `json:"id"`
	// Contains the representation of the expanded BusinessUnit. Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for BusinessUnit.
	Obj BusinessUnit `json:"obj,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *BusinessUnitReference) UnmarshalJSON(data []byte) error {
	type Alias BusinessUnitReference
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Obj != nil {
		var err error
		obj.Obj, err = mapDiscriminatorBusinessUnit(obj.Obj)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitReference) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "business-unit", Alias: (*Alias)(&obj)})
}

/**
*	[ResourceIdentifier](/../api/types#resourceidentifier) to a [BusinessUnit](ctp:api:type:BusinessUnit).
*
 */
type BusinessUnitResourceIdentifier struct {
	// Unique identifier of the referenced [BusinessUnit](ctp:api:type:BusinessUnit). Either `id` or `key` is required.
	ID *string `json:"id,omitempty"`
	// Unique key of the referenced [BusinessUnit](ctp:api:type:BusinessUnit). Either `id` or `key` is required.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "business-unit", Alias: (*Alias)(&obj)})
}

/**
*	Indicates whether the Business Unit can be edited and used in [Orders](/../api/projects/orders), [Carts](/../api/projects/carts), or [Quotes](/../api/projects/quotes).
*
 */
type BusinessUnitStatus string

const (
	BusinessUnitStatusActive   BusinessUnitStatus = "Active"
	BusinessUnitStatusInactive BusinessUnitStatus = "Inactive"
)

/**
*	Defines whether the Stores of the Business Unit are set directly on the Business Unit or are inherited from its parent unit.
*
 */
type BusinessUnitStoreMode string

const (
	BusinessUnitStoreModeExplicit   BusinessUnitStoreMode = "Explicit"
	BusinessUnitStoreModeFromParent BusinessUnitStoreMode = "FromParent"
)

/**
*	The type of the Business Unit indicating its position in a hierarchy.
*
 */
type BusinessUnitType string

const (
	BusinessUnitTypeCompany  BusinessUnitType = "Company"
	BusinessUnitTypeDivision BusinessUnitType = "Division"
)

type BusinessUnitUpdate struct {
	// Expected version of the BusinessUnit on which the changes should be applied.
	// If the expected version does not match the actual version, a [409 Conflict](/../api/errors#409-conflict) error will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the BusinessUnit.
	Actions []BusinessUnitUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *BusinessUnitUpdate) UnmarshalJSON(data []byte) error {
	type Alias BusinessUnitUpdate
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

type BusinessUnitUpdateAction interface{}

func mapDiscriminatorBusinessUnitUpdateAction(input interface{}) (BusinessUnitUpdateAction, error) {
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
		obj := BusinessUnitAddAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addAssociate":
		obj := BusinessUnitAddAssociateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addBillingAddressId":
		obj := BusinessUnitAddBillingAddressIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addShippingAddressId":
		obj := BusinessUnitAddShippingAddressIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addStore":
		obj := BusinessUnitAddStoreAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeAddress":
		obj := BusinessUnitChangeAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeAssociate":
		obj := BusinessUnitChangeAssociateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeName":
		obj := BusinessUnitChangeNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeParentUnit":
		obj := BusinessUnitChangeParentUnitAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeStatus":
		obj := BusinessUnitChangeStatusAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeAddress":
		obj := BusinessUnitRemoveAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeAssociate":
		obj := BusinessUnitRemoveAssociateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeBillingAddressId":
		obj := BusinessUnitRemoveBillingAddressIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeShippingAddressId":
		obj := BusinessUnitRemoveShippingAddressIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeStore":
		obj := BusinessUnitRemoveStoreAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAddressCustomField":
		obj := BusinessUnitSetAddressCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAddressCustomType":
		obj := BusinessUnitSetAddressCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAssociates":
		obj := BusinessUnitSetAssociatesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setContactEmail":
		obj := BusinessUnitSetContactEmailAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomField":
		obj := BusinessUnitSetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomType":
		obj := BusinessUnitSetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDefaultBillingAddress":
		obj := BusinessUnitSetDefaultBillingAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDefaultShippingAddress":
		obj := BusinessUnitSetDefaultShippingAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setStoreMode":
		obj := BusinessUnitSetStoreModeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setStores":
		obj := BusinessUnitSetStoresAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Business Unit type to represent the top level of a business.
*	Contains specific fields and values that differentiate a Company from the generic [BusinessUnit](ctp:api:type:BusinessUnit).
*
 */
type Company struct {
	// Unique identifier of the Business Unit.
	ID string `json:"id"`
	// Current version of the Business Unit.
	Version int `json:"version"`
	// Date and time (UTC) the Business Unit was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the Business Unit was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources updated after 1 February 2019 except for [events not tracked](/../api/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/../api/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// User-defined unique identifier of the Business Unit.
	Key string `json:"key"`
	// Indicates whether the Business Unit can be edited and used in [Orders](/../api/projects/orders).
	Status BusinessUnitStatus `json:"status"`
	// References to [Stores](ctp:api:type:Store) the Business Unit is associated with.
	// If empty, the Business Unit can only create [Carts](ctp:api:type:Cart), [Orders](ctp:api:type:Order), or [Quotes](/../api/quotes-overview) that have no `store` value.
	// If not empty, the Business Unit can only be linked to [Carts](ctp:api:type:Cart) and [Orders](ctp:api:type:Order) of a referenced Store.
	// Only present when `storeMode` is `Explicit`.
	Stores []StoreKeyReference `json:"stores"`
	// Is always `Explicit` since a Company does not have a parent Business Unit the Stores can be inherited from.
	StoreMode BusinessUnitStoreMode `json:"storeMode"`
	// Name of the Business Unit.
	Name string `json:"name"`
	// Email address of the Business Unit.
	ContactEmail *string `json:"contactEmail,omitempty"`
	// Custom Fields for the Business Unit.
	Custom *CustomFields `json:"custom,omitempty"`
	// Addresses used by the Business Unit.
	Addresses []Address `json:"addresses"`
	// Unique identifiers of addresses used as shipping addresses.
	ShippingAddressIds []string `json:"shippingAddressIds"`
	// Unique identifier of the address used as the default shipping address.
	DefaultShipingAddressId *string `json:"defaultShipingAddressId,omitempty"`
	// Unique identifiers of addresses used as billing addresses.
	BillingAddressIds []string `json:"billingAddressIds"`
	// Unique identifier of the address used as the default billing address.
	DefaultBillingAddressId *string `json:"defaultBillingAddressId,omitempty"`
	// Members that are part of the Business Unit in specific [roles](ctp:api:type:AssociateRole).
	Associates []Associate `json:"associates"`
	// Parent unit of the Business Unit. Only present when the `unitType` is `Division`.
	ParentUnit *BusinessUnitKeyReference `json:"parentUnit,omitempty"`
	// Top-level unit of the Business Unit. The top-level unit is of `unitType` `Company`.
	TopLevelUnit BusinessUnitKeyReference `json:"topLevelUnit"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj Company) MarshalJSON() ([]byte, error) {
	type Alias Company
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

	if raw["stores"] == nil {
		delete(raw, "stores")
	}

	if raw["shippingAddressIds"] == nil {
		delete(raw, "shippingAddressIds")
	}

	if raw["billingAddressIds"] == nil {
		delete(raw, "billingAddressIds")
	}

	return json.Marshal(raw)

}

/**
*	Draft type to represent the top level of a business. Contains the fields and values of the generic [BusinessUnitDraft](ctp:api:type:BusinessUnitDraft] that are used specifically for creating a [Company](ctp:api:type:Company).
*
 */
type CompanyDraft struct {
	// User-defined unique identifier for the Business Unit.
	Key string `json:"key"`
	// Indicates whether the Business Unit can be edited and used in [Orders](/../api/projects/orders).
	Status *BusinessUnitStatus `json:"status,omitempty"`
	// References to [Stores](ctp:api:type:Store) the Business Unit is associated with. Can only be set when `storeMode` is `Explicit`.
	// If not empty, the Business Unit can only be linked to [Carts](ctp:api:type:Cart) and [Orders](ctp:api:type:Order) of a referenced Store.
	// If empty, the Business Unit can only create [Carts](ctp:api:type:Cart), [Orders](ctp:api:type:Order), or [Quotes](/../api/quotes-overview) that have no `store` value.
	// Defaults to empty for [Companies](ctp:api:type:BusinessUnitType) and not set for [Divisions](ctp:api:type:BusinessUnitType).
	Stores []StoreKeyReference `json:"stores"`
	// Defines whether the Stores of the Business Unit are set on the Business Unit or are inherited from a parent.
	// Defaults to `Explicit` for [Companies](ctp:api:type:BusinessUnitType) and to `FromParent` for [Divisions](ctp:api:type:BusinessUnitType).
	StoreMode *BusinessUnitStoreMode `json:"storeMode,omitempty"`
	// Name of the Business Unit.
	Name string `json:"name"`
	// Email address of the Business Unit.
	ContactEmail *string `json:"contactEmail,omitempty"`
	// List of members that are part of the Business Unit in specific [roles](ctp:api:type:AssociateRole).
	Associates []AssociateDraft `json:"associates"`
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
	// Custom Fields for the Business Unit.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CompanyDraft) MarshalJSON() ([]byte, error) {
	type Alias CompanyDraft
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

	if raw["stores"] == nil {
		delete(raw, "stores")
	}

	if raw["associates"] == nil {
		delete(raw, "associates")
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

/**
*	Business Unit type to model divisions that are part of the [Company](ctp:api:type:Company) or a higher order Division.
*	Contains specific fields and values that differentiate a Division from the generic [BusinessUnit](ctp:api:type:BusinessUnit).
*
 */
type Division struct {
	// Unique identifier of the Business Unit.
	ID string `json:"id"`
	// Current version of the Business Unit.
	Version int `json:"version"`
	// Date and time (UTC) the Business Unit was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the Business Unit was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources updated after 1 February 2019 except for [events not tracked](/../api/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/../api/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// User-defined unique identifier of the Business Unit.
	Key string `json:"key"`
	// Indicates whether the Business Unit can be edited and used in [Orders](/../api/projects/orders).
	Status BusinessUnitStatus `json:"status"`
	// References to [Stores](ctp:api:type:Store) the Business Unit is associated with.
	// If empty, the Business Unit can only create [Carts](ctp:api:type:Cart), [Orders](ctp:api:type:Order), or [Quotes](/../api/quotes-overview) that have no `store` value.
	// If not empty, the Business Unit can only be linked to [Carts](ctp:api:type:Cart) and [Orders](ctp:api:type:Order) of a referenced Store.
	// Only present when `storeMode` is `Explicit`.
	Stores []StoreKeyReference `json:"stores"`
	// Defines whether the Stores of the Division are set explicitly or inherited from a parent Business Unit.
	StoreMode BusinessUnitStoreMode `json:"storeMode"`
	// Name of the Business Unit.
	Name string `json:"name"`
	// Email address of the Business Unit.
	ContactEmail *string `json:"contactEmail,omitempty"`
	// Custom Fields for the Business Unit.
	Custom *CustomFields `json:"custom,omitempty"`
	// Addresses used by the Business Unit.
	Addresses []Address `json:"addresses"`
	// Unique identifiers of addresses used as shipping addresses.
	ShippingAddressIds []string `json:"shippingAddressIds"`
	// Unique identifier of the address used as the default shipping address.
	DefaultShipingAddressId *string `json:"defaultShipingAddressId,omitempty"`
	// Unique identifiers of addresses used as billing addresses.
	BillingAddressIds []string `json:"billingAddressIds"`
	// Unique identifier of the address used as the default billing address.
	DefaultBillingAddressId *string `json:"defaultBillingAddressId,omitempty"`
	// Members that are part of the Business Unit in specific [roles](ctp:api:type:AssociateRole).
	Associates []Associate `json:"associates"`
	// Parent unit of the Division.
	ParentUnit BusinessUnitKeyReference `json:"parentUnit"`
	// Top-level unit of the Business Unit. The top-level unit is of `unitType` `Company`.
	TopLevelUnit BusinessUnitKeyReference `json:"topLevelUnit"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj Division) MarshalJSON() ([]byte, error) {
	type Alias Division
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

	if raw["stores"] == nil {
		delete(raw, "stores")
	}

	if raw["shippingAddressIds"] == nil {
		delete(raw, "shippingAddressIds")
	}

	if raw["billingAddressIds"] == nil {
		delete(raw, "billingAddressIds")
	}

	return json.Marshal(raw)

}

/**
*	Draft type to model divisions that are part of a [Company](ctp:api:type:Company) or a higher order [Division](ctp:api:type:Division).
*	Contains the fields and values of the generic [BusinessUnitDraft](ctp:api:type:BusinessUnitDraft) that are used specifically for creating a Division.
*
 */
type DivisionDraft struct {
	// User-defined unique identifier for the Business Unit.
	Key string `json:"key"`
	// Indicates whether the Business Unit can be edited and used in [Orders](/../api/projects/orders).
	Status *BusinessUnitStatus `json:"status,omitempty"`
	// References to [Stores](ctp:api:type:Store) the Business Unit is associated with. Can only be set when `storeMode` is `Explicit`.
	// If not empty, the Business Unit can only be linked to [Carts](ctp:api:type:Cart) and [Orders](ctp:api:type:Order) of a referenced Store.
	// If empty, the Business Unit can only create [Carts](ctp:api:type:Cart), [Orders](ctp:api:type:Order), or [Quotes](/../api/quotes-overview) that have no `store` value.
	// Defaults to empty for [Companies](ctp:api:type:BusinessUnitType) and not set for [Divisions](ctp:api:type:BusinessUnitType).
	Stores []StoreKeyReference `json:"stores"`
	// If not set, the Division inherits the [Stores](ctp:api:type:Store) from its `parentUnit`.
	// Set this to `Explicit` if you want to set the Stores explicitly in the `stores` field instead.
	StoreMode *BusinessUnitStoreMode `json:"storeMode,omitempty"`
	// Name of the Business Unit.
	Name string `json:"name"`
	// Email address of the Business Unit.
	ContactEmail *string `json:"contactEmail,omitempty"`
	// List of members that are part of the Business Unit in specific [roles](ctp:api:type:AssociateRole).
	Associates []AssociateDraft `json:"associates"`
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
	// Custom Fields for the Business Unit.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
	// The parent unit of this Division. Can be a Company or a Division.
	ParentUnit BusinessUnitResourceIdentifier `json:"parentUnit"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DivisionDraft) MarshalJSON() ([]byte, error) {
	type Alias DivisionDraft
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

	if raw["stores"] == nil {
		delete(raw, "stores")
	}

	if raw["associates"] == nil {
		delete(raw, "associates")
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

/**
*	Adding an address to a [Business Unit](ctp:api:type:BusinessUnit) generates a [BusinessUnitAddressAdded](ctp:api:type:BusinessUnitAddressAddedMessage) Message.
*
 */
type BusinessUnitAddAddressAction struct {
	// Address to add to the addresses of the [Business Unit](ctp:api:type:BusinessUnit).
	Address BaseAddress `json:"address"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitAddAddressAction) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitAddAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addAddress", Alias: (*Alias)(&obj)})
}

/**
*	Adding an Associate to a [Business Unit](ctp:api:type:BusinessUnit) generates a [BusinessUnitAssociateAdded](ctp:api:type:BusinessUnitAssociateAddedMessage) Message.
*
 */
type BusinessUnitAddAssociateAction struct {
	// The Associate to add.
	Associate AssociateDraft `json:"associate"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitAddAssociateAction) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitAddAssociateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addAssociate", Alias: (*Alias)(&obj)})
}

/**
*	Adding a billing address to a [Business Unit](ctp:api:type:BusinessUnit) generates a [BusinessUnitBillingAddressAdded](ctp:api:type:BusinessUnitBillingAddressAddedMessage) Message.
*
 */
type BusinessUnitAddBillingAddressIdAction struct {
	// ID of the address to add as a billing address. Either `addressId` or `addressKey` is required.
	AddressId *string `json:"addressId,omitempty"`
	// Key of the address to add as a billing address. Either `addressId` or `addressKey` is required.
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitAddBillingAddressIdAction) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitAddBillingAddressIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addBillingAddressId", Alias: (*Alias)(&obj)})
}

/**
*	Adding a shipping address to a [Business Unit](ctp:api:type:BusinessUnit) generates a [BusinessUnitShippingAddressAdded](ctp:api:type:BusinessUnitShippingAddressAddedMessage) Message.
*
 */
type BusinessUnitAddShippingAddressIdAction struct {
	// ID of the address to add as a shipping address. Either `addressId` or `addressKey` is required.
	AddressId *string `json:"addressId,omitempty"`
	// Key of the address to add as a shipping address. Either `addressId` or `addressKey` is required.
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitAddShippingAddressIdAction) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitAddShippingAddressIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addShippingAddressId", Alias: (*Alias)(&obj)})
}

/**
*	Adding a [Store](ctp:api:type:Store) to a [Business Unit](ctp:api:type:BusinessUnit) generates a [BusinessUnitStoreAdded](ctp:api:type:BusinessUnitStoreAddedMessage) Message.
*	Only applicable when `storeMode` is `Explicit`.
*
 */
type BusinessUnitAddStoreAction struct {
	// [Store](ctp:api:type:Store) to add.
	Store StoreResourceIdentifier `json:"store"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitAddStoreAction) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitAddStoreAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addStore", Alias: (*Alias)(&obj)})
}

/**
*	Changing the address on a Business Unit generates the [BusinessUnitAddressChanged](ctp:api:type:BusinessUnitAddressChangedMessage) Message.
*
 */
type BusinessUnitChangeAddressAction struct {
	// ID of the address to change. Either `addressId` or `addressKey` is required.
	AddressId *string `json:"addressId,omitempty"`
	// Key of the address to change. Either `addressId` or `addressKey` is required.
	AddressKey *string `json:"addressKey,omitempty"`
	// New address to set.
	Address BaseAddress `json:"address"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitChangeAddressAction) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitChangeAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAddress", Alias: (*Alias)(&obj)})
}

/**
*	Updating the [Associate](ctp:api:type:Associate) on a [Business Unit](ctp:api:type:BusinessUnit) generates the [BusinessUnitAssociateChanged](ctp:api:type:BusinessUnitAssociateChangedMessage) Message.
*
 */
type BusinessUnitChangeAssociateAction struct {
	// New version of an existing Associate.
	Associate AssociateDraft `json:"associate"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitChangeAssociateAction) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitChangeAssociateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAssociate", Alias: (*Alias)(&obj)})
}

/**
*	Updating the name on a [Business Unit](ctp:api:type:BusinessUnit) generates a [BusinessUnitNameChanged](ctp:api:type:BusinessUnitNameChangedMessage) Message.
*
 */
type BusinessUnitChangeNameAction struct {
	// New name to set.
	Name string `json:"name"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

/**
*	Changing the parent of a [Business Unit](ctp:api:type:BusinessUnit) generates a [BusinessUnitParentUnitChanged](ctp:api:type:BusinessUnitParentUnitChangedMessage) Message.
*
 */
type BusinessUnitChangeParentUnitAction struct {
	// New parent unit of the [Business Unit](ctp:api:type:BusinessUnit).
	ParentUnit BusinessUnitResourceIdentifier `json:"parentUnit"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitChangeParentUnitAction) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitChangeParentUnitAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeParentUnit", Alias: (*Alias)(&obj)})
}

/**
*	Changing the status of a [Business Unit](ctp:api:type:BusinessUnit) generates a [BusinessUnitStatusChanged](ctp:api:type:BusinessUnitStatusChangedMessage) Message.
*
 */
type BusinessUnitChangeStatusAction struct {
	// New status to set.
	Status string `json:"status"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitChangeStatusAction) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitChangeStatusAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeStatus", Alias: (*Alias)(&obj)})
}

/**
*	Removing the address from a [Business Unit](ctp:api:type:BusinessUnit) generates the [BusinessUnitAddressRemoved](ctp:api:type:BusinessUnitAddressRemovedMessage) Message.
*
 */
type BusinessUnitRemoveAddressAction struct {
	// ID of the address to be removed. Either `addressId` or `addressKey` is required.
	AddressId *string `json:"addressId,omitempty"`
	// Key of the address to be removed. Either `addressId` or `addressKey` is required.
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitRemoveAddressAction) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitRemoveAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeAddress", Alias: (*Alias)(&obj)})
}

/**
*	Removing an [Associate](ctp:api:type:Associate) from a [Business Unit](ctp:api:type:BusinessUnit) generates a [BusinessUnitAssociateRemoved](ctp:api:type:BusinessUnitAssociateRemovedMessage) Message.
*
 */
type BusinessUnitRemoveAssociateAction struct {
	// [Associate](ctp:api:type:Associate) to remove.
	Customer CustomerResourceIdentifier `json:"customer"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitRemoveAssociateAction) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitRemoveAssociateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeAssociate", Alias: (*Alias)(&obj)})
}

/**
*	Removing a billing address from a [Business Unit](ctp:api:type:BusinessUnit) generates a [BusinessUnitBillingAddressRemoved](ctp:api:type:BusinessUnitBillingAddressRemovedMessage) Message.
*
 */
type BusinessUnitRemoveBillingAddressIdAction struct {
	// ID of the address to be removed from `billingAddressIds`. Either `addressId` or `addressKey` is required.
	AddressId *string `json:"addressId,omitempty"`
	// Key of the address to be removed from `billingAddressIds`. Either `addressId` or `addressKey` is required.
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitRemoveBillingAddressIdAction) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitRemoveBillingAddressIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeBillingAddressId", Alias: (*Alias)(&obj)})
}

/**
*	Removing a shipping address from a [Business Unit](ctp:api:type:BusinessUnit) generates a [BusinessUnitShippingAddressRemoved](ctp:api:type:BusinessUnitShippingAddressRemovedMessage) Message.
*
 */
type BusinessUnitRemoveShippingAddressIdAction struct {
	// ID of the address to be removed from `shippingAddressIds`. Either `addressId` or `addressKey` is required.
	AddressId *string `json:"addressId,omitempty"`
	// Key of the address to be removed from `shippingAddressIds`. Either `addressId` or `addressKey` is required.
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitRemoveShippingAddressIdAction) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitRemoveShippingAddressIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeShippingAddressId", Alias: (*Alias)(&obj)})
}

/**
*	Removes a Store from the Business Unit.
*	Newly created [Carts](ctp:api:type:Cart) and [Orders](ctp:api:type:Order) can no longer reference the removed Store for the Business Unit.
*	We recommend cleaning up unordered Carts that still have the Store assigned after calling this update action since those cannot be converted to Orders.
*	Orders created prior to when the Store was removed remain unchanged.
*	Generates a [BusinessUnitStoreRemoved](ctp:api:type:BusinessUnitStoreRemovedMessage) Message.
*	Only applicable when `storeMode` is `Explicit`.
*
 */
type BusinessUnitRemoveStoreAction struct {
	// [Store](ctp:api:type:Store) to remove.
	Store StoreResourceIdentifier `json:"store"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitRemoveStoreAction) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitRemoveStoreAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeStore", Alias: (*Alias)(&obj)})
}

type BusinessUnitSetAddressCustomFieldAction struct {
	// ID of the address to be extended.
	AddressId string `json:"addressId"`
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Trying to remove a field that does not exist will fail with an [InvalidOperation](/../api/errors#general-400-invalid-operation) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitSetAddressCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitSetAddressCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAddressCustomField", Alias: (*Alias)(&obj)})
}

type BusinessUnitSetAddressCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the `address` with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the `address`.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) for the `address`.
	Fields *FieldContainer `json:"fields,omitempty"`
	// ID of the address to be extended.
	AddressId string `json:"addressId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitSetAddressCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitSetAddressCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAddressCustomType", Alias: (*Alias)(&obj)})
}

/**
*	Changes the Associates of a [Business Unit](ctp:api:type:BusinessUnit), generates a [BusinessUnitAssociatesSet](ctp:api:type:BusinessUnitAssociatesSetMessage) Message.
*
 */
type BusinessUnitSetAssociatesAction struct {
	// The new list of Associates. If not provided, any existing list is removed.
	Associates []AssociateDraft `json:"associates"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitSetAssociatesAction) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitSetAssociatesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssociates", Alias: (*Alias)(&obj)})
}

/**
*	Setting the contact email on a [Business Unit](ctp:api:type:BusinessUnit) generates a [BusinessUnitContactEmailSet](ctp:api:type:BusinessUnitContactEmailSetMessage) Message.
*
 */
type BusinessUnitSetContactEmailAction struct {
	// Email to set.
	// If `contactEmail` is absent or `null`, the existing contact email, if any, will be removed.
	ContactEmail *string `json:"contactEmail,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitSetContactEmailAction) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitSetContactEmailAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setContactEmail", Alias: (*Alias)(&obj)})
}

type BusinessUnitSetCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Trying to remove a field that does not exist will fail with an [InvalidOperation](/../api/errors#general-400-invalid-operation) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type BusinessUnitSetCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the BusinessUnit with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the BusinessUnit.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) for the BusinessUnit.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

/**
*	Setting the default billing address on a [Business Unit](ctp:api:type:BusinessUnit) generates the [BusinessUnitDefaultBillingAddressSet](ctp:api:type:BusinessUnitDefaultBillingAddressSetMessage) Message.
*
 */
type BusinessUnitSetDefaultBillingAddressAction struct {
	// ID of the address to add as a billing address. Either `addressId` or `addressKey` is required.
	AddressId *string `json:"addressId,omitempty"`
	// Key of the address to add as a billing address. Either `addressId` or `addressKey` is required.
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitSetDefaultBillingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitSetDefaultBillingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDefaultBillingAddress", Alias: (*Alias)(&obj)})
}

/**
*	Setting the default shipping address on a [Business Unit](ctp:api:type:BusinessUnit) generates a [BusinessUnitDefaultShippingAddressSet](ctp:api:type:BusinessUnitDefaultShippingAddressSetMessage) Message.
*
 */
type BusinessUnitSetDefaultShippingAddressAction struct {
	// ID of the address to add as a shipping address. Either `addressId` or `addressKey` is required.
	AddressId *string `json:"addressId,omitempty"`
	// Key of the address to add as a shipping address. Either `addressId` or `addressKey` is required.
	AddressKey *string `json:"addressKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitSetDefaultShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitSetDefaultShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDefaultShippingAddress", Alias: (*Alias)(&obj)})
}

/**
*	If `storeMode` is changed to `FromParent`, current Stores defined for the Business Unit are removed.
*	Only Business Units of type `Division` can be set to `FromParent`.
*	This update action generates a [BusinessUnitStoreModeChanged](ctp:api:type:BusinessUnitStoreModeChangedMessage) Message.
*
 */
type BusinessUnitSetStoreModeAction struct {
	// Set to `Explicit` to specify Stores for the Business Unit. Set to `FromParent` to inherit Stores from a parent.
	StoreMode BusinessUnitStoreMode `json:"storeMode"`
	// Set the [Stores](ctp:api:type:Store) the Business Unit is associated with. Can only be set if `storeMode` is `Explicit`.
	Stores []StoreResourceIdentifier `json:"stores"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitSetStoreModeAction) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitSetStoreModeAction
	data, err := json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setStoreMode", Alias: (*Alias)(&obj)})
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
*	Sets the Stores of the Business Unit. Can only be set if the Business Unit `storeMode` is `Explicit`.
*	[Carts](ctp:api:type:Cart) and [Orders](ctp:api:type:Order) created after the Set Stores update must use the new Stores of
*	the Business Unit and, if set, their [Product Selections](ctp:api:type:ProductSelection), and [Channels](ctp:api:type:Channel).
*	Orders created prior to the Set Stores update action are unchanged.
*	Setting the Stores on a Business Unit generates a [BusinessUnitStoresSet](ctp:api:type:BusinessUnitStoresSetMessage) Message.
*
 */
type BusinessUnitSetStoresAction struct {
	// [Stores](ctp:api:type:Store) to set. Overrides the current list of Stores.
	Stores []StoreResourceIdentifier `json:"stores"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitSetStoresAction) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitSetStoresAction
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
