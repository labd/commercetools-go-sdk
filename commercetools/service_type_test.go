package commercetools_test

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
)

func TestTypeCreate(t *testing.T) {
	output := testutil.RequestData{}

	client, server := testutil.MockClient(t, testutil.Fixture("type.create.json"), &output, nil)
	defer server.Close()

	input := &commercetools.TypeDraft{
		Key: "lineitemtype",
		Name: &commercetools.LocalizedString{
			"en": "lineitem",
		},
		Description: &commercetools.LocalizedString{
			"en": "description",
		},
		ResourceTypeIds: []commercetools.ResourceTypeID{"line-item"},
		FieldDefinitions: []commercetools.FieldDefinition{
			{
				Type: commercetools.CustomFieldStringType{},
				Name: "offer_name",
				Label: &commercetools.LocalizedString{
					"en": "offer_name",
				},
				Required:  false,
				InputHint: "SingleLine",
			},
		},
	}

	fmt.Println(output)

	_, err := client.Types.Create(input)
	assert.Nil(t, err)

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
		input       *commercetools.TypeUpdateInput
		requestBody string
	}{
		{
			desc: "Change key",
			input: &commercetools.TypeUpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.TypeUpdateAction{
					&commercetools.TypeChangeKeyAction{
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
			input: &commercetools.TypeUpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.TypeUpdateAction{
					&commercetools.TypeChangeNameAction{
						Name: &commercetools.LocalizedString{
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
			input: &commercetools.TypeUpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.TypeUpdateAction{
					&commercetools.TypeSetDescriptionAction{
						Description: &commercetools.LocalizedString{
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
			input: &commercetools.TypeUpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.TypeUpdateAction{
					&commercetools.TypeAddFieldDefinitionAction{
						FieldDefinition: &commercetools.FieldDefinition{
							Type: commercetools.CustomFieldBooleanType{},
							Name: "custom-field",
							Label: &commercetools.LocalizedString{
								"en": "custom-field",
							},
							Required:  true,
							InputHint: "MultiLine",
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
			input: &commercetools.TypeUpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.TypeUpdateAction{
					&commercetools.TypeRemoveFieldDefinitionAction{
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
			input: &commercetools.TypeUpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.TypeUpdateAction{
					&commercetools.TypeChangeLabelAction{
						FieldName: "custom-field",
						Label: &commercetools.LocalizedString{
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
			input: &commercetools.TypeUpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.TypeUpdateAction{
					&commercetools.TypeAddEnumValueAction{
						FieldName: "custom-field",
						Value: &commercetools.CustomFieldEnumValue{
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
			input: &commercetools.TypeUpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.TypeUpdateAction{
					&commercetools.TypeAddLocalizedEnumValueAction{
						FieldName: "custom-field",
						Value: &commercetools.CustomFieldLocalizedEnumValue{
							Key: "custom-enum",
							Label: &commercetools.LocalizedString{
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
			input: &commercetools.TypeUpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.TypeUpdateAction{
					&commercetools.TypeChangeFieldDefinitionOrderAction{
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
			input: &commercetools.TypeUpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.TypeUpdateAction{
					&commercetools.TypeChangeEnumValueOrderAction{
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
			input: &commercetools.TypeUpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.TypeUpdateAction{
					&commercetools.TypeChangeLocalizedEnumValueOrderAction{
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

			client, server := testutil.MockClient(t, testutil.Fixture("type.update.json"), &output, nil)
			defer server.Close()

			_, err := client.Types.Update(tC.input)
			assert.Nil(t, err)
			assert.Equal(t, "/unittest/types/1234", output.URL.Path)
			assert.JSONEq(t, tC.requestBody, output.JSON)
		})
	}
}

func TestTypeDeleteByID(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, testutil.Fixture("type.boolean.json"), &output, nil)
	defer server.Close()

	_, err := client.Types.DeleteByID("1234", 2)
	assert.Nil(t, err)

	params := url.Values{}
	params.Add("version", "2")
	assert.Equal(t, params, output.URL.Query())
	assert.Equal(t, "/unittest/types/1234", output.URL.Path)
}

func TestTypeDeleteByKey(t *testing.T) {
	output := testutil.RequestData{}

	client, server := testutil.MockClient(t, testutil.Fixture("type.boolean.json"), &output, nil)
	defer server.Close()

	_, err := client.Types.DeleteByKey("1234", 2)
	assert.Nil(t, err)

	params := url.Values{}
	params.Add("version", "2")
	assert.Equal(t, params, output.URL.Query())
	assert.Equal(t, "/unittest/types/key=1234", output.URL.Path)
}

func TestTypeGetByID(t *testing.T) {
	timestamp, _ := time.Parse(time.RFC3339, "2015-10-07T06:56:19.217Z")

	testCases := []struct {
		desc    string
		input   *commercetools.Type
		fixture string
	}{
		{
			desc: "Type with boolean field definition",
			input: &commercetools.Type{
				ID:      "1234",
				Version: 1,
				Key:     "lineitemtype",
				Name: &commercetools.LocalizedString{
					"en": "lineitem",
				},
				Description: &commercetools.LocalizedString{
					"en": "description",
				},
				ResourceTypeIds: []commercetools.ResourceTypeID{"line-item"},
				FieldDefinitions: []commercetools.FieldDefinition{
					{
						Type: commercetools.CustomFieldBooleanType{},
						Name: "is_special",
						Label: &commercetools.LocalizedString{
							"en": "is_special",
						},
						Required:  false,
						InputHint: "SingleLine",
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.boolean.json",
		},
		{
			desc: "Type with string field definition",
			input: &commercetools.Type{
				ID:      "1234",
				Version: 1,
				Key:     "lineitemtype",
				Name: &commercetools.LocalizedString{
					"en": "lineitem",
				},
				Description: &commercetools.LocalizedString{
					"en": "description",
				},
				ResourceTypeIds: []commercetools.ResourceTypeID{"line-item"},
				FieldDefinitions: []commercetools.FieldDefinition{
					{
						Type: commercetools.CustomFieldStringType{},
						Name: "offer_name",
						Label: &commercetools.LocalizedString{
							"en": "offer_name",
						},
						Required:  false,
						InputHint: "SingleLine",
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.string.json",
		},
		{
			desc: "Type with localized string field definition",
			input: &commercetools.Type{
				ID:      "1234",
				Version: 1,
				Key:     "lineitemtype",
				Name: &commercetools.LocalizedString{
					"en": "lineitem",
				},
				Description: &commercetools.LocalizedString{
					"en": "description",
				},
				ResourceTypeIds: []commercetools.ResourceTypeID{"line-item"},
				FieldDefinitions: []commercetools.FieldDefinition{
					{
						Type: commercetools.CustomFieldLocalizedStringType{},
						Name: "translated_offer_name",
						Label: &commercetools.LocalizedString{
							"en": "translated_offer_name",
						},
						Required:  false,
						InputHint: "SingleLine",
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.lstring.json",
		},
		{
			desc: "Type with enum field definition",
			input: &commercetools.Type{
				ID:      "1234",
				Version: 1,
				Key:     "lineitemtype",
				Name: &commercetools.LocalizedString{
					"en": "lineitem",
				},
				Description: &commercetools.LocalizedString{
					"en": "description",
				},
				ResourceTypeIds: []commercetools.ResourceTypeID{"line-item"},
				FieldDefinitions: []commercetools.FieldDefinition{
					{
						Type: commercetools.CustomFieldEnumType{
							Values: []commercetools.CustomFieldEnumValue{
								{
									Key:   "enum",
									Label: "enum",
								},
							},
						},
						Name: "offer_enum",
						Label: &commercetools.LocalizedString{
							"en": "offer_enum",
						},
						Required:  false,
						InputHint: "SingleLine",
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.enum.json",
		},
		{
			desc: "Type with localized enum field definition",
			input: &commercetools.Type{
				ID:      "1234",
				Version: 1,
				Key:     "lineitemtype",
				Name: &commercetools.LocalizedString{
					"en": "lineitem",
				},
				Description: &commercetools.LocalizedString{
					"en": "description",
				},
				ResourceTypeIds: []commercetools.ResourceTypeID{"line-item"},
				FieldDefinitions: []commercetools.FieldDefinition{
					{
						Type: commercetools.CustomFieldLocalizedEnumType{
							Values: []commercetools.CustomFieldLocalizedEnumValue{
								{
									Key: "enum",
									Label: &commercetools.LocalizedString{
										"en": "enum",
									},
								},
							},
						},
						Name: "translated_offer_enum",
						Label: &commercetools.LocalizedString{
							"en": "translated_offer_enum",
						},
						Required:  false,
						InputHint: "SingleLine",
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.lenum.json",
		},
		{
			desc: "Type with number field definition",
			input: &commercetools.Type{
				ID:      "1234",
				Version: 1,
				Key:     "lineitemtype",
				Name: &commercetools.LocalizedString{
					"en": "lineitem",
				},
				Description: &commercetools.LocalizedString{
					"en": "description",
				},
				ResourceTypeIds: []commercetools.ResourceTypeID{"line-item"},
				FieldDefinitions: []commercetools.FieldDefinition{
					{
						Type: commercetools.CustomFieldNumberType{},
						Name: "offer_number",
						Label: &commercetools.LocalizedString{
							"en": "offer_number",
						},
						Required:  false,
						InputHint: "SingleLine",
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.number.json",
		},
		{
			desc: "Type with money field definition",
			input: &commercetools.Type{
				ID:      "1234",
				Version: 1,
				Key:     "lineitemtype",
				Name: &commercetools.LocalizedString{
					"en": "lineitem",
				},
				Description: &commercetools.LocalizedString{
					"en": "description",
				},
				ResourceTypeIds: []commercetools.ResourceTypeID{"line-item"},
				FieldDefinitions: []commercetools.FieldDefinition{
					{
						Type: commercetools.CustomFieldMoneyType{},
						Name: "offer_price",
						Label: &commercetools.LocalizedString{
							"en": "offer_price",
						},
						Required:  false,
						InputHint: "SingleLine",
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.money.json",
		},
		{
			desc: "Type with date field definition",
			input: &commercetools.Type{
				ID:      "1234",
				Version: 1,
				Key:     "lineitemtype",
				Name: &commercetools.LocalizedString{
					"en": "lineitem",
				},
				Description: &commercetools.LocalizedString{
					"en": "description",
				},
				ResourceTypeIds: []commercetools.ResourceTypeID{"line-item"},
				FieldDefinitions: []commercetools.FieldDefinition{
					{
						Type: commercetools.CustomFieldDateType{},
						Name: "offer_date",
						Label: &commercetools.LocalizedString{
							"en": "offer_date",
						},
						Required:  false,
						InputHint: "SingleLine",
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.date.json",
		},
		{
			desc: "Type with time field definition",
			input: &commercetools.Type{
				ID:      "1234",
				Version: 1,
				Key:     "lineitemtype",
				Name: &commercetools.LocalizedString{
					"en": "lineitem",
				},
				Description: &commercetools.LocalizedString{
					"en": "description",
				},
				ResourceTypeIds: []commercetools.ResourceTypeID{"line-item"},
				FieldDefinitions: []commercetools.FieldDefinition{
					{
						Type: commercetools.CustomFieldTimeType{},
						Name: "offer_time",
						Label: &commercetools.LocalizedString{
							"en": "offer_time",
						},
						Required:  false,
						InputHint: "SingleLine",
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.time.json",
		},
		{
			desc: "Type with date-time field definition",
			input: &commercetools.Type{
				ID:      "1234",
				Version: 1,
				Key:     "lineitemtype",
				Name: &commercetools.LocalizedString{
					"en": "lineitem",
				},
				Description: &commercetools.LocalizedString{
					"en": "description",
				},
				ResourceTypeIds: []commercetools.ResourceTypeID{"line-item"},
				FieldDefinitions: []commercetools.FieldDefinition{
					{
						Type: commercetools.CustomFieldDateTimeType{},
						Name: "offer_date_time",
						Label: &commercetools.LocalizedString{
							"en": "offer_date_time",
						},
						Required:  false,
						InputHint: "SingleLine",
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.datetime.json",
		},
		{
			desc: "Type with reference field definition",
			input: &commercetools.Type{
				ID:      "1234",
				Version: 1,
				Key:     "lineitemtype",
				Name: &commercetools.LocalizedString{
					"en": "lineitem",
				},
				Description: &commercetools.LocalizedString{
					"en": "description",
				},
				ResourceTypeIds: []commercetools.ResourceTypeID{"line-item"},
				FieldDefinitions: []commercetools.FieldDefinition{
					{
						Type: commercetools.CustomFieldReferenceType{
							ReferenceTypeID: "product",
						},
						Name: "offer_reference",
						Label: &commercetools.LocalizedString{
							"en": "offer_reference",
						},
						Required:  false,
						InputHint: "SingleLine",
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.reference.json",
		},
		{
			desc: "Type with set field definition",
			input: &commercetools.Type{
				ID:      "1234",
				Version: 1,
				Key:     "lineitemtype",
				Name: &commercetools.LocalizedString{
					"en": "lineitem",
				},
				Description: &commercetools.LocalizedString{
					"en": "description",
				},
				ResourceTypeIds: []commercetools.ResourceTypeID{"line-item"},
				FieldDefinitions: []commercetools.FieldDefinition{
					{
						Type: commercetools.CustomFieldSetType{
							ElementType: commercetools.CustomFieldStringType{},
						},
						Name: "offer_set",
						Label: &commercetools.LocalizedString{
							"en": "offer_set",
						},
						Required:  false,
						InputHint: "SingleLine",
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.set.json",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			client, server := testutil.MockClient(t, testutil.Fixture(tC.fixture), nil, nil)
			defer server.Close()

			result, err := client.Types.GetByID("1234")
			assert.Nil(t, err)
			assert.Equal(t, tC.input, result)
		})
	}
}

func TestFieldTypes(t *testing.T) {
	testCases := []struct {
		desc   string
		object interface {
			MarshalJSON() ([]byte, error)
		}
		json string
	}{
		{
			desc:   "Boolean type",
			object: commercetools.CustomFieldBooleanType{},
			json: `{
				"name": "Boolean"
			}`,
		},
		{
			desc:   "String type",
			object: commercetools.CustomFieldStringType{},
			json: `{
				"name": "String"
			}`,
		},
		{
			desc:   "Localized string type",
			object: commercetools.CustomFieldLocalizedStringType{},
			json: `{
				"name": "LocalizedString"
			}`,
		},
		{
			desc: "Enum type",
			object: commercetools.CustomFieldEnumType{
				Values: []commercetools.CustomFieldEnumValue{
					{
						Key:   "test",
						Label: "test",
					},
				},
			},
			json: `{
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
			object: commercetools.CustomFieldLocalizedEnumType{
				Values: []commercetools.CustomFieldLocalizedEnumValue{
					{
						Key: "test",
						Label: &commercetools.LocalizedString{
							"en": "test",
						},
					},
				},
			},
			json: `{
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
			desc:   "Number type",
			object: commercetools.CustomFieldNumberType{},
			json: `{
				"name": "Number"
			}`,
		},
		{
			desc:   "Money type",
			object: commercetools.CustomFieldMoneyType{},
			json: `{
				"name": "Money"
			}`,
		},
		{
			desc:   "Date type",
			object: commercetools.CustomFieldDateType{},
			json: `{
				"name": "Date"
			}`,
		},
		{
			desc:   "Time type",
			object: commercetools.CustomFieldTimeType{},
			json: `{
				"name": "Time"
			}`,
		},
		{
			desc:   "Date time type",
			object: commercetools.CustomFieldDateTimeType{},
			json: `{
				"name": "DateTime"
			}`,
		},
		{
			desc: "Reference type",
			object: commercetools.CustomFieldReferenceType{
				ReferenceTypeID: "product",
			},
			json: `{
				"name": "Reference",
				"referenceTypeId": "product"
			}`,
		},
		{
			desc: "Set type",
			object: commercetools.CustomFieldSetType{
				ElementType: commercetools.CustomFieldBooleanType{},
			},
			json: `{
				"name": "Set",
				"elementType": {
					"name": "Boolean"
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
				"required": false,
				"type": %s,
				"inputHint": "SingleLine"
			}`, tC.json)

			fd := commercetools.FieldDefinition{}
			err := json.Unmarshal([]byte(buf), &fd)
			assert.Nil(t, err)
			assert.Equal(t, tC.object, fd.Type)
		})
	}
}
