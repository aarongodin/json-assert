package jsonAssert

import (
	"encoding/json"
	"testing"
)

// jsonMatcher is any type that allows validating an arbitrary JSON value.
type jsonMatcher interface {
	validate(t *testing.T, actual any)
}

var (
	// AnyString matches any string value, even empty ones.
	AnyString = anyStringMatcher{}
	// AnyArray matches any value with go type []any, regardless of length or contents.
	AnyArray = anyArrayMatcher{}
)

// AssertBytes unmarshals the JSON in `value` and calls `Assert()` on the result.
func AssertBytes(t *testing.T, value []byte, expected jsonMatcher) {
	var raw any
	if err := json.Unmarshal(value, &raw); err != nil {
		t.Error(err)
	}
	Assert(t, raw, expected)
}

// Assert checks a value against the passed `jsonMatcher`.
func Assert(t *testing.T, value any, expected jsonMatcher) {
	expected.validate(t, value)
}

// Validates that the value is any object matching the specified properties and values.
// The object may contain additional properties.
func ObjectContaining(args ...any) jsonMatcher {
	return objectContainingMatcher{args: args}
}

// Validates that the value is an object matching the specified properties and values.
// Values must be a jsonMatcher or valid JSON literal (string, int, []any, map[string]any).
// The object may not contain additional properties.
func Object(args ...any) jsonMatcher {
	return objectMatcher{args: args}
}
