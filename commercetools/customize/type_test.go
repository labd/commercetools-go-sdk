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
	output := testutil.RequestData{}

	client, server := testutil.MockClient(t, fixture("type.update.json"), &output, nil)
	defer server.Close()
	svc := customize.New(client)

	input := &customize.TypeUpdateInput{
		ID:      "1234",
		Version: 2,
		Actions: commercetools.UpdateActions{
			&customize.TypeChangeKey{
				Key: "123456",
			},
		},
	}

	fmt.Println(output)

	_, err := svc.TypeUpdate(input)
	assert.Equal(t, nil, err)

	expectedBody := `{
		"version": 2,
		"actions": [
			{
				"action": "changeKey",
				"key": "123456"
			}
		]
	}`

	assert.JSONEq(t, expectedBody, output.JSON)
	assert.Equal(t, "/unittest/types/1234", output.URL.Path)
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
