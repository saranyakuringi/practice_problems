package exc16

//1. Create a function that takes a number as an argument, add one to the number, and return the result

func Add(a int) int {
	return a + 1
}

//2. Create a function that takes voltage and current and returns the calculated power.

func Power(voltage, current int) int {
	power := voltage * current
	return power
}

//3. Create a function that finds the maximum range of a triangle's third edge,
//where the side lengths are all integers.

func Triangle_thirdEdge(side1, side2 int) int {
	side3 := (side1 + side2) - 1
	return side3
}

//4. Write a function that takes an integer minutes and converts it to seconds.

func Min_to_Sec_converter(min int) int {
	return min * 60
}

//5. Create a function that takes length and width and finds the perimeter of a rectangle.

func Perimeter_of_Rectange(length, width int) int {
	return 2 * (length + width)
}

//6. Two numbers are passed as parameters.The first parameter divided by the second parameter
//will have a remainder, possibly zero. Return that value.

func Remainder(a, b int) int {
	return a % b
}

//7. Write a function that takes the base and height of a triangle and return its area.

func Area_of_the_triange(base, height int) float32 {
	Area := 0.5 * float32(base*height)
	return Area
}
