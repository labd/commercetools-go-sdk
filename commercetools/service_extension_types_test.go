package commercetools_test

import (
	"encoding/json"
	"testing"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/stretchr/testify/assert"
)

func TestExtensionMarshall_basicauth(t *testing.T) {
	obj := commercetools.Extension{
		Key: "my-key",
		Destination: commercetools.ExtensionHTTPDestination{
			URL: "http://example.com",
			Authentication: commercetools.ExtensionAuthorizationHeaderAuthentication{
				HeaderValue: "Basic 12345",
			},
		},
		Triggers: []commercetools.ExtensionTrigger{
			{
				ResourceTypeID: "product",
				Actions:        []commercetools.ExtensionAction{"Create"},
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
	obj := commercetools.Extension{
		Key: "my-key",
		Destination: commercetools.ExtensionHTTPDestination{
			URL: "http://example.com",
			Authentication: commercetools.ExtensionAzureFunctionsAuthentication{
				Key: "MyAzureKey000",
			},
		},
		Triggers: []commercetools.ExtensionTrigger{
			{
				ResourceTypeID: "product",
				Actions:        []commercetools.ExtensionAction{"Create"},
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

func TestExtensionMarshall_awslambda(t *testing.T) {
	obj := commercetools.Extension{
		Key: "my-key",
		Destination: commercetools.ExtensionAWSLambdaDestination{
			Arn:          "arn:some:arn",
			AccessKey:    "12331232131",
			AccessSecret: "somesecret",
		},
		Triggers: []commercetools.ExtensionTrigger{
			{
				ResourceTypeID: "product",
				Actions:        []commercetools.ExtensionAction{"Create"},
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
		"type":"AWSLambda",
		"arn": "arn:some:arn",
		"accessKey": "12331232131",
		"accessSecret": "somesecret"
	},
	"triggers":[
		{"resourceTypeId":"product","actions":["Create"]}
	],
	"createdAt":"0001-01-01T00:00:00Z",
	"lastModifiedAt":"0001-01-01T00:00:00Z"}`, string(result))
}
