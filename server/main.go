package main

import (
	"fmt"
)

func main() {
	candies := []int{2, 3, 5, 1, 3}
	var sum []int

	extraCandies := 3
	fmt.Println(len(candies))
	// Create a copy of the original candies array
	updatedCandies := make([]int, len(candies))
	copy(updatedCandies, candies)

	for i := 0; i < len(updatedCandies); i++ {
		updatedCandies[i] = extraCandies + updatedCandies[i]
	}
	sum = append(sum, updatedCandies...)
	fmt.Println(sum)
	var results int

	for _, value := range sum {

		if value >= results {
			results = value
		}
	}

	var result []bool
	for _, value := range sum {
		result = append(result, value == results)
	}
	fmt.Println(result)
}
