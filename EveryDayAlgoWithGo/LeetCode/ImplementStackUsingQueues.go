package main

type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{[]int{}}
}
func (this *MyStack) Push(x int) {
	this.queue = append(this.queue[:0], append([]int{x}, this.queue[0:]...)...)
}

func (this *MyStack) Pop() int {
	temp := this.queue[0]
	this.queue = this.queue[1:]
	return temp
}

func (this *MyStack) Top() int {
	return this.queue[0]
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}
