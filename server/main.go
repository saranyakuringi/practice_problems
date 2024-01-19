package main

import (
	"fmt"
	"practice/server/prog"
)

func main() {
	fmt.Print("Enter the number of elements:")
	var count int
	fmt.Scan(&count)

	D := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Printf("Enter the elements %d: ", i+1)
		var value int
		fmt.Scan(&value)
		D[i] = value
	}
	prog.Exc6(D)
}
