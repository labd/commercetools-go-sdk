package types_test

import (
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/service/types"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
)

func TestTypeCreate(t *testing.T) {
	output := testutil.RequestData{}

	client, server := testutil.MockClient(t, testutil.Fixture("type.create.json"), &output, nil)
	defer server.Close()
	svc := types.New(client)

	input := &types.TypeDraft{
		Key: "lineitemtype",
		Name: commercetools.LocalizedString{
			"en": "lineitem",
		},
		Description: commercetools.LocalizedString{
			"en": "description",
		},
		ResourceTypeIds: []string{"line-item"},
		FieldDefinitions: []types.FieldDefinition{
			{
				Type: types.StringType{},
				Name: "offer_name",
				Label: commercetools.LocalizedString{
					"en": "offer_name",
				},
				Required:  false,
				InputHint: types.SingleLineTextInputHint,
			},
		},
	}

	fmt.Println(output)

	_, err := svc.Create(input)
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
		input       *types.UpdateInput
		requestBody string
	}{
		{
			desc: "Change key",
			input: &types.UpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&types.ChangeKey{
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
			input: &types.UpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&types.ChangeName{
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
			input: &types.UpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&types.SetDescription{
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
			input: &types.UpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&types.AddFieldDefinition{
						FieldDefinition: types.FieldDefinition{
							Type: types.BooleanType{},
							Name: "custom-field",
							Label: commercetools.LocalizedString{
								"en": "custom-field",
							},
							Required:  true,
							InputHint: types.MultiLineTextInputHint,
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
			input: &types.UpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&types.RemoveFieldDefinition{
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
			input: &types.UpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&types.ChangeLabel{
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
			input: &types.UpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&types.AddEnumValue{
						FieldName: "custom-field",
						Value: commercetools.EnumValue{
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
			input: &types.UpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&types.AddLocalizedEnumValue{
						FieldName: "custom-field",
						Value: commercetools.LocalizedEnumValue{
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
			input: &types.UpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&types.ChangeFieldDefinitionsOrder{
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
			input: &types.UpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&types.ChangeEnumValuesOrder{
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
			input: &types.UpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&types.ChangeLocalizedEnumValuesOrder{
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
			svc := types.New(client)

			_, err := svc.Update(tC.input)
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
	svc := types.New(client)

	_, err := svc.DeleteByID("1234", 2)
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
	svc := types.New(client)

	_, err := svc.DeleteByKey("1234", 2)
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
		input   *types.Type
		fixture string
	}{
		{
			desc: "Type with boolean field definition",
			input: &types.Type{
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
				FieldDefinitions: []types.FieldDefinition{
					{
						Type: types.BooleanType{},
						Name: "is_special",
						Label: commercetools.LocalizedString{
							"en": "is_special",
						},
						Required:  false,
						InputHint: types.SingleLineTextInputHint,
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.boolean.json",
		},
		{
			desc: "Type with string field definition",
			input: &types.Type{
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
				FieldDefinitions: []types.FieldDefinition{
					{
						Type: types.StringType{},
						Name: "offer_name",
						Label: commercetools.LocalizedString{
							"en": "offer_name",
						},
						Required:  false,
						InputHint: types.SingleLineTextInputHint,
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.string.json",
		},
		{
			desc: "Type with localized string field definition",
			input: &types.Type{
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
				FieldDefinitions: []types.FieldDefinition{
					{
						Type: types.LocalizedStringType{},
						Name: "translated_offer_name",
						Label: commercetools.LocalizedString{
							"en": "translated_offer_name",
						},
						Required:  false,
						InputHint: types.SingleLineTextInputHint,
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.lstring.json",
		},
		{
			desc: "Type with enum field definition",
			input: &types.Type{
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
				FieldDefinitions: []types.FieldDefinition{
					{
						Type: types.EnumType{
							Values: []commercetools.EnumValue{
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
						InputHint: types.SingleLineTextInputHint,
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.enum.json",
		},
		{
			desc: "Type with localized enum field definition",
			input: &types.Type{
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
				FieldDefinitions: []types.FieldDefinition{
					{
						Type: types.LocalizedEnumType{
							Values: []commercetools.LocalizedEnumValue{
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
						InputHint: types.SingleLineTextInputHint,
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.lenum.json",
		},
		{
			desc: "Type with number field definition",
			input: &types.Type{
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
				FieldDefinitions: []types.FieldDefinition{
					{
						Type: types.NumberType{},
						Name: "offer_number",
						Label: commercetools.LocalizedString{
							"en": "offer_number",
						},
						Required:  false,
						InputHint: types.SingleLineTextInputHint,
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.number.json",
		},
		{
			desc: "Type with money field definition",
			input: &types.Type{
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
				FieldDefinitions: []types.FieldDefinition{
					{
						Type: types.MoneyType{},
						Name: "offer_price",
						Label: commercetools.LocalizedString{
							"en": "offer_price",
						},
						Required:  false,
						InputHint: types.SingleLineTextInputHint,
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.money.json",
		},
		{
			desc: "Type with date field definition",
			input: &types.Type{
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
				FieldDefinitions: []types.FieldDefinition{
					{
						Type: types.DateType{},
						Name: "offer_date",
						Label: commercetools.LocalizedString{
							"en": "offer_date",
						},
						Required:  false,
						InputHint: types.SingleLineTextInputHint,
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.date.json",
		},
		{
			desc: "Type with time field definition",
			input: &types.Type{
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
				FieldDefinitions: []types.FieldDefinition{
					{
						Type: types.TimeType{},
						Name: "offer_time",
						Label: commercetools.LocalizedString{
							"en": "offer_time",
						},
						Required:  false,
						InputHint: types.SingleLineTextInputHint,
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.time.json",
		},
		{
			desc: "Type with date-time field definition",
			input: &types.Type{
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
				FieldDefinitions: []types.FieldDefinition{
					{
						Type: types.DateTimeType{},
						Name: "offer_date_time",
						Label: commercetools.LocalizedString{
							"en": "offer_date_time",
						},
						Required:  false,
						InputHint: types.SingleLineTextInputHint,
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.datetime.json",
		},
		{
			desc: "Type with reference field definition",
			input: &types.Type{
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
				FieldDefinitions: []types.FieldDefinition{
					{
						Type: types.ReferenceType{
							ReferenceTypeID: "product",
						},
						Name: "offer_reference",
						Label: commercetools.LocalizedString{
							"en": "offer_reference",
						},
						Required:  false,
						InputHint: types.SingleLineTextInputHint,
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.reference.json",
		},
		{
			desc: "Type with set field definition",
			input: &types.Type{
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
				FieldDefinitions: []types.FieldDefinition{
					{
						Type: types.SetType{
							ElementType: types.StringType{},
						},
						Name: "offer_set",
						Label: commercetools.LocalizedString{
							"en": "offer_set",
						},
						Required:  false,
						InputHint: types.SingleLineTextInputHint,
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

			svc := types.New(client)
			result, err := svc.GetByID("1234")
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
			object: types.BooleanType{},
			json: `{
				"name": "Boolean"
			}`,
		},
		{
			desc:   "String type",
			object: types.StringType{},
			json: `{
				"name": "String"
			}`,
		},
		{
			desc:   "Localized string type",
			object: types.LocalizedStringType{},
			json: `{
				"name": "LocalizedString"
			}`,
		},
		{
			desc: "Enum type",
			object: types.EnumType{
				Values: []commercetools.EnumValue{
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
			object: types.LocalizedEnumType{
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
			object: types.NumberType{},
			json: `{
				"name": "Number"
			}`,
		},
		{
			desc:   "Money type",
			object: types.MoneyType{},
			json: `{
				"name": "Money"
			}`,
		},
		{
			desc:   "Date type",
			object: types.DateType{},
			json: `{
				"name": "Date"
			}`,
		},
		{
			desc:   "Time type",
			object: types.TimeType{},
			json: `{
				"name": "Time"
			}`,
		},
		{
			desc:   "Date time type",
			object: types.DateTimeType{},
			json: `{
				"name": "DateTime"
			}`,
		},
		{
			desc: "Reference type",
			object: types.ReferenceType{
				ReferenceTypeID: "product",
			},
			json: `{
				"name": "Reference",
				"referenceTypeId": "product"
			}`,
		},
		{
			desc: "Set type",
			object: types.SetType{
				ElementType: types.BooleanType{},
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

			fd := types.FieldDefinition{}

			err := fd.UnmarshalJSON([]byte(buf))

			assert.Nil(t, err)
			assert.Equal(t, tC.object, fd.Type)
		})
	}
}
