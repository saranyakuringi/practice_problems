/*
Write a program which can compute the factorial of a given numbers.
The results should be printed in a comma-separated sequence on a single line.

Suppose the following input is supplied to the program: 8
Then, the output should be: 40320
*/

package prog

import "fmt"

func Exc2(n int) int {

	if n < 0 {
		fmt.Println("factorial of negitive numbers is not allowed")
		return 0
	}
	if n == 0 || n == 1 {
		return 1
	}
	return n * Exc2(n-1)
}
