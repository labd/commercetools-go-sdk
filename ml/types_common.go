// Generated file, please do not change!!!
package ml

import (
	"encoding/json"
	"errors"
)

// LocalizedString is something
type LocalizedString map[string]string
type Money struct {
	CentAmount int `json:"centAmount"`
	// The currency code compliant to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	CurrencyCode string `json:"currencyCode"`
}

type ReferenceTypeId string

const (
	ReferenceTypeIdCart             ReferenceTypeId = "cart"
	ReferenceTypeIdCartDiscount     ReferenceTypeId = "cart-discount"
	ReferenceTypeIdCategory         ReferenceTypeId = "category"
	ReferenceTypeIdChannel          ReferenceTypeId = "channel"
	ReferenceTypeIdCustomer         ReferenceTypeId = "customer"
	ReferenceTypeIdCustomerGroup    ReferenceTypeId = "customer-group"
	ReferenceTypeIdDiscountCode     ReferenceTypeId = "discount-code"
	ReferenceTypeIdKeyValueDocument ReferenceTypeId = "key-value-document"
	ReferenceTypeIdPayment          ReferenceTypeId = "payment"
	ReferenceTypeIdProduct          ReferenceTypeId = "product"
	ReferenceTypeIdProductType      ReferenceTypeId = "product-type"
	ReferenceTypeIdProductDiscount  ReferenceTypeId = "product-discount"
	ReferenceTypeIdOrder            ReferenceTypeId = "order"
	ReferenceTypeIdReview           ReferenceTypeId = "review"
	ReferenceTypeIdShoppingList     ReferenceTypeId = "shopping-list"
	ReferenceTypeIdShippingMethod   ReferenceTypeId = "shipping-method"
	ReferenceTypeIdState            ReferenceTypeId = "state"
	ReferenceTypeIdStore            ReferenceTypeId = "store"
	ReferenceTypeIdTaxCategory      ReferenceTypeId = "tax-category"
	ReferenceTypeIdType             ReferenceTypeId = "type"
	ReferenceTypeIdZone             ReferenceTypeId = "zone"
	ReferenceTypeIdInventoryEntry   ReferenceTypeId = "inventory-entry"
	ReferenceTypeIdOrderEdit        ReferenceTypeId = "order-edit"
)

type Reference interface{}

func mapDiscriminatorReference(input interface{}) (Reference, error) {

	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["typeId"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'typeId'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}

	switch discriminator {
	case "category":
		obj := CategoryReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "product":
		obj := ProductReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "product-type":
		obj := ProductTypeReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type CategoryReference struct {
	ID string `json:"id"`
}

// MarshalJSON override to set the discriminator value
func (obj CategoryReference) MarshalJSON() ([]byte, error) {
	type Alias CategoryReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "category", Alias: (*Alias)(&obj)})
}

type ProductReference struct {
	ID string `json:"id"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductReference) MarshalJSON() ([]byte, error) {
	type Alias ProductReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "product", Alias: (*Alias)(&obj)})
}

type ProductTypeReference struct {
	ID string `json:"id"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductTypeReference) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "product-type", Alias: (*Alias)(&obj)})
}

/**
*	The product variant that contains the image.
 */
type ProductVariant struct {
	// The product that contains this variant.
	Product ProductReference `json:"product"`
	// The state of the product variant.
	Staged bool `json:"staged"`
	// The id of the product variant.
	VariantId int `json:"variantId"`
}

type TaskStatusEnum string

const (
	TaskStatusEnumPENDING TaskStatusEnum = "PENDING"
	TaskStatusEnumSUCCESS TaskStatusEnum = "SUCCESS"
)

/**
*	Represents a URL path to poll to get the results of an Asynchronous Request.
 */
type TaskToken struct {
	// The ID for the task. Used to find the status of the task.
	TaskId string `json:"taskId"`
	// The URI path to poll for the status of the task.
	UriPath string `json:"uriPath"`
}
