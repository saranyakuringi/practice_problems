package exc12

import "fmt"

//Sum of Slice Elements:
//Write a function that takes a slice of integers and returns the sum of all elements in the slice.

func Sum(n []int) int {
	sum := 0
	for i := 0; i < len(n); i++ {
		sum += n[i]
	}
	return sum
}

// Reverse a Slice:
// Write a function that takes a slice of integers and reverses the order of its elements in-place.

func Reverse_Slice(s []int) []int {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
	return s
}

//Remove Duplicates from Slice:
//Write a function that takes a slice of strings and removes any duplicate elements, keeping only the first occurrence of each unique element.

func Unique(slice []string) []string {
	var result []string
	for i := 0; i < len(slice); i++ {
		isduplicate := false
		for j := 0; j < i; j++ {
			if slice[i] == slice[j] {
				isduplicate = true
				break
			}
		}
		if !isduplicate {
			result = append(result, slice[i])
		}
	}
	fmt.Println("Original Slice:", slice)
	fmt.Println("Unique Elements:", result)
	return result
}

//Find Maximum Element in Slice:
//Write a function that takes a slice of integers and
//returns the maximum element along with its index in the slice.

func Max(num []int) (int, int) {

	var max int = num[0]
	var idx, value int
	for idx, value = range num {
		if value > max {
			max = value
		}
	}
	fmt.Printf("index:%d, value:%d", idx, value)
	return idx, value
}

//Find Minimum Element in Slice:
//Write a function that takes a slice of integers and
//returns the maximum element along with its index in the slice.

func Min(num []int) (int, int) {

	var min int = num[0]
	var idx, value int
	for idx, value = range num {
		if value < min {
			min = value
		}
	}
	fmt.Printf("index:%d, value:%d", idx, value)
	return idx, value
}
