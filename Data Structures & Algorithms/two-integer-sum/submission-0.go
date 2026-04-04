func twoSum(nums []int, target int) []int {
    diff := map[int]int{}
    for j,n := range nums {
        curdiff := target - n
        if i,ok := diff[curdiff]; ok{
            return []int{i,j}
        }
        diff[n] = j
    }
    return []int{0,1}
}
