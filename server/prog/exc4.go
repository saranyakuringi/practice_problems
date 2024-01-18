//takes a string of comma-separated numbers and returns a slice of int

package prog

import (
	"fmt"
	"strconv"
	"strings"
)

func Exc4(input string) []int {
	//input := "12,13,14,15"

	numbers := strings.Split(input, ",")
	fmt.Println(len(numbers))
	//fmt.Println(numbers)

	num := make([]int, len(numbers))

	for i, value := range numbers {
		s := strings.Trim(value, " ")
		num[i], _ = strconv.Atoi(s)
	}
	fmt.Println(num)
	return num
}
