package importapi

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
)

/**
*	Search keywords are primarily used by the suggester but are also considered for the full-text search. SearchKeywords is a JSON object where the keys are of [IETF language tag](https://en.wikipedia.org/wiki/IETF_language_tag). The value to a language tag key is an array of SearchKeyword for the specific language.
*	```json
*	{
*	  "en": [
*	    { "text": "Multi tool" },
*	    { "text": "Swiss Army Knife", "suggestTokenizer": { "type": "whitespace" } }
*	  ],
*	  "de": [
*	    {
*	      "text": "Schweizer Messer",
*	      "suggestTokenizer": {
*	        "type": "custom",
*	        "inputs": ["schweizer messer", "offiziersmesser", "sackmesser"]
*	      }
*	    }
*	  ]
*	}
*	```
*
 */
type SearchKeywords map[string][]SearchKeyword
type SearchKeyword struct {
	Text string `json:"text"`
	// The tokenizer defines the tokens that are used for [search term suggestions](/projects/search-term-suggestions).
	SuggestTokenizer SuggestTokenizer `json:"suggestTokenizer,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *SearchKeyword) UnmarshalJSON(data []byte) error {
	type Alias SearchKeyword
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.SuggestTokenizer != nil {
		var err error
		obj.SuggestTokenizer, err = mapDiscriminatorSuggestTokenizer(obj.SuggestTokenizer)
		if err != nil {
			return err
		}
	}

	return nil
}

/**
*	The tokenizer defines the tokens that are used for [search term suggestions](/projects/search-term-suggestions).
*
 */
type SuggestTokenizer interface{}

func mapDiscriminatorSuggestTokenizer(input interface{}) (SuggestTokenizer, error) {
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
	case "custom":
		obj := CustomTokenizer{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "whitespace":
		obj := WhitespaceTokenizer{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type CustomTokenizer struct {
	Inputs []string `json:"inputs"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomTokenizer) MarshalJSON() ([]byte, error) {
	type Alias CustomTokenizer
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "custom", Alias: (*Alias)(&obj)})
}

type WhitespaceTokenizer struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj WhitespaceTokenizer) MarshalJSON() ([]byte, error) {
	type Alias WhitespaceTokenizer
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "whitespace", Alias: (*Alias)(&obj)})
}

/**
*	Represents the data used to import a Product. Once imported, this data is persisted as a [Product](ctp:api:type:Product) in the Project.
*
*	This is the minimal representation required for creating a Product in Composable Commerce.
*
 */
type ProductImport struct {
	// User-defined unique identifier. If a [Product](ctp:api:type:Product) with this `key` exists, it is updated with the imported data.
	Key string `json:"key"`
	// Maps to `ProductData.name`.
	Name LocalizedString `json:"name"`
	// Maps to `Product.productType`. If the referenced [ProductType](ctp:api:type:ProductType) does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced ProductType is created.
	ProductType ProductTypeKeyReference `json:"productType"`
	// Maps to `ProductData.slug`.
	Slug LocalizedString `json:"slug"`
	// Maps to `ProductData.description`.
	Description *LocalizedString `json:"description,omitempty"`
	// Maps to `ProductData.categories`. If the referenced [Categories](ctp:api:type:Category) do not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced Categories are created.
	Categories []CategoryKeyReference `json:"categories"`
	Attributes []Attribute            `json:"attributes"`
	// Maps to `ProductData.metaTitle`.
	MetaTitle *LocalizedString `json:"metaTitle,omitempty"`
	// Maps to `ProductData.metaDescription`.
	MetaDescription *LocalizedString `json:"metaDescription,omitempty"`
	// Maps to `ProductData.metaKeywords`.
	MetaKeywords *LocalizedString `json:"metaKeywords,omitempty"`
	// Maps to `Product.taxCategory`. If the referenced [TaxCategory](ctp:api:type:TaxCategory) does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced TaxCategory is created.
	TaxCategory *TaxCategoryKeyReference `json:"taxCategory,omitempty"`
	// Maps to `ProductData.searchKeywords`.
	SearchKeywords *SearchKeywords `json:"searchKeywords,omitempty"`
	// Maps to `Product.state`. If the referenced [State](ctp:api:type:State) does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced State is created.
	State *StateKeyReference `json:"state,omitempty"`
	// Determines the published status and current/staged projection of the Product. For more information, see [Managing the published state of Products](/import-export/best-practices#manage-published-state-of-products).
	Publish *bool `json:"publish,omitempty"`
	// Maps to `Product.priceMode`.
	PriceMode *ProductPriceModeEnum `json:"priceMode,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductImport) UnmarshalJSON(data []byte) error {
	type Alias ProductImport
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Attributes {
		var err error
		obj.Attributes[i], err = mapDiscriminatorAttribute(obj.Attributes[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductImport) MarshalJSON() ([]byte, error) {
	type Alias ProductImport
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

	if raw["categories"] == nil {
		delete(raw, "categories")
	}

	if raw["attributes"] == nil {
		delete(raw, "attributes")
	}

	return json.Marshal(raw)

}
