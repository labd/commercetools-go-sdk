// Automatically generated, do not edit

package commercetools_test

import (
	"context"
	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratedProductTypeGetWithID(t *testing.T) {
	responseData := ` {
	  "id": "c2f93298-c967-44af-8c2a-d2220bf39eb2",
	  "version": 1,
	  "name": "test_product_type",
	  "description": "Test product type.",
	  "attributes": [{
	    "type": {
	      "name": "text"
	    },
	    "isSearchable": false,
	    "inputHint": "SingleLine",
	    "name": "size",
	    "label": {
	      "en": "The right size is important."
	    },
	    "isRequired": false,
	    "attributeConstraint": "CombinationUnique"
	  }],
	  "createdAt": "1970-01-01T00:00:00.001Z",
	  "lastModifiedAt": "1970-01-01T00:00:00.001Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	product_type, err := client.ProductTypeGetWithID(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, product_type)
	assert.NotNil(t, product_type.Version)
	assert.NotEmpty(t, product_type.Name)
	assert.NotEmpty(t, product_type.LastModifiedAt)
	assert.NotEmpty(t, product_type.ID)
	assert.NotEmpty(t, product_type.Description)
	assert.NotEmpty(t, product_type.CreatedAt)

}

func TestGeneratedProductTypeGetWithKey(t *testing.T) {
	responseData := ` {
	  "id": "c2f93298-c967-44af-8c2a-d2220bf39eb2",
	  "version": 1,
	  "name": "test_product_type",
	  "description": "Test product type.",
	  "attributes": [{
	    "type": {
	      "name": "text"
	    },
	    "isSearchable": false,
	    "inputHint": "SingleLine",
	    "name": "size",
	    "label": {
	      "en": "The right size is important."
	    },
	    "isRequired": false,
	    "attributeConstraint": "CombinationUnique"
	  }],
	  "createdAt": "1970-01-01T00:00:00.001Z",
	  "lastModifiedAt": "1970-01-01T00:00:00.001Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	product_type, err := client.ProductTypeGetWithKey(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, product_type)
	assert.NotNil(t, product_type.Version)
	assert.NotEmpty(t, product_type.Name)
	assert.NotEmpty(t, product_type.LastModifiedAt)
	assert.NotEmpty(t, product_type.ID)
	assert.NotEmpty(t, product_type.Description)
	assert.NotEmpty(t, product_type.CreatedAt)

}

func TestGeneratedProductTypeDeleteWithID(t *testing.T) {
	responseData := ` {
	  "id": "c2f93298-c967-44af-8c2a-d2220bf39eb2",
	  "version": 1,
	  "name": "test_product_type",
	  "description": "Test product type.",
	  "attributes": [{
	    "type": {
	      "name": "text"
	    },
	    "isSearchable": false,
	    "inputHint": "SingleLine",
	    "name": "size",
	    "label": {
	      "en": "The right size is important."
	    },
	    "isRequired": false,
	    "attributeConstraint": "CombinationUnique"
	  }],
	  "createdAt": "1970-01-01T00:00:00.001Z",
	  "lastModifiedAt": "1970-01-01T00:00:00.001Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	product_type, err := client.ProductTypeDeleteWithID(context.TODO(), "dummy-id", 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, product_type)
	assert.NotNil(t, product_type.Version)
	assert.NotEmpty(t, product_type.Name)
	assert.NotEmpty(t, product_type.LastModifiedAt)
	assert.NotEmpty(t, product_type.ID)
	assert.NotEmpty(t, product_type.Description)
	assert.NotEmpty(t, product_type.CreatedAt)

}

func TestGeneratedProductTypeDeleteWithKey(t *testing.T) {
	responseData := ` {
	  "id": "c2f93298-c967-44af-8c2a-d2220bf39eb2",
	  "version": 1,
	  "name": "test_product_type",
	  "description": "Test product type.",
	  "attributes": [{
	    "type": {
	      "name": "text"
	    },
	    "isSearchable": false,
	    "inputHint": "SingleLine",
	    "name": "size",
	    "label": {
	      "en": "The right size is important."
	    },
	    "isRequired": false,
	    "attributeConstraint": "CombinationUnique"
	  }],
	  "createdAt": "1970-01-01T00:00:00.001Z",
	  "lastModifiedAt": "1970-01-01T00:00:00.001Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	product_type, err := client.ProductTypeDeleteWithKey(context.TODO(), "dummy-id", 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, product_type)
	assert.NotNil(t, product_type.Version)
	assert.NotEmpty(t, product_type.Name)
	assert.NotEmpty(t, product_type.LastModifiedAt)
	assert.NotEmpty(t, product_type.ID)
	assert.NotEmpty(t, product_type.Description)
	assert.NotEmpty(t, product_type.CreatedAt)

}

func TestGeneratedProductTypeQuery(t *testing.T) {
	responseData := ` {
	  "limit": 20,
	  "offset": 0,
	  "count": 2,
	  "total": 2,
	  "results": [
	    {
	      "id": "d2220bf39eb2-c2f93298-c967-44af-8c2a",
	      "version": 1,
	      "name": "test_product_type",
	      "description": "Test product type.",
	      "attributes": [{
	        "type": {
	          "name": "text"
	        },
	        "isSearchable": false,
	        "inputHint": "SingleLine",
	        "name": "size",
	        "label": {
	          "en": "The right size is important."
	        },
	        "isRequired": false,
	        "attributeConstraint": "CombinationUnique"
	      }],
	      "createdAt": "1970-01-01T00:00:00.001Z",
	      "lastModifiedAt": "1970-01-01T00:00:00.001Z"
	    },
	    {
	      "id": "c2f93298-c967-44af-8c2a-d2220bf39eb2",
	      "version": 1,
	      "name": "another-test_product_type",
	      "description": "Another test product type.",
	      "attributes": [{
	        "type": {
	          "name": "text"
	        },
	        "isSearchable": false,
	        "inputHint": "SingleLine",
	        "name": "color",
	        "label": {
	          "en": "The right color is important."
	        },
	        "isRequired": false,
	        "attributeConstraint": "CombinationUnique"
	      }],
	      "createdAt": "1970-01-01T00:00:00.001Z",
	      "lastModifiedAt": "1970-01-01T00:00:00.001Z"
	    }
	  ]
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	input := commercetools.QueryInput{}
	queryResult, err := client.ProductTypeQuery(context.TODO(), &input)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, queryResult)
	assert.NotNil(t, queryResult.Total)
	assert.NotNil(t, queryResult.Offset)
	assert.NotNil(t, queryResult.Limit)
	assert.NotNil(t, queryResult.Count)

}
