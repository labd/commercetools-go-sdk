package integrationlayer

// Generated file, please do not change!!!

type IdpMatchStrategy string

const (
	IdpMatchStrategyExternalId IdpMatchStrategy = "externalId"
	IdpMatchStrategyEmail      IdpMatchStrategy = "email"
)

type IdpClaimMapping struct {
	ExternalId string `json:"externalId"`
	Email      string `json:"email"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Phone      string `json:"phone"`
}

/**
*	The IdP login configuration, as read. The client secret is never returned.
 */
type IdpConfig struct {
	Issuer                string           `json:"issuer"`
	TokenEndpoint         string           `json:"tokenEndpoint"`
	RevocationEndpoint    string           `json:"revocationEndpoint"`
	JwksUri               string           `json:"jwksUri"`
	AuthorizationEndpoint *string          `json:"authorizationEndpoint,omitempty"`
	ClientId              string           `json:"clientId"`
	RedirectUri           string           `json:"redirectUri"`
	Claims                IdpClaimMapping  `json:"claims"`
	MatchStrategy         IdpMatchStrategy `json:"matchStrategy"`
	HasClientSecret       bool             `json:"hasClientSecret"`
}

type IdpConfigInput struct {
	Issuer                string           `json:"issuer"`
	TokenEndpoint         string           `json:"tokenEndpoint"`
	RevocationEndpoint    string           `json:"revocationEndpoint"`
	JwksUri               string           `json:"jwksUri"`
	AuthorizationEndpoint *string          `json:"authorizationEndpoint,omitempty"`
	ClientId              string           `json:"clientId"`
	RedirectUri           string           `json:"redirectUri"`
	Claims                IdpClaimMapping  `json:"claims"`
	MatchStrategy         IdpMatchStrategy `json:"matchStrategy"`
	// An empty string keeps the stored secret; the first write needs one.
	ClientSecret string `json:"clientSecret"`
}
