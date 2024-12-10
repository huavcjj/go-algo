package algorithm

import "cmp"

func QuickSort[T cmp.Ordered](arr []T) {
	if len(arr) <= 1 {
		return
	}

	pivot := arr[0]
	left, right := 1, len(arr)-1

	for left <= right {

		for left <= right && arr[right] > pivot {
			right--
		}
		for left <= right && arr[left] <= pivot {
			left++
		}
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}

	arr[0], arr[right] = arr[right], arr[0]

	QuickSort(arr[:right])
	QuickSort(arr[right+1:])
}
