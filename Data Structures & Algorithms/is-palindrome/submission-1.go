func isAlpha(b byte) bool {
    return b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' || b >= '0' && b <= '9'
}
func toLower(b byte) byte {
    if b >= 'A' && b <= 'Z' {
        return b-'A'+'a'
    }
    return b
}

func isPalindrome(s string) bool {
    l,r := 0,len(s)-1
    for l<r {
        for !isAlpha(s[l]) && l < r  {
            l++
        }
        for !isAlpha(s[r]) && l<r {
            r--
        }
        if l>=r {break}
        if toLower(s[l])!=toLower(s[r]) {
            return false
        }
        l++
        r--
    }
    return true
}
