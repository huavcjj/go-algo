package algorithm

import (
	"strconv"
	"testing"
)

func TestLRUCache(t *testing.T) {
	lru := NewLRUCache(10)

	for i := 0; i < 10; i++ {
		lru.Add(i, strconv.Itoa(i)) //9 8 7 6 5 4 3 2 1 0
	}

	for i := 0; i < 10; i += 2 {
		lru.Get(i) //8 6 4 2 0 9 7 5 3 1
	}

	for i := 10; i < 15; i++ {
		lru.Add(i, strconv.Itoa(i)) //14 13 12 11 10 8 6 4 2 0
	}

	for i := 0; i < 10; i++ {
		v, ok := lru.Get(i)
		if i%2 == 0 { // 0, 2, 4, 6, 8 はキャッシュに残るべき
			if !ok || v != strconv.Itoa(i) {
				t.Errorf("key %d not found or incorrect value: got %v, want %v", i, v, strconv.Itoa(i))
			}
		} else { // 1, 3, 5, 7, 9 はキャッシュから削除されるべき
			if ok {
				t.Errorf("key %d should not be found, got %v", i, v)
			}
		}
	}
}
