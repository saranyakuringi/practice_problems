package Array

import "strings"

//Sum Greater Than Five
//Write a function that returns the sum of elements greater than 5, in the given array.
// sumFive([1, 5, 20, 30, 4, 9, 18]) ➞ 77

// sumFive([1, 2, 3, 4]) ➞ 0

// sumFive([10, 12, 28, 47, 55, 100]) ➞ 252

func SumFive(a []int) int {
	sum := 0

	for i := 0; i < len(a); i++ {
		if a[i] > 5 {
			sum += a[i]
		}
	}

	return sum
}

//Invert an Array
// invertArray([1, 2, 3, 4, 5]) ➞ [-1, -2, -3, -4, -5]

// invertArray([1, -2, 3, -4, 5]) ➞ [-1, 2, -3, 4, -5]

// invertArray([]) ➞ []

func InvertArray(a []int) []int {

	for i := 0; i < len(a); i++ {
		a[i] = -a[i]
	}

	return a
}

// The Forbidden Letter
// Given a letter and an array of words, return whether the letter does not appear in any of the words.

// forbiddenLetter('r', ["rock", "paper", "scissors"]) ➞ false

// forbiddenLetter('a', ["spoon", "fork", "knife"]) ➞ true

// forbiddenLetter('m', []) ➞ true

func ForbiddenLetter(a []string, l string) bool {
	output := true
	for _, value := range a {
		if strings.Contains(value, l) {
			output = false
		}
	}
	return output
}

// Find the Bug: Checking Even Numbers
// // Create a function that takes in an array and returns true if all its values are even, and false otherwise.
// checkAllEven([1, 2, 3, 4]) ➞ false

// checkAllEven([2, 4, 6]) ➞ true

// checkAllEven([5, 6, 8, 10]) ➞ false

// checkAllEven([-2, 2, -2, 2]) ➞ true

func CheckAllEven(a []int) bool {
	output := true
	for i := 0; i < len(a); i++ {
		if a[i]%2 != 0 {
			output = false
			break
		}
	}
	return output
}

//Where is Bob!?!
//Write a function that searches an array of names (unsorted) for the name "Bob" and returns the location in the array. If Bob is not in the array, return -1.
//findBob(["Jimmy", "Layla", "Bob"]) ➞ 2

//findBob(["Bob", "Layla", "Kaitlyn", "Patricia"]) ➞ 0

//findBob(["Jimmy", "Layla", "James"]) ➞ -1

func FindBob(a []string) int {
	//var output int
	for i, value := range a {
		if value == "Bob" {
			return i
		}

	}
	return -1
}

//Given an array of numbers, create a function which returns the same array but with each element's index in the array added to itself.
//This means you add 0 to the number at index 0, add 1 to the number at index 1, etc...

// addIndexes([0, 0, 0, 0, 0]) ➞ [0, 1, 2, 3, 4]

// addIndexes([1, 2, 3, 4, 5]) ➞ [1, 3, 5, 7, 9]

// addIndexes([5, 4, 3, 2, 1]) ➞ [5, 5, 5, 5, 5]

func AddIndexes(a []int) []int {
	for i, value := range a {
		a[i] = value + i
	}
	return a
}

//Create a function that takes three parameters where:

//x is the start of the range (inclusive).
//y is the end of the range (inclusive).
//n is the divisor to be checked against.
//Return an ordered array with numbers in the range that are divisible by the third parameter n.
//Return an empty array if there are no numbers that are divisible by n.

// arrayOperation(1, 10, 3) ➞ [3, 6, 9]

// arrayOperation(7, 9, 2) ➞ [8]

// arrayOperation(15, 20, 7) ➞ []

func ArrayOperation(a []int) []int {
	x := a[0]
	y := a[1]
	n := a[2]
	var output []int
	for i := x; i < y; i++ {
		if i%n == 0 {
			output = append(output, i)
		}
	}
	return output
}
