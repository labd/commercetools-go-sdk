package integrationlayer

// Generated file, please do not change!!!

import (
	"encoding/json"
)

type Allowlist struct {
	Allow []string `json:"allow"`
	// Set by commercetools; read-only.
	Deny []string `json:"deny"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj Allowlist) MarshalJSON() ([]byte, error) {
	type Alias Allowlist
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["deny"] == nil {
		delete(raw, "deny")
	}

	return json.Marshal(raw)

}

type AllowlistReplaceResult struct {
	Allow   []string `json:"allow"`
	Version int      `json:"version"`
}
