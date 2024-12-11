package algorithm

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeoutCache(t *testing.T) {
	tc := NewTimeoutCache(5)
	tc.Add(1, "a", 1)
	tc.Add(2, "b", 100)
	tc.Add(3, "c", 100)

	time.Sleep(2 * time.Second)

	for _, key := range []int{1, 2, 3} {
		_, exists := tc.Get(key)
		fmt.Printf("key %d exists %v\n", key, exists)

	}
	fmt.Println()

	for i := 4; i <= 14; i++ {
		tc.Add(i, "value", 200)
	}

	for key := 1; key <= 14; key++ {
		_, exists := tc.Get(key)
		fmt.Printf("key %d exists %v\n", key, exists)
	}

}
