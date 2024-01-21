package exc8

import (
	"reflect"
	"testing"
)

func Test_Exc8(t *testing.T) {
	input := "without,hello,bag,world"
	want := []string{"bag", "hello", "without", "world"}
	got := Exc8(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Exc7() = %v, want %v", got, want)
	}
}
