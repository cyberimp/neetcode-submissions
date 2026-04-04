func maxArea(heights []int) int {
    n := len(heights)
    res := 0
    for l, r := 0, n-1; l<r; {
        res = max(res, min(heights[l], heights[r])*(r-l))

        if heights[l] > heights[r] {
            r--
        } else {
            l++
        }
    }
    return res
}
