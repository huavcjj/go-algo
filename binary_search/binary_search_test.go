package algorithm

import (
	"math/rand/v2"
	"sort"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	const L = 100
	for i := 0; i < 100; i++ {
		arr := make([]int, L)
		for j := 0; j < L; j++ {
			arr[j] = rand.IntN(L)
		}
		sort.Ints(arr)
		target := rand.IntN(L)
		index := BinarySearch[int](arr, target)

		if index >= 0 {
			if arr[index] != target {
				t.Fail()
			}
		} else {
			for _, ele := range arr {
				if ele == target {
					t.Fail()
				}
			}
		}
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	const L = 1000
	arr := make([]int, L)
	for i := 0; i < L; i++ {
		arr[i] = i
	}
	sort.Ints(arr)

	target := rand.IntN(L)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BinarySearch(arr, target)
	}
}

func BenchmarkLinearSearch(b *testing.B) {
	const L = 1000
	arr := make([]int, L)
	for i := 0; i < L; i++ {
		arr[i] = i
	}

	target := rand.IntN(L)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinearSearch(arr, target)
	}
}
