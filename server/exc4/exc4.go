//takes a string of comma-separated numbers and returns a slice of int

package exc4

import (
	"fmt"
	"strconv"
	"strings"
)

func Getinput() (input string) {
	fmt.Print("Enter the low value:")
	fmt.Scan(&input)
	return input
}
func Exc4(input string) []int {
	numbers := strings.Split(input, ",")
	fmt.Println(len(numbers))
	num := make([]int, len(numbers))

	for i, value := range numbers {
		s := strings.Trim(value, " ")
		num[i], _ = strconv.Atoi(s)
	}

	fmt.Println(num)
	return num
}
