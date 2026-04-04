func groupAnagrams(strs []string) [][]string {
    groups := map[[26]int][]string{}
    for _,s := range strs{
        key:=[26]int{}
        for _,r := range s {
            key[r-'a']++
        }
        groups[key] = append(groups[key], s)
    }
    res := [][]string{}
    for _,v := range groups {
        res = append(res, v)
    }
    return res
}
