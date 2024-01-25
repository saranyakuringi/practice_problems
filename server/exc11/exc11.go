//point [](x,y) points=[[1,2],[3,-1],[2,1],[2,3]]
//vertex=[2,3]
//int k=2
//distance the point closest to the vertex

package exc11

import (
	"fmt"
	"math"
)

func CalculateDistance(p1, p2 []int) float64 {
	dx := float64(p1[0] - p2[0])
	dy := float64(p1[1] - p2[1])
	return math.Sqrt(dx*dx + dy*dy)
}

func ClosestDistance(points [][]int, vertex []int, k float64) (point []int, distance float64) {
	//points := [][]int{{1, 2}, {3, -1}, {2, 1}, {2, 3}}
	//vertex := []int{2, 3}
	//var k float64 = 2
	//var distance float64
	var closestdistance [][]int
	//var point []int
	for _, point = range points {
		distance = CalculateDistance(vertex, point)

		if distance == k {
			closestdistance = append(closestdistance, point)
		}
	}

	fmt.Printf("Points with Distance :%.2f\n", k)
	for _, point = range closestdistance {
		fmt.Printf("Point:%v,Distance:%v", point, distance)
	}
	return point, distance
}
