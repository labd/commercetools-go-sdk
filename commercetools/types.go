package commercetools

type LocalizedString map[string]string

type Reference struct {
	TypeID string `json:"typeId`
	ID     string `json:"id"`
}

type ResourceIdentifier struct {
	ID     string `json:"id,omitempty"`
	Key    string `json:"key,omitempty`
	TypeID string `json:"typeId,omitempty`
}

type UpdateAction interface {
}

type UpdateActions []UpdateAction

type Money struct {
	Type           string `type:"type"`
	CurrencyCode   string `type:"currencyCode"`
	CentAmount     int    `type:"centAmount"`
	FractionDigits int    `type:"fractionDigits"`
}

type HighPrecisionMoney struct {
	Money
	PreciseAmount int `type:"preciseAmount"`
}

type Asset struct {
	ID          string          `type:"id"`
	Key         string          `type:"key"`
	Sources     []AssetSource   `type:"sources"`
	Name        LocalizedString `json:"name"`
	Description LocalizedString `json:"description,omitempty"`
	Tags        []string        `json:"tags,omitempty"`
	// Custom
}

type AssetDraft struct {
	Key         string          `type:"key"`
	Sources     []AssetSource   `type:"sources"`
	Name        LocalizedString `json:"name"`
	Description LocalizedString `json:"description,omitempty"`
	Tags        []string        `json:"tags,omitempty"`
	// Custom
}

type AssetSource struct {
	URI        string          `type:"uri"`
	Key        string          `type:"key"`
	Dimensions AssetDimensions `type:"dimensions,omitempty"`
}

type AssetDimensions struct {
	H int `json:"h"`
	W int `json:"w"`
}
