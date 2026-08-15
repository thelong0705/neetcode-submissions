type MinStack struct {
	sl []int
	m []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if len(this.m) == 0 {
		this.m = append(this.m, val)
	} else {
		last := this.m[len(this.m) - 1]

		if val <= last {
			this.m = append(this.m, val)
		}
	}
	

	this.sl = append(this.sl, val)
	
}

func (this *MinStack) Pop() {
	x := this.Top()
	last := this.m[len(this.m) - 1]

	if x == last {
		this.m = this.m[:len(this.m) - 1]
	}

	this.sl = this.sl[:len(this.sl) - 1]
}

func (this *MinStack) Top() int {
	return this.sl[len(this.sl) - 1]
}

func (this *MinStack) GetMin() int {
	return this.m[len(this.m) - 1]
}


// [2,3,1]