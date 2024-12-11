package algorithm

import (
	"fmt"
	"testing"
)

func TestLongestCommonSubstring(t *testing.T) {

	s1 := "ABCDE"
	s2 := "BCECDE"
	cs1, result1 := LongestCommonSubstring(s1, s2)
	cs2, result2 := LongestCommonSubstringBottomUp(s1, s2)
	cs3, result3 := LongestCommonSubstringBottomUpWithSpace(s1, s2)
	if result1 != result2 || result2 != result3 {
		fmt.Println(s1)
		fmt.Println(s2)
		fmt.Println(cs1, cs2, cs3, result1, result2, result3)
		t.Fail()
	}
}

var (
	s1 = "ABESEFHJSEFKL"
	s2 = "FEUSGHTEOHJEFK"
)

func BenchmarkLongestCommonSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LongestCommonSubstring(s1, s2)
	}
}

// 126.6 ns/op  0 B/op  0 allocs/op

func BenchmarkLongestCommonSubstringBottomUp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LongestCommonSubstringBottomUp(s1, s2)
	}
}

// 521.0 ns/op  214 B/op  15 allocs/op

func BenchmarkLongestCommonSubstringBottomUpWithSpace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LongestCommonSubstringBottomUpWithSpace(s1, s2)
	}
}

// 166.6 ns/op  224 B/op  2 allocs/op

func BenchmarkCopyRaw(b *testing.B) {
	arr1 := []byte(s1)
	arr2 := []byte(s2)
	n := len(arr1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < n; i++ {
			arr2[i] = arr1[i]
		}
	}
}

// 4,023 ns/op  0 B/op  0 allocs/op

func BenchmarkCopyStd(b *testing.B) {
	arr1 := []byte(s1)
	arr2 := []byte(s2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(arr2, arr1)
	}
}

//0.9410 ns/op  0 B/op  0 allocs/op
