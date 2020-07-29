// Automatically generated, do not edit

package commercetools_test

import (
	"context"
	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratedTaxCategoryGetWithID(t *testing.T) {
	responseData := ` {
	  "id": "c60f7377-2643-4e99-adb5-b2909657444d",
	  "version": 1,
	  "name": "test-tax-category",
	  "rates": [
	    {
	      "name": "test-tax-category",
	      "amount": 0.2,
	      "includedInPrice": true,
	      "country": "DE",
	      "id": "vWTk7VjT",
	      "subRates": []
	    }
	  ],
	  "createdAt": "2016-02-24T15:33:40.811Z",
	  "lastModifiedAt": "2016-02-24T15:33:40.811Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	tax_category, err := client.TaxCategoryGetWithID(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, tax_category)
	assert.NotNil(t, tax_category.Version)
	assert.NotEmpty(t, tax_category.Name)
	assert.NotEmpty(t, tax_category.LastModifiedAt)
	assert.NotEmpty(t, tax_category.ID)
	assert.NotEmpty(t, tax_category.CreatedAt)

}

func TestGeneratedTaxCategoryGetWithKey(t *testing.T) {
	responseData := ` {
	  "id": "c60f7377-2643-4e99-adb5-b2909657444d",
	  "version": 1,
	  "name": "test-tax-category",
	  "rates": [
	    {
	      "name": "test-tax-category",
	      "amount": 0.2,
	      "includedInPrice": true,
	      "country": "DE",
	      "id": "vWTk7VjT",
	      "subRates": []
	    }
	  ],
	  "createdAt": "2016-02-24T15:33:40.811Z",
	  "lastModifiedAt": "2016-02-24T15:33:40.811Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	tax_category, err := client.TaxCategoryGetWithKey(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, tax_category)
	assert.NotNil(t, tax_category.Version)
	assert.NotEmpty(t, tax_category.Name)
	assert.NotEmpty(t, tax_category.LastModifiedAt)
	assert.NotEmpty(t, tax_category.ID)
	assert.NotEmpty(t, tax_category.CreatedAt)

}

func TestGeneratedTaxCategoryDeleteWithID(t *testing.T) {
	responseData := ` {
	  "id": "c60f7377-2643-4e99-adb5-b2909657444d",
	  "version": 1,
	  "name": "test-tax-category",
	  "rates": [
	    {
	      "name": "test-tax-category",
	      "amount": 0.2,
	      "includedInPrice": true,
	      "country": "DE",
	      "id": "vWTk7VjT",
	      "subRates": []
	    }
	  ],
	  "createdAt": "2016-02-24T15:33:40.811Z",
	  "lastModifiedAt": "2016-02-24T15:33:40.811Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	tax_category, err := client.TaxCategoryDeleteWithID(context.TODO(), "dummy-id", 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, tax_category)
	assert.NotNil(t, tax_category.Version)
	assert.NotEmpty(t, tax_category.Name)
	assert.NotEmpty(t, tax_category.LastModifiedAt)
	assert.NotEmpty(t, tax_category.ID)
	assert.NotEmpty(t, tax_category.CreatedAt)

}

func TestGeneratedTaxCategoryDeleteWithKey(t *testing.T) {
	responseData := ` {
	  "id": "c60f7377-2643-4e99-adb5-b2909657444d",
	  "version": 1,
	  "name": "test-tax-category",
	  "rates": [
	    {
	      "name": "test-tax-category",
	      "amount": 0.2,
	      "includedInPrice": true,
	      "country": "DE",
	      "id": "vWTk7VjT",
	      "subRates": []
	    }
	  ],
	  "createdAt": "2016-02-24T15:33:40.811Z",
	  "lastModifiedAt": "2016-02-24T15:33:40.811Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	tax_category, err := client.TaxCategoryDeleteWithKey(context.TODO(), "dummy-id", 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, tax_category)
	assert.NotNil(t, tax_category.Version)
	assert.NotEmpty(t, tax_category.Name)
	assert.NotEmpty(t, tax_category.LastModifiedAt)
	assert.NotEmpty(t, tax_category.ID)
	assert.NotEmpty(t, tax_category.CreatedAt)

}

func TestGeneratedTaxCategoryQuery(t *testing.T) {
	responseData := ` {
	  "limit": 20,
	  "offset": 0,
	  "count": 1,
	  "total": 1,
	  "results": [
	    {
	      "id": "c60f7377-2643-4e99-adb5-b2909657444d",
	      "version": 1,
	      "name": "test-tax-category",
	      "rates": [
	        {
	          "name": "test-tax-category",
	          "amount": 0.2,
	          "includedInPrice": true,
	          "country": "DE",
	          "id": "vWTk7VjT",
	          "subRates": []
	        }
	      ],
	      "createdAt": "2016-02-24T15:33:40.811Z",
	      "lastModifiedAt": "2016-02-24T15:33:40.811Z"
	    }
	  ]
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	input := commercetools.QueryInput{}
	queryResult, err := client.TaxCategoryQuery(context.TODO(), &input)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, queryResult)
	assert.NotNil(t, queryResult.Total)
	assert.NotNil(t, queryResult.Offset)
	assert.NotNil(t, queryResult.Limit)
	assert.NotNil(t, queryResult.Count)

}
