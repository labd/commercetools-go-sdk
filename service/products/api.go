package products

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/labd/commercetools-go-sdk/commercetools"
)

type PriceSelection struct {
	Currency      string
	Country       string
	CustomerGroup string
	Channel       string
}

type UpdateInput struct {
	ID             string
	Version        int
	PriceSelection PriceSelection
	Actions        commercetools.UpdateActions
}

type DeleteInput struct {
	ID             string
	Key            string
	Version        int
	PriceSelection PriceSelection
}

func (ps *PriceSelection) GetQueryParameters() url.Values {
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

func (i *DeleteInput) GetQueryParameters() url.Values {
	result := i.PriceSelection.GetQueryParameters()
	if i.Version > 0 {
		result.Add("version", string(i.Version))
	}
	return result
}

func (svc *Service) Create(draft *ProductDraft) (*Product, error) {
	var result Product
	err := svc.client.Create("products", nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (svc *Service) GetByID(id string) (*Product, error) {
	var result Product
	err := svc.client.Get(fmt.Sprintf("products/%s", id), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (svc *Service) GetByKey(key string) (*Product, error) {
	return &Product{}, nil
}

func (svc *Service) Update(input *UpdateInput) (*Product, error) {
	var result Product

	endpoint := fmt.Sprintf("products/%s", input.ID)
	params := input.PriceSelection.GetQueryParameters()

	err := svc.client.Update(endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (svc *Service) DeleteByID(input *DeleteInput) (*Product, error) {
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

func (svc *Service) DeleteByKey(input *DeleteInput) (*Product, error) {
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
