package connect

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
)

/**
*	Represents a single error. Multiple errors may be included in an [ErrorResponse](ctp:connect:type:ErrorResponse).
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
	case "AuthenticationError":
		obj := AuthenticationError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AuthorizationError":
		obj := AuthorizationError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "BadGateway":
		obj := BadGatewayError{}
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
	case "DuplicateField":
		obj := DuplicateFieldError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "FieldValueNotFound":
		obj := FieldValueNotFoundError{}
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
	case "InvalidJsonInput":
		obj := InvalidJsonInputError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "InvalidPathParam":
		obj := InvalidPathParamError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "InvalidQueryParam":
		obj := InvalidQueryParamError{}
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
	case "ConnectorAlreadyCertified":
		obj := ConnectorAlreadyCertifiedError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ConnectorStagedInCertification":
		obj := ConnectorStagedInCertificationError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ConnectorStagedNotPreviewable":
		obj := ConnectorStagedNotPreviewableError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ConnectorStagedNotPrivate":
		obj := ConnectorStagedNotPrivateError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ConnectorStagedPreviewRequestUnderProcess":
		obj := ConnectorStagedPreviewRequestUnderProcessError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "GitRepositoryNotReachable":
		obj := GitRepositoryNotReachableError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ConnectorSpecificationFileNotFound":
		obj := ConnectorSpecificationFileNotFoundError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ConnectorSpecificationFileNotValid":
		obj := ConnectorSpecificationFileNotValidError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ConnectorReferenceNotFound":
		obj := ConnectorReferenceNotFoundError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "DeploymentInvalidStatusTransition":
		obj := DeploymentInvalidStatusTransitionError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "DeploymentUnknownApplicationConfiguration":
		obj := DeploymentUnknownApplicationConfigurationError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "DeploymentUnknownApplicationConfigurationKey":
		obj := DeploymentUnknownApplicationConfigurationKeyError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "DeploymentUnsupportedRegion":
		obj := DeploymentUnsupportedRegionError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "DeploymentApplicationDoNotBelong":
		obj := DeploymentEntityApplicationDoNotBelongError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "DeploymentApplicationRequired":
		obj := DeploymentEntityApplicationRequiredError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "DeploymentMustIncludeApplication":
		obj := DeploymentEntityMustIncludeApplicationError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
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
*	Returned when the client cannot be authenticated.
*
 */
type AuthenticationError struct {
	// `"Bad credientals or Client ID is not defined"`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AuthenticationError) UnmarshalJSON(data []byte) error {
	type Alias AuthenticationError
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
func (obj AuthenticationError) MarshalJSON() ([]byte, error) {
	type Alias AuthenticationError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "AuthenticationError", Alias: (*Alias)(&obj)})
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

func (obj *AuthenticationError) DecodeStruct(src map[string]interface{}) error {
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

func (obj AuthenticationError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown AuthenticationError: failed to parse error response"
}

/**
*	Returned when the client does not have sufficient permissions for this operation.
*
 */
type AuthorizationError struct {
	// `"Access denied"`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AuthorizationError) UnmarshalJSON(data []byte) error {
	type Alias AuthorizationError
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
func (obj AuthorizationError) MarshalJSON() ([]byte, error) {
	type Alias AuthorizationError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "AuthorizationError", Alias: (*Alias)(&obj)})
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

func (obj *AuthorizationError) DecodeStruct(src map[string]interface{}) error {
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

func (obj AuthorizationError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown AuthorizationError: failed to parse error response"
}

/**
*	Returned when a server-side problem is caused by scaling infrastructure resources.
*
*	The client application should retry the request with exponential backoff up to a point where further delay is unacceptable.
*
 */
type BadGatewayError struct {
	// Plain text description of the error.
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *BadGatewayError) UnmarshalJSON(data []byte) error {
	type Alias BadGatewayError
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
func (obj BadGatewayError) MarshalJSON() ([]byte, error) {
	type Alias BadGatewayError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "BadGateway", Alias: (*Alias)(&obj)})
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

func (obj *BadGatewayError) DecodeStruct(src map[string]interface{}) error {
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

func (obj BadGatewayError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown BadGatewayError: failed to parse error response"
}

/**
*	Returned when the request conflicts with the current state of the involved resources. Typically, the request attempts to modify a resource that is out of date (that is modified by another client since it was last retrieved).
*	The client application should resolve the conflict (with or without involving the end-user) before retrying the request.
*
 */
type ConcurrentModificationError struct {
	// `"Object id=$resourceId or key=$resourceKey has a different version than expected. Expected: $expectedVersion - Actual: $currentVersion)"`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// Expected version of the resource.
	ExpectedVersion int `json:"expectedVersion"`
	// Current version of the resource.
	CurrentVersion *int `json:"currentVersion,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ConcurrentModificationError) UnmarshalJSON(data []byte) error {
	type Alias ConcurrentModificationError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "expectedVersion")
	delete(obj.ExtraValues, "currentVersion")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConcurrentModificationError) MarshalJSON() ([]byte, error) {
	type Alias ConcurrentModificationError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ConcurrentModification", Alias: (*Alias)(&obj)})
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

func (obj *ConcurrentModificationError) DecodeStruct(src map[string]interface{}) error {
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

func (obj ConcurrentModificationError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ConcurrentModificationError: failed to parse error response"
}

/**
*	Returned when a field value conflicts with an existing value causing a duplicate.
*
 */
type DuplicateFieldError struct {
	// `"A duplicate value $duplicateValue exists for field $field."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// Name of the conflicting field.
	Field string `json:"field"`
	// Conflicting duplicate value.
	DuplicateValue interface{} `json:"duplicateValue"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DuplicateFieldError) UnmarshalJSON(data []byte) error {
	type Alias DuplicateFieldError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "field")
	delete(obj.ExtraValues, "duplicateValue")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DuplicateFieldError) MarshalJSON() ([]byte, error) {
	type Alias DuplicateFieldError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "DuplicateField", Alias: (*Alias)(&obj)})
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

func (obj *DuplicateFieldError) DecodeStruct(src map[string]interface{}) error {
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

func (obj DuplicateFieldError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown DuplicateFieldError: failed to parse error response"
}

/**
*	Returned when a value for a field is not found.
*
 */
type FieldValueNotFoundError struct {
	// `"The value $value for the $field field was not found"`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// Name of the field.
	Field string `json:"field"`
	// Conflicting value.
	Value interface{} `json:"value"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *FieldValueNotFoundError) UnmarshalJSON(data []byte) error {
	type Alias FieldValueNotFoundError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "field")
	delete(obj.ExtraValues, "value")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj FieldValueNotFoundError) MarshalJSON() ([]byte, error) {
	type Alias FieldValueNotFoundError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "FieldValueNotFound", Alias: (*Alias)(&obj)})
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

func (obj *FieldValueNotFoundError) DecodeStruct(src map[string]interface{}) error {
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

func (obj FieldValueNotFoundError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown FieldValueNotFoundError: failed to parse error response"
}

/**
*	Returned when a server-side problem occurs.
*
*	If you encounter this error, report it using the [Support Portal](TBC for connect).
*
 */
type GeneralError struct {
	// Description about any known details of the problem, for example, `"Write operations are temporarily unavailable"`.
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *GeneralError) UnmarshalJSON(data []byte) error {
	type Alias GeneralError
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
func (obj GeneralError) MarshalJSON() ([]byte, error) {
	type Alias GeneralError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "General", Alias: (*Alias)(&obj)})
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

func (obj *GeneralError) DecodeStruct(src map[string]interface{}) error {
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

func (obj GeneralError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown GeneralError: failed to parse error response"
}

/**
*	Returned when an invalid JSON input has been sent.
*	Either the JSON is syntactically incorrect or does not conform to the expected shape (for example is missing a required field).
*
*	The client application should validate the input according to the constraints described in the error message before sending the request.
*
 */
type InvalidJsonInputError struct {
	// `"Request body does not contain valid JSON."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// Further explanation about why the JSON is invalid.
	DetailedErrorMessage string `json:"detailedErrorMessage"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *InvalidJsonInputError) UnmarshalJSON(data []byte) error {
	type Alias InvalidJsonInputError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "detailedErrorMessage")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InvalidJsonInputError) MarshalJSON() ([]byte, error) {
	type Alias InvalidJsonInputError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "InvalidJsonInput", Alias: (*Alias)(&obj)})
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

func (obj *InvalidJsonInputError) DecodeStruct(src map[string]interface{}) error {
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

func (obj InvalidJsonInputError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown InvalidJsonInputError: failed to parse error response"
}

/**
*	Returned when an invalid path parameter has been sent.
*
 */
type InvalidPathParamError struct {
	// `"Request path param is not valid"`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// Further explanation about why the path parameter is invalid.
	DetailedErrorMessage string `json:"detailedErrorMessage"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *InvalidPathParamError) UnmarshalJSON(data []byte) error {
	type Alias InvalidPathParamError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "detailedErrorMessage")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InvalidPathParamError) MarshalJSON() ([]byte, error) {
	type Alias InvalidPathParamError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "InvalidPathParam", Alias: (*Alias)(&obj)})
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

func (obj *InvalidPathParamError) DecodeStruct(src map[string]interface{}) error {
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

func (obj InvalidPathParamError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown InvalidPathParamError: failed to parse error response"
}

/**
*	Returned when an invalid query parameter has been sent.
*
 */
type InvalidQueryParamError struct {
	// `"Request query param is not valid"`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// Further explanation about why the query parameter is invalid.
	DetailedErrorMessage string `json:"detailedErrorMessage"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *InvalidQueryParamError) UnmarshalJSON(data []byte) error {
	type Alias InvalidQueryParamError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "detailedErrorMessage")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InvalidQueryParamError) MarshalJSON() ([]byte, error) {
	type Alias InvalidQueryParamError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "InvalidQueryParam", Alias: (*Alias)(&obj)})
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

func (obj *InvalidQueryParamError) DecodeStruct(src map[string]interface{}) error {
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

func (obj InvalidQueryParamError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown InvalidQueryParamError: failed to parse error response"
}

/**
*	Returned when the resource addressed by the request URL does not exist.
*
 */
type ResourceNotFoundError struct {
	// `"Deployment with id=$resourceId or key=$resourceKey does not exist"`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ResourceNotFoundError) UnmarshalJSON(data []byte) error {
	type Alias ResourceNotFoundError
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
func (obj ResourceNotFoundError) MarshalJSON() ([]byte, error) {
	type Alias ResourceNotFoundError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ResourceNotFound", Alias: (*Alias)(&obj)})
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

func (obj *ResourceNotFoundError) DecodeStruct(src map[string]interface{}) error {
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

func (obj ResourceNotFoundError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ResourceNotFoundError: failed to parse error response"
}

/**
*	Returned when trying to certify a Connector that is already certified.
*
*	The error is returned as a failed response to the [Publish](ctp:connect:type:ConnectorPublishAction) update action only when certification is required.
*
 */
type ConnectorAlreadyCertifiedError struct {
	// `"The ConnectorStaged is already certified"`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ConnectorAlreadyCertifiedError) UnmarshalJSON(data []byte) error {
	type Alias ConnectorAlreadyCertifiedError
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
func (obj ConnectorAlreadyCertifiedError) MarshalJSON() ([]byte, error) {
	type Alias ConnectorAlreadyCertifiedError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ConnectorAlreadyCertified", Alias: (*Alias)(&obj)})
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

func (obj *ConnectorAlreadyCertifiedError) DecodeStruct(src map[string]interface{}) error {
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

func (obj ConnectorAlreadyCertifiedError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ConnectorAlreadyCertifiedError: failed to parse error response"
}

/**
*	Returned when trying to publish a Connector that requires certification but is already in the certification process.
*
*	The error is returned as a failed response to the [Publish](ctp:connect:type:ConnectorPublishAction) update actiononly when certification is required.
*
 */
type ConnectorStagedInCertificationError struct {
	// `"The ConnectorStaged is already in certification"`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ConnectorStagedInCertificationError) UnmarshalJSON(data []byte) error {
	type Alias ConnectorStagedInCertificationError
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
func (obj ConnectorStagedInCertificationError) MarshalJSON() ([]byte, error) {
	type Alias ConnectorStagedInCertificationError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ConnectorStagedInCertification", Alias: (*Alias)(&obj)})
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

func (obj *ConnectorStagedInCertificationError) DecodeStruct(src map[string]interface{}) error {
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

func (obj ConnectorStagedInCertificationError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ConnectorStagedInCertificationError: failed to parse error response"
}

/**
*	Returned when a ConnectorStaged to be deployed is not previewable.
*
*	The error is returned as a failed response to the [Create a Deployment](deployments#create-deployment) request.
*
 */
type ConnectorStagedNotPreviewableError struct {
	// `"Connector id=$resourceId or key=$resourceKey with version $version is not previewable"`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ConnectorStagedNotPreviewableError) UnmarshalJSON(data []byte) error {
	type Alias ConnectorStagedNotPreviewableError
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
func (obj ConnectorStagedNotPreviewableError) MarshalJSON() ([]byte, error) {
	type Alias ConnectorStagedNotPreviewableError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ConnectorStagedNotPreviewable", Alias: (*Alias)(&obj)})
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

func (obj *ConnectorStagedNotPreviewableError) DecodeStruct(src map[string]interface{}) error {
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

func (obj ConnectorStagedNotPreviewableError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ConnectorStagedNotPreviewableError: failed to parse error response"
}

/**
*	Returned when attempting to change the `privateProjects` of a non-private ConnectorStaged.
*
*	The error is returned as a failed response to the [Add Project to Private Connector](ctp:connect:type:ConnectorAddPrivateProjectAction) and [Remove Project from Private Connector](ctp:connect:type:ConnectorRemovePrivateProjectAction) update actions.
*
 */
type ConnectorStagedNotPrivateError struct {
	// `"The operation is not valid because ConnectorStaged with id=$resourceId or key=$resourceKey is not private"`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ConnectorStagedNotPrivateError) UnmarshalJSON(data []byte) error {
	type Alias ConnectorStagedNotPrivateError
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
func (obj ConnectorStagedNotPrivateError) MarshalJSON() ([]byte, error) {
	type Alias ConnectorStagedNotPrivateError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ConnectorStagedNotPrivate", Alias: (*Alias)(&obj)})
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

func (obj *ConnectorStagedNotPrivateError) DecodeStruct(src map[string]interface{}) error {
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

func (obj ConnectorStagedNotPrivateError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ConnectorStagedNotPrivateError: failed to parse error response"
}

/**
*	Returned when attempting to request previewable status of a ConnectorStaged that is currently being reviewed for previewable status.
*
*	The error is returned as a failed response to the [Preview Connector](ctp:connect:type:ConnectorUpdatePreviewableAction) update action.
*
 */
type ConnectorStagedPreviewRequestUnderProcessError struct {
	// `"The ConnectorStaged preview request is already in progress"`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ConnectorStagedPreviewRequestUnderProcessError) UnmarshalJSON(data []byte) error {
	type Alias ConnectorStagedPreviewRequestUnderProcessError
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
func (obj ConnectorStagedPreviewRequestUnderProcessError) MarshalJSON() ([]byte, error) {
	type Alias ConnectorStagedPreviewRequestUnderProcessError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ConnectorStagedPreviewRequestUnderProcess", Alias: (*Alias)(&obj)})
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

func (obj *ConnectorStagedPreviewRequestUnderProcessError) DecodeStruct(src map[string]interface{}) error {
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

func (obj ConnectorStagedPreviewRequestUnderProcessError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ConnectorStagedPreviewRequestUnderProcessError: failed to parse error response"
}

/**
*	Returned when the GitHub repository is unreachable or not found.
*
 */
type GitRepositoryNotReachableError struct {
	// `"Repository with the tag is not reachable"`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *GitRepositoryNotReachableError) UnmarshalJSON(data []byte) error {
	type Alias GitRepositoryNotReachableError
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
func (obj GitRepositoryNotReachableError) MarshalJSON() ([]byte, error) {
	type Alias GitRepositoryNotReachableError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "GitRepositoryNotReachable", Alias: (*Alias)(&obj)})
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

func (obj *GitRepositoryNotReachableError) DecodeStruct(src map[string]interface{}) error {
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

func (obj GitRepositoryNotReachableError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown GitRepositoryNotReachableError: failed to parse error response"
}

/**
*	Returned when the Connector specification file was not found.
*
 */
type ConnectorSpecificationFileNotFoundError struct {
	// `"The file connect.yaml at $url with the tag $tag was not found"`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ConnectorSpecificationFileNotFoundError) UnmarshalJSON(data []byte) error {
	type Alias ConnectorSpecificationFileNotFoundError
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
func (obj ConnectorSpecificationFileNotFoundError) MarshalJSON() ([]byte, error) {
	type Alias ConnectorSpecificationFileNotFoundError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ConnectorSpecificationFileNotFound", Alias: (*Alias)(&obj)})
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

func (obj *ConnectorSpecificationFileNotFoundError) DecodeStruct(src map[string]interface{}) error {
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

func (obj ConnectorSpecificationFileNotFoundError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ConnectorSpecificationFileNotFoundError: failed to parse error response"
}

/**
*	Returned when the Connector specification file is not valid.
*
 */
type ConnectorSpecificationFileNotValidError struct {
	// `"The file connect.yaml at $url with the tag $tag is not valid"`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ConnectorSpecificationFileNotValidError) UnmarshalJSON(data []byte) error {
	type Alias ConnectorSpecificationFileNotValidError
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
func (obj ConnectorSpecificationFileNotValidError) MarshalJSON() ([]byte, error) {
	type Alias ConnectorSpecificationFileNotValidError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ConnectorSpecificationFileNotValid", Alias: (*Alias)(&obj)})
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

func (obj *ConnectorSpecificationFileNotValidError) DecodeStruct(src map[string]interface{}) error {
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

func (obj ConnectorSpecificationFileNotValidError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ConnectorSpecificationFileNotValidError: failed to parse error response"
}

/**
*	Returned when the referenced Connector was not found.
*
 */
type ConnectorReferenceNotFoundError struct {
	// `"Connector '$identifier' with version $version not found"`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ConnectorReferenceNotFoundError) UnmarshalJSON(data []byte) error {
	type Alias ConnectorReferenceNotFoundError
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
func (obj ConnectorReferenceNotFoundError) MarshalJSON() ([]byte, error) {
	type Alias ConnectorReferenceNotFoundError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ConnectorReferenceNotFound", Alias: (*Alias)(&obj)})
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

func (obj *ConnectorReferenceNotFoundError) DecodeStruct(src map[string]interface{}) error {
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

func (obj ConnectorReferenceNotFoundError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ConnectorReferenceNotFoundError: failed to parse error response"
}

/**
*	Returned when the Deployment operation results in a invalid status transition.
*
*	The error is returned as a failed response to the [Redeploy](ctp:connect:type:DeploymentRedeploy) update action.
*
*	The message will contain `Already queued`, `Already deploying`, or `Already undeploying` based on the [DeploymentStatus](ctp:connect:type:DeploymentStatus) of the Deployment.
*
 */
type DeploymentInvalidStatusTransitionError struct {
	// `"Invalid status change\: Already deploying"`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DeploymentInvalidStatusTransitionError) UnmarshalJSON(data []byte) error {
	type Alias DeploymentInvalidStatusTransitionError
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
func (obj DeploymentInvalidStatusTransitionError) MarshalJSON() ([]byte, error) {
	type Alias DeploymentInvalidStatusTransitionError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "DeploymentInvalidStatusTransition", Alias: (*Alias)(&obj)})
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

func (obj *DeploymentInvalidStatusTransitionError) DecodeStruct(src map[string]interface{}) error {
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

func (obj DeploymentInvalidStatusTransitionError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown DeploymentInvalidStatusTransitionError: failed to parse error response"
}

/**
*	Returned when the Deployment contains configuration values that are not defined in the Connect application's connect.yaml file.
*
*	The error is returned as a failed response to the [Redeploy](ctp:connect:type:DeploymentRedeploy) update action and [Create a Deployment](deployments#create-deployment) request.
*
 */
type DeploymentUnknownApplicationConfigurationError struct {
	// `"Deployment does not require configuration for application named '$unknownApplicationName'."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DeploymentUnknownApplicationConfigurationError) UnmarshalJSON(data []byte) error {
	type Alias DeploymentUnknownApplicationConfigurationError
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
func (obj DeploymentUnknownApplicationConfigurationError) MarshalJSON() ([]byte, error) {
	type Alias DeploymentUnknownApplicationConfigurationError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "DeploymentUnknownApplicationConfiguration", Alias: (*Alias)(&obj)})
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

func (obj *DeploymentUnknownApplicationConfigurationError) DecodeStruct(src map[string]interface{}) error {
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

func (obj DeploymentUnknownApplicationConfigurationError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown DeploymentUnknownApplicationConfigurationError: failed to parse error response"
}

/**
*	Returned when the Deployment contains a configuration key that is not defined in the Connect application's connect.yaml file.
*
*	The error is returned as a failed response to the [Redeploy](ctp:connect:type:DeploymentRedeploy) update action and [Create a Deployment](deployments#create-deployment) request.
*
 */
type DeploymentUnknownApplicationConfigurationKeyError struct {
	// `"Deployment does not require (secret|standard) configuration with key $configurationKey for application named $applicationName."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DeploymentUnknownApplicationConfigurationKeyError) UnmarshalJSON(data []byte) error {
	type Alias DeploymentUnknownApplicationConfigurationKeyError
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
func (obj DeploymentUnknownApplicationConfigurationKeyError) MarshalJSON() ([]byte, error) {
	type Alias DeploymentUnknownApplicationConfigurationKeyError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "DeploymentUnknownApplicationConfigurationKey", Alias: (*Alias)(&obj)})
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

func (obj *DeploymentUnknownApplicationConfigurationKeyError) DecodeStruct(src map[string]interface{}) error {
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

func (obj DeploymentUnknownApplicationConfigurationKeyError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown DeploymentUnknownApplicationConfigurationKeyError: failed to parse error response"
}

/**
*	Returned when the Deployment region is not supported.
*
*	The error is returned as a failed response to the [Create a Deployment](deployments#create-deployment) request.
*
 */
type DeploymentUnsupportedRegionError struct {
	// `"Deployment unsupported region $region"`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DeploymentUnsupportedRegionError) UnmarshalJSON(data []byte) error {
	type Alias DeploymentUnsupportedRegionError
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
func (obj DeploymentUnsupportedRegionError) MarshalJSON() ([]byte, error) {
	type Alias DeploymentUnsupportedRegionError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "DeploymentUnsupportedRegion", Alias: (*Alias)(&obj)})
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

func (obj *DeploymentUnsupportedRegionError) DecodeStruct(src map[string]interface{}) error {
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

func (obj DeploymentUnsupportedRegionError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown DeploymentUnsupportedRegionError: failed to parse error response"
}

/**
*	Returned when attempting to add an application that does not belong to the Deployment.
*
 */
type DeploymentEntityApplicationDoNotBelongError struct {
	// `"Deployment with id=$resourceId or key=$resourceKey does not include application: $applicationName"`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DeploymentEntityApplicationDoNotBelongError) UnmarshalJSON(data []byte) error {
	type Alias DeploymentEntityApplicationDoNotBelongError
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
func (obj DeploymentEntityApplicationDoNotBelongError) MarshalJSON() ([]byte, error) {
	type Alias DeploymentEntityApplicationDoNotBelongError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "DeploymentApplicationDoNotBelong", Alias: (*Alias)(&obj)})
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

func (obj *DeploymentEntityApplicationDoNotBelongError) DecodeStruct(src map[string]interface{}) error {
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

func (obj DeploymentEntityApplicationDoNotBelongError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown DeploymentEntityApplicationDoNotBelongError: failed to parse error response"
}

/**
*	Returned when a Deployment does not contain any applications.
*
 */
type DeploymentEntityApplicationRequiredError struct {
	// `"A Deployment requires at least one application"`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DeploymentEntityApplicationRequiredError) UnmarshalJSON(data []byte) error {
	type Alias DeploymentEntityApplicationRequiredError
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
func (obj DeploymentEntityApplicationRequiredError) MarshalJSON() ([]byte, error) {
	type Alias DeploymentEntityApplicationRequiredError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "DeploymentApplicationRequired", Alias: (*Alias)(&obj)})
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

func (obj *DeploymentEntityApplicationRequiredError) DecodeStruct(src map[string]interface{}) error {
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

func (obj DeploymentEntityApplicationRequiredError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown DeploymentEntityApplicationRequiredError: failed to parse error response"
}

/**
*	Returned when attempting to remove an application that belongs to the Deployment.
*
 */
type DeploymentEntityMustIncludeApplicationError struct {
	// `"Deployment with id=$resourceId or key=$resourceKey must include application: $applicationName"`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DeploymentEntityMustIncludeApplicationError) UnmarshalJSON(data []byte) error {
	type Alias DeploymentEntityMustIncludeApplicationError
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
func (obj DeploymentEntityMustIncludeApplicationError) MarshalJSON() ([]byte, error) {
	type Alias DeploymentEntityMustIncludeApplicationError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "DeploymentMustIncludeApplication", Alias: (*Alias)(&obj)})
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

func (obj *DeploymentEntityMustIncludeApplicationError) DecodeStruct(src map[string]interface{}) error {
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

func (obj DeploymentEntityMustIncludeApplicationError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown DeploymentEntityMustIncludeApplicationError: failed to parse error response"
}
