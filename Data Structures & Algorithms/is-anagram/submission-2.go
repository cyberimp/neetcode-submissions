func isAnagram(s string, t string) bool {
	hashS, hashT := make(map[rune]int, 26), make(map[rune]int, 26)
	for c := 'a'; c <= 'z'; c++ {
		hashS[c] = 1
		hashT[c] = 1
	}
	for _,r := range s {
		hashS[r]++
	}
	for _,r := range t {
		hashT[r]++
	}
	for k,v := range hashS {
		if hashT[k] != v {return false}
	}
	return true
}
