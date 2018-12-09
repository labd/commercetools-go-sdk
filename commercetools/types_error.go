// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"

	mapstructure "github.com/mitchellh/mapstructure"
)

type AccessDeniedError struct {
	Message string `json:"message"`
}

func (obj AccessDeniedError) MarshalJSON() ([]byte, error) {
	type Alias AccessDeniedError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "access_denied", Alias: (*Alias)(&obj)})
}

func (e AccessDeniedError) Error() string {
	return e.Message
}

type ConcurrentModificationError struct {
	Message        string `json:"message"`
	CurrentVersion int    `json:"currentVersion,omitempty"`
}

func (obj ConcurrentModificationError) MarshalJSON() ([]byte, error) {
	type Alias ConcurrentModificationError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "ConcurrentModification", Alias: (*Alias)(&obj)})
}

func (e ConcurrentModificationError) Error() string {
	return e.Message
}

type DiscountCodeNonApplicableError struct {
	Message string `json:"message"`
}

func (obj DiscountCodeNonApplicableError) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeNonApplicableError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "DiscountCodeNonApplicable", Alias: (*Alias)(&obj)})
}

func (e DiscountCodeNonApplicableError) Error() string {
	return e.Message
}

type DuplicateAttributeValueError struct {
	Message   string     `json:"message"`
	Attribute *Attribute `json:"attribute"`
}

func (obj DuplicateAttributeValueError) MarshalJSON() ([]byte, error) {
	type Alias DuplicateAttributeValueError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "DuplicateAttributeValue", Alias: (*Alias)(&obj)})
}

func (e DuplicateAttributeValueError) Error() string {
	return e.Message
}

type DuplicateAttributeValuesError struct {
	Message    string      `json:"message"`
	Attributes []Attribute `json:"attributes"`
}

func (obj DuplicateAttributeValuesError) MarshalJSON() ([]byte, error) {
	type Alias DuplicateAttributeValuesError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "DuplicateAttributeValues", Alias: (*Alias)(&obj)})
}

func (e DuplicateAttributeValuesError) Error() string {
	return e.Message
}

type DuplicateFieldError struct {
	Message        string      `json:"message"`
	Field          string      `json:"field,omitempty"`
	DuplicateValue interface{} `json:"duplicateValue,omitempty"`
}

func (obj DuplicateFieldError) MarshalJSON() ([]byte, error) {
	type Alias DuplicateFieldError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "DuplicateField", Alias: (*Alias)(&obj)})
}

func (e DuplicateFieldError) Error() string {
	return e.Message
}

type DuplicatePriceScopeError struct {
	Message           string  `json:"message"`
	ConflictingPrices []Price `json:"conflictingPrices"`
}

func (obj DuplicatePriceScopeError) MarshalJSON() ([]byte, error) {
	type Alias DuplicatePriceScopeError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "DuplicatePriceScope", Alias: (*Alias)(&obj)})
}

func (e DuplicatePriceScopeError) Error() string {
	return e.Message
}

type DuplicateVariantValuesError struct {
	Message       string         `json:"message"`
	VariantValues *VariantValues `json:"variantValues"`
}

func (obj DuplicateVariantValuesError) MarshalJSON() ([]byte, error) {
	type Alias DuplicateVariantValuesError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "DuplicateVariantValues", Alias: (*Alias)(&obj)})
}

func (e DuplicateVariantValuesError) Error() string {
	return e.Message
}

type ErrorObject interface{}
type AbstractErrorObject struct {
	Message string `json:"message"`
}

func AbstractErrorObjectDiscriminatorMapping(input ErrorObject) ErrorObject {
	discriminator := input.(map[string]interface{})["code"]
	switch discriminator {
	case "access_denied":
		new := AccessDeniedError{}
		mapstructure.Decode(input, &new)
		return new
	case "ConcurrentModification":
		new := ConcurrentModificationError{}
		mapstructure.Decode(input, &new)
		return new
	case "DiscountCodeNonApplicable":
		new := DiscountCodeNonApplicableError{}
		mapstructure.Decode(input, &new)
		return new
	case "DuplicateAttributeValue":
		new := DuplicateAttributeValueError{}
		mapstructure.Decode(input, &new)
		return new
	case "DuplicateAttributeValues":
		new := DuplicateAttributeValuesError{}
		mapstructure.Decode(input, &new)
		return new
	case "DuplicateField":
		new := DuplicateFieldError{}
		mapstructure.Decode(input, &new)
		return new
	case "DuplicatePriceScope":
		new := DuplicatePriceScopeError{}
		mapstructure.Decode(input, &new)
		return new
	case "DuplicateVariantValues":
		new := DuplicateVariantValuesError{}
		mapstructure.Decode(input, &new)
		return new
	case "insufficient_scope":
		new := InsufficientScopeError{}
		mapstructure.Decode(input, &new)
		return new
	case "InvalidCredentials":
		new := InvalidCredentialsError{}
		mapstructure.Decode(input, &new)
		return new
	case "InvalidCurrentPassword":
		new := InvalidCurrentPasswordError{}
		mapstructure.Decode(input, &new)
		return new
	case "InvalidField":
		new := InvalidFieldError{}
		mapstructure.Decode(input, &new)
		return new
	case "InvalidInput":
		new := InvalidInputError{}
		mapstructure.Decode(input, &new)
		return new
	case "InvalidItemShippingDetails":
		new := InvalidItemShippingDetailsError{}
		mapstructure.Decode(input, &new)
		return new
	case "InvalidJsonInput":
		new := InvalidJsonInputError{}
		mapstructure.Decode(input, &new)
		return new
	case "InvalidOperation":
		new := InvalidOperationError{}
		mapstructure.Decode(input, &new)
		return new
	case "InvalidSubject":
		new := InvalidSubjectError{}
		mapstructure.Decode(input, &new)
		return new
	case "invalid_token":
		new := InvalidTokenError{}
		mapstructure.Decode(input, &new)
		return new
	case "MissingTaxRateForCountry":
		new := MissingTaxRateForCountryError{}
		mapstructure.Decode(input, &new)
		return new
	case "NoMatchingProductDiscountFound":
		new := NoMatchingProductDiscountFoundError{}
		mapstructure.Decode(input, &new)
		return new
	case "OutOfStock":
		new := OutOfStockError{}
		mapstructure.Decode(input, &new)
		return new
	case "PriceChanged":
		new := PriceChangedError{}
		mapstructure.Decode(input, &new)
		return new
	case "ReferenceExists":
		new := ReferenceExistsError{}
		mapstructure.Decode(input, &new)
		return new
	case "RequiredField":
		new := RequiredFieldError{}
		mapstructure.Decode(input, &new)
		return new
	case "ResourceNotFound":
		new := ResourceNotFoundError{}
		mapstructure.Decode(input, &new)
		return new
	case "ShippingMethodDoesNotMatchCart":
		new := ShippingMethodDoesNotMatchCartError{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}

type ErrorResponse struct {
	StatusCode        int           `json:"statusCode"`
	Message           string        `json:"message"`
	Errors            []ErrorObject `json:"errors,omitempty"`
	Error_description string        `json:"error_description,omitempty"`
	_Error            string        `json:"error,omitempty"`
}

func (obj *ErrorResponse) UnmarshalJSON(data []byte) error {
	type Alias ErrorResponse
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Errors {
		obj.Errors[i] = AbstractErrorObjectDiscriminatorMapping(obj.Errors[i])
	}

	return nil
}

func (e ErrorResponse) Error() string {
	return e.Message
}

type InsufficientScopeError struct {
	Message string `json:"message"`
}

func (obj InsufficientScopeError) MarshalJSON() ([]byte, error) {
	type Alias InsufficientScopeError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "insufficient_scope", Alias: (*Alias)(&obj)})
}

func (e InsufficientScopeError) Error() string {
	return e.Message
}

type InvalidCredentialsError struct {
	Message string `json:"message"`
}

func (obj InvalidCredentialsError) MarshalJSON() ([]byte, error) {
	type Alias InvalidCredentialsError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "InvalidCredentials", Alias: (*Alias)(&obj)})
}

func (e InvalidCredentialsError) Error() string {
	return e.Message
}

type InvalidCurrentPasswordError struct {
	Message string `json:"message"`
}

func (obj InvalidCurrentPasswordError) MarshalJSON() ([]byte, error) {
	type Alias InvalidCurrentPasswordError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "InvalidCurrentPassword", Alias: (*Alias)(&obj)})
}

func (e InvalidCurrentPasswordError) Error() string {
	return e.Message
}

type InvalidFieldError struct {
	Message       string        `json:"message"`
	InvalidValue  interface{}   `json:"invalidValue"`
	Field         string        `json:"field"`
	AllowedValues []interface{} `json:"allowedValues,omitempty"`
}

func (obj InvalidFieldError) MarshalJSON() ([]byte, error) {
	type Alias InvalidFieldError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "InvalidField", Alias: (*Alias)(&obj)})
}

func (e InvalidFieldError) Error() string {
	return e.Message
}

type InvalidInputError struct {
	Message string `json:"message"`
}

func (obj InvalidInputError) MarshalJSON() ([]byte, error) {
	type Alias InvalidInputError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "InvalidInput", Alias: (*Alias)(&obj)})
}

func (e InvalidInputError) Error() string {
	return e.Message
}

type InvalidItemShippingDetailsError struct {
	Message string `json:"message"`
	Subject string `json:"subject"`
	ItemID  string `json:"itemId"`
}

func (obj InvalidItemShippingDetailsError) MarshalJSON() ([]byte, error) {
	type Alias InvalidItemShippingDetailsError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "InvalidItemShippingDetails", Alias: (*Alias)(&obj)})
}

func (e InvalidItemShippingDetailsError) Error() string {
	return e.Message
}

type InvalidJsonInputError struct {
	Message string `json:"message"`
}

func (obj InvalidJsonInputError) MarshalJSON() ([]byte, error) {
	type Alias InvalidJsonInputError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "InvalidJsonInput", Alias: (*Alias)(&obj)})
}

func (e InvalidJsonInputError) Error() string {
	return e.Message
}

type InvalidOperationError struct {
	Message string `json:"message"`
}

func (obj InvalidOperationError) MarshalJSON() ([]byte, error) {
	type Alias InvalidOperationError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "InvalidOperation", Alias: (*Alias)(&obj)})
}

func (e InvalidOperationError) Error() string {
	return e.Message
}

type InvalidSubjectError struct {
	Message string `json:"message"`
}

func (obj InvalidSubjectError) MarshalJSON() ([]byte, error) {
	type Alias InvalidSubjectError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "InvalidSubject", Alias: (*Alias)(&obj)})
}

func (e InvalidSubjectError) Error() string {
	return e.Message
}

type InvalidTokenError struct {
	Message string `json:"message"`
}

func (obj InvalidTokenError) MarshalJSON() ([]byte, error) {
	type Alias InvalidTokenError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "invalid_token", Alias: (*Alias)(&obj)})
}

func (e InvalidTokenError) Error() string {
	return e.Message
}

type MissingTaxRateForCountryError struct {
	Message string `json:"message"`
}

func (obj MissingTaxRateForCountryError) MarshalJSON() ([]byte, error) {
	type Alias MissingTaxRateForCountryError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "MissingTaxRateForCountry", Alias: (*Alias)(&obj)})
}

func (e MissingTaxRateForCountryError) Error() string {
	return e.Message
}

type NoMatchingProductDiscountFoundError struct {
	Message string `json:"message"`
}

func (obj NoMatchingProductDiscountFoundError) MarshalJSON() ([]byte, error) {
	type Alias NoMatchingProductDiscountFoundError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "NoMatchingProductDiscountFound", Alias: (*Alias)(&obj)})
}

func (e NoMatchingProductDiscountFoundError) Error() string {
	return e.Message
}

type OutOfStockError struct {
	Message   string   `json:"message"`
	Skus      []string `json:"skus"`
	LineItems []string `json:"lineItems"`
}

func (obj OutOfStockError) MarshalJSON() ([]byte, error) {
	type Alias OutOfStockError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "OutOfStock", Alias: (*Alias)(&obj)})
}

func (e OutOfStockError) Error() string {
	return e.Message
}

type PriceChangedError struct {
	Message   string   `json:"message"`
	Shipping  bool     `json:"shipping"`
	LineItems []string `json:"lineItems"`
}

func (obj PriceChangedError) MarshalJSON() ([]byte, error) {
	type Alias PriceChangedError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "PriceChanged", Alias: (*Alias)(&obj)})
}

func (e PriceChangedError) Error() string {
	return e.Message
}

type ReferenceExistsError struct {
	Message      string          `json:"message"`
	ReferencedBy ReferenceTypeID `json:"referencedBy,omitempty"`
}

func (obj ReferenceExistsError) MarshalJSON() ([]byte, error) {
	type Alias ReferenceExistsError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "ReferenceExists", Alias: (*Alias)(&obj)})
}

func (e ReferenceExistsError) Error() string {
	return e.Message
}

type RequiredFieldError struct {
	Message string `json:"message"`
	Field   string `json:"field"`
}

func (obj RequiredFieldError) MarshalJSON() ([]byte, error) {
	type Alias RequiredFieldError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "RequiredField", Alias: (*Alias)(&obj)})
}

func (e RequiredFieldError) Error() string {
	return e.Message
}

type ResourceNotFoundError struct {
	Message string `json:"message"`
}

func (obj ResourceNotFoundError) MarshalJSON() ([]byte, error) {
	type Alias ResourceNotFoundError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "ResourceNotFound", Alias: (*Alias)(&obj)})
}

func (e ResourceNotFoundError) Error() string {
	return e.Message
}

type ShippingMethodDoesNotMatchCartError struct {
	Message string `json:"message"`
}

func (obj ShippingMethodDoesNotMatchCartError) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodDoesNotMatchCartError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "ShippingMethodDoesNotMatchCart", Alias: (*Alias)(&obj)})
}

func (e ShippingMethodDoesNotMatchCartError) Error() string {
	return e.Message
}

type VariantValues struct {
	Sku        string      `json:"sku,omitempty"`
	Prices     []Price     `json:"prices"`
	Attributes []Attribute `json:"attributes"`
}
