package commercetools_test

import (
	"testing"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/stretchr/testify/assert"
)

type ValidatorTestCase struct {
	Number      int                           `validate:"number,min=1,max=10"`
	String      string                        `validate:"string,min=2,max=12"`
	LString     commercetools.LocalizedString `validate:"lstring,min=2,max=13"`
	Skipped     bool                          `validate:"-"`
	Nonexistent bool                          `validate:"nonexistent"`
}

func TestValidators(t *testing.T) {
	testCases := []struct {
		desc           string
		input          ValidatorTestCase
		expectedLength int
	}{
		{
			desc: "Valid",
			input: ValidatorTestCase{
				Number: 8,
				String: "hello, world",
				LString: commercetools.LocalizedString{
					"en": "Hello, world",
					"nl": "Hallo, wereld",
				},
			},
			expectedLength: 0,
		},
		{
			desc: "Number too high",
			input: ValidatorTestCase{
				Number: 13, // Too high
				String: "hello, world",
				LString: commercetools.LocalizedString{
					"en": "Hello, world",
					"nl": "Hallo, wereld",
				},
			},
			expectedLength: 1,
		},
		{
			desc: "Number too low",
			input: ValidatorTestCase{
				Number: 0, // Too low
				String: "hello, world",
				LString: commercetools.LocalizedString{
					"en": "Hello, world",
					"nl": "Hallo, wereld",
				},
			},
			expectedLength: 1,
		},
		{
			desc: "String too long",
			input: ValidatorTestCase{
				Number: 8,
				String: "hello, world!", // Too long
				LString: commercetools.LocalizedString{
					"en": "Hello, world",
					"nl": "Hallo, wereld",
				},
			},
			expectedLength: 1,
		},
		{
			desc: "String too short",
			input: ValidatorTestCase{
				Number: 8,
				String: "h", // Too shoty
				LString: commercetools.LocalizedString{
					"en": "Hello, world",
					"nl": "Hallo, wereld",
				},
			},
			expectedLength: 1,
		},
		{
			desc: "LString too long",
			input: ValidatorTestCase{
				Number: 8,
				String: "hello, world",
				LString: commercetools.LocalizedString{
					"en": "Hello, world",
					"nl": "Hallo, wereld!", // Too long
				},
			},
			expectedLength: 1,
		},
		{
			desc: "LString too long",
			input: ValidatorTestCase{
				Number: 8,
				String: "hello, world",
				LString: commercetools.LocalizedString{
					"en": "h", // Too short
					"nl": "Hallo, wereld",
				},
			},
			expectedLength: 1,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			errs := commercetools.ValidateStruct(tC.input)
			assert.Equal(t, tC.expectedLength, len(errs))
		})
	}
}
