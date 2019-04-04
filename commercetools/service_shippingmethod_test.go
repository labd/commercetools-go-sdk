package commercetools_test

import (
	"fmt"
	//"net/url"
	"testing"
	//"time"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
)

func TestShippingMethodCreate(t *testing.T) {

	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, testutil.Fixture("shipping-method.create.json"), &output, nil)
	defer server.Close()

	input := &commercetools.ShippingMethodDraft{
		Name:        "Shipping method",
		Key:         "shipping-method",
		Description: "Shipping method description.",
	}
	fmt.Println(output)

	_, err := client.ShippingMethodCreate(input)
	assert.Nil(t, err)
	expectBody := `{
		"name": "Shipping method",
		"key": "shipping-method",
		"description": "Shipping method description.",
		"isDefault": false,
		"zoneRates": null,
		"taxCategory": null
	}`

	assert.JSONEq(t, expectBody, output.JSON)
}
