package helper

import (
	"testing"
)

// AssertEquals asserts that two variables are equal.
func AssertEquals(t *testing.T, actual, expected interface{}) {
	t.Helper()

	if actual != expected {
		t.Fatalf("expected: %#v\nbut result: %v", expected, actual)
	}
}
