package extensions_test

import (
	"encoding/json"
	"testing"

	"github.com/labd/commercetools-go-sdk/service/extensions"
	"github.com/stretchr/testify/assert"
)

func TestExtensionMarshall_basicauth(t *testing.T) {
	obj := extensions.Extension{
		Key: "my-key",
		Destination: extensions.DestinationHTTP{
			URL: "http://example.com",
			Authentication: extensions.DestinationAuthenticationAuth{
				HeaderValue: "Basic 12345",
			},
		},
		Triggers: []extensions.Trigger{
			extensions.Trigger{
				ResourceTypeID: "product",
				Actions:        []string{"Create"},
			},
		},
	}

	result, err := json.Marshal(obj)
	assert.Nil(t, err)
	assert.JSONEq(t, `{
	"id":"",
	"version":0,
	"key":"my-key",
	"destination":{
		"type":"HTTP",
		"url":"http://example.com",
		"authentication":{
			"type":"AuthorizationHeader",
			"headerValue":"Basic 12345"
		}
	},
	"triggers":[
		{"resourceTypeId":"product","actions":["Create"]}
	],
	"createdAt":"0001-01-01T00:00:00Z",
	"lastModifiedAt":"0001-01-01T00:00:00Z"}`, string(result))
}

func TestExtensionMarshall_azurekey(t *testing.T) {
	obj := extensions.Extension{
		Key: "my-key",
		Destination: extensions.DestinationHTTP{
			URL: "http://example.com",
			Authentication: extensions.DestinationAuthenticationAzure{
				Key: "MyAzureKey000",
			},
		},
		Triggers: []extensions.Trigger{
			extensions.Trigger{
				ResourceTypeID: "product",
				Actions:        []string{"Create"},
			},
		},
	}

	result, err := json.Marshal(obj)
	assert.Nil(t, err)
	assert.JSONEq(t, `{
	"id":"",
	"version":0,
	"key":"my-key",
	"destination":{
		"type":"HTTP",
		"url":"http://example.com",
		"authentication":{
			"type":"AzureFunctions",
			"key":"MyAzureKey000"
		}
	},
	"triggers":[
		{"resourceTypeId":"product","actions":["Create"]}
	],
	"createdAt":"0001-01-01T00:00:00Z",
	"lastModifiedAt":"0001-01-01T00:00:00Z"}`, string(result))
}
