package connectivity

import "testing"

func TestFindQueryInSet(t *testing.T) {
	testSets := getTestSets()
	for _, test := range testSets {
		exp := test.Exp
		input := test.Subject
		result := FindQuery(input[0], input[1])
		if result != exp {
			t.Errorf("expected %q, got %q", exp, result)
		}
	}
}
