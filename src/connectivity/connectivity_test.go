package connectivity

import "testing"

// test set is a struct of set/bag of items, plus
// the test subject -> 2 items
// and expected output -> true or false
type TestSet struct {
	Exp      bool
	SubjectA int
	SubjectB int
	Set      []int
}

// prepare test sets
func getTestSets() *[]TestSet {
	return &[]TestSet{
		{
			Exp:      true,
			SubjectA: 2,
			SubjectB: 3,
			Set:      []int{2, 3, 6, 7},
		},
		{
			Exp:      true,
			SubjectA: 1,
			SubjectB: 4,
			Set:      []int{1, 4, 5},
		},
		{
			Exp:      false,
			SubjectA: 2,
			SubjectB: 9,
			Set:      []int{1, 2, 3, 4, 5, 7},
		},
	}
}

func TestFindQueryInSet(t *testing.T) {
	testSets := getTestSets()
	for _, test := range *testSets {
		exp := test.Exp
		inputA := test.SubjectA
		inputB := test.SubjectB
		result := FindQuery(inputA, inputB, &test.Set)
		if result != exp {
			t.Errorf("expected %q, got %q", exp, result)
		}
	}
}
