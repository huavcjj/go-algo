package algorithm

import "cmp"

// O(2^n)
func LCSRecursive[T cmp.Ordered](s1, s2 []T) int {
	i, j := len(s1), len(s2)

	if i == 0 || j == 0 {
		return 0
	}
	if s1[i-1] == s2[j-1] {
		return 1 + LCSRecursive(s1[:i-1], s2[:j-1])
	}
	return max(LCSRecursive(s1[:i-1], s2), LCSRecursive(s1, s2[:j-1]))
}

func lcs[T cmp.Ordered](s1, s2 []T, arr [][]int) int {
	i, j := len(s1), len(s2)
	if arr[i][j] != -1 {
		return arr[i][j]
	}
	if s1[i-1] == s2[j-1] {
		arr[i][j] = 1 + lcs(s1[:i-1], s2[:j-1], arr)
	} else {
		arr[i][j] = max(lcs(s1[:i-1], s2[:j], arr), lcs(s1[:i], s2[:j-1], arr))
	}
	return arr[i][j]
}

// O(m*n)
func LCSTopDown[T cmp.Ordered](s1, s2 []T) int {
	i, j := len(s1), len(s2)
	arr := make([][]int, i+1)
	for row := range arr {
		arr[row] = make([]int, j+1)
	}
	for m := 0; m <= i; m++ {
		for n := 0; n <= j; n++ {
			if m == 0 || n == 0 {
				arr[m][n] = 0
			} else {
				arr[m][n] = -1
			}
		}
	}
	return lcs(s1, s2, arr)
}

// O(m*n)
func LCSBottomUp[T cmp.Ordered](s1, s2 []T) int {
	i, j := len(s1), len(s2)
	arr := make([][]int, i+1)
	for row := range arr {
		arr[row] = make([]int, j+1)
	}
	for m := 1; m <= i; m++ {
		for n := 1; n <= j; n++ {
			if s1[m-1] == s2[n-1] {
				arr[m][n] = arr[m-1][n-1] + 1
			} else {
				arr[m][n] = max(arr[m-1][n], arr[m][n-1])
			}
		}
	}
	return arr[i][j]
}

// O(n)
func LCSBottomUpWithSpace[T cmp.Ordered](s1, s2 []T) int {
	if len(s1) < len(s2) {
		s1, s2 = s2, s1
	}
	i, j := len(s1), len(s2)
	prevRow := make([]int, j+1)
	currRow := make([]int, j+1)

	for m := 1; m <= i; m++ {
		for n := 1; n <= j; n++ {
			if s1[m-1] == s2[n-1] {
				currRow[n] = prevRow[n-1] + 1
			} else {
				currRow[n] = max(prevRow[n], currRow[n-1])
			}
		}
		copy(prevRow, currRow)
	}
	return currRow[j]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
