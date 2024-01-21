package exc7

import (
	"reflect"
	"testing"
)

func Test_Exc7(t *testing.T) {
	a := 3
	b := 5
	want := [][]int{{0, 0, 0, 0, 0}, {0, 1, 2, 3, 4}, {0, 2, 4, 6, 8}}
	got := Exc7(a, b)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Exc7() = %v, want %v", got, want)
	}
}
