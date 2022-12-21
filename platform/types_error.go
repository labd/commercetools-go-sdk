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

/**
*	Represents a single error. Multiple errors may be included in an [ErrorResponse](ctp:api:type:ErrorResponse).
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
	case "CountryNotConfiguredInStore":
		obj := CountryNotConfiguredInStore{}
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
	case "DuplicatePriceKey":
		obj := DuplicatePriceKeyError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
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
	case "ExtensionPredicateEvaluationFailed":
		obj := ExtensionPredicateEvaluationFailedError{}
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
	case "ProductAssignmentMissing":
		obj := ProductAssignmentMissingError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ProductPresentWithDifferentVariantSelection":
		obj := ProductPresentWithDifferentVariantSelectionError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.ExistingVariantSelection != nil {
			var err error
			obj.ExistingVariantSelection, err = mapDiscriminatorProductVariantSelection(obj.ExistingVariantSelection)
			if err != nil {
				return nil, err
			}
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
	}
	return nil, nil
}

/**
*	Returned when the anonymous ID is being used by another resource.
*
*	The client application should choose another anonymous ID or retrieve an automatically generated one.
*
 */
type AnonymousIdAlreadyInUseError struct {
	// `"The given anonymous ID is already in use."`
	Message string `json:"message"`
	// Error-specific additional fields.
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

/**
*	Returned when the `name` of the [AttributeDefinition](ctp:api:type:AttributeDefinition) conflicts with an existing Attribute.
*
*	The error is returned as a failed response to the [Create ProductType](/../api/projects/productTypes#create-producttype) request or [Change AttributeDefinition Name](ctp:api:type:ProductTypeChangeAttributeNameAction) update action.
*
 */
type AttributeDefinitionAlreadyExistsError struct {
	// `"An attribute definition with name $attributeName already exists on product type $productTypeName."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// Unique identifier of the Product Type containing the conflicting name.
	ConflictingProductTypeId string `json:"conflictingProductTypeId"`
	// Name of the Product Type containing the conflicting name.
	ConflictingProductTypeName string `json:"conflictingProductTypeName"`
	// Name of the conflicting Attribute.
	ConflictingAttributeName string `json:"conflictingAttributeName"`
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

/**
*	Returned when the `type` is different for an AttributeDefinition using the same `name` in multiple Product Types.
*
*	The error is returned as a failed response to the [Create ProductType](/../api/projects/productTypes#create-producttype) request.
*
 */
type AttributeDefinitionTypeConflictError struct {
	// `"The attribute with name $attributeName has a different type on product type $productTypeName."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// Unique identifier of the Product Type containing the conflicting name.
	ConflictingProductTypeId string `json:"conflictingProductTypeId"`
	// Name of the Product Type containing the conflicting name.
	ConflictingProductTypeName string `json:"conflictingProductTypeName"`
	// Name of the conflicting Attribute.
	ConflictingAttributeName string `json:"conflictingAttributeName"`
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

/**
*	Returned when an [AttributeDefinition](ctp:api:type:AttributeDefinition) does not exist for an Attribute `name`.
*
*	The error is returned as a failed response to the [Change AttributeDefinition Name](ctp:api:type:ProductTypeChangeAttributeNameAction) update action.
*
 */
type AttributeNameDoesNotExistError struct {
	// `"Attribute definition for $attributeName does not exist on type $typeName."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// Non-existent Attribute name.
	InvalidAttributeName string `json:"invalidAttributeName"`
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
	// `"Object $resourceId has a different version than expected. Expected: $expectedVersion - Actual: $currentVersion."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
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
*	Returned when a [Cart](ctp:api:type:Cart) or an [Order](ctp:api:type:Order) in a [Store](ctp:api:type:Store) references a country that is not included in the countries configured for the Store.
*
*	The error is returned as a failed response to:
*
*	- [Create Cart in Store](/../api/projects/carts#create-a-cart-in-a-store) request and [Set Country](/../api/projects/carts#set-country) update action on Carts.
*	- [Create Order in a Store from a Cart](/../api/projects/me-orders#create-order-in-a-store-from-a-cart) requests and [Set Country](/../api/projects/me-carts#set-country) on My Orders.
*	- [Create Order by Import](/../api/projects/orders#create-order-by-import) request.
*
 */
type CountryNotConfiguredInStore struct {
	// `"The country $country is not configured for the store $store."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// Countries configured for the Store.
	StoreCountries []string `json:"storeCountries"`
	// The country that is not configured for the Store but referenced on the Cart or Order.
	Country string `json:"country"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CountryNotConfiguredInStore) UnmarshalJSON(data []byte) error {
	type Alias CountryNotConfiguredInStore
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "storeCountries")
	delete(obj.ExtraValues, "country")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CountryNotConfiguredInStore) MarshalJSON() ([]byte, error) {
	type Alias CountryNotConfiguredInStore
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "CountryNotConfiguredInStore", Alias: (*Alias)(&obj)})
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

func (obj *CountryNotConfiguredInStore) DecodeStruct(src map[string]interface{}) error {
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
*	Returned when the Cart contains a Discount Code with a [DiscountCodeState](ctp:api:type:DiscountCodeState) other than `MatchesCart`.
*
*	The error is returned as a failed response to:
*
*	- [Create Order from Cart](/../api/projects/orders#create-order-from-cart) and [Create Order from Cart in a Store](/../api/projects/orders#create-order-from-cart-in-a-store) requests on Orders.
*	- [Create Order from Cart](/../api/projects/me-orders#create-order-from-a-cart) and [Create Order in a Store from a Cart](/../api/projects/me-orders#create-order-in-a-store-from-a-cart) requests on My Orders.
*
 */
type DiscountCodeNonApplicableError struct {
	// `"The discountCode $discountCodeId cannot be applied to the cart."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// Discount Code passed to the Cart.
	DiscountCode *string `json:"discountCode,omitempty"`
	// `"DoesNotExist"` or `"TimeRangeNonApplicable"`
	Reason *string `json:"reason,omitempty"`
	// Unique identifier of the Discount Code.
	DiscountCodeId *string `json:"discountCodeId,omitempty"`
	// Date and time (UTC) from which the Discount Code is valid.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Date and time (UTC) until which the Discount Code is valid.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
	// Date and time (UTC) the Discount Code validity check was last performed.
	ValidityCheckTime *time.Time `json:"validityCheckTime,omitempty"`
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
	delete(obj.ExtraValues, "discountCodeId")
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

/**
*	Returned when the `Unique` [AttributeConstraint](ctp:api:type:AttributeConstraintEnum) criteria are not met during an [Update Product](/../api/projects/products#update-product) request.
*
 */
type DuplicateAttributeValueError struct {
	// `"Attribute can't have the same value in a different variant."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// Conflicting Attributes.
	Attribute Attribute `json:"attribute"`
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

/**
*	Returned when the `CombinationUnique` [AttributeConstraint](ctp:api:type:AttributeConstraintEnum) criteria are not met during an [Update Product](/../api/projects/products#update-product) request.
*
 */
type DuplicateAttributeValuesError struct {
	// `"The set of attributes must be unique across all variants."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// Conflicting Attributes.
	Attributes []Attribute `json:"attributes"`
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

/**
*	Returned when an [AttributeEnumType](ctp:api:type:AttributeEnumType) or [AttributeLocalizedEnumType](ctp:api:type:AttributeLocalizedEnumType) contains duplicate keys.
*
 */
type DuplicateEnumValuesError struct {
	// `"The enum values contain duplicate keys: $listOfDuplicateKeys."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// Duplicate keys.
	Duplicates []string `json:"duplicates"`
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
*	Returned when a field value conflicts with an existing value stored in a particular resource causing a duplicate.
*
 */
type DuplicateFieldWithConflictingResourceError struct {
	// `"A duplicate value $duplicateValue exists for field $field on $conflictingResource."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// Name of the conflicting field.
	Field string `json:"field"`
	// Conflicting duplicate value.
	DuplicateValue interface{} `json:"duplicateValue"`
	// Reference to the resource that has the conflicting value.
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

/**
*	Returned when a Price key conflicts with an existing key.
*
*	Keys of Embedded Prices must be unique per ProductVariant.
*
 */
type DuplicatePriceKeyError struct {
	// `"Duplicate price key: $priceKey. The price key must be unique per variant."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// Conflicting Embedded Price.
	ConflictingPrice Price `json:"conflictingPrice"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DuplicatePriceKeyError) UnmarshalJSON(data []byte) error {
	type Alias DuplicatePriceKeyError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "conflictingPrice")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DuplicatePriceKeyError) MarshalJSON() ([]byte, error) {
	type Alias DuplicatePriceKeyError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "DuplicatePriceKey", Alias: (*Alias)(&obj)})
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

func (obj *DuplicatePriceKeyError) DecodeStruct(src map[string]interface{}) error {
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

func (obj DuplicatePriceKeyError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown DuplicatePriceKeyError: failed to parse error response"
}

/**
*	Returned when a Price scope conflicts with an existing one during an [Update Product](/../api/projects/products#update-product) request.
*
*	Every Price of a Product Variant must have a distinct combination of currency, Customer Group, country, and Channel that constitute the scope of a Price.
*
 */
type DuplicatePriceScopeError struct {
	// `"Duplicate price scope: $priceScope. The combination of currency, country, customerGroup and channel must be unique for each price of a product variant."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// Conflicting Embedded Price.
	ConflictingPrice Price `json:"conflictingPrice"`
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
	delete(obj.ExtraValues, "conflictingPrice")

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

/**
*	Returned when the given Price scope conflicts with the Price scope of an existing Standalone Price.
*	Every Standalone Price associated with the same SKU must have a distinct combination of currency, country, Customer Group, Channel, and validity periods (`validFrom` and `validUntil`).
*
*	The error is returned as a failed response to the [Create StandalonePrice](/../api/projects/standalone-prices#create-standaloneprice) request.
*
 */
type DuplicateStandalonePriceScopeError struct {
	// `"Duplicate standalone price scope for SKU: $sku. The combination of SKU, currency, country, customerGroup, channel, validFrom and validUntil must be unique for each standalone price."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// Reference to the conflicting Standalone Price.
	ConflictingStandalonePrice StandalonePriceReference `json:"conflictingStandalonePrice"`
	// SKU of the [ProductVariant](ctp:api:type:ProductVariant) to which the conflicting Standalone Price is associated.
	Sku string `json:"sku"`
	// Currency code of the country.
	Currency string `json:"currency"`
	// Country code of the geographic location.
	Country *string `json:"country,omitempty"`
	// [CustomerGroup](ctp:api:type:CustomerGroup) for which the Standalone Price is valid.
	CustomerGroup *CustomerGroupResourceIdentifier `json:"customerGroup,omitempty"`
	// [Channel](ctp:api:type:Channel) for which the Standalone Price is valid.
	Channel *ChannelResourceIdentifier `json:"channel,omitempty"`
	// Date and time (UTC) from which the Standalone Price is valid.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Date and time (UTC) until which the Standalone Price is valid.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
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

/**
*	Returned when a [Product Variant](ctp:api:type:ProductVariant) value conflicts with an existing one during an [Update Product](/../api/projects/products#update-product) request.
*
 */
type DuplicateVariantValuesError struct {
	// `"A duplicate combination of the variant values (sku, key, images, prices, attributes) exists."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// Every Product Variant must have a distinct combination of SKU, prices, and custom Attribute values.
	VariantValues VariantValues `json:"variantValues"`
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

/**
*	Returned when a preview to find an appropriate Shipping Method for an OrderEdit could not be generated.
*
*	The error is returned as a failed response to the [Get Shipping Methods for an OrderEdit](/../api/projects/shippingMethods#get-shippingmethods-for-an-orderedit) request.
*
 */
type EditPreviewFailedError struct {
	// `"Error while applying staged actions. ShippingMethods could not be determined."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// State of the OrderEdit where the `stagedActions` cannot be applied to the Order.
	Result OrderEditPreviewFailure `json:"result"`
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

/**
*	Returned when an [AttributeEnumType](ctp:api:type:AttributeEnumType) or [AttributeLocalizedEnumType](ctp:api:type:AttributeLocalizedEnumType) contains a key that already exists.
*
 */
type EnumKeyAlreadyExistsError struct {
	// `"The $attributeName attribute definition already contains an enum value with the key $enumKey."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// Conflicting enum key.
	ConflictingEnumKey string `json:"conflictingEnumKey"`
	// Name of the conflicting Attribute.
	ConflictingAttributeName string `json:"conflictingAttributeName"`
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

/**
*	Returned when an [AttributeEnumType](ctp:api:type:AttributeEnumType) or [AttributeLocalizedEnumType](ctp:api:type:AttributeLocalizedEnumType) already contains a value with the given key.
*
*	The error is returned as a failed response to the [Change the key of an EnumValue](ctp:api:type:ProductTypeChangeEnumKeyAction) update action.
*
 */
type EnumKeyDoesNotExistError struct {
	// `"The $fieldName field definition does not contain an enum value with the key $enumKey."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// Conflicting enum key.
	ConflictingEnumKey string `json:"conflictingEnumKey"`
	// Name of the conflicting Attribute.
	ConflictingAttributeName string `json:"conflictingAttributeName"`
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

/**
*	Returned when an enum value cannot be removed from an Attribute as it is being used by a Product.
*
*	The error is returned as a failed response to the [Remove EnumValues from AttributeDefinition](ctp:api:type:ProductTypeRemoveEnumValuesAction) update action.
*
 */
type EnumValueIsUsedError struct {
	// `"$enumKeysTranscript is used by some products and cannot be deleted because the $attributeName attribute is required."`
	Message string `json:"message"`
	// Error-specific additional fields.
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

/**
*	Returned when during an order update of [AttributeEnumType](ctp:api:type:AttributeEnumType) or [AttributeLocalizedEnumType](ctp:api:type:AttributeLocalizedEnumType) the new enum values do not match the existing ones.
*
*	The error is returned as a failed response to the [Change the order of EnumValues](ctp:api:type:ProductTypeChangePlainEnumValueOrderAction) and [Change the order of LocalizedEnumValues](ctp:api:type:ProductTypeChangeLocalizedEnumValueOrderAction) update actions.
*
 */
type EnumValuesMustMatchError struct {
	// `"The given values must be equal to the existing enum values."`
	Message string `json:"message"`
	// Error-specific additional fields.
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
*	Represents errors related to authentication and authorization in a format conforming to the [OAuth 2.0 specification](https://tools.ietf.org/html/rfc6749#section-5.2).
*
 */
type AuthErrorResponse struct {
	// HTTP status code corresponding to the error.
	StatusCode int `json:"statusCode"`
	// First error message in the `errors` array.
	Message string `json:"message"`
	// Authentication and authorization-related errors returned for a request.
	Errors []ErrorObject `json:"errors"`
	// Error code as per the [OAuth 2.0 specification](https://tools.ietf.org/html/rfc6749#section-5.2). For example: `"access_denied"`.
	ErrorMessage string `json:"error"`
	// Plain text description of the first error.
	ErrorDescription *string `json:"error_description,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AuthErrorResponse) UnmarshalJSON(data []byte) error {
	type Alias AuthErrorResponse
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

func (obj AuthErrorResponse) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown AuthErrorResponse: failed to parse error response"
}

/**
*	Returned when the response from the API Extension could not be parsed successfully (such as a `500` HTTP status code, or an invalid JSON response).
*
 */
type ExtensionBadResponseError struct {
	// Description of the invalid Extension response. For example, `"The extension did not return the expected JSON."`.
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// User-defined localized description of the error.
	LocalizedMessage *LocalizedString `json:"localizedMessage,omitempty"`
	// Any information that should be returned to the API caller.
	ExtensionExtraInfo *interface{} `json:"extensionExtraInfo,omitempty"`
	// Additional errors related to the API Extension.
	ExtensionErrors []ExtensionError `json:"extensionErrors"`
	// The response body returned by the Extension.
	ExtensionBody *string `json:"extensionBody,omitempty"`
	// Http status code returned by the Extension.
	ExtensionStatusCode *int `json:"extensionStatusCode,omitempty"`
	// Unique identifier of the Extension.
	ExtensionId string `json:"extensionId"`
	// User-defined unique identifier of the Extension.
	ExtensionKey *string `json:"extensionKey,omitempty"`
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
	delete(obj.ExtraValues, "extensionErrors")
	delete(obj.ExtraValues, "extensionBody")
	delete(obj.ExtraValues, "extensionStatusCode")
	delete(obj.ExtraValues, "extensionId")
	delete(obj.ExtraValues, "extensionKey")

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

type ExtensionError struct {
	// Error code caused by the Extension. For example, `InvalidField`.
	Code string `json:"code"`
	// Plain text description of the error.
	Message string `json:"message"`
	// Unique identifier of the Extension.
	ExtensionId string `json:"extensionId"`
	// User-defined unique identifier of the Extension.
	ExtensionKey *string `json:"extensionKey,omitempty"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ExtensionError) UnmarshalJSON(data []byte) error {
	type Alias ExtensionError
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
func (obj ExtensionError) MarshalJSON() ([]byte, error) {
	type Alias ExtensionError
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

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj ExtensionError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ExtensionError: failed to parse error response"
}

/**
*	Returned when the API Extension does not respond within the [time limit](/../api/projects/api-extensions#time-limits), or could not be reached.
*
 */
type ExtensionNoResponseError struct {
	// `"Extension did not respond in time."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// Unique identifier of the API Extension.
	ExtensionId string `json:"extensionId"`
	// User-defined unique identifier of the API Extension, if available.
	ExtensionKey *string `json:"extensionKey,omitempty"`
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

/**
*	Returned when the predicate defined in the [ExtensionTrigger](ctp:api:type:ExtensionTrigger) could not be evaluated due to a missing field.
*
 */
type ExtensionPredicateEvaluationFailedError struct {
	// `"The compared field $fieldName is not present."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// Details about the API Extension that was involved in the error.
	ErrorByExtension ErrorByExtension `json:"errorByExtension"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ExtensionPredicateEvaluationFailedError) UnmarshalJSON(data []byte) error {
	type Alias ExtensionPredicateEvaluationFailedError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "errorByExtension")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ExtensionPredicateEvaluationFailedError) MarshalJSON() ([]byte, error) {
	type Alias ExtensionPredicateEvaluationFailedError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ExtensionPredicateEvaluationFailed", Alias: (*Alias)(&obj)})
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

func (obj *ExtensionPredicateEvaluationFailedError) DecodeStruct(src map[string]interface{}) error {
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

func (obj ExtensionPredicateEvaluationFailedError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ExtensionPredicateEvaluationFailedError: failed to parse error response"
}

/**
*	Returned when update actions could not be applied to the resource (for example, because a referenced resource does not exist).
*	This would result in a [400 Bad Request](#400-bad-request) response if the same update action was sent from a regular client.
*
 */
type ExtensionUpdateActionsFailedError struct {
	// `"The extension returned update actions that could not be executed."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// User-defined localized description of the error.
	LocalizedMessage *LocalizedString `json:"localizedMessage,omitempty"`
	// Any information that should be returned to the API caller.
	ExtensionExtraInfo *interface{} `json:"extensionExtraInfo,omitempty"`
	// Additional errors related to the API Extension.
	ExtensionErrors []ExtensionError `json:"extensionErrors"`
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
	delete(obj.ExtraValues, "extensionErrors")

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

/**
*	Returned when an [external OAuth Introspection endpoint](/../api/authorization#requesting-an-access-token-using-an-external-oauth-server) does not return a response within the [time limit](/../api/authorization#time-limits), or the response isn't compliant with [RFC 7662](https://www.rfc-editor.org/rfc/rfc7662.html) (for example, an HTTP status code like `500`).
*
 */
type ExternalOAuthFailedError struct {
	// Plain text description detailing the external OAuth error. For example, `"External OAuth did not respond in time."`.
	Message string `json:"message"`
	// Error-specific additional fields.
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

/**
*	Returned when the requested feature was removed.
*
 */
type FeatureRemovedError struct {
	// Description of the feature that is removed.
	Message string `json:"message"`
	// Error-specific additional fields.
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

/**
*	Returned when a server-side problem occurs.
*
*	If you encounter this error, report it using the [Support Portal](https://support.commercetools.com).
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

type InsufficientScopeError struct {
	// Plain text description of the cause of the error.
	Message string `json:"message"`
	// Error-specific additional fields.
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

/**
*	Returned when certain API-specific constraints were not met. For example, the specified [Discount Code](ctp:api:type:DiscountCode) was never applied and cannot be updated.
*
 */
type InternalConstraintViolatedError struct {
	// Plain text description of the constraints that were violated.
	Message string `json:"message"`
	// Error-specific additional fields.
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

/**
*	Returned when a Customer with the given credentials (matching the given email/password pair) is not found and authentication fails.
*
*	The error is returned as a failed response to:
*
*	- [Authenticate a global Customer (Sign-in)](/../api/projects/customers#authenticate-sign-in-customer) and [Authenticate Customer (Sign-in) in a Store](/../api/projects/customers#authenticate-sign-in-customer-in-store) requests on Customers.
*	- [Authenticating Customer (Sign-in)](/../api/projects/me-profile#authenticate-sign-in-customer) and [Authenticate Customer (Sign-in) in a Store](/../api/projects/me-profile#authenticate-sign-in-customer-in-store) requests on My Customer Profile.
*
 */
type InvalidCredentialsError struct {
	// `"Account with the given credentials not found."`
	Message string `json:"message"`
	// Error-specific additional fields.
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

/**
*	Returned when the current password of the Customer does not match.
*
*	The error is returned as a failed response to:
*
*	- [Change Customer Password](/../api/projects/customers#change-password-of-customer) and [Change Customer Password in a Store](/../api/projects/customers#change-password-of-customer-in-store) requests on Customers.
*	- [Change Customer Password](/../api/projects/me-profile#change-password-of-customer) and [Change Customer Password in a Store](/../api/projects/me-profile#change-password-of-customer-in-store) requests on My Customer Profile.
*
 */
type InvalidCurrentPasswordError struct {
	// `"The given current password does not match."`
	Message string `json:"message"`
	// Error-specific additional fields.
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

/**
*	Returned when a field has an invalid value.
*
 */
type InvalidFieldError struct {
	// `"The value $invalidValue is not valid for field $field."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// Name of the field with the invalid value.
	Field string `json:"field"`
	// Value invalid for the field.
	InvalidValue interface{} `json:"invalidValue"`
	// Fixed set of allowed values for the field, if any.
	AllowedValues []interface{} `json:"allowedValues"`
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

/**
*	Returned when an invalid input has been sent.
*
 */
type InvalidInputError struct {
	// Description of the constraints that are not met by the request. For example, `"Invalid $propertyName. It may be a non-empty string up to $maxLength"`.
	Message string `json:"message"`
	// Error-specific additional fields.
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

/**
*	Returned when Line Item or Custom Line Item quantities set under [ItemShippingDetails](ctp:api:type:ItemShippingDetails) do not match the sum of the quantities in their respective shipping details.
*
*	The error is returned as a failed response to the [Create Order from Cart](/../api/projects/orders#create-order-from-cart) and [Create Order from Cart in a Store](/../api/projects/orders#create-order-from-cart-in-a-store) requests.
*
 */
type InvalidItemShippingDetailsError struct {
	// `"Inconsistent shipping details for $subject with ID $itemId. $subject quantity is $itemQuantity and shippingTargets quantity sum is $quantitySum."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// `"LineItem"` or `"CustomLineItem"`
	Subject string `json:"subject"`
	// Unique identifier of the Line Item or Custom Line Item.
	ItemId string `json:"itemId"`
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
*	Returned when the resources involved in the request are not in a valid state for the operation.
*
*	The client application should validate the constraints described in the error message before sending the request.
*
 */
type InvalidOperationError struct {
	// Plain text description of the error.
	Message string `json:"message"`
	// Error-specific additional fields.
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
	// Plain text description of the cause of the error.
	Message string `json:"message"`
	// Error-specific additional fields.
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
	// Plain text description of the cause of the error.
	Message string `json:"message"`
	// Error-specific additional fields.
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

/**
*	Returned when a language cannot be removed from a Project as it is being used by a Store.
*
*	The error is returned as a failed response to the [Change Languages](ctp:api:type:ProjectChangeLanguagesAction) update action.
*
 */
type LanguageUsedInStoresError struct {
	// `"Language(s) in use by a store cannot be deleted. Remove them in all the stores of this project first."`
	Message string `json:"message"`
	// Error-specific additional fields.
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

/**
*	Returned when the Product Variant does not have a Price according to the [Product](ctp:api:type:Product) `priceMode` value for a selected currency, country, Customer Group, or Channel.
*
*	The error is returned as a failed response to:
*
*	- [Add LineItem](ctp:api:type:CartAddLineItemAction), [Add CustomLineItem](ctp:api:type:CartAddCustomLineItemAction), and [Add DiscountCode](ctp:api:type:CartAddDiscountCodeAction) update actions on Carts.
*	- [Add LineItem](ctp:api:type:StagedOrderAddLineItemAction), [Add CustomLineItem](ctp:api:type:StagedOrderAddCustomLineItemAction), and [Add DiscountCode](ctp:api:type:StagedOrderAddDiscountCodeAction) update actions on Order Edits.
*	- [Create Order from Cart](/../api/projects/orders#create-order-from-cart) and [Create Order from Cart in a Store](/../api/projects/orders#create-order-from-cart-in-a-store) requests on Orders.
*	- [Create Order from a Cart](/../api/projects/me-orders#create-order-from-a-cart) and [Create Order in a Store from a Cart](/../api/projects/me-orders#create-order-in-a-store-from-a-cart) requests on My Orders.
*
 */
type MatchingPriceNotFoundError struct {
	// `"The variant $variantId of product $productId does not contain a price for currency $currencyCode, $country, $customerGroup, $channel."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// Unique identifier of a [Product](ctp:api:type:Product).
	ProductId string `json:"productId"`
	// Unique identifier of a [ProductVariant](ctp:api:type:ProductVariant) in the Product.
	VariantId int `json:"variantId"`
	// Currency code of the country.
	Currency *string `json:"currency,omitempty"`
	// Country code of the geographic location.
	Country *string `json:"country,omitempty"`
	// Customer Group associated with the Price.
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
	// Channel associated with the Price.
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

/**
*	Returned when a resource type cannot be created as it has reached its [limits](/../api/limits).
*
*	The limits must be adjusted for this resource before sending the request again.
*
 */
type MaxResourceLimitExceededError struct {
	// `"You have exceeded the limit of $limit resources of type $resourceTypeId."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// Resource type that reached its maximum limit of configured elements (for example, 100 Zones per Project).
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

/**
*	Returned when one of the following states occur:
*
*	- [Channel](ctp:api:Channel) is added or set on a [Store](ctp:api:Store) with missing Channel `roles`.
*	- [Standalone Price](/../api/projects/standalone-prices#create-standaloneprice) references a Channel that does not contain the `ProductDistribution` role.
*
*	The error is returned as a failed response to:
*
*	- [Add Distribution Channel](ctp:api:type:StoreAddDistributionChannelAction), [Set Distribution Channel](ctp:api:type:StoreSetDistributionChannelsAction), [Add Supply Channel](ctp:api:type:StoreAddSupplyChannelAction), and [Set Supply Channel](ctp:api:type:StoreSetSupplyChannelsAction) update actions.
*	- [Create a Standalone Price](/../api/projects/standalone-prices#create-standaloneprice) request.
*
 */
type MissingRoleOnChannelError struct {
	// `"Given channel with $idOrKeyOfChannel does not have the required role $role."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a given [Channel](ctp:api:type:Channel).
	Channel *ChannelResourceIdentifier `json:"channel,omitempty"`
	// - `ProductDistribution` for Product Distribution Channels allowed for the Store. Also required for [Standalone Prices](ctp:api:type:StandalonePrice).
	// - `InventorySupply` for Inventory Supply Channels allowed for the Store.
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

/**
*	Returned when the Tax Category of at least one of the `lineItems`, `customLineItems`, or `shippingInfo` in the [Cart](ctp:api:type:Cart) is missing the [TaxRate](ctp:api:type:TaxRate) matching `country` and `state` given in the `shippingAddress` of that Cart.
*
*	The error is returned as a failed response to:
*
*	- [Set Default Shipping Address](ctp:api:type:CustomerSetDefaultShippingAddressAction), [Add LineItem](ctp:api:type:CartAddLineItemAction), [Add CustomLineItem](ctp:api:type:CartAddCustomLineItemAction), [Set Shipping Address](ctp:api:type:CartSetShippingAddressAction), [Set Customer ID](ctp:api:type:CartSetCustomerIdAction), [Add LineItem](ctp:api:type:StagedOrderAddLineItemAction), and [Add CustomLineItem](ctp:api:type:StagedOrderAddCustomLineItemAction) update actions
*	- [Create Order from Cart](/../api/projects/orders#create-order-from-cart) and [Create Order from Cart in a Store](/../api/projects/orders#create-order-from-cart-in-a-store) requests.
*
 */
type MissingTaxRateForCountryError struct {
	// `"Tax category $taxCategoryId is missing a tax rate for country $countriesAndStates."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// Unique identifier of the [TaxCategory](ctp:api:type:TaxCategory).
	TaxCategoryId string `json:"taxCategoryId"`
	// Country code of the geographic location.
	Country *string `json:"country,omitempty"`
	// State within the country, such as Texas in the United States.
	State *string `json:"state,omitempty"`
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

/**
*	Returned when a Product Discount could not be found that could be applied to the Price of a Product Variant.
*
*	The error is returned as a failed response to the [Get Matching ProductDiscount](/../api/projects/productDiscounts#get-matching-productdiscount) request.
*
 */
type NoMatchingProductDiscountFoundError struct {
	// `"Couldn't find a matching product discount for: productId=$productId, variantId=$variantId, price=$price."`
	Message string `json:"message"`
	// Error-specific additional fields.
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

/**
*	Returned when the [Project-specific category recommendations feature](/../api/projects/categoryRecommendations#project-specific-category-recommendations) is not enabled for the Project.
*
 */
type NotEnabledError struct {
	// `"The category recommendations API is not yet enabled for your project."`
	Message string `json:"message"`
	// Error-specific additional fields.
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

/**
*	Returned when the requested resource was not found.
*
 */
type ObjectNotFoundError struct {
	// `"A $resourceType with identifier $id was unexpectedly not found."`
	Message string `json:"message"`
	// Error-specific additional fields.
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

/**
*	Returned when some of the [Line Items](ctp:api:type:LineItem) are out of stock at the time of placing an [Order](ctp:api:type:Order).
*
*	The error is returned as a failed response to:
*
*	- [Create Order from Cart](/../api/projects/orders#create-order-from-cart), [Create Order from Cart in a Store](/../api/projects/orders#create-order-from-cart-in-a-store), and [Create Order by Import](/../api/projects/me-orders) requests on Orders.
*	- [Create Order from a Cart](/../api/projects/me-orders#create-order-from-a-cart) and [Create Order in a Store from Cart](/../api/projects/me-orders#create-order-in-a-store-from-a-cart) requests on My Orders.
*
 */
type OutOfStockError struct {
	// `"Some line items are out of stock at the time of placing the order: $itemSku."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// Unique identifiers of the Line Items that are out of stock.
	LineItems []string `json:"lineItems"`
	// SKUs of the Line Items that are out of stock.
	Skus []string `json:"skus"`
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

/**
*	Returned when the service is having trouble handling the load.
*
*	The client application should retry the request with exponential backoff up to a point where further delay is unacceptable.
*
 */
type OverCapacityError struct {
	// Plain text description of the error.
	Message string `json:"message"`
	// Error-specific additional fields.
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

/**
*	Returned when a given Price validity period conflicts with an existing one.
*	Every Standalone Price associated with the same SKU and with the same combination of currency, country, Customer Group, and Channel, must have non-overlapping validity periods (`validFrom` and `validUntil`).
*
*	The error is returned as a failed response to the [Create StandalonePrice](/../api/projects/standalone-prices#create-standaloneprice) request.
*
 */
type OverlappingStandalonePriceValidityError struct {
	// `Two standalone prices have overlapping validity periods."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// Reference to the conflicting Standalone Price.
	ConflictingStandalonePrice StandalonePriceReference `json:"conflictingStandalonePrice"`
	// SKU of the [ProductVariant](ctp:api:type:ProductVariant) to which the conflicting Standalone Price is associated.
	Sku string `json:"sku"`
	// Currency code of the country.
	Currency string `json:"currency"`
	// Country code of the geographic location.
	Country *string `json:"country,omitempty"`
	// [CustomerGroup](ctp:api:type:CustomerGroup) for which the Standalone Price is valid.
	CustomerGroup *CustomerGroupResourceIdentifier `json:"customerGroup,omitempty"`
	// [Channel](ctp:api:type:Channel) for which the Standalone Price is valid.
	Channel *ChannelResourceIdentifier `json:"channel,omitempty"`
	// Date and time (UTC) from which the Standalone Price is valid.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Date and time (UTC) until which the Standalone Price is valid.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
	// Date and time (UTC) from which the conflicting Standalone Price is valid.
	ConflictingValidFrom *time.Time `json:"conflictingValidFrom,omitempty"`
	// Date and time (UTC) until which the conflicting Standalone Price is valid.
	ConflictingValidUntil *time.Time `json:"conflictingValidUntil,omitempty"`
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

/**
*	Returned when a previous conflicting operation is still pending and needs to finish before the request can succeed.
*
*	The client application should retry the request with exponential backoff up to a point where further delay is unacceptable.
*	If the error persists, report it using the [Support Portal](https://support.commercetools.com).
*
 */
type PendingOperationError struct {
	// Plain text description of the error.
	Message string `json:"message"`
	// Error-specific additional fields.
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

/**
*	Returned when the Price, Tax Rate, or Shipping Rate of some Line Items changed since they were last added to the Cart.
*
*	The error is returned as a failed response to:
*
*	- [Create Order from Cart](/../api/projects/orders#create-order-from-cart) and [Create Order from Cart in a Store](/../api/projects/orders#create-order-from-cart-in-a-store) requests on Orders.
*	- [Create Order from a Cart](/../api/projects/me-orders#create-order-from-a-cart) and [Create Order in a Store from a Cart](/../api/projects/me-orders#create-order-in-a-store-from-a-cart) requests on My Orders.
*
 */
type PriceChangedError struct {
	// Plain text description of the reason for the Price change. For example, `"The price or tax of some line items changed at the time of placing the order: $lineItems."`.
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// Unique identifiers of the Line Items for which the Price or [TaxRate](ctp:api:type:TaxRate) has changed.
	LineItems []string `json:"lineItems"`
	// `true` if the [ShippingRate](ctp:api:type:ShippingRate) has changed.
	Shipping bool `json:"shipping"`
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

/**
*	Returned when a Product is not assigned to the Product Selection.
*
*	The error is returned as a failed response to the [Set Variant Selection](ctp:api:type:ProductSelectionSetVariantSelectionAction) update action.
*
 */
type ProductAssignmentMissingError struct {
	// `"A Product Variant Selection can only be set for a Product previously added to the Product Selection."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// [Reference](ctp:api:type:Reference) to the [Product](ctp:api:type:Product) for which the error was returned.
	Product ProductReference `json:"product"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductAssignmentMissingError) UnmarshalJSON(data []byte) error {
	type Alias ProductAssignmentMissingError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "product")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductAssignmentMissingError) MarshalJSON() ([]byte, error) {
	type Alias ProductAssignmentMissingError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ProductAssignmentMissing", Alias: (*Alias)(&obj)})
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

func (obj *ProductAssignmentMissingError) DecodeStruct(src map[string]interface{}) error {
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

func (obj ProductAssignmentMissingError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ProductAssignmentMissingError: failed to parse error response"
}

/**
*	Returned when a Product is already assigned to a [Product Selection](/../api/projects/product-selections), but the Product Selection has a different [Product Variant Selection](ctp:api:type:ProductVariantSelection).
*
*	The error is returned as a failed response to the [Add Product](ctp:api:type:ProductSelectionAddProductAction) update action.
*
 */
type ProductPresentWithDifferentVariantSelectionError struct {
	// `"Product is already present with the following different $variantSelections."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// [Reference](ctp:api:type:Reference) to the [Product](ctp:api:type:Product) for which the error was returned.
	Product ProductReference `json:"product"`
	// Existing Product Variant Selection for the [Product](/../api/projects/products) in the [Product Selection](/../api/projects/product-selections).
	ExistingVariantSelection ProductVariantSelection `json:"existingVariantSelection"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductPresentWithDifferentVariantSelectionError) UnmarshalJSON(data []byte) error {
	type Alias ProductPresentWithDifferentVariantSelectionError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.ExistingVariantSelection != nil {
		var err error
		obj.ExistingVariantSelection, err = mapDiscriminatorProductVariantSelection(obj.ExistingVariantSelection)
		if err != nil {
			return err
		}
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")
	delete(obj.ExtraValues, "message")
	delete(obj.ExtraValues, "product")
	delete(obj.ExtraValues, "existingVariantSelection")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductPresentWithDifferentVariantSelectionError) MarshalJSON() ([]byte, error) {
	type Alias ProductPresentWithDifferentVariantSelectionError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ProductPresentWithDifferentVariantSelection", Alias: (*Alias)(&obj)})
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

func (obj *ProductPresentWithDifferentVariantSelectionError) DecodeStruct(src map[string]interface{}) error {
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

func (obj ProductPresentWithDifferentVariantSelectionError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ProductPresentWithDifferentVariantSelectionError: failed to parse error response"
}

/**
*	Returned when the languages set for a Store are not supported by the Project.
*
*	The error is returned as a failed response to the [Set Languages](ctp:api:type:StoreSetLanguagesAction) update action.
*
 */
type ProjectNotConfiguredForLanguagesError struct {
	// `"The project is not configured for given languages."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// Languages configured for the Store.
	Languages []string `json:"languages"`
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
	Message string `json:"message"`
	// Error-specific additional fields.
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

/**
*	Returned when the query times out.
*
*	If a query constantly times out, please check if it follows the [performance best practices](/../api/predicates/query#performance-considerations).
*
 */
type QueryTimedOutError struct {
	// `"The query timed out. If your query constantly times out, please check that it follows the performance best practices (see https://docs.commercetools.com/api/predicates/query#performance-considerations)."`
	Message string `json:"message"`
	// Error-specific additional fields.
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

/**
*	Returned when a resource cannot be deleted because it is being referenced by another resource.
*
 */
type ReferenceExistsError struct {
	// `"Can not delete a $resource while it is referenced by at least one $referencedBy."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// Type of referenced resource.
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

/**
*	Returned when a resource referenced by a [Reference](ctp:api:type:Reference) or a [ResourceIdentifier](ctp:api:type:ResourceIdentifier) could not be found.
*
 */
type ReferencedResourceNotFoundError struct {
	// `"The referenced object of type $typeId $predicate was not found. It either doesn't exist, or it can't be accessed from this endpoint (e.g., if the endpoint filters by store or customer account)."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// Type of referenced resource.
	TypeId ReferenceTypeId `json:"typeId"`
	// Unique identifier of the referenced resource, if known.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced resource, if known.
	Key *string `json:"key,omitempty"`
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

/**
*	Returned when a value is not defined for a required field.
*
 */
type RequiredFieldError struct {
	// `"A value is required for field $field."`
	Message string `json:"message"`
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
	// Name of the field missing the value.
	Field string `json:"field"`
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

/**
*	Returned when the resource addressed by the request URL does not exist.
*
 */
type ResourceNotFoundError struct {
	// `"The Resource with ID $resourceId was not found."`
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
*	Returned when the resource exceeds the maximum allowed size of 16 MB.
*
 */
type ResourceSizeLimitExceededError struct {
	// `"The resource size exceeds the maximal allowed size of 16 MB."`
	Message string `json:"message"`
	// Error-specific additional fields.
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

/**
*	Returned when the indexing of Product information is deactivated in a Project.
*
*	To activate indexing, call [Change Product Search Indexing Enabled](ctp:api:type:ProjectChangeProductSearchIndexingEnabledAction) and set `enabled` to `true`.
*
 */
type SearchDeactivatedError struct {
	// `"The endpoint is deactivated for this project. Please enable it via the Project endpoint, via the Merchant Center in the Project settings, or reach out to Support to enable it."`
	Message string `json:"message"`
	// Error-specific additional fields.
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

/**
*	Returned when a search query could not be completed due to an unexpected failure.
*
 */
type SearchExecutionFailureError struct {
	// `"Something went wrong during the search query execution. In most case this happens due to usage of non-existing fields and custom product attributes. Please verify all filters and facets in your search query and make sure that all paths are correct."`
	Message string `json:"message"`
	// Error-specific additional fields.
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

/**
*	Returned when a search facet path could not be found.
*
 */
type SearchFacetPathNotFoundError struct {
	// `"Facet path $path not found."`
	Message string `json:"message"`
	// Error-specific additional fields.
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

/**
*	Returned when the indexing of Product information is still in progress for Projects that have indexing activated.
*
 */
type SearchIndexingInProgressError struct {
	// `"The indexing is currently in progress. Please wait until the status is "Activated" to execute search requests."`
	Message string `json:"message"`
	// Error-specific additional fields.
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

/**
*	Returned when a [Discount predicate](/../api/predicates/predicate-operators) or [API Extension predicate](/../api/predicates/query#using-predicates-in-conditional-api-extensions) is not semantically correct.
*
 */
type SemanticErrorError struct {
	// Plain text description of the error concerning the predicate. For example, `"Invalid country code $countryCode provided for the field $fieldDefinition."`.
	Message string `json:"message"`
	// Error-specific additional fields.
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

/**
*	Returned when the Cart contains a [ShippingMethod](ctp:api:type:ShippingMethod) that is not allowed for the [Cart](ctp:api:type:Cart). In this case, the [ShippingMethodState](ctp:api:type:ShippingMethodState) value is `DoesNotMatchCart`.
*
*	The error is returned as a failed response to the [Create Order from Cart](/../api/projects/orders#create-order-from-cart) or [Create Order from Cart in a Store](/../api/projects/orders#create-order-from-cart-in-a-store) requests.
*
 */
type ShippingMethodDoesNotMatchCartError struct {
	// `"The predicate does not match the cart."`
	Message string `json:"message"`
	// Error-specific additional fields.
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

/**
*	Returned when a [Discount predicate](/../api/predicates/predicate-operators), [API Extension predicate](/../api/predicates/query#using-predicates-in-conditional-api-extensions), or [search query](/../api/projects/products-search) does not have the correct syntax.
*
 */
type SyntaxErrorError struct {
	// `"Syntax error while parsing $fieldDefinition."`
	Message string `json:"message"`
	// Error-specific additional fields.
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
	// SKU of the [ProductVariant](ctp:api:type:ProductVariant).
	Sku *string `json:"sku,omitempty"`
	// Embedded Prices of the [ProductVariant](ctp:api:type:ProductVariant).
	Prices []PriceDraft `json:"prices"`
	// Attributes of the [ProductVariant](ctp:api:type:ProductVariant).
	Attributes []Attribute `json:"attributes"`
}
