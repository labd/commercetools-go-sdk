package importapi

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
)

/**
*	An import request batches multiple import resources of the same import resource type for processing by an import container.
*
 */
type ImportRequest interface{}

func mapDiscriminatorImportRequest(input interface{}) (ImportRequest, error) {
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
	case "category":
		obj := CategoryImportRequest{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "product":
		obj := ProductImportRequest{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "product-draft":
		obj := ProductDraftImportRequest{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "product-type":
		obj := ProductTypeImportRequest{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "product-variant":
		obj := ProductVariantImportRequest{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "price":
		obj := PriceImportRequest{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "standalone-price":
		obj := StandalonePriceImportRequest{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "order":
		obj := OrderImportRequest{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "order-patch":
		obj := OrderPatchImportRequest{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "product-variant-patch":
		obj := ProductVariantPatchRequest{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "customer":
		obj := CustomerImportRequest{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "inventory":
		obj := InventoryImportRequest{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	A list of the ID's and validation statuses of newly created [ImportOperations](#importoperation).
*	Used as a response at each resource-specific import endpoint, for example, at [Import Categories](/category#import-categories) and [Import ProductTypes](/product-type#import-producttypes).
*
 */
type ImportResponse struct {
	OperationStatus []ImportOperationStatus `json:"operationStatus"`
}

/**
*	The request body to [import Categories](#import-categories). Contains data for [Categories](/../api/projects/categories#category) to be created or updated in a Project.
*
 */
type CategoryImportRequest struct {
	// The category import resources of this request.
	Resources []CategoryImport `json:"resources"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CategoryImportRequest) MarshalJSON() ([]byte, error) {
	type Alias CategoryImportRequest
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "category", Alias: (*Alias)(&obj)})
}

/**
*	The request body to [import Products](#import-products). Contains data for [Products](/../api/projects/products#product) to be created or updated in a Project.
*
 */
type ProductImportRequest struct {
	// The product import resources of this request.
	Resources []ProductImport `json:"resources"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductImportRequest) MarshalJSON() ([]byte, error) {
	type Alias ProductImportRequest
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "product", Alias: (*Alias)(&obj)})
}

/**
*	The request body to [import ProductDrafts](#import-productdrafts). Contains data for [Products](/../api/projects/products#productdraft) to be created or updated in a Project.
*
 */
type ProductDraftImportRequest struct {
	// The product draft import resources of this request.
	Resources []ProductDraftImport `json:"resources"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDraftImportRequest) MarshalJSON() ([]byte, error) {
	type Alias ProductDraftImportRequest
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "product-draft", Alias: (*Alias)(&obj)})
}

/**
*	The request body to [import ProductTypes](#import-producttypes). Contains data for [ProductTypes](/../api/projects/productTypes#producttype) to be created or updated in a Project.
*
 */
type ProductTypeImportRequest struct {
	// The product type import resources of this request.
	Resources []ProductTypeImport `json:"resources"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTypeImportRequest) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeImportRequest
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "product-type", Alias: (*Alias)(&obj)})
}

/**
*	The request body to [import ProductVariants](#import-productvariants). Contains data for [ProductVariants](/../api/projects/products#productvariant) to be created or updated in a Project.
*
 */
type ProductVariantImportRequest struct {
	// The product variant import resources of this request.
	Resources []ProductVariantImport `json:"resources"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductVariantImportRequest) MarshalJSON() ([]byte, error) {
	type Alias ProductVariantImportRequest
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "product-variant", Alias: (*Alias)(&obj)})
}

/**
*	The request body to [import Embedded Prices](#import-embedded-prices). Contains data for [Embedded Prices](/../api/types#embedded-price) to be created or updated in a Project.
*
 */
type PriceImportRequest struct {
	// The price import resources of this request.
	Resources []PriceImport `json:"resources"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PriceImportRequest) MarshalJSON() ([]byte, error) {
	type Alias PriceImportRequest
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "price", Alias: (*Alias)(&obj)})
}

/**
*	The request body to [import Standalone Prices](#import-standalone-prices). Contains data for [Standalone Prices](/../api/projects/standalone-prices#standaloneprice) to be created or updated in a Project.
*
 */
type StandalonePriceImportRequest struct {
	// The Standalone Price import resources of this request.
	Resources []StandalonePriceImport `json:"resources"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StandalonePriceImportRequest) MarshalJSON() ([]byte, error) {
	type Alias StandalonePriceImportRequest
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "standalone-price", Alias: (*Alias)(&obj)})
}

/**
*	The request body to [import Orders](#import-orders). Contains data for [Orders](/../api/projects/orders#order) to be created or updated in a Project.
*
 */
type OrderImportRequest struct {
	// The order import resources of this request.
	Resources []OrderImport `json:"resources"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderImportRequest) MarshalJSON() ([]byte, error) {
	type Alias OrderImportRequest
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "order", Alias: (*Alias)(&obj)})
}

/**
*	The request body to [import OrderPatches](#import-orderpatches). The data to be imported are represented by [OrderPatchImport](#orderpatchimport).
*
 */
type OrderPatchImportRequest struct {
	// The order patches of this request
	Patches []OrderPatchImport `json:"patches"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderPatchImportRequest) MarshalJSON() ([]byte, error) {
	type Alias OrderPatchImportRequest
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "order-patch", Alias: (*Alias)(&obj)})
}

/**
*	The request body to [import ProductVariantPatches](#import-productvariantpatches). The data to be imported are represented by [ProductVariantPatch](#productvariantpatch).
*
 */
type ProductVariantPatchRequest struct {
	// The product variant patches of this request.
	Patches []ProductVariantPatch `json:"patches"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductVariantPatchRequest) MarshalJSON() ([]byte, error) {
	type Alias ProductVariantPatchRequest
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "product-variant-patch", Alias: (*Alias)(&obj)})
}

/**
*	The request body to [import Customers](#import-customers). Contains data for [Customers](/../api/projects/customers#customer) to be created or updated in a Project.
*
 */
type CustomerImportRequest struct {
	// The customer import resources of this request.
	Resources []CustomerImport `json:"resources"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerImportRequest) MarshalJSON() ([]byte, error) {
	type Alias CustomerImportRequest
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "customer", Alias: (*Alias)(&obj)})
}

/**
*	The request body to [import Inventories](#import-inventory). Contains data for [InventoryEntries](/../api/projects/inventory#inventoryentry) to be created or updated in a commercetools Project.
*
 */
type InventoryImportRequest struct {
	// The inventory import resources of this request.
	Resources []InventoryImport `json:"resources"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InventoryImportRequest) MarshalJSON() ([]byte, error) {
	type Alias InventoryImportRequest
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "inventory", Alias: (*Alias)(&obj)})
}
