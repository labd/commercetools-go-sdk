package frontend

// Generated file, please do not change!!!

func (c *Client) Frontastic() *FrontasticRequestBuilder {
	return &FrontasticRequestBuilder{
		client: c,
	}
}
func (c *Client) Api() *ApiRequestBuilder {
	return &ApiRequestBuilder{
		client: c,
	}
}
