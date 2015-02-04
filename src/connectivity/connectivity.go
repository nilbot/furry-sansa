package connectivity

func FindQuery(a, b int, s *[]int) bool {
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
