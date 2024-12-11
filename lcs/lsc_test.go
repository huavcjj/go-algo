package algorithm

import (
	"fmt"
	"testing"
)

func TestLCS(t *testing.T) {
	s1 := []byte("AEMQ")
	s2 := []byte("BEQ")
	l1 := LCSRecursive(s1, s2)
	l2 := LCSTopDown(s1, s2)
	l3 := LCSBottomUp(s1, s2)
	l4 := LCSBottomUpWithSpace(s1, s2)
	if l1 != 2 || l2 != 2 || l3 != 2 || l4 != 2 {
		t.Fail()
		fmt.Println(l1, l2, l3, l4)
	}
}

func BenchmarkLCS(b *testing.B) {
	s1 := []byte("AEMQEHELSH")
	s2 := []byte("BEKSFHOPE")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LCSRecursive(s1, s2)
	}
}

func BenchmarkLCSTopDown(b *testing.B) {
	s1 := []byte("AEMQEHELSH")
	s2 := []byte("BEKSFHOPE")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LCSTopDown(s1, s2)
	}
}

func BenchmarkLCSBottomUp(b *testing.B) {
	s1 := []byte("AEMQEHELSH")
	s2 := []byte("BEKSFHOPE")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LCSBottomUp(s1, s2)
	}
}

func BenchmarkLCSBottomUpWithSpace(b *testing.B) {
	s1 := []byte("AEMQEHELSH")
	s2 := []byte("BEKSFHOPE")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LCSBottomUpWithSpace(s1, s2)
	}
}
