// Automatically generated, do not edit

package commercetools_test

import (
	"context"
	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratedPaymentGetWithID(t *testing.T) {
	responseData := ` {
	  "id": "776f9240-09e1-4416-a34f-32fa80f0333f",
	  "version": 1,
	  "externalId": "123456",
	  "interfaceId": "789011",
	  "amountPlanned": {
	    "type": "centPrecision",
	    "fractionDigits": 2,
	    "currencyCode": "USD",
	    "centAmount": 1000
	  },
	  "paymentMethodInfo": {
	    "paymentInterface": "STRIPE",
	    "method": "CREDIT_CARD",
	    "name": {
	      "en": "Credit Card"
	    }
	  },
	  "paymentStatus": {},
	  "transactions": [
	    {
	      "id": "14748fe6-7f77-456a-96c8-913b1e4bbc9a",
	      "timestamp": "2015-10-20T08:54:24.000Z",
	      "type": "Charge",
	      "amount": {
	        "type": "centPrecision",
	        "fractionDigits": 2,
	        "currencyCode": "USD",
	        "centAmount": 1000
	      },
	      "state": "Pending"
	    }
	  ],
	  "interfaceInteractions": [],
	  "createdAt": "2015-10-20T08:54:11.480Z",
	  "lastModifiedAt": "2015-10-20T08:54:11.480Z",
	  "lastMessageSequenceNumber": 1
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	payment, err := client.PaymentGetWithID(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, payment)
	assert.NotNil(t, payment.Version)
	assert.NotEmpty(t, payment.LastModifiedAt)
	assert.NotEmpty(t, payment.InterfaceID)
	assert.NotEmpty(t, payment.ID)
	assert.NotEmpty(t, payment.ExternalID)
	assert.NotEmpty(t, payment.CreatedAt)

}

func TestGeneratedPaymentGetWithKey(t *testing.T) {
	responseData := ` {
	  "id": "776f9240-09e1-4416-a34f-32fa80f0333f",
	  "version": 1,
	  "externalId": "123456",
	  "interfaceId": "789011",
	  "amountPlanned": {
	    "type": "centPrecision",
	    "fractionDigits": 2,
	    "currencyCode": "USD",
	    "centAmount": 1000
	  },
	  "paymentMethodInfo": {
	    "paymentInterface": "STRIPE",
	    "method": "CREDIT_CARD",
	    "name": {
	      "en": "Credit Card"
	    }
	  },
	  "paymentStatus": {},
	  "transactions": [
	    {
	      "id": "14748fe6-7f77-456a-96c8-913b1e4bbc9a",
	      "timestamp": "2015-10-20T08:54:24.000Z",
	      "type": "Charge",
	      "amount": {
	        "type": "centPrecision",
	        "fractionDigits": 2,
	        "currencyCode": "USD",
	        "centAmount": 1000
	      },
	      "state": "Pending"
	    }
	  ],
	  "interfaceInteractions": [],
	  "createdAt": "2015-10-20T08:54:11.480Z",
	  "lastModifiedAt": "2015-10-20T08:54:11.480Z",
	  "lastMessageSequenceNumber": 1
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	payment, err := client.PaymentGetWithKey(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, payment)
	assert.NotNil(t, payment.Version)
	assert.NotEmpty(t, payment.LastModifiedAt)
	assert.NotEmpty(t, payment.InterfaceID)
	assert.NotEmpty(t, payment.ID)
	assert.NotEmpty(t, payment.ExternalID)
	assert.NotEmpty(t, payment.CreatedAt)

}

func TestGeneratedPaymentDeleteWithID(t *testing.T) {
	responseData := ` {
	  "id": "776f9240-09e1-4416-a34f-32fa80f0333f",
	  "version": 1,
	  "externalId": "123456",
	  "interfaceId": "789011",
	  "amountPlanned": {
	    "type": "centPrecision",
	    "fractionDigits": 2,
	    "currencyCode": "USD",
	    "centAmount": 1000
	  },
	  "paymentMethodInfo": {
	    "paymentInterface": "STRIPE",
	    "method": "CREDIT_CARD",
	    "name": {
	      "en": "Credit Card"
	    }
	  },
	  "paymentStatus": {},
	  "transactions": [
	    {
	      "id": "14748fe6-7f77-456a-96c8-913b1e4bbc9a",
	      "timestamp": "2015-10-20T08:54:24.000Z",
	      "type": "Charge",
	      "amount": {
	        "type": "centPrecision",
	        "fractionDigits": 2,
	        "currencyCode": "USD",
	        "centAmount": 1000
	      },
	      "state": "Pending"
	    }
	  ],
	  "interfaceInteractions": [],
	  "createdAt": "2015-10-20T08:54:11.480Z",
	  "lastModifiedAt": "2015-10-20T08:54:11.480Z",
	  "lastMessageSequenceNumber": 1
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	payment, err := client.PaymentDeleteWithID(context.TODO(), "dummy-id", 1, true)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, payment)
	assert.NotNil(t, payment.Version)
	assert.NotEmpty(t, payment.LastModifiedAt)
	assert.NotEmpty(t, payment.InterfaceID)
	assert.NotEmpty(t, payment.ID)
	assert.NotEmpty(t, payment.ExternalID)
	assert.NotEmpty(t, payment.CreatedAt)

}

func TestGeneratedPaymentDeleteWithKey(t *testing.T) {
	responseData := ` {
	  "id": "776f9240-09e1-4416-a34f-32fa80f0333f",
	  "version": 1,
	  "externalId": "123456",
	  "interfaceId": "789011",
	  "amountPlanned": {
	    "type": "centPrecision",
	    "fractionDigits": 2,
	    "currencyCode": "USD",
	    "centAmount": 1000
	  },
	  "paymentMethodInfo": {
	    "paymentInterface": "STRIPE",
	    "method": "CREDIT_CARD",
	    "name": {
	      "en": "Credit Card"
	    }
	  },
	  "paymentStatus": {},
	  "transactions": [
	    {
	      "id": "14748fe6-7f77-456a-96c8-913b1e4bbc9a",
	      "timestamp": "2015-10-20T08:54:24.000Z",
	      "type": "Charge",
	      "amount": {
	        "type": "centPrecision",
	        "fractionDigits": 2,
	        "currencyCode": "USD",
	        "centAmount": 1000
	      },
	      "state": "Pending"
	    }
	  ],
	  "interfaceInteractions": [],
	  "createdAt": "2015-10-20T08:54:11.480Z",
	  "lastModifiedAt": "2015-10-20T08:54:11.480Z",
	  "lastMessageSequenceNumber": 1
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	payment, err := client.PaymentDeleteWithKey(context.TODO(), "dummy-id", 1, true)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, payment)
	assert.NotNil(t, payment.Version)
	assert.NotEmpty(t, payment.LastModifiedAt)
	assert.NotEmpty(t, payment.InterfaceID)
	assert.NotEmpty(t, payment.ID)
	assert.NotEmpty(t, payment.ExternalID)
	assert.NotEmpty(t, payment.CreatedAt)

}

func TestGeneratedPaymentQuery(t *testing.T) {
	responseData := ` {
	  "limit": 20,
	  "offset": 0,
	  "count": 2,
	  "total": 2,
	  "results": [
	    {
	      "id": "459e32dc-74ef-4189-bbd0-932275bb027c",
	      "version": 1,
	      "externalId": "123456",
	      "interfaceId": "78901",
	      "amountPlanned": {
	        "type": "centPrecision",
	        "fractionDigits": 2,
	        "currencyCode": "USD",
	        "centAmount": 1000
	      },
	      "paymentMethodInfo": {
	        "paymentInterface": "STRIPE",
	        "method": "CREDIT_CARD",
	        "name": {
	          "en": "Credit Card"
	        }
	      },
	      "paymentStatus": {},
	      "transactions": [
	        {
	          "id": "2e318aa5-8af4-4db1-909d-e7142f7d174f",
	          "timestamp": "2015-10-20T08:51:56.000Z",
	          "type": "Charge",
	          "amount": {
	            "type": "centPrecision",
	            "fractionDigits": 2,
	            "currencyCode": "USD",
	            "centAmount": 1000
	          },
	          "state": "Pending"
	        }
	      ],
	      "interfaceInteractions": [],
	      "createdAt": "2015-10-20T08:51:43.082Z",
	      "lastModifiedAt": "2015-10-20T08:51:43.082Z",
	      "lastMessageSequenceNumber": 1
	    },
	    {
	      "id": "776f9240-09e1-4416-a34f-32fa80f0333f",
	      "version": 1,
	      "externalId": "123456",
	      "interfaceId": "789011",
	      "amountPlanned": {
	        "type": "centPrecision",
	        "fractionDigits": 2,
	        "currencyCode": "USD",
	        "centAmount": 1000
	      },
	      "paymentMethodInfo": {
	        "paymentInterface": "STRIPE",
	        "method": "CREDIT_CARD",
	        "name": {
	          "en": "Credit Card"
	        }
	      },
	      "paymentStatus": {},
	      "transactions": [
	        {
	          "id": "14748fe6-7f77-456a-96c8-913b1e4bbc9a",
	          "timestamp": "2015-10-20T08:54:24.000Z",
	          "type": "Charge",
	          "amount": {
	            "type": "centPrecision",
	            "fractionDigits": 2,
	            "currencyCode": "USD",
	            "centAmount": 1000
	          },
	          "state": "Pending"
	        }
	      ],
	      "interfaceInteractions": [],
	      "createdAt": "2015-10-20T08:54:11.480Z",
	      "lastModifiedAt": "2015-10-20T08:54:11.480Z",
	      "lastMessageSequenceNumber": 1
	    }
	  ]
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	input := commercetools.QueryInput{}
	queryResult, err := client.PaymentQuery(context.TODO(), &input)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, queryResult)
	assert.NotNil(t, queryResult.Total)
	assert.NotNil(t, queryResult.Offset)
	assert.NotNil(t, queryResult.Limit)
	assert.NotNil(t, queryResult.Count)

}
