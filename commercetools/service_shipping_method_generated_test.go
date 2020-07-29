// Automatically generated, do not edit

package commercetools_test

import (
	"context"
	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratedShippingMethodGetWithID(t *testing.T) {
	responseData := ` {
	  "id": "eb8991df-2dcd-4e24-83fe-5df46ec04422",
	  "version": 3,
	  "name": "DHL",
	  "description": "Standard delivery",
	  "taxCategory": {
	    "typeId": "tax-category",
	    "id": "5a21f15b-34f8-4b7f-9407-d1eb82a73eba"
	  },
	  "zoneRates": [
	    {
	      "zone": {
	        "typeId": "zone",
	        "id": "5cb532be-c680-43ab-ba14-b664bb03dc35"
	      },
	      "shippingRates": [
	        {
	          "price": {
	            "type": "centPrecision",
	            "fractionDigits": 2,
	            "currencyCode": "EUR",
	            "centAmount": 570
	          },
	          "tiers": []
	        }
	      ]
	    },
	    {
	      "zone": {
	        "typeId": "zone",
	        "id": "ebe01381-82be-4e63-9993-d1eb8f8e588b"
	      },
	      "shippingRates": [
	        {
	          "price": {
	            "type": "centPrecision",
	            "fractionDigits": 2,
	            "currencyCode": "USD",
	            "centAmount": 990
	          },
	          "tiers": []
	        }
	      ]
	    }
	  ],
	  "isDefault": false,
	  "createdAt": "2015-01-21T09:22:04.320Z",
	  "lastModifiedAt": "2016-02-24T13:36:56.748Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	shipping_method, err := client.ShippingMethodGetWithID(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, shipping_method)
	assert.NotNil(t, shipping_method.Version)
	assert.NotEmpty(t, shipping_method.Name)
	assert.NotEmpty(t, shipping_method.LastModifiedAt)
	assert.False(t, shipping_method.IsDefault)
	assert.NotEmpty(t, shipping_method.ID)
	assert.NotEmpty(t, shipping_method.Description)
	assert.NotEmpty(t, shipping_method.CreatedAt)

}

func TestGeneratedShippingMethodGetWithKey(t *testing.T) {
	responseData := ` {
	  "id": "eb8991df-2dcd-4e24-83fe-5df46ec04422",
	  "version": 3,
	  "name": "DHL",
	  "description": "Standard delivery",
	  "taxCategory": {
	    "typeId": "tax-category",
	    "id": "5a21f15b-34f8-4b7f-9407-d1eb82a73eba"
	  },
	  "zoneRates": [
	    {
	      "zone": {
	        "typeId": "zone",
	        "id": "5cb532be-c680-43ab-ba14-b664bb03dc35"
	      },
	      "shippingRates": [
	        {
	          "price": {
	            "type": "centPrecision",
	            "fractionDigits": 2,
	            "currencyCode": "EUR",
	            "centAmount": 570
	          },
	          "tiers": []
	        }
	      ]
	    },
	    {
	      "zone": {
	        "typeId": "zone",
	        "id": "ebe01381-82be-4e63-9993-d1eb8f8e588b"
	      },
	      "shippingRates": [
	        {
	          "price": {
	            "type": "centPrecision",
	            "fractionDigits": 2,
	            "currencyCode": "USD",
	            "centAmount": 990
	          },
	          "tiers": []
	        }
	      ]
	    }
	  ],
	  "isDefault": false,
	  "createdAt": "2015-01-21T09:22:04.320Z",
	  "lastModifiedAt": "2016-02-24T13:36:56.748Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	shipping_method, err := client.ShippingMethodGetWithKey(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, shipping_method)
	assert.NotNil(t, shipping_method.Version)
	assert.NotEmpty(t, shipping_method.Name)
	assert.NotEmpty(t, shipping_method.LastModifiedAt)
	assert.False(t, shipping_method.IsDefault)
	assert.NotEmpty(t, shipping_method.ID)
	assert.NotEmpty(t, shipping_method.Description)
	assert.NotEmpty(t, shipping_method.CreatedAt)

}

func TestGeneratedShippingMethodDeleteWithID(t *testing.T) {
	responseData := ` {
	  "id": "eb8991df-2dcd-4e24-83fe-5df46ec04422",
	  "version": 3,
	  "name": "DHL",
	  "description": "Standard delivery",
	  "taxCategory": {
	    "typeId": "tax-category",
	    "id": "5a21f15b-34f8-4b7f-9407-d1eb82a73eba"
	  },
	  "zoneRates": [
	    {
	      "zone": {
	        "typeId": "zone",
	        "id": "5cb532be-c680-43ab-ba14-b664bb03dc35"
	      },
	      "shippingRates": [
	        {
	          "price": {
	            "type": "centPrecision",
	            "fractionDigits": 2,
	            "currencyCode": "EUR",
	            "centAmount": 570
	          },
	          "tiers": []
	        }
	      ]
	    },
	    {
	      "zone": {
	        "typeId": "zone",
	        "id": "ebe01381-82be-4e63-9993-d1eb8f8e588b"
	      },
	      "shippingRates": [
	        {
	          "price": {
	            "type": "centPrecision",
	            "fractionDigits": 2,
	            "currencyCode": "USD",
	            "centAmount": 990
	          },
	          "tiers": []
	        }
	      ]
	    }
	  ],
	  "isDefault": false,
	  "createdAt": "2015-01-21T09:22:04.320Z",
	  "lastModifiedAt": "2016-02-24T13:36:56.748Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	shipping_method, err := client.ShippingMethodDeleteWithID(context.TODO(), "dummy-id", 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, shipping_method)
	assert.NotNil(t, shipping_method.Version)
	assert.NotEmpty(t, shipping_method.Name)
	assert.NotEmpty(t, shipping_method.LastModifiedAt)
	assert.False(t, shipping_method.IsDefault)
	assert.NotEmpty(t, shipping_method.ID)
	assert.NotEmpty(t, shipping_method.Description)
	assert.NotEmpty(t, shipping_method.CreatedAt)

}

func TestGeneratedShippingMethodDeleteWithKey(t *testing.T) {
	responseData := ` {
	  "id": "eb8991df-2dcd-4e24-83fe-5df46ec04422",
	  "version": 3,
	  "name": "DHL",
	  "description": "Standard delivery",
	  "taxCategory": {
	    "typeId": "tax-category",
	    "id": "5a21f15b-34f8-4b7f-9407-d1eb82a73eba"
	  },
	  "zoneRates": [
	    {
	      "zone": {
	        "typeId": "zone",
	        "id": "5cb532be-c680-43ab-ba14-b664bb03dc35"
	      },
	      "shippingRates": [
	        {
	          "price": {
	            "type": "centPrecision",
	            "fractionDigits": 2,
	            "currencyCode": "EUR",
	            "centAmount": 570
	          },
	          "tiers": []
	        }
	      ]
	    },
	    {
	      "zone": {
	        "typeId": "zone",
	        "id": "ebe01381-82be-4e63-9993-d1eb8f8e588b"
	      },
	      "shippingRates": [
	        {
	          "price": {
	            "type": "centPrecision",
	            "fractionDigits": 2,
	            "currencyCode": "USD",
	            "centAmount": 990
	          },
	          "tiers": []
	        }
	      ]
	    }
	  ],
	  "isDefault": false,
	  "createdAt": "2015-01-21T09:22:04.320Z",
	  "lastModifiedAt": "2016-02-24T13:36:56.748Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	shipping_method, err := client.ShippingMethodDeleteWithKey(context.TODO(), "dummy-id", 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, shipping_method)
	assert.NotNil(t, shipping_method.Version)
	assert.NotEmpty(t, shipping_method.Name)
	assert.NotEmpty(t, shipping_method.LastModifiedAt)
	assert.False(t, shipping_method.IsDefault)
	assert.NotEmpty(t, shipping_method.ID)
	assert.NotEmpty(t, shipping_method.Description)
	assert.NotEmpty(t, shipping_method.CreatedAt)

}

func TestGeneratedShippingMethodQuery(t *testing.T) {
	responseData := ` {
	  "limit": 20,
	  "offset": 0,
	  "count": 1,
	  "total": 1,
	  "results": [
	    {
	      "id": "eb8991df-2dcd-4e24-83fe-5df46ec04422",
	      "version": 3,
	      "name": "DHL",
	      "description": "Standard delivery",
	      "taxCategory": {
	        "typeId": "tax-category",
	        "id": "5a21f15b-34f8-4b7f-9407-d1eb82a73eba"
	      },
	      "zoneRates": [
	        {
	          "zone": {
	            "typeId": "zone",
	            "id": "5cb532be-c680-43ab-ba14-b664bb03dc35"
	          },
	          "shippingRates": [
	            {
	              "price": {
	                "type": "centPrecision",
	                "fractionDigits": 2,
	                "currencyCode": "EUR",
	                "centAmount": 570
	              },
	              "tiers": []
	            }
	          ]
	        },
	        {
	          "zone": {
	            "typeId": "zone",
	            "id": "ebe01381-82be-4e63-9993-d1eb8f8e588b"
	          },
	          "shippingRates": [
	            {
	              "price": {
	                "type": "centPrecision",
	                "fractionDigits": 2,
	                "currencyCode": "USD",
	                "centAmount": 990
	              },
	              "tiers": []
	            }
	          ]
	        }
	      ],
	      "isDefault": false,
	      "createdAt": "2015-01-21T09:22:04.320Z",
	      "lastModifiedAt": "2016-02-24T13:36:56.748Z"
	    }
	  ]
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	input := commercetools.QueryInput{}
	queryResult, err := client.ShippingMethodQuery(context.TODO(), &input)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, queryResult)
	assert.NotNil(t, queryResult.Total)
	assert.NotNil(t, queryResult.Offset)
	assert.NotNil(t, queryResult.Limit)
	assert.NotNil(t, queryResult.Count)

}
