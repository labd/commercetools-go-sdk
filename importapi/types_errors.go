package importapi

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
)

/**
*	The response in case of an error.
*
 */
type ErrorResponse struct {
	// The http status code of the response.
	StatusCode int `json:"statusCode"`
	// Describes the error.
	Message string `json:"message"`
	// This property is only used for OAuth2 errors.
	// Contains the error code.
	ErrorMessage *string `json:"error,omitempty"`
	// This property is only used for OAuth2 errors.
	// Additional information to assist the client developer in
	// understanding the error.
	ErrorDescription *string `json:"error_description,omitempty"`
	// The errors that caused this error response.
	Errors []ErrorObject `json:"errors"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ErrorResponse) UnmarshalJSON(data []byte) error {
	type Alias ErrorResponse
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ErrorResponse) MarshalJSON() ([]byte, error) {
	type Alias ErrorResponse
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

	if raw["errors"] == nil {
		delete(raw, "errors")
	}

	return json.Marshal(raw)

}

func (obj ErrorResponse) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ErrorResponse: failed to parse error response"
}

type ErrorObject interface{}

func mapDiscriminatorErrorObject(input interface{}) (ErrorObject, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["code"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'code'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "access_denied":
		obj := AccessDeniedError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "invalid_scope":
		obj := InvalidScopeError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "InvalidOperation":
		obj := InvalidOperation{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "DuplicateAttributeValue":
		obj := DuplicateAttributeValueError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Attribute != nil {
			var err error
			obj.Attribute, err = mapDiscriminatorAttribute(obj.Attribute)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "DuplicateAttributeValues":
		obj := DuplicateAttributeValuesError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		for i := range obj.Attributes {
			var err error
			obj.Attributes[i], err = mapDiscriminatorAttribute(obj.Attributes[i])
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "DuplicateField":
		obj := DuplicateFieldError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "DuplicateVariantValues":
		obj := DuplicateVariantValuesError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "insufficient_scope":
		obj := InsufficientScopeError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "InvalidCredentials":
		obj := InvalidCredentialsError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "invalid_token":
		obj := InvalidTokenError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "InvalidField":
		obj := InvalidFieldError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "InvalidJsonInput":
		obj := InvalidJsonInput{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "InvalidInput":
		obj := InvalidInput{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ResourceNotFound":
		obj := ResourceNotFoundError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ResourceCreation":
		obj := ResourceCreationError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ResourceUpdate":
		obj := ResourceUpdateError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ResourceDeletion":
		obj := ResourceDeletionError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "RequiredField":
		obj := RequiredFieldError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "InvalidTransition":
		obj := InvalidStateTransitionError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ConcurrentModification":
		obj := ConcurrentModificationError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "Contention":
		obj := ContentionError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "Generic":
		obj := GenericError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	This is the generic error code for access denied. In case of a wrong scope, an [InvalidScopeError](#invalidscopeerror) will be returned.
 */
type AccessDeniedError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AccessDeniedError) MarshalJSON() ([]byte, error) {
	type Alias AccessDeniedError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "access_denied", Alias: (*Alias)(&obj)})
}

func (obj AccessDeniedError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown AccessDeniedError: failed to parse error response"
}

/**
*	The requested scope is invalid, unknown, malformed, or exceeds the scope granted by the resource owner.
*
 */
type InvalidScopeError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InvalidScopeError) MarshalJSON() ([]byte, error) {
	type Alias InvalidScopeError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "invalid_scope", Alias: (*Alias)(&obj)})
}

func (obj InvalidScopeError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown InvalidScopeError: failed to parse error response"
}

/**
*	The resources in the request are not in the valid state for the operation.
*	The client application should validate the constraints described in the error message before sending the request again.
*
 */
type InvalidOperation struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InvalidOperation) MarshalJSON() ([]byte, error) {
	type Alias InvalidOperation
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "InvalidOperation", Alias: (*Alias)(&obj)})
}

/**
*	The `Unique` [Attribute Constraint](/../api/projects/productTypes#attributeconstraint-enum) was violated.
 */
type DuplicateAttributeValueError struct {
	Message string `json:"message"`
	// The attribute in conflict.
	Attribute Attribute `json:"attribute"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DuplicateAttributeValueError) UnmarshalJSON(data []byte) error {
	type Alias DuplicateAttributeValueError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Attribute != nil {
		var err error
		obj.Attribute, err = mapDiscriminatorAttribute(obj.Attribute)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DuplicateAttributeValueError) MarshalJSON() ([]byte, error) {
	type Alias DuplicateAttributeValueError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "DuplicateAttributeValue", Alias: (*Alias)(&obj)})
}

func (obj DuplicateAttributeValueError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown DuplicateAttributeValueError: failed to parse error response"
}

/**
*	The `CombinationUnique` [Attribute Constraint](/../api/projects/productTypes#attributeconstraint-enum) was violated.
 */
type DuplicateAttributeValuesError struct {
	Message    string      `json:"message"`
	Attributes []Attribute `json:"attributes"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DuplicateAttributeValuesError) UnmarshalJSON(data []byte) error {
	type Alias DuplicateAttributeValuesError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Attributes {
		var err error
		obj.Attributes[i], err = mapDiscriminatorAttribute(obj.Attributes[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DuplicateAttributeValuesError) MarshalJSON() ([]byte, error) {
	type Alias DuplicateAttributeValuesError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "DuplicateAttributeValues", Alias: (*Alias)(&obj)})
}

func (obj DuplicateAttributeValuesError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown DuplicateAttributeValuesError: failed to parse error response"
}

/**
*	The given value already exists for a field that is checked for unique values.
 */
type DuplicateFieldError struct {
	Message string `json:"message"`
	// The name of the field.
	Field *string `json:"field,omitempty"`
	// The offending duplicate value.
	DuplicateValue interface{} `json:"duplicateValue,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DuplicateFieldError) MarshalJSON() ([]byte, error) {
	type Alias DuplicateFieldError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "DuplicateField", Alias: (*Alias)(&obj)})
}

func (obj DuplicateFieldError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown DuplicateFieldError: failed to parse error response"
}

/**
*	The given combination of values of a [Product Variant](/../api/projects/products#productvariant) conflicts with an existing one.
*	Every [Product Variant](/../api/projects/products#productvariant) must have a distinct combination of SKU, prices, and custom attribute values.
*
 */
type DuplicateVariantValuesError struct {
	Message string `json:"message"`
	// The offending variant values.
	VariantValues VariantValues `json:"variantValues"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DuplicateVariantValuesError) MarshalJSON() ([]byte, error) {
	type Alias DuplicateVariantValuesError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "DuplicateVariantValues", Alias: (*Alias)(&obj)})
}

func (obj DuplicateVariantValuesError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown DuplicateVariantValuesError: failed to parse error response"
}

type VariantValues struct {
	Sku        *string       `json:"sku,omitempty"`
	Prices     []PriceImport `json:"prices"`
	Attributes []Attribute   `json:"attributes"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *VariantValues) UnmarshalJSON(data []byte) error {
	type Alias VariantValues
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Attributes {
		var err error
		obj.Attributes[i], err = mapDiscriminatorAttribute(obj.Attributes[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type InsufficientScopeError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InsufficientScopeError) MarshalJSON() ([]byte, error) {
	type Alias InsufficientScopeError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "insufficient_scope", Alias: (*Alias)(&obj)})
}

func (obj InsufficientScopeError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown InsufficientScopeError: failed to parse error response"
}

type InvalidCredentialsError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InvalidCredentialsError) MarshalJSON() ([]byte, error) {
	type Alias InvalidCredentialsError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "InvalidCredentials", Alias: (*Alias)(&obj)})
}

func (obj InvalidCredentialsError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown InvalidCredentialsError: failed to parse error response"
}

type InvalidTokenError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InvalidTokenError) MarshalJSON() ([]byte, error) {
	type Alias InvalidTokenError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "invalid_token", Alias: (*Alias)(&obj)})
}

func (obj InvalidTokenError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown InvalidTokenError: failed to parse error response"
}

/**
*	A given field is not supported.
*	This error occurs, for example, if the field `variants`, which is not supported by [Product Import](/product#productimport), is sent to the Product Import endpoint.
*
 */
type InvalidFieldError struct {
	Message string `json:"message"`
	// The name of the field.
	Field string `json:"field"`
	// The invalid value.
	InvalidValue interface{} `json:"invalidValue"`
	// The set of allowed values for the field, if any.
	AllowedValues []interface{} `json:"allowedValues"`
	ResourceIndex *int          `json:"resourceIndex,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InvalidFieldError) MarshalJSON() ([]byte, error) {
	type Alias InvalidFieldError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "InvalidField", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["allowedValues"] == nil {
		delete(raw, "allowedValues")
	}

	return json.Marshal(raw)

}

func (obj InvalidFieldError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown InvalidFieldError: failed to parse error response"
}

/**
*	An invalid JSON input has been sent to the service.
*	Either the JSON is syntactically incorrect or the JSON has an unexpected shape, for example, a required field is missing.
*	The client application should validate the input according to the constraints described in the error message before sending the request again.
*
 */
type InvalidJsonInput struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InvalidJsonInput) MarshalJSON() ([]byte, error) {
	type Alias InvalidJsonInput
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "InvalidJsonInput", Alias: (*Alias)(&obj)})
}

/**
*	An invalid input has been sent to the service. The client application should validate the input according to the
*	constraints described in the error message before sending the request again.
*
 */
type InvalidInput struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InvalidInput) MarshalJSON() ([]byte, error) {
	type Alias InvalidInput
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "InvalidInput", Alias: (*Alias)(&obj)})
}

type ResourceNotFoundError struct {
	Message  string      `json:"message"`
	Resource interface{} `json:"resource,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ResourceNotFoundError) MarshalJSON() ([]byte, error) {
	type Alias ResourceNotFoundError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ResourceNotFound", Alias: (*Alias)(&obj)})
}

func (obj ResourceNotFoundError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ResourceNotFoundError: failed to parse error response"
}

type ResourceCreationError struct {
	Message  string      `json:"message"`
	Resource interface{} `json:"resource,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ResourceCreationError) MarshalJSON() ([]byte, error) {
	type Alias ResourceCreationError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ResourceCreation", Alias: (*Alias)(&obj)})
}

func (obj ResourceCreationError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ResourceCreationError: failed to parse error response"
}

type ResourceUpdateError struct {
	Message  string      `json:"message"`
	Resource interface{} `json:"resource,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ResourceUpdateError) MarshalJSON() ([]byte, error) {
	type Alias ResourceUpdateError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ResourceUpdate", Alias: (*Alias)(&obj)})
}

func (obj ResourceUpdateError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ResourceUpdateError: failed to parse error response"
}

type ResourceDeletionError struct {
	Message  string      `json:"message"`
	Resource interface{} `json:"resource,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ResourceDeletionError) MarshalJSON() ([]byte, error) {
	type Alias ResourceDeletionError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ResourceDeletion", Alias: (*Alias)(&obj)})
}

func (obj ResourceDeletionError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ResourceDeletionError: failed to parse error response"
}

/**
*	A required field is missing a value.
 */
type RequiredFieldError struct {
	Message string `json:"message"`
	// The name of the field.
	Field string `json:"field"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RequiredFieldError) MarshalJSON() ([]byte, error) {
	type Alias RequiredFieldError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "RequiredField", Alias: (*Alias)(&obj)})
}

func (obj RequiredFieldError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown RequiredFieldError: failed to parse error response"
}

type InvalidStateTransitionError struct {
	Message string `json:"message"`
	// Every [Import Operation](/import-operation) is assigned with one of the following states.
	CurrentState ProcessingState `json:"currentState"`
	// Every [Import Operation](/import-operation) is assigned with one of the following states.
	NewState ProcessingState `json:"newState"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InvalidStateTransitionError) MarshalJSON() ([]byte, error) {
	type Alias InvalidStateTransitionError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "InvalidTransition", Alias: (*Alias)(&obj)})
}

func (obj InvalidStateTransitionError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown InvalidStateTransitionError: failed to parse error response"
}

/**
*	The request conflicts with the current state of the involved resources.
*	This error typically occurs when the request attempts to modify a resource that is out of date, that is, it has been modified by another client since the last time it was retrieved by the system attempting to update it.
*	The client application should resolve the conflict (with or without involving the end-user) before retrying the request.
*
 */
type ConcurrentModificationError struct {
	Message string `json:"message"`
	// The version specified in the failed request.
	SpecifiedVersion *int `json:"specifiedVersion,omitempty"`
	// The current version of the resource.
	CurrentVersion int `json:"currentVersion"`
	// The resource in conflict.
	ConflictedResource interface{} `json:"conflictedResource,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConcurrentModificationError) MarshalJSON() ([]byte, error) {
	type Alias ConcurrentModificationError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ConcurrentModification", Alias: (*Alias)(&obj)})
}

func (obj ConcurrentModificationError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ConcurrentModificationError: failed to parse error response"
}

type ContentionError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ContentionError) MarshalJSON() ([]byte, error) {
	type Alias ContentionError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "Contention", Alias: (*Alias)(&obj)})
}

func (obj ContentionError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ContentionError: failed to parse error response"
}

type GenericError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj GenericError) MarshalJSON() ([]byte, error) {
	type Alias GenericError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "Generic", Alias: (*Alias)(&obj)})
}

func (obj GenericError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown GenericError: failed to parse error response"
}
