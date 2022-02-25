package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type AssignedProductReference struct {
	// Reference to a Product that is assigned to the Product Selection.
	Product ProductReference `json:"product"`
}

type AssignedProductSelection struct {
	// Reference to the Product Selection that this assignment is part of.
	ProductSelection ProductSelectionReference `json:"productSelection"`
}

/**
*	[PagedQueryResult](/general-concepts#pagedqueryresult) containing an array of [AssignedProductSelection](ctp:api:type:AssignedProductSelection).
*
 */
type AssignedProductSelectionPagedQueryResponse struct {
	// Number of results requested in the query request.
	Limit int `json:"limit"`
	// Offset supplied by the client or the server default.
	// It is the number of elements skipped, not a page number.
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/general-concepts#strong-consistency).
	// Unlike other endpoints, the Product Selection endpoint does not return this field by default.
	// To get `total`, pass the query parameter `withTotal` set to `true`.
	// When the results are filtered with a [Query Predicate](/predicates/query), `total` is subject to a [limit](/limits#queries).
	Total *int `json:"total,omitempty"`
	// References to Product Selection that are assigned to the Product.
	Results []AssignedProductSelection `json:"results"`
}

type ProductSelection struct {
	// Unique ID of the Product Selection.
	ID string `json:"id"`
	// Current version of the Product Selection.
	Version int `json:"version"`
	// Date and time (UTC) the Product Selection was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the Product Selection was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources updated after 1/02/2019 except for events not
	// tracked.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1/02/2019 except for events not
	// tracked.
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// User-defined unique identifier for the Product Selection.
	Key *string `json:"key,omitempty"`
	// Name of the Product Selection.
	Name LocalizedString `json:"name"`
	// Number of Products that are currently assigned to this Product Selection.
	ProductCount int `json:"productCount"`
	// Specifies in which way the Products are assigned to the Product Selection. Currently, the only way of doing this is to specify each Product individually. Hence, the type is fixed to `individual` for now, but we have plans to add other types in the future.
	Type ProductSelectionTypeEnum `json:"type"`
	// Custom Fields of this Product Selection.
	Custom *CustomFields `json:"custom,omitempty"`
}

/**
*	Specifies which Product is assigned to which Product Selection.
 */
type ProductSelectionAssignment struct {
	// Reference to a Product that is assigned to the Product Selection.
	Product ProductReference `json:"product"`
	// Reference to the Product Selection that this assignment is part of.
	ProductSelection ProductSelectionReference `json:"productSelection"`
}

type ProductSelectionDraft struct {
	// User-defined unique identifier for the Product Selection. You can use `key` besides `ID` to reference the Product Selection.
	Key *string `json:"key,omitempty"`
	// Name of the Product Selection. Not checked for uniqueness, but distinct names are recommended.
	Name LocalizedString `json:"name"`
	// Custom Fields of this Product Selection.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
}

/**
*	[PagedQueryResult](/general-concepts#pagedqueryresult) containing an array of [ProductSelection](ctp:api:type:ProductSelection).
*
 */
type ProductSelectionPagedQueryResponse struct {
	// Number of results requested in the query request.
	Limit int `json:"limit"`
	// Offset supplied by the client or the server default.
	// It is the number of elements skipped, not a page number.
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/general-concepts#strong-consistency).
	// Unlike other endpoints, the Product Selection endpoint does not return this field by default.
	// To get `total`, pass the query parameter `withTotal` set to `true`.
	// When the results are filtered with a [Query Predicate](/predicates/query), `total` is subject to a [limit](/limits#queries).
	Total *int `json:"total,omitempty"`
	// The Product Selections matching the query.
	Results []ProductSelection `json:"results"`
}

/**
*	[PagedQueryResult](/general-concepts#pagedqueryresult) containing an array of [AssignedProductReference](ctp:api:type:AssignedProductReference).
*
 */
type ProductSelectionProductPagedQueryResponse struct {
	// Number of results requested in the query request.
	Limit int `json:"limit"`
	// Offset supplied by the client or the server default.
	// It is the number of elements skipped, not a page number.
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/general-concepts#strong-consistency).
	// Unlike other endpoints, the Product Selection endpoint does not return this field by default.
	// To get `total`, pass the query parameter `withTotal` set to `true`.
	// When the results are filtered with a [Query Predicate](/predicates/query), `total` is subject to a [limit](/limits#queries).
	Total *int `json:"total,omitempty"`
	// References to Products that are assigned to the Product Selection.
	Results []AssignedProductReference `json:"results"`
}

type ProductSelectionReference struct {
	// Unique ID of the Product Selection.
	ID string `json:"id"`
	// Contains the representation of the expanded Product Selection. Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for Product Selection.
	Obj *ProductSelection `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSelectionReference) MarshalJSON() ([]byte, error) {
	type Alias ProductSelectionReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "product-selection", Alias: (*Alias)(&obj)})
}

type ProductSelectionResourceIdentifier struct {
	// Unique ID of the referenced resource. Either `id` or `key` is required.
	ID *string `json:"id,omitempty"`
	// Unique key of the referenced resource. Either `id` or `key` is required.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSelectionResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias ProductSelectionResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "product-selection", Alias: (*Alias)(&obj)})
}

type ProductSelectionType interface{}

func mapDiscriminatorProductSelectionType(input interface{}) (ProductSelectionType, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "individual":
		obj := IndividualProductSelectionType{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type IndividualProductSelectionType struct {
	// The name of the Product Selection which is recommended to be unique.
	Name LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj IndividualProductSelectionType) MarshalJSON() ([]byte, error) {
	type Alias IndividualProductSelectionType
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "individual", Alias: (*Alias)(&obj)})
}

/**
*	The following type of Product Selections is supported:
*
 */
type ProductSelectionTypeEnum string

const (
	ProductSelectionTypeEnumIndividual ProductSelectionTypeEnum = "individual"
)

type ProductSelectionUpdate struct {
	Version int                            `json:"version"`
	Actions []ProductSelectionUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductSelectionUpdate) UnmarshalJSON(data []byte) error {
	type Alias ProductSelectionUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorProductSelectionUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}
	return nil
}

type ProductSelectionUpdateAction interface{}

func mapDiscriminatorProductSelectionUpdateAction(input interface{}) (ProductSelectionUpdateAction, error) {
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
	case "addProduct":
		obj := ProductSelectionAddProductAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeName":
		obj := ProductSelectionChangeNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeProduct":
		obj := ProductSelectionRemoveProductAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomField":
		obj := ProductSelectionSetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomType":
		obj := ProductSelectionSetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := ProductSelectionSetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	[PagedQueryResult](/general-concepts#pagedqueryresult) containing an array of [ProductSelectionAssignment](ctp:api:type:ProductSelectionAssignment).
*
 */
type ProductsInStorePagedQueryResponse struct {
	// Number of results requested in the query request.
	Limit int `json:"limit"`
	// Offset supplied by the client or the server default.
	// It is the number of elements skipped, not a page number.
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/general-concepts#strong-consistency).
	// Unlike other endpoints, the Product Selection endpoint does not return this field by default.
	// To get `total`, pass the query parameter `withTotal` set to `true`.
	// When the results are filtered with a [Query Predicate](/predicates/query), `total` is subject to a [limit](/limits#queries).
	Total *int `json:"total,omitempty"`
	// Product Selection Assignments.
	Results []ProductSelectionAssignment `json:"results"`
}

type ProductSelectionAddProductAction struct {
	// ResourceIdentifier to Product
	Product ProductResourceIdentifier `json:"product"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSelectionAddProductAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSelectionAddProductAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addProduct", Alias: (*Alias)(&obj)})
}

type ProductSelectionChangeNameAction struct {
	// The new name to be set for the Product Selection.
	Name LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSelectionChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSelectionChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type ProductSelectionRemoveProductAction struct {
	// ResourceIdentifier to Product
	Product ProductResourceIdentifier `json:"product"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSelectionRemoveProductAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSelectionRemoveProductAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeProduct", Alias: (*Alias)(&obj)})
}

type ProductSelectionSetCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Trying to remove a field that does not exist will fail with an [InvalidOperation](/../api/errors#general-400-invalid-operation) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSelectionSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSelectionSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type ProductSelectionSetCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the ProductSelection with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the ProductSelection.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the ProductSelection.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSelectionSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSelectionSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type ProductSelectionSetKeyAction struct {
	// If `key` is absent or `null`, the existing key, if any, will be removed.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSelectionSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSelectionSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}
