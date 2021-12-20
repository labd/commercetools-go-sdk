package platform_test

import (
	"encoding/json"
	"testing"

	"github.com/labd/commercetools-go-sdk/ctutils"
	"github.com/labd/commercetools-go-sdk/platform"
	"github.com/stretchr/testify/assert"
)

func TestSetKeyMarshall(t *testing.T) {
	obj := platform.ExtensionSetKeyAction{
		Key: ctutils.StringRef("my-key"),
	}

	result, err := json.Marshal(obj)
	assert.Nil(t, err)
	assert.JSONEq(t, `{
		"key":"my-key",
		"action":"setKey"
	}`, string(result))
}

func TestChangeTriggersMarshall(t *testing.T) {
	obj := platform.ExtensionChangeTriggersAction{
		Triggers: []platform.ExtensionTrigger{
			{
				ResourceTypeId: "product",
				Actions:        []platform.ExtensionAction{"Create", "Change"},
			},
		},
	}

	result, err := json.Marshal(obj)
	assert.Nil(t, err)
	assert.JSONEq(t, `{
		"triggers": [
			{
				"resourceTypeId": "product",
				"actions": ["Create", "Change"]
			}
		],
		"action": "changeTriggers"
	}`, string(result))
}

func TestChangeDestinationMarshall(t *testing.T) {
	obj := platform.ExtensionChangeDestinationAction{
		Destination: platform.ExtensionHttpDestination{
			Url: "http://example.com",
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
