// Automatically generated, do not edit

package commercetools_test

import (
	"context"
	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratedExtensionGetWithID(t *testing.T) {
	responseData := ` {
	  "id": "8062243c-46fc-40b5-88a4-75e2216aef75",
	  "version": 1,
	  "createdAt": "2017-01-25T14:14:22.417Z",
	  "lastModifiedAt": "2017-01-25T14:14:22.417Z",
	  "destination": {
	    "type": "HTTP",
	    "url": "https://example.azurewebsites.net/api/extension",
	    "authentication": {
	      "type": "AzureFunctions",
	      "key": "some-azure-function-code"
	    }
	  },
	  "triggers": [{
	    "resourceTypeId": "cart",
	    "actions": ["Create", "Update"]
	  }],
	  "key": "my-extension"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	extension, err := client.ExtensionGetWithID(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, extension)
	assert.NotNil(t, extension.Version)
	assert.NotEmpty(t, extension.LastModifiedAt)
	assert.NotEmpty(t, extension.Key)
	assert.NotEmpty(t, extension.ID)
	assert.NotEmpty(t, extension.CreatedAt)

}

func TestGeneratedExtensionGetWithKey(t *testing.T) {
	responseData := ` {
	  "id": "8062243c-46fc-40b5-88a4-75e2216aef75",
	  "version": 1,
	  "createdAt": "2017-01-25T14:14:22.417Z",
	  "lastModifiedAt": "2017-01-25T14:14:22.417Z",
	  "destination": {
	    "type": "HTTP",
	    "url": "https://example.azurewebsites.net/api/extension",
	    "authentication": {
	      "type": "AzureFunctions",
	      "key": "some-azure-function-code"
	    }
	  },
	  "triggers": [{
	    "resourceTypeId": "cart",
	    "actions": ["Create", "Update"]
	  }],
	  "key": "my-extension"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	extension, err := client.ExtensionGetWithKey(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, extension)
	assert.NotNil(t, extension.Version)
	assert.NotEmpty(t, extension.LastModifiedAt)
	assert.NotEmpty(t, extension.Key)
	assert.NotEmpty(t, extension.ID)
	assert.NotEmpty(t, extension.CreatedAt)

}

func TestGeneratedExtensionDeleteWithID(t *testing.T) {
	responseData := ` {
	  "id": "8062243c-46fc-40b5-88a4-75e2216aef75",
	  "version": 1,
	  "createdAt": "2017-01-25T14:14:22.417Z",
	  "lastModifiedAt": "2017-01-25T14:14:22.417Z",
	  "destination": {
	    "type": "HTTP",
	    "url": "https://example.azurewebsites.net/api/extension",
	    "authentication": {
	      "type": "AzureFunctions",
	      "key": "some-azure-function-code"
	    }
	  },
	  "triggers": [{
	    "resourceTypeId": "cart",
	    "actions": ["Create", "Update"]
	  }],
	  "key": "my-extension"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	extension, err := client.ExtensionDeleteWithID(context.TODO(), "dummy-id", 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, extension)
	assert.NotNil(t, extension.Version)
	assert.NotEmpty(t, extension.LastModifiedAt)
	assert.NotEmpty(t, extension.Key)
	assert.NotEmpty(t, extension.ID)
	assert.NotEmpty(t, extension.CreatedAt)

}

func TestGeneratedExtensionDeleteWithKey(t *testing.T) {
	responseData := ` {
	  "id": "8062243c-46fc-40b5-88a4-75e2216aef75",
	  "version": 1,
	  "createdAt": "2017-01-25T14:14:22.417Z",
	  "lastModifiedAt": "2017-01-25T14:14:22.417Z",
	  "destination": {
	    "type": "HTTP",
	    "url": "https://example.azurewebsites.net/api/extension",
	    "authentication": {
	      "type": "AzureFunctions",
	      "key": "some-azure-function-code"
	    }
	  },
	  "triggers": [{
	    "resourceTypeId": "cart",
	    "actions": ["Create", "Update"]
	  }],
	  "key": "my-extension"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	extension, err := client.ExtensionDeleteWithKey(context.TODO(), "dummy-id", 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, extension)
	assert.NotNil(t, extension.Version)
	assert.NotEmpty(t, extension.LastModifiedAt)
	assert.NotEmpty(t, extension.Key)
	assert.NotEmpty(t, extension.ID)
	assert.NotEmpty(t, extension.CreatedAt)

}

func TestGeneratedExtensionQuery(t *testing.T) {
	responseData := ` {
	  "limit": 20,
	  "offset": 0,
	  "count": 1,
	  "total": 1,
	  "results": [
	    {
	      "id": "8062243c-46fc-40b5-88a4-75e2216aef75",
	      "version": 1,
	      "createdAt": "2017-01-25T14:14:22.417Z",
	      "lastModifiedAt": "2017-01-25T14:14:22.417Z",
	      "destination": {
	        "type": "HTTP",
	        "url": "https://example.org/extension"
	      },
	      "triggers": [{
	        "resourceTypeId": "cart",
	        "actions": ["Create"]
	      }]
	    }
	  ]
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	input := commercetools.QueryInput{}
	queryResult, err := client.ExtensionQuery(context.TODO(), &input)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, queryResult)
	assert.NotNil(t, queryResult.Total)
	assert.NotNil(t, queryResult.Offset)
	assert.NotNil(t, queryResult.Limit)
	assert.NotNil(t, queryResult.Count)

}
