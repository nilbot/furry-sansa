package connectivity

import "testing"

// test set is a struct of set/bag of items, plus
// the test subject -> 2 items
// and expected output -> true or false
type TestSet struct {
	Exp      bool
	SubjectA int
	SubjectB int
	Set      Component
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
		result := test.Set.FindQueryInSingleSetNaive(inputA, inputB)
		if result != exp {
			t.Errorf("expected %q, got %q", exp, result)
		}
	}
}

// we have multiple components (sets)
// FindQuery is expected to Find if the 2 input items are
// in the same component.
// We need a struct for Components just like our previous TestSet
type TestConnectivity struct {
	Expected bool
	A        int
	B        int
	Components
}

func getTestComponents() []Component {
	return []Component{
		Component{2, 3, 6, 7}, Component{1, 4}, Component{5, 8, 9},
	}
}

func getTestSuits() *[]TestConnectivity {
	return &[]TestConnectivity{
		{
			Expected:   true,
			A:          2,
			B:          3,
			Components: getTestComponents(),
		},
		{
			Expected:   true,
			A:          1,
			B:          4,
			Components: getTestComponents(),
		},
		{
			Expected:   false,
			A:          2,
			B:          9,
			Components: getTestComponents(),
		},
	}
}

// test findquery in labyrinth
func TestFindQuery(t *testing.T) {
	testSuits := getTestSuits()
	for _, test := range *testSuits {
		exp := test.Expected
		a := test.A
		b := test.B
		result := test.Components.FindQuery(a, b)
		if result != exp {
			t.Errorf("expected %q, got %q", exp, result)
		}
	}
}
