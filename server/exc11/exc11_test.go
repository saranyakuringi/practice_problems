package exc11

import (
	"reflect"
	"testing"
)

func Test_Exc11(t *testing.T) {
	points := [][]int{{1, 2}, {3, -1}, {2, 1}, {2, 3}}
	vertex := []int{2, 3}
	var k float64 = 2

	want1 := []int{2, 1}
	var want2 float64 = 0
	got1, got2 := ClosestDistance(points, vertex, k)

	if !reflect.DeepEqual(got1, want1) || got2 != want2 {
		t.Errorf("Exc10() = %v, %v, want = %v, %v", got1, got2, want1, want2)
	}
}
