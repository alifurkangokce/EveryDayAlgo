package main

type OrderedStream struct {
	stream []string
	ptr    int
}

func Constructor2(n int) OrderedStream {
	return OrderedStream{
		stream: make([]string, n),
		ptr:    0,
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.stream[idKey-1] = value // stream[null,null,ccccc] ptr=0
	result := []string{}

	for this.ptr < len(this.stream) && this.stream[this.ptr] != "" {
		result = append(result, this.stream[this.ptr])
		this.ptr++
	}

	return result
}
