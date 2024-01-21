/*Write a program that accepts sequence of lines as input and
prints the lines after making all characters in the sentence capitalized.

Suppose the following input is supplied to the program:*/

package exc9

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Getinput() (input []string) {
	fmt.Print("Enter the input string(type 'end' on a new line to finish):")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "end" {
			break
		}
		input = append(input, line)
	}
	return input
}

func Exc9(input []string) []string {
	var result []string
	for _, value := range input {
		result = append(result, strings.ToUpper(value))
	}
	fmt.Println("Original String:", input)
	fmt.Println("String in Uppercase:", result)

	return result

}
