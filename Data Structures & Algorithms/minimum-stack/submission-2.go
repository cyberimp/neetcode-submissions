type MinStack struct {
	stack []int
	minim []int
}

func Constructor() MinStack {
	return MinStack{stack: []int{}, minim: []int{}}
}

func (this *MinStack) Push(val int) {
	m := math.MaxInt
	if len(this.minim) > 0 {
		m = this.minim[len(this.minim)-1]
	}
	this.stack = append(this.stack, val)
	this.minim = append(this.minim, min(m,val))
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minim = this.minim[:len(this.minim)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minim[len(this.minim)-1]
}
