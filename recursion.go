package main

import(
	"fmt"
)

func factorial(i int) int{
	if i <=1 {
		return 1
	}

	return i * factorial(i - 1)
}

func fibonacci(j int) int {
	if j == 0 {
		return 0
	}
	if j == 1 {
		return 1
	}

	return fibonacci(j - 1) + fibonacci(j - 2) 
}

func main() {
	var i int = 10
	fmt.Printf("Factorial of %d is %d", i, factorial(i))

	fmt.Println("\nFibonacci series upto 10...\n")

	var j int
	for j = 0; j < 10; j++ {
		fmt.Printf("%d", fibonacci(j))
	}
}