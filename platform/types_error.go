// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type ErrorByExtension struct {
	Id  string `json:"id"`
	Key string `json:"key,omitEmpty"`
}

type ErrorObject interface{}

func mapDiscriminatorErrorObject(input interface{}) (ErrorObject, error) {

	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["code"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'code'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}

	switch discriminator {
	case "access_denied":
		new := AccessDeniedError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "AnonymousIdAlreadyInUse":
		new := AnonymousIdAlreadyInUseError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "AttributeDefinitionAlreadyExists":
		new := AttributeDefinitionAlreadyExistsError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "AttributeDefinitionTypeConflict":
		new := AttributeDefinitionTypeConflictError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "AttributeNameDoesNotExist":
		new := AttributeNameDoesNotExistError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ConcurrentModification":
		new := ConcurrentModificationError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "DiscountCodeNonApplicable":
		new := DiscountCodeNonApplicableError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "DuplicateAttributeValue":
		new := DuplicateAttributeValueError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "DuplicateAttributeValues":
		new := DuplicateAttributeValuesError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "DuplicateEnumValues":
		new := DuplicateEnumValuesError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "DuplicateField":
		new := DuplicateFieldError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "DuplicateFieldWithConflictingResource":
		new := DuplicateFieldWithConflictingResourceError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "DuplicatePriceScope":
		new := DuplicatePriceScopeError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "DuplicateVariantValues":
		new := DuplicateVariantValuesError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "EditPreviewFailed":
		new := EditPreviewFailedError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "EnumKeyAlreadyExists":
		new := EnumKeyAlreadyExistsError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "EnumKeyDoesNotExist":
		new := EnumKeyDoesNotExistError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "EnumValueIsUsed":
		new := EnumValueIsUsedError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "EnumValuesMustMatch":
		new := EnumValuesMustMatchError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ExtensionBadResponse":
		new := ExtensionBadResponseError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ExtensionNoResponse":
		new := ExtensionNoResponseError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ExtensionUpdateActionsFailed":
		new := ExtensionUpdateActionsFailedError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ExternalOAuthFailed":
		new := ExternalOAuthFailedError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "FeatureRemoved":
		new := FeatureRemovedError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "General":
		new := GeneralError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "insufficient_scope":
		new := InsufficientScopeError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "InternalConstraintViolated":
		new := InternalConstraintViolatedError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "InvalidCredentials":
		new := InvalidCredentialsError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "InvalidCurrentPassword":
		new := InvalidCurrentPasswordError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "InvalidField":
		new := InvalidFieldError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "InvalidInput":
		new := InvalidInputError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "InvalidItemShippingDetails":
		new := InvalidItemShippingDetailsError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "InvalidJsonInput":
		new := InvalidJsonInputError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "InvalidOperation":
		new := InvalidOperationError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "InvalidSubject":
		new := InvalidSubjectError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "invalid_token":
		new := InvalidTokenError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "LanguageUsedInStores":
		new := LanguageUsedInStoresError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "MatchingPriceNotFound":
		new := MatchingPriceNotFoundError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "MaxResourceLimitExceeded":
		new := MaxResourceLimitExceededError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "MissingRoleOnChannel":
		new := MissingRoleOnChannelError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "MissingTaxRateForCountry":
		new := MissingTaxRateForCountryError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "NoMatchingProductDiscountFound":
		new := NoMatchingProductDiscountFoundError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "NotEnabled":
		new := NotEnabledError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ObjectNotFound":
		new := ObjectNotFoundError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OutOfStock":
		new := OutOfStockError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OverCapacity":
		new := OverCapacityError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "PendingOperation":
		new := PendingOperationError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "PriceChanged":
		new := PriceChangedError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProjectNotConfiguredForLanguages":
		new := ProjectNotConfiguredForLanguagesError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "QueryComplexityLimitExceeded":
		new := QueryComplexityLimitExceededError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "QueryTimedOut":
		new := QueryTimedOutError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ReferenceExists":
		new := ReferenceExistsError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ReferencedResourceNotFound":
		new := ReferencedResourceNotFoundError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "RequiredField":
		new := RequiredFieldError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ResourceNotFound":
		new := ResourceNotFoundError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ResourceSizeLimitExceeded":
		new := ResourceSizeLimitExceededError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "SearchDeactivated":
		new := SearchDeactivatedError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "SearchExecutionFailure":
		new := SearchExecutionFailureError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "SearchFacetPathNotFound":
		new := SearchFacetPathNotFoundError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "SearchIndexingInProgress":
		new := SearchIndexingInProgressError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "SemanticError":
		new := SemanticErrorError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ShippingMethodDoesNotMatchCart":
		new := ShippingMethodDoesNotMatchCartError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "SyntaxError":
		new := SyntaxErrorError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "WeakPassword":
		new := WeakPasswordError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type AccessDeniedError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj AccessDeniedError) MarshalJSON() ([]byte, error) {
	type Alias AccessDeniedError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "access_denied", Alias: (*Alias)(&obj)})
}
func (obj AccessDeniedError) Error() string {
	return obj.Message
}

type AnonymousIdAlreadyInUseError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj AnonymousIdAlreadyInUseError) MarshalJSON() ([]byte, error) {
	type Alias AnonymousIdAlreadyInUseError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "AnonymousIdAlreadyInUse", Alias: (*Alias)(&obj)})
}
func (obj AnonymousIdAlreadyInUseError) Error() string {
	return obj.Message
}

type AttributeDefinitionAlreadyExistsError struct {
	Message                    string `json:"message"`
	ConflictingProductTypeId   string `json:"conflictingProductTypeId"`
	ConflictingProductTypeName string `json:"conflictingProductTypeName"`
	ConflictingAttributeName   string `json:"conflictingAttributeName"`
}

// MarshalJSON override to set the discriminator value
func (obj AttributeDefinitionAlreadyExistsError) MarshalJSON() ([]byte, error) {
	type Alias AttributeDefinitionAlreadyExistsError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "AttributeDefinitionAlreadyExists", Alias: (*Alias)(&obj)})
}
func (obj AttributeDefinitionAlreadyExistsError) Error() string {
	return obj.Message
}

type AttributeDefinitionTypeConflictError struct {
	Message                    string `json:"message"`
	ConflictingProductTypeId   string `json:"conflictingProductTypeId"`
	ConflictingProductTypeName string `json:"conflictingProductTypeName"`
	ConflictingAttributeName   string `json:"conflictingAttributeName"`
}

// MarshalJSON override to set the discriminator value
func (obj AttributeDefinitionTypeConflictError) MarshalJSON() ([]byte, error) {
	type Alias AttributeDefinitionTypeConflictError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "AttributeDefinitionTypeConflict", Alias: (*Alias)(&obj)})
}
func (obj AttributeDefinitionTypeConflictError) Error() string {
	return obj.Message
}

type AttributeNameDoesNotExistError struct {
	Message              string `json:"message"`
	InvalidAttributeName string `json:"invalidAttributeName"`
}

// MarshalJSON override to set the discriminator value
func (obj AttributeNameDoesNotExistError) MarshalJSON() ([]byte, error) {
	type Alias AttributeNameDoesNotExistError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "AttributeNameDoesNotExist", Alias: (*Alias)(&obj)})
}
func (obj AttributeNameDoesNotExistError) Error() string {
	return obj.Message
}

type ConcurrentModificationError struct {
	Message        string `json:"message"`
	CurrentVersion int    `json:"currentVersion,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj ConcurrentModificationError) MarshalJSON() ([]byte, error) {
	type Alias ConcurrentModificationError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ConcurrentModification", Alias: (*Alias)(&obj)})
}
func (obj ConcurrentModificationError) Error() string {
	return obj.Message
}

type DiscountCodeNonApplicableError struct {
	Message           string    `json:"message"`
	DiscountCode      string    `json:"discountCode,omitEmpty"`
	Reason            string    `json:"reason,omitEmpty"`
	DicountCodeId     string    `json:"dicountCodeId,omitEmpty"`
	ValidFrom         time.Time `json:"validFrom,omitEmpty"`
	ValidUntil        time.Time `json:"validUntil,omitEmpty"`
	ValidityCheckTime time.Time `json:"validityCheckTime,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj DiscountCodeNonApplicableError) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeNonApplicableError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "DiscountCodeNonApplicable", Alias: (*Alias)(&obj)})
}
func (obj DiscountCodeNonApplicableError) Error() string {
	return obj.Message
}

type DuplicateAttributeValueError struct {
	Message   string    `json:"message"`
	Attribute Attribute `json:"attribute"`
}

// MarshalJSON override to set the discriminator value
func (obj DuplicateAttributeValueError) MarshalJSON() ([]byte, error) {
	type Alias DuplicateAttributeValueError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "DuplicateAttributeValue", Alias: (*Alias)(&obj)})
}
func (obj DuplicateAttributeValueError) Error() string {
	return obj.Message
}

type DuplicateAttributeValuesError struct {
	Message    string      `json:"message"`
	Attributes []Attribute `json:"attributes"`
}

// MarshalJSON override to set the discriminator value
func (obj DuplicateAttributeValuesError) MarshalJSON() ([]byte, error) {
	type Alias DuplicateAttributeValuesError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "DuplicateAttributeValues", Alias: (*Alias)(&obj)})
}
func (obj DuplicateAttributeValuesError) Error() string {
	return obj.Message
}

type DuplicateEnumValuesError struct {
	Message    string   `json:"message"`
	Duplicates []string `json:"duplicates"`
}

// MarshalJSON override to set the discriminator value
func (obj DuplicateEnumValuesError) MarshalJSON() ([]byte, error) {
	type Alias DuplicateEnumValuesError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "DuplicateEnumValues", Alias: (*Alias)(&obj)})
}
func (obj DuplicateEnumValuesError) Error() string {
	return obj.Message
}

type DuplicateFieldError struct {
	Message             string      `json:"message"`
	Field               string      `json:"field,omitEmpty"`
	DuplicateValue      interface{} `json:"duplicateValue,omitEmpty"`
	ConflictingResource Reference   `json:"conflictingResource,omitEmpty"`
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

// MarshalJSON override to set the discriminator value
func (obj DuplicateFieldError) MarshalJSON() ([]byte, error) {
	type Alias DuplicateFieldError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "DuplicateField", Alias: (*Alias)(&obj)})
}
func (obj DuplicateFieldError) Error() string {
	return obj.Message
}

type DuplicateFieldWithConflictingResourceError struct {
	Message             string      `json:"message"`
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

// MarshalJSON override to set the discriminator value
func (obj DuplicateFieldWithConflictingResourceError) MarshalJSON() ([]byte, error) {
	type Alias DuplicateFieldWithConflictingResourceError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "DuplicateFieldWithConflictingResource", Alias: (*Alias)(&obj)})
}
func (obj DuplicateFieldWithConflictingResourceError) Error() string {
	return obj.Message
}

type DuplicatePriceScopeError struct {
	Message           string  `json:"message"`
	ConflictingPrices []Price `json:"conflictingPrices"`
}

// MarshalJSON override to set the discriminator value
func (obj DuplicatePriceScopeError) MarshalJSON() ([]byte, error) {
	type Alias DuplicatePriceScopeError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "DuplicatePriceScope", Alias: (*Alias)(&obj)})
}
func (obj DuplicatePriceScopeError) Error() string {
	return obj.Message
}

type DuplicateVariantValuesError struct {
	Message       string        `json:"message"`
	VariantValues VariantValues `json:"variantValues"`
}

// MarshalJSON override to set the discriminator value
func (obj DuplicateVariantValuesError) MarshalJSON() ([]byte, error) {
	type Alias DuplicateVariantValuesError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "DuplicateVariantValues", Alias: (*Alias)(&obj)})
}
func (obj DuplicateVariantValuesError) Error() string {
	return obj.Message
}

type EditPreviewFailedError struct {
	Message string                  `json:"message"`
	Result  OrderEditPreviewFailure `json:"result"`
}

// MarshalJSON override to set the discriminator value
func (obj EditPreviewFailedError) MarshalJSON() ([]byte, error) {
	type Alias EditPreviewFailedError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "EditPreviewFailed", Alias: (*Alias)(&obj)})
}
func (obj EditPreviewFailedError) Error() string {
	return obj.Message
}

type EnumKeyAlreadyExistsError struct {
	Message                  string `json:"message"`
	ConflictingEnumKey       string `json:"conflictingEnumKey"`
	ConflictingAttributeName string `json:"conflictingAttributeName"`
}

// MarshalJSON override to set the discriminator value
func (obj EnumKeyAlreadyExistsError) MarshalJSON() ([]byte, error) {
	type Alias EnumKeyAlreadyExistsError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "EnumKeyAlreadyExists", Alias: (*Alias)(&obj)})
}
func (obj EnumKeyAlreadyExistsError) Error() string {
	return obj.Message
}

type EnumKeyDoesNotExistError struct {
	Message                  string `json:"message"`
	ConflictingEnumKey       string `json:"conflictingEnumKey"`
	ConflictingAttributeName string `json:"conflictingAttributeName"`
}

// MarshalJSON override to set the discriminator value
func (obj EnumKeyDoesNotExistError) MarshalJSON() ([]byte, error) {
	type Alias EnumKeyDoesNotExistError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "EnumKeyDoesNotExist", Alias: (*Alias)(&obj)})
}
func (obj EnumKeyDoesNotExistError) Error() string {
	return obj.Message
}

type EnumValueIsUsedError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj EnumValueIsUsedError) MarshalJSON() ([]byte, error) {
	type Alias EnumValueIsUsedError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "EnumValueIsUsed", Alias: (*Alias)(&obj)})
}
func (obj EnumValueIsUsedError) Error() string {
	return obj.Message
}

type EnumValuesMustMatchError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj EnumValuesMustMatchError) MarshalJSON() ([]byte, error) {
	type Alias EnumValuesMustMatchError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "EnumValuesMustMatch", Alias: (*Alias)(&obj)})
}
func (obj EnumValuesMustMatchError) Error() string {
	return obj.Message
}

type ErrorResponse struct {
	StatusCode       int           `json:"statusCode"`
	Message          string        `json:"message"`
	ErrorMessage     string        `json:"error,omitEmpty"`
	ErrorDescription string        `json:"error_description,omitEmpty"`
	Errors           []ErrorObject `json:"errors,omitEmpty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ErrorResponse) UnmarshalJSON(data []byte) error {
	type Alias ErrorResponse
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	return nil
}

func (obj ErrorResponse) Error() string {
	return obj.Message
}

type ExtensionBadResponseError struct {
	Message            string           `json:"message"`
	LocalizedMessage   *LocalizedString `json:"localizedMessage,omitEmpty"`
	ExtensionExtraInfo *interface{}     `json:"extensionExtraInfo,omitEmpty"`
	ErrorByExtension   ErrorByExtension `json:"errorByExtension"`
}

// MarshalJSON override to set the discriminator value
func (obj ExtensionBadResponseError) MarshalJSON() ([]byte, error) {
	type Alias ExtensionBadResponseError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ExtensionBadResponse", Alias: (*Alias)(&obj)})
}
func (obj ExtensionBadResponseError) Error() string {
	return obj.Message
}

type ExtensionNoResponseError struct {
	Message      string `json:"message"`
	ExtensionId  string `json:"extensionId"`
	ExtensionKey string `json:"extensionKey,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj ExtensionNoResponseError) MarshalJSON() ([]byte, error) {
	type Alias ExtensionNoResponseError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ExtensionNoResponse", Alias: (*Alias)(&obj)})
}
func (obj ExtensionNoResponseError) Error() string {
	return obj.Message
}

type ExtensionUpdateActionsFailedError struct {
	Message            string           `json:"message"`
	LocalizedMessage   *LocalizedString `json:"localizedMessage,omitEmpty"`
	ExtensionExtraInfo *interface{}     `json:"extensionExtraInfo,omitEmpty"`
	ErrorByExtension   ErrorByExtension `json:"errorByExtension"`
}

// MarshalJSON override to set the discriminator value
func (obj ExtensionUpdateActionsFailedError) MarshalJSON() ([]byte, error) {
	type Alias ExtensionUpdateActionsFailedError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ExtensionUpdateActionsFailed", Alias: (*Alias)(&obj)})
}
func (obj ExtensionUpdateActionsFailedError) Error() string {
	return obj.Message
}

type ExternalOAuthFailedError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj ExternalOAuthFailedError) MarshalJSON() ([]byte, error) {
	type Alias ExternalOAuthFailedError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ExternalOAuthFailed", Alias: (*Alias)(&obj)})
}
func (obj ExternalOAuthFailedError) Error() string {
	return obj.Message
}

type FeatureRemovedError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj FeatureRemovedError) MarshalJSON() ([]byte, error) {
	type Alias FeatureRemovedError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "FeatureRemoved", Alias: (*Alias)(&obj)})
}
func (obj FeatureRemovedError) Error() string {
	return obj.Message
}

type GeneralError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj GeneralError) MarshalJSON() ([]byte, error) {
	type Alias GeneralError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "General", Alias: (*Alias)(&obj)})
}
func (obj GeneralError) Error() string {
	return obj.Message
}

type InsufficientScopeError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj InsufficientScopeError) MarshalJSON() ([]byte, error) {
	type Alias InsufficientScopeError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "insufficient_scope", Alias: (*Alias)(&obj)})
}
func (obj InsufficientScopeError) Error() string {
	return obj.Message
}

type InternalConstraintViolatedError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj InternalConstraintViolatedError) MarshalJSON() ([]byte, error) {
	type Alias InternalConstraintViolatedError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "InternalConstraintViolated", Alias: (*Alias)(&obj)})
}
func (obj InternalConstraintViolatedError) Error() string {
	return obj.Message
}

type InvalidCredentialsError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj InvalidCredentialsError) MarshalJSON() ([]byte, error) {
	type Alias InvalidCredentialsError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "InvalidCredentials", Alias: (*Alias)(&obj)})
}
func (obj InvalidCredentialsError) Error() string {
	return obj.Message
}

type InvalidCurrentPasswordError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj InvalidCurrentPasswordError) MarshalJSON() ([]byte, error) {
	type Alias InvalidCurrentPasswordError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "InvalidCurrentPassword", Alias: (*Alias)(&obj)})
}
func (obj InvalidCurrentPasswordError) Error() string {
	return obj.Message
}

type InvalidFieldError struct {
	Message       string        `json:"message"`
	Field         string        `json:"field"`
	InvalidValue  interface{}   `json:"invalidValue"`
	AllowedValues []interface{} `json:"allowedValues,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj InvalidFieldError) MarshalJSON() ([]byte, error) {
	type Alias InvalidFieldError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "InvalidField", Alias: (*Alias)(&obj)})
}
func (obj InvalidFieldError) Error() string {
	return obj.Message
}

type InvalidInputError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj InvalidInputError) MarshalJSON() ([]byte, error) {
	type Alias InvalidInputError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "InvalidInput", Alias: (*Alias)(&obj)})
}
func (obj InvalidInputError) Error() string {
	return obj.Message
}

type InvalidItemShippingDetailsError struct {
	Message string `json:"message"`
	Subject string `json:"subject"`
	ItemId  string `json:"itemId"`
}

// MarshalJSON override to set the discriminator value
func (obj InvalidItemShippingDetailsError) MarshalJSON() ([]byte, error) {
	type Alias InvalidItemShippingDetailsError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "InvalidItemShippingDetails", Alias: (*Alias)(&obj)})
}
func (obj InvalidItemShippingDetailsError) Error() string {
	return obj.Message
}

type InvalidJsonInputError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj InvalidJsonInputError) MarshalJSON() ([]byte, error) {
	type Alias InvalidJsonInputError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "InvalidJsonInput", Alias: (*Alias)(&obj)})
}
func (obj InvalidJsonInputError) Error() string {
	return obj.Message
}

type InvalidOperationError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj InvalidOperationError) MarshalJSON() ([]byte, error) {
	type Alias InvalidOperationError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "InvalidOperation", Alias: (*Alias)(&obj)})
}
func (obj InvalidOperationError) Error() string {
	return obj.Message
}

type InvalidSubjectError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj InvalidSubjectError) MarshalJSON() ([]byte, error) {
	type Alias InvalidSubjectError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "InvalidSubject", Alias: (*Alias)(&obj)})
}
func (obj InvalidSubjectError) Error() string {
	return obj.Message
}

type InvalidTokenError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj InvalidTokenError) MarshalJSON() ([]byte, error) {
	type Alias InvalidTokenError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "invalid_token", Alias: (*Alias)(&obj)})
}
func (obj InvalidTokenError) Error() string {
	return obj.Message
}

type LanguageUsedInStoresError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj LanguageUsedInStoresError) MarshalJSON() ([]byte, error) {
	type Alias LanguageUsedInStoresError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "LanguageUsedInStores", Alias: (*Alias)(&obj)})
}
func (obj LanguageUsedInStoresError) Error() string {
	return obj.Message
}

type MatchingPriceNotFoundError struct {
	Message       string                 `json:"message"`
	ProductId     string                 `json:"productId"`
	VariantId     int                    `json:"variantId"`
	Currency      string                 `json:"currency,omitEmpty"`
	Country       string                 `json:"country,omitEmpty"`
	CustomerGroup CustomerGroupReference `json:"customerGroup,omitEmpty"`
	Channel       ChannelReference       `json:"channel,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj MatchingPriceNotFoundError) MarshalJSON() ([]byte, error) {
	type Alias MatchingPriceNotFoundError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "MatchingPriceNotFound", Alias: (*Alias)(&obj)})
}
func (obj MatchingPriceNotFoundError) Error() string {
	return obj.Message
}

type MaxResourceLimitExceededError struct {
	Message          string          `json:"message"`
	ExceededResource ReferenceTypeId `json:"exceededResource"`
}

// MarshalJSON override to set the discriminator value
func (obj MaxResourceLimitExceededError) MarshalJSON() ([]byte, error) {
	type Alias MaxResourceLimitExceededError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "MaxResourceLimitExceeded", Alias: (*Alias)(&obj)})
}
func (obj MaxResourceLimitExceededError) Error() string {
	return obj.Message
}

type MissingRoleOnChannelError struct {
	Message     string                    `json:"message"`
	Channel     ChannelResourceIdentifier `json:"channel,omitEmpty"`
	MissingRole ChannelRoleEnum           `json:"missingRole"`
}

// MarshalJSON override to set the discriminator value
func (obj MissingRoleOnChannelError) MarshalJSON() ([]byte, error) {
	type Alias MissingRoleOnChannelError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "MissingRoleOnChannel", Alias: (*Alias)(&obj)})
}
func (obj MissingRoleOnChannelError) Error() string {
	return obj.Message
}

type MissingTaxRateForCountryError struct {
	Message       string `json:"message"`
	TaxCategoryId string `json:"taxCategoryId"`
	Country       string `json:"country,omitEmpty"`
	State         string `json:"state,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj MissingTaxRateForCountryError) MarshalJSON() ([]byte, error) {
	type Alias MissingTaxRateForCountryError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "MissingTaxRateForCountry", Alias: (*Alias)(&obj)})
}
func (obj MissingTaxRateForCountryError) Error() string {
	return obj.Message
}

type NoMatchingProductDiscountFoundError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj NoMatchingProductDiscountFoundError) MarshalJSON() ([]byte, error) {
	type Alias NoMatchingProductDiscountFoundError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "NoMatchingProductDiscountFound", Alias: (*Alias)(&obj)})
}
func (obj NoMatchingProductDiscountFoundError) Error() string {
	return obj.Message
}

type NotEnabledError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj NotEnabledError) MarshalJSON() ([]byte, error) {
	type Alias NotEnabledError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "NotEnabled", Alias: (*Alias)(&obj)})
}
func (obj NotEnabledError) Error() string {
	return obj.Message
}

type ObjectNotFoundError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj ObjectNotFoundError) MarshalJSON() ([]byte, error) {
	type Alias ObjectNotFoundError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ObjectNotFound", Alias: (*Alias)(&obj)})
}
func (obj ObjectNotFoundError) Error() string {
	return obj.Message
}

type OutOfStockError struct {
	Message   string   `json:"message"`
	LineItems []string `json:"lineItems"`
	Skus      []string `json:"skus"`
}

// MarshalJSON override to set the discriminator value
func (obj OutOfStockError) MarshalJSON() ([]byte, error) {
	type Alias OutOfStockError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "OutOfStock", Alias: (*Alias)(&obj)})
}
func (obj OutOfStockError) Error() string {
	return obj.Message
}

type OverCapacityError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj OverCapacityError) MarshalJSON() ([]byte, error) {
	type Alias OverCapacityError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "OverCapacity", Alias: (*Alias)(&obj)})
}
func (obj OverCapacityError) Error() string {
	return obj.Message
}

type PendingOperationError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj PendingOperationError) MarshalJSON() ([]byte, error) {
	type Alias PendingOperationError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "PendingOperation", Alias: (*Alias)(&obj)})
}
func (obj PendingOperationError) Error() string {
	return obj.Message
}

type PriceChangedError struct {
	Message   string   `json:"message"`
	LineItems []string `json:"lineItems"`
	Shipping  bool     `json:"shipping"`
}

// MarshalJSON override to set the discriminator value
func (obj PriceChangedError) MarshalJSON() ([]byte, error) {
	type Alias PriceChangedError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "PriceChanged", Alias: (*Alias)(&obj)})
}
func (obj PriceChangedError) Error() string {
	return obj.Message
}

type ProjectNotConfiguredForLanguagesError struct {
	Message   string   `json:"message"`
	Languages []string `json:"languages,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProjectNotConfiguredForLanguagesError) MarshalJSON() ([]byte, error) {
	type Alias ProjectNotConfiguredForLanguagesError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ProjectNotConfiguredForLanguages", Alias: (*Alias)(&obj)})
}
func (obj ProjectNotConfiguredForLanguagesError) Error() string {
	return obj.Message
}

type QueryComplexityLimitExceededError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj QueryComplexityLimitExceededError) MarshalJSON() ([]byte, error) {
	type Alias QueryComplexityLimitExceededError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "QueryComplexityLimitExceeded", Alias: (*Alias)(&obj)})
}
func (obj QueryComplexityLimitExceededError) Error() string {
	return obj.Message
}

type QueryTimedOutError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj QueryTimedOutError) MarshalJSON() ([]byte, error) {
	type Alias QueryTimedOutError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "QueryTimedOut", Alias: (*Alias)(&obj)})
}
func (obj QueryTimedOutError) Error() string {
	return obj.Message
}

type ReferenceExistsError struct {
	Message      string          `json:"message"`
	ReferencedBy ReferenceTypeId `json:"referencedBy,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj ReferenceExistsError) MarshalJSON() ([]byte, error) {
	type Alias ReferenceExistsError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ReferenceExists", Alias: (*Alias)(&obj)})
}
func (obj ReferenceExistsError) Error() string {
	return obj.Message
}

type ReferencedResourceNotFoundError struct {
	Message string          `json:"message"`
	TypeId  ReferenceTypeId `json:"typeId"`
	Id      string          `json:"id,omitEmpty"`
	Key     string          `json:"key,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj ReferencedResourceNotFoundError) MarshalJSON() ([]byte, error) {
	type Alias ReferencedResourceNotFoundError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ReferencedResourceNotFound", Alias: (*Alias)(&obj)})
}
func (obj ReferencedResourceNotFoundError) Error() string {
	return obj.Message
}

type RequiredFieldError struct {
	Message string `json:"message"`
	Field   string `json:"field"`
}

// MarshalJSON override to set the discriminator value
func (obj RequiredFieldError) MarshalJSON() ([]byte, error) {
	type Alias RequiredFieldError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "RequiredField", Alias: (*Alias)(&obj)})
}
func (obj RequiredFieldError) Error() string {
	return obj.Message
}

type ResourceNotFoundError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj ResourceNotFoundError) MarshalJSON() ([]byte, error) {
	type Alias ResourceNotFoundError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ResourceNotFound", Alias: (*Alias)(&obj)})
}
func (obj ResourceNotFoundError) Error() string {
	return obj.Message
}

type ResourceSizeLimitExceededError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj ResourceSizeLimitExceededError) MarshalJSON() ([]byte, error) {
	type Alias ResourceSizeLimitExceededError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ResourceSizeLimitExceeded", Alias: (*Alias)(&obj)})
}
func (obj ResourceSizeLimitExceededError) Error() string {
	return obj.Message
}

type SearchDeactivatedError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj SearchDeactivatedError) MarshalJSON() ([]byte, error) {
	type Alias SearchDeactivatedError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "SearchDeactivated", Alias: (*Alias)(&obj)})
}
func (obj SearchDeactivatedError) Error() string {
	return obj.Message
}

type SearchExecutionFailureError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj SearchExecutionFailureError) MarshalJSON() ([]byte, error) {
	type Alias SearchExecutionFailureError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "SearchExecutionFailure", Alias: (*Alias)(&obj)})
}
func (obj SearchExecutionFailureError) Error() string {
	return obj.Message
}

type SearchFacetPathNotFoundError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj SearchFacetPathNotFoundError) MarshalJSON() ([]byte, error) {
	type Alias SearchFacetPathNotFoundError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "SearchFacetPathNotFound", Alias: (*Alias)(&obj)})
}
func (obj SearchFacetPathNotFoundError) Error() string {
	return obj.Message
}

type SearchIndexingInProgressError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj SearchIndexingInProgressError) MarshalJSON() ([]byte, error) {
	type Alias SearchIndexingInProgressError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "SearchIndexingInProgress", Alias: (*Alias)(&obj)})
}
func (obj SearchIndexingInProgressError) Error() string {
	return obj.Message
}

type SemanticErrorError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj SemanticErrorError) MarshalJSON() ([]byte, error) {
	type Alias SemanticErrorError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "SemanticError", Alias: (*Alias)(&obj)})
}
func (obj SemanticErrorError) Error() string {
	return obj.Message
}

type ShippingMethodDoesNotMatchCartError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodDoesNotMatchCartError) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodDoesNotMatchCartError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ShippingMethodDoesNotMatchCart", Alias: (*Alias)(&obj)})
}
func (obj ShippingMethodDoesNotMatchCartError) Error() string {
	return obj.Message
}

type SyntaxErrorError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj SyntaxErrorError) MarshalJSON() ([]byte, error) {
	type Alias SyntaxErrorError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "SyntaxError", Alias: (*Alias)(&obj)})
}
func (obj SyntaxErrorError) Error() string {
	return obj.Message
}

type VariantValues struct {
	Sku        string       `json:"sku,omitEmpty"`
	Prices     []PriceDraft `json:"prices"`
	Attributes []Attribute  `json:"attributes"`
}

type WeakPasswordError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj WeakPasswordError) MarshalJSON() ([]byte, error) {
	type Alias WeakPasswordError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "WeakPassword", Alias: (*Alias)(&obj)})
}
func (obj WeakPasswordError) Error() string {
	return obj.Message
}
