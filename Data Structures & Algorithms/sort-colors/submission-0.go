func sortColors(nums []int) {
    var buckets [3]int
	for _,n := range nums {
		buckets[n]++
	}
	p:=0
	for i:= 0; i<3; i++{
		for j := 0; j < buckets[i]; j++ {
			nums[p] = i
			p++
		}
	}
}
