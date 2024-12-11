package algorithm

func LongestCommonSubstring(s1, s2 string) (string, int) {
	var maxLen int
	var endIndex int

	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			prefixLen := 0
			for k := 0; k < len(s2)-j && k < len(s1)-i; k++ {
				if s1[i+k] == s2[j+k] {
					prefixLen++
					if prefixLen > maxLen {
						maxLen = prefixLen
						endIndex = i + k + 1
					}
				} else {
					break
				}
			}
		}
	}

	if maxLen == 0 {
		return "", 0
	}
	return s1[endIndex-maxLen : endIndex], maxLen
}

func LongestCommonSubstringBottomUp(s1, s2 string) (string, int) {
	dp := make([][]int, len(s1)+1)
	for i := range dp {
		dp[i] = make([]int, len(s2)+1)
	}

	var maxLen int
	var endIndex int

	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			if s1[i] == s2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
				if dp[i+1][j+1] > maxLen {
					maxLen = dp[i+1][j+1]
					endIndex = i + 1
				}
			}
		}
	}
	if maxLen == 0 {
		return "", 0
	}
	return s1[endIndex-maxLen : endIndex], maxLen
}

func LongestCommonSubstringBottomUpWithSpace(s1, s2 string) (string, int) {
	if len(s1) < len(s2) {
		s1, s2 = s2, s1
	}

	prevRow := make([]int, len(s2)+1)
	currRow := make([]int, len(s2)+1)

	var maxLen int
	var endIndex int

	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			if s1[i] == s2[j] {
				currRow[j+1] = prevRow[j] + 1
				if currRow[j+1] > maxLen {
					maxLen = currRow[j+1]
					endIndex = i + 1
				}
			} else {
				currRow[j+1] = 0
			}
		}
		copy(prevRow, currRow)
	}
	if maxLen == 0 {
		return "", 0
	}
	return s1[endIndex-maxLen : endIndex], maxLen
}
