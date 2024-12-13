package algorithm

import (
	"container/list"
	"math/rand/v2"
)

type AliasSampler struct {
	accept []float64
	alias  []int
}

func NewAliasSampler(probs []float64) *AliasSampler {
	if len(probs) <= 0 {
		return nil
	}

	sum := 0.0
	for _, prob := range probs {
		sum += prob
	}
	for i := range probs {
		probs[i] /= sum
	}

	n := float64(len(probs))
	arr := make([]float64, len(probs))
	for i, prob := range probs {
		arr[i] = n * prob
	}

	accept := make([]float64, len(probs))
	alias := make([]int, len(probs))

	for i := 0; i < len(probs); i++ {
		accept[i] = 1.0
		alias[i] = -1
	}

	largeStack := list.New()
	smallStack := list.New()

	for i, value := range arr {
		if value < 1 {
			smallStack.PushFront(i)
		} else {
			largeStack.PushFront(i)
		}
	}

	for largeStack.Len() > 0 && smallStack.Len() > 0 {
		smallIndex := smallStack.Front()
		smallStack.Remove(smallIndex)

		largeIndex := largeStack.Front()
		largeStack.Remove(largeIndex)

		smallIdx := smallIndex.Value.(int)
		largeIdx := largeIndex.Value.(int)

		accept[smallIdx] = arr[smallIdx]
		alias[smallIdx] = largeIdx

		arr[largeIdx] -= (1 - arr[smallIdx])

		if arr[largeIdx] < 1 {
			smallStack.PushFront(largeIdx)
		} else if arr[largeIdx] > 1 {
			largeStack.PushFront(largeIdx)
		}

	}
	return &AliasSampler{
		accept: accept,
		alias:  alias,
	}
}

func (as AliasSampler) Sample() int {
	i := rand.IntN(len(as.alias))
	if rand.Float64() < as.accept[i] {
		return i
	}
	return as.alias[i]
}
