// Automatically generated, do not edit

package commercetools_test

import (
	"context"
	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratedCustomerGroupGetWithID(t *testing.T) {
	responseData := ` {
	  "id": "aef9cf41-94ad-4794-8122-62d308900430",
	  "version": 2,
	  "name": "Webshop user",
	  "createdAt": "2017-01-10T06:51:25.896Z",
	  "lastModifiedAt": "2017-01-10T06:51:25.946Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	customer_group, err := client.CustomerGroupGetWithID(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, customer_group)
	assert.NotNil(t, customer_group.Version)
	assert.NotEmpty(t, customer_group.Name)
	assert.NotEmpty(t, customer_group.LastModifiedAt)
	assert.NotEmpty(t, customer_group.ID)
	assert.NotEmpty(t, customer_group.CreatedAt)

}

func TestGeneratedCustomerGroupGetWithKey(t *testing.T) {
	responseData := ` {
	  "id": "aef9cf41-94ad-4794-8122-62d308900430",
	  "version": 2,
	  "name": "Webshop user",
	  "createdAt": "2017-01-10T06:51:25.896Z",
	  "lastModifiedAt": "2017-01-10T06:51:25.946Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	customer_group, err := client.CustomerGroupGetWithKey(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, customer_group)
	assert.NotNil(t, customer_group.Version)
	assert.NotEmpty(t, customer_group.Name)
	assert.NotEmpty(t, customer_group.LastModifiedAt)
	assert.NotEmpty(t, customer_group.ID)
	assert.NotEmpty(t, customer_group.CreatedAt)

}

func TestGeneratedCustomerGroupDeleteWithID(t *testing.T) {
	responseData := ` {
	  "id": "aef9cf41-94ad-4794-8122-62d308900430",
	  "version": 2,
	  "name": "Webshop user",
	  "createdAt": "2017-01-10T06:51:25.896Z",
	  "lastModifiedAt": "2017-01-10T06:51:25.946Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	customer_group, err := client.CustomerGroupDeleteWithID(context.TODO(), "dummy-id", 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, customer_group)
	assert.NotNil(t, customer_group.Version)
	assert.NotEmpty(t, customer_group.Name)
	assert.NotEmpty(t, customer_group.LastModifiedAt)
	assert.NotEmpty(t, customer_group.ID)
	assert.NotEmpty(t, customer_group.CreatedAt)

}

func TestGeneratedCustomerGroupDeleteWithKey(t *testing.T) {
	responseData := ` {
	  "id": "aef9cf41-94ad-4794-8122-62d308900430",
	  "version": 2,
	  "name": "Webshop user",
	  "createdAt": "2017-01-10T06:51:25.896Z",
	  "lastModifiedAt": "2017-01-10T06:51:25.946Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	customer_group, err := client.CustomerGroupDeleteWithKey(context.TODO(), "dummy-id", 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, customer_group)
	assert.NotNil(t, customer_group.Version)
	assert.NotEmpty(t, customer_group.Name)
	assert.NotEmpty(t, customer_group.LastModifiedAt)
	assert.NotEmpty(t, customer_group.ID)
	assert.NotEmpty(t, customer_group.CreatedAt)

}

func TestGeneratedCustomerGroupQuery(t *testing.T) {
	responseData := ` {
	  "limit": 20,
	  "offset": 0,
	  "count": 1,
	  "total": 1,
	  "results": [
	    {
	      "id": "aef9cf41-94ad-4794-8122-62d308900430",
	      "version": 2,
	      "name": "Webshop user",
	      "createdAt": "2017-01-10T06:51:25.896Z",
	      "lastModifiedAt": "2017-01-10T06:51:25.946Z"
	    }
	  ]
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	input := commercetools.QueryInput{}
	queryResult, err := client.CustomerGroupQuery(context.TODO(), &input)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, queryResult)
	assert.NotNil(t, queryResult.Total)
	assert.NotNil(t, queryResult.Offset)
	assert.NotNil(t, queryResult.Limit)
	assert.NotNil(t, queryResult.Count)

}
