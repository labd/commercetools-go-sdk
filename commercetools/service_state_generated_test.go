// Automatically generated, do not edit

package commercetools_test

import (
	"context"
	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratedStateGetWithID(t *testing.T) {
	responseData := ` {
	  "id": "7c2e2694-aefe-43d7-888e-6a99514caaca",
	  "version": 1,
	  "key": "Initial",
	  "type": "LineItemState",
	  "roles": [],
	  "name": {
	    "en": "Initial"
	  },
	  "description": {
	    "en": "Initial is the first that (custom) line item gets after it's creation"
	  },
	  "builtIn": true,
	  "initial": true,
	  "createdAt": "2015-01-21T09:22:03.906Z",
	  "lastModifiedAt": "2015-01-21T09:22:03.906Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	state, err := client.StateGetWithID(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, state)
	assert.NotNil(t, state.Version)
	assert.NotEmpty(t, state.Type)
	assert.NotEmpty(t, state.LastModifiedAt)
	assert.NotEmpty(t, state.Key)
	assert.True(t, state.Initial)
	assert.NotEmpty(t, state.ID)
	assert.NotEmpty(t, state.CreatedAt)
	assert.True(t, state.BuiltIn)

}

func TestGeneratedStateDeleteWithID(t *testing.T) {
	responseData := ` {
	  "id": "7c2e2694-aefe-43d7-888e-6a99514caaca",
	  "version": 1,
	  "key": "Initial",
	  "type": "LineItemState",
	  "roles": [],
	  "name": {
	    "en": "Initial"
	  },
	  "description": {
	    "en": "Initial is the first that (custom) line item gets after it's creation"
	  },
	  "builtIn": true,
	  "initial": true,
	  "createdAt": "2015-01-21T09:22:03.906Z",
	  "lastModifiedAt": "2015-01-21T09:22:03.906Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	state, err := client.StateDeleteWithID(context.TODO(), "dummy-id", 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, state)
	assert.NotNil(t, state.Version)
	assert.NotEmpty(t, state.Type)
	assert.NotEmpty(t, state.LastModifiedAt)
	assert.NotEmpty(t, state.Key)
	assert.True(t, state.Initial)
	assert.NotEmpty(t, state.ID)
	assert.NotEmpty(t, state.CreatedAt)
	assert.True(t, state.BuiltIn)

}

func TestGeneratedStateQuery(t *testing.T) {
	responseData := ` {
	  "limit": 20,
	  "offset": 0,
	  "count": 2,
	  "total": 2,
	  "results": [
	    {
	      "id": "7c2e2694-aefe-43d7-888e-6a99514caaca",
	      "version": 1,
	      "key": "Initial",
	      "type": "LineItemState",
	      "roles": [],
	      "name": {
	        "en": "Initial"
	      },
	      "description": {
	        "en": "Initial is the first that (custom) line item gets after it's creation"
	      },
	      "builtIn": true,
	      "initial": true,
	      "createdAt": "2015-01-21T09:22:03.906Z",
	      "lastModifiedAt": "2015-01-21T09:22:03.906Z"
	    },
	    {
	      "id": "fb2b1abd-7598-40f1-906c-9e5e7c6f8ebc",
	      "version": 1,
	      "key": "test-state",
	      "type": "ProductState",
	      "roles": [],
	      "builtIn": false,
	      "initial": true,
	      "createdAt": "2016-02-18T17:18:53.338Z",
	      "lastModifiedAt": "2016-02-18T17:18:53.338Z"
	    }
	  ]
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	input := commercetools.QueryInput{}
	queryResult, err := client.StateQuery(context.TODO(), &input)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, queryResult)
	assert.NotNil(t, queryResult.Total)
	assert.NotNil(t, queryResult.Offset)
	assert.NotNil(t, queryResult.Limit)
	assert.NotNil(t, queryResult.Count)

}
