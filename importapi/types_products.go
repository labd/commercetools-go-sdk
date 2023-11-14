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
	// The tokenizer defines the tokens that are used to match against the [Suggest Query](/../products-suggestions#suggest-query) input.
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
*	The tokenizer defines the tokens that are used to match against the [Suggest Query](/../products-suggestions#suggest-query) input.
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
*	The data representation for a Product to be imported that is persisted as a [Product](/../api/projects/products#product) in the Project.
*
*	This is the minimal representation required for creating a [Product](/../api/projects/products#product) in commercetools.
*
 */
type ProductImport struct {
	// User-defined unique identifier. If a [Product](/../api/projects/products#product) with this `key` exists, it will be updated with the imported data.
	Key string `json:"key"`
	// Maps to `Product.name`.
	Name LocalizedString `json:"name"`
	// The `productType` of a [Product](/../api/projects/products#product).
	// Maps to `Product.productType`.
	// The Reference to the [ProductType](/../api/projects/productTypes#producttype) with which the Product is associated.
	// If referenced ProductType does not exist, the `state` of the [ImportOperation](/import-operation#importoperation) will be set to `unresolved` until the necessary ProductType is created.
	ProductType ProductTypeKeyReference `json:"productType"`
	// Human-readable identifiers usually used as deep-link URL to the related product. Each slug must be unique across a Project,
	// but a product can have the same slug for different languages. Allowed are alphabetic, numeric, underscore (_) and hyphen (-) characters.
	Slug LocalizedString `json:"slug"`
	// Maps to `Product.description`.
	Description *LocalizedString `json:"description,omitempty"`
	// Maps to `Product.categories`.
	// The References to the [Categories](/../api/projects/categories#category) with which the Product is associated.
	// If referenced Categories do not exist, the `state` of the [ImportOperation](/import-operation#importoperation) will be set to `unresolved` until the necessary Categories are created.
	Categories []CategoryKeyReference `json:"categories"`
	// A localized string is a JSON object where the keys are of [IETF language tag](https://en.wikipedia.org/wiki/IETF_language_tag), and the values the corresponding strings used for that language.
	// ```json
	// {
	//   "de": "Hundefutter",
	//   "en": "dog food"
	// }
	// ```
	MetaTitle *LocalizedString `json:"metaTitle,omitempty"`
	// A localized string is a JSON object where the keys are of [IETF language tag](https://en.wikipedia.org/wiki/IETF_language_tag), and the values the corresponding strings used for that language.
	// ```json
	// {
	//   "de": "Hundefutter",
	//   "en": "dog food"
	// }
	// ```
	MetaDescription *LocalizedString `json:"metaDescription,omitempty"`
	// A localized string is a JSON object where the keys are of [IETF language tag](https://en.wikipedia.org/wiki/IETF_language_tag), and the values the corresponding strings used for that language.
	// ```json
	// {
	//   "de": "Hundefutter",
	//   "en": "dog food"
	// }
	// ```
	MetaKeywords *LocalizedString `json:"metaKeywords,omitempty"`
	// The Reference to the [TaxCategory](/../api/projects/taxCategories#taxcategory) with which the Product is associated.
	// If referenced TaxCategory does not exist, the `state` of the [ImportOperation](/import-operation#importoperation) will be set to `unresolved` until the necessary TaxCategory is created.
	TaxCategory *TaxCategoryKeyReference `json:"taxCategory,omitempty"`
	// Search keywords are primarily used by the suggester but are also considered for the full-text search. SearchKeywords is a JSON object where the keys are of [IETF language tag](https://en.wikipedia.org/wiki/IETF_language_tag). The value to a language tag key is an array of SearchKeyword for the specific language.
	// ```json
	// {
	//   "en": [
	//     { "text": "Multi tool" },
	//     { "text": "Swiss Army Knife", "suggestTokenizer": { "type": "whitespace" } }
	//   ],
	//   "de": [
	//     {
	//       "text": "Schweizer Messer",
	//       "suggestTokenizer": {
	//         "type": "custom",
	//         "inputs": ["schweizer messer", "offiziersmesser", "sackmesser"]
	//       }
	//     }
	//   ]
	// }
	// ```
	SearchKeywords *SearchKeywords `json:"searchKeywords,omitempty"`
	// The Reference to the [State](/../api/projects/states#state) with which the Product is associated.
	// If referenced State does not exist, the `state` of the [ImportOperation](/import-operation#importoperation) will be set to `unresolved` until the necessary State is created.
	State *StateKeyReference `json:"state,omitempty"`
	// If `publish` is set to either `true` or `false`, both staged and current projections are set to the same value provided by the import data.
	// If `publish` is not set, the staged projection is set to the provided import data, but the current projection stays unchanged.
	// However, if the import data contains no update, that is, if it matches the staged projection of the existing Product, the import induces no change in the existing Product whether `publish` is set or not.
	Publish *bool `json:"publish,omitempty"`
	// Determines the type of Prices the API uses. See [ProductPriceMode](/../api/projects/products#productpricemode) for more details. If not provided, the existing `Product.priceMode` is not changed.
	PriceMode *ProductPriceModeEnum `json:"priceMode,omitempty"`
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

	return json.Marshal(raw)

}
