package collection

import (
	"math/rand"
	"testing"
)

func TestNewBag(t *testing.T) {
	_, err := NewBag()
	if err != nil {
		t.Errorf("error in new bag: %v", err)
	}

}

func getATestBag() *Bag {
	b, _ := NewBag()
	return b
}

func TestBagAdd(t *testing.T) {
	b := getATestBag()
	err := b.Add(int(1))
	if err != nil {
		t.Errorf("got error during add: %v", err)
	}
	if b.Size() != 1 {
		t.Errorf("expected one item, got %v", b.Size())
	}
}

func TestBagReportSize(t *testing.T) {

	titems := []struct {
		number int
		items  []int
	}{
		{3, []int{-6, 8, 677}},
		{2, []int{-2, 4}},
	}

	for _, test := range titems {
		b := getATestBag()
		for _, vals := range test.items {
			b.Add(int(vals))
		}
		if b.Size() != int(test.number) {
			t.Errorf("Size wrong, expected %q, got %q", test.number, b.Size())
		}
	}

	length := rand.Intn(1e2)
	nb := getATestBag()
	for i := 0; i != length; i++ {
		nb.Add(rand.Intn(1e1))
	}
	if nb.Size() != int(length) {
		t.Errorf("random fill: expected %q, got %q", length, nb.Size())
	}
}

func TestBagCalculateMean(t *testing.T) {
	testdata := []struct {
		data []int
		mean float64
	}{
		{[]int{8, 5, 6, 8, 4, 6, 6, 9, 5, 5}, 6.2},
		{[]int{14, 8, 16, 5, 12, -15, 13, -6, 15, 1}, 6.3},
	}
	for _, test := range testdata {
		b := getATestBag()
		for _, number := range test.data {
			b.Add(number)
		}
		if b.Mean() != test.mean {
			t.Errorf("expected %v, got %v", test.mean, b.Mean())
		}
	}
}
