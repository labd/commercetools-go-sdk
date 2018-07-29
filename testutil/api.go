package testutil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/commercetools/credentials"
)

type RequestData struct {
	URL    string
	Body   []byte
	Params url.Values
}

func MockClient(
	t *testing.T,
	fixture string,
	output *RequestData,
	callback func([]byte)) (*commercetools.Client, *httptest.Server) {

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture)

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if output != nil {
			output.Body = body
		}
		json.Unmarshal(body, &output)
	}

	ts := httptest.NewServer(http.HandlerFunc(handler))

	client, err := commercetools.NewClient(&commercetools.Config{
		ProjectKey:   "unittest",
		Region:       ts.URL,
		AuthProvider: credentials.NewDummyCredentialsProvider("Bearer unittest"),
	})
	if err != nil {
		t.Fatal(err)
	}

	return client, ts
}
