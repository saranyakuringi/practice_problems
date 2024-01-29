package exc16

//Create a function that takes a number as an argument, add one to the number, and return the result

func Add(a int) int {
	return a + 1
}

//Create a function that takes voltage and current and returns the calculated power.

func Power(voltage, current int) int {
	power := voltage * current
	return power
}

//Create a function that finds the maximum range of a triangle's third edge,
//where the side lengths are all integers.

func Triangle_thirdEdge(side1, side2 int) int {
	side3 := (side1 + side2) - 1
	return side3
}
