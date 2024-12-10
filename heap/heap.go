package algorithm

import (
	"cmp"
	"errors"
	"slices"
)

type Heap[T cmp.Ordered] struct {
	arr []T
}

func NewHeap[T cmp.Ordered](arr []T) *Heap[T] {
	brr := slices.Clone(arr)
	return &Heap[T]{arr: brr}
}

func (heap *Heap[T]) heapifyDown(parent int) {
	left := 2*parent + 1
	if left >= len(heap.arr) {
		return
	}
	minIndex := parent
	if heap.arr[left] < heap.arr[minIndex] {
		minIndex = left
	}
	right := 2*parent + 2
	if right < len(heap.arr) && heap.arr[right] < heap.arr[minIndex] {
		minIndex = right
	}
	if minIndex != parent {
		heap.arr[minIndex], heap.arr[parent] = heap.arr[parent], heap.arr[minIndex]
		heap.heapifyDown(minIndex)
	}
}

func (heap *Heap[T]) heapifyUp(child int) {
	if child == 0 {
		return
	}
	parent := (child - 1) / 2
	if heap.arr[child] < heap.arr[parent] {
		heap.arr[child], heap.arr[parent] = heap.arr[parent], heap.arr[child]
		heap.heapifyUp(parent)
	}
}

func (heap *Heap[T]) Build() {
	n := len(heap.arr)
	if n <= 1 {
		return
	}
	lastIndex := (n - 1) / 2
	for i := lastIndex; i >= 0; i-- {
		heap.heapifyDown(i)
	}
}

func (heap *Heap[T]) Push(value T) {
	heap.arr = append(heap.arr, value)
	heap.heapifyUp(len(heap.arr) - 1)
}

func (heap *Heap[T]) Pop() (T, error) {
	if len(heap.arr) == 0 {
		var v T
		return v, errors.New("heap is empty")
	}
	root := heap.arr[0]
	heap.arr[0] = heap.arr[len(heap.arr)-1]
	heap.arr = heap.arr[:len(heap.arr)-1]

	heap.heapifyDown(0)
	return root, nil
}

func (heap *Heap[T]) Top() (T, error) {
	if len(heap.arr) == 0 {
		return *new(T), errors.New("heap is empty")
	}
	return heap.arr[0], nil
}

func (heap *Heap[T]) Size() int {
	return len(heap.arr)
}

func (heap *Heap[T]) GetAll() []T {
	return slices.Clone(heap.arr)
}

func (heap *Heap[T]) ReplaceTop(value T) {
	if len(heap.arr) == 0 {
		heap.Push(value)
		return
	}
	heap.arr[0] = value
	heap.heapifyDown(0)
}
