// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"errors"
	"time"
)

// ErrorObject uses code as discriminator attribute
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
		if new.ConflictingResource != nil {
			new.ConflictingResource, err = mapDiscriminatorReference(new.ConflictingResource)
			if err != nil {
				return nil, err
			}
		}
		return new, nil
	case "DuplicateFieldWithConflictingResource":
		new := DuplicateFieldWithConflictingResourceError{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		if new.ConflictingResource != nil {
			new.ConflictingResource, err = mapDiscriminatorReference(new.ConflictingResource)
			if err != nil {
				return nil, err
			}
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
		new := InvalidJSONInputError{}
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

// AnonymousIdAlreadyInUseError implements the interface ErrorObject
type AnonymousIdAlreadyInUseError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj AnonymousIdAlreadyInUseError) MarshalJSON() ([]byte, error) {
	type Alias AnonymousIdAlreadyInUseError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "AnonymousIdAlreadyInUse", Alias: (*Alias)(&obj)})
}

func (obj AnonymousIdAlreadyInUseError) Error() string {
	return obj.Message
}

// AttributeDefinitionAlreadyExistsError implements the interface ErrorObject
type AttributeDefinitionAlreadyExistsError struct {
	Message                    string `json:"message"`
	ConflictingProductTypeName string `json:"conflictingProductTypeName"`
	ConflictingProductTypeID   string `json:"conflictingProductTypeId"`
	ConflictingAttributeName   string `json:"conflictingAttributeName"`
}

// MarshalJSON override to set the discriminator value
func (obj AttributeDefinitionAlreadyExistsError) MarshalJSON() ([]byte, error) {
	type Alias AttributeDefinitionAlreadyExistsError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "AttributeDefinitionAlreadyExists", Alias: (*Alias)(&obj)})
}

func (obj AttributeDefinitionAlreadyExistsError) Error() string {
	return obj.Message
}

// AttributeDefinitionTypeConflictError implements the interface ErrorObject
type AttributeDefinitionTypeConflictError struct {
	Message                    string `json:"message"`
	ConflictingProductTypeName string `json:"conflictingProductTypeName"`
	ConflictingProductTypeID   string `json:"conflictingProductTypeId"`
	ConflictingAttributeName   string `json:"conflictingAttributeName"`
}

// MarshalJSON override to set the discriminator value
func (obj AttributeDefinitionTypeConflictError) MarshalJSON() ([]byte, error) {
	type Alias AttributeDefinitionTypeConflictError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "AttributeDefinitionTypeConflict", Alias: (*Alias)(&obj)})
}

func (obj AttributeDefinitionTypeConflictError) Error() string {
	return obj.Message
}

// AttributeNameDoesNotExistError implements the interface ErrorObject
type AttributeNameDoesNotExistError struct {
	Message              string `json:"message"`
	InvalidAttributeName string `json:"invalidAttributeName"`
}

// MarshalJSON override to set the discriminator value
func (obj AttributeNameDoesNotExistError) MarshalJSON() ([]byte, error) {
	type Alias AttributeNameDoesNotExistError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "AttributeNameDoesNotExist", Alias: (*Alias)(&obj)})
}

func (obj AttributeNameDoesNotExistError) Error() string {
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
	Message           string     `json:"message"`
	ValidityCheckTime *time.Time `json:"validityCheckTime,omitempty"`
	ValidUntil        *time.Time `json:"validUntil,omitempty"`
	ValidFrom         *time.Time `json:"validFrom,omitempty"`
	Reason            string     `json:"reason,omitempty"`
	DiscountCode      string     `json:"discountCode,omitempty"`
	DicountCodeID     string     `json:"dicountCodeId,omitempty"`
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

// DuplicateEnumValuesError implements the interface ErrorObject
type DuplicateEnumValuesError struct {
	Message    string   `json:"message"`
	Duplicates []string `json:"duplicates"`
}

// MarshalJSON override to set the discriminator value
func (obj DuplicateEnumValuesError) MarshalJSON() ([]byte, error) {
	type Alias DuplicateEnumValuesError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "DuplicateEnumValues", Alias: (*Alias)(&obj)})
}

func (obj DuplicateEnumValuesError) Error() string {
	return obj.Message
}

// DuplicateFieldError implements the interface ErrorObject
type DuplicateFieldError struct {
	Message             string      `json:"message"`
	Field               string      `json:"field,omitempty"`
	DuplicateValue      interface{} `json:"duplicateValue,omitempty"`
	ConflictingResource Reference   `json:"conflictingResource,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj DuplicateFieldError) MarshalJSON() ([]byte, error) {
	type Alias DuplicateFieldError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "DuplicateField", Alias: (*Alias)(&obj)})
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

func (obj DuplicateFieldError) Error() string {
	return obj.Message
}

// DuplicateFieldWithConflictingResourceError implements the interface ErrorObject
type DuplicateFieldWithConflictingResourceError struct {
	Message             string      `json:"message"`
	Field               string      `json:"field"`
	DuplicateValue      interface{} `json:"duplicateValue"`
	ConflictingResource Reference   `json:"conflictingResource"`
}

// MarshalJSON override to set the discriminator value
func (obj DuplicateFieldWithConflictingResourceError) MarshalJSON() ([]byte, error) {
	type Alias DuplicateFieldWithConflictingResourceError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "DuplicateFieldWithConflictingResource", Alias: (*Alias)(&obj)})
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

func (obj DuplicateFieldWithConflictingResourceError) Error() string {
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

// EditPreviewFailedError implements the interface ErrorObject
type EditPreviewFailedError struct {
	Message string                   `json:"message"`
	Result  *OrderEditPreviewFailure `json:"result"`
}

// MarshalJSON override to set the discriminator value
func (obj EditPreviewFailedError) MarshalJSON() ([]byte, error) {
	type Alias EditPreviewFailedError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "EditPreviewFailed", Alias: (*Alias)(&obj)})
}

func (obj EditPreviewFailedError) Error() string {
	return obj.Message
}

// EnumKeyAlreadyExistsError implements the interface ErrorObject
type EnumKeyAlreadyExistsError struct {
	Message                  string `json:"message"`
	ConflictingEnumKey       string `json:"conflictingEnumKey"`
	ConflictingAttributeName string `json:"conflictingAttributeName"`
}

// MarshalJSON override to set the discriminator value
func (obj EnumKeyAlreadyExistsError) MarshalJSON() ([]byte, error) {
	type Alias EnumKeyAlreadyExistsError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "EnumKeyAlreadyExists", Alias: (*Alias)(&obj)})
}

func (obj EnumKeyAlreadyExistsError) Error() string {
	return obj.Message
}

// EnumKeyDoesNotExistError implements the interface ErrorObject
type EnumKeyDoesNotExistError struct {
	Message                  string `json:"message"`
	ConflictingEnumKey       string `json:"conflictingEnumKey"`
	ConflictingAttributeName string `json:"conflictingAttributeName"`
}

// MarshalJSON override to set the discriminator value
func (obj EnumKeyDoesNotExistError) MarshalJSON() ([]byte, error) {
	type Alias EnumKeyDoesNotExistError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "EnumKeyDoesNotExist", Alias: (*Alias)(&obj)})
}

func (obj EnumKeyDoesNotExistError) Error() string {
	return obj.Message
}

// EnumValueIsUsedError implements the interface ErrorObject
type EnumValueIsUsedError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj EnumValueIsUsedError) MarshalJSON() ([]byte, error) {
	type Alias EnumValueIsUsedError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "EnumValueIsUsed", Alias: (*Alias)(&obj)})
}

func (obj EnumValueIsUsedError) Error() string {
	return obj.Message
}

// EnumValuesMustMatchError implements the interface ErrorObject
type EnumValuesMustMatchError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj EnumValuesMustMatchError) MarshalJSON() ([]byte, error) {
	type Alias EnumValuesMustMatchError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "EnumValuesMustMatch", Alias: (*Alias)(&obj)})
}

func (obj EnumValuesMustMatchError) Error() string {
	return obj.Message
}

// ErrorByExtension is a standalone struct
type ErrorByExtension struct {
	Key string `json:"key,omitempty"`
	ID  string `json:"id"`
}

// ErrorResponse is a standalone struct
type ErrorResponse struct {
	StatusCode       int           `json:"statusCode"`
	Message          string        `json:"message"`
	Errors           []ErrorObject `json:"errors,omitempty"`
	ErrorDescription string        `json:"error_description,omitempty"`
	ErrorMessage     string        `json:"error,omitempty"`
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

func (obj ErrorResponse) Error() string {
	return obj.Message
}

// ExtensionBadResponseError implements the interface ErrorObject
type ExtensionBadResponseError struct {
	Message            string            `json:"message"`
	LocalizedMessage   *LocalizedString  `json:"localizedMessage,omitempty"`
	ExtensionExtraInfo interface{}       `json:"extensionExtraInfo,omitempty"`
	ErrorByExtension   *ErrorByExtension `json:"errorByExtension"`
}

// MarshalJSON override to set the discriminator value
func (obj ExtensionBadResponseError) MarshalJSON() ([]byte, error) {
	type Alias ExtensionBadResponseError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "ExtensionBadResponse", Alias: (*Alias)(&obj)})
}

func (obj ExtensionBadResponseError) Error() string {
	return obj.Message
}

// ExtensionNoResponseError implements the interface ErrorObject
type ExtensionNoResponseError struct {
	Message      string `json:"message"`
	ExtensionKey string `json:"extensionKey,omitempty"`
	ExtensionID  string `json:"extensionId"`
}

// MarshalJSON override to set the discriminator value
func (obj ExtensionNoResponseError) MarshalJSON() ([]byte, error) {
	type Alias ExtensionNoResponseError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "ExtensionNoResponse", Alias: (*Alias)(&obj)})
}

func (obj ExtensionNoResponseError) Error() string {
	return obj.Message
}

// ExtensionUpdateActionsFailedError implements the interface ErrorObject
type ExtensionUpdateActionsFailedError struct {
	Message            string            `json:"message"`
	LocalizedMessage   *LocalizedString  `json:"localizedMessage,omitempty"`
	ExtensionExtraInfo interface{}       `json:"extensionExtraInfo,omitempty"`
	ErrorByExtension   *ErrorByExtension `json:"errorByExtension"`
}

// MarshalJSON override to set the discriminator value
func (obj ExtensionUpdateActionsFailedError) MarshalJSON() ([]byte, error) {
	type Alias ExtensionUpdateActionsFailedError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "ExtensionUpdateActionsFailed", Alias: (*Alias)(&obj)})
}

func (obj ExtensionUpdateActionsFailedError) Error() string {
	return obj.Message
}

// ExternalOAuthFailedError implements the interface ErrorObject
type ExternalOAuthFailedError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj ExternalOAuthFailedError) MarshalJSON() ([]byte, error) {
	type Alias ExternalOAuthFailedError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "ExternalOAuthFailed", Alias: (*Alias)(&obj)})
}

func (obj ExternalOAuthFailedError) Error() string {
	return obj.Message
}

// FeatureRemovedError implements the interface ErrorObject
type FeatureRemovedError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj FeatureRemovedError) MarshalJSON() ([]byte, error) {
	type Alias FeatureRemovedError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "FeatureRemoved", Alias: (*Alias)(&obj)})
}

func (obj FeatureRemovedError) Error() string {
	return obj.Message
}

// GeneralError implements the interface ErrorObject
type GeneralError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj GeneralError) MarshalJSON() ([]byte, error) {
	type Alias GeneralError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "General", Alias: (*Alias)(&obj)})
}

func (obj GeneralError) Error() string {
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

// InternalConstraintViolatedError implements the interface ErrorObject
type InternalConstraintViolatedError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj InternalConstraintViolatedError) MarshalJSON() ([]byte, error) {
	type Alias InternalConstraintViolatedError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "InternalConstraintViolated", Alias: (*Alias)(&obj)})
}

func (obj InternalConstraintViolatedError) Error() string {
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

// LanguageUsedInStoresError implements the interface ErrorObject
type LanguageUsedInStoresError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj LanguageUsedInStoresError) MarshalJSON() ([]byte, error) {
	type Alias LanguageUsedInStoresError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "LanguageUsedInStores", Alias: (*Alias)(&obj)})
}

func (obj LanguageUsedInStoresError) Error() string {
	return obj.Message
}

// MatchingPriceNotFoundError implements the interface ErrorObject
type MatchingPriceNotFoundError struct {
	Message       string                  `json:"message"`
	VariantID     int                     `json:"variantId"`
	ProductID     string                  `json:"productId"`
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
	Currency      string                  `json:"currency,omitempty"`
	Country       string                  `json:"country,omitempty"`
	Channel       *ChannelReference       `json:"channel,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MatchingPriceNotFoundError) MarshalJSON() ([]byte, error) {
	type Alias MatchingPriceNotFoundError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "MatchingPriceNotFound", Alias: (*Alias)(&obj)})
}

func (obj MatchingPriceNotFoundError) Error() string {
	return obj.Message
}

// MaxResourceLimitExceededError implements the interface ErrorObject
type MaxResourceLimitExceededError struct {
	Message          string          `json:"message"`
	ExceededResource ReferenceTypeID `json:"exceededResource"`
}

// MarshalJSON override to set the discriminator value
func (obj MaxResourceLimitExceededError) MarshalJSON() ([]byte, error) {
	type Alias MaxResourceLimitExceededError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "MaxResourceLimitExceeded", Alias: (*Alias)(&obj)})
}

func (obj MaxResourceLimitExceededError) Error() string {
	return obj.Message
}

// MissingRoleOnChannelError implements the interface ErrorObject
type MissingRoleOnChannelError struct {
	Message     string                     `json:"message"`
	MissingRole ChannelRoleEnum            `json:"missingRole"`
	Channel     *ChannelResourceIdentifier `json:"channel,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MissingRoleOnChannelError) MarshalJSON() ([]byte, error) {
	type Alias MissingRoleOnChannelError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "MissingRoleOnChannel", Alias: (*Alias)(&obj)})
}

func (obj MissingRoleOnChannelError) Error() string {
	return obj.Message
}

// MissingTaxRateForCountryError implements the interface ErrorObject
type MissingTaxRateForCountryError struct {
	Message       string `json:"message"`
	TaxCategoryID string `json:"taxCategoryId"`
	State         string `json:"state,omitempty"`
	Country       string `json:"country,omitempty"`
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

// NotEnabledError implements the interface ErrorObject
type NotEnabledError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj NotEnabledError) MarshalJSON() ([]byte, error) {
	type Alias NotEnabledError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "NotEnabled", Alias: (*Alias)(&obj)})
}

func (obj NotEnabledError) Error() string {
	return obj.Message
}

// ObjectNotFoundError implements the interface ErrorObject
type ObjectNotFoundError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj ObjectNotFoundError) MarshalJSON() ([]byte, error) {
	type Alias ObjectNotFoundError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "ObjectNotFound", Alias: (*Alias)(&obj)})
}

func (obj ObjectNotFoundError) Error() string {
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

// OverCapacityError implements the interface ErrorObject
type OverCapacityError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj OverCapacityError) MarshalJSON() ([]byte, error) {
	type Alias OverCapacityError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "OverCapacity", Alias: (*Alias)(&obj)})
}

func (obj OverCapacityError) Error() string {
	return obj.Message
}

// PendingOperationError implements the interface ErrorObject
type PendingOperationError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj PendingOperationError) MarshalJSON() ([]byte, error) {
	type Alias PendingOperationError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "PendingOperation", Alias: (*Alias)(&obj)})
}

func (obj PendingOperationError) Error() string {
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

// ProjectNotConfiguredForLanguagesError implements the interface ErrorObject
type ProjectNotConfiguredForLanguagesError struct {
	Message   string   `json:"message"`
	Languages []string `json:"languages,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProjectNotConfiguredForLanguagesError) MarshalJSON() ([]byte, error) {
	type Alias ProjectNotConfiguredForLanguagesError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "ProjectNotConfiguredForLanguages", Alias: (*Alias)(&obj)})
}

func (obj ProjectNotConfiguredForLanguagesError) Error() string {
	return obj.Message
}

// QueryComplexityLimitExceededError implements the interface ErrorObject
type QueryComplexityLimitExceededError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj QueryComplexityLimitExceededError) MarshalJSON() ([]byte, error) {
	type Alias QueryComplexityLimitExceededError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "QueryComplexityLimitExceeded", Alias: (*Alias)(&obj)})
}

func (obj QueryComplexityLimitExceededError) Error() string {
	return obj.Message
}

// QueryTimedOutError implements the interface ErrorObject
type QueryTimedOutError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj QueryTimedOutError) MarshalJSON() ([]byte, error) {
	type Alias QueryTimedOutError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "QueryTimedOut", Alias: (*Alias)(&obj)})
}

func (obj QueryTimedOutError) Error() string {
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

// ReferencedResourceNotFoundError implements the interface ErrorObject
type ReferencedResourceNotFoundError struct {
	Message string          `json:"message"`
	TypeID  ReferenceTypeID `json:"typeId"`
	Key     string          `json:"key,omitempty"`
	ID      string          `json:"id,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ReferencedResourceNotFoundError) MarshalJSON() ([]byte, error) {
	type Alias ReferencedResourceNotFoundError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "ReferencedResourceNotFound", Alias: (*Alias)(&obj)})
}

func (obj ReferencedResourceNotFoundError) Error() string {
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

// ResourceSizeLimitExceededError implements the interface ErrorObject
type ResourceSizeLimitExceededError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj ResourceSizeLimitExceededError) MarshalJSON() ([]byte, error) {
	type Alias ResourceSizeLimitExceededError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "ResourceSizeLimitExceeded", Alias: (*Alias)(&obj)})
}

func (obj ResourceSizeLimitExceededError) Error() string {
	return obj.Message
}

// SearchDeactivatedError implements the interface ErrorObject
type SearchDeactivatedError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj SearchDeactivatedError) MarshalJSON() ([]byte, error) {
	type Alias SearchDeactivatedError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "SearchDeactivated", Alias: (*Alias)(&obj)})
}

func (obj SearchDeactivatedError) Error() string {
	return obj.Message
}

// SearchExecutionFailureError implements the interface ErrorObject
type SearchExecutionFailureError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj SearchExecutionFailureError) MarshalJSON() ([]byte, error) {
	type Alias SearchExecutionFailureError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "SearchExecutionFailure", Alias: (*Alias)(&obj)})
}

func (obj SearchExecutionFailureError) Error() string {
	return obj.Message
}

// SearchFacetPathNotFoundError implements the interface ErrorObject
type SearchFacetPathNotFoundError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj SearchFacetPathNotFoundError) MarshalJSON() ([]byte, error) {
	type Alias SearchFacetPathNotFoundError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "SearchFacetPathNotFound", Alias: (*Alias)(&obj)})
}

func (obj SearchFacetPathNotFoundError) Error() string {
	return obj.Message
}

// SearchIndexingInProgressError implements the interface ErrorObject
type SearchIndexingInProgressError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj SearchIndexingInProgressError) MarshalJSON() ([]byte, error) {
	type Alias SearchIndexingInProgressError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "SearchIndexingInProgress", Alias: (*Alias)(&obj)})
}

func (obj SearchIndexingInProgressError) Error() string {
	return obj.Message
}

// SemanticErrorError implements the interface ErrorObject
type SemanticErrorError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj SemanticErrorError) MarshalJSON() ([]byte, error) {
	type Alias SemanticErrorError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "SemanticError", Alias: (*Alias)(&obj)})
}

func (obj SemanticErrorError) Error() string {
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

// SyntaxErrorError implements the interface ErrorObject
type SyntaxErrorError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj SyntaxErrorError) MarshalJSON() ([]byte, error) {
	type Alias SyntaxErrorError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "SyntaxError", Alias: (*Alias)(&obj)})
}

func (obj SyntaxErrorError) Error() string {
	return obj.Message
}

// VariantValues is a standalone struct
type VariantValues struct {
	SKU        string       `json:"sku,omitempty"`
	Prices     []PriceDraft `json:"prices"`
	Attributes []Attribute  `json:"attributes"`
}

// WeakPasswordError implements the interface ErrorObject
type WeakPasswordError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj WeakPasswordError) MarshalJSON() ([]byte, error) {
	type Alias WeakPasswordError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "WeakPassword", Alias: (*Alias)(&obj)})
}

func (obj WeakPasswordError) Error() string {
	return obj.Message
}
