package commercetools

import (
	"errors"
	"fmt"
	"net/url"
)

// ProductPriceSelection is used to select a price from the Product prices array
// for displaying it to the user.
type ProductPriceSelection struct {
	Currency      string
	Country       string
	CustomerGroup string
	Channel       string
}

// ProductUpdateInput provides the data required to update a product.
type ProductUpdateInput struct {
	ID             string
	Version        int
	PriceSelection ProductPriceSelection
	Actions        []ProductUpdateAction
}

// ProductDeleteInput provides the data required to delete a product.
type ProductDeleteInput struct {
	ID             string
	Key            string
	Version        int
	PriceSelection ProductPriceSelection
}

// ProductURLPath is the commercetools API product path.
const ProductURLPath = "products"

func (ps *ProductPriceSelection) getQueryParameters() url.Values {
	result := url.Values{}

	if ps.Currency != "" {
		result.Add("priceCurrency", ps.Currency)
	}
	if ps.Country != "" {
		result.Add("priceCountry", ps.Country)
	}
	if ps.CustomerGroup != "" {
		result.Add("priceCustomerGroup", ps.CustomerGroup)
	}
	if ps.Channel != "" {
		result.Add("priceChannel", ps.Channel)
	}
	return result
}

func (i *ProductDeleteInput) getQueryParameters() url.Values {
	result := i.PriceSelection.getQueryParameters()
	if i.Version > 0 {
		result.Add("version", string(i.Version))
	}
	return result
}

// ProductCreate will create a new product from a draft, and return the newly
// created product. OAuth2 Scopes: manage_products:{projectKey}
func (client *Client) ProductCreate(draft *ProductDraft) (*Product, error) {
	var result Product
	err := client.Create(ProductURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ProductGetByID will return a product matching the provided ID. OAuth2 Scopes:
// view_products:{projectKey}
func (client *Client) ProductGetByID(id string) (*Product, error) {
	var result Product
	err := client.Get(fmt.Sprintf("%s/%s", ProductURLPath, id), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ProductGetByKey will return a product matching the provided Key. OAuth2 Scopes:
// view_products:{projectKey}
func (client *Client) ProductGetByKey(key string) (*Product, error) {
	return &Product{}, nil
}

// ProductUpdate will update a product matching the provided ID with the
// defined ProductUpdateActions. OAuth2 Scopes: manage_product:{projectKey}
func (client *Client) ProductUpdate(input *ProductUpdateInput) (*Product, error) {
	var result Product

	endpoint := fmt.Sprintf("%s/%s", ProductURLPath, input.ID)
	params := input.PriceSelection.getQueryParameters()

	err := client.Update(endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ProductDeleteByID will delete a product matching the provided ID. OAuth2
// Scopes: manage_products:{projectKey}
func (client *Client) ProductDeleteByID(input *ProductDeleteInput) (*Product, error) {
	if input.ID == "" {
		return nil, errors.New("Missing required field ID")
	}
	var result Product
	endpoint := fmt.Sprintf("%s/%s", ProductURLPath, input.ID)
	params := input.getQueryParameters()
	err := client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ProductDeleteByKey will delete a product matching the provided Key. OAuth2
// Scopes: manage_products:{projectKey}
func (client *Client) ProductDeleteByKey(input *ProductDeleteInput) (*Product, error) {
	if input.Key == "" {
		return nil, errors.New("Missing required field Key")
	}
	var result Product
	endpoint := fmt.Sprintf("%s/key=%s", ProductURLPath, input.Key)
	params := input.getQueryParameters()
	err := client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ProductQuery will query products.
// OAuth2 Scopes: view_products:{projectKey}
func (client *Client) ProductQuery(input *QueryInput) (result *ProductPagedQueryResponse, err error) {
	err = client.Query(ProductURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
