package history

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
)

/**
*	Provides descriptive information specific to the resource.
 */
type Label interface{}

func mapDiscriminatorLabel(input interface{}) (Label, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "CustomObjectLabel":
		obj := CustomObjectLabel{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "CustomerLabel":
		obj := CustomerLabel{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "LocalizedLabel":
		obj := LocalizedLabel{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "OrderLabel":
		obj := OrderLabel{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "PaymentLabel":
		obj := PaymentLabel{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ProductLabel":
		obj := ProductLabel{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "QuoteLabel":
		obj := QuoteLabel{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "QuoteRequestLabel":
		obj := QuoteRequestLabel{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ReviewLabel":
		obj := ReviewLabel{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "StagedQuoteLabel":
		obj := StagedQuoteLabel{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "StringLabel":
		obj := StringLabel{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type CustomObjectLabel struct {
	// User-defined unique identifier of the CustomObject within the defined `container`.
	Key string `json:"key"`
	// Namespace to group Custom Objects.
	Container string `json:"container"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomObjectLabel) MarshalJSON() ([]byte, error) {
	type Alias CustomObjectLabel
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomObjectLabel", Alias: (*Alias)(&obj)})
}

type CustomerLabel struct {
	// Given name (first name) of the Customer.
	FirstName string `json:"firstName"`
	// Family name (last name) of the Customer.
	LastName string `json:"lastName"`
	// User-defined unique identifier of the [Customer](ctp:api:type:Customer).
	CustomerNumber string `json:"customerNumber"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerLabel) MarshalJSON() ([]byte, error) {
	type Alias CustomerLabel
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerLabel", Alias: (*Alias)(&obj)})
}

type LocalizedLabel struct {
	// Changed value.
	Value LocalizedString `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj LocalizedLabel) MarshalJSON() ([]byte, error) {
	type Alias LocalizedLabel
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "LocalizedLabel", Alias: (*Alias)(&obj)})
}

type OrderLabel struct {
	// Email address of the Customer that the Order belongs to.
	CustomerEmail string `json:"customerEmail"`
	// User-defined unique identifier of the Order that is unique across a Project.
	OrderNumber string `json:"orderNumber"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderLabel) MarshalJSON() ([]byte, error) {
	type Alias OrderLabel
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderLabel", Alias: (*Alias)(&obj)})
}

type PaymentLabel struct {
	// User-defined unique identifier of the Payment.
	Key string `json:"key"`
	// Money value the Payment intends to receive from the Customer.
	AmountPlanned Money `json:"amountPlanned"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentLabel) MarshalJSON() ([]byte, error) {
	type Alias PaymentLabel
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "PaymentLabel", Alias: (*Alias)(&obj)})
}

type ProductLabel struct {
	// User-defined identifier used in a deep-link URL for the Product.
	Slug LocalizedString `json:"slug"`
	// Name of the Product.
	Name LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductLabel) MarshalJSON() ([]byte, error) {
	type Alias ProductLabel
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductLabel", Alias: (*Alias)(&obj)})
}

type QuoteLabel struct {
	// User-defined unique identifier of the Quote.
	Key string `json:"key"`
	// The [Buyer](/../api/quotes-overview#buyer) who requested the Quote.
	Customer Reference `json:"customer"`
	// Staged Quote related to the Quote.
	StagedQuote Reference `json:"stagedQuote"`
	// Quote Request related to the Quote.
	QuoteRequest Reference `json:"quoteRequest"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj QuoteLabel) MarshalJSON() ([]byte, error) {
	type Alias QuoteLabel
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "QuoteLabel", Alias: (*Alias)(&obj)})
}

type QuoteRequestLabel struct {
	// User-defined unique identifier of the Quote Request.
	Key string `json:"key"`
	// The [Buyer](/../api/quotes-overview#buyer) who raised the Quote Request.
	Customer Reference `json:"customer"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj QuoteRequestLabel) MarshalJSON() ([]byte, error) {
	type Alias QuoteRequestLabel
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "QuoteRequestLabel", Alias: (*Alias)(&obj)})
}

type ReviewLabel struct {
	// User-defined unique identifier of the Review.
	Key string `json:"key"`
	// Title of the Review.
	Title string `json:"title"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ReviewLabel) MarshalJSON() ([]byte, error) {
	type Alias ReviewLabel
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ReviewLabel", Alias: (*Alias)(&obj)})
}

type StagedQuoteLabel struct {
	// User-defined unique identifier of the Staged Quote.
	Key string `json:"key"`
	// The [Buyer](/../api/quotes-overview#buyer) who requested the Quote.
	Customer Reference `json:"customer"`
	// Quote Request related to the Staged Quote.
	QuoteRequest Reference `json:"quoteRequest"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedQuoteLabel) MarshalJSON() ([]byte, error) {
	type Alias StagedQuoteLabel
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StagedQuoteLabel", Alias: (*Alias)(&obj)})
}

type StringLabel struct {
	// Changed value.
	Value string `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StringLabel) MarshalJSON() ([]byte, error) {
	type Alias StringLabel
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StringLabel", Alias: (*Alias)(&obj)})
}
