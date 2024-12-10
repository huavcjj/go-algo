package algorithm

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 2, 9, 1, 5, 6}, []int{1, 2, 5, 5, 6, 9}},
		{[]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}, []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
		{[]int{10, -1, 2, 0, -5, 8}, []int{-5, -1, 0, 2, 8, 10}},
	}

	for _, test := range tests {
		arr := append([]int(nil), test.input...)
		QuickSort(arr)

		for i, v := range arr {
			if v != test.expected[i] {
				t.Errorf("For input %v, expected %v but got %v", test.input, test.expected, arr)
				break
			}
		}
	}
}
