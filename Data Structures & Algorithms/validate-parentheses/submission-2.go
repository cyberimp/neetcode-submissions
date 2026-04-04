func isValid(s string) bool {
	valid := map[rune]rune{'[':']', '(':')', '{':'}'}
	var pop rune
	stack := []rune{}
	for _,r := range s {
		if p,ok := valid[r];ok {//нашли открывающую
			stack = append(stack, p)
		} else {//видим закрывающую
			n := len(stack)
			if n == 0 {
				return false
			}
			pop, stack = stack[n-1], stack[:n-1]
			if pop != r {
				return false
			}
		}
	}
	return len(stack) == 0
}
