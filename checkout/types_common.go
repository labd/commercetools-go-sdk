package checkout

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
)

/**
*	A Reference represents a loose reference to another resource in the same Project identified by its `id`. The `typeId` indicates the type of the referenced resource. Each resource type has its corresponding Reference type, like [ApplicationReference](ctp:checkout:type:ApplicationReference).
*
 */
type Reference interface{}

func mapDiscriminatorReference(input interface{}) (Reference, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["typeId"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'typeId'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "application":
		obj := ApplicationReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "cart":
		obj := CartReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "order":
		obj := OrderReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "payment-method":
		obj := PaymentMethodReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "payment":
		obj := PaymentReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "payment-integration":
		obj := PaymentIntegrationReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "recurring-order":
		obj := RecurringOrderReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "recurring-payment":
		obj := RecurringPaymentReference{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Reference to a [PaymentMethod](ctp:api:type:PaymentMethod).
*
 */
type PaymentMethodReference struct {
	// Unique identifier of the referenced [PaymentMethod](ctp:api:type:PaymentMethod).
	ID string `json:"id"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentMethodReference) MarshalJSON() ([]byte, error) {
	type Alias PaymentMethodReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "payment-method", Alias: (*Alias)(&obj)})
}

/**
*	Type of resource the value should reference. Supported resource type identifiers are:
*
 */
type ReferenceTypeId string

const (
	ReferenceTypeIdCart               ReferenceTypeId = "cart"
	ReferenceTypeIdOrder              ReferenceTypeId = "order"
	ReferenceTypeIdPayment            ReferenceTypeId = "payment"
	ReferenceTypeIdApplication        ReferenceTypeId = "application"
	ReferenceTypeIdPaymentIntegration ReferenceTypeId = "payment-integration"
	ReferenceTypeIdDeployment         ReferenceTypeId = "deployment"
	ReferenceTypeIdPaymentMethod      ReferenceTypeId = "payment-method"
	ReferenceTypeIdRecurringPayment   ReferenceTypeId = "recurring-payment"
	ReferenceTypeIdRecurringOrder     ReferenceTypeId = "recurring-order"
)

/**
*	Draft type to create a [Reference](ctp:checkout:type:Reference) to a resource. Provide either the `id` or (wherever supported) the `key` of the resource to reference, but depending on the API endpoint the response returns a Reference.
*
*	Each resource type has its corresponding ResourceIdentifier, like [PaymentResourceIdentifier](ctp:checkout:type:PaymentResourceIdentifier).
*
 */
type ResourceIdentifier interface{}

func mapDiscriminatorResourceIdentifier(input interface{}) (ResourceIdentifier, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["typeId"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'typeId'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "application":
		obj := ApplicationResourceIdentifier{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "cart":
		obj := CartResourceIdentifier{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "order":
		obj := OrderResourceIdentifier{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "payment":
		obj := PaymentResourceIdentifier{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "payment-integration":
		obj := PaymentIntegrationResourceIdentifier{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Money in cent amounts for a specific currency.
*
 */
type Amount struct {
	// Amount in the smallest indivisible unit of a currency, such as:
	//
	// * Cents for EUR and USD, pence for GBP, or centime for CHF (5 CHF is specified as `500`).
	// * The value in the major unit for currencies without minor units, like JPY (5 JPY is specified as `5`).
	CentAmount int `json:"centAmount"`
	// Currency code compliant to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	CurrencyCode string `json:"currencyCode"`
}

/**
*	JSON object where the keys are of type [Locale](ctp:checkout:type:Locale), and the values are the strings used for the corresponding language.
*
 */
type LocalizedString map[string]string

/**
*	A localized URL that provides different URL values for different locales.
*
 */
type LocalizedUrl map[string]string

/**
*	The Region in which the Checkout application is [hosted](/hosts-and-authorization).
*
 */
type Region string

const (
	RegionEuropeWest1Gcp         Region = "europe-west1.gcp"
	RegionEuCentral1Aws          Region = "eu-central-1.aws"
	RegionUsCentral1Gcp          Region = "us-central1.gcp"
	RegionUsEast2Aws             Region = "us-east-2.aws"
	RegionAustraliaSoutheast1Gcp Region = "australia-southeast1.gcp"
)
