// Automatically generated, do not edit

package commercetools_test

import (
	"context"
	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratedStoreGetWithID(t *testing.T) {
	responseData := ` {
	  "id": "c2f93299-c967-44af-8c2a-d2220bf39eb2",
	  "version": 1,
	  "key": "random-key-12314",
	  "name": {
	    "en": "big store"
	  },
	  "createdAt": "1971-01-01T00:00:00.001Z",
	  "lastModifiedAt": "1972-01-01T00:00:00.001Z",
	  "distributionChannels": []
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	store, err := client.StoreGetWithID(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, store)
	assert.NotNil(t, store.Version)
	assert.NotEmpty(t, store.LastModifiedAt)
	assert.NotEmpty(t, store.Key)
	assert.NotEmpty(t, store.ID)
	assert.NotEmpty(t, store.CreatedAt)

}

func TestGeneratedStoreGetWithKey(t *testing.T) {
	responseData := ` {
	  "id": "c2f93299-c967-44af-8c2a-d2220bf39eb2",
	  "version": 1,
	  "key": "random-key-12314",
	  "name": {
	    "en": "big store"
	  },
	  "createdAt": "1971-01-01T00:00:00.001Z",
	  "lastModifiedAt": "1972-01-01T00:00:00.001Z",
	  "distributionChannels": []
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	store, err := client.StoreGetWithKey(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, store)
	assert.NotNil(t, store.Version)
	assert.NotEmpty(t, store.LastModifiedAt)
	assert.NotEmpty(t, store.Key)
	assert.NotEmpty(t, store.ID)
	assert.NotEmpty(t, store.CreatedAt)

}

func TestGeneratedStoreDeleteWithID(t *testing.T) {
	responseData := ` {
	  "id": "c2f93299-c967-44af-8c2a-d2220bf39eb2",
	  "version": 1,
	  "key": "random-key-12314",
	  "name": {
	    "en": "big store"
	  },
	  "createdAt": "1971-01-01T00:00:00.001Z",
	  "lastModifiedAt": "1972-01-01T00:00:00.001Z",
	  "distributionChannels": []
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	store, err := client.StoreDeleteWithID(context.TODO(), "dummy-id", 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, store)
	assert.NotNil(t, store.Version)
	assert.NotEmpty(t, store.LastModifiedAt)
	assert.NotEmpty(t, store.Key)
	assert.NotEmpty(t, store.ID)
	assert.NotEmpty(t, store.CreatedAt)

}

func TestGeneratedStoreDeleteWithKey(t *testing.T) {
	responseData := ` {
	  "id": "c2f93299-c967-44af-8c2a-d2220bf39eb2",
	  "version": 1,
	  "key": "random-key-12314",
	  "name": {
	    "en": "big store"
	  },
	  "createdAt": "1971-01-01T00:00:00.001Z",
	  "lastModifiedAt": "1972-01-01T00:00:00.001Z",
	  "distributionChannels": []
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	store, err := client.StoreDeleteWithKey(context.TODO(), "dummy-id", 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, store)
	assert.NotNil(t, store.Version)
	assert.NotEmpty(t, store.LastModifiedAt)
	assert.NotEmpty(t, store.Key)
	assert.NotEmpty(t, store.ID)
	assert.NotEmpty(t, store.CreatedAt)

}

func TestGeneratedStoreQuery(t *testing.T) {
	responseData := ` {
	  "limit": 20,
	  "offset": 0,
	  "count": 2,
	  "total": 2,
	  "results": [
	    {
	      "id": "c2f93299-c967-44af-8c2a-d2220bf39eb2",
	      "version": 1,
	      "key": "random-key-12314",
	      "name": {
	        "en": "big store"
	      },
	      "createdAt": "1971-01-01T00:00:00.001Z",
	      "lastModifiedAt": "1972-01-01T00:00:00.001Z",
	      "distributionChannels": []
	    },
	    {
	      "id": "c2f93209-c967-44af-8c2a-d2220bf39eb2",
	      "version": 1,
	      "key": "random-key-1234314",
	      "name": {
	        "en": "second store"
	      },
	      "createdAt": "1973-01-01T00:00:00.001Z",
	      "lastModifiedAt": "1973-01-01T00:00:00.001Z",
	      "distributionChannels": []
	    }
	  ]
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	input := commercetools.QueryInput{}
	queryResult, err := client.StoreQuery(context.TODO(), &input)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, queryResult)
	assert.NotNil(t, queryResult.Total)
	assert.NotNil(t, queryResult.Offset)
	assert.NotNil(t, queryResult.Limit)
	assert.NotNil(t, queryResult.Count)

}
