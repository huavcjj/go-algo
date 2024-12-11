package algorithm

import "cmp"

func EditDistance[T cmp.Ordered](s1, s2 []T) int {
	i := len(s1)
	j := len(s2)
	if i == 0 {
		return j
	}
	if j == 0 {
		return i
	}

	if s1[i-1] == s2[j-1] {
		return EditDistance(s1[:i-1], s2[:j-1])
	}
	return 1 + min(EditDistance(s1[:i-1], s2[:j-1]), EditDistance(s1[:i-1], s2[:j]), EditDistance(s1[:i], s2[:j-1]))
}

func EditDistanceTopDown[T cmp.Ordered](s1, s2 []T) int {
	i, j := len(s1), len(s2)
	arr := make([][]int, i+1)
	for row := range arr {
		arr[row] = make([]int, j+1)
	}
	for m := 0; m <= i; m++ {
		for n := 0; n <= j; n++ {
			if m == 0 {
				arr[m][n] = n
			} else if n == 0 {
				arr[m][n] = m
			} else {
				arr[m][n] = -1
			}
		}
	}
	return ed(s1, s2, arr)
}

func ed[T cmp.Ordered](s1, s2 []T, arr [][]int) int {
	i, j := len(s1), len(s2)
	if arr[i][j] != -1 {
		return arr[i][j]
	}
	if s1[i-1] == s2[j-1] {
		arr[i][j] = ed(s1[:i-1], s2[:j-1], arr)
	} else {
		arr[i][j] = 1 + min(ed(s1[:i-1], s2[:j-1], arr), ed(s1[:i-1], s2[:j], arr), ed(s1[:i], s2[:j-1], arr))
	}
	return arr[i][j]
}

// O(m*n)
func EditDistanceBottomUp[T cmp.Ordered](s1, s2 []T) int {
	i, j := len(s1), len(s2)
	arr := make([][]int, i+1)
	for row := range arr {
		arr[row] = make([]int, j+1)
	}
	for m := 0; m <= i; m++ {
		for n := 0; n <= j; n++ {
			if m == 0 {
				arr[m][n] = n
			} else if n == 0 {
				arr[m][n] = m
			} else if s1[m-1] == s2[n-1] {
				arr[m][n] = arr[m-1][n-1]
			} else {
				arr[m][n] = 1 + min(arr[m-1][n-1], arr[m-1][n], arr[m][n-1])
			}
		}
	}
	return arr[i][j]
}

// O(n)
func EditDistanceBottomUpWithSpace[T cmp.Ordered](s1, s2 []T) int {
	if len(s1) < len(s2) {
		s1, s2 = s2, s1
	}
	i, j := len(s1), len(s2)
	prevRow := make([]int, j+1)
	currRow := make([]int, j+1)
	for col := 0; col <= j; col++ {
		prevRow[col] = col
	}
	for row := 1; row <= i; row++ {
		currRow[0] = row
		for col := 1; col <= j; col++ {
			if s1[row-1] == s2[col-1] {
				currRow[col] = prevRow[col-1]
			} else {
				currRow[col] = 1 + min(
					prevRow[col-1],
					prevRow[col],
					currRow[col-1],
				)
			}
		}
		prevRow, currRow = currRow, prevRow
	}
	return prevRow[j]
}

func min(nums ...int) int {
	if len(nums) == 0 {
		panic("min requires at least one argument")
	}
	m := nums[0]
	for _, n := range nums[1:] {
		if n < m {
			m = n
		}
	}
	return m
}
