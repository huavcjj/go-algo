package algorithm

import (
	"cmp"
	"fmt"
)

type ListNode[T cmp.Ordered] struct {
	Value T
	Prev  *ListNode[T]
	Next  *ListNode[T]
}

type DoubleList[T cmp.Ordered] struct {
	Head   *ListNode[T]
	Tail   *ListNode[T]
	Length int
}

func (list *DoubleList[T]) Traverse() {
	curr := list.Head
	for curr != nil {
		fmt.Printf("%v ", curr.Value)
		curr = curr.Next
	}
}

func (list *DoubleList[T]) ReverseTraverse() {
	curr := list.Tail
	for curr != nil {
		fmt.Printf("%v ", curr.Value)
		curr = curr.Prev
	}
}

func (list *DoubleList[T]) PushFront(value T) {
	node := &ListNode[T]{Value: value}
	if list.Head == nil {
		list.Head = node
		list.Tail = node
	} else {
		node.Next = list.Head
		list.Head.Prev = node
		list.Head = node
	}
	list.Length++
}

func (list *DoubleList[T]) PushBack(value T) {
	node := &ListNode[T]{Value: value}
	if list.Tail == nil {
		list.Head = node
		list.Tail = node
	} else {
		node.Prev = list.Tail
		list.Tail.Next = node
		list.Tail = node
	}
	list.Length++
}

func (list *DoubleList[T]) InsertAfter(value T, node *ListNode[T]) {
	if node == nil {
		return
	}
	newNode := &ListNode[T]{Value: value}
	newNode.Prev = node
	newNode.Next = node.Next

	if node.Next != nil {
		node.Next.Prev = newNode
	} else {
		list.Tail = newNode
	}
	node.Next = newNode
	list.Length++
}

func (list *DoubleList[T]) InsertBefore(value T, node *ListNode[T]) {
	if node == nil {
		return
	}
	newNode := &ListNode[T]{Value: value}
	newNode.Prev = node.Prev
	newNode.Next = node

	if node.Prev != nil {
		node.Prev.Next = newNode
	} else {
		list.Head = newNode
	}
	node.Prev = newNode
	list.Length++
}

func (list *DoubleList[T]) Remove(node *ListNode[T]) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		list.Head = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		list.Tail = node.Prev
	}
	list.Length--
}

func (list *DoubleList[T]) Find(value T) *ListNode[T] {
	curr := list.Head
	for curr != nil {
		if curr.Value == value {
			return curr
		}
		curr = curr.Next
	}
	return nil
}

func (list *DoubleList[T]) Get(index int) *ListNode[T] {
	if index < 0 || index >= list.Length {
		return nil
	}
	if index < list.Length/2 {
		curr := list.Head
		for i := 0; i < index; i++ {
			curr = curr.Next
		}
		return curr
	} else {
		curr := list.Tail
		for i := list.Length - 1; i > index; i-- {
			curr = curr.Prev
		}
		return curr
	}
}

//container/list
