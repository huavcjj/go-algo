package algorithm

func Fib(n int) int {
	if n <= 1 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func FibTopDown(n int) int {
	if n <= 1 {
		return n
	}
	cache := make(map[int]int)
	cache[0] = 0
	cache[1] = 1

	return f(n, cache)
}

func f(n int, cache map[int]int) int {
	if val, exists := cache[n]; exists {
		return val
	}
	cache[n] = f(n-1, cache) + f(n-2, cache)
	return cache[n]
}

func FibBottomUp(n int) int {
	if n <= 1 {
		return n
	}
	arr := make([]int, n+1)
	arr[0] = 0
	arr[1] = 1
	for i := 2; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n]
}

func FibBottomUpWithSpace(n int) int {
	if n <= 1 {
		return n
	}
	prev2 := 0
	prev1 := 1
	for i := 2; i <= n; i++ {
		curr := prev1 + prev2
		prev2 = prev1
		prev1 = curr
	}
	return prev1
}

func Steps(n int) int {
	if n <= 2 {
		return n
	}
	prev2 := 1
	prev1 := 2
	for i := 3; i <= n; i++ {
		curr := prev1 + prev2
		prev2 = prev1
		prev1 = curr
	}
	return prev1
}
