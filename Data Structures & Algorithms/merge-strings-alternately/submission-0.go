func mergeAlternately(word1 string, word2 string) string {
	b := strings.Builder{}
	n,m := len(word1), len(word2)
	var i int
	for i = 0; i<n && i<m ; i++{
		b.WriteByte(word1[i])
		b.WriteByte(word2[i])
	}
	if n > m {
		for i<n {
			b.WriteByte(word1[i])
			i++
		}
	}else if n < m {
		for i<m {
			b.WriteByte(word2[i])
			i++
		}
	}
	return b.String()
}
