package jsonAssert

import (
	"errors"
	"testing"
)

type objectContainingMatcher struct {
	args []any
}

func (m objectContainingMatcher) validate(t *testing.T, actual any) {
	validateObject(t, m.args, actual)
}

type objectMatcher struct {
	args []any
}

func (m objectMatcher) validate(t *testing.T, actual any) {
	validateObject(t, m.args, actual)

	// Check that there are not any additional properties
	propertyNames := make([]string, len(m.args)/2)
	for i := 0; i < len(m.args); i += 2 {
		propertyName, ok := m.args[i].(string)
		if !ok {
			t.Errorf("expected odd-numbered arg to be a property name as a string, got: %v", m.args[i])
		}
		propertyNames[i/2] = propertyName
	}

	for p := range actual.(map[string]any) {
		found := false
		for _, propertyName := range propertyNames {
			if p == propertyName {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected object to not contain additional property: %s", p)
		}
	}
}

// Shared object validation between all object matchers
func validateObject(t *testing.T, args []any, actual any) {
	if len(args)%2 != 0 {
		t.Error(errors.New("expected object matcher to be given an even number of args"))
	}

	actualMap, ok := actual.(map[string]any)
	if !ok {
		t.Errorf("expected value to be an object: %v", actual)
	}

	// Given the properties the user wants to validate, check each of those against what is on the map
	for i := 0; i < len(args)/2; i++ {
		propertyName, ok := args[i*2].(string)
		matcher := args[i*2+1]
		if !ok {
			t.Errorf("expected odd-numbered argument to be a property name as a string")
		}

		value, ok := actualMap[propertyName]
		if !ok {
			t.Errorf("expected object value to have property: %s", propertyName)
		}

		switch objectPropertyMatcher := matcher.(type) {
		case jsonMatcher:
			objectPropertyMatcher.validate(t, value)
		case string:
			if objectPropertyMatcher != value {
				t.Errorf(`expected strings to match; expected "%s", received "%s" `, objectPropertyMatcher, value)
			}
		case int:
			if float64(objectPropertyMatcher) != value {
				t.Errorf("expected numbers to be equal; expected %d, received %v", objectPropertyMatcher, value)
			}
		case nil:
			if value != nil {
				t.Errorf("expected value to be nil, got: %v", value)
			}
		default:
			t.Errorf("expected even-numbered argument to object matcher to be a json matcher or json literal value")
		}
	}
}
