package commercetools_test

import (
	"context"
	"testing"

	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
)

func TestCustomApplicationGetWithID(t *testing.T) {
	jsonData := `{
		"data":{
			"projectExtension": {
				"id":"ckcmb5qvfzper07140kmkkvuc",
				"applications":[
					{
						"id":"ckcnc6hydu2ou08131y504crj",
						"createdAt":"2020-07-15T12:24:10.593Z",
						"updatedAt":"2020-07-15T12:24:10.593Z",
						"isActive":false,
						"name":"noooooo-test",
						"description":"test terraform---xxxx-",
						"url":"http://labdigital.nl",
						"navbarMenu":{
							"uriPath":"sort-test",
							"icon":"",
							"labelAllLocales":[
								{
									"locale":"NL",
									"value":"Michael"
								}
							],
							"permissions":["ViewCategories"],
							"submenu":[]
						}
					}
				]
			}
		}
	}`

	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, jsonData, &output, nil)
	defer server.Close()

	_, err := client.CustomApplicationGetWithID(context.TODO(), "1234")
	assert.Nil(t, err)

	assert.Equal(t, "/graphql", output.URL.Path)
}
