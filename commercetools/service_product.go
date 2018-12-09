package commercetools

import (
	"errors"
	"fmt"
	"net/url"
)

type ProductPriceSelection struct {
	Currency      string
	Country       string
	CustomerGroup string
	Channel       string
}

type ProductUpdateInput struct {
	ID             string
	Version        int
	PriceSelection ProductPriceSelection
	Actions        []ProductUpdateAction
}

type ProductDeleteInput struct {
	ID             string
	Key            string
	Version        int
	PriceSelection ProductPriceSelection
}

func (ps *ProductPriceSelection) GetQueryParameters() url.Values {
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

func (i *ProductDeleteInput) GetQueryParameters() url.Values {
	result := i.PriceSelection.GetQueryParameters()
	if i.Version > 0 {
		result.Add("version", string(i.Version))
	}
	return result
}

// Service contains client information and bundles all actions.
type ProductService struct {
	client *Client
}

func (svc *ProductService) Create(draft *ProductDraft) (*Product, error) {
	var result Product
	err := svc.client.Create("products", nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (svc *ProductService) GetByID(id string) (*Product, error) {
	var result Product
	err := svc.client.Get(fmt.Sprintf("products/%s", id), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (svc *ProductService) GetByKey(key string) (*Product, error) {
	return &Product{}, nil
}

func (svc *ProductService) Update(input *ProductUpdateInput) (*Product, error) {
	var result Product

	endpoint := fmt.Sprintf("products/%s", input.ID)
	params := input.PriceSelection.GetQueryParameters()

	err := svc.client.Update(endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (svc *ProductService) DeleteByID(input *ProductDeleteInput) (*Product, error) {
	if input.ID == "" {
		return nil, errors.New("Missing required field ID")
	}
	var result Product
	endpoint := fmt.Sprintf("products/%s", input.ID)
	params := input.GetQueryParameters()
	err := svc.client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (svc *ProductService) DeleteByKey(input *ProductDeleteInput) (*Product, error) {
	if input.Key == "" {
		return nil, errors.New("Missing required field Key")
	}
	var result Product
	endpoint := fmt.Sprintf("products/key=%s", input.Key)
	params := input.GetQueryParameters()
	err := svc.client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return &result, nil
}
