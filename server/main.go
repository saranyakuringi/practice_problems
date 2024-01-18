package main

import (
	"fmt"
	"practice/server/prog"
)

func main() {
	//prog.Exc1(2000, 3200)
	// var num int
	// fmt.Print("Enter the number:")
	// fmt.Scanln(&num)

	// //fmt.Printf("The Factorial of %d is:%d", num, prog.Exc2(num))
	// fmt.Printf("The square of numbers upto %d is: %d", num, prog.Exc3(num))
	// input := "12, 23, 34, 45, 56, 67, 78, 89, 90"
	// prog.Exc4(input)

	// input := prog.Exc5_ReadString()

	// print := prog.PrintString(input)

	input := prog.Exc5_ReadString()

	print := prog.PrintString(input)

	fmt.Println("Original string:", input)
	fmt.Println("Upper case string:", print)

}
