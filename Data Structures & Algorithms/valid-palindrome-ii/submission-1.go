func isAlpha(b byte) bool {
	return b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' || b >='0' && b <= '9'
}

func isSame(a, b byte) bool {
	a1, b1 := a, b
	if a >= 'A' && a <= 'Z' {
		a1 = a - 'A' + 'a'
	}
	if b >= 'A' && b <= 'Z' {
		b1 = b - 'A' + 'a'
	}
	return a1 == b1
}

func isPal(s string, l, r int) bool {
	for l < r {
		if !isSame(s[l], s[r]) {
			return false
		}
		l++
		r--
	}
	return true
}

func validPalindrome(s string) bool {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		if !isSame(s[l], s[r]) {
			return isPal(s, l+1, r) || isPal(s, l, r-1)
		}
	}
	return true
}