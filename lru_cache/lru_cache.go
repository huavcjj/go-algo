package algorithm

import "container/list"

type LRUCache struct {
	cache map[int]*list.Element
	lst   *list.List
	cap   int
}

type entry struct {
	key   int
	value string
}

func NewLRUCache(cap int) *LRUCache {
	return &LRUCache{
		cache: make(map[int]*list.Element, cap),
		lst:   list.New(),
		cap:   cap,
	}
}

func (lru *LRUCache) Add(key int, value string) {
	if ele, exists := lru.cache[key]; exists {
		ele.Value.(*entry).value = value
		lru.lst.MoveToFront(ele)
		return
	}

	if len(lru.cache) == lru.cap {
		back := lru.lst.Back()
		if back != nil {
			delete(lru.cache, back.Value.(*entry).key)
			lru.lst.Remove(back)
		}
	}
	ele := lru.lst.PushFront(&entry{key: key, value: value})
	lru.cache[key] = ele
}

func (lru *LRUCache) Get(key int) (string, bool) {
	if ele, exists := lru.cache[key]; exists {
		lru.lst.MoveToFront(ele)
		return ele.Value.(*entry).value, true
	}
	return "", false
}
