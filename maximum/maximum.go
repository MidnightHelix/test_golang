package main

import "fmt"

func findMax(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	max := arr[0]
	for _, value := range arr {
		if value > max {
			max = value
		}
	}
	return max
}

func main() {
	// Contoh penggunaan fungsi findMax
	numbers := []int{5, 2, 9, 1, 7, 3, 10, 11}
	maxValue := findMax(numbers)

	fmt.Println(maxValue)
}
