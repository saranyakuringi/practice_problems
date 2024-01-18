/*Exercise 001:

Write a program which will find all such numbers which are divisible by 7
but are not a multiple of 5, between 2000 and 3200 (both included).
The numbers obtained should be printed in a comma-separated sequence on a single line.
*/

package prog

import (
	"fmt"
	"strconv"
	"strings"
)

func Exc1(low, high int) string {
	var result []string
	for i := low; i <= high; i++ {
		if i%7 == 0 && i%5 != 0 {
			result = append(result, strconv.Itoa(i))
		}
	}
	results := strings.Join(result, ",")
	fmt.Println(results)
	return results
}
