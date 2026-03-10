package importapi

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
)

/**
*	Determines whether the AssociateRoleAssignment can be inherited by child Business Units.
*
 */
type AssociateRoleInheritanceMode string

const (
	AssociateRoleInheritanceModeEnabled  AssociateRoleInheritanceMode = "Enabled"
	AssociateRoleInheritanceModeDisabled AssociateRoleInheritanceMode = "Disabled"
)

/**
*	Indicates whether the Business Unit can be used.
 */
type BusinessUnitStatus string

const (
	BusinessUnitStatusActive   BusinessUnitStatus = "Active"
	BusinessUnitStatusInactive BusinessUnitStatus = "Inactive"
)

type BusinessUnitAssociateMode string

const (
	BusinessUnitAssociateModeExplicit              BusinessUnitAssociateMode = "Explicit"
	BusinessUnitAssociateModeExplicitAndFromParent BusinessUnitAssociateMode = "ExplicitAndFromParent"
)

/**
*	Determines whether a Business Unit can inherit [Approval Rules](/projects/approval-rules) from a parent. Only Business Units of type `Division` can use `ExplicitAndFromParent`.
*
 */
type BusinessUnitApprovalRuleMode string

const (
	BusinessUnitApprovalRuleModeExplicit              BusinessUnitApprovalRuleMode = "Explicit"
	BusinessUnitApprovalRuleModeExplicitAndFromParent BusinessUnitApprovalRuleMode = "ExplicitAndFromParent"
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
*	Defines the type of the Business Unit.
 */
type BusinessUnitType string

const (
	BusinessUnitTypeCompany  BusinessUnitType = "Company"
	BusinessUnitTypeDivision BusinessUnitType = "Division"
)

/**
*	The role of an Associate in a Business Unit.
 */
type AssociateRoleAssignmentDraft struct {
	// The role to assign to the Associate.
	AssociateRole AssociateRoleKeyReference `json:"associateRole"`
	// Determines whether the AssociateRole is inherited. If `Disabled`, the AssociateRole is not inherited from a parent Business Unit.
	Inheritance AssociateRoleInheritanceMode `json:"inheritance"`
}

/**
*	Draft for an Associate in a Business Unit.
 */
type AssociateDraft struct {
	// The Customer to be part of the Business Unit.
	Customer CustomerKeyReference `json:"customer"`
	// The roles to assign to the Associate.
	AssociateRoleAssignments []AssociateRoleAssignmentDraft `json:"associateRoleAssignments"`
}

/**
*	Represents the data used to import a BusinessUnit. Can be of type [Company](ctp:api:type:Company) or [Division](ctp:api:type:Division).
*
 */
type BusinessUnitImport interface{}

func mapDiscriminatorBusinessUnitImport(input interface{}) (BusinessUnitImport, error) {
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
		obj := CompanyBusinessUnitImport{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "Division":
		obj := DivisionBusinessUnitImport{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Represents a [Company](ctp:api:type:Company), the top-level of a business.
 */
type CompanyBusinessUnitImport struct {
	// User-defined unique identifier. If a [BusinessUnit](ctp:api:type:BusinessUnit) with this `key` exists, it is updated with the imported data.
	Key string `json:"key"`
	// The name of the Business Unit.
	Name string `json:"name"`
	// The status of the Business Unit.
	Status *BusinessUnitStatus `json:"status,omitempty"`
	// The contact email address for the Business Unit.
	ContactEmail *string `json:"contactEmail,omitempty"`
	// List of Associates to be assigned to the Business Unit.
	Associates []AssociateDraft `json:"associates"`
	// The addresses for the Business Unit.
	Addresses []Address `json:"addresses"`
	// The indices of the shipping addresses in the `addresses` array.
	ShippingAddresses []int `json:"shippingAddresses"`
	// The index of the default shipping address in the `addresses` array.
	DefaultShippingAddress *int `json:"defaultShippingAddress,omitempty"`
	// The indices of the billing addresses in the `addresses` array.
	BillingAddresses []int `json:"billingAddresses"`
	// The index of the default billing address in the `addresses` array.
	DefaultBillingAddress *int `json:"defaultBillingAddress,omitempty"`
	// The Stores of the Business Unit.
	Stores []StoreKeyReference `json:"stores"`
	// Custom fields for the Business Unit.
	Custom    *Custom                `json:"custom,omitempty"`
	StoreMode *BusinessUnitStoreMode `json:"storeMode,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CompanyBusinessUnitImport) MarshalJSON() ([]byte, error) {
	type Alias CompanyBusinessUnitImport
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

	if raw["stores"] == nil {
		delete(raw, "stores")
	}

	return json.Marshal(raw)

}

/**
*	Represents a [Division](ctp:api:type:Division), a sub-unit of a [Company](ctp:api:type:Company) or another Division.
 */
type DivisionBusinessUnitImport struct {
	// User-defined unique identifier. If a [BusinessUnit](ctp:api:type:BusinessUnit) with this `key` exists, it is updated with the imported data.
	Key string `json:"key"`
	// The name of the Business Unit.
	Name string `json:"name"`
	// The status of the Business Unit.
	Status *BusinessUnitStatus `json:"status,omitempty"`
	// The contact email address for the Business Unit.
	ContactEmail *string `json:"contactEmail,omitempty"`
	// List of Associates to be assigned to the Business Unit.
	Associates []AssociateDraft `json:"associates"`
	// The addresses for the Business Unit.
	Addresses []Address `json:"addresses"`
	// The indices of the shipping addresses in the `addresses` array.
	ShippingAddresses []int `json:"shippingAddresses"`
	// The index of the default shipping address in the `addresses` array.
	DefaultShippingAddress *int `json:"defaultShippingAddress,omitempty"`
	// The indices of the billing addresses in the `addresses` array.
	BillingAddresses []int `json:"billingAddresses"`
	// The index of the default billing address in the `addresses` array.
	DefaultBillingAddress *int `json:"defaultBillingAddress,omitempty"`
	// The Stores of the Business Unit.
	Stores []StoreKeyReference `json:"stores"`
	// Custom fields for the Business Unit.
	Custom *Custom `json:"custom,omitempty"`
	// If `Explicit`, the `stores` field cannot be empty and the Business Unit is explicitly associated with the given Stores. If `FromParent`, the Business Unit inherits the Stores from its parent.
	StoreMode *BusinessUnitStoreMode `json:"storeMode,omitempty"`
	// The parent Business Unit of this Division.
	ParentUnit BusinessUnitKeyReference `json:"parentUnit"`
	// If `Explicit`, Associates are not inherited from the parent. If `ExplicitAndFromParent`, Associates are inherited from the parent.
	AssociateMode *BusinessUnitAssociateMode `json:"associateMode,omitempty"`
	// If `Explicit`, approval rules are not inherited from the parent. If `ExplicitAndFromParent`, approval rules are inherited from the parent.
	ApprovalRuleMode *BusinessUnitApprovalRuleMode `json:"approvalRuleMode,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DivisionBusinessUnitImport) MarshalJSON() ([]byte, error) {
	type Alias DivisionBusinessUnitImport
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

	if raw["stores"] == nil {
		delete(raw, "stores")
	}

	return json.Marshal(raw)

}
