package extensions_test

import (
	"encoding/json"
	"testing"

	"github.com/labd/commercetools-go-sdk/service/extensions"
	"github.com/stretchr/testify/assert"
)

func TestSetKeyMarshall(t *testing.T) {
	obj := extensions.SetKey{
		Key: "my-key",
	}

	result, err := json.Marshal(obj)
	assert.Nil(t, err)
	assert.JSONEq(t, `{
	"key":"my-key",
	"action":"setKey"}`, string(result))
}

func TestChangeTriggersMarshall(t *testing.T) {
	obj := extensions.ChangeTriggers{
		Triggers: []extensions.Trigger{
			extensions.Trigger{
				ResourceTypeID: "product",
				Actions:        []string{"Create", "Change"},
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
	obj := extensions.ChangeDestination{
		Destination: extensions.DestinationHTTP{
			URL: "http://example.com",
		},
	}

	result, err := json.Marshal(obj)
	assert.Nil(t, err)
	assert.JSONEq(t, `{
	"destination":{
		"type": "HTTP",
		"url":"http://example.com",
		"authentication":null
	},
	"action":"changeDestination"}`, string(result))
}
