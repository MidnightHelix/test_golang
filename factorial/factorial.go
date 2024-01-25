package main

import "fmt"

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var bilangan int
	bilangan = 5

	result := factorial(bilangan)
	fmt.Printf("Input: %d\n", bilangan)
	fmt.Printf("Output: %d\n", result)

}
