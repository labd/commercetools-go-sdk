package integrationlayer_test

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/labd/commercetools-go-sdk/integrationlayer"
)

// newClient serves handler and returns a client for its "test" Project.
func newClient(t *testing.T, handler http.HandlerFunc) *integrationlayer.ByProjectKeyRequestBuilder {
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)
	client, err := integrationlayer.NewClient(&integrationlayer.ClientConfig{URL: server.URL})
	require.NoError(t, err)
	return client.WithProjectKey("test")
}

func TestSettingsGet(t *testing.T) {
	client := newClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/test/main/settings", r.URL.Path)
		_, _ = w.Write([]byte(`{"project":{"label":"Shop","language":"nl-NL","currency":"EUR","country":"NL"},
			"platform":{"regionId":"europe-west1"},"presentment":{"languages":["nl-NL"],"countries":["NL"],"currencies":["EUR"]},
			"version":3}`))
	})

	view, err := client.Main().Settings().Get().Execute(context.Background())
	require.NoError(t, err)
	assert.Equal(t, "NL", *view.Project.Country)
	assert.Equal(t, "europe-west1", view.Platform.RegionId)
	assert.Equal(t, []string{"EUR"}, view.Presentment.Currencies)
}

func TestConfigPatchDeletesWithNull(t *testing.T) {
	var body map[string]any
	client := newClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method)
		data, _ := io.ReadAll(r.Body)
		require.NoError(t, json.Unmarshal(data, &body))
		_, _ = w.Write([]byte(`{"entries":[]}`))
	})

	_, err := client.Main().Extension().Config().Patch(integrationlayer.ExtensionConfigPatch{
		Entries: []integrationlayer.ConfigEntryPatch{{Key: "OLD_KEY", Value: nil}},
	}).Execute(context.Background())
	require.NoError(t, err)
	assert.Equal(t, map[string]any{"entries": []any{map[string]any{"key": "OLD_KEY", "value": nil}}}, body)
}

func TestAllowlistPutSendsABareArray(t *testing.T) {
	var body []string
	client := newClient(t, func(w http.ResponseWriter, r *http.Request) {
		data, _ := io.ReadAll(r.Body)
		require.NoError(t, json.Unmarshal(data, &body))
		_, _ = w.Write([]byte(`{"allow":["cdn.contentful.com"],"version":4}`))
	})

	result, err := client.Main().Extension().Allowlist().Put([]string{"cdn.contentful.com"}).Execute(context.Background())
	require.NoError(t, err)
	assert.Equal(t, []string{"cdn.contentful.com"}, body)
	assert.Equal(t, 4, result.Version)
}

func TestErrors(t *testing.T) {
	client := newClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/test/main/idp" {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(`{"error":"No IdP configured"}`))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error":"invalid_request","error_description":"all four settings are required"}`))
	})

	_, err := client.Main().Idp().Get().Execute(context.Background())
	assert.True(t, errors.Is(err, integrationlayer.ErrNotFound))

	_, err = client.Main().Settings().Put(integrationlayer.ProjectSettings{}).Execute(context.Background())
	var response integrationlayer.ErrorResponse
	require.True(t, errors.As(err, &response))
	assert.Equal(t, "invalid_request: all four settings are required", err.Error())
}
