package easy

func numberOfMatches(n int) int {
	if n <= 1 {
		return 0
	}
	return n/2 + numberOfMatches(n/2+n%2)
}
