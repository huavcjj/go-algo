package algorithm

import (
	"fmt"
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
		if v, ok := lru.Get(i); !ok || v != strconv.Itoa(i) {
			fmt.Printf("key %d not found\n", i)
		}
	}
}
