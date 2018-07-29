package customize_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/labd/commercetools-go-sdk/commercetools/customize"
	"github.com/labd/commercetools-go-sdk/testutil"
)

func TestExtensionCreate(t *testing.T) {
	client, server := testutil.MockClient(t, fixture("extension.azure.json"), nil, nil)
	defer server.Close()
	svc := customize.New(client)

	draft := &customize.ExtensionDraft{
		Key: "test",
		Destination: customize.ExtensionDestinationHTTP{
			URL: "http://example.com",
		},
		Triggers: []customize.Trigger{
			customize.Trigger{
				ResourceTypeID: "product",
				Actions:        []string{"Create"},
			},
		},
	}

	_, err := svc.ExtensionCreate(draft)
	assert.Equal(t, nil, err)
}

func TestExtensionUpdate(t *testing.T) {
	var output map[string]interface{}

	client, server := testutil.MockClient(t, fixture("extension.azure.json"), &output, nil)
	defer server.Close()
	svc := customize.New(client)

	input := &customize.ExtensionUpdateInput{}

	fmt.Println(output)

	_, err := svc.ExtensionUpdate(input)
	assert.Equal(t, nil, err)
}

func TestExtensionDeleteByID(t *testing.T) {
	client, server := testutil.MockClient(t, fixture("extension.azure.json"), nil, nil)
	defer server.Close()
	svc := customize.New(client)

	_, err := svc.ExtensionDeleteByID("1234", 2)
	assert.Equal(t, nil, err)
}

func TestExtensionDeleteByKey(t *testing.T) {
	client, server := testutil.MockClient(t, fixture("extension.azure.json"), nil, nil)
	defer server.Close()
	svc := customize.New(client)

	_, err := svc.ExtensionDeleteByKey("1234", 2)
	assert.Equal(t, nil, err)
}
