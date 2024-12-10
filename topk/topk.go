package algorithm

import (
	"cmp"
	"slices"
)

func partition[T cmp.Ordered](slice []T, k int, result []T) []T {
	if len(slice) <= k {
		return append(result, slice...)
	}

	pivot := 0
	left, right := 1, len(slice)-1

	for left <= right {
		for left <= right && slice[right] > slice[pivot] {
			right--
		}
		for left <= right && slice[left] <= slice[pivot] {
			left++
		}
		if left < right {
			slice[left], slice[right] = slice[right], slice[left]
		}
	}

	slice[pivot], slice[right] = slice[right], slice[pivot]

	m := len(slice) - right
	n := m - 1
	if m == k {
		result = append(result, slice[right:]...)
	} else if n == k {
		result = append(result, slice[right+1:]...)
	} else if n > k {
		result = partition(slice[right+1:], k, result)
	} else {
		result = append(result, slice[right:]...)
		result = partition(slice[:right], k-m, result)
	}
	return result
}

func TopKByPartition[T cmp.Ordered](list []T, k int) []T {
	if k <= 0 || k > len(list) {
		return list
	}
	result := make([]T, 0, k)
	brr := slices.Clone(list)
	return partition(brr, k, result)
}
