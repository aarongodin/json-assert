package jsonAssert

import "testing"

type anyStringMatcher struct{}

func (m anyStringMatcher) validate(t *testing.T, actual any) {
	if _, ok := actual.(string); !ok {
		t.Errorf("expected value to be any string, got: %v", actual)
	}
}
