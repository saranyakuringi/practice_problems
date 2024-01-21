/*
Create a separate file (module) which has at least two methods:

- ReadString: to read a string from console input
- PrintString: to return the string in upper case.

Also create a `main.go` file that acts as calling class.
*/

package exc5

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Exc5_ReadString() string {
	fmt.Print("Enter the string:")
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	if scanner.Scan() {
		input = scanner.Text()
		fmt.Println("Original String:", input)
		fmt.Println("String in Uppercase:", PrintString(input))
		return input
	}

	return ""
}

func PrintString(input string) string {
	return strings.ToUpper(input)
}
