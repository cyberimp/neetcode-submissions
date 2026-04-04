func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	res,l,r:=0,0,len(people)-1
	for l<=r {
		remain := limit-people[r]
		r--
		res++
		if l <= r && people[l] <= remain {
			l++
		}
	}
	return res
}
