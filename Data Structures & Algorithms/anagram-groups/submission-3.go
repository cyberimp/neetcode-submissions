func groupAnagrams(strs []string) [][]string {
	res := map[[26]int][]string{}
	for _,str := range strs{
		key := [26]int{}
		for _,r := range str {
			key[r-'a']++
		}
		res[key] = append(res[key], str)
	}
	arr_res := [][]string{}
	for _,v := range res{
		arr_res = append(arr_res, v)
	}
	return arr_res
}
