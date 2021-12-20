package platform_test

import (
	"context"
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/labd/commercetools-go-sdk/ctutils"
	"github.com/labd/commercetools-go-sdk/platform"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
)

func TestZoneCreate(t *testing.T) {
	output := testutil.RequestData{}
	response := testutil.ResponseData{
		StatusCode: 201,
		Body: `
		{
			"name": "shipping-zone",
			"locations": [
			  {
				"country": "DE"
			  }
			]
		}
		`,
	}

	client, server := testutil.MockClient(t, response, &output, nil)
	defer server.Close()

	input := platform.ZoneDraft{
		Name:        "Primary zone",
		Description: ctutils.StringRef("Primary zone description."),
		Locations: []platform.Location{
			{
				Country: "NL",
			},
			{
				Country: "DE",
			},
		},
	}

	fmt.Println(output)

	_, err := client.WithProjectKey("unittest").Zones().Post(input).Execute(context.TODO())
	assert.Nil(t, err)

	expectedBody := `{
        "name": "Primary zone",
        "description": "Primary zone description.",
        "locations": [
            {
                "country": "NL"
            },
            {
                "country": "DE"
            }
        ]
    }`

	assert.JSONEq(t, expectedBody, output.JSON)
}

func TestZoneUpdate(t *testing.T) {
	testCases := []struct {
		desc        string
		input       platform.ZoneUpdate
		requestBody string
	}{
		{
			desc: "Change name",
			input: platform.ZoneUpdate{
				Version: 2,
				Actions: []platform.ZoneUpdateAction{
					&platform.ZoneChangeNameAction{
						Name: "New name",
					},
				},
			},
			requestBody: `{
                "version": 2,
                "actions": [
                    {
                        "action": "changeName",
                        "name": "New name"
                    }
                ]
            }`,
		},
		{
			desc: "Set description",
			input: platform.ZoneUpdate{
				Version: 2,
				Actions: []platform.ZoneUpdateAction{
					&platform.ZoneSetDescriptionAction{
						Description: ctutils.StringRef("A shipping zone description"),
					},
				},
			},
			requestBody: `{
                "version": 2,
                "actions": [
                    {
                        "action": "setDescription",
                        "description": "A shipping zone description"
                    }
                ]
            }`,
		},
		{
			desc: "Add location ",
			input: platform.ZoneUpdate{
				Version: 2,
				Actions: []platform.ZoneUpdateAction{
					&platform.ZoneAddLocationAction{
						Location: platform.Location{
							Country: "US",
						},
					},
				},
			},
			requestBody: `{
                "version": 2,
                "actions": [
                    {
                        "action": "addLocation",
                        "location": {
                            "country": "US"
                        }
                    }
                ]
            }`,
		},
		{
			desc: "Remove location",
			input: platform.ZoneUpdate{
				Version: 2,
				Actions: []platform.ZoneUpdateAction{
					&platform.ZoneRemoveLocationAction{
						Location: platform.Location{
							Country: "US",
						},
					},
				},
			},
			requestBody: `{
                "version": 2,
                "actions": [
                    {
                        "action": "removeLocation",
                        "location": {
                            "country": "US"
                        }
                    }
                ]
            }`,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			output := testutil.RequestData{}
			response := testutil.Fixture("shipping-zone.update.json", 200)

			client, server := testutil.MockClient(t, response, &output, nil)
			defer server.Close()

			_, err := client.
				WithProjectKey("unittest").
				Zones().
				WithId("1234").
				Post(tC.input).
				Execute(context.TODO())
			assert.Nil(t, err)
			assert.Equal(t, "/unittest/zones/1234", output.URL.Path)
			assert.JSONEq(t, tC.requestBody, output.JSON)
		})
	}
}

func TestZoneDeleteByID(t *testing.T) {
	output := testutil.RequestData{}
	response := testutil.ResponseData{StatusCode: 200, Body: "{}"}
	client, server := testutil.MockClient(t, response, &output, nil)
	defer server.Close()

	_, err := client.WithProjectKey("unittest").
		Zones().
		WithId("1234").
		Delete().
		WithQueryParams(platform.ByProjectKeyZonesByIDRequestMethodDeleteInput{
			Version: 2,
		}).
		Execute(context.TODO())

	assert.Nil(t, err)

	params := url.Values{}
	params.Add("version", "2")
	assert.Equal(t, params, output.URL.Query())
	assert.Equal(t, "/unittest/zones/1234", output.URL.Path)
}

func TestZoneGetByID(t *testing.T) {
	timestamp, _ := time.Parse(time.RFC3339, "2016-02-24T15:33:40.811Z")
	response := testutil.Fixture("shipping-zone.json", 200)

	client, server := testutil.MockClient(t, response, nil, nil)
	defer server.Close()

	input := &platform.Zone{
		ID:      "c60f7377-2643-4e99-adb5-b2909657444d",
		Version: 1,
		Name:    "shipping-zone",
		Locations: []platform.Location{
			{
				Country: "DE",
			},
		},
		CreatedAt:      timestamp,
		LastModifiedAt: timestamp,
	}

	result, err := client.WithProjectKey("unittest").Zones().WithId("1234").Get().Execute(context.TODO())
	assert.Nil(t, err)
	assert.Equal(t, input, result)
}

func TestZoneQuery(t *testing.T) {
	output := testutil.RequestData{}
	response := testutil.ResponseData{StatusCode: 200, Body: "{}"}

	client, server := testutil.MockClient(t, response, &output, nil)
	defer server.Close()

	_, err := client.
		WithProjectKey("unittest").
		Zones().
		Get().
		WithQueryParams(platform.ByProjectKeyZonesRequestMethodGetInput{
			Limit: ctutils.IntRef(500),
		}).
		Execute(context.TODO())

	assert.Nil(t, err)
	assert.Equal(t, url.Values{"limit": []string{"500"}}, output.URL.Query())
}
