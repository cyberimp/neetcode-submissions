func evalRPN(tokens []string) int {
	eval := map[string]func(a, b int)int{
		"+": func(a, b int) int{return a+b},
		"-": func(a, b int) int{return b-a},
		"*": func(a, b int) int{return a*b},
		"/": func(a, b int) int{return b/a},
	}
	stack := []int{}
	for _, token := range tokens {
		if v,err:= strconv.Atoi(token); err != nil {
			stack = append(stack[:len(stack)-2],eval[token](stack[len(stack)-1],stack[len(stack)-2]))
		} else {
			stack = append(stack,v)
		}
	}
	return stack[0]
}
