func twoSum(nums []int, target int) []int {
    hash := map[int]int{}
	for i, v := range nums {
		delta := target - v
		if j,ok := hash[delta]; ok && i!=j {
			return []int{j,i}
		}
		hash[v] = i
	}
	return []int{}
}
