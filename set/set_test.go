package set

import (
	"slices"
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

	var s2 Set[int]
	s2.Add(2)
	if !s2.Contains(2) {
		t.Error("expected set to contain 2 but it did not")
	}

	var s3 Set[uint32]
	if s3.Contains(2) {
		t.Error("expected set to not contain 2 but it did")
	}
}

func TestSetRange(t *testing.T) {
	expected := []int{1, 3, 4}

	s := New(expected...)

	var out []int

	yield := func(a int) bool {
		out = append(out, a)
		return true
	}

	s.All()(yield)

	slices.Sort(out)

	if !slices.Equal(out, expected) {
		t.Errorf("iteration produced %v instead of the expected %v", out, expected)
	}
}
