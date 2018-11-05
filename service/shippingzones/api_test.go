package shippingzones_test

import (
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/service/shippingzones"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
)

func TestShippingZoneCreate(t *testing.T) {
	output := testutil.RequestData{}

	client, server := testutil.MockClient(t, testutil.Fixture("shipping-zone.create.json"), &output, nil)
	defer server.Close()
	svc := shippingzones.New(client)

	input := &shippingzones.ShippingZoneDraft{
		Name:        "Primary zone",
		Description: "Primary zone description.",
		Locations: []shippingzones.Location{
			{
				Country: "NL",
			},
			{
				Country: "DE",
			},
		},
	}

	fmt.Println(output)

	_, err := svc.Create(input)
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

func TestShippingZoneUpdate(t *testing.T) {
	testCases := []struct {
		desc        string
		input       *shippingzones.UpdateInput
		requestBody string
	}{
		{
			desc: "Change name",
			input: &shippingzones.UpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&shippingzones.ChangeName{
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
			input: &shippingzones.UpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&shippingzones.SetDescription{
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
			input: &shippingzones.UpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&shippingzones.AddLocation{
						Location: shippingzones.Location{
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
			input: &shippingzones.UpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&shippingzones.RemoveLocation{
						Location: shippingzones.Location{
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
			svc := shippingzones.New(client)

			_, err := svc.Update(tC.input)
			assert.Nil(t, err)
			assert.Equal(t, "/unittest/zones/1234", output.URL.Path)
			assert.JSONEq(t, tC.requestBody, output.JSON)
		})
	}
}

func TestShippingZoneDeleteByID(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, "{}", &output, nil)
	defer server.Close()
	svc := shippingzones.New(client)

	_, err := svc.DeleteByID("1234", 2)
	assert.Nil(t, err)

	params := url.Values{}
	params.Add("version", "2")
	assert.Equal(t, params, output.URL.Query())
	assert.Equal(t, "/unittest/zones/1234", output.URL.Path)
}

func TestShippingZoneGetByID(t *testing.T) {
	timestamp, _ := time.Parse(time.RFC3339, "2016-02-24T15:33:40.811Z")

	client, server := testutil.MockClient(t, testutil.Fixture("shipping-zone.json"), nil, nil)
	defer server.Close()

	input := &shippingzones.ShippingZone{
		ID:      "c60f7377-2643-4e99-adb5-b2909657444d",
		Version: 1,
		Name:    "shipping-zone",
		Locations: []shippingzones.Location{
			{
				Country: "DE",
			},
		},
		CreatedAt:      timestamp,
		LastModifiedAt: timestamp,
	}

	svc := shippingzones.New(client)
	result, err := svc.GetByID("1234")
	assert.Nil(t, err)
	assert.Equal(t, input, result)
}
