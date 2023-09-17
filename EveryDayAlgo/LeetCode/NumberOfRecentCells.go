package main

type RecentCounter struct {
	calls []int
}

// func Constructor() RecentCounter {
// 	return RecentCounter{}
// }

func (this *RecentCounter) Ping(t int) int {
	var trim int
	for _, call := range this.calls {
		if t-3000 <= call {
			break
		}
		trim++
	}
	this.calls = this.calls[trim:]
	this.calls = append(this.calls, t)
	return len(this.calls)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
