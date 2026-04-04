func hasDuplicate(nums []int) bool {
    hash := map[int]struct{}{}
    for _, n := range nums {
        if _, ok := hash[n]; ok { return true }
        hash[n] = struct{}{}
    }
    return false
}
