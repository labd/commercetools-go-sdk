package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type ErrorByExtension struct {
	ID  string  `json:"id"`
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
	Message string `json:"message"`
	// interface{} `json:"//"`
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

type AnonymousIdAlreadyInUseError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AnonymousIdAlreadyInUseError) MarshalJSON() ([]byte, error) {
	type Alias AnonymousIdAlreadyInUseError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "AnonymousIdAlreadyInUse", Alias: (*Alias)(&obj)})
}
func (obj AnonymousIdAlreadyInUseError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown AnonymousIdAlreadyInUseError: failed to parse error response"
}

type AttributeDefinitionAlreadyExistsError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
	ConflictingProductTypeId   string `json:"conflictingProductTypeId"`
	ConflictingProductTypeName string `json:"conflictingProductTypeName"`
	ConflictingAttributeName   string `json:"conflictingAttributeName"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AttributeDefinitionAlreadyExistsError) MarshalJSON() ([]byte, error) {
	type Alias AttributeDefinitionAlreadyExistsError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "AttributeDefinitionAlreadyExists", Alias: (*Alias)(&obj)})
}
func (obj AttributeDefinitionAlreadyExistsError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown AttributeDefinitionAlreadyExistsError: failed to parse error response"
}

type AttributeDefinitionTypeConflictError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
	ConflictingProductTypeId   string `json:"conflictingProductTypeId"`
	ConflictingProductTypeName string `json:"conflictingProductTypeName"`
	ConflictingAttributeName   string `json:"conflictingAttributeName"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AttributeDefinitionTypeConflictError) MarshalJSON() ([]byte, error) {
	type Alias AttributeDefinitionTypeConflictError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "AttributeDefinitionTypeConflict", Alias: (*Alias)(&obj)})
}
func (obj AttributeDefinitionTypeConflictError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown AttributeDefinitionTypeConflictError: failed to parse error response"
}

type AttributeNameDoesNotExistError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
	InvalidAttributeName string `json:"invalidAttributeName"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AttributeNameDoesNotExistError) MarshalJSON() ([]byte, error) {
	type Alias AttributeNameDoesNotExistError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "AttributeNameDoesNotExist", Alias: (*Alias)(&obj)})
}
func (obj AttributeNameDoesNotExistError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown AttributeNameDoesNotExistError: failed to parse error response"
}

type BadGatewayError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BadGatewayError) MarshalJSON() ([]byte, error) {
	type Alias BadGatewayError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "BadGateway", Alias: (*Alias)(&obj)})
}
func (obj BadGatewayError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown BadGatewayError: failed to parse error response"
}

type ConcurrentModificationError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
	CurrentVersion *int `json:"currentVersion,omitempty"`
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

type DiscountCodeNonApplicableError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
	DiscountCode      *string    `json:"discountCode,omitempty"`
	Reason            *string    `json:"reason,omitempty"`
	DicountCodeId     *string    `json:"dicountCodeId,omitempty"`
	ValidFrom         *time.Time `json:"validFrom,omitempty"`
	ValidUntil        *time.Time `json:"validUntil,omitempty"`
	ValidityCheckTime *time.Time `json:"validityCheckTime,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeNonApplicableError) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeNonApplicableError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "DiscountCodeNonApplicable", Alias: (*Alias)(&obj)})
}
func (obj DiscountCodeNonApplicableError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown DiscountCodeNonApplicableError: failed to parse error response"
}

type DuplicateAttributeValueError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
	Attribute Attribute `json:"attribute"`
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

type DuplicateAttributeValuesError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
	Attributes []Attribute `json:"attributes"`
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

type DuplicateEnumValuesError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
	Duplicates []string `json:"duplicates"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DuplicateEnumValuesError) MarshalJSON() ([]byte, error) {
	type Alias DuplicateEnumValuesError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "DuplicateEnumValues", Alias: (*Alias)(&obj)})
}
func (obj DuplicateEnumValuesError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown DuplicateEnumValuesError: failed to parse error response"
}

type DuplicateFieldError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
	Field               *string     `json:"field,omitempty"`
	DuplicateValue      interface{} `json:"duplicateValue,omitempty"`
	ConflictingResource Reference   `json:"conflictingResource,omitempty"`
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
	return nil
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

type DuplicateFieldWithConflictingResourceError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
	Field               string      `json:"field"`
	DuplicateValue      interface{} `json:"duplicateValue"`
	ConflictingResource Reference   `json:"conflictingResource"`
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
	return nil
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

type DuplicatePriceScopeError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
	ConflictingPrices []Price `json:"conflictingPrices"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DuplicatePriceScopeError) MarshalJSON() ([]byte, error) {
	type Alias DuplicatePriceScopeError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "DuplicatePriceScope", Alias: (*Alias)(&obj)})
}
func (obj DuplicatePriceScopeError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown DuplicatePriceScopeError: failed to parse error response"
}

type DuplicateVariantValuesError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
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

type EditPreviewFailedError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
	Result OrderEditPreviewFailure `json:"result"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj EditPreviewFailedError) MarshalJSON() ([]byte, error) {
	type Alias EditPreviewFailedError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "EditPreviewFailed", Alias: (*Alias)(&obj)})
}
func (obj EditPreviewFailedError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown EditPreviewFailedError: failed to parse error response"
}

type EnumKeyAlreadyExistsError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
	ConflictingEnumKey       string `json:"conflictingEnumKey"`
	ConflictingAttributeName string `json:"conflictingAttributeName"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj EnumKeyAlreadyExistsError) MarshalJSON() ([]byte, error) {
	type Alias EnumKeyAlreadyExistsError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "EnumKeyAlreadyExists", Alias: (*Alias)(&obj)})
}
func (obj EnumKeyAlreadyExistsError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown EnumKeyAlreadyExistsError: failed to parse error response"
}

type EnumKeyDoesNotExistError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
	ConflictingEnumKey       string `json:"conflictingEnumKey"`
	ConflictingAttributeName string `json:"conflictingAttributeName"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj EnumKeyDoesNotExistError) MarshalJSON() ([]byte, error) {
	type Alias EnumKeyDoesNotExistError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "EnumKeyDoesNotExist", Alias: (*Alias)(&obj)})
}
func (obj EnumKeyDoesNotExistError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown EnumKeyDoesNotExistError: failed to parse error response"
}

type EnumValueIsUsedError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj EnumValueIsUsedError) MarshalJSON() ([]byte, error) {
	type Alias EnumValueIsUsedError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "EnumValueIsUsed", Alias: (*Alias)(&obj)})
}
func (obj EnumValueIsUsedError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown EnumValueIsUsedError: failed to parse error response"
}

type EnumValuesMustMatchError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj EnumValuesMustMatchError) MarshalJSON() ([]byte, error) {
	type Alias EnumValuesMustMatchError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "EnumValuesMustMatch", Alias: (*Alias)(&obj)})
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

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["errors"] == nil {
		delete(target, "errors")
	}

	return json.Marshal(target)
}
func (obj ErrorResponse) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ErrorResponse: failed to parse error response"
}

type ExtensionBadResponseError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
	LocalizedMessage   *LocalizedString `json:"localizedMessage,omitempty"`
	ExtensionExtraInfo *interface{}     `json:"extensionExtraInfo,omitempty"`
	ErrorByExtension   ErrorByExtension `json:"errorByExtension"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ExtensionBadResponseError) MarshalJSON() ([]byte, error) {
	type Alias ExtensionBadResponseError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ExtensionBadResponse", Alias: (*Alias)(&obj)})
}
func (obj ExtensionBadResponseError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ExtensionBadResponseError: failed to parse error response"
}

type ExtensionNoResponseError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
	ExtensionId  string  `json:"extensionId"`
	ExtensionKey *string `json:"extensionKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ExtensionNoResponseError) MarshalJSON() ([]byte, error) {
	type Alias ExtensionNoResponseError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ExtensionNoResponse", Alias: (*Alias)(&obj)})
}
func (obj ExtensionNoResponseError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ExtensionNoResponseError: failed to parse error response"
}

type ExtensionUpdateActionsFailedError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
	LocalizedMessage   *LocalizedString `json:"localizedMessage,omitempty"`
	ExtensionExtraInfo *interface{}     `json:"extensionExtraInfo,omitempty"`
	ErrorByExtension   ErrorByExtension `json:"errorByExtension"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ExtensionUpdateActionsFailedError) MarshalJSON() ([]byte, error) {
	type Alias ExtensionUpdateActionsFailedError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ExtensionUpdateActionsFailed", Alias: (*Alias)(&obj)})
}
func (obj ExtensionUpdateActionsFailedError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ExtensionUpdateActionsFailedError: failed to parse error response"
}

type ExternalOAuthFailedError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ExternalOAuthFailedError) MarshalJSON() ([]byte, error) {
	type Alias ExternalOAuthFailedError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ExternalOAuthFailed", Alias: (*Alias)(&obj)})
}
func (obj ExternalOAuthFailedError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ExternalOAuthFailedError: failed to parse error response"
}

type FeatureRemovedError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj FeatureRemovedError) MarshalJSON() ([]byte, error) {
	type Alias FeatureRemovedError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "FeatureRemoved", Alias: (*Alias)(&obj)})
}
func (obj FeatureRemovedError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown FeatureRemovedError: failed to parse error response"
}

type GeneralError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
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

type InsufficientScopeError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
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

type InternalConstraintViolatedError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
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

type InvalidCredentialsError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
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

type InvalidCurrentPasswordError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InvalidCurrentPasswordError) MarshalJSON() ([]byte, error) {
	type Alias InvalidCurrentPasswordError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "InvalidCurrentPassword", Alias: (*Alias)(&obj)})
}
func (obj InvalidCurrentPasswordError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown InvalidCurrentPasswordError: failed to parse error response"
}

type InvalidFieldError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
	Field         string        `json:"field"`
	InvalidValue  interface{}   `json:"invalidValue"`
	AllowedValues []interface{} `json:"allowedValues"`
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

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["allowedValues"] == nil {
		delete(target, "allowedValues")
	}

	return json.Marshal(target)
}
func (obj InvalidFieldError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown InvalidFieldError: failed to parse error response"
}

type InvalidInputError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
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

type InvalidItemShippingDetailsError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
	Subject string `json:"subject"`
	ItemId  string `json:"itemId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InvalidItemShippingDetailsError) MarshalJSON() ([]byte, error) {
	type Alias InvalidItemShippingDetailsError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "InvalidItemShippingDetails", Alias: (*Alias)(&obj)})
}
func (obj InvalidItemShippingDetailsError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown InvalidItemShippingDetailsError: failed to parse error response"
}

type InvalidJsonInputError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
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

type InvalidOperationError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
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

type InvalidSubjectError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InvalidSubjectError) MarshalJSON() ([]byte, error) {
	type Alias InvalidSubjectError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "InvalidSubject", Alias: (*Alias)(&obj)})
}
func (obj InvalidSubjectError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown InvalidSubjectError: failed to parse error response"
}

type InvalidTokenError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
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

type LanguageUsedInStoresError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj LanguageUsedInStoresError) MarshalJSON() ([]byte, error) {
	type Alias LanguageUsedInStoresError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "LanguageUsedInStores", Alias: (*Alias)(&obj)})
}
func (obj LanguageUsedInStoresError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown LanguageUsedInStoresError: failed to parse error response"
}

type MatchingPriceNotFoundError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
	ProductId string  `json:"productId"`
	VariantId int     `json:"variantId"`
	Currency  *string `json:"currency,omitempty"`
	Country   *string `json:"country,omitempty"`
	// [Reference](/types#reference) to a [CustomerGroup](ctp:api:type:CustomerGroup).
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
	// [Reference](/../api/types#reference) to a [Channel](ctp:api:type:Channel).
	Channel *ChannelReference `json:"channel,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MatchingPriceNotFoundError) MarshalJSON() ([]byte, error) {
	type Alias MatchingPriceNotFoundError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "MatchingPriceNotFound", Alias: (*Alias)(&obj)})
}
func (obj MatchingPriceNotFoundError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown MatchingPriceNotFoundError: failed to parse error response"
}

type MaxResourceLimitExceededError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
	ExceededResource ReferenceTypeId `json:"exceededResource"`
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

type MissingRoleOnChannelError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
	// [ResourceIdentifier](/../api/types#resourceidentifier) to a [Channel](ctp:api:type:Channel).
	Channel *ChannelResourceIdentifier `json:"channel,omitempty"`
	// Describes the purpose and type of the Channel. A Channel can have one or more roles.
	MissingRole ChannelRoleEnum `json:"missingRole"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MissingRoleOnChannelError) MarshalJSON() ([]byte, error) {
	type Alias MissingRoleOnChannelError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "MissingRoleOnChannel", Alias: (*Alias)(&obj)})
}
func (obj MissingRoleOnChannelError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown MissingRoleOnChannelError: failed to parse error response"
}

type MissingTaxRateForCountryError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
	TaxCategoryId string  `json:"taxCategoryId"`
	Country       *string `json:"country,omitempty"`
	State         *string `json:"state,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MissingTaxRateForCountryError) MarshalJSON() ([]byte, error) {
	type Alias MissingTaxRateForCountryError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "MissingTaxRateForCountry", Alias: (*Alias)(&obj)})
}
func (obj MissingTaxRateForCountryError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown MissingTaxRateForCountryError: failed to parse error response"
}

type NoMatchingProductDiscountFoundError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj NoMatchingProductDiscountFoundError) MarshalJSON() ([]byte, error) {
	type Alias NoMatchingProductDiscountFoundError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "NoMatchingProductDiscountFound", Alias: (*Alias)(&obj)})
}
func (obj NoMatchingProductDiscountFoundError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown NoMatchingProductDiscountFoundError: failed to parse error response"
}

type NotEnabledError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj NotEnabledError) MarshalJSON() ([]byte, error) {
	type Alias NotEnabledError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "NotEnabled", Alias: (*Alias)(&obj)})
}
func (obj NotEnabledError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown NotEnabledError: failed to parse error response"
}

type ObjectNotFoundError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ObjectNotFoundError) MarshalJSON() ([]byte, error) {
	type Alias ObjectNotFoundError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ObjectNotFound", Alias: (*Alias)(&obj)})
}
func (obj ObjectNotFoundError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ObjectNotFoundError: failed to parse error response"
}

type OutOfStockError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
	LineItems []string `json:"lineItems"`
	Skus      []string `json:"skus"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OutOfStockError) MarshalJSON() ([]byte, error) {
	type Alias OutOfStockError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "OutOfStock", Alias: (*Alias)(&obj)})
}
func (obj OutOfStockError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown OutOfStockError: failed to parse error response"
}

type OverCapacityError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OverCapacityError) MarshalJSON() ([]byte, error) {
	type Alias OverCapacityError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "OverCapacity", Alias: (*Alias)(&obj)})
}
func (obj OverCapacityError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown OverCapacityError: failed to parse error response"
}

type PendingOperationError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PendingOperationError) MarshalJSON() ([]byte, error) {
	type Alias PendingOperationError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "PendingOperation", Alias: (*Alias)(&obj)})
}
func (obj PendingOperationError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown PendingOperationError: failed to parse error response"
}

type PriceChangedError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
	LineItems []string `json:"lineItems"`
	Shipping  bool     `json:"shipping"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PriceChangedError) MarshalJSON() ([]byte, error) {
	type Alias PriceChangedError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "PriceChanged", Alias: (*Alias)(&obj)})
}
func (obj PriceChangedError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown PriceChangedError: failed to parse error response"
}

type ProjectNotConfiguredForLanguagesError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
	Languages []string `json:"languages"`
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

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["languages"] == nil {
		delete(target, "languages")
	}

	return json.Marshal(target)
}
func (obj ProjectNotConfiguredForLanguagesError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ProjectNotConfiguredForLanguagesError: failed to parse error response"
}

type QueryComplexityLimitExceededError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj QueryComplexityLimitExceededError) MarshalJSON() ([]byte, error) {
	type Alias QueryComplexityLimitExceededError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "QueryComplexityLimitExceeded", Alias: (*Alias)(&obj)})
}
func (obj QueryComplexityLimitExceededError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown QueryComplexityLimitExceededError: failed to parse error response"
}

type QueryTimedOutError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj QueryTimedOutError) MarshalJSON() ([]byte, error) {
	type Alias QueryTimedOutError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "QueryTimedOut", Alias: (*Alias)(&obj)})
}
func (obj QueryTimedOutError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown QueryTimedOutError: failed to parse error response"
}

type ReferenceExistsError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
	ReferencedBy *ReferenceTypeId `json:"referencedBy,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ReferenceExistsError) MarshalJSON() ([]byte, error) {
	type Alias ReferenceExistsError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ReferenceExists", Alias: (*Alias)(&obj)})
}
func (obj ReferenceExistsError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ReferenceExistsError: failed to parse error response"
}

type ReferencedResourceNotFoundError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
	TypeId ReferenceTypeId `json:"typeId"`
	ID     *string         `json:"id,omitempty"`
	Key    *string         `json:"key,omitempty"`
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

type RequiredFieldError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
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

type ResourceNotFoundError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
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

type ResourceSizeLimitExceededError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ResourceSizeLimitExceededError) MarshalJSON() ([]byte, error) {
	type Alias ResourceSizeLimitExceededError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ResourceSizeLimitExceeded", Alias: (*Alias)(&obj)})
}
func (obj ResourceSizeLimitExceededError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ResourceSizeLimitExceededError: failed to parse error response"
}

type SearchDeactivatedError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SearchDeactivatedError) MarshalJSON() ([]byte, error) {
	type Alias SearchDeactivatedError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "SearchDeactivated", Alias: (*Alias)(&obj)})
}
func (obj SearchDeactivatedError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown SearchDeactivatedError: failed to parse error response"
}

type SearchExecutionFailureError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SearchExecutionFailureError) MarshalJSON() ([]byte, error) {
	type Alias SearchExecutionFailureError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "SearchExecutionFailure", Alias: (*Alias)(&obj)})
}
func (obj SearchExecutionFailureError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown SearchExecutionFailureError: failed to parse error response"
}

type SearchFacetPathNotFoundError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SearchFacetPathNotFoundError) MarshalJSON() ([]byte, error) {
	type Alias SearchFacetPathNotFoundError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "SearchFacetPathNotFound", Alias: (*Alias)(&obj)})
}
func (obj SearchFacetPathNotFoundError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown SearchFacetPathNotFoundError: failed to parse error response"
}

type SearchIndexingInProgressError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SearchIndexingInProgressError) MarshalJSON() ([]byte, error) {
	type Alias SearchIndexingInProgressError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "SearchIndexingInProgress", Alias: (*Alias)(&obj)})
}
func (obj SearchIndexingInProgressError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown SearchIndexingInProgressError: failed to parse error response"
}

type SemanticErrorError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SemanticErrorError) MarshalJSON() ([]byte, error) {
	type Alias SemanticErrorError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "SemanticError", Alias: (*Alias)(&obj)})
}
func (obj SemanticErrorError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown SemanticErrorError: failed to parse error response"
}

type ShippingMethodDoesNotMatchCartError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShippingMethodDoesNotMatchCartError) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodDoesNotMatchCartError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ShippingMethodDoesNotMatchCart", Alias: (*Alias)(&obj)})
}
func (obj ShippingMethodDoesNotMatchCartError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ShippingMethodDoesNotMatchCartError: failed to parse error response"
}

type SyntaxErrorError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
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

type VariantValues struct {
	Sku        *string      `json:"sku,omitempty"`
	Prices     []PriceDraft `json:"prices"`
	Attributes []Attribute  `json:"attributes"`
}

type WeakPasswordError struct {
	Message string `json:"message"`
	// interface{} `json:"//"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj WeakPasswordError) MarshalJSON() ([]byte, error) {
	type Alias WeakPasswordError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "WeakPassword", Alias: (*Alias)(&obj)})
}
func (obj WeakPasswordError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown WeakPasswordError: failed to parse error response"
}
