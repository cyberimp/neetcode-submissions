func isAnagram(s string, t string) bool {
	hashS, hashT := make(map[rune]int, 26), make(map[rune]int, 26)

	for _,r := range s {
		hashS[r]++
	}
	for _,r := range t {
		hashT[r]++
	}
	for c := 'a'; c <= 'z'; c++ {
		if hashS[c] != hashT[c] { return false }
	}
	return true
}
