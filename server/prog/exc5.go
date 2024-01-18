/*
Create a separate file (module) which has at least two methods:

- ReadString: to read a string from console input
- PrintString: to return the string in upper case.

Also create a `main.go` file that acts as calling class.
*/

package prog

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Exc5_ReadString() string {
	fmt.Print("Enter the string:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func PrintString(input string) string {
	return strings.ToUpper(input)
}
