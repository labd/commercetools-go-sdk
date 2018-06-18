package extension

import "fmt"

type GetSubscriptionInput struct {
	ID  string
	Key string
}

type DeleteSubscriptionInput struct {
	ID      string
	Version int
}

func (svc *Service) GetSubscriptionByID(id string) (*Subscription, error) {
	var result Subscription
	err := svc.client.Get(fmt.Sprintf("subscriptions/%s", id), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (svc *Service) CreateSubscription(draft *SubscriptionDraft) (*Subscription, error) {
	var result Subscription
	err := svc.client.Create("subscriptions", nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (svc *Service) DeleteSubscription(input *DeleteSubscriptionInput) (*Subscription, error) {
	/*
			endpoint := fmt.Sprintf("subscriptions/%s", input.ID)

			params := url.Values{}
			params.Add("version", fmt.Sprintf("%d", input.Version))

			responseData, err := c.doRequest("DELETE", params, nil, endpoint)

			result := Subscription{}
			err = json.Unmarshal(responseData, &result)
		    return &result, err
	*/
	return nil, nil
}
