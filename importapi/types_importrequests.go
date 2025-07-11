package importapi

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
)

/**
*	An import request batches multiple import resources of the same import resource type for processing by an Import Container.
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
	case "type":
		obj := TypeImportRequest{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "discount-code":
		obj := DiscountCodeImportRequest{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "product-selection":
		obj := ProductSelectionImportRequest{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	The response of each Import Request.
*
 */
type ImportResponse struct {
	// The identifiers and status of the [ImportOperations](ctp:import:type:ImportOperation) created by the ImportRequest.
	OperationStatus []ImportOperationStatus `json:"operationStatus"`
}

/**
*	The request body to [import Categories](ctp:import:endpoint:/{projectKey}/categories/import-containers/{importContainerKey}:POST). Contains data for [Categories](ctp:api:type:Category) to be created or updated in a Project.
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
*	The request body to [import Products](ctp:import:endpoint:/{projectKey}/products/import-containers/{importContainerKey}:POST). Contains data for [Products](ctp:api:type:Product) to be created or updated in a Project.
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
*	The request body to [import ProductDrafts](ctp:import:endpoint:/{projectKey}/product-drafts/import-containers/{importContainerKey}:POST). Contains data for [Products](ctp:api:type:Product) to be created or updated in a Project.
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
*	The request body to [import ProductTypes](ctp:import:endpoint:/{projectKey}/product-types/import-containers/{importContainerKey}:POST). Contains data for [ProductTypes](ctp:api:type:ProductType) to be created or updated in a Project.
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
*	The request body to [import ProductVariants](ctp:import:endpoint:/{projectKey}/product-variants/import-containers/{importContainerKey}:POST). Contains data for [ProductVariants](ctp:api:type:ProductVariant) to be created or updated in a Project.
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
*	The request body to [import Embedded Prices](ctp:import:endpoint:/{projectKey}/prices/import-containers/{importContainerKey}:POST). Contains data for [Embedded Prices](/../api/types#price) to be created or updated in a Project.
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
*	The request body to [import Standalone Prices](ctp:import:endpoint:/{projectKey}/standalone-prices/import-containers/{importContainerKey}:POST). Contains data for [Standalone Prices](ctp:api:type:StandalonePrice) to be created or updated in a Project.
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
*	The request body to [import Orders](ctp:import:endpoint:/{projectKey}/orders/import-containers/{importContainerKey}:POST). Contains data for [Orders](ctp:api:type:Order) to be created in a Project.
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
*	The request body to [import OrderPatches](ctp:import:endpoint:/{projectKey}/order-patches/import-containers/{importContainerKey}:POST). The data to be imported are represented by [OrderPatchImport](ctp:import:type:OrderPatchImport).
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
*	The request body to [import ProductVariantPatches](ctp:import:endpoint:/{projectKey}/product-variant-patches/import-containers/{importContainerKey}:POST). The data to be imported are represented by [ProductVariantPatch](ctp:import:type:ProductVariantPatch).
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
*	The request body to [import Customers](ctp:import:endpoint:/{projectKey}/customers/import-containers/{importContainerKey}:POST). Contains data for [Customers](ctp:api:type:Customer) to be created or updated in a Project.
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
*	The request body to [import Inventories](ctp:import:endpoint:/{projectKey}/inventories/import-containers/{importContainerKey}:POST). Contains data for [InventoryEntries](ctp:api:type:InventoryEntry) to be created or updated in a commercetools Project.
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

/**
*	The request body to [import Types](ctp:import:endpoint:/{projectKey}/types/import-containers/{importContainerKey}:POST). Contains data for [Types](ctp:api:type:Type) to be created or updated in a Project.
*
 */
type TypeImportRequest struct {
	// The type import resources of this request.
	Resources []TypeImport `json:"resources"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TypeImportRequest) MarshalJSON() ([]byte, error) {
	type Alias TypeImportRequest
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "type", Alias: (*Alias)(&obj)})
}

/**
*	The request body to [import DiscountCodes](ctp:import:endpoint:/{projectKey}/discount-codes/import-containers/{importContainerKey}:POST). Contains data for [Discount Codes](ctp:api:type:DiscountCode) to be created or updated in a Project.
*
 */
type DiscountCodeImportRequest struct {
	// The Discount Code import resources of this request.
	Resources []DiscountCodeImport `json:"resources"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeImportRequest) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeImportRequest
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "discount-code", Alias: (*Alias)(&obj)})
}

/**
*	The request body to [import Product Selections](ctp:import:endpoint:/{projectKey}/product-selections/import-containers/{importContainerKey}:POST). Contains data for [Product Selections](ctp:api:type:ProductSelection) to be created or updated in a Project.
*
 */
type ProductSelectionImportRequest struct {
	// The Product Selection import resources of this request.
	Resources []ProductSelectionImport `json:"resources"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSelectionImportRequest) MarshalJSON() ([]byte, error) {
	type Alias ProductSelectionImportRequest
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "product-selection", Alias: (*Alias)(&obj)})
}
