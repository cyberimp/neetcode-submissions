func longestConsecutive(nums []int) int {
    if len(nums) == 0 { return 0 }
    numSet := map[int]struct{}{}
    for _, num := range nums {
        numSet[num] = struct{}{}
    }
    longest := 1
    for k := range numSet {
        var length int
        if _,ok := numSet[k-1]; !ok {
            length = 1
            cur := k + 1
            _,ok = numSet[cur] 
            for ok {
                length++
                cur++
                _,ok = numSet[cur] 
            }
        }
        longest = max(longest,length)
    }
    return longest
}
