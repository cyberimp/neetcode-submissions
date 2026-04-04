func isAlphaN(b byte) bool {
    return b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' || b >= '0' && b <= '9'
}

func toLower(b byte) byte {
	if b >= 'A' && b <='Z' {
		return b-'A'+'a'
	} else {
		return b
	}
}

func isPalindrome(s string) bool {
	l,r := 0, len(s) - 1
	for l < r {
		for l < r && !isAlphaN(s[l]) {
			l++
		}
		for l < r && !isAlphaN(s[r]) {
			r--
		}
        if l>=r {break}
		if toLower(s[l]) != toLower(s[r]) {
			return false
		}
		l++
		r--
	}
	return true
}
