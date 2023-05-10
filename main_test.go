package scratchpad

import (
	"testing"
)

func Test_fibonacci(t *testing.T) {
	exp := 55
	act := fibonacci(10)

	if act != exp {
		t.Errorf("expected: %d got: %d", exp, act)
	}
}

func Test_fibonacciDP(t *testing.T) {
	exp := 55
	act := fibonacciDP(10)

	if act != exp {
		t.Errorf("expected: %d got: %d", exp, act)
	}
}

var fib_array []int

func init() {
	fib_array = []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811, 514229, 832040, 1346269, 2178309, 3524578}

	// create and initialise a new array called my_array of 20 random integers from 1 to 10000000

}

func Benchmark_fibonacci(b *testing.B) {

	for i := 0; i < b.N; i++ {
		fibonacci(196418)
		// for x := range fib_array {
		// 	fibonacci(x)
		// }
	}
}

func Benchmark_fibonacciDP(b *testing.B) {

	for i := 0; i < b.N; i++ {
		fibonacciDP(196418)
		// for x := range fib_array {
		// 	fibonacciDP(x)
		// }
	}
}
