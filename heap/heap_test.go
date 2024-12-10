package algorithm

import (
	"container/heap"
	"testing"
)

func TestHeap(t *testing.T) {
	h := NewHeap[int]([]int{50, 27, 19, 62, 15, 23, 17, 10})
	h.Build()
	h.Push(5)
	for h.Size() > 0 {
		t.Log(h.Pop())
	}
}

func TestStdHeap(t *testing.T) {
	pq := make(PriorityQueue[int], 0, 10)
	pq.Push(&Item[int]{Value: 50})
	pq.Push(&Item[int]{Value: 27})
	pq.Push(&Item[int]{Value: 19})
	heap.Init(&pq)
	heap.Push(&pq, &Item[int]{Value: 62})

	for pq.Len() > 0 {
		t.Log(heap.Pop(&pq))
	}

}
