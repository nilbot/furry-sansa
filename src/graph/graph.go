package Graph

type Bag struct {
	AdjList []int
}

type Graph struct {
	V   int
	E   int
	Adj *Bag
}

func NewGraph(vertices int) (*Graph, error) {
	if vertices < 0 {
		return nil, "Number of Vertices must be non-negative"
	}
	result := &Graph{V: vertices, E: 0, Adj: &Bag{AdjList: {}}}
	return result, nil
}
