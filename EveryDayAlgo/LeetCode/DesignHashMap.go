package main

type MyHashMap struct {
	arr []int
}

/** Initialize your data structure here. */
func ConstructorHH() MyHashMap {
	m := new(MyHashMap)
	m.arr = make([]int, 10000, 1000001)
	for i, _ := range m.arr {
		m.arr[i] = -1
	}
	return *m
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	l := len(this.arr)
	if key >= l {
		newlen := 2 * key
		if newlen > 1000001 {
			newlen = 1000001
		}
		this.arr = this.arr[:newlen]
		for i := l; i < len(this.arr); i++ {
			this.arr[i] = -1
		}
	}
	this.arr[key] = value
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	if key >= len(this.arr) {
		return -1
	}
	return this.arr[key]
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	if key >= len(this.arr) {
		return
	}
	this.arr[key] = -1
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
