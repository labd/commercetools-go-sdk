// Automatically generated, do not edit

package commercetools_test

import (
	"context"
	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratedShoppingListGetWithID(t *testing.T) {
	responseData := ` {
	  "id": "9693f04b-5aec-467f-baa1-fc74da7d0c3d",
	  "version": 1,
	  "name": {
	    "en": "test"
	  },
	  "key": "test",
	  "lineItems": [],
	  "textLineItems": [],
	  "createdAt": "2017-03-30T11:49:40.904Z",
	  "lastModifiedAt": "2017-03-30T11:49:40.904Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	shopping_list, err := client.ShoppingListGetWithID(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, shopping_list)
	assert.NotNil(t, shopping_list.Version)
	assert.NotEmpty(t, shopping_list.LastModifiedAt)
	assert.NotEmpty(t, shopping_list.Key)
	assert.NotEmpty(t, shopping_list.ID)
	assert.NotEmpty(t, shopping_list.CreatedAt)

}

func TestGeneratedShoppingListGetWithKey(t *testing.T) {
	responseData := ` {
	  "id": "9693f04b-5aec-467f-baa1-fc74da7d0c3d",
	  "version": 1,
	  "name": {
	    "en": "test"
	  },
	  "key": "test",
	  "lineItems": [],
	  "textLineItems": [],
	  "createdAt": "2017-03-30T11:49:40.904Z",
	  "lastModifiedAt": "2017-03-30T11:49:40.904Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	shopping_list, err := client.ShoppingListGetWithKey(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, shopping_list)
	assert.NotNil(t, shopping_list.Version)
	assert.NotEmpty(t, shopping_list.LastModifiedAt)
	assert.NotEmpty(t, shopping_list.Key)
	assert.NotEmpty(t, shopping_list.ID)
	assert.NotEmpty(t, shopping_list.CreatedAt)

}

func TestGeneratedShoppingListDeleteWithID(t *testing.T) {
	responseData := ` {
	  "id": "9693f04b-5aec-467f-baa1-fc74da7d0c3d",
	  "version": 1,
	  "name": {
	    "en": "test"
	  },
	  "key": "test",
	  "lineItems": [],
	  "textLineItems": [],
	  "createdAt": "2017-03-30T11:49:40.904Z",
	  "lastModifiedAt": "2017-03-30T11:49:40.904Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	shopping_list, err := client.ShoppingListDeleteWithID(context.TODO(), "dummy-id", 1, true)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, shopping_list)
	assert.NotNil(t, shopping_list.Version)
	assert.NotEmpty(t, shopping_list.LastModifiedAt)
	assert.NotEmpty(t, shopping_list.Key)
	assert.NotEmpty(t, shopping_list.ID)
	assert.NotEmpty(t, shopping_list.CreatedAt)

}

func TestGeneratedShoppingListDeleteWithKey(t *testing.T) {
	responseData := ` {
	  "id": "9693f04b-5aec-467f-baa1-fc74da7d0c3d",
	  "version": 1,
	  "name": {
	    "en": "test"
	  },
	  "key": "test",
	  "lineItems": [],
	  "textLineItems": [],
	  "createdAt": "2017-03-30T11:49:40.904Z",
	  "lastModifiedAt": "2017-03-30T11:49:40.904Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	shopping_list, err := client.ShoppingListDeleteWithKey(context.TODO(), "dummy-id", 1, true)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, shopping_list)
	assert.NotNil(t, shopping_list.Version)
	assert.NotEmpty(t, shopping_list.LastModifiedAt)
	assert.NotEmpty(t, shopping_list.Key)
	assert.NotEmpty(t, shopping_list.ID)
	assert.NotEmpty(t, shopping_list.CreatedAt)

}

func TestGeneratedShoppingListQuery(t *testing.T) {
	responseData := ` {
	  "limit": 20,
	  "offset": 0,
	  "count": 1,
	  "total": 1,
	  "results": [
	    {
	      "id": "9693f04b-5aec-467f-baa1-fc74da7d0c3d",
	      "version": 1,
	      "name": {
	        "en": "test"
	      },
	      "key": "test",
	      "lineItems": [],
	      "textLineItems": [],
	      "createdAt": "2017-03-30T11:49:40.904Z",
	      "lastModifiedAt": "2017-03-30T11:49:40.904Z"
	    }
	  ]
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	input := commercetools.QueryInput{}
	queryResult, err := client.ShoppingListQuery(context.TODO(), &input)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, queryResult)
	assert.NotNil(t, queryResult.Total)
	assert.NotNil(t, queryResult.Offset)
	assert.NotNil(t, queryResult.Limit)
	assert.NotNil(t, queryResult.Count)

}
