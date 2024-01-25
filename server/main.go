package main

import "practice/server/exc11"

func main() {
	points := [][]int{{1, 2}, {3, -1}, {2, 1}, {2, 3}}
	vertex := []int{2, 3}
	var k float64 = 2
	exc11.ClosestDistance(points, vertex, k)
}
