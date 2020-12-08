// Automatically generated, do not edit

package commercetools_test

import (
	"context"
	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratedCustomerGetWithEmailToken(t *testing.T) {
	responseData := ` {
	      "addresses": [],
	      "email": "<random>@example.com",
	      "firstName": "John",
	      "id": "some_123_id",
	      "isEmailVerified": false,
	      "lastName": "Doe",
	      "password": "secret123",
	      "version": 1,
	      "createdAt": "2015-07-06T13:22:33.339Z",
	      "lastModifiedAt": "2015-07-06T13:22:33.339Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	customer, err := client.CustomerGetWithEmailToken(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, customer)
	assert.NotNil(t, customer.Version)
	assert.NotEmpty(t, customer.Password)
	assert.NotEmpty(t, customer.LastName)
	assert.NotEmpty(t, customer.LastModifiedAt)
	assert.False(t, customer.IsEmailVerified)
	assert.NotEmpty(t, customer.ID)
	assert.NotEmpty(t, customer.FirstName)
	assert.NotEmpty(t, customer.Email)
	assert.NotEmpty(t, customer.CreatedAt)

}

func TestGeneratedCustomerGetWithID(t *testing.T) {
	responseData := ` { "customer": {
	      "addresses": [],
	      "email": "<random>@example.com",
	      "firstName": "John",
	      "id": "some_123_id",
	      "isEmailVerified": false,
	      "lastName": "Doe",
	      "password": "secret123",
	      "version": 1,
	      "createdAt": "2015-07-06T13:22:33.339Z",
	      "lastModifiedAt": "2015-07-06T13:22:33.339Z"
	} }`
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	customer, err := client.CustomerGetWithID(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, customer)
	assert.NotNil(t, customer.Customer.Version)
	assert.NotEmpty(t, customer.Customer.Password)
	assert.NotEmpty(t, customer.Customer.LastName)
	assert.NotEmpty(t, customer.Customer.LastModifiedAt)
	assert.False(t, customer.Customer.IsEmailVerified)
	assert.NotEmpty(t, customer.Customer.ID)
	assert.NotEmpty(t, customer.Customer.FirstName)
	assert.NotEmpty(t, customer.Customer.Email)
	assert.NotEmpty(t, customer.Customer.CreatedAt)

}

func TestGeneratedCustomerGetWithKey(t *testing.T) {
	responseData := ` {
	      "addresses": [],
	      "email": "<random>@example.com",
	      "firstName": "John",
	      "id": "some_123_id",
	      "isEmailVerified": false,
	      "lastName": "Doe",
	      "password": "secret123",
	      "version": 1,
	      "createdAt": "2015-07-06T13:22:33.339Z",
	      "lastModifiedAt": "2015-07-06T13:22:33.339Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	customer, err := client.CustomerGetWithKey(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, customer)
	assert.NotNil(t, customer.Version)
	assert.NotEmpty(t, customer.Password)
	assert.NotEmpty(t, customer.LastName)
	assert.NotEmpty(t, customer.LastModifiedAt)
	assert.False(t, customer.IsEmailVerified)
	assert.NotEmpty(t, customer.ID)
	assert.NotEmpty(t, customer.FirstName)
	assert.NotEmpty(t, customer.Email)
	assert.NotEmpty(t, customer.CreatedAt)

}

func TestGeneratedCustomerGetWithPasswordToken(t *testing.T) {
	responseData := ` {
	      "addresses": [],
	      "email": "<random>@example.com",
	      "firstName": "John",
	      "id": "some_123_id",
	      "isEmailVerified": false,
	      "lastName": "Doe",
	      "password": "secret123",
	      "version": 1,
	      "createdAt": "2015-07-06T13:22:33.339Z",
	      "lastModifiedAt": "2015-07-06T13:22:33.339Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	customer, err := client.CustomerGetWithPasswordToken(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, customer)
	assert.NotNil(t, customer.Version)
	assert.NotEmpty(t, customer.Password)
	assert.NotEmpty(t, customer.LastName)
	assert.NotEmpty(t, customer.LastModifiedAt)
	assert.False(t, customer.IsEmailVerified)
	assert.NotEmpty(t, customer.ID)
	assert.NotEmpty(t, customer.FirstName)
	assert.NotEmpty(t, customer.Email)
	assert.NotEmpty(t, customer.CreatedAt)

}

func TestGeneratedCustomerDeleteWithID(t *testing.T) {
	responseData := ` {
	      "addresses": [],
	      "email": "<random>@example.com",
	      "firstName": "John",
	      "id": "some_123_id",
	      "isEmailVerified": false,
	      "lastName": "Doe",
	      "password": "secret123",
	      "version": 1,
	      "createdAt": "2015-07-06T13:22:33.339Z",
	      "lastModifiedAt": "2015-07-06T13:22:33.339Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	customer, err := client.CustomerDeleteWithID(context.TODO(), "dummy-id", 1, true)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, customer)
	assert.NotNil(t, customer.Version)
	assert.NotEmpty(t, customer.Password)
	assert.NotEmpty(t, customer.LastName)
	assert.NotEmpty(t, customer.LastModifiedAt)
	assert.False(t, customer.IsEmailVerified)
	assert.NotEmpty(t, customer.ID)
	assert.NotEmpty(t, customer.FirstName)
	assert.NotEmpty(t, customer.Email)
	assert.NotEmpty(t, customer.CreatedAt)

}

func TestGeneratedCustomerDeleteWithKey(t *testing.T) {
	responseData := ` {
	      "addresses": [],
	      "email": "<random>@example.com",
	      "firstName": "John",
	      "id": "some_123_id",
	      "isEmailVerified": false,
	      "lastName": "Doe",
	      "password": "secret123",
	      "version": 1,
	      "createdAt": "2015-07-06T13:22:33.339Z",
	      "lastModifiedAt": "2015-07-06T13:22:33.339Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	customer, err := client.CustomerDeleteWithKey(context.TODO(), "dummy-id", 1, true)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, customer)
	assert.NotNil(t, customer.Version)
	assert.NotEmpty(t, customer.Password)
	assert.NotEmpty(t, customer.LastName)
	assert.NotEmpty(t, customer.LastModifiedAt)
	assert.False(t, customer.IsEmailVerified)
	assert.NotEmpty(t, customer.ID)
	assert.NotEmpty(t, customer.FirstName)
	assert.NotEmpty(t, customer.Email)
	assert.NotEmpty(t, customer.CreatedAt)

}

func TestGeneratedCustomerQuery(t *testing.T) {
	responseData := ` {
	      "limit": 20,
	      "offset": 0,
	      "count": 1,
	      "total": 1,
	      "results": [{
	            "addresses": [],
	            "email": "<random>@example.com",
	            "firstName": "John",
	            "id": "some_123_id",
	            "isEmailVerified": false,
	            "lastName": "Doe",
	            "password": "secret123",
	            "version": 1,
	            "createdAt": "2015-07-06T13:22:33.339Z",
	            "lastModifiedAt": "2015-07-06T13:22:33.339Z"
	      }]
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	input := commercetools.QueryInput{}
	queryResult, err := client.CustomerQuery(context.TODO(), &input)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, queryResult)
	assert.NotNil(t, queryResult.Total)
	assert.NotNil(t, queryResult.Offset)
	assert.NotNil(t, queryResult.Limit)
	assert.NotNil(t, queryResult.Count)

}
