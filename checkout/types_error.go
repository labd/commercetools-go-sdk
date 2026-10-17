package checkout

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
)

/**
*	This is the representation of a single error.
 */
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
	case "ConcurrentModification":
		obj := ConcurrentModificationError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ConnectorFailed":
		obj := ConnectorFailedError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ConnectorTimeout":
		obj := ConnectorTimeoutError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "DuplicateFieldWithConflictingResource":
		obj := DuplicateFieldWithConflictingResourceError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "General":
		obj := GeneralError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "InternalConstraintViolated":
		obj := InternalConstraintViolatedError{}
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
	case "InvalidInput":
		obj := InvalidInputError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "InvalidJsonInput":
		obj := InvalidJsonInputError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "InvalidOperation":
		obj := InvalidOperationError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "MaxResourceLimitExceeded":
		obj := MaxResourceLimitExceededError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "MissingProjectKey":
		obj := MissingProjectKeyError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "MultipleActionsNotAllowed":
		obj := MultipleActionsNotAllowedError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "PaymentFailure":
		obj := PaymentFailureError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ReferencedResourceNotFound":
		obj := ReferencedResourceNotFoundError{}
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
	case "ResourceNotFound":
		obj := ResourceNotFoundError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ServiceUnavailable":
		obj := ServiceUnavailableError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SyntaxError":
		obj := SyntaxErrorError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Returned when the request conflicts with the current state of the involved resources. Typically, the request attempts to modify a resource that is out of date (that is modified by another client since it was last retrieved).
*	The client application should resolve the conflict (with or without involving the end user) before retrying the request.
*
 */
type ConcurrentModificationError struct {
	// `"Object $resourceId has a different version than expected. Expected: $expectedVersion - Actual: $currentVersion."`
	Message string `json:"message"`
	// The ID of the resource that has a version conflict.
	ResourceId string `json:"resourceId"`
	// Expected version of the resource.
	ExpectedVersion int `json:"expectedVersion"`
	// Current version of the resource.
	CurrentVersion int `json:"currentVersion"`
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

/**
*	Returned when the payment [Connector](/connectors-and-applications#payment-connectors) cannot be reached.
*
 */
type ConnectorFailedError struct {
	// `"The connector could not be reached."`
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConnectorFailedError) MarshalJSON() ([]byte, error) {
	type Alias ConnectorFailedError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ConnectorFailed", Alias: (*Alias)(&obj)})
}

func (obj ConnectorFailedError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ConnectorFailedError: failed to parse error response"
}

/**
*	Returned when the payment [Connector](/checkout/connectors-and-applications#payment-connectors) does not respond within the configured timeout of 30 seconds. The Connector may still be processing the request.
*
 */
type ConnectorTimeoutError struct {
	// `"The connector did not respond within the configured timeout."`
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConnectorTimeoutError) MarshalJSON() ([]byte, error) {
	type Alias ConnectorTimeoutError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ConnectorTimeout", Alias: (*Alias)(&obj)})
}

func (obj ConnectorTimeoutError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ConnectorTimeoutError: failed to parse error response"
}

/**
*	Returned when a field value conflicts with an existing value stored in a particular resource causing a duplicate.
*
 */
type DuplicateFieldWithConflictingResourceError struct {
	// `"A duplicate value $duplicateValue exists for field $field in request."`
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DuplicateFieldWithConflictingResourceError) MarshalJSON() ([]byte, error) {
	type Alias DuplicateFieldWithConflictingResourceError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "DuplicateFieldWithConflictingResource", Alias: (*Alias)(&obj)})
}

func (obj DuplicateFieldWithConflictingResourceError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown DuplicateFieldWithConflictingResourceError: failed to parse error response"
}

/**
*	Base representation of an error response containing common fields to all errors.
*
 */
type ErrorResponse struct {
	// HTTP status code corresponding to the error.
	StatusCode int `json:"statusCode"`
	// First error message in the `errors` array.
	Message string `json:"message"`
	// Errors returned for a request.
	//
	// A single error response can contain multiple errors if the errors are related to the same HTTP status code such as `400`.
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

/**
*	Returned when a server-side problem occurs. In some cases, the requested action may successfully complete after the error is returned. Therefore, it is recommended to verify the status of the requested resource after receiving a 500 error.
*
*	If you encounter this error, report it to the [commercetools support team](https://support.commercetools.com/).
*
 */
type GeneralError struct {
	// Description about any known details of the problem, for example, `"Write operations are temporarily unavailable"`.
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj GeneralError) MarshalJSON() ([]byte, error) {
	type Alias GeneralError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "General", Alias: (*Alias)(&obj)})
}

func (obj GeneralError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown GeneralError: failed to parse error response"
}

/**
*	Returned when the referenced resources violate a constraint required for the operation, for example, when a Cart and a PaymentMethod are expected to belong to the same customer but do not.
*
 */
type InternalConstraintViolatedError struct {
	// Description of the error.
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InternalConstraintViolatedError) MarshalJSON() ([]byte, error) {
	type Alias InternalConstraintViolatedError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "InternalConstraintViolated", Alias: (*Alias)(&obj)})
}

func (obj InternalConstraintViolatedError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown InternalConstraintViolatedError: failed to parse error response"
}

/**
*	Returned when a field has an invalid value.
*
 */
type InvalidFieldError struct {
	// `"The value $invalidValue is not valid for field $field."`
	Message string `json:"message"`
	// Name of the field with the invalid value.
	Field string `json:"field"`
	// Value invalid for the field.
	InvalidValue interface{} `json:"invalidValue"`
	// Fixed set of allowed values for the field, if any.
	AllowedValues []string `json:"allowedValues"`
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
*	Returned when an invalid input has been sent.
*
 */
type InvalidInputError struct {
	// Description of the invalid input error.
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InvalidInputError) MarshalJSON() ([]byte, error) {
	type Alias InvalidInputError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "InvalidInput", Alias: (*Alias)(&obj)})
}

func (obj InvalidInputError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown InvalidInputError: failed to parse error response"
}

/**
*	Returned when an invalid JSON input has been sent.
*	Either the JSON is syntactically incorrect or does not conform to the expected shape, for example, it is missing a required field.
*
*	The client application should validate the input according to the constraints described in the error message before sending the request.
*
 */
type InvalidJsonInputError struct {
	// `"Request body does not contain valid JSON."`
	Message string `json:"message"`
	// Further explanation about why the JSON is invalid.
	DetailedErrorMessage *string `json:"detailedErrorMessage,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InvalidJsonInputError) MarshalJSON() ([]byte, error) {
	type Alias InvalidJsonInputError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "InvalidJsonInput", Alias: (*Alias)(&obj)})
}

func (obj InvalidJsonInputError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown InvalidJsonInputError: failed to parse error response"
}

/**
*	Returned when the resources involved in the request are not in a valid state for the operation.
*
*	The client application should validate the constraints described in the error message before sending the request.
*
 */
type InvalidOperationError struct {
	// Description of the error.
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InvalidOperationError) MarshalJSON() ([]byte, error) {
	type Alias InvalidOperationError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "InvalidOperation", Alias: (*Alias)(&obj)})
}

func (obj InvalidOperationError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown InvalidOperationError: failed to parse error response"
}

/**
*	Returned when a resource type cannot be created as it has reached its [limits](/limits).
*
*	The limits must be adjusted for this resource before sending the request again.
*
 */
type MaxResourceLimitExceededError struct {
	// `"You have exceeded the limit of $limit resources of type $resourceTypeId."`
	Message string `json:"message"`
	// The limit that was exceeded.
	Limit float64 `json:"limit"`
	// The resource type that reached its limit.
	ResourceTypeId string `json:"resourceTypeId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MaxResourceLimitExceededError) MarshalJSON() ([]byte, error) {
	type Alias MaxResourceLimitExceededError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "MaxResourceLimitExceeded", Alias: (*Alias)(&obj)})
}

func (obj MaxResourceLimitExceededError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown MaxResourceLimitExceededError: failed to parse error response"
}

/**
*	Returned when the project key is missing from the request.
*
 */
type MissingProjectKeyError struct {
	// `"Missing project key in the request path."`
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MissingProjectKeyError) MarshalJSON() ([]byte, error) {
	type Alias MissingProjectKeyError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "MissingProjectKey", Alias: (*Alias)(&obj)})
}

func (obj MissingProjectKeyError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown MissingProjectKeyError: failed to parse error response"
}

/**
*	Returned when `actions` in the request body contains more than one object.
*
 */
type MultipleActionsNotAllowedError struct {
	// `"Actions accepts only one action at time. Array size must be 1."`
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MultipleActionsNotAllowedError) MarshalJSON() ([]byte, error) {
	type Alias MultipleActionsNotAllowedError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "MultipleActionsNotAllowed", Alias: (*Alias)(&obj)})
}

func (obj MultipleActionsNotAllowedError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown MultipleActionsNotAllowedError: failed to parse error response"
}

/**
*	Returned when the payment cannot be completed successfully.
*
 */
type PaymentFailureError struct {
	// `"Payment could not be completed successfully."`
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentFailureError) MarshalJSON() ([]byte, error) {
	type Alias PaymentFailureError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "PaymentFailure", Alias: (*Alias)(&obj)})
}

func (obj PaymentFailureError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown PaymentFailureError: failed to parse error response"
}

/**
*	Returned when a resource referenced by a [Reference](ctp:checkout:type:Reference) or a [ResourceIdentifier](ctp:checkout:type:ResourceIdentifier) could not be found.
*
 */
type ReferencedResourceNotFoundError struct {
	// `"The referenced object of type $typeId $id || key was not found. It either doesn't exist, or it can't be accessed from this endpoint (e.g., if the endpoint filters by store or customer account)."`
	Message string `json:"message"`
	// Type of referenced resource.
	TypeId string `json:"typeId"`
	// Unique identifier of the referenced resource, if known.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced resource, if known.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ReferencedResourceNotFoundError) MarshalJSON() ([]byte, error) {
	type Alias ReferencedResourceNotFoundError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ReferencedResourceNotFound", Alias: (*Alias)(&obj)})
}

func (obj ReferencedResourceNotFoundError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ReferencedResourceNotFoundError: failed to parse error response"
}

/**
*	Returned when a value is not defined for a required field.
*
 */
type RequiredFieldError struct {
	// `"A value is required for field $field."`
	Message string `json:"message"`
	// Name of the field missing the value.
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

/**
*	Returned when the resource addressed by the request URL does not exist.
*
 */
type ResourceNotFoundError struct {
	// `"The Resource with $resourceIdentifier $resourceId was not found."`
	Message string `json:"message"`
	// The identifier type used (e.g., `id`, `key`).
	ResourceIdentifier string `json:"resourceIdentifier"`
	// The actual identifier value.
	ResourceId string `json:"resourceId"`
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

/**
*	Returned when the service is unavailable, for example when the Notification Service is down.
*
 */
type ServiceUnavailableError struct {
	// `"Service is unavailable."`
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ServiceUnavailableError) MarshalJSON() ([]byte, error) {
	type Alias ServiceUnavailableError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ServiceUnavailable", Alias: (*Alias)(&obj)})
}

func (obj ServiceUnavailableError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ServiceUnavailableError: failed to parse error response"
}

/**
*	Returned when a [Payment Predicate](/payment-integration-predicates) does not have the correct syntax.
*
 */
type SyntaxErrorError struct {
	// `"Syntax error while parsing $fieldDefinition."`
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SyntaxErrorError) MarshalJSON() ([]byte, error) {
	type Alias SyntaxErrorError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "SyntaxError", Alias: (*Alias)(&obj)})
}

func (obj SyntaxErrorError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown SyntaxErrorError: failed to parse error response"
}
