package set

import (
	"testing"
)

func TestSet(t *testing.T) {
	s := New[string]()

	s.Add("45")
	s.Add("xdd")

	if !s.Contains("45") {
		t.Error("expected set to contain 45 but it did not")
	}

	if !s.Contains("xdd") {
		t.Error("expected set to contain xdd but it did not")
	}

	if s.Contains("xpp") {
		t.Error("set contains xpp but this was never added")
	}
}
