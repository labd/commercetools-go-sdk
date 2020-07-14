package commercetools

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	mapstructure "github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

type GraphQLQuery struct {
	client   *Client
	endpoint string
	query    string
	vars     GraphQLVariablesMap
	headers  map[string]string
}

// NewGraphQLQuery creates a new GraphQLQuery object which can be used to
// execute a GraphQL query
func (client *Client) NewGraphQLQuery(query string) *GraphQLQuery {

	queryObject := GraphQLQuery{
		client:   client,
		endpoint: fmt.Sprintf("%s/%s/graphql", client.Endpoints().API, client.projectKey),
		query:    query,
	}

	return &queryObject
}

// Bind variables to the GraphQL query
func (gql *GraphQLQuery) Bind(key string, value interface{}) {
	if gql.vars == nil {
		gql.vars = make(GraphQLVariablesMap)
	}
	gql.vars[key] = value
}

type GraphQLOption func(o *GraphQLQuery)

func (gql *GraphQLQuery) ForMerchantCenter() GraphQLOption {
	return func(o *GraphQLQuery) {

		o.endpoint = fmt.Sprintf("%s/graphql", o.client.Endpoints().MerchantCenterAPI)
		o.headers = make(map[string]string)
		o.headers["X-Project-Key"] = gql.client.ProjectKey()
		o.headers["X-GraphQL-Target"] = "settings"
	}
}

// Execute the GraphQL query
func (gql *GraphQLQuery) Execute(respData interface{}, opts ...GraphQLOption) error {
	var body bytes.Buffer

	requestObj := GraphQLRequest{
		Query:     gql.query,
		Variables: &gql.vars,
	}

	for _, opt := range opts {
		opt(gql)
	}

	if err := json.NewEncoder(&body).Encode(requestObj); err != nil {
		return errors.Wrap(err, "encode body")
	}
	ctx := context.TODO()
	resp, err := gql.client.getResponse(ctx, "POST", gql.endpoint, nil, &body, gql.headers)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return gql.processResponse(resp, respData)
}

func (gql *GraphQLQuery) processResponse(resp *http.Response, dest interface{}) error {
	body, err := ioutil.ReadAll(resp.Body)

	var output GraphQLResponse
	err = json.Unmarshal(body, &output)
	if err != nil {
		return err
	}

	if output.Errors != nil {
		return output.Errors[0]
	}

	if output.Data != nil {
		mapstructure.Decode(output.Data, &dest)
	}
	return nil
}
