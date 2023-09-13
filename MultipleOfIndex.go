
func MultipleOfIndex(ints []int) []int {
	// good luck
	var arr = []int{}
	for i, num := range ints {
		if i == 0 {
			continue
		}
		if num%i == 0 {
			arr = append(arr, num)
		} else {
			continue
		}
	}
	return arr
}