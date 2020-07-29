// Automatically generated, do not edit

package commercetools_test

import (
	"context"
	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratedTypeGetWithID(t *testing.T) {
	responseData := ` {
	  "id": "3ae9bcca-df23-443e-bd22-0c592f9694fa",
	  "version": 1,
	  "key": "lineitemtype",
	  "name": {
	    "en": "lineitem"
	  },
	  "description": {
	    "en": "description"
	  },
	  "resourceTypeIds": [
	    "line-item"
	  ],
	  "fieldDefinitions": [
	    {
	      "name": "offer_name",
	      "label": {
	        "en": "offer_name"
	      },
	      "required": false,
	      "type": {
	        "name": "String"
	      },
	      "inputHint": "SingleLine"
	    }
	  ],
	  "createdAt": "2015-10-07T06:56:19.217Z",
	  "lastModifiedAt": "2015-10-07T06:56:19.217Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	rType, err := client.TypeGetWithID(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, rType)
	assert.NotNil(t, rType.Version)
	assert.NotEmpty(t, rType.LastModifiedAt)
	assert.NotEmpty(t, rType.Key)
	assert.NotEmpty(t, rType.ID)
	assert.NotEmpty(t, rType.CreatedAt)

}

func TestGeneratedTypeGetWithKey(t *testing.T) {
	responseData := ` {
	  "id": "3ae9bcca-df23-443e-bd22-0c592f9694fa",
	  "version": 1,
	  "key": "lineitemtype",
	  "name": {
	    "en": "lineitem"
	  },
	  "description": {
	    "en": "description"
	  },
	  "resourceTypeIds": [
	    "line-item"
	  ],
	  "fieldDefinitions": [
	    {
	      "name": "offer_name",
	      "label": {
	        "en": "offer_name"
	      },
	      "required": false,
	      "type": {
	        "name": "String"
	      },
	      "inputHint": "SingleLine"
	    }
	  ],
	  "createdAt": "2015-10-07T06:56:19.217Z",
	  "lastModifiedAt": "2015-10-07T06:56:19.217Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	rType, err := client.TypeGetWithKey(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, rType)
	assert.NotNil(t, rType.Version)
	assert.NotEmpty(t, rType.LastModifiedAt)
	assert.NotEmpty(t, rType.Key)
	assert.NotEmpty(t, rType.ID)
	assert.NotEmpty(t, rType.CreatedAt)

}

func TestGeneratedTypeDeleteWithID(t *testing.T) {
	responseData := ` {
	  "id": "3ae9bcca-df23-443e-bd22-0c592f9694fa",
	  "version": 1,
	  "key": "lineitemtype",
	  "name": {
	    "en": "lineitem"
	  },
	  "description": {
	    "en": "description"
	  },
	  "resourceTypeIds": [
	    "line-item"
	  ],
	  "fieldDefinitions": [
	    {
	      "name": "offer_name",
	      "label": {
	        "en": "offer_name"
	      },
	      "required": false,
	      "type": {
	        "name": "String"
	      },
	      "inputHint": "SingleLine"
	    }
	  ],
	  "createdAt": "2015-10-07T06:56:19.217Z",
	  "lastModifiedAt": "2015-10-07T06:56:19.217Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	rType, err := client.TypeDeleteWithID(context.TODO(), "dummy-id", 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, rType)
	assert.NotNil(t, rType.Version)
	assert.NotEmpty(t, rType.LastModifiedAt)
	assert.NotEmpty(t, rType.Key)
	assert.NotEmpty(t, rType.ID)
	assert.NotEmpty(t, rType.CreatedAt)

}

func TestGeneratedTypeDeleteWithKey(t *testing.T) {
	responseData := ` {
	  "id": "3ae9bcca-df23-443e-bd22-0c592f9694fa",
	  "version": 1,
	  "key": "lineitemtype",
	  "name": {
	    "en": "lineitem"
	  },
	  "description": {
	    "en": "description"
	  },
	  "resourceTypeIds": [
	    "line-item"
	  ],
	  "fieldDefinitions": [
	    {
	      "name": "offer_name",
	      "label": {
	        "en": "offer_name"
	      },
	      "required": false,
	      "type": {
	        "name": "String"
	      },
	      "inputHint": "SingleLine"
	    }
	  ],
	  "createdAt": "2015-10-07T06:56:19.217Z",
	  "lastModifiedAt": "2015-10-07T06:56:19.217Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	rType, err := client.TypeDeleteWithKey(context.TODO(), "dummy-id", 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, rType)
	assert.NotNil(t, rType.Version)
	assert.NotEmpty(t, rType.LastModifiedAt)
	assert.NotEmpty(t, rType.Key)
	assert.NotEmpty(t, rType.ID)
	assert.NotEmpty(t, rType.CreatedAt)

}

func TestGeneratedTypeQuery(t *testing.T) {
	responseData := ` {
	  "limit": 20,
	  "offset": 0,
	  "count": 1,
	  "total": 1,
	  "results": [
	    {
	      "id": "3ae9bcca-df23-443e-bd22-0c592f9694fa",
	      "version": 1,
	      "key": "lineitemtype",
	      "name": {
	        "en": "lineitem"
	      },
	      "description": {
	        "en": "description"
	      },
	      "resourceTypeIds": [
	        "line-item"
	      ],
	      "fieldDefinitions": [
	        {
	          "name": "offer_name",
	          "label": {
	            "en": "offer_name"
	          },
	          "required": false,
	          "type": {
	            "name": "String"
	          },
	          "inputHint": "SingleLine"
	        }
	      ],
	      "createdAt": "2015-10-07T06:56:19.217Z",
	      "lastModifiedAt": "2015-10-07T06:56:19.217Z"
	    }
	  ]
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	input := commercetools.QueryInput{}
	queryResult, err := client.TypeQuery(context.TODO(), &input)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, queryResult)
	assert.NotNil(t, queryResult.Total)
	assert.NotNil(t, queryResult.Offset)
	assert.NotNil(t, queryResult.Limit)
	assert.NotNil(t, queryResult.Count)

}
