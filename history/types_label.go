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
			return nil, errors.New("Error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}

	switch discriminator {
	case "CustomerLabel":
		new := CustomerLabel{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "LocalizedLabel":
		new := LocalizedLabel{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderLabel":
		new := OrderLabel{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "PaymentLabel":
		new := PaymentLabel{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductLabel":
		new := ProductLabel{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ReviewLabel":
		new := ReviewLabel{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "StringLabel":
		new := StringLabel{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type CustomerLabel struct {
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	CustomerNumber string `json:"customerNumber"`
}

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
func (obj StringLabel) MarshalJSON() ([]byte, error) {
	type Alias StringLabel
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StringLabel", Alias: (*Alias)(&obj)})
}
