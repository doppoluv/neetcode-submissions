type MinStack struct {
	stack    []int
	minElems []int
}

func Constructor() MinStack {
	st := new(MinStack)
	st.stack = make([]int, 0)
	st.minElems = make([]int, 0)
	return *st
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)

	if len(this.minElems) == 0 || val <= this.minElems[len(this.minElems)-1] {
		this.minElems = append(this.minElems, val)
	}
}

func (this *MinStack) Pop() {
	val := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]

	if this.minElems[len(this.minElems)-1] == val {
		this.minElems = this.minElems[:len(this.minElems)-1]
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minElems[len(this.minElems)-1]
}
