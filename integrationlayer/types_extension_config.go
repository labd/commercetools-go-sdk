package integrationlayer

// Generated file, please do not change!!!

import (
	"encoding/json"
)

type ConfigEntry struct {
	Key string `json:"key"`
	// Withheld for a secret entry.
	Value  interface{} `json:"value,omitempty"`
	Secret bool        `json:"secret"`
}

type ConfigEntryPatch struct {
	Key string `json:"key"`
	// `null` deletes the entry.
	Value  interface{} `json:"value"`
	Secret *bool       `json:"secret,omitempty"`
}

type ExtensionConfig struct {
	Entries                    []ConfigEntry `json:"entries"`
	MaskExtensionGraphQlErrors *bool         `json:"maskExtensionGraphQLErrors,omitempty"`
}

type ExtensionConfigPatch struct {
	Entries []ConfigEntryPatch `json:"entries"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ExtensionConfigPatch) MarshalJSON() ([]byte, error) {
	type Alias ExtensionConfigPatch
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

	if raw["entries"] == nil {
		delete(raw, "entries")
	}

	return json.Marshal(raw)

}
