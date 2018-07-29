package credentials

type DummyCredentialsProvider struct {
	token string
}

func NewDummyCredentialsProvider(token string) AuthProvider {
	return &DummyCredentialsProvider{token}
}

func (c *DummyCredentialsProvider) GetAuthToken() (string, error) {
	return c.token, nil
}
