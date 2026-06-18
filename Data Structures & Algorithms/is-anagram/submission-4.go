func isAnagram(s string, t string) bool {
    var key,lock [26]int
	if len(s) != len(t) { return false }
	for _,r := range s {
		key[r-'a']++
	}
	for _,r := range t {
		key[r-'a']--
	}
	return key == lock
}
