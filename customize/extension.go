package customize

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/labd/commercetools-go-sdk/common"
)

type ExtensionDeleteInput struct {
	ID      string
	Version int
}

type ExtensionUpdateInput struct {
	ID      string
	Version int
	Actions common.UpdateActions
}

type Extension struct {
	ID             string               `json:"id"`
	Version        int                  `json:"version"`
	Key            string               `json:"key"`
	Destination    ExtensionDestination `json:"destination"`
	Triggers       []Trigger            `json:"messages"`
	CreatedAt      time.Time            `json:"createdAt"`
	LastModifiedAt time.Time            `json:"lastModifiedAt"`
}

type ExtensionDraft struct {
	Key         string               `json:"key"`
	Destination ExtensionDestination `json:"destination"`
	Triggers    []Trigger            `json:"messages"`
}

type ExtensionDestination interface {
}

// ExtensionDestinationHTTP implementation

type ExtensionDestinationHTTP struct {
	URL string `json:"url"`
}

func (ed *ExtensionDestinationHTTP) Type() string {
	return "HTTP"
}

func (ed *ExtensionDestinationHTTP) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type string `json:"type"`
		ExtensionDestinationHTTP
	}{
		Type: ed.Type(),
		ExtensionDestinationHTTP: *ed,
	})
}

// ExtensionDestinationAWSLambda implementation

type ExtensionDestinationAWSLambda struct{}

func (ed *ExtensionDestinationAWSLambda) Type() string {
	return "AWSLambda"
}

func (ed *ExtensionDestinationAWSLambda) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type string `json:"type"`
		ExtensionDestinationAWSLambda
	}{
		Type: ed.Type(),
		ExtensionDestinationAWSLambda: *ed,
	})
}

type ExtensionDestinationAuthentication struct {
}

type ExtensionDestinationAuthenticationAzure struct{}

func (ed *ExtensionDestinationAuthenticationAzure) Type() string {
	return "AzureFunctions"
}

type ExtensionDestinationAuthenticationAuth struct {
}

func (ed *ExtensionDestinationAuthenticationAuth) Type() string {
	return "AuthorizationHeader"
}

type Trigger struct {
	ResourceTypeID string   `json:"resourceTypeId"`
	Actions        []string `json:"actions"`
}

func (svc *Service) ExtensionGetByID(id string) (*Extension, error) {
	var result Extension
	err := svc.client.Get(fmt.Sprintf("extensions/%s", id), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (svc *Service) ExtensionCreate(draft *ExtensionDraft) (*Extension, error) {
	var result Extension
	err := svc.client.Create("extensions", nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (svc *Service) ExtensionUpdate(input *ExtensionUpdateInput) (*Extension, error) {
	var result Extension

	endpoint := fmt.Sprintf("products/%s", input.ID)
	err := svc.client.Update(endpoint, nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (svc *Service) ExtensionDeleteByID(id string, version int) (*Extension, error) {
	var result Extension
	endpoint := fmt.Sprintf("extensions/%s", id)
	params := url.Values{}
	params.Set("version", string(version))
	err := svc.client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (svc *Service) ExtensionDeleteByKey(key string, version int) (*Extension, error) {
	var result Extension
	endpoint := fmt.Sprintf("extensions/key=%s", key)
	params := url.Values{}
	params.Set("version", string(version))
	err := svc.client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return &result, nil
}
