package algorithm

import (
	"math/rand/v2"
	"testing"
)

func TestQuickSort(t *testing.T) {
	for i := 0; i < 100; i++ {
		slice := make([]int, 20)
		for j := 0; j < 20; j++ {
			slice[j] = rand.IntN(100)
		}
		QuickSort(slice)

		for j := 0; j < len(slice)-1; j++ {
			if slice[j] > slice[j+1] {
				t.Fail()
			}
		}
	}
}
