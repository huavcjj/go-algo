package algorithm

import (
	"container/heap"
	"time"
)

type HeapNode struct {
	key      int
	deadline int
}

type ExpirationHeap []*HeapNode

func (h ExpirationHeap) Len() int {
	return len(h)
}

func (h ExpirationHeap) Less(i, j int) bool {
	return h[i].deadline < h[j].deadline
}

func (h ExpirationHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ExpirationHeap) Push(x interface{}) {
	*h = append(*h, x.(*HeapNode))
}

func (h *ExpirationHeap) Pop() interface{} {
	n := len(*h)
	last := (*h)[n-1]
	*h = (*h)[:n-1]
	return last
}

type TimeoutCache struct {
	cache map[int]interface{}
	hp    ExpirationHeap
	cap   int
}

func NewTimeoutCache(cap int) *TimeoutCache {
	tc := &TimeoutCache{
		cache: make(map[int]interface{}, cap),
		hp:    make(ExpirationHeap, 0, 10),
		cap:   cap,
	}
	heap.Init(&tc.hp)
	go tc.monitor()
	return tc
}

func (tc *TimeoutCache) Add(key int, value interface{}, ttl int) {

	deadline := int(time.Now().Unix()) + ttl
	if len(tc.cache) == tc.cap {
		top := tc.hp[0]
		if top.deadline <= deadline {
			heap.Pop(&tc.hp)
			delete(tc.cache, top.key)
		} else {
			return
		}
	}
	tc.cache[key] = value
	heap.Push(&tc.hp, &HeapNode{key: key, deadline: deadline})
}

func (tc *TimeoutCache) Get(key int) (interface{}, bool) {
	value, exists := tc.cache[key]
	return value, exists
}

func (tc *TimeoutCache) monitor() {
	for {
		if tc.hp.Len() == 0 {
			time.Sleep(100 * time.Millisecond)
			continue
		}
		now := int(time.Now().Unix())
		top := tc.hp[0]
		if top.deadline < now {
			heap.Pop(&tc.hp)
			delete(tc.cache, top.key)
		} else {
			time.Sleep(100 * time.Millisecond)
		}
	}
}
