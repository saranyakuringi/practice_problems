package main

import (
	"practice/server/prog"
	"reflect"
	"testing"
)

func TestExc_03(t *testing.T) {
	want := map[int]int{1: 1, 2: 4, 3: 9, 4: 16, 5: 25, 6: 36, 7: 49, 8: 64}
	got := prog.Exc3(8)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Ex003() = %v, want %v", got, want)
	}
}
