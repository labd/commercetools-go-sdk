package platform_test

// import (
// 	"encoding/json"
// 	"testing"

// 	"github.com/labd/commercetools-go-sdk/commercetools"
// 	"github.com/stretchr/testify/assert"
// )

// func TestExtensionMarshall_basicauth(t *testing.T) {
// 	obj := platform.Extension{
// 		Key: "my-key",
// 		Destination: platform.ExtensionHTTPDestination{
// 			URL: "http://example.com",
// 			Authentication: platform.ExtensionAuthorizationHeaderAuthentication{
// 				HeaderValue: "Basic 12345",
// 			},
// 		},
// 		Triggers: []platform.ExtensionTrigger{
// 			{
// 				ResourceTypeID: "product",
// 				Actions:        []platform.ExtensionAction{"Create"},
// 			},
// 		},
// 	}

// 	result, err := json.Marshal(obj)
// 	assert.Nil(t, err)
// 	assert.JSONEq(t, `{
// 	"id":"",
// 	"version":0,
// 	"key":"my-key",
// 	"destination":{
// 		"type":"HTTP",
// 		"url":"http://example.com",
// 		"authentication":{
// 			"type":"AuthorizationHeader",
// 			"headerValue":"Basic 12345"
// 		}
// 	},
// 	"triggers":[
// 		{"resourceTypeId":"product","actions":["Create"]}
// 	],
// 	"createdAt":"0001-01-01T00:00:00Z",
// 	"lastModifiedAt":"0001-01-01T00:00:00Z"}`, string(result))
// }

// func TestExtensionMarshall_azurekey(t *testing.T) {
// 	obj := platform.Extension{
// 		Key: "my-key",
// 		Destination: platform.ExtensionHTTPDestination{
// 			URL: "http://example.com",
// 			Authentication: platform.ExtensionAzureFunctionsAuthentication{
// 				Key: "MyAzureKey000",
// 			},
// 		},
// 		Triggers: []platform.ExtensionTrigger{
// 			{
// 				ResourceTypeID: "product",
// 				Actions:        []platform.ExtensionAction{"Create"},
// 			},
// 		},
// 	}

// 	result, err := json.Marshal(obj)
// 	assert.Nil(t, err)
// 	assert.JSONEq(t, `{
// 	"id":"",
// 	"version":0,
// 	"key":"my-key",
// 	"destination":{
// 		"type":"HTTP",
// 		"url":"http://example.com",
// 		"authentication":{
// 			"type":"AzureFunctions",
// 			"key":"MyAzureKey000"
// 		}
// 	},
// 	"triggers":[
// 		{"resourceTypeId":"product","actions":["Create"]}
// 	],
// 	"createdAt":"0001-01-01T00:00:00Z",
// 	"lastModifiedAt":"0001-01-01T00:00:00Z"}`, string(result))
// }

// func TestExtensionMarshall_awslambda(t *testing.T) {
// 	obj := platform.Extension{
// 		Key: "my-key",
// 		Destination: platform.ExtensionAWSLambdaDestination{
// 			Arn:          "arn:some:arn",
// 			AccessKey:    "12331232131",
// 			AccessSecret: "somesecret",
// 		},
// 		Triggers: []platform.ExtensionTrigger{
// 			{
// 				ResourceTypeID: "product",
// 				Actions:        []platform.ExtensionAction{"Create"},
// 			},
// 		},
// 	}

// 	result, err := json.Marshal(obj)
// 	assert.Nil(t, err)
// 	assert.JSONEq(t, `{
// 	"id":"",
// 	"version":0,
// 	"key":"my-key",
// 	"destination":{
// 		"type":"AWSLambda",
// 		"arn": "arn:some:arn",
// 		"accessKey": "12331232131",
// 		"accessSecret": "somesecret"
// 	},
// 	"triggers":[
// 		{"resourceTypeId":"product","actions":["Create"]}
// 	],
// 	"createdAt":"0001-01-01T00:00:00Z",
// 	"lastModifiedAt":"0001-01-01T00:00:00Z"}`, string(result))
// }
