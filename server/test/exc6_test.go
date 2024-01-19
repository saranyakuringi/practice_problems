package main

import (
	"practice/server/prog"
	"reflect"
	"testing"
)

func TestExc_06(t *testing.T) {
	want := []int{18, 22, 24}
	input := []int{100, 150, 180}
	got := prog.Exc6(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Exc6() = %v, want %v", got, want)
	}
}
