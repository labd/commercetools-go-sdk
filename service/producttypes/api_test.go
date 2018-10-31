package producttypes_test

import (
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/stretchr/testify/assert"

	"github.com/labd/commercetools-go-sdk/service/producttypes"
	"github.com/labd/commercetools-go-sdk/testutil"
)

func TestProductTypeCreate(t *testing.T) {
	output := testutil.RequestData{}

	client, server := testutil.MockClient(t, testutil.Fixture("producttype.create.json"), &output, nil)
	defer server.Close()
	svc := producttypes.New(client)

	input := &producttypes.ProductTypeDraft{
		Name:        "test_product_type",
		Description: "Test product type.",
		Attributes: []producttypes.AttributeDefinitionDraft{
			{
				Type:         producttypes.TextType{},
				IsSearchable: false,
				InputHint:    commercetools.SingleLineTextInputHint,
				Name:         "size",
				Label: commercetools.LocalizedString{
					"en": "The right size is important.",
				},
				IsRequired:          false,
				AttributeConstraint: producttypes.CombinationUniqueAttributeConstraint,
			},
		},
	}

	fmt.Println(output)

	_, err := svc.Create(input)
	assert.Nil(t, err)

	expectedBody := `{
		"name": "test_product_type",
		"description": "Test product type.",
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
	}`
	assert.JSONEq(t, expectedBody, output.JSON)
}

func TestProductTypeUpdate(t *testing.T) {
	testCases := []struct {
		desc        string
		input       *producttypes.UpdateInput
		requestBody string
	}{
		{
			desc: "Set key",
			input: &producttypes.UpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: commercetools.UpdateActions{
					&producttypes.SetKey{
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
			input: &producttypes.UpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: commercetools.UpdateActions{
					&producttypes.ChangeName{
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
			input: &producttypes.UpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: commercetools.UpdateActions{
					&producttypes.ChangeDescription{
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
			input: &producttypes.UpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: commercetools.UpdateActions{
					&producttypes.AddAttributeDefinition{
						Attribute: producttypes.AttributeDefinitionDraft{
							Type: producttypes.BooleanType{},
							Name: "attribute",
							Label: commercetools.LocalizedString{
								"en": "attribute",
							},
							IsRequired:          false,
							AttributeConstraint: producttypes.NoneAttributeConstraint,
							InputTip: commercetools.LocalizedString{
								"en": "A simple hint",
							},
							InputHint:    commercetools.SingleLineTextInputHint,
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
			input: &producttypes.UpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: commercetools.UpdateActions{
					&producttypes.RemoveAttributeDefinition{
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
			input: &producttypes.UpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: commercetools.UpdateActions{
					&producttypes.ChangeAttributeDefinitionName{
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
			input: &producttypes.UpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: commercetools.UpdateActions{
					&producttypes.ChangeAttributeDefinitionLabel{
						AttributeName: "attribute",
						Label: commercetools.LocalizedString{
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
			input: &producttypes.UpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: commercetools.UpdateActions{
					&producttypes.SetAttributeDefinitionInputTip{
						AttributeName: "attribute",
						InputTip: commercetools.LocalizedString{
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
			input: &producttypes.UpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: commercetools.UpdateActions{
					&producttypes.AddPlainEnumValue{
						AttributeName: "attribute",
						Value: commercetools.EnumValue{
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
			input: &producttypes.UpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: commercetools.UpdateActions{
					&producttypes.AddLocalizedEnumValue{
						AttributeName: "attribute",
						Value: commercetools.LocalizedEnumValue{
							Key: "enum",
							Label: commercetools.LocalizedString{
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
			input: &producttypes.UpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: commercetools.UpdateActions{
					&producttypes.RemoveEnumValues{
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
			input: &producttypes.UpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: commercetools.UpdateActions{
					&producttypes.ChangeAttributeDefinitionsOrder{
						Attributes: []producttypes.AttributeDefinition{
							{
								Type:         producttypes.TextType{},
								IsSearchable: false,
								InputHint:    commercetools.SingleLineTextInputHint,
								Name:         "size",
								Label: commercetools.LocalizedString{
									"en": "The right size is important.",
								},
								IsRequired:          false,
								AttributeConstraint: producttypes.CombinationUniqueAttributeConstraint,
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
			input: &producttypes.UpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: commercetools.UpdateActions{
					&producttypes.ChangePlainEnumValuesOrder{
						AttributeName: "attribute",
						Values: []commercetools.EnumValue{
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
			input: &producttypes.UpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: commercetools.UpdateActions{
					&producttypes.ChangeLocalizedEnumValuesOrder{
						AttributeName: "attribute",
						Values: []commercetools.LocalizedEnumValue{
							{
								Key: "enum",
								Label: commercetools.LocalizedString{
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
			input: &producttypes.UpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: commercetools.UpdateActions{
					&producttypes.ChangePlainEnumValueKey{
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
			input: &producttypes.UpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: commercetools.UpdateActions{
					&producttypes.ChangePlainEnumValueLabel{
						AttributeName: "attribute",
						NewValue: commercetools.EnumValue{
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
			input: &producttypes.UpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: commercetools.UpdateActions{
					&producttypes.ChangeLocalizedEnumValueLabel{
						AttributeName: "attribute",
						NewValue: commercetools.LocalizedEnumValue{
							Key: "enum",
							Label: commercetools.LocalizedString{
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
			input: &producttypes.UpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: commercetools.UpdateActions{
					&producttypes.ChangeIsSearchable{
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
			input: &producttypes.UpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: commercetools.UpdateActions{
					&producttypes.ChangeInputHint{
						AttributeName: "attribute",
						NewValue:      commercetools.SingleLineTextInputHint,
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
			input: &producttypes.UpdateInput{
				ID:      "1234",
				Version: 3,
				Actions: commercetools.UpdateActions{
					&producttypes.ChangeAttributeConstraint{
						AttributeName: "attribute",
						NewValue:      producttypes.CombinationUniqueAttributeConstraint,
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
			svc := producttypes.New(client)

			_, err := svc.Update(tC.input)
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
	svc := producttypes.New(client)

	_, err := svc.DeleteByID("1234", 3)
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
	svc := producttypes.New(client)

	_, err := svc.DeleteByKey("1234", 3)
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
		input   *producttypes.ProductType
		fixture string
	}{
		{
			desc: "Type with size text attribute definition",
			input: &producttypes.ProductType{
				ID:          "1234",
				Version:     1,
				Name:        "test_product_type",
				Description: "Test product type.",
				Attributes: []producttypes.AttributeDefinition{
					{
						Type:         producttypes.TextType{},
						IsSearchable: false,
						InputHint:    commercetools.SingleLineTextInputHint,
						Name:         "size",
						Label: commercetools.LocalizedString{
							"en": "The right size is important.",
						},
						IsRequired:          false,
						AttributeConstraint: producttypes.CombinationUniqueAttributeConstraint,
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "producttype.size.json",
		},
		{
			desc: "Type with color text attribute definition",
			input: &producttypes.ProductType{
				ID:          "1234",
				Version:     1,
				Name:        "another-test_product_type",
				Description: "Another test product type.",
				Attributes: []producttypes.AttributeDefinition{
					{
						Type:         producttypes.LocalizedTextType{},
						IsSearchable: false,
						InputHint:    commercetools.SingleLineTextInputHint,
						Name:         "color",
						Label: commercetools.LocalizedString{
							"en": "The right color is important.",
						},
						IsRequired:          false,
						AttributeConstraint: producttypes.CombinationUniqueAttributeConstraint,
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

			svc := producttypes.New(client)
			result, err := svc.GetByID("1234")
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
			object: producttypes.BooleanType{},
			json: `{
				"name": "boolean"
			}`,
		},
		{
			desc:   "Text type",
			object: producttypes.TextType{},
			json: `{
				"name": "text"
			}`,
		},
		{
			desc:   "Localized text type",
			object: producttypes.LocalizedTextType{},
			json: `{
				"name": "ltext"
			}`,
		},
		{
			desc: "Enum type",
			object: producttypes.EnumType{
				Values: []commercetools.EnumValue{
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
			object: producttypes.LocalizedEnumType{
				Values: []commercetools.LocalizedEnumValue{
					{
						Key: "test",
						Label: commercetools.LocalizedString{
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
			object: producttypes.NumberType{},
			json: `{
				"name": "number"
			}`,
		},
		{
			desc:   "Money type",
			object: producttypes.MoneyType{},
			json: `{
				"name": "money"
			}`,
		},
		{
			desc:   "Date type",
			object: producttypes.DateType{},
			json: `{
				"name": "date"
			}`,
		},
		{
			desc:   "Time type",
			object: producttypes.TimeType{},
			json: `{
				"name": "time"
			}`,
		},
		{
			desc:   "Date time type",
			object: producttypes.DateTimeType{},
			json: `{
				"name": "datetime"
			}`,
		},
		{
			desc: "Reference type",
			object: producttypes.ReferenceType{
				ReferenceTypeID: "product",
			},
			json: `{
				"name": "reference",
				"referenceTypeId": "product"
			}`,
		},
		{
			desc: "Set type",
			object: producttypes.SetType{
				ElementType: producttypes.BooleanType{},
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
			object: producttypes.NestedType{
				TypeReference: commercetools.Reference{ID: "b2ad48b7-c65b-4743-810c-dc898bf934cc", TypeID: "product-type"},
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

			ad := producttypes.AttributeDefinition{}

			err := ad.UnmarshalJSON([]byte(buf))
			assert.Nil(t, err)
			assert.Equal(t, tC.object, ad.Type)
		})
	}
}
