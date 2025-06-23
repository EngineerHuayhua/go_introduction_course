package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Printf("i: %d, sum: %d\n", i, sum)
	}

	sum = 0 // Reset sum for the next loop
	// Using a while-like loop with for
	for sum < 100 {
		sum += 10
		fmt.Printf("sum: %d\n", sum)
	}

	sum = 0 // Reset sum for the next loop
	// Infinite loop (be careful with this in practice)
	for {
		if sum > 100 {
			fmt.Println("Breaking out of infinite loop")
			break // Break the loop when sum exceeds 100
		}
		sum += 20
		fmt.Printf("sum: %d\n", sum)
	}

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Iterating over array:")
	// Using range to iterate over an array
	for i := range arr {
		fmt.Printf("Index: %d, Value: %d\n", i, arr[i])
	}

	fmt.Println("Iterating over array with value:")
	// Using range to iterate over an array with value
	for i, value := range arr {
		fmt.Printf("Index: %d, Value: %d\n", i, value)
	}

	fmt.Println("Mapas:")
	map2 := map[string]float64{
		"one":   1.0,
		"two":   2.0,
		"three": 3.0,
	}
	for key, value := range map2 {
		fmt.Printf("Key: %s, Value: %.2f\n", key, value)
	}

	fmt.Println()
	map3 := map[string][]int{
		"A": nil,
		"B": {1, 2, 3},
		"C": {4, 5, 6},
	}

	for key, value := range map3 {
		fmt.Println("key:", key)
		for _, v := range value {
			fmt.Printf("Value: %d\n", v)
		}
		fmt.Println()
	}
}
