package commercetools_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/stretchr/testify/assert"
)

type SampleBaseType interface{}

type SampleStruct struct {
	NonNullableDate time.Time
	NullableDate    *time.Time
	NullDate        *time.Time
}

type SampleContainer struct {
	Data SampleBaseType
}

func TestDateTimeDecode(t *testing.T) {
	body := `{
		"data":{
			"nullableDate": "2020-03-26T13:43:43.364Z",
			"nonNullableDate": "2020-03-26T13:43:43.364Z",
			"nullDate": null
		}
	}`

	var obj SampleContainer
	err := json.Unmarshal([]byte(body), &obj)
	assert.Nil(t, err)

	output := SampleStruct{}

	err = commercetools.Decode(obj.Data, &output)
	assert.Nil(t, err)

	resDate, err := time.Parse(time.RFC3339, "2020-03-26T13:43:43.364Z")
	assert.Nil(t, err)

	assert.Equal(t, resDate, *output.NullableDate)
	assert.Equal(t, resDate, output.NonNullableDate)
	assert.Nil(t, output.NullDate)
}
