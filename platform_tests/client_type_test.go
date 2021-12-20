package platform_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/labd/commercetools-go-sdk/ctutils"
	"github.com/labd/commercetools-go-sdk/platform"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
)

func TestTypeCreate(t *testing.T) {
	output := testutil.RequestData{}
	response := testutil.Fixture("type.create.json", 201)

	client, server := testutil.MockClient(t, response, &output, nil)
	defer server.Close()

	inputHint := platform.TypeTextInputHint("SingleLine")

	input := platform.TypeDraft{
		Key: "lineitemtype",
		Name: platform.LocalizedString{
			"en": "lineitem",
		},
		Description: &platform.LocalizedString{
			"en": "description",
		},
		ResourceTypeIds: []platform.ResourceTypeId{"line-item"},
		FieldDefinitions: []platform.FieldDefinition{
			{
				Type: platform.CustomFieldStringType{},
				Name: "offer_name",
				Label: platform.LocalizedString{
					"en": "offer_name",
				},
				Required:  false,
				InputHint: &inputHint,
			},
		},
	}

	_, err := client.WithProjectKey("unittest").Types().Post(input).Execute(context.TODO())
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
	inputHint := platform.TypeTextInputHint("MultiLine")

	testCases := []struct {
		desc        string
		input       platform.TypeUpdate
		requestBody string
	}{
		{
			desc: "Change key",
			input: platform.TypeUpdate{
				Version: 2,
				Actions: []platform.TypeUpdateAction{
					platform.TypeChangeKeyAction{
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
			input: platform.TypeUpdate{
				Version: 2,
				Actions: []platform.TypeUpdateAction{
					platform.TypeChangeNameAction{
						Name: platform.LocalizedString{
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
			input: platform.TypeUpdate{
				Version: 2,
				Actions: []platform.TypeUpdateAction{
					platform.TypeSetDescriptionAction{
						Description: &platform.LocalizedString{
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
			input: platform.TypeUpdate{
				Version: 2,
				Actions: []platform.TypeUpdateAction{
					platform.TypeAddFieldDefinitionAction{
						FieldDefinition: platform.FieldDefinition{
							Type: platform.CustomFieldBooleanType{},
							Name: "custom-field",
							Label: platform.LocalizedString{
								"en": "custom-field",
							},
							Required:  true,
							InputHint: &inputHint,
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
			input: platform.TypeUpdate{
				Version: 2,
				Actions: []platform.TypeUpdateAction{
					platform.TypeRemoveFieldDefinitionAction{
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
			input: platform.TypeUpdate{
				Version: 2,
				Actions: []platform.TypeUpdateAction{
					platform.TypeChangeLabelAction{
						FieldName: "custom-field",
						Label: platform.LocalizedString{
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
			input: platform.TypeUpdate{
				Version: 2,
				Actions: []platform.TypeUpdateAction{
					platform.TypeAddEnumValueAction{
						FieldName: "custom-field",
						Value: platform.CustomFieldEnumValue{
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
			input: platform.TypeUpdate{
				Version: 2,
				Actions: []platform.TypeUpdateAction{
					platform.TypeAddLocalizedEnumValueAction{
						FieldName: "custom-field",
						Value: platform.CustomFieldLocalizedEnumValue{
							Key: "custom-enum",
							Label: platform.LocalizedString{
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
			input: platform.TypeUpdate{
				Version: 2,
				Actions: []platform.TypeUpdateAction{
					platform.TypeChangeFieldDefinitionOrderAction{
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
			input: platform.TypeUpdate{
				Version: 2,
				Actions: []platform.TypeUpdateAction{
					platform.TypeChangeEnumValueOrderAction{
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
			input: platform.TypeUpdate{
				Version: 2,
				Actions: []platform.TypeUpdateAction{
					platform.TypeChangeLocalizedEnumValueOrderAction{
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
			response := testutil.Fixture("type.create.json", 200)

			client, server := testutil.MockClient(t, response, &output, nil)
			defer server.Close()

			_, err := client.WithProjectKey("unittest").Types().WithId("1234").Post(tC.input).Execute(context.TODO())
			assert.Nil(t, err)
			assert.Equal(t, "/unittest/types/1234", output.URL.Path)
			assert.JSONEq(t, tC.requestBody, output.JSON)
		})
	}
}

func TestTypeDeleteByID(t *testing.T) {
	output := testutil.RequestData{}
	response := testutil.Fixture("type.create.json", 200)
	client, server := testutil.MockClient(t, response, &output, nil)
	defer server.Close()

	_, err := client.
		WithProjectKey("unittest").
		Types().
		WithId("1234").
		Delete().
		WithQueryParams(platform.ByProjectKeyTypesByIDRequestMethodDeleteInput{
			Version: 2,
		}).
		Execute(context.TODO())
	assert.Nil(t, err)

	params := url.Values{}
	params.Add("version", "2")
	assert.Equal(t, params, output.URL.Query())
	assert.Equal(t, "/unittest/types/1234", output.URL.Path)
}

func TestTypeDeleteByKey(t *testing.T) {
	output := testutil.RequestData{}
	response := testutil.Fixture("type.create.json", 200)
	client, server := testutil.MockClient(t, response, &output, nil)
	defer server.Close()

	_, err := client.
		WithProjectKey("unittest").
		Types().
		WithKey("1234").
		Delete().
		WithQueryParams(platform.ByProjectKeyTypesKeyByKeyRequestMethodDeleteInput{
			Version: 2,
		}).
		Execute(context.TODO())
	assert.Nil(t, err)

	params := url.Values{}
	params.Add("version", "2")
	assert.Equal(t, params, output.URL.Query())
	assert.Equal(t, "/unittest/types/key=1234", output.URL.Path)
}

func TestTypeGetByID(t *testing.T) {
	timestamp, _ := time.Parse(time.RFC3339, "2015-10-07T06:56:19.217Z")
	inputHint := platform.TypeTextInputHint("SingleLine")

	testCases := []struct {
		desc    string
		input   platform.Type
		fixture string
	}{
		{
			desc: "Type with boolean field definition",
			input: platform.Type{
				ID:      "1234",
				Version: 1,
				Key:     "lineitemtype",
				Name: platform.LocalizedString{
					"en": "lineitem",
				},
				Description: &platform.LocalizedString{
					"en": "description",
				},
				ResourceTypeIds: []platform.ResourceTypeId{"line-item"},
				FieldDefinitions: []platform.FieldDefinition{
					{
						Type: platform.CustomFieldBooleanType{},
						Name: "is_special",
						Label: platform.LocalizedString{
							"en": "is_special",
						},
						Required:  false,
						InputHint: &inputHint,
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.boolean.json",
		},
		{
			desc: "Type with string field definition",
			input: platform.Type{
				ID:      "1234",
				Version: 1,
				Key:     "lineitemtype",
				Name: platform.LocalizedString{
					"en": "lineitem",
				},
				Description: &platform.LocalizedString{
					"en": "description",
				},
				ResourceTypeIds: []platform.ResourceTypeId{"line-item"},
				FieldDefinitions: []platform.FieldDefinition{
					{
						Type: platform.CustomFieldStringType{},
						Name: "offer_name",
						Label: platform.LocalizedString{
							"en": "offer_name",
						},
						Required:  false,
						InputHint: &inputHint,
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.string.json",
		},
		{
			desc: "Type with localized string field definition",
			input: platform.Type{
				ID:      "1234",
				Version: 1,
				Key:     "lineitemtype",
				Name: platform.LocalizedString{
					"en": "lineitem",
				},
				Description: &platform.LocalizedString{
					"en": "description",
				},
				ResourceTypeIds: []platform.ResourceTypeId{"line-item"},
				FieldDefinitions: []platform.FieldDefinition{
					{
						Type: platform.CustomFieldLocalizedStringType{},
						Name: "translated_offer_name",
						Label: platform.LocalizedString{
							"en": "translated_offer_name",
						},
						Required:  false,
						InputHint: &inputHint,
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.lstring.json",
		},
		{
			desc: "Type with enum field definition",
			input: platform.Type{
				ID:      "1234",
				Version: 1,
				Key:     "lineitemtype",
				Name: platform.LocalizedString{
					"en": "lineitem",
				},
				Description: &platform.LocalizedString{
					"en": "description",
				},
				ResourceTypeIds: []platform.ResourceTypeId{"line-item"},
				FieldDefinitions: []platform.FieldDefinition{
					{
						Type: platform.CustomFieldEnumType{
							Values: []platform.CustomFieldEnumValue{
								{
									Key:   "enum",
									Label: "enum",
								},
							},
						},
						Name: "offer_enum",
						Label: platform.LocalizedString{
							"en": "offer_enum",
						},
						Required:  false,
						InputHint: &inputHint,
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.enum.json",
		},
		{
			desc: "Type with localized enum field definition",
			input: platform.Type{
				ID:      "1234",
				Version: 1,
				Key:     "lineitemtype",
				Name: platform.LocalizedString{
					"en": "lineitem",
				},
				Description: &platform.LocalizedString{
					"en": "description",
				},
				ResourceTypeIds: []platform.ResourceTypeId{"line-item"},
				FieldDefinitions: []platform.FieldDefinition{
					{
						Type: platform.CustomFieldLocalizedEnumType{
							Values: []platform.CustomFieldLocalizedEnumValue{
								{
									Key: "enum",
									Label: platform.LocalizedString{
										"en": "enum",
									},
								},
							},
						},
						Name: "translated_offer_enum",
						Label: platform.LocalizedString{
							"en": "translated_offer_enum",
						},
						Required:  false,
						InputHint: &inputHint,
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.lenum.json",
		},
		{
			desc: "Type with number field definition",
			input: platform.Type{
				ID:      "1234",
				Version: 1,
				Key:     "lineitemtype",
				Name: platform.LocalizedString{
					"en": "lineitem",
				},
				Description: &platform.LocalizedString{
					"en": "description",
				},
				ResourceTypeIds: []platform.ResourceTypeId{"line-item"},
				FieldDefinitions: []platform.FieldDefinition{
					{
						Type: platform.CustomFieldNumberType{},
						Name: "offer_number",
						Label: platform.LocalizedString{
							"en": "offer_number",
						},
						Required:  false,
						InputHint: &inputHint,
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.number.json",
		},
		{
			desc: "Type with money field definition",
			input: platform.Type{
				ID:      "1234",
				Version: 1,
				Key:     "lineitemtype",
				Name: platform.LocalizedString{
					"en": "lineitem",
				},
				Description: &platform.LocalizedString{
					"en": "description",
				},
				ResourceTypeIds: []platform.ResourceTypeId{"line-item"},
				FieldDefinitions: []platform.FieldDefinition{
					{
						Type: platform.CustomFieldMoneyType{},
						Name: "offer_price",
						Label: platform.LocalizedString{
							"en": "offer_price",
						},
						Required:  false,
						InputHint: &inputHint,
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.money.json",
		},
		{
			desc: "Type with date field definition",
			input: platform.Type{
				ID:      "1234",
				Version: 1,
				Key:     "lineitemtype",
				Name: platform.LocalizedString{
					"en": "lineitem",
				},
				Description: &platform.LocalizedString{
					"en": "description",
				},
				ResourceTypeIds: []platform.ResourceTypeId{"line-item"},
				FieldDefinitions: []platform.FieldDefinition{
					{
						Type: platform.CustomFieldDateType{},
						Name: "offer_date",
						Label: platform.LocalizedString{
							"en": "offer_date",
						},
						Required:  false,
						InputHint: &inputHint,
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.date.json",
		},
		{
			desc: "Type with time field definition",
			input: platform.Type{
				ID:      "1234",
				Version: 1,
				Key:     "lineitemtype",
				Name: platform.LocalizedString{
					"en": "lineitem",
				},
				Description: &platform.LocalizedString{
					"en": "description",
				},
				ResourceTypeIds: []platform.ResourceTypeId{"line-item"},
				FieldDefinitions: []platform.FieldDefinition{
					{
						Type: platform.CustomFieldTimeType{},
						Name: "offer_time",
						Label: platform.LocalizedString{
							"en": "offer_time",
						},
						Required:  false,
						InputHint: &inputHint,
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.time.json",
		},
		{
			desc: "Type with date-time field definition",
			input: platform.Type{
				ID:      "1234",
				Version: 1,
				Key:     "lineitemtype",
				Name: platform.LocalizedString{
					"en": "lineitem",
				},
				Description: &platform.LocalizedString{
					"en": "description",
				},
				ResourceTypeIds: []platform.ResourceTypeId{"line-item"},
				FieldDefinitions: []platform.FieldDefinition{
					{
						Type: platform.CustomFieldDateTimeType{},
						Name: "offer_date_time",
						Label: platform.LocalizedString{
							"en": "offer_date_time",
						},
						Required:  false,
						InputHint: &inputHint,
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.datetime.json",
		},
		{
			desc: "Type with reference field definition",
			input: platform.Type{
				ID:      "1234",
				Version: 1,
				Key:     "lineitemtype",
				Name: platform.LocalizedString{
					"en": "lineitem",
				},
				Description: &platform.LocalizedString{
					"en": "description",
				},
				ResourceTypeIds: []platform.ResourceTypeId{"line-item"},
				FieldDefinitions: []platform.FieldDefinition{
					{
						Type: platform.CustomFieldReferenceType{
							ReferenceTypeId: "product",
						},
						Name: "offer_reference",
						Label: platform.LocalizedString{
							"en": "offer_reference",
						},
						Required:  false,
						InputHint: &inputHint,
					},
				},
				CreatedAt:      timestamp,
				LastModifiedAt: timestamp,
			},
			fixture: "type.reference.json",
		},
		{
			desc: "Type with set field definition",
			input: platform.Type{
				ID:      "1234",
				Version: 1,
				Key:     "lineitemtype",
				Name: platform.LocalizedString{
					"en": "lineitem",
				},
				Description: &platform.LocalizedString{
					"en": "description",
				},
				ResourceTypeIds: []platform.ResourceTypeId{"line-item"},
				FieldDefinitions: []platform.FieldDefinition{
					{
						Type: platform.CustomFieldSetType{
							ElementType: platform.CustomFieldStringType{},
						},
						Name: "offer_set",
						Label: platform.LocalizedString{
							"en": "offer_set",
						},
						Required:  false,
						InputHint: &inputHint,
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
			client, server := testutil.MockClient(t, testutil.Fixture(tC.fixture, 200), nil, nil)
			defer server.Close()

			result, err := client.WithProjectKey("unittest").Types().WithId("1234").Get().Execute(context.TODO())
			assert.Nil(t, err)

			assert.Equal(t, tC.input, *result)
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
			object: platform.CustomFieldBooleanType{},
			json: `{
				"name": "Boolean"
			}`,
		},
		{
			desc:   "String type",
			object: platform.CustomFieldStringType{},
			json: `{
				"name": "String"
			}`,
		},
		{
			desc:   "Localized string type",
			object: platform.CustomFieldLocalizedStringType{},
			json: `{
				"name": "LocalizedString"
			}`,
		},
		{
			desc: "Enum type",
			object: platform.CustomFieldEnumType{
				Values: []platform.CustomFieldEnumValue{
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
			object: platform.CustomFieldLocalizedEnumType{
				Values: []platform.CustomFieldLocalizedEnumValue{
					{
						Key: "test",
						Label: platform.LocalizedString{
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
			object: platform.CustomFieldNumberType{},
			json: `{
				"name": "Number"
			}`,
		},
		{
			desc:   "Money type",
			object: platform.CustomFieldMoneyType{},
			json: `{
				"name": "Money"
			}`,
		},
		{
			desc:   "Date type",
			object: platform.CustomFieldDateType{},
			json: `{
				"name": "Date"
			}`,
		},
		{
			desc:   "Time type",
			object: platform.CustomFieldTimeType{},
			json: `{
				"name": "Time"
			}`,
		},
		{
			desc:   "Date time type",
			object: platform.CustomFieldDateTimeType{},
			json: `{
				"name": "DateTime"
			}`,
		},
		{
			desc: "Reference type",
			object: platform.CustomFieldReferenceType{
				ReferenceTypeId: "product",
			},
			json: `{
				"name": "Reference",
				"referenceTypeId": "product"
			}`,
		},
		{
			desc: "Set type",
			object: platform.CustomFieldSetType{
				ElementType: platform.CustomFieldBooleanType{},
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

			fd := platform.FieldDefinition{}
			err := json.Unmarshal([]byte(buf), &fd)
			assert.Nil(t, err)
			assert.Equal(t, tC.object, fd.Type)
		})
	}
}

func TestTypeQuery(t *testing.T) {
	output := testutil.RequestData{}
	response := testutil.ResponseData{StatusCode: 200, Body: "{}"}

	client, server := testutil.MockClient(t, response, &output, nil)
	defer server.Close()

	_, err := client.
		WithProjectKey("unittest").
		Types().
		Get().
		WithQueryParams(platform.ByProjectKeyTypesRequestMethodGetInput{
			Limit: ctutils.IntRef(500),
		}).
		Execute(context.TODO())

	assert.Nil(t, err)
	assert.Equal(t, url.Values{"limit": []string{"500"}}, output.URL.Query())
}
