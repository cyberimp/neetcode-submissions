func majorityElement(nums []int) int {
    hash := map[int]int {}
	maj := int(len(nums)/2)
	for _,n := range nums{
		hash[n]++
		if hash[n] > maj {
			return n
		}
	}
	return nums[0]
}
