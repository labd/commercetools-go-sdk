// Generated file, please do not change!!!
package history

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
	case "ReviewLabel":
		obj := ReviewLabel{}
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
	Key       string `json:"key"`
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
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
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
	CustomerEmail string `json:"customerEmail"`
	OrderNumber   string `json:"orderNumber"`
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
	Key           string `json:"key"`
	AmountPlanned Money  `json:"amountPlanned"`
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
	Slug LocalizedString `json:"slug"`
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

type ReviewLabel struct {
	Key   string `json:"key"`
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

type StringLabel struct {
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
