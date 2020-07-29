// Automatically generated, do not edit

package commercetools_test

import (
	"context"
	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratedAPIClientGetWithID(t *testing.T) {
	responseData := ` {
	  "id": "api-client-id",
	  "name": "api-client-name",
	  "scope": "view_products",
	  "createdAt": "2018-01-01T00:00:00.001Z",
	  "lastUsedAt": "2017-09-10",
	  "secret": "secret-passphrase"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	api_client, err := client.APIClientGetWithID(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, api_client)

}

func TestGeneratedAPIClientDeleteWithID(t *testing.T) {
	responseData := ` {
	  "id": "api-client-id",
	  "name": "api-client-name",
	  "scope": "view_products",
	  "createdAt": "2018-01-01T00:00:00.001Z",
	  "lastUsedAt": "2017-09-10",
	  "secret": "secret-passphrase"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	api_client, err := client.APIClientDeleteWithID(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, api_client)

}

func TestGeneratedAPIClientQuery(t *testing.T) {
	responseData := ` {
	  "limit": 20,
	  "offset": 0,
	  "count": 1,
	  "total": 1,
	  "results": [
	    {
	      "id": "api-client-id",
	      "name": "api-client-name",
	      "scope": "view_products",
	      "createdAt": "2018-01-01T00:00:00.001Z",
	      "lastUsedAt": "2017-09-10",
	      "secret": "secret-passphrase"
	    }
	  ]
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	input := commercetools.QueryInput{}
	queryResult, err := client.APIClientQuery(context.TODO(), &input)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, queryResult)

}
