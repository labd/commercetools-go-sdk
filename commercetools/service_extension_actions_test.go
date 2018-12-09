package commercetools_test

import (
	"encoding/json"
	"testing"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/stretchr/testify/assert"
)

func TestSetKeyMarshall(t *testing.T) {
	obj := commercetools.ExtensionSetKeyAction{
		Key: "my-key",
	}

	result, err := json.Marshal(obj)
	assert.Nil(t, err)
	assert.JSONEq(t, `{
	"key":"my-key",
	"action":"setKey"}`, string(result))
}

func TestChangeTriggersMarshall(t *testing.T) {
	obj := commercetools.ExtensionChangeTriggersAction{
		Triggers: []commercetools.ExtensionTrigger{
			commercetools.ExtensionTrigger{
				ResourceTypeID: "product",
				Actions:        []commercetools.ExtensionAction{"Create", "Change"},
			},
		},
	}

	result, err := json.Marshal(obj)
	assert.Nil(t, err)
	assert.JSONEq(t, `{
	"triggers":[{
		"resourceTypeId":"product",
		"actions":["Create", "Change"]
	}],
	"action":"changeTriggers"}`, string(result))
}

func TestChangeDestinationMarshall(t *testing.T) {
	obj := commercetools.ExtensionChangeDestinationAction{
		Destination: commercetools.ExtensionHttpDestination{
			URL: "http://example.com",
		},
	}

	result, err := json.Marshal(obj)
	assert.Nil(t, err)
	assert.JSONEq(t, `{
	"destination":{
		"type": "HTTP",
		"url":"http://example.com"
	},
	"action":"changeDestination"}`, string(result))
}
