func ReverseSeq(n int) []int {

	result := make([]int, n)
	j := n
	for i := 0; i < n; i++ {
		result[i] = j
		j--
	}
	return result
}