package commercetools_test

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/stretchr/testify/assert"

	"github.com/labd/commercetools-go-sdk/testutil"
)

func TestProductTypeCreate(t *testing.T) {
	output := testutil.RequestData{}

	client, server := testutil.MockClient(t, testutil.Fixture("producttype.create.json"), &output, nil)
	defer server.Close()
	input := &commercetools.ProductTypeDraft{
		Name:        "test_product_type",
		Description: "Test product type.",
		Attributes: []commercetools.AttributeDefinitionDraft{
			{
				Type:         commercetools.AttributeTextType{},
				IsSearchable: false,
				InputHint:    "SingleLine",
				Name:         "size",
				Label: &commercetools.LocalizedString{
					"en": "The right size is important.",
				},
				IsRequired:          false,
				AttributeConstraint: commercetools.AttributeConstraintEnum("CombinationUnique"),
			},
		},
	}

	fmt.Println(output)

	_, err := client.ProductTypeCreate(input)
	assert.Nil(t, err)

	expectedBody := `{
		"name": "test_product_type",
		"description": "Test product type.",
		"attributes": [
		  	{
				"type": {
			  		"name": "text"
				},
				"inputHint": "SingleLine",
				"name": "size",
				"label": {
			  		"en": "The right size is important."
				},
				"isRequired": false,
				"isSearchable": false,
				"attributeConstraint": "CombinationUnique"
		  	}
		]
	}`
	assert.JSONEq(t, expectedBody, output.JSON)
}

func TestProductTypeUpdate(t *testing.T) {
	testCases := []struct {
		desc        string
		input       *commercetools.ProductTypeUpdateInput
		requestBody string
	}{
		{
			desc: "Set key",
			input: &commercetools.ProductTypeUpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: []commercetools.ProductTypeUpdateAction{
					&commercetools.ProductTypeSetKeyAction{
						Key: "123456",
					},
				},
			},
			requestBody: `{
				"version": 3,
				"actions": [
					{
						"action": "setKey",
						"key": "123456"
					}
				]
			}`,
		},
		{
			desc: "Change name",
			input: &commercetools.ProductTypeUpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: []commercetools.ProductTypeUpdateAction{
					&commercetools.ProductTypeChangeNameAction{
						Name: "Changed name",
					},
				},
			},
			requestBody: `{
				"version": 3,
				"actions": [
					{
						"action": "changeName",
						"name": "Changed name"
					}
				]
			}`,
		},
		{
			desc: "Change description",
			input: &commercetools.ProductTypeUpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: []commercetools.ProductTypeUpdateAction{
					&commercetools.ProductTypeChangeDescriptionAction{
						Description: "Changed description",
					},
				},
			},
			requestBody: `{
				"version": 3,
				"actions": [
					{
						"action": "changeDescription",
						"description": "Changed description"
					}
				]
			}`,
		},
		{
			desc: "Add attribute definition",
			input: &commercetools.ProductTypeUpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: []commercetools.ProductTypeUpdateAction{
					&commercetools.ProductTypeAddAttributeDefinitionAction{
						Attribute: &commercetools.AttributeDefinitionDraft{
							Type: commercetools.AttributeBooleanType{},
							Name: "attribute",
							Label: &commercetools.LocalizedString{
								"en": "attribute",
							},
							IsRequired:          false,
							AttributeConstraint: commercetools.AttributeConstraintEnum("None"),
							InputTip: &commercetools.LocalizedString{
								"en": "A simple hint",
							},
							InputHint:    "SingleLine",
							IsSearchable: true,
						},
					},
				},
			},
			requestBody: `{
				"version": 3,
				"actions": [
					{
						"action": "addAttributeDefinition",
						"attribute": {
							"type": {
								"name": "boolean"
							},
							"name": "attribute",
							"label": {
								"en": "attribute"
							},
							"isRequired": false,
							"attributeConstraint": "None",
							"inputTip": {
								"en": "A simple hint"
							},
							"inputHint": "SingleLine",
							"isSearchable": true
						}
					}
				]
			}`,
		},
		{
			desc: "Remove attribute definition",
			input: &commercetools.ProductTypeUpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: []commercetools.ProductTypeUpdateAction{
					&commercetools.ProductTypeRemoveAttributeDefinitionAction{
						Name: "attribute",
					},
				},
			},
			requestBody: `{
				"version": 3,
				"actions": [
					{
						"action": "removeAttributeDefinition",
						"name": "attribute"
					}
				]
			}`,
		},
		{
			desc: "Change attribute definition name",
			input: &commercetools.ProductTypeUpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: []commercetools.ProductTypeUpdateAction{
					&commercetools.ProductTypeChangeAttributeNameAction{
						AttributeName:    "attribute",
						NewAttributeName: "new-attribute",
					},
				},
			},
			requestBody: `{
				"version": 3,
				"actions": [
					{
						"action": "changeAttributeName",
						"attributeName": "attribute",
						"newAttributeName": "new-attribute"
					}
				]
			}`,
		},
		{
			desc: "Change attribute definition label",
			input: &commercetools.ProductTypeUpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: []commercetools.ProductTypeUpdateAction{
					&commercetools.ProductTypeChangeLabelAction{
						AttributeName: "attribute",
						Label: &commercetools.LocalizedString{
							"en": "label",
						},
					},
				},
			},
			requestBody: `{
				"version": 3,
				"actions": [
					{
						"action": "changeLabel",
						"attributeName": "attribute",
						"label": {
							"en": "label"
						}
					}
				]
			}`,
		},
		{
			desc: "Set attribute definition input tip",
			input: &commercetools.ProductTypeUpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: []commercetools.ProductTypeUpdateAction{
					&commercetools.ProductTypeSetInputTipAction{
						AttributeName: "attribute",
						InputTip: &commercetools.LocalizedString{
							"en": "Small but helpful hint.",
						},
					},
				},
			},
			requestBody: `{
				"version": 3,
				"actions": [
					{
						"action": "setInputTip",
						"attributeName": "attribute",
						"inputTip": {
							"en": "Small but helpful hint."
						}
					}
				]
			}`,
		},
		{
			desc: "Add plain enum value",
			input: &commercetools.ProductTypeUpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: []commercetools.ProductTypeUpdateAction{
					&commercetools.ProductTypeAddPlainEnumValueAction{
						AttributeName: "attribute",
						Value: &commercetools.AttributePlainEnumValue{
							Key:   "enum",
							Label: "A plain enum value.",
						},
					},
				},
			},
			requestBody: `{
				"version": 3,
				"actions": [
					{
						"action": "addPlainEnumValue",
						"attributeName": "attribute",
						"value": {
							"key": "enum",
							"label": "A plain enum value."
						}
					}
				]
			}`,
		},
		{
			desc: "Add localized enum value",
			input: &commercetools.ProductTypeUpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: []commercetools.ProductTypeUpdateAction{
					&commercetools.ProductTypeAddLocalizedEnumValueAction{
						AttributeName: "attribute",
						Value: &commercetools.AttributeLocalizedEnumValue{
							Key: "enum",
							Label: &commercetools.LocalizedString{
								"en": "A localized enum value.",
							},
						},
					},
				},
			},
			requestBody: `{
				"version": 3,
				"actions": [
					{
						"action": "addLocalizedEnumValue",
						"attributeName": "attribute",
						"value": {
							"key": "enum",
							"label": {
								"en": "A localized enum value."
							}
						}
					}
				]
			}`,
		},
		{
			desc: "Remove enum values",
			input: &commercetools.ProductTypeUpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: []commercetools.ProductTypeUpdateAction{
					&commercetools.ProductTypeRemoveEnumValuesAction{
						AttributeName: "attribute",
						Keys:          []string{"key-1", "key-2"},
					},
				},
			},
			requestBody: `{
				"version": 3,
				"actions": [
					{
						"action": "removeEnumValues",
						"attributeName": "attribute",
						"keys": ["key-1", "key-2"]
					}
				]
			}`,
		},
		{
			desc: "Change attribute definitions order",
			input: &commercetools.ProductTypeUpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: []commercetools.ProductTypeUpdateAction{
					&commercetools.ProductTypeChangeAttributeOrderAction{
						Attributes: []commercetools.AttributeDefinition{
							{
								Type:         commercetools.AttributeTextType{},
								IsSearchable: false,
								InputHint:    "SingleLine",
								Name:         "size",
								Label: &commercetools.LocalizedString{
									"en": "The right size is important.",
								},
								IsRequired:          false,
								AttributeConstraint: commercetools.AttributeConstraintEnum("CombinationUnique"),
							},
						},
					},
				},
			},
			requestBody: `{
				"version": 3,
				"actions": [
					{
						"action": "changeAttributeOrder",
						"attributes": [
							{
								"type": {
									"name": "text"
								},
								"isSearchable": false,
								"inputHint": "SingleLine",
								"name": "size",
								"label": {
									"en": "The right size is important."
								},
								"isRequired": false,
								"attributeConstraint": "CombinationUnique"
							}
						]
					}
				]
			}`,
		},
		{
			desc: "Change plain enum value order",
			input: &commercetools.ProductTypeUpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: []commercetools.ProductTypeUpdateAction{
					&commercetools.ProductTypeChangePlainEnumValueOrderAction{
						AttributeName: "attribute",
						Values: []commercetools.AttributePlainEnumValue{
							{
								Key:   "enum",
								Label: "A plain enum value.",
							},
						},
					},
				},
			},
			requestBody: `{
				"version": 3,
				"actions": [
					{
						"action": "changePlainEnumValueOrder",
						"attributeName": "attribute",
						"values": [
							{
								"key": "enum",
								"label": "A plain enum value."
							}
						]
					}
				]
			}`,
		},
		{
			desc: "Change localized enum value order",
			input: &commercetools.ProductTypeUpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: []commercetools.ProductTypeUpdateAction{
					&commercetools.ProductTypeChangeLocalizedEnumValueOrderAction{
						AttributeName: "attribute",
						Values: []commercetools.AttributeLocalizedEnumValue{
							{
								Key: "enum",
								Label: &commercetools.LocalizedString{
									"en": "A localized enum value.",
								},
							},
						},
					},
				},
			},
			requestBody: `{
				"version": 3,
				"actions": [
					{
						"action": "changeLocalizedEnumValueOrder",
						"attributeName": "attribute",
						"values": [
							{
								"key": "enum",
								"label": {
									"en": "A localized enum value."
								}
							}
						]
					}
				]
			}`,
		},
		{
			desc: "Change enum value key",
			input: &commercetools.ProductTypeUpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: []commercetools.ProductTypeUpdateAction{
					&commercetools.ProductTypeChangeEnumKeyAction{
						AttributeName: "attribute",
						Key:           "enum",
						NewKey:        "new-enum",
					},
				},
			},
			requestBody: `{
				"version": 3,
				"actions": [
					{
						"action": "changeEnumKey",
						"attributeName": "attribute",
						"key": "enum",
						"newKey": "new-enum"
					}
				]
			}`,
		},
		{
			desc: "Change plain enum value label",
			input: &commercetools.ProductTypeUpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: []commercetools.ProductTypeUpdateAction{
					&commercetools.ProductTypeChangePlainEnumValueLabelAction{
						AttributeName: "attribute",
						NewValue: &commercetools.AttributePlainEnumValue{
							Key:   "enum",
							Label: "A descriptive label.",
						},
					},
				},
			},
			requestBody: `{
				"version": 3,
				"actions": [
					{
						"action": "changePlainEnumValueLabel",
						"attributeName": "attribute",
						"newValue": {
							"key": "enum",
							"label": "A descriptive label."
						}
					}
				]
			}`,
		},
		{
			desc: "Change localized enum value label",
			input: &commercetools.ProductTypeUpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: []commercetools.ProductTypeUpdateAction{
					&commercetools.ProductTypeChangeLocalizedEnumValueLabelAction{
						AttributeName: "attribute",
						NewValue: &commercetools.AttributeLocalizedEnumValue{
							Key: "enum",
							Label: &commercetools.LocalizedString{
								"en": "A descriptive label.",
							},
						},
					},
				},
			},
			requestBody: `{
				"version": 3,
				"actions": [
					{
						"action": "changeLocalizedEnumValueLabel",
						"attributeName": "attribute",
						"newValue": {
							"key": "enum",
							"label": {
								"en": "A descriptive label."
							}
						}
					}
				]
			}`,
		},
		{
			desc: "Change is searchable",
			input: &commercetools.ProductTypeUpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: []commercetools.ProductTypeUpdateAction{
					&commercetools.ProductTypeChangeIsSearchableAction{
						AttributeName: "attribute",
						IsSearchable:  true,
					},
				},
			},
			requestBody: `{
				"version": 3,
				"actions": [
					{
						"action": "changeIsSearchable",
						"attributeName": "attribute",
						"isSearchable": true
					}
				]
			}`,
		},
		{
			desc: "Change input hint",
			input: &commercetools.ProductTypeUpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: []commercetools.ProductTypeUpdateAction{
					&commercetools.ProductTypeChangeInputHintAction{
						AttributeName: "attribute",
						NewValue:      "SingleLine",
					},
				},
			},
			requestBody: `{
				"version": 3,
				"actions": [
					{
						"action": "changeInputHint",
						"attributeName": "attribute",
						"newValue": "SingleLine"
					}
				]
			}`,
		},
		{
			desc: "Change attribute constraint",
			input: &commercetools.ProductTypeUpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: []commercetools.ProductTypeUpdateAction{
					&commercetools.ProductTypeChangeAttributeConstraintAction{
						AttributeName: "attribute",
						NewValue:      commercetools.AttributeConstraintEnumDraft("CombinationUnique"),
					},
				},
			},
			requestBody: `{
				"version": 3,
				"actions": [
					{
						"action": "changeAttributeConstraint",
						"attributeName": "attribute",
						"newValue": "CombinationUnique"
					}
				]
			}`,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			output := testutil.RequestData{}

			client, server := testutil.MockClient(t, testutil.Fixture("producttype.update.json"), &output, nil)
			defer server.Close()

			_, err := client.ProductTypeUpdate(tC.input)
			assert.Nil(t, err)
			assert.Equal(t, "/unittest/product-types/1234", output.URL.Path)
			assert.JSONEq(t, tC.requestBody, output.JSON)
		})
	}
}

func TestProductTypeDeleteByID(t *testing.T) {
	output := testutil.RequestData{}

	client, server := testutil.MockClient(t, "{}", &output, nil)
	defer server.Close()

	_, err := client.ProductTypeDeleteByID("1234", 3)
	assert.Nil(t, err)

	params := url.Values{}
	params.Add("version", "3")
	assert.Equal(t, params, output.URL.Query())
	assert.Equal(t, "/unittest/product-types/1234", output.URL.Path)
}

func TestProductTypeDeleteByKey(t *testing.T) {
	output := testutil.RequestData{}

	client, server := testutil.MockClient(t, "{}", &output, nil)
	defer server.Close()

	_, err := client.ProductTypeDeleteByKey("1234", 3)
	assert.Nil(t, err)

	params := url.Values{}
	params.Add("version", "3")
	assert.Equal(t, params, output.URL.Query())
	assert.Equal(t, "/unittest/product-types/key=1234", output.URL.Path)
}

func TestProductTypeGetByID(t *testing.T) {
	timestamp, _ := time.Parse(time.RFC3339, "1970-01-01T00:00:00.001Z")

	testCases := []struct {
		desc    string
		input   *commercetools.ProductType
		fixture string
	}{
		{
			desc: "Type with size text attribute definition",
			input: &commercetools.ProductType{
				ID:          "1234",
				Version:     1,
				Name:        "test_product_type",
				Description: "Test product type.",
				Attributes: []commercetools.AttributeDefinition{
					{
						Type:         commercetools.AttributeTextType{},
						IsSearchable: false,
						InputHint:    "SingleLine",
						Name:         "size",
						Label: &commercetools.LocalizedString{
							"en": "The right size is important.",
						},
						IsRequired:          false,
						AttributeConstraint: commercetools.AttributeConstraintEnum("CombinationUnique"),
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "producttype.size.json",
		},
		{
			desc: "Type with color text attribute definition",
			input: &commercetools.ProductType{
				ID:          "1234",
				Version:     1,
				Name:        "another-test_product_type",
				Description: "Another test product type.",
				Attributes: []commercetools.AttributeDefinition{
					{
						Type:         commercetools.AttributeLocalizableTextType{},
						IsSearchable: false,
						InputHint:    "SingleLine",
						Name:         "color",
						Label: &commercetools.LocalizedString{
							"en": "The right color is important.",
						},
						IsRequired:          false,
						AttributeConstraint: commercetools.AttributeConstraintEnum("CombinationUnique"),
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "producttype.color.json",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			client, server := testutil.MockClient(t, testutil.Fixture(tC.fixture), nil, nil)
			defer server.Close()

			result, err := client.ProductTypeGetByID("1234")
			assert.Nil(t, err)
			assert.Equal(t, tC.input, result)
		})
	}
}

func TestAttributeTypes(t *testing.T) {
	testCases := []struct {
		desc   string
		object interface {
			MarshalJSON() ([]byte, error)
		}
		json string
	}{
		{
			desc:   "Boolean type",
			object: commercetools.AttributeBooleanType{},
			json: `{
				"name": "boolean"
			}`,
		},
		{
			desc:   "Text type",
			object: commercetools.AttributeTextType{},
			json: `{
				"name": "text"
			}`,
		},
		{
			desc:   "Localized text type",
			object: commercetools.AttributeLocalizableTextType{},
			json: `{
				"name": "ltext"
			}`,
		},
		{
			desc: "Enum type",
			object: commercetools.AttributeEnumType{
				Values: []commercetools.AttributePlainEnumValue{
					{
						Key:   "test",
						Label: "test",
					},
				},
			},
			json: `{
				"name": "enum",
				"values": [
					{
						"key": "test",
						"label": "test"
					}
				]
			}`,
		},
		{
			desc: "Localized enum type",
			object: commercetools.AttributeLocalizedEnumType{
				Values: []commercetools.AttributeLocalizedEnumValue{
					{
						Key: "test",
						Label: &commercetools.LocalizedString{
							"en": "test",
						},
					},
				},
			},
			json: `{
				"name": "lenum",
				"values": [
					{
						"key": "test",
						"label": {
							"en": "test"
						}
					}
				]
			}`,
		},
		{
			desc:   "Number type",
			object: commercetools.AttributeNumberType{},
			json: `{
				"name": "number"
			}`,
		},
		{
			desc:   "Money type",
			object: commercetools.AttributeMoneyType{},
			json: `{
				"name": "money"
			}`,
		},
		{
			desc:   "Date type",
			object: commercetools.AttributeDateType{},
			json: `{
				"name": "date"
			}`,
		},
		{
			desc:   "Time type",
			object: commercetools.AttributeTimeType{},
			json: `{
				"name": "time"
			}`,
		},
		{
			desc:   "Date time type",
			object: commercetools.AttributeDateTimeType{},
			json: `{
				"name": "datetime"
			}`,
		},
		{
			desc: "Reference type",
			object: commercetools.AttributeReferenceType{
				ReferenceTypeID: "product",
			},
			json: `{
				"name": "reference",
				"referenceTypeId": "product"
			}`,
		},
		{
			desc: "Set type",
			object: commercetools.AttributeSetType{
				ElementType: commercetools.AttributeBooleanType{},
			},
			json: `{
				"name": "set",
				"elementType": {
					"name": "boolean"
				}
			}`,
		},
		{
			desc: "Nested type",
			object: commercetools.AttributeNestedType{
				TypeReference: &commercetools.ProductTypeReference{
					ID: "b2ad48b7-c65b-4743-810c-dc898bf934cc",
				},
			},
			json: `{
				"name": "nested",
				"typeReference": {
					"id": "b2ad48b7-c65b-4743-810c-dc898bf934cc",
					"typeId": "product-type"
				}
			}`,
		},
	}

	for _, tC := range testCases {
		// Validate Marshaling (object -> json)
		t.Run(tC.desc, func(t *testing.T) {
			output, err := tC.object.MarshalJSON()
			assert.Nil(t, err)
			assert.JSONEq(t, tC.json, string(output))
		})

		// Validate Umarshalling (json -> object)
		t.Run(tC.desc, func(t *testing.T) {
			buf := fmt.Sprintf(`
			{
				"name": "test",
				"label": {
					"en": "test"
				},
				"isRequired": false,
				"attributeConstraint": "None",
				"type": %s,
				"inputTip": {
					"en": "test"
				},
				"inputHint": "SingleLine",
				"isSearchable": true
			}`, tC.json)

			ad := commercetools.AttributeDefinition{}

			err := json.Unmarshal([]byte(buf), &ad)
			assert.Nil(t, err)
			assert.Equal(t, tC.object, ad.Type)
		})
	}
}
