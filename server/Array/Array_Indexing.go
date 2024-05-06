package Array

import "math"

func Array_Indexing(array []int, index int) int {
	i := int(index / 2)
	return array[i]
}

//Create a function that takes a vector and returns the sum of all numbers in the vector.

func Sum(array []int) int {
	sum := 0
	for i := 0; i < len(array); i++ {
		sum = sum + array[i]
	}
	return sum
}

//Given two arrays, which represent two sandwiches, return whether both sandwiches use the same type of bread.
//The bread will always be found at the start and end of the array.

func HasSameBread(a1, a2 []string) bool {
	output := false

	for i := 0; i < len(a1); i++ {
		for j := 0; j < len(a2); j++ {
			if a1[0] == a2[0] && a1[len(a1)-1] == a2[len(a2)-1] {
				output = true
			}

		}

	}
	return output
}

//test case
// hasSameBread(
// 	["white bread", "lettuce", "white bread"],
// 	["white bread", "tomato", "white bread"]
//   ) ➞ true

//   hasSameBread(
// 	["brown bread", "chicken", "brown bread"],
// 	["white bread", "chicken", "white bread"]
//   ) ➞ false

//   hasSameBread(
// 	["toast", "cheese", "toast"],
// 	["brown bread", "cheese", "toast"]
//   ) ➞ false

//Help fix all the bugs in the function incrementItems! It is intended to add 1 to every element in the array!

//incrementItems([0, 1, 2, 3]) ➞ [1, 2, 3, 4]

//incrementItems([2, 4, 6, 8]) ➞ [3, 5, 7, 9]

//incrementItems([-1, -2, -3, -4]) ➞ [0, -1, -2, -3]

func IncrementItems(a []int) []int {
	for i := 0; i < len(a); i++ {
		a[i] += 1
	}
	return a
}

//Create a function that accepts an array of numbers and returns the last number in the array.

//getLastItem([1, 2, 3], 3) ➞ 3

//getLastItem([0], 1) ➞ 0

//getLastItem([-1, -3], 2) ➞ -3

func GetLastItem(a []int) int {
	return a[len(a)-1]
}

//Index Multiplier
//Return the sum of all items in an array, where each item is multiplied by its index (zero-based).
//For empty arrays, return 0.

//indexMultiplier([1, 2, 3, 4, 5]) ➞ 40
// (1*0 + 2*1 + 3*2 + 4*3 + 5*4)

//indexMultiplier([-3, 0, 8, -6]) ➞ -2
// (-3*0 + 0*1 + 8*2 + -6*3)

func IndexMultiplier(a []int) int {
	sum := 0
	for i, value := range a {
		a[i] = i * value
		sum = sum + a[i]
	}
	return sum

}

//Less Than 100 Array Remix
//Given an array of numbers, return true if the sum of the array is less than 100; otherwise return false.

//arrayLessThan100([5, 57]) ➞ true

//arrayLessThan100([77, 30]) ➞ false

//arrayLessThan100([0]) ➞ true

func ArrayLessThan100(a []int) bool {
	sum := 0
	output := true
	for i := 0; i < len(a); i++ {
		sum = sum + a[i]
		//fmt.Println(sum)
		if sum > 100 {
			output = false
		}
	}
	return output
}

//Given an array of integers, determine whether the sum of its elements is even or odd.

//The return value should be a string ("odd" or "even").

//If the input array is empty, consider it as an array with a zero ([0]).

//evenOrOdd([0]) ➞ "even"

//evenOrOdd([1]) ➞ "odd"

//evenOrOdd([]) ➞ "even"

//evenOrOdd([0, 1, 5]) ➞ "even"

func EvenOrOdd(a []int) string {
	var output string
	sum := 0

	for i := 0; i < len(a); i++ {
		sum = sum + a[i]
	}

	if sum%2 == 0 {
		output = "even"
	} else {
		output = "odd"
	}
	return output
}

//Mary wants to run a 25-mile marathon. When she attempts to sign up for the marathon,
//she notices the sign-up sheet doesn't directly state the marathon's length.
//Instead, the marathon's length is listed in small, different portions. Help Mary find out how long the marathon actually is.

//Return true if the marathon is 25 miles long, otherwise, return false.

//marathonDistance([1, 2, 3, 4]) ➞ false

// marathonDistance([1, 9, 5, 8, 2]) ➞ true

// marathonDistance([-6, 15, 4]) ➞ true

func MarathonDistance(a []int) bool {
	output := false
	sum := 0

	for i := 0; i < len(a); i++ {
		sum = sum + int(math.Abs(float64(a[i])))
		if sum == 25 {
			output = true
		}
	}

	return output
}
