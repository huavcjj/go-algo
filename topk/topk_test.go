package algorithm

import (
	"math/rand/v2"
	"testing"
)

func TestTopKByPartition(t *testing.T) {
	const K = 10
	const L = 20
	for i := 0; i < 100; i++ {
		list := make([]int, L)
		for j := 0; j < L; j++ {
			list[j] = rand.IntN(5 * L)
		}
		topK := TopKByPartition(list, K)

		min := topK[0]
		for _, v := range topK {
			if min > v {
				min = v
			}
		}

		small, equal := 0, 0
		for _, v := range list {
			if v < min {
				small++
			} else if v == min {
				equal++
			}
		}

		if small > len(list)-K || small+equal < len(list)-K+1 {
			t.Fail()
		}
	}
}
