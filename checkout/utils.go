package checkout

// Generated file, please do not change!!!

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

func serializeInput(input interface{}) (io.Reader, error) {
	m, err := json.MarshalIndent(input, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("unable to serialize content: %w", err)
	}
	data := bytes.NewReader(m)
	return data, nil
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
	meta := &mapstructure.Metadata{}
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: meta,
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

	if val, ok := result.(DecodeStruct); ok {
		if raw, ok := input.(map[string]interface{}); ok {
			unused := make(map[string]interface{})
			for _, key := range meta.Unused {
				unused[key] = raw[key]
			}
			val.DecodeStruct(unused)
		}
	}

	return err
}

type DecodeStruct interface {
	DecodeStruct(map[string]interface{}) error
}
