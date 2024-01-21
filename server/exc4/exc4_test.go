package exc4

import (
	"reflect"
	"testing"
)

func TestExc_04(t *testing.T) {
	input := "12, 23, 34, 45, 56, 67, 78, 89, 90"
	want := []int{12, 23, 34, 45, 56, 67, 78, 89, 90}
	got := Exc4(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Exc4() = %v, want %v", got, want)
	}
}
