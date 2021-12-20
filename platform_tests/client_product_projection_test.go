package platform_test

import (
	"context"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/labd/commercetools-go-sdk/platform"
	"github.com/labd/commercetools-go-sdk/testutil"
)

func TestProductProjectionSearch(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, testutil.ResponseData{StatusCode: 200, Body: "{}"}, &output, nil)
	defer server.Close()

	input := platform.ByProjectKeyProductProjectionsSearchRequestMethodGetInput{
		Text:   map[string][]string{"text.nl": []string{"foobar"}},
		Filter: []string{"category.id:foo", "category.id:bar"},
	}
	_, err := client.WithProjectKey("unittest").
		ProductProjections().
		Search().
		Get().
		Offset(20). // Removed by next func
		WithQueryParams(input).
		Limit(30).
		Execute(context.TODO())

	assert.Nil(t, err)
	assert.Equal(t, url.Values{
		"text.nl": {"foobar"},
		"filter":  {"category.id:foo", "category.id:bar"},
		"limit":   {"30"},
	}, output.URL.Query())
}
