package main

import "container/heap"

type any = interface{}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
	hp *IntHeap
	k  int
}

func ConstructorRes(k int, nums []int) KthLargest {
	intHeap := IntHeap(nums)
	kl := KthLargest{k: k, hp: &intHeap}
	heap.Init(kl.hp)
	for kl.hp.Len() > k {
		heap.Pop(kl.hp)
	}
	return kl
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.hp, val)
	for this.hp.Len() > this.k {
		heap.Pop(this.hp)
	}
	return (*this.hp)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
