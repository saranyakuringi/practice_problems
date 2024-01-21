/*Write a program that accepts a comma separated sequence of words as input
and prints the words in a comma-separated sequence after sorting them alphabetically.

Suppose the following input is supplied to the program:

input: without,hello,bag,world
output: bag,hello,without,world*/

package exc8

import (
	"fmt"
	"sort"
	"strings"
)

func Getinput() (input string) {
	fmt.Print("Enter the input:")
	fmt.Scan(&input)
	return input
}

func Exc8(input string) []string {
	words := strings.Split(input, ",")
	fmt.Println(len(words))
	sort.Strings(words)
	fmt.Println(words)
	return words

}
