func hasDuplicate(nums []int) bool {
    hash:=map[int]struct{}{}
    for _,i := range nums{
        if _,ok := hash[i]; ok{ return true }
        hash[i] = struct{}{}
    }
    return false
}
