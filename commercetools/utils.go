package commercetools

import (
	"net/url"
	"reflect"
	"regexp"
	"time"

	"github.com/appscode/go-querystring/query"
	mapstructure "github.com/mitchellh/mapstructure"
)

func serializeQueryParams(v interface{}) url.Values {
	re := regexp.MustCompile(`([^\]]+)\[([^\]]+)]$`)
	values, _ := query.Values(v)

	newValues := url.Values{}
	for key, values := range values {
		newKey := re.ReplaceAllString(key, "$1.$2")
		for _, value := range values {
			newValues.Add(newKey, value)
		}
	}
	return newValues
}

func toTimeHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if t != reflect.TypeOf(time.Time{}) {
			return data, nil
		}

		switch f.Kind() {
		case reflect.String:
			return time.Parse(time.RFC3339, data.(string))
		case reflect.Float64:
			return time.Unix(0, int64(data.(float64))*int64(time.Millisecond)), nil
		case reflect.Int64:
			return time.Unix(0, data.(int64)*int64(time.Millisecond)), nil
		default:
			return data, nil
		}
		// Convert it by parsing
	}
}

func decodeStruct(input interface{}, result interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: nil,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			toTimeHookFunc()),
		Result: result,
	})
	if err != nil {
		return err
	}

	if err := decoder.Decode(input); err != nil {
		return err
	}
	return err
}
