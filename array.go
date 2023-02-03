package jsonAssert

import "testing"

type anyArrayMatcher struct{}

func (m anyArrayMatcher) validate(t *testing.T, actual any) {
	if _, ok := actual.([]any); !ok {
		t.Errorf("expected value to be any array, got: %v", actual)
	}
}
