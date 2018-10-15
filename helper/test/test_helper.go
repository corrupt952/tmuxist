package helper

import (
	"testing"
)

func AssertEquals(t *testing.T, r, e interface{}) {
	t.Helper()

	if r != e {
		t.Fatalf("expected: %#v\nbut result: %v", e, r)
	}
}
