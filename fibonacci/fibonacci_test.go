package algorithm

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	n := 20
	a, b, c, d := Fib(n), FibTopDown(n), FibBottomUp(n), FibBottomUpWithSpace(n)
	if a != b || b != c || c != d {
		fmt.Println(a, b, c, d)
		t.Fail()
	}
}

func BenchmarkFib(b *testing.B) {
	n := 20
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFibTopDown(b *testing.B) {
	n := 20
	for i := 0; i < b.N; i++ {
		FibTopDown(n)
	}
}

func BenchmarkFibButtomUp(b *testing.B) {
	n := 20
	for i := 0; i < b.N; i++ {
		FibBottomUp(n)
	}
}

func BenchmarkFibButtomUp01(b *testing.B) {
	n := 20
	for i := 0; i < b.N; i++ {
		FibBottomUpWithSpace(n)
	}
}
