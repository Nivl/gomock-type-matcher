package matcher_test

import (
	"testing"

	matcher "github.com/Nivl/gomock-type-matcher"
	"github.com/stretchr/testify/assert"
)

func TestType(t *testing.T) {
	testCases := []struct {
		description    string
		typ            string
		toMatch        interface{}
		expectedResult bool
	}{
		{`a string with a string should work`, "string", "tomatch", true},
		{`an int with an int should work`, "int", 4, true},
		{`a bool with a string should fail`, "bool", "string", false},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()
			m := matcher.Type(tc.typ)
			expectedString := "is of type " + tc.typ
			assert.Equal(t, expectedString, m.String())

			result := m.Matches(tc.toMatch)
			assert.Equal(t, tc.expectedResult, result, "Matches() did not return the expected value")
		})
	}
}
