func calPoints(operations []string) int {
	points := []int{}
	for _, op := range operations{
		num, err := strconv.Atoi(op)
		n := len(points)
		switch {
		case err == nil:
			points = append(points, num)
		case op == "+":
			num1, num2 := points[n-1],points[n-2]
			points = append(points, num1+num2)
		case op == "D":
			points = append(points, points[n-1]*2)
		case op == "C":
			points = points[:n-1]
		}
	}
	sum := 0
	for _,n := range points {
		sum+=n
	}
	return sum
}
