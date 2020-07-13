package commercetools_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
)

type OutputData struct{}

func errorHandler(statusCode int, returnValue string, encoding string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		w.Write([]byte(returnValue))
	}
}

func TestClientGetBadRequestJson(t *testing.T) {
	body := `,`
	client, server := testutil.MockClient(
		t, "", nil, errorHandler(http.StatusBadRequest, body, "application/json"))
	defer server.Close()

	output := OutputData{}

	err := client.Get(context.TODO(), "/", nil, &output)
	assert.Equal(t, "invalid character ',' looking for beginning of value", err.Error())
}

func TestClientNotFound(t *testing.T) {
	body := ``
	client, server := testutil.MockClient(
		t, "", nil, errorHandler(http.StatusNotFound, body, "application/json"))
	defer server.Close()

	output := OutputData{}

	err := client.Get(context.TODO(), "/", nil, &output)
	assert.Equal(t, "Not Found (404): ResourceNotFound", err.Error())

	ctErr, ok := err.(commercetools.ErrorResponse)
	assert.Equal(t, true, ok)
	assert.Equal(t, 404, ctErr.StatusCode)
}

func TestOAuth2TokenError(t *testing.T) {
	body := `{
		"statusCode": 401,
		"message": "Please provide valid client credentials using HTTP Basic Authentication.",
		"errors": [
			{
				"code": "invalid_client",
				"message": "Please provide valid client credentials using HTTP Basic Authentication."
			}
		],
		"error": "invalid_client",
		"error_description": "Please provide valid client credentials using HTTP Basic Authentication."
	}`

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, body)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	client, err := commercetools.NewClient(&commercetools.ClientConfig{
		ProjectKey: "dummy",
		Credentials: &commercetools.ClientCredentials{
			ClientID:     "foo",
			ClientSecret: "bar",
			Scopes:       []string{"manage_project:dummy"},
		},
		Endpoints: &commercetools.ClientEndpoints{
			API:  server.URL,
			Auth: fmt.Sprintf("%s/oauth/token", server.URL),
		},
	})
	assert.NoError(t, err)

	// Call a random endpoint since the oauth2 exchange happens on the first
	// request.
	draft := &commercetools.APIClientDraft{
		Name:  "test",
		Scope: "manage_orders:my-ct-project-key manage_payments:my-ct-project-key",
	}
	_, err = client.APIClientCreate(context.TODO(), draft)

	// Validate that we have received a valid error response
	assert.Equal(t, "Please provide valid client credentials using HTTP Basic Authentication.", err.Error())

	ctErr, ok := err.(commercetools.ErrorResponse)
	assert.Equal(t, true, ok)
	assert.Equal(t, ctErr.StatusCode, 401)
	assert.Equal(t, ctErr.ErrorDescription, "Please provide valid client credentials using HTTP Basic Authentication.", err.Error())
}

func TestClientRequest(t *testing.T) {
	oAuthBody := `{
		"access_token": "dummy-token",
		"scope": "user",
		"token_type": "bearer",
		"expires_in": 86400
	}`

	productsBody := `{
		"limit": 20,
		"offset": 0,
		"count": 0,
		"total": 0,
		"results": []
	}`

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if r.URL.Path == "/oauth/token" {
			fmt.Fprint(w, oAuthBody)
		} else {
			assert.Equal(t, r.Header["Authorization"], []string{"Bearer dummy-token"})
			fmt.Fprint(w, productsBody)
		}
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	client, err := commercetools.NewClient(&commercetools.ClientConfig{
		ProjectKey: "dummy",
		Credentials: &commercetools.ClientCredentials{
			ClientID:     "foo",
			ClientSecret: "bar",
			Scopes:       []string{"manage_project:dummy"},
		},
		Endpoints: &commercetools.ClientEndpoints{
			API:  server.URL,
			Auth: fmt.Sprintf("%s/oauth/token", server.URL),
		},
	})
	assert.NoError(t, err)

	query, err := client.ProductQuery(context.TODO(), &commercetools.QueryInput{})
	assert.NoError(t, err)
	assert.Equal(t, query.Total, 0)
}

func TestClientRequestCustomHTTPClient(t *testing.T) {
	oAuthBody := `{
		"access_token": "dummy-token",
		"scope": "user",
		"token_type": "bearer",
		"expires_in": 86400
	}`

	productsBody := `{
		"limit": 20,
		"offset": 0,
		"count": 0,
		"total": 0,
		"results": []
	}`

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if r.URL.Path == "/oauth/token" {
			fmt.Fprint(w, oAuthBody)
		} else {
			assert.Equal(t, r.Header["Authorization"], []string{"Bearer dummy-token"})
			fmt.Fprint(w, productsBody)
		}
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	httpClient := &http.Client{}

	client, err := commercetools.NewClient(&commercetools.ClientConfig{
		ProjectKey: "dummy",
		Credentials: &commercetools.ClientCredentials{
			ClientID:     "foo",
			ClientSecret: "bar",
			Scopes:       []string{"manage_project:dummy"},
		},
		Endpoints: &commercetools.ClientEndpoints{
			API:  server.URL,
			Auth: fmt.Sprintf("%s/oauth/token", server.URL),
		},
		HTTPClient: httpClient,
	})
	assert.NoError(t, err)

	query, err := client.ProductQuery(context.TODO(), &commercetools.QueryInput{})
	assert.NoError(t, err)
	assert.Equal(t, query.Total, 0)
}

func TestAuthError(t *testing.T) {
	body := `{
		"statusCode": 403,
		"message": "Insufficient scope",
		"errors": [
			{
				"code": "insufficient_scope",
				"message": "Insufficient scope"
			}
		],
		"error": "insufficient_scope",
		"error_description": "Insufficient scope"
	}`
	client, server := testutil.MockClient(
		t, "", nil, errorHandler(http.StatusForbidden, body, "application/json"))

	defer server.Close()

	output := OutputData{}

	err := client.Get(context.TODO(), "/", nil, &output)

	assert.Equal(t, "Insufficient scope", err.Error())

	ctErr, ok := err.(commercetools.ErrorResponse)
	assert.Equal(t, true, ok)

	ctChildErr, ok := ctErr.Errors[0].(commercetools.InsufficientScopeError)
	assert.Equal(t, true, ok)
	assert.Equal(t, "Insufficient scope", ctChildErr.Message)
}

func TestInvalidJsonError(t *testing.T) {
	body := `{
		"statusCode": 400,
		"message": "Request body does not contain valid JSON.",
		"errors": [
			{
				"code": "InvalidJsonInput",
				"message": "Request body does not contain valid JSON.",
				"detailedErrorMessage": "No content to map due to end-of-input"
			}
		]
	}`
	client, server := testutil.MockClient(
		t, "", nil, errorHandler(http.StatusBadRequest, body, "application/json"))

	defer server.Close()

	output := OutputData{}

	err := client.Get(context.TODO(), "/", nil, &output)

	assert.Equal(t, "Request body does not contain valid JSON.", err.Error())

	ctErr, ok := err.(commercetools.ErrorResponse)
	assert.Equal(t, true, ok)

	ctChildErr, ok := ctErr.Errors[0].(commercetools.InvalidJSONInputError)
	assert.Equal(t, "Request body does not contain valid JSON.", ctChildErr.Error())
	// assert.Equal(t, commercetools.ErrInvalidJSONInput, ctErr.Errors[0].Code())
}

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
				Expand: "taxCategory",
			},
			query: url.Values{
				"expand": []string{"taxCategory"},
			},
			rawQuery: "expand=taxCategory",
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
