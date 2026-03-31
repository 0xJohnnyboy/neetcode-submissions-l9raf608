func hasDuplicate(nums []int) bool {
    valueMap := map[int]int{}
	for _, v := range nums {
		valueMap[v]++
		if valueMap[v] > 1 {
			return true
		}
	}

	return false
}
