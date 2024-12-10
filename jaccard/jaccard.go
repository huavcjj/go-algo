package algorithm

import (
	"cmp"
)

func JaccardUsingHashMap[T cmp.Ordered](collection1, collection2 []T) float64 {
	if len(collection1) == 0 || len(collection2) == 0 {
		return 0.0
	}

	map1 := make(map[T]struct{}, len(collection1))
	for _, item := range collection1 {
		map1[item] = struct{}{}
	}

	intersection := 0
	for _, item := range collection2 {
		if _, ok := map1[item]; ok {
			intersection++
		}
	}

	union := len(collection1) + len(collection2) - intersection
	return float64(intersection) / float64(union)
}

func JaccardSorted[T cmp.Ordered](collection1, collection2 []T) float64 {
	if len(collection1) == 0 || len(collection2) == 0 {
		return 0.0
	}

	intersection := 0
	i, j := 0, 0
	for i < len(collection1) && j < len(collection2) {
		if collection1[i] == collection2[j] {
			intersection++
			i++
			j++
		} else if collection1[i] < collection2[j] {
			i++
		} else {
			j++
		}
	}

	union := len(collection1) + len(collection2) - intersection
	return float64(intersection) / float64(union)
}
