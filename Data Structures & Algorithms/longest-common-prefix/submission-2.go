func longestCommonPrefix(strs []string) string {
	minl := 200
	res := strings.Builder{}
    for _,s := range strs {
		minl = min(len(s), minl)
	}
	if minl == 0 {return ""}
	for i := 0; i<minl; i++ {
		b:=strs[0][i]
		for j := 1; j < len(strs); j++{
			if strs[j][i] != b {
				return res.String()
			}
		}
		res.WriteByte(b)
	}
	return res.String()
}
