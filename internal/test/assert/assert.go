package assert

import "testing"

func Equal(t *testing.T, observed interface{}, expected interface{}) {
	t.Helper()
	if observed != expected {
		t.Fatalf("%s != %s", observed, expected)
	}
}

func True(t *testing.T, value bool, msgAndArgs ...interface{}) bool {
	t.Helper()
	if !value {
		t.Fatalf("Should be true")
	}

	return value
}
