package main

import (
	"fmt"
)

// Variadic function
func Mul(ints ...int) int {
	result := 1
	for _, i := range ints {
		result *= i
	}
	return result
}

// function that has a function argument
func Doit(f func(a, b int) int, a, b int) int {
	return f(a, b)
}

func Sum(a, b int) int {
	return a + b
}

func Min(a, b int) int {
	return a - b
}

// function that return a function
// Closure
func Accumulator(increment int) func() int {
	i := 0
	return func() int {
		i += increment
		return i
	}
}

func main() {
	fmt.Println(Mul(1, 2, 3, 4, 5))
	fmt.Println(Mul([]int{1, 2, 3, 4, 5}...))

	fmt.Println(Doit(Sum, 5, 4))
	fmt.Println(Doit(Min, 5, 4))

	f1 := Accumulator(2)
	f2 := Accumulator(4)

	for i := 0; i < 4; i++ {
		fmt.Printf("output f1 = %d, output f2 = %d\n", f1(), f2())
	}

}
