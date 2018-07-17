package testutil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labd/commercetools-go-sdk/common"
	"github.com/labd/commercetools-go-sdk/credentials"
)

func MockClient(t *testing.T, fixture string, output *map[string]interface{}, callback func([]byte)) (*common.Client, *httptest.Server) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture)

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		json.Unmarshal(body, &output)
	}

	ts := httptest.NewServer(http.HandlerFunc(handler))

	client, err := common.NewClient(&common.Config{
		ProjectKey:   "unittest",
		Region:       ts.URL,
		AuthProvider: credentials.NewDummyCredentialsProvider("Bearer unittest"),
	})
	if err != nil {
		t.Fatal(err)
	}

	return client, ts
}
