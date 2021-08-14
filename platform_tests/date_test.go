package platform_test

import (
	"encoding/json"
	"testing"

	"github.com/labd/commercetools-go-sdk/platform"
	"github.com/stretchr/testify/assert"
)

func TestDateSerialization(t *testing.T) {

	data := struct {
		String   string         `json:"string"`
		Value    *platform.Date `json:"value"`
		Optional *platform.Date `json:"optional",omitempty`
	}{}
	data.String = "foobar"
	data.Value = &platform.Date{Year: 1983, Month: 5, Day: 10}
	data.Optional = &platform.Date{Year: 1983, Month: 5, Day: 10}

	m, err := json.Marshal(data)
	assert.NoError(t, err)

	expected := `{"string":"foobar","value":"1983-05-10","optional":"1983-05-10"}`
	assert.Equal(t, expected, string(m))
}

func TestDateSerializationEmpty(t *testing.T) {

	data := struct {
		String   string         `json:"string"`
		Value    *platform.Date `json:"value"`
		Optional *platform.Date `json:"optional,omitempty"`
	}{}
	data.String = "foobar"

	m, err := json.Marshal(data)
	assert.NoError(t, err)

	expected := `{"string":"foobar","value":null}`
	assert.Equal(t, expected, string(m))
}

func TestDateDeserialization(t *testing.T) {

	data := struct {
		String   string         `json:"string"`
		Value    *platform.Date `json:"value"`
		Optional *platform.Date `json:"optional",omitempty`
	}{}

	value := `{"string":"foobar","value":"1983-05-10"}`
	err := json.Unmarshal([]byte(value), &data)
	assert.NoError(t, err)
	assert.Equal(t, "foobar", data.String)
	assert.Equal(t, &platform.Date{Year: 1983, Month: 5, Day: 10}, data.Value)
	assert.Equal(t, (*platform.Date)(nil), data.Optional)
}
