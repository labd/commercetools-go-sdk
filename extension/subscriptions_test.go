package extension_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labd/commercetools-go-sdk/common"
	"github.com/labd/commercetools-go-sdk/credentials"
	"github.com/labd/commercetools-go-sdk/extension"
)

func fixture(path string) string {
	b, err := ioutil.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func TestSubscriptionCreate(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("subscription.example.json"))

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		var result map[string]interface{}
		json.Unmarshal(body, &result)
	}

	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	client, err := common.NewClient(&common.Config{
		ProjectKey:   "unittest",
		Region:       ts.URL,
		AuthProvider: credentials.NewDummyCredentialsProvider("Bearer unittest"),
	})
	if err != nil {
		t.Fatal(err)
	}
	svc := extension.New(client)

	draft := &extension.SubscriptionDraft{
		Key: "test",
		Destination: extension.SubscriptionAWSSQSDestination{
			Type:         "SQS",
			QueueURL:     "http://example.com/",
			AccessKey:    "A1234567890",
			AccessSecret: "S1234567800",
			Region:       "eu-central-1",
		},
		Messages: []extension.MessageSubscription{
			extension.MessageSubscription{
				ResourceTypeID: "product",
			},
		},
	}

	_, err = svc.SubscriptionCreate(draft)
}
