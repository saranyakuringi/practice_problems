/*Write a program that calculates and prints the value according to the given formula:

Q = Square root of [(2 * C * D)/H]

Following are the fixed values of C and H:

C is 50. H is 30.

D is the variable whose values should be input to your program in a comma-separated sequence.

Example:

Let us assume the following comma separated input sequence is given to the program: 100,150,180

The output of the program should be: 18,22,24

1.1. Hints:
use math.Sqrt for the square root

use math.Round for rounding a float operation*/

package prog

import (
	"fmt"
	"math"
)

func Exc6(D []int) []int {

	C := 50
	H := 30

	var data int
	var result []int
	for i := 0; i < len(D); i++ {
		data = int(math.Round(math.Sqrt(float64((2 * C * D[i]) / H))))
		result = append(result, data)
	}

	fmt.Println(result)
	return result
}
