func characterReplacement(s string, k int) int {
	res := 0
	charSet := map[byte]struct{}{}
	
	for i := range s {
		charSet[s[i]] = struct{}{}
	}

	for c := range charSet {
		count, l := 0, 0
		for r := range s {
			if s[r] == c {
				count++
			}
		

			for (r - l + 1) - count > k {
				if s[l] == c {
					count--
				}
				l++
			}
			
			res = max(res, r-l+1)
		}
	}

	return res
}
