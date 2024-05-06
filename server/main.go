package main

import (
	"fmt"
	"practice/server/Array"
)

func main() {
	array := []int{1, 2, 3, 4, 5, 6}
	index := int(8.0)
	fmt.Println(Array.Array_Indexing(array, index))
}
