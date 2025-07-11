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
	// The Variants of the Product that are included from the Product Selection.
	//
	// This field may exist only in Product Selections with `Individual` [ProductSelectionMode](ctp:api:type:ProductSelectionMode).
	// In absence of this field, all Variants are deemed to be included.
	VariantSelection ProductVariantSelection `json:"variantSelection,omitempty"`
	// The Variants of the Product that are excluded from the Product Selection.
	//
	// This field may exist only in Product Selections with `IndividualExclusion` [ProductSelectionMode](ctp:api:type:ProductSelectionMode).
	// In absence of this field, all Variants are deemed to be excluded.
	VariantExclusion *ProductVariantExclusion `json:"variantExclusion,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AssignedProductReference) UnmarshalJSON(data []byte) error {
	type Alias AssignedProductReference
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.VariantSelection != nil {
		var err error
		obj.VariantSelection, err = mapDiscriminatorProductVariantSelection(obj.VariantSelection)
		if err != nil {
			return err
		}
	}

	return nil
}

type AssignedProductSelection struct {
	// Reference to the Product Selection that this assignment is part of.
	ProductSelection ProductSelectionReference `json:"productSelection"`
	// Defines which Variants of the Product will be included in the Product Selection.
	//
	// This field is only available for assignments to a Product Selection with `Individual` [ProductSelectionMode](ctp:api:type:ProductSelectionMode).
	VariantSelection ProductVariantSelection `json:"variantSelection,omitempty"`
	// Defines which Variants of the Product will be excluded from the Product Selection.
	//
	// This field is only available for assignments to a Product Selection with `IndividualExclusion` [ProductSelectionMode](ctp:api:type:ProductSelectionMode).
	VariantExclusion *ProductVariantExclusion `json:"variantExclusion,omitempty"`
	// Date and time (UTC) this assignment was initially created.
	CreatedAt time.Time `json:"createdAt"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AssignedProductSelection) UnmarshalJSON(data []byte) error {
	type Alias AssignedProductSelection
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.VariantSelection != nil {
		var err error
		obj.VariantSelection, err = mapDiscriminatorProductVariantSelection(obj.VariantSelection)
		if err != nil {
			return err
		}
	}

	return nil
}

/**
*	[PagedQueryResult](/general-concepts#pagedqueryresult) containing an array of [AssignedProductSelection](ctp:api:type:AssignedProductSelection).
*
 */
type AssignedProductSelectionPagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// Present only when the `withTotal` query parameter is set to `true`.
	// This number is an estimation that is not [strongly consistent](/../api/general-concepts#strong-consistency).
	// When the results are filtered with a [Query Predicate](/../api/predicates/query), `total` is subject to a [limit](/../api/limits#queries).
	Total *int `json:"total,omitempty"`
	// References to ProductSelection that are assigned to the Product.
	Results []AssignedProductSelection `json:"results"`
}

type ProductSelection struct {
	// Unique identifier of the ProductSelection.
	ID string `json:"id"`
	// Current version of the ProductSelection.
	Version int `json:"version"`
	// Date and time (UTC) the ProductSelection was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the ProductSelection was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources updated after 1/02/2019 except for [events not tracked](/../api/general-concepts#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1/02/2019 except for [events not tracked](/../api/general-concepts#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// User-defined unique identifier of the ProductSelection.
	Key *string `json:"key,omitempty"`
	// Name of the ProductSelection.
	Name LocalizedString `json:"name"`
	// Number of Products that are currently assigned to this ProductSelection.
	ProductCount int `json:"productCount"`
	// Specifies in which way the Products are assigned to the ProductSelection.
	// Currently, the only way of doing this is to specify each Product individually, either by [including or excluding](ctp:api:type:ProductSelectionMode) them explicitly.
	Type *ProductSelectionTypeEnum `json:"type,omitempty"`
	// Specifies in which way the Products are assigned to the ProductSelection.
	// Currently, the only way of doing this is to specify each Product individually, either by [including or excluding](ctp:api:type:ProductSelectionMode) them explicitly.
	Mode ProductSelectionMode `json:"mode"`
	// Custom Fields of the ProductSelection.
	Custom *CustomFields `json:"custom,omitempty"`
}

/**
*
*	Given the mode of Product Selection, this assignment refers to, it may contain:
*
*	- `variantSelection` field for a Product Selection with `Individual` [ProductSelectionMode](ctp:api:type:ProductSelectionMode).
*	- `variantExclusion` field for a Product Selection with `IndividualExclusion` [ProductSelectionMode](ctp:api:type:ProductSelectionMode).
 */
type ProductSelectionAssignment struct {
	// Reference to a Product that is assigned to the ProductSelection.
	Product ProductReference `json:"product"`
	// Reference to the Product Selection that this assignment is part of.
	ProductSelection ProductSelectionReference `json:"productSelection"`
	// Defines which particular Variants of the Product are included in the Product Selection.
	// If undefined all Variants of the referenced Product are included.
	//
	// This field is only available for assignments to a Product Selection with `Individual` [ProductSelectionMode](ctp:api:type:ProductSelectionMode).
	// The list of SKUs will be updated automatically on any change of those performed on the respective Product itself.
	VariantSelection ProductVariantSelection `json:"variantSelection,omitempty"`
	// Defines which particular Variants of the Product are excluded from the Product Selection.
	// If undefined all Variants of the referenced Product are excluded.
	//
	// This field is only available for assignments to a Product Selection with `IndividualExclusion` [ProductSelectionMode](ctp:api:type:ProductSelectionMode).
	// The list of SKUs will be updated automatically on any change of those performed on the respective Product itself.
	VariantExclusion *ProductVariantExclusion `json:"variantExclusion,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductSelectionAssignment) UnmarshalJSON(data []byte) error {
	type Alias ProductSelectionAssignment
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.VariantSelection != nil {
		var err error
		obj.VariantSelection, err = mapDiscriminatorProductVariantSelection(obj.VariantSelection)
		if err != nil {
			return err
		}
	}

	return nil
}

type ProductSelectionDraft struct {
	// User-defined unique identifier for the ProductSelection.
	Key *string `json:"key,omitempty"`
	// Name of the ProductSelection. Not checked for uniqueness, but distinct names are recommended.
	Name LocalizedString `json:"name"`
	// Custom Fields of this ProductSelection.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
	// Type of the Product Selection.
	Type *ProductSelectionTypeEnum `json:"type,omitempty"`
	// Mode of the Product Selection.
	Mode *ProductSelectionMode `json:"mode,omitempty"`
}

/**
*	Product Selections can have the following modes:
*
 */
type ProductSelectionMode string

const (
	ProductSelectionModeIndividual          ProductSelectionMode = "Individual"
	ProductSelectionModeIndividualExclusion ProductSelectionMode = "IndividualExclusion"
)

/**
*	[PagedQueryResult](/general-concepts#pagedqueryresult) containing an array of [ProductSelection](ctp:api:type:ProductSelection).
*
 */
type ProductSelectionPagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/general-concepts#strong-consistency).
	// This field is returned by default.
	// For improved performance, calculating this field can be deactivated by using the query parameter `withTotal=false`.
	// When the results are filtered with a [Query Predicate](/predicates/query), `total` is subject to a [limit](/limits#queries).
	Total *int `json:"total,omitempty"`
	// [ProductSelections](ctp:api:type:ProductSelection) matching the query.
	Results []ProductSelection `json:"results"`
}

/**
*	[PagedQueryResult](/general-concepts#pagedqueryresult) containing an array of [AssignedProductReference](ctp:api:type:AssignedProductReference).
*
 */
type ProductSelectionProductPagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// Present only when the `withTotal` query parameter is set to `true`.
	// This number is an estimation that is not [strongly consistent](/../api/general-concepts#strong-consistency).
	// When the results are filtered with a [Query Predicate](/../api/predicates/query), `total` is subject to a [limit](/../api/limits#queries).
	Total *int `json:"total,omitempty"`
	// References to Products that are assigned to the ProductSelection.
	Results []AssignedProductReference `json:"results"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [ProductSelection](ctp:api:type:ProductSelection).
*
 */
type ProductSelectionReference struct {
	// Unique identifier of the referenced [ProductSelection](ctp:api:type:ProductSelection).
	ID string `json:"id"`
	// Contains the representation of the expanded ProductSelection. Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for ProductSelections.
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

/**
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [ProductSelection](ctp:api:type:ProductSelection). Either `id` or `key` is required. If both are set, an [InvalidJsonInput](/../api/errors#invalidjsoninput) error is returned.
*
 */
type ProductSelectionResourceIdentifier struct {
	// Unique identifier of the referenced [ProductSelection](ctp:api:type:ProductSelection). Required if `key` is absent.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced [ProductSelection](ctp:api:type:ProductSelection). Required if `id` is absent.
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
	case "individualExclusion":
		obj := IndividualExclusionProductSelectionType{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "individual":
		obj := IndividualProductSelectionType{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type IndividualExclusionProductSelectionType struct {
	// The name of the ProductSelection which is recommended to be unique.
	Name LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj IndividualExclusionProductSelectionType) MarshalJSON() ([]byte, error) {
	type Alias IndividualExclusionProductSelectionType
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "individualExclusion", Alias: (*Alias)(&obj)})
}

type IndividualProductSelectionType struct {
	// The name of the ProductSelection which is recommended to be unique.
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
*	The following types of Product Selections are supported:
*
 */
type ProductSelectionTypeEnum string

const (
	ProductSelectionTypeEnumIndividual          ProductSelectionTypeEnum = "individual"
	ProductSelectionTypeEnumIndividualExclusion ProductSelectionTypeEnum = "individualExclusion"
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
		if obj.VariantSelection != nil {
			var err error
			obj.VariantSelection, err = mapDiscriminatorProductVariantSelection(obj.VariantSelection)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "changeName":
		obj := ProductSelectionChangeNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "excludeProduct":
		obj := ProductSelectionExcludeProductAction{}
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
	case "setVariantExclusion":
		obj := ProductSelectionSetVariantExclusionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setVariantSelection":
		obj := ProductSelectionSetVariantSelectionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.VariantSelection != nil {
			var err error
			obj.VariantSelection, err = mapDiscriminatorProductVariantSelection(obj.VariantSelection)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Only Product Variants with the explicitly listed SKUs are part of a Product Selection with `IndividualExclusion` [ProductSelectionMode](ctp:api:type:ProductSelectionMode).
*
 */
type ProductVariantExclusion struct {
	// Non-empty array of SKUs representing Product Variants to be included in the Product Selection with `IndividualExclusion` [ProductSelectionMode](ctp:api:type:ProductSelectionMode).
	Skus []string `json:"skus"`
}

/**
*	Polymorphic base type for Product Variant Selections. The actual type is determined by the `type` field.
*
 */
type ProductVariantSelection interface{}

func mapDiscriminatorProductVariantSelection(input interface{}) (ProductVariantSelection, error) {
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
	case "exclusion":
		obj := ProductVariantSelectionExclusion{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "includeAllExcept":
		obj := ProductVariantSelectionIncludeAllExcept{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "includeOnly":
		obj := ProductVariantSelectionIncludeOnly{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "inclusion":
		obj := ProductVariantSelectionInclusion{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	All Product Variants except the explicitly stated SKUs are part of the Product Selection.
*
 */
type ProductVariantSelectionExclusion struct {
	// Non-empty array of SKUs representing Product Variants to be excluded from the Product Selection.
	Skus []string `json:"skus"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductVariantSelectionExclusion) MarshalJSON() ([]byte, error) {
	type Alias ProductVariantSelectionExclusion
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "exclusion", Alias: (*Alias)(&obj)})
}

/**
*	All Product Variants except the explicitly stated SKUs are part of the Product Selection.
*
 */
type ProductVariantSelectionIncludeAllExcept struct {
	// Non-empty array of SKUs representing Product Variants to be excluded from the Product Selection.
	Skus []string `json:"skus"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductVariantSelectionIncludeAllExcept) MarshalJSON() ([]byte, error) {
	type Alias ProductVariantSelectionIncludeAllExcept
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "includeAllExcept", Alias: (*Alias)(&obj)})
}

/**
*	Only Product Variants with explicitly stated SKUs are part of the Product Selection.
*
 */
type ProductVariantSelectionIncludeOnly struct {
	// Non-empty array of SKUs representing Product Variants to be included into the Product Selection.
	Skus []string `json:"skus"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductVariantSelectionIncludeOnly) MarshalJSON() ([]byte, error) {
	type Alias ProductVariantSelectionIncludeOnly
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "includeOnly", Alias: (*Alias)(&obj)})
}

/**
*	Only Product Variants with explicitly stated SKUs are part of the Product Selection.
*
 */
type ProductVariantSelectionInclusion struct {
	// Non-empty array of SKUs representing Product Variants to be included into the Product Selection.
	Skus []string `json:"skus"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductVariantSelectionInclusion) MarshalJSON() ([]byte, error) {
	type Alias ProductVariantSelectionInclusion
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "inclusion", Alias: (*Alias)(&obj)})
}

type ProductVariantSelectionTypeEnum string

const (
	ProductVariantSelectionTypeEnumInclusion        ProductVariantSelectionTypeEnum = "inclusion"
	ProductVariantSelectionTypeEnumExclusion        ProductVariantSelectionTypeEnum = "exclusion"
	ProductVariantSelectionTypeEnumIncludeOnly      ProductVariantSelectionTypeEnum = "includeOnly"
	ProductVariantSelectionTypeEnumIncludeAllExcept ProductVariantSelectionTypeEnum = "includeAllExcept"
)

/**
*	[PagedQueryResult](/general-concepts#pagedqueryresult) containing an array of [ProductSelectionAssignment](ctp:api:type:ProductSelectionAssignment).
*
 */
type ProductsInStorePagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// Present only when the `withTotal` query parameter is set to `true`.
	// This number is an estimation that is not [strongly consistent](/../api/general-concepts#strong-consistency).
	// When the results are filtered with a [Query Predicate](/../api/predicates/query), `total` is subject to a [limit](/../api/limits#queries).
	Total *int `json:"total,omitempty"`
	// ProductSelectionAssignments matching the query.
	Results []ProductSelectionAssignment `json:"results"`
}

/**
*	Adds a Product to the Product Selection.
*
*	If the specified Product is already assigned to the Product Selection, but the existing Product Selection has a different Product Variant Selection, a [ProductPresentWithDifferentVariantSelection](ctp:api:type:ProductPresentWithDifferentVariantSelectionError) error is returned.
*
 */
type ProductSelectionAddProductAction struct {
	// ResourceIdentifier of the Product
	Product ProductResourceIdentifier `json:"product"`
	// Defines which Variants of the Product will be included in the Product Selection.
	// If not supplied all Variants are deemed to be included.
	VariantSelection ProductVariantSelection `json:"variantSelection,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductSelectionAddProductAction) UnmarshalJSON(data []byte) error {
	type Alias ProductSelectionAddProductAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.VariantSelection != nil {
		var err error
		obj.VariantSelection, err = mapDiscriminatorProductVariantSelection(obj.VariantSelection)
		if err != nil {
			return err
		}
	}

	return nil
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
	// The new name to be set for the ProductSelection.
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

/**
*	Excludes a Product from a Product Selection with `IndividualExclusion` [ProductSelectionMode](ctp:api:type:ProductSelectionMode).
*
*	If the specified Product is already assigned to the Product Selection, but the existing Product Selection has a different Product Variant Exclusion, a [ProductPresentWithDifferentVariantSelection](ctp:api:type:ProductPresentWithDifferentVariantSelectionError) error is returned.
*
 */
type ProductSelectionExcludeProductAction struct {
	// ResourceIdentifier of the Product
	Product ProductResourceIdentifier `json:"product"`
	// Defines which Variants of the Product will be excluded from the Product Selection.
	// If not supplied all Variants are deemed to be excluded.
	VariantExclusion *ProductVariantExclusion `json:"variantExclusion,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSelectionExcludeProductAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSelectionExcludeProductAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "excludeProduct", Alias: (*Alias)(&obj)})
}

type ProductSelectionRemoveProductAction struct {
	// ResourceIdentifier of the Product
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
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
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

/**
*	Updates the Product Variant Exclusion of an existing [Product Selection Assignment](ctp:api:type:ProductSelectionAssignment).
*	A [ProductVariantExclusion](ctp:api:type:ProductVariantExclusion) can only be set if the [Product](ctp:api:type:Product) has already been excluded from the [Product Selection](ctp:api:type:ProductSelection) with `IndividualExclusion` [ProductSelectionMode](ctp:api:type:ProductSelectionMode).
*
*	If the specified Product is not assigned to the Product Selection, a [ProductAssignmentMissing](ctp:api:type:ProductAssignmentMissingError) error is returned.
*
 */
type ProductSelectionSetVariantExclusionAction struct {
	// ResourceIdentifier of the Product
	Product ProductResourceIdentifier `json:"product"`
	// Determines which Variants of the previously excluded Product are to be included in the Product Selection with `IndividualExclusion` [ProductSelectionMode](ctp:api:type:ProductSelectionMode).
	// Leave it empty to unset an existing Variant Exclusion.
	VariantExclusion *ProductVariantExclusion `json:"variantExclusion,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSelectionSetVariantExclusionAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSelectionSetVariantExclusionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setVariantExclusion", Alias: (*Alias)(&obj)})
}

/**
*	Updates the Product Variant Selection of an existing [Product Selection Assignment](ctp:api:type:ProductSelectionAssignment).
*	A [ProductVariantSelection](ctp:api:type:ProductVariantSelection) can only be set if the [Product](ctp:api:type:Product) has already been included in the Product Selection with `Individual` [ProductSelectionMode](ctp:api:type:ProductSelectionMode).
*
*	If the specified Product is not assigned to the Product Selection, a [ProductAssignmentMissing](ctp:api:type:ProductAssignmentMissingError) error is returned.
*
 */
type ProductSelectionSetVariantSelectionAction struct {
	// ResourceIdentifier of the Product
	Product ProductResourceIdentifier `json:"product"`
	// Determines which Variants of the previously added Product are to be included in, or excluded from, the Product Selection.
	// Leave it empty to unset an existing Variant Selection.
	VariantSelection ProductVariantSelection `json:"variantSelection,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductSelectionSetVariantSelectionAction) UnmarshalJSON(data []byte) error {
	type Alias ProductSelectionSetVariantSelectionAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.VariantSelection != nil {
		var err error
		obj.VariantSelection, err = mapDiscriminatorProductVariantSelection(obj.VariantSelection)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSelectionSetVariantSelectionAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSelectionSetVariantSelectionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setVariantSelection", Alias: (*Alias)(&obj)})
}
