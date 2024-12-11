package algorithm

import (
	"testing"
)

func TestEditDistance(t *testing.T) {
	s1 := []byte("word")
	s2 := []byte("world")
	if EditDistance(s1, s2) != 1 {
		t.Fail()
	}
	if EditDistanceTopDown(s1, s2) != 1 {
		t.Fail()
	}
	if EditDistanceBottomUp(s1, s2) != 1 {
		t.Fail()
	}
	if EditDistanceBottomUpWithSpace(s1, s2) != 1 {
		t.Fail()
	}

	s1 = []byte("horse")
	s2 = []byte("ros")
	if EditDistance(s1, s2) != 3 {
		t.Fail()
	}
	if EditDistanceTopDown(s1, s2) != 3 {
		t.Fail()
	}
	if EditDistanceBottomUp(s1, s2) != 3 {
		t.Fail()
	}
	if EditDistanceBottomUpWithSpace(s1, s2) != 3 {
		t.Fail()
	}
}

func BenchmarkEditDistance(b *testing.B) {
	s1 := []byte("kitten")
	s2 := []byte("sitting")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EditDistance(s1, s2)
	}
}

func BenchmarkEditDistanceTopDown(b *testing.B) {
	s1 := []byte("kitten")
	s2 := []byte("sitting")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EditDistanceTopDown(s1, s2)
	}
}

func BenchmarkEditDistanceBottomUp(b *testing.B) {
	s1 := []byte("kitten")
	s2 := []byte("sitting")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EditDistanceBottomUp(s1, s2)
	}
}

func BenchmarkEditDistanceBottomUpWithSpace(b *testing.B) {
	s1 := []byte("kitten")
	s2 := []byte("sitting")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EditDistanceBottomUpWithSpace(s1, s2)
	}
}
