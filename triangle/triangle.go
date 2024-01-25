package main

import "fmt"

func printTriangle(height int) {
	for i := 1; i <= height; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	var bilangan int
	bilangan = 9

	fmt.Printf("Input: %d\n\n", bilangan)
	fmt.Printf("Output: \n")
	printTriangle(bilangan)
}
