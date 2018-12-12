// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"

	mapstructure "github.com/mitchellh/mapstructure"
)

// ErrorObject uses code as discriminator attribute
type ErrorObject interface{}

func mapDiscriminatorErrorObject(input ErrorObject) ErrorObject {
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
		new := InvalidJSONInputError{}
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

// AccessDeniedError implements the interface ErrorObject
type AccessDeniedError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj AccessDeniedError) MarshalJSON() ([]byte, error) {
	type Alias AccessDeniedError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "access_denied", Alias: (*Alias)(&obj)})
}

func (obj AccessDeniedError) Error() string {
	return obj.Message
}

// ConcurrentModificationError implements the interface ErrorObject
type ConcurrentModificationError struct {
	Message        string `json:"message"`
	CurrentVersion int    `json:"currentVersion,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ConcurrentModificationError) MarshalJSON() ([]byte, error) {
	type Alias ConcurrentModificationError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "ConcurrentModification", Alias: (*Alias)(&obj)})
}

func (obj ConcurrentModificationError) Error() string {
	return obj.Message
}

// DiscountCodeNonApplicableError implements the interface ErrorObject
type DiscountCodeNonApplicableError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj DiscountCodeNonApplicableError) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeNonApplicableError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "DiscountCodeNonApplicable", Alias: (*Alias)(&obj)})
}

func (obj DiscountCodeNonApplicableError) Error() string {
	return obj.Message
}

// DuplicateAttributeValueError implements the interface ErrorObject
type DuplicateAttributeValueError struct {
	Message   string     `json:"message"`
	Attribute *Attribute `json:"attribute"`
}

// MarshalJSON override to set the discriminator value
func (obj DuplicateAttributeValueError) MarshalJSON() ([]byte, error) {
	type Alias DuplicateAttributeValueError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "DuplicateAttributeValue", Alias: (*Alias)(&obj)})
}

func (obj DuplicateAttributeValueError) Error() string {
	return obj.Message
}

// DuplicateAttributeValuesError implements the interface ErrorObject
type DuplicateAttributeValuesError struct {
	Message    string      `json:"message"`
	Attributes []Attribute `json:"attributes"`
}

// MarshalJSON override to set the discriminator value
func (obj DuplicateAttributeValuesError) MarshalJSON() ([]byte, error) {
	type Alias DuplicateAttributeValuesError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "DuplicateAttributeValues", Alias: (*Alias)(&obj)})
}

func (obj DuplicateAttributeValuesError) Error() string {
	return obj.Message
}

// DuplicateFieldError implements the interface ErrorObject
type DuplicateFieldError struct {
	Message        string      `json:"message"`
	Field          string      `json:"field,omitempty"`
	DuplicateValue interface{} `json:"duplicateValue,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj DuplicateFieldError) MarshalJSON() ([]byte, error) {
	type Alias DuplicateFieldError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "DuplicateField", Alias: (*Alias)(&obj)})
}

func (obj DuplicateFieldError) Error() string {
	return obj.Message
}

// DuplicatePriceScopeError implements the interface ErrorObject
type DuplicatePriceScopeError struct {
	Message           string  `json:"message"`
	ConflictingPrices []Price `json:"conflictingPrices"`
}

// MarshalJSON override to set the discriminator value
func (obj DuplicatePriceScopeError) MarshalJSON() ([]byte, error) {
	type Alias DuplicatePriceScopeError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "DuplicatePriceScope", Alias: (*Alias)(&obj)})
}

func (obj DuplicatePriceScopeError) Error() string {
	return obj.Message
}

// DuplicateVariantValuesError implements the interface ErrorObject
type DuplicateVariantValuesError struct {
	Message       string         `json:"message"`
	VariantValues *VariantValues `json:"variantValues"`
}

// MarshalJSON override to set the discriminator value
func (obj DuplicateVariantValuesError) MarshalJSON() ([]byte, error) {
	type Alias DuplicateVariantValuesError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "DuplicateVariantValues", Alias: (*Alias)(&obj)})
}

func (obj DuplicateVariantValuesError) Error() string {
	return obj.Message
}

// ErrorResponse is a standalone struct
type ErrorResponse struct {
	StatusCode       int           `json:"statusCode"`
	Message          string        `json:"message"`
	Errors           []ErrorObject `json:"errors,omitempty"`
	ErrorDescription string        `json:"error_description,omitempty"`
	_Error           string        `json:"error,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ErrorResponse) UnmarshalJSON(data []byte) error {
	type Alias ErrorResponse
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Errors {
		obj.Errors[i] = mapDiscriminatorErrorObject(obj.Errors[i])
	}

	return nil
}

func (obj ErrorResponse) Error() string {
	return obj.Message
}

// InsufficientScopeError implements the interface ErrorObject
type InsufficientScopeError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj InsufficientScopeError) MarshalJSON() ([]byte, error) {
	type Alias InsufficientScopeError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "insufficient_scope", Alias: (*Alias)(&obj)})
}

func (obj InsufficientScopeError) Error() string {
	return obj.Message
}

// InvalidCredentialsError implements the interface ErrorObject
type InvalidCredentialsError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj InvalidCredentialsError) MarshalJSON() ([]byte, error) {
	type Alias InvalidCredentialsError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "InvalidCredentials", Alias: (*Alias)(&obj)})
}

func (obj InvalidCredentialsError) Error() string {
	return obj.Message
}

// InvalidCurrentPasswordError implements the interface ErrorObject
type InvalidCurrentPasswordError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj InvalidCurrentPasswordError) MarshalJSON() ([]byte, error) {
	type Alias InvalidCurrentPasswordError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "InvalidCurrentPassword", Alias: (*Alias)(&obj)})
}

func (obj InvalidCurrentPasswordError) Error() string {
	return obj.Message
}

// InvalidFieldError implements the interface ErrorObject
type InvalidFieldError struct {
	Message       string        `json:"message"`
	InvalidValue  interface{}   `json:"invalidValue"`
	Field         string        `json:"field"`
	AllowedValues []interface{} `json:"allowedValues,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj InvalidFieldError) MarshalJSON() ([]byte, error) {
	type Alias InvalidFieldError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "InvalidField", Alias: (*Alias)(&obj)})
}

func (obj InvalidFieldError) Error() string {
	return obj.Message
}

// InvalidInputError implements the interface ErrorObject
type InvalidInputError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj InvalidInputError) MarshalJSON() ([]byte, error) {
	type Alias InvalidInputError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "InvalidInput", Alias: (*Alias)(&obj)})
}

func (obj InvalidInputError) Error() string {
	return obj.Message
}

// InvalidItemShippingDetailsError implements the interface ErrorObject
type InvalidItemShippingDetailsError struct {
	Message string `json:"message"`
	Subject string `json:"subject"`
	ItemID  string `json:"itemId"`
}

// MarshalJSON override to set the discriminator value
func (obj InvalidItemShippingDetailsError) MarshalJSON() ([]byte, error) {
	type Alias InvalidItemShippingDetailsError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "InvalidItemShippingDetails", Alias: (*Alias)(&obj)})
}

func (obj InvalidItemShippingDetailsError) Error() string {
	return obj.Message
}

// InvalidJSONInputError implements the interface ErrorObject
type InvalidJSONInputError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj InvalidJSONInputError) MarshalJSON() ([]byte, error) {
	type Alias InvalidJSONInputError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "InvalidJsonInput", Alias: (*Alias)(&obj)})
}

func (obj InvalidJSONInputError) Error() string {
	return obj.Message
}

// InvalidOperationError implements the interface ErrorObject
type InvalidOperationError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj InvalidOperationError) MarshalJSON() ([]byte, error) {
	type Alias InvalidOperationError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "InvalidOperation", Alias: (*Alias)(&obj)})
}

func (obj InvalidOperationError) Error() string {
	return obj.Message
}

// InvalidSubjectError implements the interface ErrorObject
type InvalidSubjectError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj InvalidSubjectError) MarshalJSON() ([]byte, error) {
	type Alias InvalidSubjectError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "InvalidSubject", Alias: (*Alias)(&obj)})
}

func (obj InvalidSubjectError) Error() string {
	return obj.Message
}

// InvalidTokenError implements the interface ErrorObject
type InvalidTokenError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj InvalidTokenError) MarshalJSON() ([]byte, error) {
	type Alias InvalidTokenError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "invalid_token", Alias: (*Alias)(&obj)})
}

func (obj InvalidTokenError) Error() string {
	return obj.Message
}

// MissingTaxRateForCountryError implements the interface ErrorObject
type MissingTaxRateForCountryError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj MissingTaxRateForCountryError) MarshalJSON() ([]byte, error) {
	type Alias MissingTaxRateForCountryError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "MissingTaxRateForCountry", Alias: (*Alias)(&obj)})
}

func (obj MissingTaxRateForCountryError) Error() string {
	return obj.Message
}

// NoMatchingProductDiscountFoundError implements the interface ErrorObject
type NoMatchingProductDiscountFoundError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj NoMatchingProductDiscountFoundError) MarshalJSON() ([]byte, error) {
	type Alias NoMatchingProductDiscountFoundError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "NoMatchingProductDiscountFound", Alias: (*Alias)(&obj)})
}

func (obj NoMatchingProductDiscountFoundError) Error() string {
	return obj.Message
}

// OutOfStockError implements the interface ErrorObject
type OutOfStockError struct {
	Message   string   `json:"message"`
	Skus      []string `json:"skus"`
	LineItems []string `json:"lineItems"`
}

// MarshalJSON override to set the discriminator value
func (obj OutOfStockError) MarshalJSON() ([]byte, error) {
	type Alias OutOfStockError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "OutOfStock", Alias: (*Alias)(&obj)})
}

func (obj OutOfStockError) Error() string {
	return obj.Message
}

// PriceChangedError implements the interface ErrorObject
type PriceChangedError struct {
	Message   string   `json:"message"`
	Shipping  bool     `json:"shipping"`
	LineItems []string `json:"lineItems"`
}

// MarshalJSON override to set the discriminator value
func (obj PriceChangedError) MarshalJSON() ([]byte, error) {
	type Alias PriceChangedError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "PriceChanged", Alias: (*Alias)(&obj)})
}

func (obj PriceChangedError) Error() string {
	return obj.Message
}

// ReferenceExistsError implements the interface ErrorObject
type ReferenceExistsError struct {
	Message      string          `json:"message"`
	ReferencedBy ReferenceTypeID `json:"referencedBy,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ReferenceExistsError) MarshalJSON() ([]byte, error) {
	type Alias ReferenceExistsError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "ReferenceExists", Alias: (*Alias)(&obj)})
}

func (obj ReferenceExistsError) Error() string {
	return obj.Message
}

// RequiredFieldError implements the interface ErrorObject
type RequiredFieldError struct {
	Message string `json:"message"`
	Field   string `json:"field"`
}

// MarshalJSON override to set the discriminator value
func (obj RequiredFieldError) MarshalJSON() ([]byte, error) {
	type Alias RequiredFieldError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "RequiredField", Alias: (*Alias)(&obj)})
}

func (obj RequiredFieldError) Error() string {
	return obj.Message
}

// ResourceNotFoundError implements the interface ErrorObject
type ResourceNotFoundError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj ResourceNotFoundError) MarshalJSON() ([]byte, error) {
	type Alias ResourceNotFoundError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "ResourceNotFound", Alias: (*Alias)(&obj)})
}

func (obj ResourceNotFoundError) Error() string {
	return obj.Message
}

// ShippingMethodDoesNotMatchCartError implements the interface ErrorObject
type ShippingMethodDoesNotMatchCartError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodDoesNotMatchCartError) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodDoesNotMatchCartError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "ShippingMethodDoesNotMatchCart", Alias: (*Alias)(&obj)})
}

func (obj ShippingMethodDoesNotMatchCartError) Error() string {
	return obj.Message
}

// VariantValues is a standalone struct
type VariantValues struct {
	SKU        string      `json:"sku,omitempty"`
	Prices     []Price     `json:"prices"`
	Attributes []Attribute `json:"attributes"`
}
