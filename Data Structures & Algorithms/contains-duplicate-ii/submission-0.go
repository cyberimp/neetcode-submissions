func containsNearbyDuplicate(nums []int, k int) bool {
	window := map[int]struct{}{}
	l := 0
	for r := 0; r < len(nums); r++ {
		if r-l > k {
			delete(window, nums[l])
			l++
		}
		if _,ok:=window[nums[r]];ok {
			return true
		}
		window[nums[r]] = struct{}{}
	}

	return false
}
