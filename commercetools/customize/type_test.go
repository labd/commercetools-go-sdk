package customize_test

import (
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/commercetools/customize"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
)

func TestTypeCreate(t *testing.T) {
	output := testutil.RequestData{}

	client, server := testutil.MockClient(t, fixture("type.create.json"), &output, nil)
	defer server.Close()
	svc := customize.New(client)

	input := &customize.TypeDraft{
		Key: "lineitemtype",
		Name: commercetools.LocalizedString{
			"en": "lineitem",
		},
		Description: commercetools.LocalizedString{
			"en": "description",
		},
		ResourceTypeIds: []string{"line-item"},
		FieldDefinitions: []customize.FieldDefinition{
			customize.FieldDefinition{
				Type: customize.StringType{},
				Name: "offer_name",
				Label: commercetools.LocalizedString{
					"en": "offer_name",
				},
				Required:  false,
				InputHint: customize.SingleLineTextInputHint,
			},
		},
	}

	fmt.Println(output)

	_, err := svc.TypeCreate(input)
	assert.Equal(t, nil, err)

	expectedBody := `{
		"key": "lineitemtype",
		"name": {
			"en": "lineitem"
		},
		"description": {
			"en": "description"
		},
		"resourceTypeIds": [
			"line-item"
		],
		"fieldDefinitions": [
			{
				"name": "offer_name",
				"label": {
					"en": "offer_name"
				},
				"required": false,
				"type": {
					"name": "String"
				},
				"inputHint": "SingleLine"
			}
		]
	}`

	assert.JSONEq(t, expectedBody, output.JSON)
}

func TestTypeUpdate(t *testing.T) {
	testCases := []struct {
		desc        string
		input       *customize.TypeUpdateInput
		requestBody string
	}{
		{
			desc: "Change key",
			input: &customize.TypeUpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&customize.TypeChangeKey{
						Key: "123456",
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "changeKey",
						"key": "123456"
					}
				]
			}`,
		},
		{
			desc: "Change name",
			input: &customize.TypeUpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&customize.TypeChangeName{
						Name: commercetools.LocalizedString{
							"en": "Product",
						},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "changeName",
						"name": {
							"en": "Product"
						}
					}
				]
			}`,
		},
		{
			desc: "Set description",
			input: &customize.TypeUpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&customize.TypeSetDescription{
						Description: commercetools.LocalizedString{
							"en": "Description",
						},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "setDescription",
						"description": {
							"en": "Description"
						}
					}
				]
			}`,
		},
		{
			desc: "Add field definition",
			input: &customize.TypeUpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&customize.TypeAddFieldDefinition{
						FieldDefinition: customize.FieldDefinition{
							Type: customize.BooleanType{},
							Name: "custom-field",
							Label: commercetools.LocalizedString{
								"en": "custom-field",
							},
							Required:  true,
							InputHint: customize.MultiLineTextInputHint,
						},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "addFieldDefinition",
						"fieldDefinition": {
							"type": {
								"name": "Boolean"
							},
							"name": "custom-field",
							"label": {
								"en": "custom-field"
							},
							"required": true,
							"inputHint": "MultiLine"
						}
					}
				]
			}`,
		},
		{
			desc: "Remove field definition",
			input: &customize.TypeUpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&customize.TypeRemoveFieldDefinition{
						FieldName: "custom-field",
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "removeFieldDefinition",
						"fieldName": "custom-field"
					}
				]
			}`,
		},
		{
			desc: "Change label",
			input: &customize.TypeUpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&customize.TypeChangeLabel{
						FieldName: "custom-field",
						Label: commercetools.LocalizedString{
							"en": "custom-field",
						},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "changeLabel",
						"fieldName": "custom-field",
						"label": {
							"en": "custom-field"
						}
					}
				]
			}`,
		},
		{
			desc: "Add enum value",
			input: &customize.TypeUpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&customize.TypeAddEnumValue{
						FieldName: "custom-field",
						Value: customize.EnumValue{
							Key:   "custom-enum",
							Label: "enum-label",
						},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "addEnumValue",
						"fieldName": "custom-field",
						"value": {
							"key": "custom-enum",
							"label": "enum-label"
						}
					}
				]
			}`,
		},
		{
			desc: "Add localized enum value",
			input: &customize.TypeUpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&customize.TypeAddLocalizedEnumValue{
						FieldName: "custom-field",
						Value: customize.LocalizedEnumValue{
							Key: "custom-enum",
							Label: commercetools.LocalizedString{
								"en": "custom-label",
							},
						},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "addLocalizedEnumValue",
						"fieldName": "custom-field",
						"value": {
							"key": "custom-enum",
							"label": {
								"en": "custom-label"
							}
						}
					}
				]
			}`,
		},
		{
			desc: "Change field definition order",
			input: &customize.TypeUpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&customize.TypeChangeFieldDefinitionsOrder{
						FieldNames: []string{"field-1", "field-2"},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "changeFieldDefinitionOrder",
						"fieldNames": ["field-1", "field-2"]
					}
				]
			}`,
		},
		{
			desc: "Change enum value order",
			input: &customize.TypeUpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&customize.TypeChangeEnumValuesOrder{
						FieldName: "custom-field",
						Keys:      []string{"key-1", "key-2"},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "changeEnumValueOrder",
						"fieldName": "custom-field",
						"keys": ["key-1", "key-2"]
					}
				]
			}`,
		},
		{
			desc: "Change localized enum value order",
			input: &customize.TypeUpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&customize.TypeChangeLocalizedEnumValuesOrder{
						FieldName: "custom-field",
						Keys:      []string{"key-1", "key-2"},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "changeLocalizedEnumValueOrder",
						"fieldName": "custom-field",
						"keys": ["key-1", "key-2"]
					}
				]
			}`,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			output := testutil.RequestData{}

			client, server := testutil.MockClient(t, fixture("type.update.json"), &output, nil)
			defer server.Close()
			svc := customize.New(client)

			_, err := svc.TypeUpdate(tC.input)
			assert.Equal(t, nil, err)
			assert.Equal(t, "/unittest/types/1234", output.URL.Path)
			assert.JSONEq(t, tC.requestBody, output.JSON)
		})
	}
}

func TestTypeDeleteByID(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, fixture("type.json"), &output, nil)
	defer server.Close()
	svc := customize.New(client)

	_, err := svc.TypeDeleteByID("1234", 2)
	assert.Equal(t, nil, err)

	params := url.Values{}
	params.Add("version", "2")
	assert.Equal(t, params, output.URL.Query())
	assert.Equal(t, "/unittest/types/1234", output.URL.Path)
}

func TestTypeDeleteByKey(t *testing.T) {
	output := testutil.RequestData{}

	client, server := testutil.MockClient(t, fixture("type.json"), &output, nil)
	defer server.Close()
	svc := customize.New(client)

	_, err := svc.TypeDeleteByKey("1234", 2)
	assert.Equal(t, nil, err)

	params := url.Values{}
	params.Add("version", "2")
	assert.Equal(t, params, output.URL.Query())
	assert.Equal(t, "/unittest/types/key=1234", output.URL.Path)
}

func TestTypeTypeGetByID(t *testing.T) {
	output := testutil.RequestData{}

	client, server := testutil.MockClient(t, fixture("type.json"), &output, nil)
	defer server.Close()
	svc := customize.New(client)

	result, err := svc.TypeGetByID("1234")
	assert.Equal(t, nil, err)

	timestamp, _ := time.Parse(time.RFC3339, "2015-10-07T06:56:19.217Z")

	expected := &customize.Type{
		ID:      "1234",
		Version: 1,
		Key:     "lineitemtype",
		Name: commercetools.LocalizedString{
			"en": "lineitem",
		},
		Description: commercetools.LocalizedString{
			"en": "description",
		},
		ResourceTypeIds: []string{"line-item"},
		FieldDefinitions: []customize.FieldDefinition{
			customize.FieldDefinition{
				Type: customize.StringType{},
				Name: "offer_name",
				Label: commercetools.LocalizedString{
					"en": "offer_name",
				},
				Required:  false,
				InputHint: customize.SingleLineTextInputHint,
			},
		},
		CreatedAt:      timestamp,
		LastModifiedAt: timestamp,
	}
	assert.Equal(t, result, expected)
}

func TestFieldTypes(t *testing.T) {
	testCases := []struct {
		desc  string
		input interface {
			MarshalJSON() ([]byte, error)
		}
		output string
	}{
		{
			desc:  "Boolean type",
			input: customize.BooleanType{},
			output: `{
				"name": "Boolean"
			}`,
		},
		{
			desc:  "String type",
			input: customize.StringType{},
			output: `{
				"name": "String"
			}`,
		},
		{
			desc:  "Localized string type",
			input: customize.LocalizedStringType{},
			output: `{
				"name": "LocalizedString"
			}`,
		},
		{
			desc: "Enum type",
			input: customize.EnumType{
				Values: []customize.EnumValue{
					customize.EnumValue{
						Key:   "test",
						Label: "test",
					},
				},
			},
			output: `{
				"name": "Enum",
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
			input: customize.LocalizedEnumType{
				Values: []customize.LocalizedEnumValue{
					customize.LocalizedEnumValue{
						Key: "test",
						Label: commercetools.LocalizedString{
							"en": "test",
						},
					},
				},
			},
			output: `{
				"name": "LocalizedEnum",
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
			desc:  "Number type",
			input: customize.NumberType{},
			output: `{
				"name": "Number"
			}`,
		},
		{
			desc:  "Money type",
			input: customize.MoneyType{},
			output: `{
				"name": "Money"
			}`,
		},
		{
			desc:  "Date type",
			input: customize.DateType{},
			output: `{
				"name": "Date"
			}`,
		},
		{
			desc:  "Time type",
			input: customize.TimeType{},
			output: `{
				"name": "Time"
			}`,
		},
		{
			desc:  "Date time type",
			input: customize.DateTimeType{},
			output: `{
				"name": "DateTime"
			}`,
		},
		{
			desc: "Reference type",
			input: customize.ReferenceType{
				ReferenceTypeId: "product",
			},
			output: `{
				"name": "Reference",
				"referenceTypeId": "product"
			}`,
		},
		{
			desc: "Set type",
			input: customize.SetType{
				ElementType: customize.BooleanType{},
			},
			output: `{
				"name": "Set",
				"elementType": {
					"name": "Boolean"
				}
			}`,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			output, err := tC.input.MarshalJSON()
			assert.Equal(t, nil, err)
			assert.JSONEq(t, tC.output, string(output))
		})
	}
}
