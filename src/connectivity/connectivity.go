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

type Connectivity struct {
	Size     int
	elements []int
	Components
}

// Construct a array of number 0...N-1
func NewConnectivity(N int) *Connectivity {
	result := &Connectivity{
		Size: N,
	}
	items := make([]int, N)
	components := make([]Component, N)
	for index := 0; index != N; index++ {
		items[index] = index
		components[index] = []int{index}
	}
	result.elements = items
	result.Components = components
	return result
}
