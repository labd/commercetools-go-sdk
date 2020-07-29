// Automatically generated, do not edit

package commercetools_test

import (
	"context"
	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratedCategoryGetWithID(t *testing.T) {
	responseData := ` {
	  "id": "c2f93298-c967-44af-8c2a-d2220bf39eb2",
	  "version": 1,
	  "name": {
	    "en": "Hats"
	  },
	  "slug": {
	    "en": "hats"
	  },
	  "parent": {
	    "typeId": "category",
	    "id": "123456"
	  },
	  "ancestors": [],
	  "orderHint": "0.000013833188755841543727840",
	  "createdAt": "1970-01-01T00:00:00.001Z",
	  "lastModifiedAt": "1970-01-01T00:00:00.001Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	category, err := client.CategoryGetWithID(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, category)
	assert.NotNil(t, category.Version)
	assert.NotEmpty(t, category.OrderHint)
	assert.NotEmpty(t, category.LastModifiedAt)
	assert.NotEmpty(t, category.ID)
	assert.NotEmpty(t, category.CreatedAt)

}

func TestGeneratedCategoryGetWithKey(t *testing.T) {
	responseData := ` {
	  "id": "c2f93298-c967-44af-8c2a-d2220bf39eb2",
	  "version": 1,
	  "name": {
	    "en": "Hats"
	  },
	  "slug": {
	    "en": "hats"
	  },
	  "parent": {
	    "typeId": "category",
	    "id": "123456"
	  },
	  "ancestors": [],
	  "orderHint": "0.000013833188755841543727840",
	  "createdAt": "1970-01-01T00:00:00.001Z",
	  "lastModifiedAt": "1970-01-01T00:00:00.001Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	category, err := client.CategoryGetWithKey(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, category)
	assert.NotNil(t, category.Version)
	assert.NotEmpty(t, category.OrderHint)
	assert.NotEmpty(t, category.LastModifiedAt)
	assert.NotEmpty(t, category.ID)
	assert.NotEmpty(t, category.CreatedAt)

}

func TestGeneratedCategoryDeleteWithID(t *testing.T) {
	responseData := ` {
	  "id": "c2f93298-c967-44af-8c2a-d2220bf39eb2",
	  "version": 1,
	  "name": {
	    "en": "Hats"
	  },
	  "slug": {
	    "en": "hats"
	  },
	  "parent": {
	    "typeId": "category",
	    "id": "123456"
	  },
	  "ancestors": [],
	  "orderHint": "0.000013833188755841543727840",
	  "createdAt": "1970-01-01T00:00:00.001Z",
	  "lastModifiedAt": "1970-01-01T00:00:00.001Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	category, err := client.CategoryDeleteWithID(context.TODO(), "dummy-id", 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, category)
	assert.NotNil(t, category.Version)
	assert.NotEmpty(t, category.OrderHint)
	assert.NotEmpty(t, category.LastModifiedAt)
	assert.NotEmpty(t, category.ID)
	assert.NotEmpty(t, category.CreatedAt)

}

func TestGeneratedCategoryDeleteWithKey(t *testing.T) {
	responseData := ` {
	  "id": "c2f93298-c967-44af-8c2a-d2220bf39eb2",
	  "version": 1,
	  "name": {
	    "en": "Hats"
	  },
	  "slug": {
	    "en": "hats"
	  },
	  "parent": {
	    "typeId": "category",
	    "id": "123456"
	  },
	  "ancestors": [],
	  "orderHint": "0.000013833188755841543727840",
	  "createdAt": "1970-01-01T00:00:00.001Z",
	  "lastModifiedAt": "1970-01-01T00:00:00.001Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	category, err := client.CategoryDeleteWithKey(context.TODO(), "dummy-id", 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, category)
	assert.NotNil(t, category.Version)
	assert.NotEmpty(t, category.OrderHint)
	assert.NotEmpty(t, category.LastModifiedAt)
	assert.NotEmpty(t, category.ID)
	assert.NotEmpty(t, category.CreatedAt)

}

func TestGeneratedCategoryQuery(t *testing.T) {
	responseData := ` {
	  "limit": 20,
	  "offset": 0,
	  "count": 2,
	  "total": 2,
	  "results": [
	    {
	      "id": "c2f93298-c967-44af-8c2a-d2220bf39eb2",
	      "version": 1,
	      "name": {
	        "en": "Hats"
	      },
	      "slug": {
	        "en": "hats"
	      },
	      "ancestors": [{
	        "typeId": "category",
	        "id": "123456"
	      }],
	      "orderHint": "0.000013833188755841543727840",
	      "createdAt": "1970-01-01T00:00:00.001Z",
	      "lastModifiedAt": "1970-01-01T00:00:00.001Z"
	    },
	    {
	      "id": "1bae3aa3-1e02-49d2-b719-4c5020f50638",
	      "version": 1,
	      "name": {
	        "en": "Long sleeves"
	      },
	      "slug": {
	        "en": "long-sleeves"
	      },
	      "ancestors": [],
	      "orderHint": "0.000013833188755781786424864",
	      "createdAt": "1970-01-01T00:00:00.001Z",
	      "lastModifiedAt": "1970-01-01T00:00:00.001Z"
	    }
	  ]
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	input := commercetools.QueryInput{}
	queryResult, err := client.CategoryQuery(context.TODO(), &input)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, queryResult)
	assert.NotNil(t, queryResult.Total)
	assert.NotNil(t, queryResult.Offset)
	assert.NotNil(t, queryResult.Limit)
	assert.NotNil(t, queryResult.Count)

}
