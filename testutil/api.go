package testutil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/commercetools/credentials"
)

type RequestData struct {
	URL  url.URL
	Body []byte
	JSON string
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

		// Check if the body is valid JSON
		var dummy map[string]interface{}
		if err := json.Unmarshal(body, &dummy); err != nil {
			log.Println(err)
		}

		if output != nil {
			output.JSON = string(body)
			output.URL = *r.URL
		}
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
