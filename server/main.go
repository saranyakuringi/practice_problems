package main

import (
	"fmt"
	"practice/server/prog"
)

func main() {
	fmt.Print("Enter the number of elements in x:")
	var x int
	fmt.Scan(&x)

	fmt.Print("Enter the number of elements in y:")
	var y int
	fmt.Scan(&y)

	// D := make([]int, count)
	// for i := 0; i < count; i++ {
	// 	fmt.Printf("Enter the elements %d: ", i+1)
	// 	var value int
	// 	fmt.Scan(&value)
	// 	D[i] = value
	// }
	// prog.Exc6(D)

	prog.Exc7(x, y)
}
