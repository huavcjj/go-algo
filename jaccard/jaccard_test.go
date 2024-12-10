package algorithm

import (
	"fmt"
	"slices"
	"testing"
)

func TestJaccardUsingHashMap(t *testing.T) {
	l1 := []string{"go", "python", "mysql", "java", "server", "AI"}
	l2 := []string{"c#", "AI", "mysql", "typescript", "sql", "server"}
	fmt.Println(JaccardUsingHashMap(l1, l2))
}

func TestJaccardSorted(t *testing.T) {
	l1 := []string{"go", "python", "mysql", "java", "server", "AI"}
	l2 := []string{"c#", "AI", "mysql", "typescript", "sql", "server"}
	slices.Sort(l1)
	slices.Sort(l2)
	fmt.Println(JaccardSorted(l1, l2))
}
