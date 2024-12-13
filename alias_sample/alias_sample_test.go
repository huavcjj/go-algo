package algorithm

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestAliasSample(t *testing.T) {
	endpoints := []string{"127.0.0.1:1234", "127.0.0.1:2345", "127.0.0.1:3456"}
	probs := []float64{1, 2, 3}
	useCount := make([]int32, len(endpoints))
	balancer := NewAliasSampler(probs)
	const P = 100
	wg := sync.WaitGroup{}
	wg.Add(P)
	for i := 0; i < P; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				idx := balancer.Sample()
				atomic.AddInt32(&useCount[idx], 1)
				time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
			}
		}()
	}
	wg.Wait()

	fmt.Println(useCount)
}
