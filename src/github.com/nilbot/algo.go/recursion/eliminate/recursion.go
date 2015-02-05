package eliminate

func TailRecursion(m, n int) int {
	if m == 0 || n == 0 {
		return 1
	} else {
		return 2 * TailRecursion(m-1, n+1)
	}
}

func TailRecEliminated(m, n int) int {
	ret := 1
	for m != 0 && n != 0 {
		ret *= 2
		m--
		n++
	}
	return ret
}
