func twoSum(nums []int, target int) []int {
    hash := map[int]int{}
	for i, v := range nums {
		hash[v] = i
	}
	for i_index, v := range nums {
		if j_index,ok := hash[target-v]; ok && i_index!=j_index {
			return []int{min(i_index, j_index), max(i_index, j_index)}
		}
	}
	return []int{}
}
