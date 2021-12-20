package platform_test

import (
	"context"
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/labd/commercetools-go-sdk/ctutils"
	"github.com/labd/commercetools-go-sdk/platform"
	"github.com/labd/commercetools-go-sdk/testutil"
)

func TestStoreCreate(t *testing.T) {
	output := testutil.RequestData{}
	response := testutil.Fixture("store.json", 201)

	client, server := testutil.MockClient(t, response, &output, nil)
	defer server.Close()

	input := platform.StoreDraft{
		Key: "test123",
		Name: &platform.LocalizedString{
			"en": "test123",
		},
	}

	store, err := client.WithProjectKey("unittest").Stores().Post(input).Execute(context.TODO())
	fmt.Printf("%#v", store)
	assert.Nil(t, err)

	assert.Equal(t, "test123", store.Key)
	assert.Equal(t, "test123", (*store.Name)["en"])
}

func TestStoreUpdate(t *testing.T) {
	testCases := []struct {
		desc        string
		input       platform.StoreUpdate
		requestBody string
	}{
		{
			desc: "Set locale",
			input: platform.StoreUpdate{
				Version: 2,
				Actions: []platform.StoreUpdateAction{
					&platform.StoreSetNameAction{
						Name: &platform.LocalizedString{
							"en": "foo",
						},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "setName",
						"name": {
							"en": "foo"
						}
					}
				]
			}`,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			output := testutil.RequestData{}
			response := testutil.ResponseData{StatusCode: 200, Body: "{}"}

			client, server := testutil.MockClient(t, response, &output, nil)
			defer server.Close()

			_, err := client.
				WithProjectKey("unittest").
				Stores().
				WithId("1234").
				Post(tC.input).
				Execute(context.TODO())
			assert.Nil(t, err)
			assert.Equal(t, "/unittest/stores/1234", output.URL.Path)
			assert.JSONEq(t, tC.requestBody, output.JSON)
		})
	}
}

func TestStoreUpdateByKey(t *testing.T) {
	testCases := []struct {
		desc        string
		input       platform.StoreUpdate
		requestBody string
	}{
		{
			desc: "Set locale",
			input: platform.StoreUpdate{
				Version: 2,
				Actions: []platform.StoreUpdateAction{
					&platform.StoreSetNameAction{
						Name: &platform.LocalizedString{
							"en": "foo",
						},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "setName",
						"name": {
							"en": "foo"
						}
					}
				]
			}`,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			output := testutil.RequestData{}
			response := testutil.ResponseData{StatusCode: 200, Body: "{}"}

			client, server := testutil.MockClient(t, response, &output, nil)
			defer server.Close()

			_, err := client.
				WithProjectKey("unittest").
				Stores().
				WithKey("test123").
				Post(tC.input).
				Execute(context.TODO())
			assert.Nil(t, err)
			assert.Equal(t, "/unittest/stores/key=test123", output.URL.Path)
			assert.JSONEq(t, tC.requestBody, output.JSON)
		})
	}
}

func TestStoreDelete(t *testing.T) {
	output := testutil.RequestData{}
	response := testutil.ResponseData{StatusCode: 200, Body: "{}"}
	client, server := testutil.MockClient(t, response, &output, nil)
	defer server.Close()

	_, err := client.
		WithProjectKey("unittest").
		Stores().
		WithId("ba8d47e5-6591-4ca2-af4c-d547f062bf35").
		Delete().
		WithQueryParams(platform.ByProjectKeyStoresByIDRequestMethodDeleteInput{
			Version: 2,
		}).
		Execute(context.TODO())
	assert.Nil(t, err)

	params := url.Values{}
	params.Add("version", "2")
	assert.Equal(t, params, output.URL.Query())
	assert.Equal(t, "/unittest/stores/ba8d47e5-6591-4ca2-af4c-d547f062bf35", output.URL.Path)
}

func TestStoreDeleteByKey(t *testing.T) {
	output := testutil.RequestData{}
	response := testutil.ResponseData{StatusCode: 200, Body: "{}"}
	client, server := testutil.MockClient(t, response, &output, nil)
	defer server.Close()

	_, err := client.
		WithProjectKey("unittest").
		Stores().
		WithKey("test123").
		Delete().
		WithQueryParams(platform.ByProjectKeyStoresKeyByKeyRequestMethodDeleteInput{
			Version: 2,
		}).
		Execute(context.TODO())
	assert.Nil(t, err)

	params := url.Values{}
	params.Add("version", "2")
	assert.Equal(t, params, output.URL.Query())
	assert.Equal(t, "/unittest/stores/key=test123", output.URL.Path)
}

func TestStoreGetByID(t *testing.T) {
	response := testutil.Fixture("store.json", 200)
	client, server := testutil.MockClient(t, response, nil, nil)
	defer server.Close()

	store, err := client.
		WithProjectKey("unittest").
		Stores().
		WithKey("test123").
		Get().
		Execute(context.TODO())

	assert.Nil(t, err)
	assert.Equal(t, "ba8d47e5-6591-4ca2-af4c-d547f062bf35", store.ID)
}

func TestStoreGetByKey(t *testing.T) {
	response := testutil.Fixture("store.json", 200)
	client, server := testutil.MockClient(t, response, nil, nil)
	defer server.Close()

	store, err := client.
		WithProjectKey("unittest").
		Stores().
		WithKey("test123").
		Get().
		Execute(context.TODO())

	assert.Nil(t, err)
	assert.Equal(t, "ba8d47e5-6591-4ca2-af4c-d547f062bf35", store.ID)
}

func TestStoreQuery(t *testing.T) {
	output := testutil.RequestData{}
	response := testutil.ResponseData{StatusCode: 200, Body: "{}"}

	client, server := testutil.MockClient(t, response, &output, nil)
	defer server.Close()

	_, err := client.
		WithProjectKey("unittest").
		Stores().
		Get().
		WithQueryParams(platform.ByProjectKeyStoresRequestMethodGetInput{
			Limit: ctutils.IntRef(500),
		}).
		Execute(context.TODO())

	assert.Nil(t, err)
	assert.Equal(t, url.Values{"limit": []string{"500"}}, output.URL.Query())
}
