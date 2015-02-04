package connectivity

func (s *Component) FindQueryInSingleSetNaive(a, b int) bool {
	//naive
	var first, second bool
	for _, it := range *s {
		if a == it {
			first = true
		}
		if b == it {
			second = true
		}
	}
	return first && second
}

type Component []int
type Components []Component

func (s *Components) FindQuery(a, b int) bool {
	for _, c := range *s {
		if c.FindQueryInSingleSetNaive(a, b) {
			return true
		}
	}
	return false
}

func (s *Components) queryComponent(q int) Component {
	for _, c := range *s {
		for _, d := range c {
			if d == q {
				return c
			}
		}
	}
	return nil
}
