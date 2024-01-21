package exc10

import (
	"reflect"
	"testing"
)

func Test_Exc10(t *testing.T) {
	x := 20
	y := 30
	want := []string{"20", "22", "24", "26", "28"}
	got := Exc10(x, y)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Exc10() = %v, want %v", got, want)
	}
}
