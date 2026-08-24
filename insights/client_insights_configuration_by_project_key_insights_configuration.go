package insights

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInsightsConfigurationRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Retrieve the Platform Insights configuration for your Project.
*
 */
func (rb *ByProjectKeyInsightsConfigurationRequestBuilder) Get() *ByProjectKeyInsightsConfigurationRequestMethodGet {
	return &ByProjectKeyInsightsConfigurationRequestMethodGet{
		url:    fmt.Sprintf("/%s/insights-configuration", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Perform specific update actions, such as enabling or disabling an existing Platform Insights configuration. After you enable or disable a Platform Insights configuration, the action is performed in an [eventually consistent](/general-concepts#eventual-consistency) way.
*
 */
func (rb *ByProjectKeyInsightsConfigurationRequestBuilder) Post(body ConfigurationUpdate) *ByProjectKeyInsightsConfigurationRequestMethodPost {
	return &ByProjectKeyInsightsConfigurationRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/insights-configuration", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Create or replace a Platform Insights configuration. After you create a Platform Insights configuration, it becomes active in an [eventually consistent](/general-concepts#eventual-consistency) way.
*
 */
func (rb *ByProjectKeyInsightsConfigurationRequestBuilder) Put(body ProjectConfigurationDraft) *ByProjectKeyInsightsConfigurationRequestMethodPut {
	return &ByProjectKeyInsightsConfigurationRequestMethodPut{
		body:   body,
		url:    fmt.Sprintf("/%s/insights-configuration", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Delete the Platform Insights configuration details of your Project. After you delete a Platform Insights configuration, it is deleted in an [eventually consistent](/general-concepts#eventual-consistency) way.
*
 */
func (rb *ByProjectKeyInsightsConfigurationRequestBuilder) Delete() *ByProjectKeyInsightsConfigurationRequestMethodDelete {
	return &ByProjectKeyInsightsConfigurationRequestMethodDelete{
		url:    fmt.Sprintf("/%s/insights-configuration", rb.projectKey),
		client: rb.client,
	}
}
