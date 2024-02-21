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
	case "General":
		obj := GeneralError{}
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
	}
	return nil, nil
}

/**
*	Returned when a server-side problem occurs. In some cases, the requested action may successfully complete after the error is returned. Therefore, it is recommended to verify the status of the requested resource after receiving a 500 error.
*
*	If you encounter this error, report it using the [Support Portal](https://commercetools.atlassian.net/servicedesk/customer/portal/30).
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
	// `"The Resource with ID $resourceId was not found."`
	Message string `json:"message"`
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
