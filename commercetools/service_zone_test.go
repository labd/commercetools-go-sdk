package commercetools_test

import (
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
)

func TestZoneCreate(t *testing.T) {
	output := testutil.RequestData{}

	client, server := testutil.MockClient(t, testutil.Fixture("shipping-zone.create.json"), &output, nil)
	defer server.Close()
	input := &commercetools.ZoneDraft{
		Name:        "Primary zone",
		Description: "Primary zone description.",
		Locations: []commercetools.Location{
			{
				Country: "NL",
			},
			{
				Country: "DE",
			},
		},
	}

	fmt.Println(output)

	_, err := client.ZoneCreate(input)
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
		input       *commercetools.ZoneUpdateWithIdInput
		requestBody string
	}{
		{
			desc: "Change name",
			input: &commercetools.ZoneUpdateWithIdInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.ZoneUpdateAction{
					&commercetools.ZoneChangeNameAction{
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
			input: &commercetools.ZoneUpdateWithIdInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.ZoneUpdateAction{
					&commercetools.ZoneSetDescriptionAction{
						Description: "A shipping zone description",
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
			input: &commercetools.ZoneUpdateWithIdInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.ZoneUpdateAction{
					&commercetools.ZoneAddLocationAction{
						Location: &commercetools.Location{
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
			input: &commercetools.ZoneUpdateWithIdInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.ZoneUpdateAction{
					&commercetools.ZoneRemoveLocationAction{
						Location: &commercetools.Location{
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

			client, server := testutil.MockClient(t, testutil.Fixture("shipping-zone.update.json"), &output, nil)
			defer server.Close()

			_, err := client.ZoneUpdateWithId(tC.input)
			assert.Nil(t, err)
			assert.Equal(t, "/unittest/zones/1234", output.URL.Path)
			assert.JSONEq(t, tC.requestBody, output.JSON)
		})
	}
}

func TestZoneDeleteByID(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, "{}", &output, nil)
	defer server.Close()

	_, err := client.ZoneDeleteWithId("1234", 2)
	assert.Nil(t, err)

	params := url.Values{}
	params.Add("version", "2")
	assert.Equal(t, params, output.URL.Query())
	assert.Equal(t, "/unittest/zones/1234", output.URL.Path)
}

func TestZoneGetByID(t *testing.T) {
	timestamp, _ := time.Parse(time.RFC3339, "2016-02-24T15:33:40.811Z")

	client, server := testutil.MockClient(t, testutil.Fixture("shipping-zone.json"), nil, nil)
	defer server.Close()

	input := &commercetools.Zone{
		ID:      "c60f7377-2643-4e99-adb5-b2909657444d",
		Version: 1,
		Name:    "shipping-zone",
		Locations: []commercetools.Location{
			{
				Country: "DE",
			},
		},
		CreatedAt:      timestamp,
		LastModifiedAt: timestamp,
	}

	result, err := client.ZoneGetWithId("1234")
	assert.Nil(t, err)
	assert.Equal(t, input, result)
}

func TestZoneQuery(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, "{}", &output, nil)
	defer server.Close()

	queryInput := commercetools.QueryInput{
		Limit: 500,
	}
	_, err := client.ZoneQuery(&queryInput)
	assert.Nil(t, err)

	assert.Equal(t, url.Values{"limit": []string{"500"}}, output.URL.Query())
}
