func groupAnagrams(strs []string) [][]string {
    groups := map[[26]int][]string{}
    for _,str := range strs {
        key := [26]int{}
        for _,r := range str {
            key[r - 'a']++
        }
        groups[key] = append(groups[key], str)
    }
    res := [][]string{}
    for _, v := range groups{
        res = append(res, v)
    }
    return res
}
