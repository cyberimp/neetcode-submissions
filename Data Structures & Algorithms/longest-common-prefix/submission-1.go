func longestCommonPrefix(strs []string) string {
    prefix := []rune(strs[0])
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < len(prefix) {
			prefix = prefix[:len(strs[i])]
		}
		for n,r := range strs[i] {
			if n >= len(prefix) {
				break
			}
			if r != prefix[n] {
				prefix = prefix[:n]
				break
			}
		}
	}
	return string(prefix)
}
