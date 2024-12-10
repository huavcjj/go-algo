package algorithm

import "testing"

func TestHeap(t *testing.T) {
	h := NewHeap[int]([]int{50, 27, 19, 62, 15, 23, 17, 10})
	h.Build()
	h.Push(5)
	for h.Size() > 0 {
		t.Log(h.Pop())
	}
}
