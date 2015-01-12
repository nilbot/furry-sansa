package eliminate

import (
	"testing"
)

func TestTailRecursion(t *testing.T) {
	cases := []struct {
		first, second, expected int
	}{
		{3, -6, 8},
		{2, -2, 4},
	}
	for _, c := range cases {
		got := TailRecursion(c.first, c.second)
		if got != c.expected {
			t.Errorf("Recursion(%q,%q) == %q, expected %q", c.first, c.second, got, c.expected)
		}
	}
}
