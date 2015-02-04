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

//helper func for checking 2 components
func checkComponent(a, b Component) bool {
	if a == nil && b == nil {
		return true
	}
	if len(a) == len(b) {
		index := 0
		for index != len(a) {
			if a[index] != b[index] {
				return false
			}
			index++
		}
		return true
	}
	return false
}

// test helper func queryComponent
func TestQueryComponent(t *testing.T) {
	c := Components(getTestComponents())
	if !checkComponent(c.queryComponent(1), c[1]) {
		t.Errorf("expected component %v, got %v", c[1], c.queryComponent(1))
	}
	if !checkComponent(c.queryComponent(2), c[0]) {
		t.Errorf("expected component %v, got %v", c[0], c.queryComponent(2))
	}
	if !checkComponent(c.queryComponent(5), c[2]) {
		t.Errorf("expected component %v, got %v", c[2], c.queryComponent(5))
	}
	// not found digit should return nil instead of empty set, because it would
	// mess with FindQuery
	if !checkComponent(c.queryComponent(0), nil) {
		t.Errorf("expected component %v, got %v", nil, c.queryComponent(0))
	}
}

// test constructor
func TestConstructorNewConnectivity(t *testing.T) {
	test := NewConnectivity(10)
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	index := 0
	for index != len(test.elements) {
		if expected[index] != test.elements[index] {
			t.Errorf("expected %v, got %v", expected[index], test.elements[index])
		}
		index++
	}
}
