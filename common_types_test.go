package matcher_test

import (
	"testing"

	matcher "github.com/Nivl/gomock-type-matcher"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	// Test the String() method
	m := matcher.String()
	assert.Equal(t, "is of type string", m.String())

	testCases := []struct {
		description    string
		toMatch        interface{}
		expectedResult bool
	}{
		{"a string should work", "str", true},
		{"an int should fail", 1, false},
		{"a float should fail", 1.8, false},
		{"a bool should fail", false, false},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()
			m := matcher.String()
			result := m.Matches(tc.toMatch)
			assert.Equal(t, tc.expectedResult, result, "Matches() did not return the expected value")
		})
	}
}

func TestInt(t *testing.T) {
	// Test the String() method
	m := matcher.Int()
	assert.Equal(t, "is of type int", m.String())

	testCases := []struct {
		description    string
		toMatch        interface{}
		expectedResult bool
	}{
		{"a string should fail", "str", false},
		{"an int should work", 1, true},
		{"an int64 should fail", int64(1), false},
		{"a float should fail", 1.8, false},
		{"a bool should fail", false, false},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()
			m := matcher.Int()
			result := m.Matches(tc.toMatch)
			assert.Equal(t, tc.expectedResult, result, "Matches() did not return the expected value")
		})
	}
}

func TestInt64(t *testing.T) {
	// Test the String() method
	m := matcher.Int64()
	assert.Equal(t, "is of type int64", m.String())

	testCases := []struct {
		description    string
		toMatch        interface{}
		expectedResult bool
	}{
		{"a string should fail", "str", false},
		{"an int should work", 1, false},
		{"an int64 should fail", int64(1), true},
		{"a float should fail", 1.8, false},
		{"a bool should fail", false, false},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()
			m := matcher.Int64()
			result := m.Matches(tc.toMatch)
			assert.Equal(t, tc.expectedResult, result, "Matches() did not return the expected value")
		})
	}
}

func TestInterface(t *testing.T) {
	testCases := []struct {
		description    string
		typ            interface{}
		typStr         string
		toMatch        interface{}
		expectedResult bool
	}{
		{`a string with a string should work`, "typ", "string", "tomatch", true},
		{`an int with a string should fail`, 2, "int", "string", false},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()
			m := matcher.Interface(tc.typ)
			expectedString := "is of type " + tc.typStr
			assert.Equal(t, expectedString, m.String())

			result := m.Matches(tc.toMatch)
			assert.Equal(t, tc.expectedResult, result, "Matches() did not return the expected value")
		})
	}
}
