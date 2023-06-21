package frontend

// Generated file, please do not change!!!

type FrontasticActionRequestBuilder struct {
	client *Client
}

func (rb *FrontasticActionRequestBuilder) WithNamespaceValueWithActionValue(namespace string, action string) *FrontasticActionByNamespaceByActionRequestBuilder {
	return &FrontasticActionByNamespaceByActionRequestBuilder{
		namespace: namespace,
		action:    action,
		client:    rb.client,
	}
}
