// Automatically generated, do not edit

package commercetools_test

import (
	"context"
	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratedReviewGetWithID(t *testing.T) {
	responseData := ` {
	  "id": "919d5ab5-bf99-4bda-a6fd-ab3ee8123456",
	  "version": 1,
	  "includedInStatistics": true,
	  "authorName": "John Doe",
	  "title": "Incredible",
	  "text": "Best product ever",
	  "rating": 5,
	  "target": {
	    "typeId": "product",
	    "id": "8fddacac-6ef5-4e66-af6e-124452123456"
	  },
	  "customer": {
	    "typeId": "customer",
	    "id": "8a8b3e43-b9b0-4b30-8c27-58148123456"
	  },
	  "custom": {
	    "type": {
	      "typeId": "type",
	      "id": "3939dd9c-0884-4bfa-99c2-40b426123456"
	    },
	    "fields": {
	      "authorMail": "john.doe@example.com"
	    }
	  },
	  "createdAt": "2016-10-20T06:48:53.829Z",
	  "lastModifiedAt": "2016-10-20T06:48:53.829Z",
	  "lastMessageSequenceNumber": 1
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	review, err := client.ReviewGetWithID(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, review)
	assert.NotNil(t, review.Version)
	assert.NotEmpty(t, review.Title)
	assert.NotEmpty(t, review.Text)
	assert.NotNil(t, review.Rating)
	assert.NotEmpty(t, review.LastModifiedAt)
	assert.True(t, review.IncludedInStatistics)
	assert.NotEmpty(t, review.ID)
	assert.NotEmpty(t, review.CreatedAt)
	assert.NotEmpty(t, review.AuthorName)

}

func TestGeneratedReviewGetWithKey(t *testing.T) {
	responseData := ` {
	  "id": "919d5ab5-bf99-4bda-a6fd-ab3ee8123456",
	  "version": 1,
	  "includedInStatistics": true,
	  "authorName": "John Doe",
	  "title": "Incredible",
	  "text": "Best product ever",
	  "rating": 5,
	  "target": {
	    "typeId": "product",
	    "id": "8fddacac-6ef5-4e66-af6e-124452123456"
	  },
	  "customer": {
	    "typeId": "customer",
	    "id": "8a8b3e43-b9b0-4b30-8c27-58148123456"
	  },
	  "custom": {
	    "type": {
	      "typeId": "type",
	      "id": "3939dd9c-0884-4bfa-99c2-40b426123456"
	    },
	    "fields": {
	      "authorMail": "john.doe@example.com"
	    }
	  },
	  "createdAt": "2016-10-20T06:48:53.829Z",
	  "lastModifiedAt": "2016-10-20T06:48:53.829Z",
	  "lastMessageSequenceNumber": 1
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	review, err := client.ReviewGetWithKey(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, review)
	assert.NotNil(t, review.Version)
	assert.NotEmpty(t, review.Title)
	assert.NotEmpty(t, review.Text)
	assert.NotNil(t, review.Rating)
	assert.NotEmpty(t, review.LastModifiedAt)
	assert.True(t, review.IncludedInStatistics)
	assert.NotEmpty(t, review.ID)
	assert.NotEmpty(t, review.CreatedAt)
	assert.NotEmpty(t, review.AuthorName)

}

func TestGeneratedReviewDeleteWithID(t *testing.T) {
	responseData := ` {
	  "id": "919d5ab5-bf99-4bda-a6fd-ab3ee8123456",
	  "version": 1,
	  "includedInStatistics": true,
	  "authorName": "John Doe",
	  "title": "Incredible",
	  "text": "Best product ever",
	  "rating": 5,
	  "target": {
	    "typeId": "product",
	    "id": "8fddacac-6ef5-4e66-af6e-124452123456"
	  },
	  "customer": {
	    "typeId": "customer",
	    "id": "8a8b3e43-b9b0-4b30-8c27-58148123456"
	  },
	  "custom": {
	    "type": {
	      "typeId": "type",
	      "id": "3939dd9c-0884-4bfa-99c2-40b426123456"
	    },
	    "fields": {
	      "authorMail": "john.doe@example.com"
	    }
	  },
	  "createdAt": "2016-10-20T06:48:53.829Z",
	  "lastModifiedAt": "2016-10-20T06:48:53.829Z",
	  "lastMessageSequenceNumber": 1
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	review, err := client.ReviewDeleteWithID(context.TODO(), "dummy-id", 1, true)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, review)
	assert.NotNil(t, review.Version)
	assert.NotEmpty(t, review.Title)
	assert.NotEmpty(t, review.Text)
	assert.NotNil(t, review.Rating)
	assert.NotEmpty(t, review.LastModifiedAt)
	assert.True(t, review.IncludedInStatistics)
	assert.NotEmpty(t, review.ID)
	assert.NotEmpty(t, review.CreatedAt)
	assert.NotEmpty(t, review.AuthorName)

}

func TestGeneratedReviewDeleteWithKey(t *testing.T) {
	responseData := ` {
	  "id": "919d5ab5-bf99-4bda-a6fd-ab3ee8123456",
	  "version": 1,
	  "includedInStatistics": true,
	  "authorName": "John Doe",
	  "title": "Incredible",
	  "text": "Best product ever",
	  "rating": 5,
	  "target": {
	    "typeId": "product",
	    "id": "8fddacac-6ef5-4e66-af6e-124452123456"
	  },
	  "customer": {
	    "typeId": "customer",
	    "id": "8a8b3e43-b9b0-4b30-8c27-58148123456"
	  },
	  "custom": {
	    "type": {
	      "typeId": "type",
	      "id": "3939dd9c-0884-4bfa-99c2-40b426123456"
	    },
	    "fields": {
	      "authorMail": "john.doe@example.com"
	    }
	  },
	  "createdAt": "2016-10-20T06:48:53.829Z",
	  "lastModifiedAt": "2016-10-20T06:48:53.829Z",
	  "lastMessageSequenceNumber": 1
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	review, err := client.ReviewDeleteWithKey(context.TODO(), "dummy-id", 1, true)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, review)
	assert.NotNil(t, review.Version)
	assert.NotEmpty(t, review.Title)
	assert.NotEmpty(t, review.Text)
	assert.NotNil(t, review.Rating)
	assert.NotEmpty(t, review.LastModifiedAt)
	assert.True(t, review.IncludedInStatistics)
	assert.NotEmpty(t, review.ID)
	assert.NotEmpty(t, review.CreatedAt)
	assert.NotEmpty(t, review.AuthorName)

}

func TestGeneratedReviewQuery(t *testing.T) {
	responseData := ` {
	  "limit": 20,
	  "offset": 0,
	  "count": 2,
	  "total": 2,
	  "results": [
	    {
	      "id": "919d5ab5-bf99-4bda-a6fd-ab3ee8123456",
	      "version": 1,
	      "includedInStatistics": true,
	      "authorName": "John Doe",
	      "title": "Incredible",
	      "text": "Best product ever",
	      "rating": 5,
	      "target": {
	        "typeId": "product",
	        "id": "8fddacac-6ef5-4e66-af6e-124452123456"
	      },
	      "customer": {
	        "typeId": "customer",
	        "id": "8a8b3e43-b9b0-4b30-8c27-58148123456"
	      },
	      "custom": {
	        "type": {
	          "typeId": "type",
	          "id": "3939dd9c-0884-4bfa-99c2-40b426123456"
	        },
	        "fields": {
	          "authorMail": "john.doe@example.com"
	        }
	      },
	      "createdAt": "2016-10-20T06:48:53.829Z",
	      "lastModifiedAt": "2016-10-20T06:48:53.829Z",
	      "lastMessageSequenceNumber": 1
	    },
	    {
	      "id": "2f60e06c-7672-47fc-962a-1eafa1123456",
	      "version": 1,
	      "includedInStatistics": true,
	      "authorName": ".",
	      "title": ".",
	      "text": ".",
	      "rating": 5,
	      "target": {
	        "typeId": "product",
	        "id": "2c155644-bcde-426c-b021-a2aab123456"
	      },
	      "custom": {
	        "type": {
	          "typeId": "type",
	          "id": "3939dd9c-0884-4bfa-99c2-40b42123456"
	        },
	        "fields": {
	          "authorMail": "."
	        }
	      },
	      "createdAt": "2016-10-20T06:51:26.795Z",
	      "lastModifiedAt": "2016-10-20T06:51:26.795Z",
	      "lastMessageSequenceNumber": 1
	    }
	  ]
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	input := commercetools.QueryInput{}
	queryResult, err := client.ReviewQuery(context.TODO(), &input)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, queryResult)
	assert.NotNil(t, queryResult.Total)
	assert.NotNil(t, queryResult.Offset)
	assert.NotNil(t, queryResult.Limit)
	assert.NotNil(t, queryResult.Count)

}
