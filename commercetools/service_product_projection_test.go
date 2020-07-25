package commercetools_test

import (
	"context"
	"net/url"
	"testing"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/stretchr/testify/assert"

	"github.com/labd/commercetools-go-sdk/testutil"
)

func TestProductProjectionSearch(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, "{}", &output, nil)
	defer server.Close()

	queryInput := commercetools.ProductProjectionSearchInput{
		Text: map[string]string{"nl": "foobar"},
	}
	_, err := client.ProductProjectionSearch(context.TODO(), &queryInput)

	assert.Nil(t, err)
	assert.Equal(t, url.Values{"text.nl": []string{"foobar"}}, output.URL.Query())
}
