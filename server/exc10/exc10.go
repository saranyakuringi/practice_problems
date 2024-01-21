/*Write a program, which will find all such numbers between 100 and 300 (both included)
such that each digit of the number is an even number.
The numbers obtained should be printed in a comma-separated sequence on a single line.*/

package exc10

import (
	"fmt"
)

func Getinput() (x, y int) {

	fmt.Print("Enter the x value:")
	fmt.Scan(&x)
	fmt.Print("Enter the y value")
	fmt.Scan(&y)

	return x, y

}

func Exc10(x, y int) []string {
	var values []string
	for i := x; i < y; i++ {
		if i%2 == 0 {
			values = append(values, fmt.Sprint(i))
		}
	}
	fmt.Println(values)
	return values
}
