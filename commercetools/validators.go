package commercetools

import (
	"fmt"
	"reflect"
	"strings"
)

const tagName = "validate"

// Validator is a generic interface for struct field validators.
type Validator interface {
	// Validate method performs validation and returns result and optional error.
	Validate(interface{}) (bool, error)
}

// DefaultValidator does not perform any validations.
type DefaultValidator struct{}

// Validate verifies the supplied value and returns it's results.
func (v DefaultValidator) Validate(value interface{}) (bool, error) {
	return true, nil
}

// StringValidator validates string length.
type StringValidator struct {
	Min int
	Max int
}

// Validate verifies the supplied value and returns it's results.
func (v StringValidator) Validate(value interface{}) (bool, error) {
	l := len(value.(string))

	if l < v.Min {
		return false, fmt.Errorf("should be at least %v chars long", v.Min)
	}

	if v.Max >= v.Min && l > v.Max {
		return false, fmt.Errorf("should be less than %v chars long", v.Max)
	}

	return true, nil
}

// LStringValidator validates localized string length.
type LStringValidator struct {
	Min int
	Max int
}

// Validate verifies the supplied value and returns it's results.
func (v LStringValidator) Validate(value interface{}) (bool, error) {
	for _, content := range value.(LocalizedString) {
		l := len(content)

		if l < v.Min {
			return false, fmt.Errorf("should be at least %v chars long", v.Min)
		}

		if v.Max >= v.Min && l > v.Max {
			return false, fmt.Errorf("should be less than %v chars long", v.Max)
		}
	}
	return true, nil
}

// NumberValidator performs numerical value validation.
// Its limited to int type for simplicity.
type NumberValidator struct {
	Min int
	Max int
}

// Validate verifies the supplied value and returns it's results.
func (v NumberValidator) Validate(value interface{}) (bool, error) {
	num := value.(int)

	if num < v.Min {
		return false, fmt.Errorf("should be greater than %v", v.Min)
	}

	if v.Max >= v.Min && num > v.Max {
		return false, fmt.Errorf("should be less than %v", v.Max)
	}

	return true, nil
}

// Returns validator struct corresponding to validation type
func getValidatorFromTag(tag string) Validator {
	args := strings.Split(tag, ",")

	switch args[0] {
	case "number":
		validator := NumberValidator{}
		fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
		return validator
	case "string":
		validator := StringValidator{}
		fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
		return validator
	case "lstring":
		validator := LStringValidator{}
		fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
		return validator
	}

	return DefaultValidator{}
}

// ValidateStruct performs actual data validation using validator definitions on the struct
func ValidateStruct(s interface{}) []error {
	errs := []error{}

	// ValueOf returns a Value representing the run-time data
	v := reflect.ValueOf(s)

	for i := 0; i < v.NumField(); i++ {
		// Get the field tag value
		tag := v.Type().Field(i).Tag.Get(tagName)

		// Skip if tag is not defined or ignored
		if tag == "" || tag == "-" {
			continue
		}

		// Get a validator that corresponds to a tag
		validator := getValidatorFromTag(tag)

		// Perform validation
		valid, err := validator.Validate(v.Field(i).Interface())

		// Append error to results
		if !valid && err != nil {
			errs = append(errs, fmt.Errorf("%s %s", v.Type().Field(i).Name, err.Error()))
		}
	}

	return errs
}
