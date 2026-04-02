func twoSum(nums []int, target int) []int {
	valueMap := map[int]int{}

	for i, val := range nums {
		diff := target - val

		if j, ok := valueMap[diff]; ok{
			return []int{j, i}
		}

		valueMap[val] = i
	}

	return []int{}
}
