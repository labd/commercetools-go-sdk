package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type ErrorByExtension struct {
	// Unique identifier of the Extension.
	ID string `json:"id"`
	// User-defined unique identifier of the Extension.
	Key *string `json:"key,omitempty"`
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
	case "AnonymousIdAlreadyInUse":
		obj := AnonymousIdAlreadyInUseError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AttributeDefinitionAlreadyExists":
		obj := AttributeDefinitionAlreadyExistsError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AttributeDefinitionTypeConflict":
		obj := AttributeDefinitionTypeConflictError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AttributeNameDoesNotExist":
		obj := AttributeNameDoesNotExistError{}
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
	case "DiscountCodeNonApplicable":
		obj := DiscountCodeNonApplicableError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "DuplicateAttributeValue":
		obj := DuplicateAttributeValueError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "DuplicateAttributeValues":
		obj := DuplicateAttributeValuesError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "DuplicateEnumValues":
		obj := DuplicateEnumValuesError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "DuplicateField":
		obj := DuplicateFieldError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.ConflictingResource != nil {
			var err error
			obj.ConflictingResource, err = mapDiscriminatorReference(obj.ConflictingResource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "DuplicateFieldWithConflictingResource":
		obj := DuplicateFieldWithConflictingResourceError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.ConflictingResource != nil {
			var err error
			obj.ConflictingResource, err = mapDiscriminatorReference(obj.ConflictingResource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "DuplicatePriceScope":
		obj := DuplicatePriceScopeError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "DuplicateStandalonePriceScope":
		obj := DuplicateStandalonePriceScopeError{}
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
	case "EditPreviewFailed":
		obj := EditPreviewFailedError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "EnumKeyAlreadyExists":
		obj := EnumKeyAlreadyExistsError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "EnumKeyDoesNotExist":
		obj := EnumKeyDoesNotExistError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "EnumValueIsUsed":
		obj := EnumValueIsUsedError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "EnumValuesMustMatch":
		obj := EnumValuesMustMatchError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ExtensionBadResponse":
		obj := ExtensionBadResponseError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ExtensionNoResponse":
		obj := ExtensionNoResponseError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ExtensionUpdateActionsFailed":
		obj := ExtensionUpdateActionsFailedError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ExternalOAuthFailed":
		obj := ExternalOAuthFailedError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "FeatureRemoved":
		obj := FeatureRemovedError{}
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
	case "insufficient_scope":
		obj := InsufficientScopeError{}
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
	case "InvalidCredentials":
		obj := InvalidCredentialsError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "InvalidCurrentPassword":
		obj := InvalidCurrentPasswordError{}
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
	case "InvalidItemShippingDetails":
		obj := InvalidItemShippingDetailsError{}
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
	case "InvalidSubject":
		obj := InvalidSubjectError{}
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
	case "LanguageUsedInStores":
		obj := LanguageUsedInStoresError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "MatchingPriceNotFound":
		obj := MatchingPriceNotFoundError{}
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
	case "MissingRoleOnChannel":
		obj := MissingRoleOnChannelError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "MissingTaxRateForCountry":
		obj := MissingTaxRateForCountryError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "NoMatchingProductDiscountFound":
		obj := NoMatchingProductDiscountFoundError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "NotEnabled":
		obj := NotEnabledError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ObjectNotFound":
		obj := ObjectNotFoundError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "OutOfStock":
		obj := OutOfStockError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "OverCapacity":
		obj := OverCapacityError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "OverlappingStandalonePriceValidity":
		obj := OverlappingStandalonePriceValidityError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "PendingOperation":
		obj := PendingOperationError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "PriceChanged":
		obj := PriceChangedError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ProjectNotConfiguredForLanguages":
		obj := ProjectNotConfiguredForLanguagesError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "QueryComplexityLimitExceeded":
		obj := QueryComplexityLimitExceededError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "QueryTimedOut":
		obj := QueryTimedOutError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ReferenceExists":
		obj := ReferenceExistsError{}
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
	case "ResourceSizeLimitExceeded":
		obj := ResourceSizeLimitExceededError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SearchDeactivated":
		obj := SearchDeactivatedError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SearchExecutionFailure":
		obj := SearchExecutionFailureError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SearchFacetPathNotFound":
		obj := SearchFacetPathNotFoundError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SearchIndexingInProgress":
		obj := SearchIndexingInProgressError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SemanticError":
		obj := SemanticErrorError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ShippingMethodDoesNotMatchCart":
		obj := ShippingMethodDoesNotMatchCartError{}
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
	case "WeakPassword":
		obj := WeakPasswordError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type AccessDeniedError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AccessDeniedError) UnmarshalJSON(data []byte) error {
	type Alias AccessDeniedError
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
func (obj AccessDeniedError) MarshalJSON() ([]byte, error) {
	type Alias AccessDeniedError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "access_denied", Alias: (*Alias)(&obj)})
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

func (obj *AccessDeniedError) DecodeStruct(src map[string]interface{}) error {
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

func (obj AccessDeniedError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown AccessDeniedError: failed to parse error response"
}

type AnonymousIdAlreadyInUseError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AnonymousIdAlreadyInUseError) UnmarshalJSON(data []byte) error {
	type Alias AnonymousIdAlreadyInUseError
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
func (obj AnonymousIdAlreadyInUseError) MarshalJSON() ([]byte, error) {
	type Alias AnonymousIdAlreadyInUseError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "AnonymousIdAlreadyInUse", Alias: (*Alias)(&obj)})
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

func (obj *AnonymousIdAlreadyInUseError) DecodeStruct(src map[string]interface{}) error {
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

func (obj AnonymousIdAlreadyInUseError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown AnonymousIdAlreadyInUseError: failed to parse error response"
}

type AttributeDefinitionAlreadyExistsError struct {
	Message                    string                 `json:"message"`
	ExtraValues                map[string]interface{} `json:"-"`
	ConflictingProductTypeId   string                 `json:"conflictingProductTypeId"`
	ConflictingProductTypeName string                 `json:"conflictingProductTypeName"`
	ConflictingAttributeName   string                 `json:"conflictingAttributeName"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AttributeDefinitionAlreadyExistsError) UnmarshalJSON(data []byte) error {
	type Alias AttributeDefinitionAlreadyExistsError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "conflictingProductTypeId")
	delete(obj.ExtraValues, "conflictingProductTypeName")
	delete(obj.ExtraValues, "conflictingAttributeName")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AttributeDefinitionAlreadyExistsError) MarshalJSON() ([]byte, error) {
	type Alias AttributeDefinitionAlreadyExistsError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "AttributeDefinitionAlreadyExists", Alias: (*Alias)(&obj)})
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

func (obj *AttributeDefinitionAlreadyExistsError) DecodeStruct(src map[string]interface{}) error {
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

func (obj AttributeDefinitionAlreadyExistsError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown AttributeDefinitionAlreadyExistsError: failed to parse error response"
}

type AttributeDefinitionTypeConflictError struct {
	Message                    string                 `json:"message"`
	ExtraValues                map[string]interface{} `json:"-"`
	ConflictingProductTypeId   string                 `json:"conflictingProductTypeId"`
	ConflictingProductTypeName string                 `json:"conflictingProductTypeName"`
	ConflictingAttributeName   string                 `json:"conflictingAttributeName"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AttributeDefinitionTypeConflictError) UnmarshalJSON(data []byte) error {
	type Alias AttributeDefinitionTypeConflictError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "conflictingProductTypeId")
	delete(obj.ExtraValues, "conflictingProductTypeName")
	delete(obj.ExtraValues, "conflictingAttributeName")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AttributeDefinitionTypeConflictError) MarshalJSON() ([]byte, error) {
	type Alias AttributeDefinitionTypeConflictError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "AttributeDefinitionTypeConflict", Alias: (*Alias)(&obj)})
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

func (obj *AttributeDefinitionTypeConflictError) DecodeStruct(src map[string]interface{}) error {
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

func (obj AttributeDefinitionTypeConflictError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown AttributeDefinitionTypeConflictError: failed to parse error response"
}

type AttributeNameDoesNotExistError struct {
	Message              string                 `json:"message"`
	ExtraValues          map[string]interface{} `json:"-"`
	InvalidAttributeName string                 `json:"invalidAttributeName"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AttributeNameDoesNotExistError) UnmarshalJSON(data []byte) error {
	type Alias AttributeNameDoesNotExistError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "invalidAttributeName")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AttributeNameDoesNotExistError) MarshalJSON() ([]byte, error) {
	type Alias AttributeNameDoesNotExistError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "AttributeNameDoesNotExist", Alias: (*Alias)(&obj)})
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

func (obj *AttributeNameDoesNotExistError) DecodeStruct(src map[string]interface{}) error {
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

func (obj AttributeNameDoesNotExistError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown AttributeNameDoesNotExistError: failed to parse error response"
}

type BadGatewayError struct {
	Message     string                 `json:"message"`
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

type ConcurrentModificationError struct {
	Message        string                 `json:"message"`
	ExtraValues    map[string]interface{} `json:"-"`
	CurrentVersion *int                   `json:"currentVersion,omitempty"`
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

type DiscountCodeNonApplicableError struct {
	Message           string                 `json:"message"`
	ExtraValues       map[string]interface{} `json:"-"`
	DiscountCode      *string                `json:"discountCode,omitempty"`
	Reason            *string                `json:"reason,omitempty"`
	DicountCodeId     *string                `json:"dicountCodeId,omitempty"`
	ValidFrom         *time.Time             `json:"validFrom,omitempty"`
	ValidUntil        *time.Time             `json:"validUntil,omitempty"`
	ValidityCheckTime *time.Time             `json:"validityCheckTime,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DiscountCodeNonApplicableError) UnmarshalJSON(data []byte) error {
	type Alias DiscountCodeNonApplicableError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "discountCode")
	delete(obj.ExtraValues, "reason")
	delete(obj.ExtraValues, "dicountCodeId")
	delete(obj.ExtraValues, "validFrom")
	delete(obj.ExtraValues, "validUntil")
	delete(obj.ExtraValues, "validityCheckTime")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeNonApplicableError) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeNonApplicableError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "DiscountCodeNonApplicable", Alias: (*Alias)(&obj)})
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

func (obj *DiscountCodeNonApplicableError) DecodeStruct(src map[string]interface{}) error {
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

func (obj DiscountCodeNonApplicableError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown DiscountCodeNonApplicableError: failed to parse error response"
}

type DuplicateAttributeValueError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
	Attribute   Attribute              `json:"attribute"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DuplicateAttributeValueError) UnmarshalJSON(data []byte) error {
	type Alias DuplicateAttributeValueError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "attribute")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DuplicateAttributeValueError) MarshalJSON() ([]byte, error) {
	type Alias DuplicateAttributeValueError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "DuplicateAttributeValue", Alias: (*Alias)(&obj)})
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

func (obj *DuplicateAttributeValueError) DecodeStruct(src map[string]interface{}) error {
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

func (obj DuplicateAttributeValueError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown DuplicateAttributeValueError: failed to parse error response"
}

type DuplicateAttributeValuesError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
	Attributes  []Attribute            `json:"attributes"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DuplicateAttributeValuesError) UnmarshalJSON(data []byte) error {
	type Alias DuplicateAttributeValuesError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "attributes")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DuplicateAttributeValuesError) MarshalJSON() ([]byte, error) {
	type Alias DuplicateAttributeValuesError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "DuplicateAttributeValues", Alias: (*Alias)(&obj)})
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

func (obj *DuplicateAttributeValuesError) DecodeStruct(src map[string]interface{}) error {
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

func (obj DuplicateAttributeValuesError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown DuplicateAttributeValuesError: failed to parse error response"
}

type DuplicateEnumValuesError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
	Duplicates  []string               `json:"duplicates"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DuplicateEnumValuesError) UnmarshalJSON(data []byte) error {
	type Alias DuplicateEnumValuesError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "duplicates")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DuplicateEnumValuesError) MarshalJSON() ([]byte, error) {
	type Alias DuplicateEnumValuesError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "DuplicateEnumValues", Alias: (*Alias)(&obj)})
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

func (obj *DuplicateEnumValuesError) DecodeStruct(src map[string]interface{}) error {
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

func (obj DuplicateEnumValuesError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown DuplicateEnumValuesError: failed to parse error response"
}

type DuplicateFieldError struct {
	Message        string                 `json:"message"`
	ExtraValues    map[string]interface{} `json:"-"`
	Field          *string                `json:"field,omitempty"`
	DuplicateValue interface{}            `json:"duplicateValue,omitempty"`
	// A Reference represents a loose reference to another resource in the same Project identified by its `id`. The `typeId` indicates the type of the referenced resource. Each resource type has its corresponding Reference type, like [ChannelReference](ctp:api:type:ChannelReference).  A referenced resource can be embedded through [Reference Expansion](/general-concepts#reference-expansion). The expanded reference is the value of an additional `obj` field then.
	ConflictingResource Reference `json:"conflictingResource,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DuplicateFieldError) UnmarshalJSON(data []byte) error {
	type Alias DuplicateFieldError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.ConflictingResource != nil {
		var err error
		obj.ConflictingResource, err = mapDiscriminatorReference(obj.ConflictingResource)
		if err != nil {
			return err
		}
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "field")
	delete(obj.ExtraValues, "duplicateValue")
	delete(obj.ExtraValues, "conflictingResource")

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

type DuplicateFieldWithConflictingResourceError struct {
	Message        string                 `json:"message"`
	ExtraValues    map[string]interface{} `json:"-"`
	Field          string                 `json:"field"`
	DuplicateValue interface{}            `json:"duplicateValue"`
	// A Reference represents a loose reference to another resource in the same Project identified by its `id`. The `typeId` indicates the type of the referenced resource. Each resource type has its corresponding Reference type, like [ChannelReference](ctp:api:type:ChannelReference).  A referenced resource can be embedded through [Reference Expansion](/general-concepts#reference-expansion). The expanded reference is the value of an additional `obj` field then.
	ConflictingResource Reference `json:"conflictingResource"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DuplicateFieldWithConflictingResourceError) UnmarshalJSON(data []byte) error {
	type Alias DuplicateFieldWithConflictingResourceError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.ConflictingResource != nil {
		var err error
		obj.ConflictingResource, err = mapDiscriminatorReference(obj.ConflictingResource)
		if err != nil {
			return err
		}
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "field")
	delete(obj.ExtraValues, "duplicateValue")
	delete(obj.ExtraValues, "conflictingResource")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DuplicateFieldWithConflictingResourceError) MarshalJSON() ([]byte, error) {
	type Alias DuplicateFieldWithConflictingResourceError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "DuplicateFieldWithConflictingResource", Alias: (*Alias)(&obj)})
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

func (obj *DuplicateFieldWithConflictingResourceError) DecodeStruct(src map[string]interface{}) error {
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

func (obj DuplicateFieldWithConflictingResourceError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown DuplicateFieldWithConflictingResourceError: failed to parse error response"
}

type DuplicatePriceScopeError struct {
	Message           string                 `json:"message"`
	ExtraValues       map[string]interface{} `json:"-"`
	ConflictingPrices []Price                `json:"conflictingPrices"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DuplicatePriceScopeError) UnmarshalJSON(data []byte) error {
	type Alias DuplicatePriceScopeError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "conflictingPrices")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DuplicatePriceScopeError) MarshalJSON() ([]byte, error) {
	type Alias DuplicatePriceScopeError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "DuplicatePriceScope", Alias: (*Alias)(&obj)})
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

func (obj *DuplicatePriceScopeError) DecodeStruct(src map[string]interface{}) error {
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

func (obj DuplicatePriceScopeError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown DuplicatePriceScopeError: failed to parse error response"
}

type DuplicateStandalonePriceScopeError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
	// [Reference](/../api/types#reference) to a [StandalonePrice](ctp:api:type:StandalonePrice).
	ConflictingStandalonePrice StandalonePriceReference `json:"conflictingStandalonePrice"`
	Sku                        string                   `json:"sku"`
	Currency                   string                   `json:"currency"`
	Country                    *string                  `json:"country,omitempty"`
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [CustomerGroup](ctp:api:type:CustomerGroup).
	CustomerGroup *CustomerGroupResourceIdentifier `json:"customerGroup,omitempty"`
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [Channel](ctp:api:type:Channel).
	Channel    *ChannelResourceIdentifier `json:"channel,omitempty"`
	ValidFrom  *time.Time                 `json:"validFrom,omitempty"`
	ValidUntil *time.Time                 `json:"validUntil,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DuplicateStandalonePriceScopeError) UnmarshalJSON(data []byte) error {
	type Alias DuplicateStandalonePriceScopeError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "conflictingStandalonePrice")
	delete(obj.ExtraValues, "sku")
	delete(obj.ExtraValues, "currency")
	delete(obj.ExtraValues, "country")
	delete(obj.ExtraValues, "customerGroup")
	delete(obj.ExtraValues, "channel")
	delete(obj.ExtraValues, "validFrom")
	delete(obj.ExtraValues, "validUntil")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DuplicateStandalonePriceScopeError) MarshalJSON() ([]byte, error) {
	type Alias DuplicateStandalonePriceScopeError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "DuplicateStandalonePriceScope", Alias: (*Alias)(&obj)})
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

func (obj *DuplicateStandalonePriceScopeError) DecodeStruct(src map[string]interface{}) error {
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

func (obj DuplicateStandalonePriceScopeError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown DuplicateStandalonePriceScopeError: failed to parse error response"
}

type DuplicateVariantValuesError struct {
	Message       string                 `json:"message"`
	ExtraValues   map[string]interface{} `json:"-"`
	VariantValues VariantValues          `json:"variantValues"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DuplicateVariantValuesError) UnmarshalJSON(data []byte) error {
	type Alias DuplicateVariantValuesError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "variantValues")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DuplicateVariantValuesError) MarshalJSON() ([]byte, error) {
	type Alias DuplicateVariantValuesError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "DuplicateVariantValues", Alias: (*Alias)(&obj)})
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

func (obj *DuplicateVariantValuesError) DecodeStruct(src map[string]interface{}) error {
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

func (obj DuplicateVariantValuesError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown DuplicateVariantValuesError: failed to parse error response"
}

type EditPreviewFailedError struct {
	Message     string                  `json:"message"`
	ExtraValues map[string]interface{}  `json:"-"`
	Result      OrderEditPreviewFailure `json:"result"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *EditPreviewFailedError) UnmarshalJSON(data []byte) error {
	type Alias EditPreviewFailedError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "result")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj EditPreviewFailedError) MarshalJSON() ([]byte, error) {
	type Alias EditPreviewFailedError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "EditPreviewFailed", Alias: (*Alias)(&obj)})
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

func (obj *EditPreviewFailedError) DecodeStruct(src map[string]interface{}) error {
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

func (obj EditPreviewFailedError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown EditPreviewFailedError: failed to parse error response"
}

type EnumKeyAlreadyExistsError struct {
	Message                  string                 `json:"message"`
	ExtraValues              map[string]interface{} `json:"-"`
	ConflictingEnumKey       string                 `json:"conflictingEnumKey"`
	ConflictingAttributeName string                 `json:"conflictingAttributeName"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *EnumKeyAlreadyExistsError) UnmarshalJSON(data []byte) error {
	type Alias EnumKeyAlreadyExistsError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "conflictingEnumKey")
	delete(obj.ExtraValues, "conflictingAttributeName")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj EnumKeyAlreadyExistsError) MarshalJSON() ([]byte, error) {
	type Alias EnumKeyAlreadyExistsError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "EnumKeyAlreadyExists", Alias: (*Alias)(&obj)})
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

func (obj *EnumKeyAlreadyExistsError) DecodeStruct(src map[string]interface{}) error {
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

func (obj EnumKeyAlreadyExistsError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown EnumKeyAlreadyExistsError: failed to parse error response"
}

type EnumKeyDoesNotExistError struct {
	Message                  string                 `json:"message"`
	ExtraValues              map[string]interface{} `json:"-"`
	ConflictingEnumKey       string                 `json:"conflictingEnumKey"`
	ConflictingAttributeName string                 `json:"conflictingAttributeName"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *EnumKeyDoesNotExistError) UnmarshalJSON(data []byte) error {
	type Alias EnumKeyDoesNotExistError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "conflictingEnumKey")
	delete(obj.ExtraValues, "conflictingAttributeName")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj EnumKeyDoesNotExistError) MarshalJSON() ([]byte, error) {
	type Alias EnumKeyDoesNotExistError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "EnumKeyDoesNotExist", Alias: (*Alias)(&obj)})
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

func (obj *EnumKeyDoesNotExistError) DecodeStruct(src map[string]interface{}) error {
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

func (obj EnumKeyDoesNotExistError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown EnumKeyDoesNotExistError: failed to parse error response"
}

type EnumValueIsUsedError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *EnumValueIsUsedError) UnmarshalJSON(data []byte) error {
	type Alias EnumValueIsUsedError
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
func (obj EnumValueIsUsedError) MarshalJSON() ([]byte, error) {
	type Alias EnumValueIsUsedError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "EnumValueIsUsed", Alias: (*Alias)(&obj)})
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

func (obj *EnumValueIsUsedError) DecodeStruct(src map[string]interface{}) error {
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

func (obj EnumValueIsUsedError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown EnumValueIsUsedError: failed to parse error response"
}

type EnumValuesMustMatchError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *EnumValuesMustMatchError) UnmarshalJSON(data []byte) error {
	type Alias EnumValuesMustMatchError
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
func (obj EnumValuesMustMatchError) MarshalJSON() ([]byte, error) {
	type Alias EnumValuesMustMatchError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "EnumValuesMustMatch", Alias: (*Alias)(&obj)})
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

func (obj *EnumValuesMustMatchError) DecodeStruct(src map[string]interface{}) error {
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

func (obj EnumValuesMustMatchError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown EnumValuesMustMatchError: failed to parse error response"
}

type ErrorResponse struct {
	StatusCode       int           `json:"statusCode"`
	Message          string        `json:"message"`
	ErrorMessage     *string       `json:"error,omitempty"`
	ErrorDescription *string       `json:"error_description,omitempty"`
	Errors           []ErrorObject `json:"errors"`
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

type ExtensionBadResponseError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
	// JSON object where the keys are of type [Locale](ctp:api:type:Locale), and the values are the strings used for the corresponding language.
	LocalizedMessage   *LocalizedString `json:"localizedMessage,omitempty"`
	ExtensionExtraInfo *interface{}     `json:"extensionExtraInfo,omitempty"`
	ErrorByExtension   ErrorByExtension `json:"errorByExtension"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ExtensionBadResponseError) UnmarshalJSON(data []byte) error {
	type Alias ExtensionBadResponseError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "localizedMessage")
	delete(obj.ExtraValues, "extensionExtraInfo")
	delete(obj.ExtraValues, "errorByExtension")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ExtensionBadResponseError) MarshalJSON() ([]byte, error) {
	type Alias ExtensionBadResponseError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ExtensionBadResponse", Alias: (*Alias)(&obj)})
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

func (obj *ExtensionBadResponseError) DecodeStruct(src map[string]interface{}) error {
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

func (obj ExtensionBadResponseError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ExtensionBadResponseError: failed to parse error response"
}

type ExtensionNoResponseError struct {
	Message      string                 `json:"message"`
	ExtraValues  map[string]interface{} `json:"-"`
	ExtensionId  string                 `json:"extensionId"`
	ExtensionKey *string                `json:"extensionKey,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ExtensionNoResponseError) UnmarshalJSON(data []byte) error {
	type Alias ExtensionNoResponseError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "extensionId")
	delete(obj.ExtraValues, "extensionKey")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ExtensionNoResponseError) MarshalJSON() ([]byte, error) {
	type Alias ExtensionNoResponseError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ExtensionNoResponse", Alias: (*Alias)(&obj)})
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

func (obj *ExtensionNoResponseError) DecodeStruct(src map[string]interface{}) error {
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

func (obj ExtensionNoResponseError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ExtensionNoResponseError: failed to parse error response"
}

type ExtensionUpdateActionsFailedError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
	// JSON object where the keys are of type [Locale](ctp:api:type:Locale), and the values are the strings used for the corresponding language.
	LocalizedMessage   *LocalizedString `json:"localizedMessage,omitempty"`
	ExtensionExtraInfo *interface{}     `json:"extensionExtraInfo,omitempty"`
	ErrorByExtension   ErrorByExtension `json:"errorByExtension"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ExtensionUpdateActionsFailedError) UnmarshalJSON(data []byte) error {
	type Alias ExtensionUpdateActionsFailedError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "localizedMessage")
	delete(obj.ExtraValues, "extensionExtraInfo")
	delete(obj.ExtraValues, "errorByExtension")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ExtensionUpdateActionsFailedError) MarshalJSON() ([]byte, error) {
	type Alias ExtensionUpdateActionsFailedError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ExtensionUpdateActionsFailed", Alias: (*Alias)(&obj)})
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

func (obj *ExtensionUpdateActionsFailedError) DecodeStruct(src map[string]interface{}) error {
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

func (obj ExtensionUpdateActionsFailedError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ExtensionUpdateActionsFailedError: failed to parse error response"
}

type ExternalOAuthFailedError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ExternalOAuthFailedError) UnmarshalJSON(data []byte) error {
	type Alias ExternalOAuthFailedError
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
func (obj ExternalOAuthFailedError) MarshalJSON() ([]byte, error) {
	type Alias ExternalOAuthFailedError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ExternalOAuthFailed", Alias: (*Alias)(&obj)})
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

func (obj *ExternalOAuthFailedError) DecodeStruct(src map[string]interface{}) error {
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

func (obj ExternalOAuthFailedError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ExternalOAuthFailedError: failed to parse error response"
}

type FeatureRemovedError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *FeatureRemovedError) UnmarshalJSON(data []byte) error {
	type Alias FeatureRemovedError
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
func (obj FeatureRemovedError) MarshalJSON() ([]byte, error) {
	type Alias FeatureRemovedError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "FeatureRemoved", Alias: (*Alias)(&obj)})
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

func (obj *FeatureRemovedError) DecodeStruct(src map[string]interface{}) error {
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

func (obj FeatureRemovedError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown FeatureRemovedError: failed to parse error response"
}

type GeneralError struct {
	Message     string                 `json:"message"`
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

type InsufficientScopeError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *InsufficientScopeError) UnmarshalJSON(data []byte) error {
	type Alias InsufficientScopeError
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
func (obj InsufficientScopeError) MarshalJSON() ([]byte, error) {
	type Alias InsufficientScopeError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "insufficient_scope", Alias: (*Alias)(&obj)})
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

func (obj *InsufficientScopeError) DecodeStruct(src map[string]interface{}) error {
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

func (obj InsufficientScopeError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown InsufficientScopeError: failed to parse error response"
}

type InternalConstraintViolatedError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *InternalConstraintViolatedError) UnmarshalJSON(data []byte) error {
	type Alias InternalConstraintViolatedError
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
func (obj InternalConstraintViolatedError) MarshalJSON() ([]byte, error) {
	type Alias InternalConstraintViolatedError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "InternalConstraintViolated", Alias: (*Alias)(&obj)})
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

func (obj *InternalConstraintViolatedError) DecodeStruct(src map[string]interface{}) error {
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

func (obj InternalConstraintViolatedError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown InternalConstraintViolatedError: failed to parse error response"
}

type InvalidCredentialsError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *InvalidCredentialsError) UnmarshalJSON(data []byte) error {
	type Alias InvalidCredentialsError
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
func (obj InvalidCredentialsError) MarshalJSON() ([]byte, error) {
	type Alias InvalidCredentialsError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "InvalidCredentials", Alias: (*Alias)(&obj)})
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

func (obj *InvalidCredentialsError) DecodeStruct(src map[string]interface{}) error {
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

func (obj InvalidCredentialsError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown InvalidCredentialsError: failed to parse error response"
}

type InvalidCurrentPasswordError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *InvalidCurrentPasswordError) UnmarshalJSON(data []byte) error {
	type Alias InvalidCurrentPasswordError
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
func (obj InvalidCurrentPasswordError) MarshalJSON() ([]byte, error) {
	type Alias InvalidCurrentPasswordError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "InvalidCurrentPassword", Alias: (*Alias)(&obj)})
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

func (obj *InvalidCurrentPasswordError) DecodeStruct(src map[string]interface{}) error {
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

func (obj InvalidCurrentPasswordError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown InvalidCurrentPasswordError: failed to parse error response"
}

type InvalidFieldError struct {
	Message       string                 `json:"message"`
	ExtraValues   map[string]interface{} `json:"-"`
	Field         string                 `json:"field"`
	InvalidValue  interface{}            `json:"invalidValue"`
	AllowedValues []interface{}          `json:"allowedValues"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *InvalidFieldError) UnmarshalJSON(data []byte) error {
	type Alias InvalidFieldError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "field")
	delete(obj.ExtraValues, "invalidValue")
	delete(obj.ExtraValues, "allowedValues")

	return nil
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

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *InvalidFieldError) DecodeStruct(src map[string]interface{}) error {
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

func (obj InvalidFieldError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown InvalidFieldError: failed to parse error response"
}

type InvalidInputError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *InvalidInputError) UnmarshalJSON(data []byte) error {
	type Alias InvalidInputError
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
func (obj InvalidInputError) MarshalJSON() ([]byte, error) {
	type Alias InvalidInputError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "InvalidInput", Alias: (*Alias)(&obj)})
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

func (obj *InvalidInputError) DecodeStruct(src map[string]interface{}) error {
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

func (obj InvalidInputError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown InvalidInputError: failed to parse error response"
}

type InvalidItemShippingDetailsError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
	Subject     string                 `json:"subject"`
	ItemId      string                 `json:"itemId"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *InvalidItemShippingDetailsError) UnmarshalJSON(data []byte) error {
	type Alias InvalidItemShippingDetailsError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "subject")
	delete(obj.ExtraValues, "itemId")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InvalidItemShippingDetailsError) MarshalJSON() ([]byte, error) {
	type Alias InvalidItemShippingDetailsError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "InvalidItemShippingDetails", Alias: (*Alias)(&obj)})
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

func (obj *InvalidItemShippingDetailsError) DecodeStruct(src map[string]interface{}) error {
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

func (obj InvalidItemShippingDetailsError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown InvalidItemShippingDetailsError: failed to parse error response"
}

type InvalidJsonInputError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
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

type InvalidOperationError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *InvalidOperationError) UnmarshalJSON(data []byte) error {
	type Alias InvalidOperationError
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
func (obj InvalidOperationError) MarshalJSON() ([]byte, error) {
	type Alias InvalidOperationError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "InvalidOperation", Alias: (*Alias)(&obj)})
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

func (obj *InvalidOperationError) DecodeStruct(src map[string]interface{}) error {
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

func (obj InvalidOperationError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown InvalidOperationError: failed to parse error response"
}

type InvalidSubjectError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *InvalidSubjectError) UnmarshalJSON(data []byte) error {
	type Alias InvalidSubjectError
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
func (obj InvalidSubjectError) MarshalJSON() ([]byte, error) {
	type Alias InvalidSubjectError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "InvalidSubject", Alias: (*Alias)(&obj)})
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

func (obj *InvalidSubjectError) DecodeStruct(src map[string]interface{}) error {
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

func (obj InvalidSubjectError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown InvalidSubjectError: failed to parse error response"
}

type InvalidTokenError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *InvalidTokenError) UnmarshalJSON(data []byte) error {
	type Alias InvalidTokenError
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
func (obj InvalidTokenError) MarshalJSON() ([]byte, error) {
	type Alias InvalidTokenError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "invalid_token", Alias: (*Alias)(&obj)})
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

func (obj *InvalidTokenError) DecodeStruct(src map[string]interface{}) error {
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

func (obj InvalidTokenError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown InvalidTokenError: failed to parse error response"
}

type LanguageUsedInStoresError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *LanguageUsedInStoresError) UnmarshalJSON(data []byte) error {
	type Alias LanguageUsedInStoresError
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
func (obj LanguageUsedInStoresError) MarshalJSON() ([]byte, error) {
	type Alias LanguageUsedInStoresError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "LanguageUsedInStores", Alias: (*Alias)(&obj)})
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

func (obj *LanguageUsedInStoresError) DecodeStruct(src map[string]interface{}) error {
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

func (obj LanguageUsedInStoresError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown LanguageUsedInStoresError: failed to parse error response"
}

type MatchingPriceNotFoundError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
	ProductId   string                 `json:"productId"`
	VariantId   int                    `json:"variantId"`
	Currency    *string                `json:"currency,omitempty"`
	Country     *string                `json:"country,omitempty"`
	// [Reference](ctp:api:type:Reference) to a [CustomerGroup](ctp:api:type:CustomerGroup).
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
	// [Reference](ctp:api:type:Reference) to a [Channel](ctp:api:type:Channel).
	Channel *ChannelReference `json:"channel,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *MatchingPriceNotFoundError) UnmarshalJSON(data []byte) error {
	type Alias MatchingPriceNotFoundError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "productId")
	delete(obj.ExtraValues, "variantId")
	delete(obj.ExtraValues, "currency")
	delete(obj.ExtraValues, "country")
	delete(obj.ExtraValues, "customerGroup")
	delete(obj.ExtraValues, "channel")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MatchingPriceNotFoundError) MarshalJSON() ([]byte, error) {
	type Alias MatchingPriceNotFoundError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "MatchingPriceNotFound", Alias: (*Alias)(&obj)})
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

func (obj *MatchingPriceNotFoundError) DecodeStruct(src map[string]interface{}) error {
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

func (obj MatchingPriceNotFoundError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown MatchingPriceNotFoundError: failed to parse error response"
}

type MaxResourceLimitExceededError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
	// Type of resource the value should reference. Supported resource type identifiers are:
	ExceededResource ReferenceTypeId `json:"exceededResource"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *MaxResourceLimitExceededError) UnmarshalJSON(data []byte) error {
	type Alias MaxResourceLimitExceededError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "exceededResource")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MaxResourceLimitExceededError) MarshalJSON() ([]byte, error) {
	type Alias MaxResourceLimitExceededError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "MaxResourceLimitExceeded", Alias: (*Alias)(&obj)})
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

func (obj *MaxResourceLimitExceededError) DecodeStruct(src map[string]interface{}) error {
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

func (obj MaxResourceLimitExceededError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown MaxResourceLimitExceededError: failed to parse error response"
}

type MissingRoleOnChannelError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [Channel](ctp:api:type:Channel).
	Channel *ChannelResourceIdentifier `json:"channel,omitempty"`
	// Describes the purpose and type of the Channel. A Channel can have one or more roles.
	MissingRole ChannelRoleEnum `json:"missingRole"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *MissingRoleOnChannelError) UnmarshalJSON(data []byte) error {
	type Alias MissingRoleOnChannelError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "channel")
	delete(obj.ExtraValues, "missingRole")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MissingRoleOnChannelError) MarshalJSON() ([]byte, error) {
	type Alias MissingRoleOnChannelError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "MissingRoleOnChannel", Alias: (*Alias)(&obj)})
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

func (obj *MissingRoleOnChannelError) DecodeStruct(src map[string]interface{}) error {
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

func (obj MissingRoleOnChannelError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown MissingRoleOnChannelError: failed to parse error response"
}

type MissingTaxRateForCountryError struct {
	Message       string                 `json:"message"`
	ExtraValues   map[string]interface{} `json:"-"`
	TaxCategoryId string                 `json:"taxCategoryId"`
	Country       *string                `json:"country,omitempty"`
	State         *string                `json:"state,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *MissingTaxRateForCountryError) UnmarshalJSON(data []byte) error {
	type Alias MissingTaxRateForCountryError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "taxCategoryId")
	delete(obj.ExtraValues, "country")
	delete(obj.ExtraValues, "state")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MissingTaxRateForCountryError) MarshalJSON() ([]byte, error) {
	type Alias MissingTaxRateForCountryError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "MissingTaxRateForCountry", Alias: (*Alias)(&obj)})
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

func (obj *MissingTaxRateForCountryError) DecodeStruct(src map[string]interface{}) error {
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

func (obj MissingTaxRateForCountryError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown MissingTaxRateForCountryError: failed to parse error response"
}

type NoMatchingProductDiscountFoundError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *NoMatchingProductDiscountFoundError) UnmarshalJSON(data []byte) error {
	type Alias NoMatchingProductDiscountFoundError
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
func (obj NoMatchingProductDiscountFoundError) MarshalJSON() ([]byte, error) {
	type Alias NoMatchingProductDiscountFoundError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "NoMatchingProductDiscountFound", Alias: (*Alias)(&obj)})
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

func (obj *NoMatchingProductDiscountFoundError) DecodeStruct(src map[string]interface{}) error {
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

func (obj NoMatchingProductDiscountFoundError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown NoMatchingProductDiscountFoundError: failed to parse error response"
}

type NotEnabledError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *NotEnabledError) UnmarshalJSON(data []byte) error {
	type Alias NotEnabledError
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
func (obj NotEnabledError) MarshalJSON() ([]byte, error) {
	type Alias NotEnabledError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "NotEnabled", Alias: (*Alias)(&obj)})
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

func (obj *NotEnabledError) DecodeStruct(src map[string]interface{}) error {
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

func (obj NotEnabledError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown NotEnabledError: failed to parse error response"
}

type ObjectNotFoundError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ObjectNotFoundError) UnmarshalJSON(data []byte) error {
	type Alias ObjectNotFoundError
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
func (obj ObjectNotFoundError) MarshalJSON() ([]byte, error) {
	type Alias ObjectNotFoundError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ObjectNotFound", Alias: (*Alias)(&obj)})
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

func (obj *ObjectNotFoundError) DecodeStruct(src map[string]interface{}) error {
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

func (obj ObjectNotFoundError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ObjectNotFoundError: failed to parse error response"
}

type OutOfStockError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
	LineItems   []string               `json:"lineItems"`
	Skus        []string               `json:"skus"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OutOfStockError) UnmarshalJSON(data []byte) error {
	type Alias OutOfStockError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "lineItems")
	delete(obj.ExtraValues, "skus")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OutOfStockError) MarshalJSON() ([]byte, error) {
	type Alias OutOfStockError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "OutOfStock", Alias: (*Alias)(&obj)})
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

func (obj *OutOfStockError) DecodeStruct(src map[string]interface{}) error {
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

func (obj OutOfStockError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown OutOfStockError: failed to parse error response"
}

type OverCapacityError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OverCapacityError) UnmarshalJSON(data []byte) error {
	type Alias OverCapacityError
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
func (obj OverCapacityError) MarshalJSON() ([]byte, error) {
	type Alias OverCapacityError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "OverCapacity", Alias: (*Alias)(&obj)})
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

func (obj *OverCapacityError) DecodeStruct(src map[string]interface{}) error {
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

func (obj OverCapacityError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown OverCapacityError: failed to parse error response"
}

type OverlappingStandalonePriceValidityError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
	// [Reference](/../api/types#reference) to a [StandalonePrice](ctp:api:type:StandalonePrice).
	ConflictingStandalonePrice StandalonePriceReference `json:"conflictingStandalonePrice"`
	Sku                        string                   `json:"sku"`
	Currency                   string                   `json:"currency"`
	Country                    *string                  `json:"country,omitempty"`
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [CustomerGroup](ctp:api:type:CustomerGroup).
	CustomerGroup *CustomerGroupResourceIdentifier `json:"customerGroup,omitempty"`
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [Channel](ctp:api:type:Channel).
	Channel               *ChannelResourceIdentifier `json:"channel,omitempty"`
	ValidFrom             *time.Time                 `json:"validFrom,omitempty"`
	ValidUntil            *time.Time                 `json:"validUntil,omitempty"`
	ConflictingValidFrom  *time.Time                 `json:"conflictingValidFrom,omitempty"`
	ConflictingValidUntil *time.Time                 `json:"conflictingValidUntil,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OverlappingStandalonePriceValidityError) UnmarshalJSON(data []byte) error {
	type Alias OverlappingStandalonePriceValidityError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "conflictingStandalonePrice")
	delete(obj.ExtraValues, "sku")
	delete(obj.ExtraValues, "currency")
	delete(obj.ExtraValues, "country")
	delete(obj.ExtraValues, "customerGroup")
	delete(obj.ExtraValues, "channel")
	delete(obj.ExtraValues, "validFrom")
	delete(obj.ExtraValues, "validUntil")
	delete(obj.ExtraValues, "conflictingValidFrom")
	delete(obj.ExtraValues, "conflictingValidUntil")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OverlappingStandalonePriceValidityError) MarshalJSON() ([]byte, error) {
	type Alias OverlappingStandalonePriceValidityError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "OverlappingStandalonePriceValidity", Alias: (*Alias)(&obj)})
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

func (obj *OverlappingStandalonePriceValidityError) DecodeStruct(src map[string]interface{}) error {
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

func (obj OverlappingStandalonePriceValidityError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown OverlappingStandalonePriceValidityError: failed to parse error response"
}

type PendingOperationError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *PendingOperationError) UnmarshalJSON(data []byte) error {
	type Alias PendingOperationError
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
func (obj PendingOperationError) MarshalJSON() ([]byte, error) {
	type Alias PendingOperationError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "PendingOperation", Alias: (*Alias)(&obj)})
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

func (obj *PendingOperationError) DecodeStruct(src map[string]interface{}) error {
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

func (obj PendingOperationError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown PendingOperationError: failed to parse error response"
}

type PriceChangedError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
	LineItems   []string               `json:"lineItems"`
	Shipping    bool                   `json:"shipping"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *PriceChangedError) UnmarshalJSON(data []byte) error {
	type Alias PriceChangedError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "lineItems")
	delete(obj.ExtraValues, "shipping")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PriceChangedError) MarshalJSON() ([]byte, error) {
	type Alias PriceChangedError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "PriceChanged", Alias: (*Alias)(&obj)})
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

func (obj *PriceChangedError) DecodeStruct(src map[string]interface{}) error {
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

func (obj PriceChangedError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown PriceChangedError: failed to parse error response"
}

type ProjectNotConfiguredForLanguagesError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
	Languages   []string               `json:"languages"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProjectNotConfiguredForLanguagesError) UnmarshalJSON(data []byte) error {
	type Alias ProjectNotConfiguredForLanguagesError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "languages")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProjectNotConfiguredForLanguagesError) MarshalJSON() ([]byte, error) {
	type Alias ProjectNotConfiguredForLanguagesError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ProjectNotConfiguredForLanguages", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["languages"] == nil {
		delete(raw, "languages")
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *ProjectNotConfiguredForLanguagesError) DecodeStruct(src map[string]interface{}) error {
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

func (obj ProjectNotConfiguredForLanguagesError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ProjectNotConfiguredForLanguagesError: failed to parse error response"
}

type QueryComplexityLimitExceededError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *QueryComplexityLimitExceededError) UnmarshalJSON(data []byte) error {
	type Alias QueryComplexityLimitExceededError
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
func (obj QueryComplexityLimitExceededError) MarshalJSON() ([]byte, error) {
	type Alias QueryComplexityLimitExceededError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "QueryComplexityLimitExceeded", Alias: (*Alias)(&obj)})
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

func (obj *QueryComplexityLimitExceededError) DecodeStruct(src map[string]interface{}) error {
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

func (obj QueryComplexityLimitExceededError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown QueryComplexityLimitExceededError: failed to parse error response"
}

type QueryTimedOutError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *QueryTimedOutError) UnmarshalJSON(data []byte) error {
	type Alias QueryTimedOutError
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
func (obj QueryTimedOutError) MarshalJSON() ([]byte, error) {
	type Alias QueryTimedOutError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "QueryTimedOut", Alias: (*Alias)(&obj)})
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

func (obj *QueryTimedOutError) DecodeStruct(src map[string]interface{}) error {
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

func (obj QueryTimedOutError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown QueryTimedOutError: failed to parse error response"
}

type ReferenceExistsError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
	// Type of resource the value should reference. Supported resource type identifiers are:
	ReferencedBy *ReferenceTypeId `json:"referencedBy,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ReferenceExistsError) UnmarshalJSON(data []byte) error {
	type Alias ReferenceExistsError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "referencedBy")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ReferenceExistsError) MarshalJSON() ([]byte, error) {
	type Alias ReferenceExistsError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ReferenceExists", Alias: (*Alias)(&obj)})
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

func (obj *ReferenceExistsError) DecodeStruct(src map[string]interface{}) error {
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

func (obj ReferenceExistsError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ReferenceExistsError: failed to parse error response"
}

type ReferencedResourceNotFoundError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
	// Type of resource the value should reference. Supported resource type identifiers are:
	TypeId ReferenceTypeId `json:"typeId"`
	ID     *string         `json:"id,omitempty"`
	Key    *string         `json:"key,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ReferencedResourceNotFoundError) UnmarshalJSON(data []byte) error {
	type Alias ReferencedResourceNotFoundError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "typeId")
	delete(obj.ExtraValues, "id")
	delete(obj.ExtraValues, "key")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ReferencedResourceNotFoundError) MarshalJSON() ([]byte, error) {
	type Alias ReferencedResourceNotFoundError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ReferencedResourceNotFound", Alias: (*Alias)(&obj)})
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

func (obj *ReferencedResourceNotFoundError) DecodeStruct(src map[string]interface{}) error {
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

func (obj ReferencedResourceNotFoundError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ReferencedResourceNotFoundError: failed to parse error response"
}

type RequiredFieldError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
	Field       string                 `json:"field"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *RequiredFieldError) UnmarshalJSON(data []byte) error {
	type Alias RequiredFieldError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "field")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RequiredFieldError) MarshalJSON() ([]byte, error) {
	type Alias RequiredFieldError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "RequiredField", Alias: (*Alias)(&obj)})
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

func (obj *RequiredFieldError) DecodeStruct(src map[string]interface{}) error {
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

func (obj RequiredFieldError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown RequiredFieldError: failed to parse error response"
}

type ResourceNotFoundError struct {
	Message     string                 `json:"message"`
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

type ResourceSizeLimitExceededError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ResourceSizeLimitExceededError) UnmarshalJSON(data []byte) error {
	type Alias ResourceSizeLimitExceededError
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
func (obj ResourceSizeLimitExceededError) MarshalJSON() ([]byte, error) {
	type Alias ResourceSizeLimitExceededError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ResourceSizeLimitExceeded", Alias: (*Alias)(&obj)})
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

func (obj *ResourceSizeLimitExceededError) DecodeStruct(src map[string]interface{}) error {
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

func (obj ResourceSizeLimitExceededError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ResourceSizeLimitExceededError: failed to parse error response"
}

type SearchDeactivatedError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *SearchDeactivatedError) UnmarshalJSON(data []byte) error {
	type Alias SearchDeactivatedError
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
func (obj SearchDeactivatedError) MarshalJSON() ([]byte, error) {
	type Alias SearchDeactivatedError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "SearchDeactivated", Alias: (*Alias)(&obj)})
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

func (obj *SearchDeactivatedError) DecodeStruct(src map[string]interface{}) error {
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

func (obj SearchDeactivatedError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown SearchDeactivatedError: failed to parse error response"
}

type SearchExecutionFailureError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *SearchExecutionFailureError) UnmarshalJSON(data []byte) error {
	type Alias SearchExecutionFailureError
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
func (obj SearchExecutionFailureError) MarshalJSON() ([]byte, error) {
	type Alias SearchExecutionFailureError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "SearchExecutionFailure", Alias: (*Alias)(&obj)})
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

func (obj *SearchExecutionFailureError) DecodeStruct(src map[string]interface{}) error {
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

func (obj SearchExecutionFailureError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown SearchExecutionFailureError: failed to parse error response"
}

type SearchFacetPathNotFoundError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *SearchFacetPathNotFoundError) UnmarshalJSON(data []byte) error {
	type Alias SearchFacetPathNotFoundError
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
func (obj SearchFacetPathNotFoundError) MarshalJSON() ([]byte, error) {
	type Alias SearchFacetPathNotFoundError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "SearchFacetPathNotFound", Alias: (*Alias)(&obj)})
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

func (obj *SearchFacetPathNotFoundError) DecodeStruct(src map[string]interface{}) error {
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

func (obj SearchFacetPathNotFoundError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown SearchFacetPathNotFoundError: failed to parse error response"
}

type SearchIndexingInProgressError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *SearchIndexingInProgressError) UnmarshalJSON(data []byte) error {
	type Alias SearchIndexingInProgressError
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
func (obj SearchIndexingInProgressError) MarshalJSON() ([]byte, error) {
	type Alias SearchIndexingInProgressError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "SearchIndexingInProgress", Alias: (*Alias)(&obj)})
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

func (obj *SearchIndexingInProgressError) DecodeStruct(src map[string]interface{}) error {
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

func (obj SearchIndexingInProgressError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown SearchIndexingInProgressError: failed to parse error response"
}

type SemanticErrorError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *SemanticErrorError) UnmarshalJSON(data []byte) error {
	type Alias SemanticErrorError
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
func (obj SemanticErrorError) MarshalJSON() ([]byte, error) {
	type Alias SemanticErrorError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "SemanticError", Alias: (*Alias)(&obj)})
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

func (obj *SemanticErrorError) DecodeStruct(src map[string]interface{}) error {
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

func (obj SemanticErrorError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown SemanticErrorError: failed to parse error response"
}

type ShippingMethodDoesNotMatchCartError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ShippingMethodDoesNotMatchCartError) UnmarshalJSON(data []byte) error {
	type Alias ShippingMethodDoesNotMatchCartError
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
func (obj ShippingMethodDoesNotMatchCartError) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodDoesNotMatchCartError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ShippingMethodDoesNotMatchCart", Alias: (*Alias)(&obj)})
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

func (obj *ShippingMethodDoesNotMatchCartError) DecodeStruct(src map[string]interface{}) error {
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

func (obj ShippingMethodDoesNotMatchCartError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ShippingMethodDoesNotMatchCartError: failed to parse error response"
}

type SyntaxErrorError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *SyntaxErrorError) UnmarshalJSON(data []byte) error {
	type Alias SyntaxErrorError
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
func (obj SyntaxErrorError) MarshalJSON() ([]byte, error) {
	type Alias SyntaxErrorError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "SyntaxError", Alias: (*Alias)(&obj)})
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

func (obj *SyntaxErrorError) DecodeStruct(src map[string]interface{}) error {
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

func (obj SyntaxErrorError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown SyntaxErrorError: failed to parse error response"
}

type VariantValues struct {
	Sku        *string      `json:"sku,omitempty"`
	Prices     []PriceDraft `json:"prices"`
	Attributes []Attribute  `json:"attributes"`
}

type WeakPasswordError struct {
	Message     string                 `json:"message"`
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *WeakPasswordError) UnmarshalJSON(data []byte) error {
	type Alias WeakPasswordError
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
func (obj WeakPasswordError) MarshalJSON() ([]byte, error) {
	type Alias WeakPasswordError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "WeakPassword", Alias: (*Alias)(&obj)})
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

func (obj *WeakPasswordError) DecodeStruct(src map[string]interface{}) error {
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

func (obj WeakPasswordError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown WeakPasswordError: failed to parse error response"
}
