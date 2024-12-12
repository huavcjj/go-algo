package algorithm

import (
	"container/list"
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	list := &DoubleList[int]{}
	list.PushBack(1)  // 1
	list.PushBack(2)  // 1 -> 2
	list.PushFront(3) // 3 -> 1 -> 2
	list.PushFront(4) // 4 -> 3 -> 1 -> 2

	third := list.Get(2)        //1
	list.InsertAfter(8, third)  // 4 -> 3 -> 1 -> 8 -> 2
	list.InsertBefore(9, third) // 4 -> 3 -> 9 -> 1 -> 8 -> 2
	list.Remove(third)          // 4 -> 3 -> 9 -> 8 -> 2

	fmt.Println(list.Length) // 5

	list.Traverse()        // 4 -> 3 -> 9 -> 8 -> 2
	list.ReverseTraverse() // 2 -> 8 -> 9 -> 3 -> 4
}

func TestStdList(t *testing.T) {
	list := list.New()
	list.PushBack(1)  // 1
	list.PushBack(2)  // 1 -> 2
	list.PushFront(3) // 3 -> 1 -> 2
	list.PushFront(4) // 4 -> 3 -> 1 -> 2

	third := list.Front().Next().Next() // 1
	list.InsertAfter(8, third)          // 4 -> 3 -> 1 -> 8 -> 2
	list.InsertBefore(9, third)         // 4 -> 3 -> 9 -> 1 -> 8 -> 2
	list.Remove(third)                  // 4 -> 3 -> 9 -> 8 -> 2

	fmt.Println(list.Len()) // 5

	expected := []int{4, 3, 9, 8, 2}
	actual := make([]int, 0, list.Len())
	for e := list.Front(); e != nil; e = e.Next() {
		actual = append(actual, e.Value.(int))
	}

	for i, v := range actual {
		if v != expected[i] {
			t.Errorf("expected %v at index %d, got %v", expected[i], i, v)
		}
	}
}
