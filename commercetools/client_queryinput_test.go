package commercetools_test

import (
	"context"
	"net/url"
	"testing"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
)

func TestQueryInput(t *testing.T) {
	testCases := []struct {
		desc     string
		input    *commercetools.QueryInput
		query    url.Values
		rawQuery string
	}{
		{
			desc: "Where",
			input: &commercetools.QueryInput{
				Where: "not (name = 'Peter' and age < 42)",
			},
			query: url.Values{
				"where": []string{"not (name = 'Peter' and age < 42)"},
			},
			rawQuery: "where=not+%28name+%3D+%27Peter%27+and+age+%3C+42%29",
		},
		{
			desc: "Sort",
			input: &commercetools.QueryInput{
				Sort: []string{"name desc", "dog.age asc"},
			},
			query: url.Values{
				"sort": []string{"name desc", "dog.age asc"},
			},
			rawQuery: "sort=name+desc&sort=dog.age+asc",
		},
		{
			desc: "Expand",
			input: &commercetools.QueryInput{
				Expand: []string{"taxCategory", "categories[*]"},
			},
			query: url.Values{
				"expand": []string{"taxCategory", "categories[*]"},
			},
			rawQuery: "expand=taxCategory&expand=categories%5B%2A%5D",
		},
		{
			desc: "Limit",
			input: &commercetools.QueryInput{
				Limit: 20,
			},
			query: url.Values{
				"limit": []string{"20"},
			},
			rawQuery: "limit=20",
		},
		{
			desc: "Offset",
			input: &commercetools.QueryInput{
				Offset: 20,
			},
			query: url.Values{
				"offset": []string{"20"},
			},
			rawQuery: "offset=20",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			output := testutil.RequestData{}

			client, server := testutil.MockClient(t, "{}", &output, nil)
			defer server.Close()

			_, err := client.TaxCategoryQuery(context.TODO(), tC.input)

			assert.Nil(t, err)
			assert.Equal(t, tC.query, output.URL.Query())
			assert.Equal(t, tC.rawQuery, output.URL.RawQuery)
		})
	}
}
