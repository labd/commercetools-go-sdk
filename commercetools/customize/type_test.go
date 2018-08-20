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
			{
				Type: customize.StringType{},
				Name: "offer_name%",
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
				"name": "offer_name%",
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
	client, server := testutil.MockClient(t, fixture("type.boolean.json"), &output, nil)
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

	client, server := testutil.MockClient(t, fixture("type.boolean.json"), &output, nil)
	defer server.Close()
	svc := customize.New(client)

	_, err := svc.TypeDeleteByKey("1234", 2)
	assert.Equal(t, nil, err)

	params := url.Values{}
	params.Add("version", "2")
	assert.Equal(t, params, output.URL.Query())
	assert.Equal(t, "/unittest/types/key=1234", output.URL.Path)
}

func TestTypeGetByID(t *testing.T) {
	timestamp, _ := time.Parse(time.RFC3339, "2015-10-07T06:56:19.217Z")

	testCases := []struct {
		desc    string
		input   *customize.Type
		fixture string
	}{
		{
			desc: "Type with boolean field definition",
			input: &customize.Type{
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
					{
						Type: customize.BooleanType{},
						Name: "is_special",
						Label: commercetools.LocalizedString{
							"en": "is_special",
						},
						Required:  false,
						InputHint: customize.SingleLineTextInputHint,
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.boolean.json",
		},
		{
			desc: "Type with string field definition",
			input: &customize.Type{
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
					{
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
			},
			fixture: "type.string.json",
		},
		{
			desc: "Type with localized string field definition",
			input: &customize.Type{
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
					{
						Type: customize.LocalizedStringType{},
						Name: "translated_offer_name",
						Label: commercetools.LocalizedString{
							"en": "translated_offer_name",
						},
						Required:  false,
						InputHint: customize.SingleLineTextInputHint,
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.lstring.json",
		},
		{
			desc: "Type with enum field definition",
			input: &customize.Type{
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
					{
						Type: customize.EnumType{
							Values: []customize.EnumValue{
								{
									Key:   "enum",
									Label: "enum",
								},
							},
						},
						Name: "offer_enum",
						Label: commercetools.LocalizedString{
							"en": "offer_enum",
						},
						Required:  false,
						InputHint: customize.SingleLineTextInputHint,
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.enum.json",
		},
		{
			desc: "Type with localized enum field definition",
			input: &customize.Type{
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
					{
						Type: customize.LocalizedEnumType{
							Values: []customize.LocalizedEnumValue{
								{
									Key: "enum",
									Label: commercetools.LocalizedString{
										"en": "enum",
									},
								},
							},
						},
						Name: "translated_offer_enum",
						Label: commercetools.LocalizedString{
							"en": "translated_offer_enum",
						},
						Required:  false,
						InputHint: customize.SingleLineTextInputHint,
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.lenum.json",
		},
		{
			desc: "Type with number field definition",
			input: &customize.Type{
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
					{
						Type: customize.NumberType{},
						Name: "offer_number",
						Label: commercetools.LocalizedString{
							"en": "offer_number",
						},
						Required:  false,
						InputHint: customize.SingleLineTextInputHint,
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.number.json",
		},
		{
			desc: "Type with money field definition",
			input: &customize.Type{
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
					{
						Type: customize.MoneyType{},
						Name: "offer_price",
						Label: commercetools.LocalizedString{
							"en": "offer_price",
						},
						Required:  false,
						InputHint: customize.SingleLineTextInputHint,
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.money.json",
		},
		{
			desc: "Type with date field definition",
			input: &customize.Type{
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
					{
						Type: customize.DateType{},
						Name: "offer_date",
						Label: commercetools.LocalizedString{
							"en": "offer_date",
						},
						Required:  false,
						InputHint: customize.SingleLineTextInputHint,
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.date.json",
		},
		{
			desc: "Type with time field definition",
			input: &customize.Type{
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
					{
						Type: customize.TimeType{},
						Name: "offer_time",
						Label: commercetools.LocalizedString{
							"en": "offer_time",
						},
						Required:  false,
						InputHint: customize.SingleLineTextInputHint,
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.time.json",
		},
		{
			desc: "Type with date-time field definition",
			input: &customize.Type{
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
					{
						Type: customize.DateTimeType{},
						Name: "offer_date_time",
						Label: commercetools.LocalizedString{
							"en": "offer_date_time",
						},
						Required:  false,
						InputHint: customize.SingleLineTextInputHint,
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.datetime.json",
		},
		{
			desc: "Type with reference field definition",
			input: &customize.Type{
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
					{
						Type: customize.ReferenceType{
							ReferenceTypeID: "product",
						},
						Name: "offer_reference",
						Label: commercetools.LocalizedString{
							"en": "offer_reference",
						},
						Required:  false,
						InputHint: customize.SingleLineTextInputHint,
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.reference.json",
		},
		{
			desc: "Type with set field definition",
			input: &customize.Type{
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
					{
						Type: customize.SetType{
							ElementType: customize.StringType{},
						},
						Name: "offer_set",
						Label: commercetools.LocalizedString{
							"en": "offer_set",
						},
						Required:  false,
						InputHint: customize.SingleLineTextInputHint,
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
			client, server := testutil.MockClient(t, fixture(tC.fixture), nil, nil)
			defer server.Close()

			for i, err := range commercetools.ValidateStruct(*tC.input) {
				fmt.Printf("\t%d. %s\n", i+1, err.Error())
			}

			svc := customize.New(client)
			result, err := svc.TypeGetByID("1234")
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
			object: customize.BooleanType{},
			json: `{
				"name": "Boolean"
			}`,
		},
		{
			desc:   "String type",
			object: customize.StringType{},
			json: `{
				"name": "String"
			}`,
		},
		{
			desc:   "Localized string type",
			object: customize.LocalizedStringType{},
			json: `{
				"name": "LocalizedString"
			}`,
		},
		{
			desc: "Enum type",
			object: customize.EnumType{
				Values: []customize.EnumValue{
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
			object: customize.LocalizedEnumType{
				Values: []customize.LocalizedEnumValue{
					{
						Key: "test",
						Label: commercetools.LocalizedString{
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
			object: customize.NumberType{},
			json: `{
				"name": "Number"
			}`,
		},
		{
			desc:   "Money type",
			object: customize.MoneyType{},
			json: `{
				"name": "Money"
			}`,
		},
		{
			desc:   "Date type",
			object: customize.DateType{},
			json: `{
				"name": "Date"
			}`,
		},
		{
			desc:   "Time type",
			object: customize.TimeType{},
			json: `{
				"name": "Time"
			}`,
		},
		{
			desc:   "Date time type",
			object: customize.DateTimeType{},
			json: `{
				"name": "DateTime"
			}`,
		},
		{
			desc: "Reference type",
			object: customize.ReferenceType{
				ReferenceTypeID: "product",
			},
			json: `{
				"name": "Reference",
				"referenceTypeId": "product"
			}`,
		},
		{
			desc: "Set type",
			object: customize.SetType{
				ElementType: customize.BooleanType{},
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

		// Validate Marshalling (object -> json)
		t.Run(tC.desc, func(t *testing.T) {
			output, err := tC.object.MarshalJSON()
			assert.Equal(t, nil, err)
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

			fd := customize.FieldDefinition{}

			err := fd.UnmarshalJSON([]byte(buf))

			assert.Equal(t, nil, err)
			assert.Equal(t, tC.object, fd.Type)
		})
	}
}
