package exc9

import (
	"reflect"
	"testing"
)

func Test_Exc9(t *testing.T) {
	want := []string{"HELLO WORLD", "PRACTICE MAKES PERFECT"}
	input := []string{"hello world", "practice makes perfect"}
	got := Exc9(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Exc9() = %v, want %v", got, want)
	}
}
