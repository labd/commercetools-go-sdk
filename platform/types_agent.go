package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"io"
)

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when the customer is associated with multiple [Business Units](ctp:api:type:BusinessUnit) and none could be automatically selected. To resolve this error, specify an explicit `businessUnit` reference in the request.
*
 */
type AgentBusinessUnitAmbiguousError struct {
	// Plain text description of the error.
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AgentBusinessUnitAmbiguousError) UnmarshalJSON(data []byte) error {
	type Alias AgentBusinessUnitAmbiguousError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AgentBusinessUnitAmbiguousError) MarshalJSON() ([]byte, error) {
	type Alias AgentBusinessUnitAmbiguousError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "BusinessUnitAmbiguous", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *AgentBusinessUnitAmbiguousError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}

func (obj AgentBusinessUnitAmbiguousError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown AgentBusinessUnitAmbiguousError: failed to parse error response"
}

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when the customer is associated with more [Business Units](ctp:api:type:BusinessUnit) than can be resolved automatically. To resolve this error, specify an explicit `businessUnit` reference in the request.
*
 */
type AgentBusinessUnitLimitExceededError struct {
	// Plain text description of the error.
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AgentBusinessUnitLimitExceededError) UnmarshalJSON(data []byte) error {
	type Alias AgentBusinessUnitLimitExceededError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AgentBusinessUnitLimitExceededError) MarshalJSON() ([]byte, error) {
	type Alias AgentBusinessUnitLimitExceededError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "BusinessUnitLimitExceeded", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *AgentBusinessUnitLimitExceededError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}

func (obj AgentBusinessUnitLimitExceededError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown AgentBusinessUnitLimitExceededError: failed to parse error response"
}

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when the [BusinessUnit](ctp:api:type:BusinessUnit) could not be resolved from the input. Pin it with an explicit `businessUnit` reference in the request.
*
 */
type AgentBusinessUnitUnresolvedError struct {
	// Plain text description of the error.
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AgentBusinessUnitUnresolvedError) UnmarshalJSON(data []byte) error {
	type Alias AgentBusinessUnitUnresolvedError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AgentBusinessUnitUnresolvedError) MarshalJSON() ([]byte, error) {
	type Alias AgentBusinessUnitUnresolvedError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "BusinessUnitUnresolved", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *AgentBusinessUnitUnresolvedError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}

func (obj AgentBusinessUnitUnresolvedError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown AgentBusinessUnitUnresolvedError: failed to parse error response"
}

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when the input could not be processed due to an internal error.
*
 */
type AgentExtractionFailedError struct {
	// Plain text description of the error.
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AgentExtractionFailedError) UnmarshalJSON(data []byte) error {
	type Alias AgentExtractionFailedError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AgentExtractionFailedError) MarshalJSON() ([]byte, error) {
	type Alias AgentExtractionFailedError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ExtractionFailed", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *AgentExtractionFailedError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}

func (obj AgentExtractionFailedError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown AgentExtractionFailedError: failed to parse error response"
}

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when Intake is not enabled for the Project.
*
 */
type AgentFeatureDisabledError struct {
	// Plain text description of the error.
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AgentFeatureDisabledError) UnmarshalJSON(data []byte) error {
	type Alias AgentFeatureDisabledError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AgentFeatureDisabledError) MarshalJSON() ([]byte, error) {
	type Alias AgentFeatureDisabledError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "FeatureDisabled", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *AgentFeatureDisabledError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}

func (obj AgentFeatureDisabledError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown AgentFeatureDisabledError: failed to parse error response"
}

/**
*	Returned when an uploaded file in an Intake [API request](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) could not be parsed.
*
 */
type AgentFileNotProcessedWarning struct {
	// Plain text description of why the file could not be parsed.
	Message string `json:"message"`
	// Name of the file that failed to parse.
	FileName string `json:"fileName"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AgentFileNotProcessedWarning) MarshalJSON() ([]byte, error) {
	type Alias AgentFileNotProcessedWarning
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "FileNotProcessed", Alias: (*Alias)(&obj)})
}

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request with `outputType` set to `Cart` or `QuoteRequest` when a shipping country was not provided and could not be inferred from the input.
*
 */
type AgentMissingCountryError struct {
	// Plain text description of the error.
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AgentMissingCountryError) UnmarshalJSON(data []byte) error {
	type Alias AgentMissingCountryError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AgentMissingCountryError) MarshalJSON() ([]byte, error) {
	type Alias AgentMissingCountryError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "MissingCountry", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *AgentMissingCountryError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}

func (obj AgentMissingCountryError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown AgentMissingCountryError: failed to parse error response"
}

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when the customer's email address was not provided and could not be inferred from the input.
*
 */
type AgentMissingCustomerEmailError struct {
	// Plain text description of the error.
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AgentMissingCustomerEmailError) UnmarshalJSON(data []byte) error {
	type Alias AgentMissingCustomerEmailError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AgentMissingCustomerEmailError) MarshalJSON() ([]byte, error) {
	type Alias AgentMissingCustomerEmailError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "MissingCustomerEmail", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *AgentMissingCustomerEmailError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}

func (obj AgentMissingCustomerEmailError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown AgentMissingCustomerEmailError: failed to parse error response"
}

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when the input did not specify whether to create a Cart, a Quote Request, or a Shopping List. Set `outputType` explicitly in the request.
*
 */
type AgentMissingEntityTypeError struct {
	// Plain text description of the error.
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AgentMissingEntityTypeError) UnmarshalJSON(data []byte) error {
	type Alias AgentMissingEntityTypeError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AgentMissingEntityTypeError) MarshalJSON() ([]byte, error) {
	type Alias AgentMissingEntityTypeError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "MissingEntityType", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *AgentMissingEntityTypeError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}

func (obj AgentMissingEntityTypeError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown AgentMissingEntityTypeError: failed to parse error response"
}

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request with `outputType` set to `ShoppingList` when a name for the Shopping List is not present in the input. A name is never invented.
*
 */
type AgentMissingShoppingListNameError struct {
	// Plain text description of the error.
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AgentMissingShoppingListNameError) UnmarshalJSON(data []byte) error {
	type Alias AgentMissingShoppingListNameError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AgentMissingShoppingListNameError) MarshalJSON() ([]byte, error) {
	type Alias AgentMissingShoppingListNameError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "MissingShoppingListName", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *AgentMissingShoppingListNameError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}

func (obj AgentMissingShoppingListNameError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown AgentMissingShoppingListNameError: failed to parse error response"
}

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when no Line Items could be extracted from the provided `prompt` or files.
*
 */
type AgentNoLineItemsExtractedError struct {
	// Plain text description of the error.
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AgentNoLineItemsExtractedError) UnmarshalJSON(data []byte) error {
	type Alias AgentNoLineItemsExtractedError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AgentNoLineItemsExtractedError) MarshalJSON() ([]byte, error) {
	type Alias AgentNoLineItemsExtractedError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "NoLineItemsExtracted", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *AgentNoLineItemsExtractedError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}

func (obj AgentNoLineItemsExtractedError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown AgentNoLineItemsExtractedError: failed to parse error response"
}

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when the input does not contain enough information to create the requested entity. For a Cart or Quote Request, no products, no customer email address, and no shipping country could be determined. For a Shopping List, no products, no customer email address, and no name could be determined.
*
 */
type AgentOutOfScopeError struct {
	// Plain text description of the error.
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AgentOutOfScopeError) UnmarshalJSON(data []byte) error {
	type Alias AgentOutOfScopeError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AgentOutOfScopeError) MarshalJSON() ([]byte, error) {
	type Alias AgentOutOfScopeError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "OutOfScope", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *AgentOutOfScopeError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}

func (obj AgentOutOfScopeError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown AgentOutOfScopeError: failed to parse error response"
}

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when Product Search is not enabled for the project.
*
 */
type AgentProductSearchNotEnabledError struct {
	// Plain text description of the error.
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AgentProductSearchNotEnabledError) UnmarshalJSON(data []byte) error {
	type Alias AgentProductSearchNotEnabledError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AgentProductSearchNotEnabledError) MarshalJSON() ([]byte, error) {
	type Alias AgentProductSearchNotEnabledError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ProductSearchNotEnabled", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *AgentProductSearchNotEnabledError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}

func (obj AgentProductSearchNotEnabledError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown AgentProductSearchNotEnabledError: failed to parse error response"
}

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when none of the requested products could be matched to the catalog.
*
 */
type AgentProductsNotFoundError struct {
	// Plain text description of the error.
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AgentProductsNotFoundError) UnmarshalJSON(data []byte) error {
	type Alias AgentProductsNotFoundError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AgentProductsNotFoundError) MarshalJSON() ([]byte, error) {
	type Alias AgentProductsNotFoundError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ProductsNotFound", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *AgentProductsNotFoundError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}

func (obj AgentProductsNotFoundError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown AgentProductsNotFoundError: failed to parse error response"
}

/**
*	Returned when one or more Products referenced in an Intake [API request](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) could not be matched to the catalog or are unavailable in the bound Store. The omitted Products are excluded from the created entity.
*
 */
type AgentProductsNotFoundWarning struct {
	// Plain text description of the omitted Products.
	Message string `json:"message"`
	// Identifiers of the Products that could not be matched to the catalog or are unavailable in the bound Store, as they appeared in the input.
	Products []string `json:"products"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AgentProductsNotFoundWarning) MarshalJSON() ([]byte, error) {
	type Alias AgentProductsNotFoundWarning
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ProductsNotFound", Alias: (*Alias)(&obj)})
}

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when the underlying [Cart](ctp:api:type:Cart) was created but the [QuoteRequest](ctp:api:type:QuoteRequest) could not be created from it, for example because the Cart has no verified Customer.
*
 */
type AgentQuoteRequestCreationFailedError struct {
	// Plain text description of the error.
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// `id` of the orphaned Cart, for retry or cleanup.
	CartId string `json:"cartId"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AgentQuoteRequestCreationFailedError) UnmarshalJSON(data []byte) error {
	type Alias AgentQuoteRequestCreationFailedError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "cartId")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AgentQuoteRequestCreationFailedError) MarshalJSON() ([]byte, error) {
	type Alias AgentQuoteRequestCreationFailedError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "QuoteRequestCreationFailed", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *AgentQuoteRequestCreationFailedError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}

func (obj AgentQuoteRequestCreationFailedError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown AgentQuoteRequestCreationFailedError: failed to parse error response"
}

/**
*	Authentication or authorization failure from a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request. Uses the standard [AuthErrorResponse](ctp:api:type:AuthErrorResponse) shape (`statusCode`, `message`, `errors`, plus OAuth `error` / `error_description`) rather than the [AgentResponsesErrorResponse](ctp:api:type:AgentResponsesErrorResponse) envelope used for other errors from this endpoint.
*
 */
type AgentResponsesAuthError struct {
	// HTTP status code corresponding to the error.
	StatusCode int `json:"statusCode"`
	// First error message in the `errors` array.
	Message string `json:"message"`
	// Authentication and authorization-related errors returned for a request.
	Errors []ErrorObject `json:"errors"`
	// Error code as per the [OAuth 2.0 specification](https://datatracker.ietf.org/doc/html/rfc6749#section-5.2). For example: `"access_denied"`.
	ErrorMessage string `json:"error"`
	// Plain text description of the first error.
	ErrorDescription *string `json:"error_description,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AgentResponsesAuthError) UnmarshalJSON(data []byte) error {
	type Alias AgentResponsesAuthError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Errors {
		var err error
		obj.Errors[i], err = mapDiscriminatorErrorObject(obj.Errors[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (obj AgentResponsesAuthError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown AgentResponsesAuthError: failed to parse error response"
}

/**
*	Error response from a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request. Extends [ErrorResponse](ctp:api:type:ErrorResponse) with a `threadId` for support.
*
 */
type AgentResponsesErrorResponse struct {
	// HTTP status code corresponding to the error.
	StatusCode int `json:"statusCode"`
	// First error message in the `errors` array.
	Message string `json:"message"`
	// Errors returned for a request.
	//
	// A single error response can contain multiple errors if the errors are related to the same HTTP status code such as `400`.
	Errors []ErrorObject `json:"errors"`
	// Identifier of the workflow run that raised the error. Present on errors raised once the workflow has started. Absent on errors raised while parsing the request, before the workflow starts.
	ThreadId *string `json:"threadId,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AgentResponsesErrorResponse) UnmarshalJSON(data []byte) error {
	type Alias AgentResponsesErrorResponse
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Errors {
		var err error
		obj.Errors[i], err = mapDiscriminatorErrorObject(obj.Errors[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (obj AgentResponsesErrorResponse) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown AgentResponsesErrorResponse: failed to parse error response"
}

/**
*	`multipart/form-data` body for a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when files are attached. The JSON fields go in the `payload` part; files go in `files` parts.
*
 */
type AgentResponsesMultipartRequest struct {
	// JSON document matching [AgentResponsesPayload](ctp:api:type:AgentResponsesPayload). Must include `outputType` and `locale`. At least one of `prompt` or `files` must be provided.
	Payload AgentResponsesPayload `json:"payload"`
	// Supporting files. Repeat this part for each file. Maximum of 3 files, 20 MB each. Allowed content types: `application/pdf`, `text/plain`, `text/csv`, `application/vnd.openxmlformats-officedocument.wordprocessingml.document`, `application/vnd.openxmlformats-officedocument.spreadsheetml.sheet`, `message/rfc822`, `text/markdown`, `text/html`.
	Files []io.Reader `json:"files"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AgentResponsesMultipartRequest) MarshalJSON() ([]byte, error) {
	type Alias AgentResponsesMultipartRequest
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["files"] == nil {
		delete(raw, "files")
	}

	return json.Marshal(raw)

}

/**
*	Entity that Intake creates from a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request. Never inferred from the input.
*
 */
type AgentResponsesOutputType string

const (
	AgentResponsesOutputTypeCart         AgentResponsesOutputType = "Cart"
	AgentResponsesOutputTypeQuoteRequest AgentResponsesOutputType = "QuoteRequest"
	AgentResponsesOutputTypeShoppingList AgentResponsesOutputType = "ShoppingList"
)

/**
*	Structured fields for the `payload` part of a `multipart/form-data` [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request to Intake. `prompt` is optional when files are attached.
*
*	For an `application/json` request, use [AgentResponsesRequest](ctp:api:type:AgentResponsesRequest) instead.
*
 */
type AgentResponsesPayload struct {
	// The entity to create.
	OutputType AgentResponsesOutputType `json:"outputType"`
	// A valid BCP 47 language tag, for example `en-US`.
	Locale string `json:"locale"`
	// Natural-language description of the entity to create, such as the body of an email or a note. Required if no file is attached. For a Shopping List, the name of the list must be stated in either the prompt or an attached file.
	Prompt *string `json:"prompt,omitempty"`
	// [BusinessUnit](ctp:api:type:BusinessUnit) to use for the request. If omitted, Intake resolves it from the input.
	BusinessUnit *BusinessUnitResourceIdentifier `json:"businessUnit,omitempty"`
	// [Store](ctp:api:type:Store) to use for the request. If omitted, Intake resolves it from the input.
	Store *StoreResourceIdentifier `json:"store,omitempty"`
}

/**
*	JSON body for an `application/json` [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request to Intake. `prompt` is required.
*
*	To attach files, send a `multipart/form-data` request with an [AgentResponsesPayload](ctp:api:type:AgentResponsesPayload) as the `payload` part.
*
 */
type AgentResponsesRequest struct {
	// The entity to create.
	OutputType AgentResponsesOutputType `json:"outputType"`
	// A valid BCP 47 language tag, for example `en-US`.
	Locale string `json:"locale"`
	// Natural-language description of the entity to create, such as the body of an email or a note. For a Shopping List, it must state the name of the list.
	Prompt string `json:"prompt"`
	// [BusinessUnit](ctp:api:type:BusinessUnit) to use for the request. If omitted, Intake resolves it from `prompt`.
	BusinessUnit *BusinessUnitResourceIdentifier `json:"businessUnit,omitempty"`
	// [Store](ctp:api:type:Store) to use for the request. If omitted, Intake resolves it from `prompt`.
	Store *StoreResourceIdentifier `json:"store,omitempty"`
}

/**
*	Successful `201` response from a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request. Returns the created entity in commercetools REST representation, together with any non-fatal `warnings`.
*
*	The concrete type depends on `entityType`:
*
*	- [AgentResponsesCartSuccess](ctp:api:type:AgentResponsesCartSuccess) when a [Cart](ctp:api:type:Cart) was created
*	- [AgentResponsesQuoteRequestSuccess](ctp:api:type:AgentResponsesQuoteRequestSuccess) when a [QuoteRequest](ctp:api:type:QuoteRequest) was created
*	- [AgentResponsesShoppingListSuccess](ctp:api:type:AgentResponsesShoppingListSuccess) when a [ShoppingList](ctp:api:type:ShoppingList) was created
*
 */
type AgentResponsesSuccess interface{}

func mapDiscriminatorAgentResponsesSuccess(input interface{}) (AgentResponsesSuccess, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["entityType"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'entityType'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "Cart":
		obj := AgentResponsesCartSuccess{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		for i := range obj.Warnings {
			var err error
			obj.Warnings[i], err = mapDiscriminatorWarningObject(obj.Warnings[i])
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "QuoteRequest":
		obj := AgentResponsesQuoteRequestSuccess{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		for i := range obj.Warnings {
			var err error
			obj.Warnings[i], err = mapDiscriminatorWarningObject(obj.Warnings[i])
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ShoppingList":
		obj := AgentResponsesShoppingListSuccess{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		for i := range obj.Warnings {
			var err error
			obj.Warnings[i], err = mapDiscriminatorWarningObject(obj.Warnings[i])
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Successful `201` response from a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when `outputType` is `Cart`.
*
 */
type AgentResponsesCartSuccess struct {
	// Non-fatal issues encountered while processing the request. Present only when at least one warning is returned.
	Warnings []WarningObject `json:"warnings"`
	// Identifier of the workflow run that produced this response.
	ThreadId string `json:"threadId"`
	// The created [Cart](ctp:api:type:Cart) in full commercetools REST representation.
	Entity Cart `json:"entity"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AgentResponsesCartSuccess) UnmarshalJSON(data []byte) error {
	type Alias AgentResponsesCartSuccess
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Warnings {
		var err error
		obj.Warnings[i], err = mapDiscriminatorWarningObject(obj.Warnings[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AgentResponsesCartSuccess) MarshalJSON() ([]byte, error) {
	type Alias AgentResponsesCartSuccess
	data, err := json.Marshal(struct {
		Action string `json:"entityType"`
		*Alias
	}{Action: "Cart", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["warnings"] == nil {
		delete(raw, "warnings")
	}

	return json.Marshal(raw)

}

/**
*	Successful `201` response from a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when `outputType` is `QuoteRequest`.
*
 */
type AgentResponsesQuoteRequestSuccess struct {
	// Non-fatal issues encountered while processing the request. Present only when at least one warning is returned.
	Warnings []WarningObject `json:"warnings"`
	// Identifier of the workflow run that produced this response.
	ThreadId string `json:"threadId"`
	// The created [QuoteRequest](ctp:api:type:QuoteRequest) in full commercetools REST representation.
	Entity QuoteRequest `json:"entity"`
	// `id` of the underlying [Cart](ctp:api:type:Cart).
	CartId string `json:"cartId"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AgentResponsesQuoteRequestSuccess) UnmarshalJSON(data []byte) error {
	type Alias AgentResponsesQuoteRequestSuccess
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Warnings {
		var err error
		obj.Warnings[i], err = mapDiscriminatorWarningObject(obj.Warnings[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AgentResponsesQuoteRequestSuccess) MarshalJSON() ([]byte, error) {
	type Alias AgentResponsesQuoteRequestSuccess
	data, err := json.Marshal(struct {
		Action string `json:"entityType"`
		*Alias
	}{Action: "QuoteRequest", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["warnings"] == nil {
		delete(raw, "warnings")
	}

	return json.Marshal(raw)

}

/**
*	Successful `201` response from a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when `outputType` is `ShoppingList`.
*
 */
type AgentResponsesShoppingListSuccess struct {
	// Non-fatal issues encountered while processing the request. Present only when at least one warning is returned.
	Warnings []WarningObject `json:"warnings"`
	// Identifier of the workflow run that produced this response.
	ThreadId string `json:"threadId"`
	// The created [ShoppingList](ctp:api:type:ShoppingList) in full commercetools REST representation.
	Entity ShoppingList `json:"entity"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AgentResponsesShoppingListSuccess) UnmarshalJSON(data []byte) error {
	type Alias AgentResponsesShoppingListSuccess
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Warnings {
		var err error
		obj.Warnings[i], err = mapDiscriminatorWarningObject(obj.Warnings[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AgentResponsesShoppingListSuccess) MarshalJSON() ([]byte, error) {
	type Alias AgentResponsesShoppingListSuccess
	data, err := json.Marshal(struct {
		Action string `json:"entityType"`
		*Alias
	}{Action: "ShoppingList", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["warnings"] == nil {
		delete(raw, "warnings")
	}

	return json.Marshal(raw)

}

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when the [ShoppingList](ctp:api:type:ShoppingList) could not be created, for example because the request has no verified Customer or because commercetools rejected the ShoppingList draft. No entity is created, so there is nothing to clean up.
*
 */
type AgentShoppingListCreationFailedError struct {
	// Plain text description of the error.
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AgentShoppingListCreationFailedError) UnmarshalJSON(data []byte) error {
	type Alias AgentShoppingListCreationFailedError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AgentShoppingListCreationFailedError) MarshalJSON() ([]byte, error) {
	type Alias AgentShoppingListCreationFailedError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ShoppingListCreationFailed", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *AgentShoppingListCreationFailedError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}

func (obj AgentShoppingListCreationFailedError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown AgentShoppingListCreationFailedError: failed to parse error response"
}

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when the resolved [Business Unit](ctp:api:type:BusinessUnit) grants multiple [Stores](ctp:api:type:Store) and none could be automatically selected. To resolve this error, specify an explicit `store` reference in the request.
*
 */
type AgentStoreAmbiguousError struct {
	// Plain text description of the error.
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AgentStoreAmbiguousError) UnmarshalJSON(data []byte) error {
	type Alias AgentStoreAmbiguousError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AgentStoreAmbiguousError) MarshalJSON() ([]byte, error) {
	type Alias AgentStoreAmbiguousError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "StoreAmbiguous", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *AgentStoreAmbiguousError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}

func (obj AgentStoreAmbiguousError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown AgentStoreAmbiguousError: failed to parse error response"
}

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when the resolved [Store](ctp:api:type:Store) has more than one entry in the `distributionChannels` array, which is not supported. To resolve this error, choose a Store with zero or one distribution channel.
*
 */
type AgentStoreDistributionChannelsUnsupportedError struct {
	// Plain text description of the error.
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AgentStoreDistributionChannelsUnsupportedError) UnmarshalJSON(data []byte) error {
	type Alias AgentStoreDistributionChannelsUnsupportedError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AgentStoreDistributionChannelsUnsupportedError) MarshalJSON() ([]byte, error) {
	type Alias AgentStoreDistributionChannelsUnsupportedError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "StoreDistributionChannelsUnsupported", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *AgentStoreDistributionChannelsUnsupportedError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}

func (obj AgentStoreDistributionChannelsUnsupportedError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown AgentStoreDistributionChannelsUnsupportedError: failed to parse error response"
}

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when the [Store](ctp:api:type:Store) could not be resolved from the input. Pin it with an explicit `store` reference in the request.
*
 */
type AgentStoreUnresolvedError struct {
	// Plain text description of the error.
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AgentStoreUnresolvedError) UnmarshalJSON(data []byte) error {
	type Alias AgentStoreUnresolvedError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AgentStoreUnresolvedError) MarshalJSON() ([]byte, error) {
	type Alias AgentStoreUnresolvedError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "StoreUnresolved", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *AgentStoreUnresolvedError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}

func (obj AgentStoreUnresolvedError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown AgentStoreUnresolvedError: failed to parse error response"
}

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when the customer is associated with multiple [Business Units](ctp:api:type:BusinessUnit) and none could be automatically selected. To resolve this error, specify an explicit `businessUnit` reference in the request.
*
 */
type GraphQLAgentBusinessUnitAmbiguousError struct {
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *GraphQLAgentBusinessUnitAmbiguousError) UnmarshalJSON(data []byte) error {
	type Alias GraphQLAgentBusinessUnitAmbiguousError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj GraphQLAgentBusinessUnitAmbiguousError) MarshalJSON() ([]byte, error) {
	type Alias GraphQLAgentBusinessUnitAmbiguousError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "BusinessUnitAmbiguous", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *GraphQLAgentBusinessUnitAmbiguousError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when the customer is associated with more [Business Units](ctp:api:type:BusinessUnit) than can be resolved automatically. To resolve this error, specify an explicit `businessUnit` reference in the request.
*
 */
type GraphQLAgentBusinessUnitLimitExceededError struct {
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *GraphQLAgentBusinessUnitLimitExceededError) UnmarshalJSON(data []byte) error {
	type Alias GraphQLAgentBusinessUnitLimitExceededError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj GraphQLAgentBusinessUnitLimitExceededError) MarshalJSON() ([]byte, error) {
	type Alias GraphQLAgentBusinessUnitLimitExceededError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "BusinessUnitLimitExceeded", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *GraphQLAgentBusinessUnitLimitExceededError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when the [BusinessUnit](ctp:api:type:BusinessUnit) could not be resolved from the input. Pin it with an explicit `businessUnit` reference in the request.
*
 */
type GraphQLAgentBusinessUnitUnresolvedError struct {
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *GraphQLAgentBusinessUnitUnresolvedError) UnmarshalJSON(data []byte) error {
	type Alias GraphQLAgentBusinessUnitUnresolvedError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj GraphQLAgentBusinessUnitUnresolvedError) MarshalJSON() ([]byte, error) {
	type Alias GraphQLAgentBusinessUnitUnresolvedError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "BusinessUnitUnresolved", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *GraphQLAgentBusinessUnitUnresolvedError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when the input could not be processed due to an internal error.
*
 */
type GraphQLAgentExtractionFailedError struct {
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *GraphQLAgentExtractionFailedError) UnmarshalJSON(data []byte) error {
	type Alias GraphQLAgentExtractionFailedError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj GraphQLAgentExtractionFailedError) MarshalJSON() ([]byte, error) {
	type Alias GraphQLAgentExtractionFailedError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ExtractionFailed", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *GraphQLAgentExtractionFailedError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when Intake is not enabled for the Project.
*
 */
type GraphQLAgentFeatureDisabledError struct {
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *GraphQLAgentFeatureDisabledError) UnmarshalJSON(data []byte) error {
	type Alias GraphQLAgentFeatureDisabledError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj GraphQLAgentFeatureDisabledError) MarshalJSON() ([]byte, error) {
	type Alias GraphQLAgentFeatureDisabledError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "FeatureDisabled", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *GraphQLAgentFeatureDisabledError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request with `outputType` set to `Cart` or `QuoteRequest` when a shipping country was not provided and could not be inferred from the input.
*
 */
type GraphQLAgentMissingCountryError struct {
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *GraphQLAgentMissingCountryError) UnmarshalJSON(data []byte) error {
	type Alias GraphQLAgentMissingCountryError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj GraphQLAgentMissingCountryError) MarshalJSON() ([]byte, error) {
	type Alias GraphQLAgentMissingCountryError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "MissingCountry", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *GraphQLAgentMissingCountryError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when the customer's email address was not provided and could not be inferred from the input.
*
 */
type GraphQLAgentMissingCustomerEmailError struct {
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *GraphQLAgentMissingCustomerEmailError) UnmarshalJSON(data []byte) error {
	type Alias GraphQLAgentMissingCustomerEmailError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj GraphQLAgentMissingCustomerEmailError) MarshalJSON() ([]byte, error) {
	type Alias GraphQLAgentMissingCustomerEmailError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "MissingCustomerEmail", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *GraphQLAgentMissingCustomerEmailError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when the input did not specify whether to create a Cart, a Quote Request, or a Shopping List. Set `outputType` explicitly in the request.
*
 */
type GraphQLAgentMissingEntityTypeError struct {
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *GraphQLAgentMissingEntityTypeError) UnmarshalJSON(data []byte) error {
	type Alias GraphQLAgentMissingEntityTypeError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj GraphQLAgentMissingEntityTypeError) MarshalJSON() ([]byte, error) {
	type Alias GraphQLAgentMissingEntityTypeError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "MissingEntityType", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *GraphQLAgentMissingEntityTypeError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request with `outputType` set to `ShoppingList` when a name for the Shopping List is not present in the input. A name is never invented.
*
 */
type GraphQLAgentMissingShoppingListNameError struct {
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *GraphQLAgentMissingShoppingListNameError) UnmarshalJSON(data []byte) error {
	type Alias GraphQLAgentMissingShoppingListNameError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj GraphQLAgentMissingShoppingListNameError) MarshalJSON() ([]byte, error) {
	type Alias GraphQLAgentMissingShoppingListNameError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "MissingShoppingListName", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *GraphQLAgentMissingShoppingListNameError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when no Line Items could be extracted from the provided `prompt` or files.
*
 */
type GraphQLAgentNoLineItemsExtractedError struct {
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *GraphQLAgentNoLineItemsExtractedError) UnmarshalJSON(data []byte) error {
	type Alias GraphQLAgentNoLineItemsExtractedError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj GraphQLAgentNoLineItemsExtractedError) MarshalJSON() ([]byte, error) {
	type Alias GraphQLAgentNoLineItemsExtractedError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "NoLineItemsExtracted", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *GraphQLAgentNoLineItemsExtractedError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when the input does not contain enough information to create the requested entity. For a Cart or Quote Request, no products, no customer email address, and no shipping country could be determined. For a Shopping List, no products, no customer email address, and no name could be determined.
*
 */
type GraphQLAgentOutOfScopeError struct {
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *GraphQLAgentOutOfScopeError) UnmarshalJSON(data []byte) error {
	type Alias GraphQLAgentOutOfScopeError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj GraphQLAgentOutOfScopeError) MarshalJSON() ([]byte, error) {
	type Alias GraphQLAgentOutOfScopeError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "OutOfScope", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *GraphQLAgentOutOfScopeError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when Product Search is not enabled for the project.
*
 */
type GraphQLAgentProductSearchNotEnabledError struct {
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *GraphQLAgentProductSearchNotEnabledError) UnmarshalJSON(data []byte) error {
	type Alias GraphQLAgentProductSearchNotEnabledError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj GraphQLAgentProductSearchNotEnabledError) MarshalJSON() ([]byte, error) {
	type Alias GraphQLAgentProductSearchNotEnabledError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ProductSearchNotEnabled", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *GraphQLAgentProductSearchNotEnabledError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when none of the requested products could be matched to the catalog.
*
 */
type GraphQLAgentProductsNotFoundError struct {
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *GraphQLAgentProductsNotFoundError) UnmarshalJSON(data []byte) error {
	type Alias GraphQLAgentProductsNotFoundError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj GraphQLAgentProductsNotFoundError) MarshalJSON() ([]byte, error) {
	type Alias GraphQLAgentProductsNotFoundError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ProductsNotFound", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *GraphQLAgentProductsNotFoundError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when the underlying [Cart](ctp:api:type:Cart) was created but the [QuoteRequest](ctp:api:type:QuoteRequest) could not be created from it, for example because the Cart has no verified Customer.
*
 */
type GraphQLAgentQuoteRequestCreationFailedError struct {
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// `id` of the orphaned Cart, for retry or cleanup.
	CartId string `json:"cartId"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *GraphQLAgentQuoteRequestCreationFailedError) UnmarshalJSON(data []byte) error {
	type Alias GraphQLAgentQuoteRequestCreationFailedError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "cartId")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj GraphQLAgentQuoteRequestCreationFailedError) MarshalJSON() ([]byte, error) {
	type Alias GraphQLAgentQuoteRequestCreationFailedError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "QuoteRequestCreationFailed", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *GraphQLAgentQuoteRequestCreationFailedError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when the [ShoppingList](ctp:api:type:ShoppingList) could not be created, for example because the request has no verified Customer or because commercetools rejected the ShoppingList draft. No entity is created, so there is nothing to clean up.
*
 */
type GraphQLAgentShoppingListCreationFailedError struct {
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *GraphQLAgentShoppingListCreationFailedError) UnmarshalJSON(data []byte) error {
	type Alias GraphQLAgentShoppingListCreationFailedError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj GraphQLAgentShoppingListCreationFailedError) MarshalJSON() ([]byte, error) {
	type Alias GraphQLAgentShoppingListCreationFailedError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ShoppingListCreationFailed", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *GraphQLAgentShoppingListCreationFailedError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when the resolved [Business Unit](ctp:api:type:BusinessUnit) grants multiple [Stores](ctp:api:type:Store) and none could be automatically selected. To resolve this error, specify an explicit `store` reference in the request.
*
 */
type GraphQLAgentStoreAmbiguousError struct {
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *GraphQLAgentStoreAmbiguousError) UnmarshalJSON(data []byte) error {
	type Alias GraphQLAgentStoreAmbiguousError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj GraphQLAgentStoreAmbiguousError) MarshalJSON() ([]byte, error) {
	type Alias GraphQLAgentStoreAmbiguousError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "StoreAmbiguous", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *GraphQLAgentStoreAmbiguousError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when the resolved [Store](ctp:api:type:Store) has more than one entry in the `distributionChannels` array, which is not supported. To resolve this error, choose a Store with zero or one distribution channel.
*
 */
type GraphQLAgentStoreDistributionChannelsUnsupportedError struct {
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *GraphQLAgentStoreDistributionChannelsUnsupportedError) UnmarshalJSON(data []byte) error {
	type Alias GraphQLAgentStoreDistributionChannelsUnsupportedError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj GraphQLAgentStoreDistributionChannelsUnsupportedError) MarshalJSON() ([]byte, error) {
	type Alias GraphQLAgentStoreDistributionChannelsUnsupportedError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "StoreDistributionChannelsUnsupported", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *GraphQLAgentStoreDistributionChannelsUnsupportedError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}

/**
*	Returned by a [/responses](/api/projects/intake#create-a-cart-quote-request-or-shopping-list) request when the [Store](ctp:api:type:Store) could not be resolved from the input. Pin it with an explicit `store` reference in the request.
*
 */
type GraphQLAgentStoreUnresolvedError struct {
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *GraphQLAgentStoreUnresolvedError) UnmarshalJSON(data []byte) error {
	type Alias GraphQLAgentStoreUnresolvedError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj GraphQLAgentStoreUnresolvedError) MarshalJSON() ([]byte, error) {
	type Alias GraphQLAgentStoreUnresolvedError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "StoreUnresolved", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *GraphQLAgentStoreUnresolvedError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}
