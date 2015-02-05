package collection

import "fmt"

type Bag struct {
	items []int
	size  int
}

func NewBag() (*Bag, error) {
	ret := &Bag{}
	return ret, nil
}

func (b *Bag) Size() int {
	return b.size
}

func (b *Bag) Add(item interface{}) error {
	switch t := item.(type) {
	case int:
		b.items = append(b.items, t)
		b.size++
	default:
		return fmt.Errorf("type mismatch, expected int")
	}
	return nil
}

func (b *Bag) Mean() float64 {
	sum := 0
	for _, n := range b.items {
		sum += n
	}
	return float64(sum) / float64(len(b.items))
}
