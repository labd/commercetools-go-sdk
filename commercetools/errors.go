package commercetools

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

// Add the errors as listed in the commercetools documentation
const (
	// Custom error
	ErrNotSpecified     = "NotSpecified"
	ErrInvalidJSONInput = "InvalidJsonInput" // Not documented

	// General
	ErrGeneral                               = "General"
	ErrOverCapacity                          = "OverCapacity"
	ErrPendingOperation                      = "PendingOperation"
	ErrResourceNotFound                      = "ResourceNotFound"
	ErrInvalidInput                          = "InvalidInput"
	ErrInvalidOperation                      = "InvalidOperation"
	ErrInvalidField                          = "InvalidField"
	ErrRequiredField                         = "RequiredField"
	ErrDuplicateField                        = "DuplicateField"
	ErrDuplicateFieldWithConflictingResource = "DuplicateFieldWithConflictingResource"
	ErrConcurrentModification                = "ConcurrentModification"

	// Extensions
	ErrExtensionBadResponse         = "ExtensionBadResponse"
	ErrExtensionUpdateActionsFailed = "ExtensionUpdateActionsFailed"
	ErrExtensionNoResponse          = "ExtensionNoResponse"

	// Products
	ErrDuplicatePriceScope      = "DuplicatePriceScope"
	ErrDuplicateVariantValues   = "DuplicateVariantValues"
	ErrDuplicateAttributeValue  = "DuplicateAttributeValue"
	ErrDuplicateAttributeValues = "DuplicateAttributeValues"

	// Product Types
	ErrAttributeDefinitionAlreadyExists = "AttributeDefinitionAlreadyExists"
	ErrAttributeDefinitionTypeConflict  = "AttributeDefinitionTypeConflict"
	ErrDuplicateEnumValues              = "DuplicateEnumValues"
	ErrEnumKeyAlreadyExists             = "EnumKeyAlreadyExists"
	ErrEnumKeyDoesNotExist              = "EnumKeyDoesNotExist"
	ErrAttributeNameDoesNotExist        = "AttributeNameDoesNotExist"
	ErrEnumValuesMustMatch              = "EnumValuesMustMatch"

	// Orders
	ErrOutOfStock                     = "OutOfStock"
	ErrPriceChanged                   = "PriceChanged"
	ErrDiscountCodeNonApplicable      = "DiscountCodeNonApplicable"
	ErrShippingMethodDoesNotMatchCart = "ShippingMethodDoesNotMatchCart"
	ErrInvalidItemShippingDetails     = "InvalidItemShippingDetails"
	ErrMatchingPriceNotFound          = "MatchingPriceNotFound"
	ErrMissingTaxRateForCountry       = "MissingTaxRateForCountry"

	// Customers
	ErrInvalidCredentials     = "InvalidCredentials"
	ErrInvalidCurrentPassword = "InvalidCurrentPassword"
	// ErrMissingTaxRateForCountry = "MissingTaxRateForCountry" // (duplicate)

	// Product Discounts
	ErrNoMatchingProductDiscountFound = "NoMatchingProductDiscountFound"

	// Shipping Methods
	ErrEditPreviewFailed = "EditPreviewFailed"
)

type Error struct {
	message    string
	statusCode int
	Errors     []ErrorItem
}

func (e *Error) Code() string {
	if len(e.Errors) > 0 {
		return e.Errors[0].Code()
	}
	return ErrNotSpecified
}

func (e *Error) Message() string {
	var message string

	switch len(e.Errors) {
	case 0:
		message = e.message
	case 1:
		message = e.Errors[0].Message()
	default:
		{
			messages := make([]string, len(e.Errors))
			for i, item := range e.Errors {
				messages[i] = fmt.Sprintf(" %d. %s", i+1, item.Message())
			}
			message = "Mutiple errors returned:\n" + strings.Join(messages, "\n")
		}
	}

	if message == "" {
		return fmt.Sprintf("%s (%d)", http.StatusText(e.statusCode), e.statusCode)
	}
	return fmt.Sprintf("%s (%d): %s", http.StatusText(e.statusCode), e.statusCode, message)
}

// StatusCode contains the HTTP Status code returned by the commercetools
// platform.
func (e *Error) StatusCode() int {
	return e.statusCode
}

func (e Error) Error() string {
	return e.Message()
}

type ErrorItem interface {
	Code() string
	Message() string
	Extra() map[string]string
}

type ErrorObject struct {
	code    string
	message string
	extra   map[string]string
}

func (e ErrorObject) Code() string {
	return e.code
}

func (e ErrorObject) Extra() map[string]string {
	return e.extra
}

func (e ErrorObject) Message() string {
	if e.message == "" {
		return e.code
	}
	return fmt.Sprintf("%s - %s", e.code, e.message)
}

type ErrorInvalidJson struct {
	*ErrorObject
}

func (e ErrorInvalidJson) Message() string {
	return fmt.Sprintf("%s - %s (%s)", e.code, e.message, e.DetailedErrorMessage())
}

func (e ErrorInvalidJson) DetailedErrorMessage() string {
	return e.Extra()["detailedErrorMessage"]
}

func parseErrorResponse(response *http.Response, body []byte) error {
	if len(body) == 0 {
		return Error{
			statusCode: response.StatusCode,
			Errors:     []ErrorItem{ErrorObject{code: ErrResourceNotFound}},
		}
	}

	data := make(map[string]interface{})
	err := json.Unmarshal(body, &data)
	if err != nil {
		return errors.Wrap(err, "Unmarshalling error response failed")
	}

	if val, ok := data["message"]; ok {
		errors := parseErrorResponseObjects(data["errors"])
		return Error{
			message:    val.(string),
			statusCode: response.StatusCode,
			Errors:     errors,
		}
	}

	return Error{
		message:    string(body),
		statusCode: response.StatusCode,
	}
}

func parseErrorResponseObjects(errors interface{}) []ErrorItem {
	errSlice := errors.([]interface{})
	result := make([]ErrorItem, len(errSlice))

	for i, item := range errSlice {
		data := item.(map[string]interface{})
		code := data["code"].(string)
		msg := data["message"].(string)

		// Extra error fields
		extra := map[string]string{}
		for k, v := range data {
			if k != "code" && k != "message" {
				extra[k] = v.(string)
			}
		}

		var item ErrorItem
		base := ErrorObject{code: code, message: msg, extra: extra}

		switch code {
		case ErrInvalidJSONInput:
			item = ErrorInvalidJson{&base}
		default:
			item = base
		}
		result[i] = item
	}
	return result
}
