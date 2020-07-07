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
}

// NewGraphQLQuery creates a new GraphQLQuery object which can be used to
// execute a GraphQL query
func (client *Client) NewGraphQLQuery(query string) *GraphQLQuery {

	queryObject := GraphQLQuery{
		client:   client,
		endpoint: fmt.Sprintf("%s/%s/graphql", client.url, client.projectKey),
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

// Execute the GraphQL query
func (gql *GraphQLQuery) Execute(respData interface{}) error {
	var body bytes.Buffer

	requestObj := GraphQLRequest{
		Query:     gql.query,
		Variables: &gql.vars,
	}

	if err := json.NewEncoder(&body).Encode(requestObj); err != nil {
		return errors.Wrap(err, "encode body")
	}
	ctx := context.TODO()
	resp, err := gql.client.getResponse(ctx, "POST", gql.endpoint, nil, &body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

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
